import type { Config } from "@jest/types";
import { defaults } from "jest-config";

export default async (): Promise<Config.InitialOptions> => {
  return {
    preset: "ts-jest",
    testEnvironment: "node",
    verbose: true,
    moduleFileExtensions: [...defaults.moduleFileExtensions, ".ts", ".tsx", ".js"],
    moduleNameMapper: {
      // "^image![a-zA-Z0-9$_-]+$": "GlobalImageStub",
      // "^[./a-zA-Z0-9$_-]+\\.png$": "<rootDir>/RelativeImageStub.js",
      // "module_name_(.*)": "<rootDir>/substituted_module_$1.js",
      // "assets/(.*)": ["<rootDir>/images/$1", "<rootDir>/photos/$1", "<rootDir>/recipes/$1"],

    },
  };
};
