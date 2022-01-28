const projectName = 'bq-pipeline';

export enum Environments {
  PROD = 'prod',
  DEV = 'dev',
  TEST = 'test',
  V1DEV = 'v1dev',
}
export interface EnvironmentVariables {}

const EnvironmentVariablesSetting: {
  [key in Environments]: EnvironmentVariables;
} = {
  [Environments.PROD]: {},
  [Environments.DEV]: {},
  [Environments.V1DEV]: {},
  [Environments.TEST]: {},
};

export function valueOf(env: Environments): EnvironmentVariables {
  return EnvironmentVariablesSetting[env];
}
export const withEnvPrefix = (target: Environments, str: string) =>
  `${projectName}-${str}-${target}`;

export const isProd = (target: Environments) => {
  return target === Environments.PROD;
};
export const isDev = (target: Environments) => {
  return target === Environments.DEV;
};
