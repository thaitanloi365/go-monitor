import React, { PureComponent, Fragment } from 'react';
import { connect } from 'umi';
import { IConnectState } from 'types';
import { Button, Row, Input, Form } from 'antd';
import config from 'utils/config';
import styles from './index.less';

const FormItem = Form.Item;

@connect(({ loading, dispatch }: IConnectState) => ({ loading, dispatch }))
class Login extends PureComponent<IConnectState> {
  render() {
    const { dispatch, loading } = this.props;

    const handleOk = (values: any) => {
      dispatch({ type: 'login/login', payload: values });
    };

    return (
      <Fragment>
        <div className={styles.form}>
          <div className={styles.logo}>
            <img alt="logo" src={config.logoPath} />
            <span>{config.siteName}</span>
          </div>
          <Form onFinish={handleOk}>
            <FormItem name="username" rules={[{ required: true }]} hasFeedback>
              <Input placeholder={'Username'} />
            </FormItem>
            <FormItem name="password" rules={[{ required: true }]} hasFeedback>
              <Input type="password" placeholder={'Password'} />
            </FormItem>
            <Row>
              <Button type="primary" htmlType="submit" loading={loading.effects.login}>
                Sign in
              </Button>
              <p>
                <span className="margin-right">Username ：guest</span>
                <span>Password ：guest</span>
              </p>
            </Row>
          </Form>
        </div>
      </Fragment>
    );
  }
}

export default Login;
