import * as cdktf from "cdktf";
import { Environments, valueOf } from "../env";
import { BigqueryStack } from "./bigquery";

const target: Environments = process.env.TARGET as Environments;
if (!target || !valueOf(target)) throw new Error("Invalid target environment");

const app = new cdktf.App();

//TODO use target
new BigqueryStack(app, "terraform-cdk");
app.synth();
