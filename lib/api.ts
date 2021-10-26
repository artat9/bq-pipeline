import {
  AuthorizationType,
  DynamoDbDataSource,
  FieldLogLevel,
  GraphqlApi,
  GraphqlApiProps,
  LambdaDataSource,
  MappingTemplate,
  Schema,
} from "@aws-cdk/aws-appsync";
import * as dynamodb from "@aws-cdk/aws-dynamodb";
import * as cdk from "@aws-cdk/core";
import { join } from "path";
import * as environment from "./env";
import { lambdaFunction } from "./function";
import IAM = require("@aws-cdk/aws-iam");
interface ApiProps extends cdk.StackProps {
  ddb: dynamodb.Table;
}

export class ApiStack extends cdk.Stack {
  public readonly appSyncLog: string;
  constructor(
    scope: cdk.Construct,
    id: string,
    target: environment.Environments,
    props: ApiProps
  ) {
    super(scope, id, props);
    const api = new GraphqlApi(this, "GraphqlApi", apiProps(target));
    withDDBResolvers(api, dDBDataSource(api, target, props));
    api.createResolver({
      typeName: "Mutation",
      fieldName: "uploadImage",
      dataSource: new LambdaDataSource(this, "uploadimage", {
        api: api,
        lambdaFunction: lambdaFunction(this, "issuepresign", target),
      }),
    });
    api.createResolver({
      typeName: "Mutation",
      fieldName: "applyForMediaAccount",
      dataSource: new LambdaDataSource(this, "createmediaaccount", {
        api: api,
        lambdaFunction: lambdaFunction(this, "mediaaccount", target),
      }),
    });
  }
}

const withDDBResolvers = (
  api: GraphqlApi,
  ddbDataSource: DynamoDbDataSource
) => {
  api.createResolver({
    typeName: "Query",
    fieldName: "getAccount",
    dataSource: ddbDataSource,
    requestMappingTemplate: MappingTemplate.fromFile(
      join(
        __dirname,
        "schema",
        "resolvers",
        "public",
        "getAccount" + ".request.vtl"
      )
    ),
    responseMappingTemplate: MappingTemplate.fromFile(
      join(
        __dirname,
        "schema",
        "resolvers",
        "public",
        "getAccount" + ".response.vtl"
      )
    ),
  });
};

const dDBDataSource = (
  api: GraphqlApi,
  target: environment.Environments,
  props: ApiProps
) => {
  return new DynamoDbDataSource(api, "main", {
    api: api,
    table: props.ddb,
    name: "ddb",
    readOnlyAccess: true,
  });
};

class QueryResolverInput {
  stack: ApiStack;
  resolverNames: string[];
  target: environment.Environments;
  functionRole: IAM.PolicyStatement;
  constructor(
    stack: ApiStack,
    resolverNames: string[],
    target: environment.Environments,
    functionRole: IAM.PolicyStatement
  ) {
    this.stack = stack;
    this.resolverNames = resolverNames;
    this.target = target;
    this.functionRole = functionRole;
  }
}

const apiDefaultRole = (
  api: GraphqlApi,
  target: environment.Environments,
  props: ApiProps
) => {
  const serviceRole = new IAM.Role(api, "appsyncServiceRole", {
    assumedBy: new IAM.ServicePrincipal("appsync.amazonaws.com"),
    path: "/service-role/",
    roleName: environment.withEnvPrefix(target, "appsyncDefaultRole"),
  });
  return serviceRole;
};

const withLambdaQueryResolvers = (
  api: GraphqlApi,
  input: QueryResolverInput,
  target: environment.Environments
) => {
  input.resolverNames.forEach((n) => {
    const f = lambdaFunction(input.stack, n, target);
    f.addToRolePolicy(input.functionRole);
    api.addLambdaDataSource(n, f).createResolver(queryResolver(n));
  });
};

const resolver = (typeName: string, fieldName: string) => {
  return {
    typeName: typeName,
    fieldName: fieldName,
    responseMappingTemplate: MappingTemplate.lambdaResult(),
  };
};

const queryResolver = (fieldName: string) => {
  return resolver("Query", fieldName);
};

const apiProps = (target: environment.Environments): GraphqlApiProps => {
  return {
    xrayEnabled: true,
    name: environment.withEnvPrefix(target, "GraphqlApi"),
    logConfig: {
      fieldLogLevel: FieldLogLevel.ALL,
      excludeVerboseContent: false,
    },
    schema: Schema.fromAsset(join(__dirname, "schema/schema.graphql")),
    authorizationConfig: {
      defaultAuthorization: {
        authorizationType: AuthorizationType.API_KEY,
        apiKeyConfig: {
          expires: cdk.Expiration.after(cdk.Duration.days(365)),
          description: "default",
        },
      },
    },
  };
};
