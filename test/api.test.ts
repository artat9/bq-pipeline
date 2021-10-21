import { SynthUtils } from "@aws-cdk/assert";
import { App } from "@aws-cdk/core";
import { Environments } from "../lib/env";
import { ApiStack } from "./../lib/api";
import { DataSourceStack } from "./../lib/datasource";

test("resource created", () => {
  const app = new App();
  const target = Environments.TEST;
  const ds = new DataSourceStack(app, "ds", target);
  const stack = new ApiStack(app, "test", target, {
    ddb: ds.ddb,
  });
  expect(SynthUtils.toCloudFormation(stack)).toMatchSnapshot();
});
