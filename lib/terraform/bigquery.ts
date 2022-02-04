import {
  BigqueryDataset,
  BigqueryDatasetConfig,
  BigqueryTable,
  BigqueryTableConfig,
  GoogleProvider,
} from "@cdktf/provider-google";
import { TerraformStack } from "cdktf";
import { Construct } from "constructs";
import * as fs from "fs";
import * as path from "path";
import * as environment from "../env";
import { environmentParameters } from "../function";

export class BigqueryStack extends TerraformStack {
  constructor(
    scope: Construct,
    name: string
    //target: environment.Environments
  ) {
    super(scope, name);

    // TODO FIX Path(Github Action?)
    const credentialsPath = "/Users/oyanagijun/.gcp/toranomon/account.json";
    console.log(credentialsPath);
    const credentials = fs.existsSync(credentialsPath)
      ? fs.readFileSync(credentialsPath).toString()
      : "{}";

    console.log(credentials);

    //TODO Fix solid projectId writing
    const projectId = "os-tmp";

    const datasetId = "sample_terraform_dataset_2";
    const tableId = "sample_terraform_user_table_2";

    const labels = environmentParameters(environment.Environments.DEV);

    const schemaPath = path.join(
      process.cwd(),
      "schema/sample_user_tables_schema.json"
    );
    const schema = fs.existsSync(schemaPath)
      ? fs.readFileSync(credentialsPath).toString()
      : "{}";
    const bigqueryDatasetConfig: BigqueryDatasetConfig = {
      datasetId: datasetId,
      labels: labels,
    };
    const bigqueryTableConfig: BigqueryTableConfig = {
      datasetId: datasetId,
      tableId: tableId,
      labels: labels,
      schema: schema,
    };

    new GoogleProvider(this, "Google", {
      region: "us-central1",
      zone: "us-central1-a",
      project: projectId,
      credentials,
    });

    new BigqueryDataset(this, projectId, bigqueryDatasetConfig);
    new BigqueryTable(this, tableId, bigqueryTableConfig);
  }
}
