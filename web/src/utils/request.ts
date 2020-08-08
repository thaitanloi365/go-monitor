import { message } from 'antd';
import axios, { AxiosRequestConfig } from 'axios';
import { cloneDeep } from 'lodash';
import { compile, parse } from 'path-to-regexp';
import { CANCEL_REQUEST_MESSAGE } from 'utils/constant';

const { CancelToken } = axios;
// @ts-ignore
window.cancelRequest = new Map();

export default function request(options: AxiosRequestConfig) {
  let { data, url } = options;
  const cloneData = cloneDeep(data);

  try {
    let domain = '';
    const urlMatch = url.match(/[a-zA-z]+:\/\/[^/]*/);
    if (urlMatch) {
      [domain] = urlMatch;
      url = url.slice(domain.length);
    }

    const match = parse(url);
    url = compile(url)(data);

    for (const item of match) {
      if (item instanceof Object && item.name in cloneData) {
        delete cloneData[item.name];
      }
    }
    url = domain + url;
  } catch (e) {
    message.error(e.message);
  }

  options.url = url;
  options.params = cloneData;
  options.cancelToken = new CancelToken((cancel) => {
    // @ts-ignore
    window.cancelRequest.set(Symbol(Date.now()), {
      pathname: window.location.pathname,
      cancel,
    });
  });

  return axios(options)
    .then((response) => {
      const { statusText, status, data } = response;

      let result = {};
      if (typeof data === 'object') {
        result = data;
        if (Array.isArray(data)) {
          // @ts-ignore
          result.list = data;
        }
      } else {
        // @ts-ignore
        result.data = data;
      }

      return Promise.resolve({
        success: true,
        message: statusText,
        statusCode: status,
        ...result,
      });
    })
    .catch((error) => {
      const { response, message } = error;

      if (String(message) === CANCEL_REQUEST_MESSAGE) {
        return {
          success: false,
        };
      }

      let msg;
      let statusCode;

      if (response && response instanceof Object) {
        console.log(response);
        const { data, statusText } = response;
        statusCode = response.status;
        msg = data.message || statusText;
      } else {
        statusCode = 600;
        msg = error.message || 'Network Error';
      }

      /* eslint-disable */
      return Promise.reject({
        success: false,
        statusCode,
        message: msg,
      });
    });
}
