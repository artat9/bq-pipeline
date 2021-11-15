import { Table } from '@aws-cdk/aws-dynamodb';
import { Topic } from '@aws-cdk/aws-sns';
import { Construct, Stack, StackProps } from '@aws-cdk/core';
import * as environment from './env';

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
    new Topic(this, 'application-created', {
      topicName: applicationCreatedTopicName(target),
    });
  }
}

export const applicationCreatedTopicName = (
  target: environment.Environments
) => {
  return environment.withEnvPrefix(target, 'media-application-created');
};
