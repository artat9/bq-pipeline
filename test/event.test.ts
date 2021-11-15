import { SynthUtils } from '@aws-cdk/assert';
import { App } from '@aws-cdk/core';
import { Environments } from '../lib/env';
import { EventStack } from '../lib/event';

test('resource created', () => {
  const app = new App();
  const target = Environments.TEST;
  const stack = new EventStack(app, 'test', target);
  expect(SynthUtils.toCloudFormation(stack)).toMatchSnapshot();
});
