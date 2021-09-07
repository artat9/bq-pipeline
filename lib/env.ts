const projectName = "kaleido-backend";

export enum Environments {
  PROD = "prod",
  DEV = "dev",
  TEST = "test",
}
export interface EnvironmentVariables {
  rootDomain: string;
  hostedZoneId: string;
  txtRecords?: string[];
  allowedOrigin: string;
  deployBranchName: string;
  certificateArn: string;
}

const EnvironmentVariablesSetting: {
  [key in Environments]: EnvironmentVariables;
} = {
  [Environments.PROD]: {
    certificateArn:
      "arn:aws:acm:us-east-1:495476032358:certificate/98c65aba-1677-404e-9bac-1b2142602ad2",
    rootDomain: "kaleidodao.org",
    hostedZoneId: "Z07302833NYNT52C5SA30",
    allowedOrigin: "https://kaleidodao.org",
    deployBranchName: "main",
  },
  [Environments.DEV]: {
    certificateArn:
      "arn:aws:acm:us-east-1:495476032358:certificate/c635e3e6-7cc5-4777-b0a9-57317a70e9c2",
    rootDomain: "kaleido-dev.tk",
    hostedZoneId: "Z04820191RSTGGR5Z7MK0",
    allowedOrigin: "https://kaleido-webfront-git-dev-squard.vercel.app",
    deployBranchName: "dev",
  },
  [Environments.TEST]: {
    certificateArn: "",
    rootDomain: "test.example.com",
    hostedZoneId: "XXX",
    allowedOrigin: "http://localhost:3001",
    deployBranchName: "test",
  },
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
