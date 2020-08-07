import React, { PureComponent } from 'react';
import { connect } from 'umi';
import { Row, Col } from 'antd';
import { Page } from 'components';
import styles from './index.less';
import store from 'store';

const bodyStyle = {
  bodyStyle: {
    height: 432,
    background: '#fff',
  },
};

@connect(({ app, dashboard, loading }) => ({
  dashboard,
  loading,
}))
class Dashboard extends PureComponent {
  render() {
    const userDetail = store.get('user');
    return (
      <Page
        // loading={loading.models.dashboard && sales.length === 0}
        className={styles.dashboard}
      >
        <Row gutter={24}></Row>
      </Page>
    );
  }
}

export default Dashboard;
