import React, { Component, ComponentProps } from 'react';
import classnames from 'classnames';
import styles from './index.less';
import Loader from 'components/Loader';

export interface IPageProps extends ComponentProps<any> {
  loading?: boolean;
  inner?: boolean;
}
export default class Page extends Component<IPageProps, {}> {
  render() {
    const { className, children, loading = false, inner = false } = this.props;
    const loadingStyle = {
      height: 'calc(100vh - 184px)',
      overflow: 'hidden',
    };
    return (
      <div
        className={classnames(className, {
          [styles.contentInner]: inner,
        })}
        style={loading ? loadingStyle : null}
      >
        {loading ? <Loader spinning /> : ''}
        {children}
      </div>
    );
  }
}
