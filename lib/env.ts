const projectName = "bq-pipeline";

export enum Environments {
  PROD = "prod",
  DEV = "dev",
  TEST = "test",
  V1DEV = "v1dev",
}

//TODO FIX
export interface EnvironmentVariables {
  projectId: string;
  db: {
    user: {
      datasetId: string;
      tables: {
        sample1: {
          tableId: string;
          schemaJsonPath: string;
        };
        sample2: {
          tableId: string;
          schemaJsonPath: string;
        };
      };
    };
  };
}

const EnvironmentVariablesSetting: {
  [key in Environments]: EnvironmentVariables;
} = {
  [Environments.PROD]: {
    projectId: "",
    db: {
      user: {
        datasetId: "",
        tables: {
          sample1: {
            tableId: "",
            schemaJsonPath: "",
          },
          sample2: {
            tableId: "",
            schemaJsonPath: "",
          },
        },
      },
    },
  },
  [Environments.DEV]: {
    projectId: "",
    db: {
      user: {
        datasetId: "",
        tables: {
          sample1: {
            tableId: "",
            schemaJsonPath: "",
          },
          sample2: {
            tableId: "",
            schemaJsonPath: "",
          },
        },
      },
    },
  },
  [Environments.V1DEV]: {
    projectId: "xxxxxxxxxxxxxxxxx",
    db: {
      user: {
        datasetId: "sample_terraform_dataset_1",
        tables: {
          sample1: {
            tableId: "sample_terraform_user_table_1",
            schemaJsonPath: "schema/sample_user_tables_schema.json",
          },
          sample2: {
            tableId: "sample_terraform_user_table_1_2",
            schemaJsonPath: "schema/sample_user_table_schema.json",
          },
        },
      },
    },
  },
  [Environments.TEST]: {
    projectId: "",
    db: {
      user: {
        datasetId: "",
        tables: {
          sample1: {
            tableId: "",
            schemaJsonPath: "",
          },
          sample2: {
            tableId: "",
            schemaJsonPath: "",
          },
        },
      },
    },
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
