import * as cdk from '@aws-cdk/core';
import { ApiStack } from '../lib/api';
import * as environment from '../lib/env';
const app = new cdk.App();
const target: environment.Environments = app.node.tryGetContext(
  'target'
) as environment.Environments;
if (!target || !environment.valueOf(target))
  throw new Error('Invalid target environment');

new ApiStack(app, environment.withEnvPrefix(target, 'graphqlapi'), target, {});
