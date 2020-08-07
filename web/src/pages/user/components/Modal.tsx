import React, { PureComponent } from 'react';
import { Form, Input, InputNumber, Radio, Modal } from 'antd';

const FormItem = Form.Item;

const formItemLayout = {
  labelCol: {
    span: 6,
  },
  wrapperCol: {
    span: 14,
  },
};

class UserModal extends PureComponent {
  formRef = React.createRef();

  handleOk = () => {
    const { item = {}, onOk } = this.props;

    this.formRef.current
      .validateFields()
      .then((values) => {
        const data = {
          ...values,
          key: item.key,
        };
        data.address = data.address.join(' ');
        onOk(data);
      })
      .catch((errorInfo) => {
        console.log(errorInfo);
      });
  };

  render() {
    const { item = {}, onOk, form, ...modalProps } = this.props;

    return (
      <Modal {...modalProps} onOk={this.handleOk}>
        <Form
          ref={this.formRef}
          name="control-ref"
          initialValues={{
            ...item,
            address: item.address && item.address.split(' '),
          }}
          layout="horizontal"
        >
          <FormItem name="name" rules={[{ required: true }]} label="Name" hasFeedback {...formItemLayout}>
            <Input />
          </FormItem>
          <FormItem name="nickName" rules={[{ required: true }]} label="NickName" hasFeedback {...formItemLayout}>
            <Input />
          </FormItem>
          <FormItem name="isMale" rules={[{ required: true }]} label="Gender" hasFeedback {...formItemLayout}>
            <Radio.Group>
              <Radio value>Male</Radio>
              <Radio value={false}>Female</Radio>
            </Radio.Group>
          </FormItem>
          <FormItem name="age" label="Age" hasFeedback {...formItemLayout}>
            <InputNumber min={18} max={100} />
          </FormItem>
          <FormItem
            name="phone"
            rules={[
              {
                required: true,
                pattern: /^1[34578]\d{9}$/,
                message: 'The input is not valid phone!',
              },
            ]}
            label="Phone"
            hasFeedback
            {...formItemLayout}
          >
            <Input />
          </FormItem>
          <FormItem
            name="email"
            rules={[
              {
                required: true,
                pattern: /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/,
                message: 'The input is not valid E-mail!',
              },
            ]}
            label="Email"
            hasFeedback
            {...formItemLayout}
          >
            <Input />
          </FormItem>
        </Form>
      </Modal>
    );
  }
}

export default UserModal;
