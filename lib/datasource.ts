import * as dynamodb from "@aws-cdk/aws-dynamodb";
import {
  Domain,
  DomainProps,
  ElasticsearchVersion,
} from "@aws-cdk/aws-elasticsearch";
import {
  AccountPrincipal,
  Effect,
  PolicyStatement,
  ServicePrincipal,
} from "@aws-cdk/aws-iam";
import { LogGroup, RetentionDays } from "@aws-cdk/aws-logs";
import * as cdk from "@aws-cdk/core";
import * as environment from "./env";

export class DataSourceStack extends cdk.Stack {
  public readonly es: Domain;
  public readonly ddb: dynamodb.Table;
  constructor(
    scope: cdk.Construct,
    id: string,
    target: environment.Environments
  ) {
    super(scope, id);
    const tableName = environment.withEnvPrefix(target, "main");
    //this.es = new Domain(this, "Domain", domainProps(this, target));
    this.ddb = new dynamodb.Table(this, tableName, ddbProps(tableName, target));
    withIndexes(this.ddb);
  }
}

const withIndexes = (table: dynamodb.Table) => {
  table.addGlobalSecondaryIndex({
    indexName: "GSI-1",
    partitionKey: {
      name: "SK",
      type: dynamodb.AttributeType.STRING,
    },
    sortKey: {
      name: "Data",
      type: dynamodb.AttributeType.STRING,
    },
  });
  table.addGlobalSecondaryIndex({
    indexName: "GSI-2",
    partitionKey: {
      name: "PK",
      type: dynamodb.AttributeType.STRING,
    },
    sortKey: {
      name: "Number",
      type: dynamodb.AttributeType.NUMBER,
    },
  });
};

const ddbProps = (
  tableName: string,
  target: environment.Environments
): dynamodb.TableProps => {
  return {
    tableName: tableName,
    partitionKey: {
      name: `PK`,
      type: dynamodb.AttributeType.STRING,
    },
    sortKey: {
      name: `SK`,
      type: dynamodb.AttributeType.STRING,
    },
    billingMode: dynamodb.BillingMode.PAY_PER_REQUEST,
    removalPolicy: cdk.RemovalPolicy.RETAIN,
    pointInTimeRecovery: true,
  };
};

const domainProps = (
  scope: cdk.Construct,
  target: environment.Environments
): DomainProps => {
  const domainName = environment.withEnvPrefix(target, "data");
  const appLogGroup = esLogGroup(scope, target, "app");
  const auditLogGroup = esLogGroup(scope, target, "audit");
  const slowIndexLogGroup = esLogGroup(scope, target, "slowIndex");
  const slowSearchLogGroup = esLogGroup(scope, target, "slowSearch");

  return {
    domainName: domainName,
    version: ElasticsearchVersion.V7_9,
    capacity: {
      dataNodeInstanceType: esInstanceType(target),
      masterNodes: 3,
      masterNodeInstanceType: esInstanceType(target),
      dataNodes: 3,
    },
    removalPolicy: cdk.RemovalPolicy.RETAIN,

    zoneAwareness: {
      availabilityZoneCount: 3,
      enabled: true,
    },
    enforceHttps: true,
    nodeToNodeEncryption: true,
    encryptionAtRest: {
      enabled: true,
    },

    accessPolicies: [
      allowMyAccountToAccessEs(scope, domainName),
      allowEsToWriteLogToCWLogs(scope, [
        appLogGroup,
        auditLogGroup,
        slowIndexLogGroup,
        slowSearchLogGroup,
      ]),
    ],

    logging: {
      appLogEnabled: true,
      appLogGroup: appLogGroup,
      //      auditLogEnabled: true,
      //      auditLogGroup: auditLogGroup,
      slowIndexLogEnabled: true,
      slowIndexLogGroup: slowIndexLogGroup,
      slowSearchLogEnabled: true,
      slowSearchLogGroup: slowSearchLogGroup,
    },
    automatedSnapshotStartHour: 0,
  };
};

const esInstanceType = (target: environment.Environments) => {
  if (environment.isProd(target)) {
    return "t3.medium.elasticsearch";
  }
  return "t3.small.elasticsearch";
};

const allowMyAccountToAccessEs = (
  scope: cdk.Construct,
  domainName: string
): PolicyStatement => {
  return new PolicyStatement({
    effect: Effect.ALLOW,
    actions: ["es:*"],
    principals: [new AccountPrincipal(cdk.Stack.of(scope).account)],
    resources: [
      `arn:aws:es:${cdk.Stack.of(scope).region}:${
        cdk.Stack.of(scope).account
      }:domain/${domainName}/*`,
    ],
  });
};

const allowEsToWriteLogToCWLogs = (
  scope: cdk.Construct,
  logGroups: LogGroup[]
): PolicyStatement => {
  return new PolicyStatement({
    effect: Effect.ALLOW,
    actions: ["logs:PutLogEvents", "logs:CreateLogStream"],
    principals: [new ServicePrincipal("es.amazonaws.com")],
    resources: logGroups.map((lg) => lg.logGroupArn),
  });
};

const esLogGroup = (
  scope: cdk.Construct,
  target: environment.Environments,
  logName: string
): LogGroup => {
  return new LogGroup(scope, `es${logName}`, {
    retention: RetentionDays.SIX_MONTHS,
    removalPolicy: cdk.RemovalPolicy.DESTROY,
    logGroupName: `/aws/es/${target.toString()}/${logName}`,
  });
};
