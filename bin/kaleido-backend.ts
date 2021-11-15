import * as cdk from '@aws-cdk/core';
import { CertificateStack } from '../lib/certificate';
import * as environment from '../lib/env';
import { RestApiStack } from '../lib/restapi';
import { Route53Stack } from '../lib/route53';
import { ApiStack } from './../lib/api';
import { AsstesStack } from './../lib/asset';
import { DataSourceStack } from './../lib/datasource';
import { EventStack } from './../lib/event';

const app = new cdk.App();
const target: environment.Environments = app.node.tryGetContext(
  'target'
) as environment.Environments;
if (!target || !environment.valueOf(target))
  throw new Error('Invalid target environment');

const route53 = new Route53Stack(
  app,
  environment.withEnvPrefix(target, 'route53'),
  target
);
const cert = new CertificateStack(
  app,
  environment.withEnvPrefix(target, `certificate-additional`),
  target
);
new AsstesStack(app, environment.withEnvPrefix(target, 'asset'), target, {
  certificate: cert.certificate,
  hostedZone: route53.hostedZone,
});
new RestApiStack(app, environment.withEnvPrefix(target, 'restapi'), target);
const ds = new DataSourceStack(
  app,
  environment.withEnvPrefix(target, 'datasource'),
  target
);
new ApiStack(app, environment.withEnvPrefix(target, 'graphqlapi'), target, {
  ddb: ds.ddb,
});
new EventStack(app, environment.withEnvPrefix(target, 'event'), target);
