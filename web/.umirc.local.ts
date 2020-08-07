import { IConfig } from 'umi-types';

// ref: https://umijs.org/config/
const config: IConfig = {
  define: {
    ENV: 'local',
    VERSION: 'v1.0.0',
    API_BASE_URL: 'http://localhost:8080',
  },
};

export default config;
