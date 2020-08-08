import _ from 'lodash';
declare module '*.css';
declare module '*.png';
declare module '*.less';
declare module 'enquire-js';
declare module 'dva-model-extend';
declare var BUILD_ENV: string;
declare var API_BASE_URL: string;
declare var VERSION: string;
declare global {
  const _: typeof _;
}
