import React, { PureComponent, Fragment } from 'react';
import { Icon as LegacyIcon } from '@ant-design/compatible';
import { Menu } from 'antd';
import { NavLink, withRouter } from 'umi';
import { arrayToTree, queryAncestors } from 'utils';
import { pathToRegexp } from 'path-to-regexp';
import store from 'store';
import { RouteComponentProps } from 'types/umi';

const { SubMenu } = Menu;

interface ISiderMenuProps extends RouteComponentProps {
  menus: any[];
  theme: 'light' | 'dark';
  isMobile: boolean;
  onCollapseChange: (collapse: boolean) => void;
  collapsed: boolean;
}

@withRouter
class SiderMenu extends PureComponent<ISiderMenuProps> {
  state = {
    openKeys: store.get('openKeys') || [],
  };

  onOpenChange = (openKeys: Array<any>) => {
    const { menus } = this.props;
    const rootSubmenuKeys = menus.filter((_) => !_.menuParentId).map((_) => _.id);

    const latestOpenKey = openKeys.find((key) => this.state.openKeys.indexOf(key) === -1);

    let newOpenKeys = openKeys;
    if (rootSubmenuKeys.indexOf(latestOpenKey) !== -1) {
      newOpenKeys = latestOpenKey ? [latestOpenKey] : [];
    }

    this.setState({
      openKeys: newOpenKeys,
    });
    store.set('openKeys', newOpenKeys);
  };

  generateMenus = (data: Array<any>) => {
    return data.map((item) => {
      if (item.children) {
        return (
          <SubMenu
            key={item.id}
            title={
              <Fragment>
                {item.icon && <LegacyIcon type={item.icon} />}
                <span>{item.name}</span>
              </Fragment>
            }
          >
            {this.generateMenus(item.children)}
          </SubMenu>
        );
      }
      return (
        <Menu.Item key={item.id}>
          <NavLink to={item.route || '#'}>
            {item.icon && <LegacyIcon type={item.icon} />}
            <span>{item.name}</span>
          </NavLink>
        </Menu.Item>
      );
    });
  };

  render() {
    const { collapsed, theme, menus, location, isMobile, onCollapseChange } = this.props;

    // Generating tree-structured data for menu content.
    const menuTree = arrayToTree(menus, 'id', 'menuParentId');

    // Find a menu that matches the pathname.
    const currentMenu = menus.find((_) => _.route && pathToRegexp(_.route).exec(location.pathname));

    // Find the key that should be selected according to the current menu.
    const selectedKeys = currentMenu ? queryAncestors(menus, currentMenu, 'menuParentId').map((_) => _.id) : [];

    const menuProps = collapsed
      ? {}
      : {
          openKeys: this.state.openKeys,
        };

    return (
      <Menu
        mode="inline"
        theme={theme}
        onOpenChange={this.onOpenChange}
        selectedKeys={selectedKeys}
        onClick={
          isMobile
            ? () => {
                onCollapseChange(true);
              }
            : undefined
        }
        {...menuProps}
      >
        {this.generateMenus(menuTree)}
      </Menu>
    );
  }
}

export default SiderMenu;
