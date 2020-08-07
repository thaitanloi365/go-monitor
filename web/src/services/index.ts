import {Method} from 'axios';
import config from 'utils/config';
import request from 'utils/request';

import api from './api';

export type Response = Promise <|{
  success: boolean;
  message: string;
  statusCode: number;
}
|{
  success: boolean;
}
> ;

const gen = (params: string) => {
  let url = config.apiPrefix + params;
  let method: any = 'GET';

  const paramsArray = params.split(' ');
  if (paramsArray.length === 2) {
    method = paramsArray[0];
    url = config.apiPrefix + paramsArray[1];
  }

  return async function(data: any) {
    try {
      const resp =
          await request({url, data, baseURL: config.apiBaseURL, method});
      console.log('**** resp', resp)
      return resp
    } catch (error) {
      console.log('**** errr', error)
      return error
    }
  };
};

type APIMap = {
  [key in keyof typeof api]: Response
};

const APIFunction = {};
for (const key in api) {
  APIFunction[key] = gen(api[key]);
}

export default APIFunction as APIMap;
