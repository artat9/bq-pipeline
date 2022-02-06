import {
  BigqueryDataset,
  BigqueryDatasetConfig,
  BigqueryTable,
  BigqueryTableConfig,
  GoogleProvider,
  ProjectIamBinding,
  ServiceAccount,
} from "@cdktf/provider-google";
import { TerraformStack } from "cdktf";
import { Construct } from "constructs";
import * as fs from "fs";
import * as path from "path";
import * as environment from "../env";
import { credentials, environmentParameters } from "../function";

export class BigqueryStack extends TerraformStack {
  constructor(
    scope: Construct,
    name: string,
    target: environment.Environments
  ) {
    super(scope, name);

    // TODO FIX Path(Github Action?)
    const credentialsPath = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx";
    const gcpCredentials = credentials(credentialsPath);
    const envVariables = environment.valueOf(target);
    const projectId = envVariables.projectId;

    new GoogleProvider(this, "Google", {
      region: "us-central1",
      zone: "us-central1-a",
      project: projectId,
      credentials: gcpCredentials,
    });

    const datasetId = envVariables.db.user.datasetId;
    const labels = environmentParameters(target, {}, true);
    const bigqueryDatasetConfig: BigqueryDatasetConfig = {
      datasetId: datasetId,
      labels: labels,
    };
    const dataset = new BigqueryDataset(this, datasetId, bigqueryDatasetConfig);

    const tableId = envVariables.db.user.tables.sample1.tableId;
    const schemaPath = path.join(
      process.cwd(),
      envVariables.db.user.tables.sample1.schemaJsonPath
    );
    const schema = fs.existsSync(schemaPath)
      ? fs.readFileSync(schemaPath).toString()
      : "{}";
    const bigqueryTableConfig: BigqueryTableConfig = {
      datasetId: dataset.datasetId,
      tableId: tableId,
      labels: labels,
      schema: schema,
      deletionProtection: false,
    };
    new BigqueryTable(this, tableId, bigqueryTableConfig);

    const serviceAccount = new ServiceAccount(
      this,
      "bq-pipelinee-service-account",
      {
        accountId: "bq-pipeline-member",
        displayName: "bq-pipeline-member",
      }
    );

    console.log(serviceAccount.accountId);

    new ProjectIamBinding(this, "jobUserBinding", {
      members: [`serviceAccount:${serviceAccount.email}`],
      role: "roles/bigquery.jobUser",
    });
    new ProjectIamBinding(this, "dataEditorBinding", {
      members: [`serviceAccount:${serviceAccount.email}`],
      role: "roles/bigquery.dataEditor",
    });
  }
}
