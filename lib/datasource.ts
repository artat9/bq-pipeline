import * as dynamodb from "@aws-cdk/aws-dynamodb";
import * as cdk from "@aws-cdk/core";
import * as environment from "./env";

export class DataSourceStack extends cdk.Stack {
  public readonly ddb: dynamodb.Table;
  constructor(
    scope: cdk.Construct,
    id: string,
    target: environment.Environments
  ) {
    super(scope, id);
    const tableName = environment.withEnvPrefix(target, "main");
    this.ddb = new dynamodb.Table(this, tableName, ddbProps(tableName));
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

const ddbProps = (tableName: string): dynamodb.TableProps => {
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
