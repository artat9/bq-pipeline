import {
  ApiKeySourceType,
  DomainName,
  EndpointType,
  IResource,
  LambdaIntegration,
  MockIntegration,
  PassthroughBehavior,
  RestApi,
  SecurityPolicy,
} from "@aws-cdk/aws-apigateway";
import * as certificatemanager from "@aws-cdk/aws-certificatemanager";
import { Effect, PolicyStatement } from "@aws-cdk/aws-iam";
import { Code, Function, Runtime, Tracing } from "@aws-cdk/aws-lambda";
import * as route53 from "@aws-cdk/aws-route53";
import * as alias from "@aws-cdk/aws-route53-targets";
import { Bucket, IBucket } from "@aws-cdk/aws-s3";
import * as cdk from "@aws-cdk/core";
import { CfnOutput, Construct, Duration } from "@aws-cdk/core";
import { resolve } from "path";
import * as environment from "./env";
import { hostedZoneFromId } from "./route53";

export class RestApiStack extends cdk.Stack {
  constructor(
    scope: cdk.Construct,
    id: string,
    target: environment.Environments,
    props?: cdk.StackProps
  ) {
    super(scope, id, props);
    const bucket = new Bucket(this, "AssetsBucket", {
      bucketName: environment.withEnvPrefix(target, "assets"),
      blockPublicAccess: {
        blockPublicAcls: true,
        blockPublicPolicy: true,
        ignorePublicAcls: true,
        restrictPublicBuckets: true,
      },
    });

    // ApiGateway
    const api = new RestApi(this, "RestApi", {
      restApiName: environment.withEnvPrefix(target, "restApi"),
      apiKeySourceType: ApiKeySourceType.HEADER,
    });
    const apiKey = api.addApiKey("APIKey", {
      apiKeyName: environment.withEnvPrefix(target, "RestApiKey"),
    });
    api
      .addUsagePlan("UsagePlanForAPIKey", {
        apiKey: apiKey,
      })
      .addApiStage({
        stage: api.deploymentStage,
      });
    new CfnOutput(this, environment.withEnvPrefix(target, "RestAPIKey"), {
      value: apiKey.keyId,
    });
    nftFunctions(this, ["post", "ad"], target, bucket, api);
    const customDomain = withCustomDomain(this, api, target);
    const getad = api.root
      .resourceForPath("ad")
      .addResource("{account}")
      .addResource("{metadata}");
    getad.addMethod(
      "GET",
      new LambdaIntegration(lambdaFunction(this, "getad", target))
    );

    addCorsOptions(getad);
    const account = api.root.addResource("account");
    const sign = account.addResource("sign");
    sign.addMethod(
      "POST",
      new LambdaIntegration(lambdaFunction(this, "sign", target))
    );
    addCorsOptions(sign);
    const reflesh = account.addResource("reflesh");
    reflesh.addMethod(
      "POST",
      new LambdaIntegration(lambdaFunction(this, "reflesh", target))
    );
    addCorsOptions(reflesh);
    aRecord(this, target, customDomain);
  }
}

const lambdaFunction = (
  scope: Construct,
  id: string,
  target: environment.Environments
) => {
  const func = new Function(scope, id, {
    functionName: environment.withEnvPrefix(target, id),
    code: code(id),
    handler: "bin/main",
    timeout: Duration.minutes(1),
    runtime: Runtime.GO_1_X,
    tracing: Tracing.ACTIVE,
    environment: {
      EnvironmentId: target.toString(),
      AllowedOrigin: environment.valueOf(target).allowedOrigin,
    },
  });
  func.addToRolePolicy(
    new PolicyStatement({
      effect: Effect.ALLOW,
      actions: ["lambda:*"],
      resources: ["*"],
    })
  );
  func.addToRolePolicy(
    new PolicyStatement({
      effect: Effect.ALLOW,
      actions: ["ssm:Get*"],
      resources: ["*"],
    })
  );
  func.addToRolePolicy(
    new PolicyStatement({
      effect: Effect.ALLOW,
      actions: ["ssm:Get*"],
      resources: ["*"],
    })
  );
  func.addToRolePolicy(
    new PolicyStatement({
      effect: Effect.ALLOW,
      actions: ["dynamodb:*"],
      resources: ["*"],
    })
  );
  return func;
};

const nftFunctions = (
  scope: Construct,
  resources: string[],
  target: environment.Environments,
  assetBucket: IBucket,
  api: RestApi
) => {
  return resources.map((r) => {
    const resource = api.root.addResource(r);
    const func = lambdaFunction(scope, r, target);
    resource.addMethod("POST", new LambdaIntegration(func));
    addCorsOptions(resource);
    return resource;
  });
};

const withCustomDomain = (
  stack: cdk.Stack,
  api: RestApi,
  target: environment.Environments
) => {
  const customDomain = api.addDomainName(
    environment.withEnvPrefix(target, "domain"),
    customDomainProps(stack, target)
  );
  return customDomain;
};

const restApiDomainName = (target: environment.Environments) => {
  return `api.` + environment.valueOf(target).rootDomain;
};

const aRecord = (
  stack: RestApiStack,
  target: environment.Environments,
  customDomain: DomainName
) => {
  new route53.ARecord(stack, "RestApiARecord", {
    zone: hostedZoneFromId(stack, target),
    recordName: restApiDomainName(target),
    target: route53.RecordTarget.fromAlias(
      new alias.ApiGatewayDomain(customDomain)
    ),
  });
};

const customDomainProps = (
  stack: cdk.Stack,
  target: environment.Environments
) => {
  return {
    domainName: restApiDomainName(target),
    certificate: certificatemanager.Certificate.fromCertificateArn(
      stack,
      "Cert",
      environment.valueOf(target).certificateArn
    ),
    securityPolicy: SecurityPolicy.TLS_1_2,
    endpointType: EndpointType.REGIONAL,
  };
};

function addCorsOptions(apiResource: IResource) {
  apiResource.addMethod(
    "OPTIONS",
    new MockIntegration({
      integrationResponses: [
        {
          statusCode: "200",
          responseParameters: {
            "method.response.header.Access-Control-Allow-Headers":
              "'Content-Type,X-Amz-Date,Authorization,X-Api-Key,X-Amz-Security-Token,X-Amz-User-Agent'",
            "method.response.header.Access-Control-Allow-Origin": "'*'",
            "method.response.header.Access-Control-Allow-Credentials":
              "'false'",
            "method.response.header.Access-Control-Allow-Methods":
              "'OPTIONS,GET,PUT,POST,DELETE'",
          },
        },
      ],
      passthroughBehavior: PassthroughBehavior.NEVER,
      requestTemplates: {
        "application/json": '{"statusCode": 200}',
      },
    }),
    {
      methodResponses: [
        {
          statusCode: "200",
          responseParameters: {
            "method.response.header.Access-Control-Allow-Headers": true,
            "method.response.header.Access-Control-Allow-Methods": true,
            "method.response.header.Access-Control-Allow-Credentials": true,
            "method.response.header.Access-Control-Allow-Origin": true,
          },
        },
      ],
    }
  );
}
const code = (dirname: string) => {
  return Code.fromAsset(
    resolve(`${__dirname}/../`, "lib", "functions", dirname, "bin", "main.zip")
  );
};
