import { Reducer } from 'react';
import { Effect, Subscription } from 'dva';
import { RouteComponentProps } from 'react-router';

export type ITheme = 'dark' | 'light';
export type IEffect = Effect;
export type IReducer<T> = Reducer<T, any>;
export type ISubscription = Subscription;

export interface IModel<T> {
  namespace?: string;
  state?: T;
  effects?: { [key: string]: Effect };
  reducers?: { [key: string]: Reducer<T> };
  subscriptions?: { [key: string]: Subscription };
}

export interface ILoadingState {
  global: boolean;
  effects: { [key: string]: boolean | undefined };
  models: {
    global?: boolean;
    menu?: boolean;
    setting?: boolean;
    user?: boolean;
    login?: boolean;
  };
}

export interface IConnectState extends Partial<RouteComponentProps> {
  loading?: ILoadingState;
  app?: any;

  // Fix typing check
  dispatch?<K = any>(action: AnyAction): K;
}
