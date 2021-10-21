import * as cdk from "@aws-cdk/core";
import { CertificateStack } from "../lib/certificate";
import * as environment from "../lib/env";
import { RestApiStack } from "../lib/restapi";
import { Route53Stack } from "../lib/route53";
import { ApiStack } from "./../lib/api";

const app = new cdk.App();
const target: environment.Environments = app.node.tryGetContext(
  "target"
) as environment.Environments;
if (!target || !environment.valueOf(target))
  throw new Error("Invalid target environment");

new Route53Stack(app, environment.withEnvPrefix(target, "route53"), target);
new CertificateStack(
  app,
  environment.withEnvPrefix(target, `certificate-additional`),
  target
);
new RestApiStack(app, environment.withEnvPrefix(target, "restapi"), target);
new ApiStack(app, environment.withEnvPrefix(target, "graphqlapi"), target);
