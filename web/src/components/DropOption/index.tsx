import React from 'react';
import { BarsOutlined, DownOutlined } from '@ant-design/icons';
import { Dropdown, Button, Menu } from 'antd';
import { DropDownProps } from 'antd/lib/dropdown';

export interface IDropOptionProps extends DropDownProps {
  onMenuClick: (item: any) => void;
  menuOptions: Array<{ key: string; name: string }>;
  buttonStyle: React.CSSProperties;
  dropdownProps: any;
}
const DropOption: React.SFC<IDropOptionProps> = (props) => {
  const { onMenuClick, menuOptions = [], buttonStyle, dropdownProps } = props;
  const menu = menuOptions.map((item) => <Menu.Item key={item.key}>{item.name}</Menu.Item>);
  return (
    <Dropdown overlay={<Menu onClick={onMenuClick}>{menu}</Menu>} {...dropdownProps}>
      <Button style={{ border: 'none', ...buttonStyle }}>
        <BarsOutlined style={{ marginRight: 2 }} />
        <DownOutlined />
      </Button>
    </Dropdown>
  );
};

export default DropOption;
