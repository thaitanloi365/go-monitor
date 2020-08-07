import { IConfig } from 'umi-types';

// ref: https://umijs.org/config/
const config: IConfig = {
  define: {
    ENV: 'staging',
    VERSION: 'v1.0.0',
    API_BASE_URL: 'https://api-staging.sipscience.com',
  },
};

export default config;
