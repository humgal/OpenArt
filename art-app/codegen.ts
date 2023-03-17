
import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,
  schema: "http://localhost:8080/graphql",
  // documents: "src/**/*.tsx",
  generates: {
    "src/gql/": {
      preset: "client",
      plugins: ['typescript']
    },
    "./graphql/schema.json": {
      plugins: ["introspection"]
    }
  }
};

export default config;
