import { Table } from '@aws-cdk/aws-dynamodb';
import { IEventSource } from '@aws-cdk/aws-lambda';
import { SnsEventSource } from '@aws-cdk/aws-lambda-event-sources';
import { Topic } from '@aws-cdk/aws-sns';
import { Construct, Stack, StackProps } from '@aws-cdk/core';
import * as environment from './env';
import { lambdaFunction } from './function';

type EventProps = StackProps & {
  ddb: Table;
  esEndPoint: string;
  esArn: string;
  publicApiLogGroup: string;
};

export class EventStack extends Stack {
  constructor(
    scope: Construct,
    id: string,
    target: environment.Environments,
    props?: StackProps
  ) {
    super(scope, id, props);
    const applicationCreated = new Topic(this, 'application-created', {
      topicName: applicationCreatedTopicName(target),
    });
    subscribeFunctions(
      this,
      new SnsEventSource(applicationCreated),
      ['accountmail', 'notifytoslack'],
      target
    );
  }
}

const subscribeFunctions = (
  scope: Construct,
  source: IEventSource,
  functions: string[],
  target: environment.Environments
) => {
  functions.forEach((f) => {
    const func = lambdaFunction(scope, f, target);
    func.addEventSource(source);
  });
};

export const applicationCreatedTopicName = (
  target: environment.Environments
) => {
  return environment.withEnvPrefix(target, 'media-application-created');
};
