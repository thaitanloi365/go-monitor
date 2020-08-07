import { IConfig } from 'umi-types';

// ref: https://umijs.org/config/
const config: IConfig = {
  define: {
    ENV: 'production',
    VERSION: 'v1.0.0',
    API_BASE_URL: 'https://api.sipscience.com',
  },
};

export default config;
