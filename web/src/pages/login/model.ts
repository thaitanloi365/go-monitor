import store from 'store';
import { IEffect, IModel } from 'types';
import { history } from 'umi';
import { parseFromUrl } from 'utils';
import APIFunction from 'services';

export interface ILoginModelState {}
export interface ILoginModelType extends IModel<ILoginModelState> {
  effects: {
    login: IEffect;
  };
}

const Model: ILoginModelType = {
  namespace: 'login',

  state: {},

  effects: {
    *login({ payload }, { put, call, select }) {
      const { success, data } = yield call(APIFunction.login, payload);
      if (success && data) {
        store.set('token', data.token);
        store.set('user', data.user);

        var value = parseFromUrl(location?.search);
        console.log(value?.from);
        if (value?.from && value.from != 'login') {
          history.push(value?.from);
          return;
        }

        history.push('/dashboard');
      } else {
        throw data;
      }
    },
  },
};

export default Model;
