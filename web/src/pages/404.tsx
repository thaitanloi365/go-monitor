import React from 'react';
import { FrownOutlined } from '@ant-design/icons';
import styles from './404.less';
import { Page } from 'components';

const Error = () => (
  <Page inner>
    <div className={styles.error}>
      <FrownOutlined />
      <h1>404 Not Found</h1>
    </div>
  </Page>
);

export default Error;
