import { Effect, PolicyStatement } from '@aws-cdk/aws-iam';
import { Code, Function, Runtime, Tracing } from '@aws-cdk/aws-lambda';
import { Construct, Duration } from '@aws-cdk/core';
import { resolve } from 'path';
import * as environment from './env';
import { applicationCreatedTopicName } from './event';

export const lambdaFunction = (
  scope: Construct,
  id: string,
  target: environment.Environments,
  additionalEnvParams?: { [key: string]: string }
) => {
  const func = new Function(scope, id, {
    functionName: environment.withEnvPrefix(target, id),
    code: code(id),
    handler: 'bin/main',
    timeout: Duration.minutes(1),
    runtime: Runtime.GO_1_X,
    tracing: Tracing.ACTIVE,
    environment: environmentParameters(target, additionalEnvParams),
  });
  // TODO: restrict access
  func.addToRolePolicy(
    new PolicyStatement({
      effect: Effect.ALLOW,
      actions: ['lambda:*'],
      resources: ['*'],
    })
  );
  func.addToRolePolicy(
    new PolicyStatement({
      effect: Effect.ALLOW,
      actions: ['ssm:Get*'],
      resources: ['*'],
    })
  );
  func.addToRolePolicy(
    new PolicyStatement({
      effect: Effect.ALLOW,
      actions: ['dynamodb:*'],
      resources: ['*'],
    })
  );
  func.addToRolePolicy(
    new PolicyStatement({
      effect: Effect.ALLOW,
      actions: ['sns:Publish'],
      resources: ['*'],
    })
  );
  func.addToRolePolicy(
    new PolicyStatement({
      effect: Effect.ALLOW,
      actions: ['ses:SendEmail'],
      resources: ['*'],
    })
  );
  func.addToRolePolicy(
    new PolicyStatement({
      effect: Effect.ALLOW,
      actions: ['s3:*'],
      resources: ['*'],
    })
  );
  return func;
};
const code = (dirname: string) => {
  return Code.fromAsset(
    resolve(`${__dirname}/../`, 'cmd', dirname, 'bin', 'main.zip')
  );
};

const environmentParameters = (
  target: environment.Environments,
  additionalEnvParams?: {
    [key: string]: string;
  }
): { [key: string]: string } => {
  var stack: { [key: string]: string } = {
    EnvironmentId: target.toString(),
    AllowedOrigin: environment.valueOf(target).allowedOrigin,
    ApplicationCreatedTopicName: applicationCreatedTopicName(target),
    RootDomain: environment.valueOf(target).rootDomain,
  };
  if (!additionalEnvParams) {
    return stack;
  }
  Object.keys(additionalEnvParams).forEach((k) => {
    stack[k] = additionalEnvParams[k];
  });
  return stack;
};
