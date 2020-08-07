import React, { PureComponent, Fragment } from 'react';
import { connect } from 'umi';
import { Helmet } from 'react-helmet';
import { Loader } from 'components';
import { queryLayout } from 'utils';
import NProgress from 'nprogress';
import { withRouter } from 'umi';
import PublicLayout from './PublicLayout';
import PrimaryLayout from './PrimaryLayout';
import './BaseLayout.less';
import config from 'utils/config';
import { IConnectState } from 'types';

const LayoutMap = {
  primary: PrimaryLayout,
  public: PublicLayout,
};

@withRouter
@connect(({ loading }: IConnectState) => ({ loading }))
class BaseLayout extends PureComponent<IConnectState> {
  previousPath = '';

  render() {
    const { loading, children, location } = this.props;
    const Container = LayoutMap[queryLayout(config.layouts, location.pathname)];

    const currentPath = location.pathname + location.search;
    if (currentPath !== this.previousPath) {
      NProgress.start();
    }

    if (!loading.global) {
      NProgress.done();
      this.previousPath = currentPath;
    }

    return (
      <Fragment>
        <Helmet>
          <title>{config.siteName}</title>
        </Helmet>
        <Loader fullScreen spinning={loading.effects['app/query']} />
        <Container>{children}</Container>
      </Fragment>
    );
  }
}

export default BaseLayout;
