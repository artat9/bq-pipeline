import { Table } from '@aws-cdk/aws-dynamodb';
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
    const mailFunc = lambdaFunction(this, 'accountmail', target);
    mailFunc.addEventSource(new SnsEventSource(applicationCreated));
  }
}

export const applicationCreatedTopicName = (
  target: environment.Environments
) => {
  return environment.withEnvPrefix(target, 'media-application-created');
};
