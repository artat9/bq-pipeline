import { SynthUtils } from "@aws-cdk/assert";
import { App } from "@aws-cdk/core";
import { AsstesStack } from "../lib/asset";
import { Environments } from "../lib/env";
import { CertificateStack } from "./../lib/certificate";
import { Route53Stack } from "./../lib/route53";

test("resource created", () => {
  const app = new App();
  const target = Environments.TEST;
  const cert = new CertificateStack(app, "cert", target);
  const r53 = new Route53Stack(app, "r53", target);
  const stack = new AsstesStack(app, "test", target, {
    certificate: cert.certificate,
    hostedZone: r53.hostedZone,
  });
  expect(SynthUtils.toCloudFormation(stack)).toMatchSnapshot();
});
