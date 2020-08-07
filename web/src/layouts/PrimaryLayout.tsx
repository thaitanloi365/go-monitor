import React, { PureComponent, Fragment } from 'react';
import { withRouter, connect } from 'umi';
import { MyLayout, GlobalFooter } from 'components';
import { BackTop, Layout, Drawer } from 'antd';
import { enquireScreen, unenquireScreen } from 'enquire-js';
import Error from '../pages/404';
import styles from './PrimaryLayout.less';
import store from 'store';
import config from 'utils/config';
import constant from 'utils/constant';
import { IConnectState } from 'types';
import { IHeaderProps } from 'components/Layout/Header';
import { ISiderProps } from 'components/Layout/Sider';

const { Content } = Layout;
const { Header, Bread, Sider } = MyLayout;
@withRouter
@connect(({ app, loading }: IConnectState) => ({ app, loading }))
class PrimaryLayout extends PureComponent<IConnectState> {
  state = {
    isMobile: false,
  };
  enquireHandler: any;

  componentDidMount() {
    this.enquireHandler = enquireScreen((mobile: boolean) => {
      const { isMobile } = this.state;
      if (isMobile !== mobile) {
        this.setState({
          isMobile: mobile,
        });
      }
    });
  }

  componentWillUnmount() {
    unenquireScreen(this.enquireHandler);
  }

  onCollapseChange = (collapsed: boolean) => {
    this.props.dispatch({
      type: 'app/handleCollapseChange',
      payload: collapsed,
    });
  };

  renderFooterInfo() {
    return (
      <div className={styles.info}>
        <span>{config.copyright}</span>
      </div>
    );
  }

  render() {
    const { app, dispatch, children } = this.props;
    const { theme, collapsed, notifications } = app || {};
    const user = store.get('user') || {};
    const token = store.get('token') || '';
    const { isMobile } = this.state;
    const { onCollapseChange } = this;

    // Query whether you have permission to enter this page
    const hasPermission = token !== '';

    // MenuParentId is equal to -1 is not a available menu.
    const menus = constant.menus.filter((item) => item.menuParentId !== '-1');

    const headerProps: IHeaderProps = {
      menus,
      collapsed,
      notifications,
      onCollapseChange,
      fixed: config.fixedHeader,
      onAllNotificationsRead() {
        dispatch({ type: 'app/allNotificationsRead' });
      },
      onSignOut() {
        dispatch({ type: 'app/signOut' });
      },
      onEdit() {
        dispatch({ type: 'app/signOut' });
      },
      user,
    };

    const siderProps: ISiderProps = {
      theme,
      menus,
      isMobile,
      collapsed,
      onCollapseChange,
      onThemeChange(theme: 'light' | 'dark') {
        dispatch({
          type: 'app/handleThemeChange',
          payload: theme,
        });
      },
    };

    return (
      <Fragment>
        <Layout>
          {isMobile ? (
            <Drawer
              maskClosable={true}
              closable={false}
              onClose={onCollapseChange.bind(this, !collapsed)}
              visible={!collapsed}
              placement="left"
              width={200}
              style={{
                padding: 0,
                height: '100vh',
              }}
            >
              <Sider {...siderProps} collapsed={false} />
            </Drawer>
          ) : (
            <Sider {...siderProps} />
          )}
          <div className={styles.container} style={{ paddingTop: config.fixedHeader ? 72 : 0 }} id="primaryLayout">
            <Header {...headerProps} />
            <Content className={styles.content}>
              <Bread routeList={constant.menus} />
              {hasPermission ? children : <Error />}
            </Content>
            <BackTop
              className={styles.backTop}
              target={() => document.querySelector<HTMLHtmlElement>('#primaryLayout')}
            />
            <GlobalFooter className={styles.footer} copyright={this.renderFooterInfo()} />
          </div>
        </Layout>
      </Fragment>
    );
  }
}

export default PrimaryLayout;
