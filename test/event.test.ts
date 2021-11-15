import { SynthUtils } from '@aws-cdk/assert';
import { App } from '@aws-cdk/core';
import { DataSourceStack } from '../lib/datasource';
import { Environments } from '../lib/env';
import { EventStack } from '../lib/event';

test('resource created', () => {
  const app = new App();
  const target = Environments.TEST;
  const ds = new DataSourceStack(app, 'ds', target);
  const stack = new EventStack(app, 'test', target);
  expect(SynthUtils.toCloudFormation(stack)).toMatchSnapshot();
});
