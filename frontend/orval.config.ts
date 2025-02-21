import { defineConfig } from "orval";
import { loadEnvFile } from "process";

loadEnvFile(".env");

export default defineConfig({
  api: {
    output: {
      // mode: 'tags-split',
      target: "./src/api/generated/taskManagerApis.ts",
      schemas: "./src/api/models",
      client: "react-query",
      prettier: true,
      clean: true,
      override: {
        fetch: {
          includeHttpResponseReturnType: false,
        },
      },
      baseUrl: `${process.env.API_URL}`,
      allParamsOptional: false,
      unionAddMissingProperties: true,
    },
    input: {
      target: `${process.env.SWAGGER_URL}`,
      validation: false,
    },
  },
});
