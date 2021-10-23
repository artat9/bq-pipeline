import { ICertificate } from "@aws-cdk/aws-certificatemanager";
import {
  CloudFrontAllowedMethods,
  CloudFrontWebDistribution,
  OriginAccessIdentity,
  PriceClass,
  SSLMethod,
} from "@aws-cdk/aws-cloudfront";
import {
  CanonicalUserPrincipal,
  Effect,
  PolicyStatement,
} from "@aws-cdk/aws-iam";
import { S3EventSource } from "@aws-cdk/aws-lambda-event-sources";
import { ARecord, IHostedZone, RecordTarget } from "@aws-cdk/aws-route53";
import { CloudFrontTarget } from "@aws-cdk/aws-route53-targets";
import {
  BlockPublicAccess,
  Bucket,
  BucketAccessControl,
  EventType,
  HttpMethods,
} from "@aws-cdk/aws-s3";
import {
  Construct,
  Duration,
  RemovalPolicy,
  Stack,
  StackProps,
} from "@aws-cdk/core";
import * as environment from "./env";
import { lambdaFunction } from "./function";
type AsstesStackProps = StackProps & {
  certificate: ICertificate;
  hostedZone: IHostedZone;
};
export class AsstesStack extends Stack {
  public readonly assetBucket: Bucket;
  constructor(
    scope: Construct,
    id: string,
    target: environment.Environments,
    props: AsstesStackProps
  ) {
    super(scope, id);
    permanentBucket(this, target, this.account, props);
  }
}

export const assetBucketName = (target: environment.Environments) => {
  return environment.withEnvPrefix(target, "asset");
};

const permanentBucket = (
  scope: Construct,
  target: environment.Environments,
  account: string,
  props: AsstesStackProps
) => {
  const b = new Bucket(scope, "AssetBucket", {
    accessControl: BucketAccessControl.PRIVATE,
    blockPublicAccess: BlockPublicAccess.BLOCK_ALL,
    bucketName: assetBucketName(target),
    publicReadAccess: false,
    removalPolicy: RemovalPolicy.DESTROY,
    autoDeleteObjects: true,
    cors: cors(),
  });
  const oai = new OriginAccessIdentity(scope, "OAI-AssetsBucket");
  const myBucketPolicy = new PolicyStatement({
    effect: Effect.ALLOW,
    actions: ["s3:GetObject"],
    principals: [
      new CanonicalUserPrincipal(
        oai.cloudFrontOriginAccessIdentityS3CanonicalUserId
      ),
    ],
    resources: [`${b.bucketArn}/*`],
  });
  b.addToResourcePolicy(myBucketPolicy);
  const assetsDist = new CloudFrontWebDistribution(scope, "AssetDistribution", {
    viewerCertificate: {
      aliases: [assetDomainName(target)],
      props: {
        acmCertificateArn: props.certificate.certificateArn,
        minimumProtocolVersion: "TLSv1",
        sslSupportMethod: SSLMethod.SNI,
      },
    },
    priceClass: PriceClass.PRICE_CLASS_ALL,
    originConfigs: [
      {
        s3OriginSource: {
          s3BucketSource: b,
          originAccessIdentity: oai,
        },
        behaviors: [
          {
            isDefaultBehavior: true,
            allowedMethods: CloudFrontAllowedMethods.GET_HEAD_OPTIONS,
            minTtl: Duration.days(365),
            maxTtl: Duration.days(365),
            defaultTtl: Duration.days(365),
          },
        ],
      },
    ],
  });
  const invalidateFunc = lambdaFunction(scope, "updateimage", target, {
    AssetDistributionId: assetsDist.distributionId,
  });
  invalidateFunc.addToRolePolicy(
    new PolicyStatement({
      effect: Effect.ALLOW,
      actions: ["cloudfront:CreateInvalidation"],
      resources: [
        `arn:aws:cloudfront::${account}:distribution/${assetsDist.distributionId}`,
        `arn:aws:cloudfront::${account}:distribution/${assetsDist.distributionId}/*`,
      ],
    })
  );
  invalidateFunc.addEventSource(
    new S3EventSource(b, {
      events: [EventType.OBJECT_CREATED],
    })
  );
  new ARecord(scope, `AssetCDNArecord`, {
    recordName: assetDomainName(target),
    target: RecordTarget.fromAlias(new CloudFrontTarget(assetsDist)),
    zone: props.hostedZone,
  });
  return b;
};

const assetDomainName = (target: environment.Environments) => {
  const { rootDomain } = environment.valueOf(target);
  return `assets.${rootDomain}`;
};

const cors = () => {
  return [
    {
      allowedMethods: [HttpMethods.GET, HttpMethods.HEAD],
      allowedOrigins: ["*"],
      allowedHeaders: ["*"],
      maxAge: 3000,
      exposedHeaders: [
        "x-amz-server-side-encryption",
        "x-amz-request-id",
        "x-amz-id-2",
        "ETag",
      ],
    },
  ];
};
