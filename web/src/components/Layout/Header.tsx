import React, { PureComponent, Fragment } from 'react';
import { Menu, Layout, Avatar, Popover, Badge, List } from 'antd';
import { IUser, IMenus } from 'types/app';
import { Ellipsis } from 'components';
import { Icon as LegacyIcon } from '@ant-design/compatible';
import { BellOutlined, RightOutlined } from '@ant-design/icons';
import moment from 'moment';
import classnames from 'classnames';
import styles from './Header.less';

const { SubMenu } = Menu;

export interface IHeaderProps {
  fixed: boolean;
  user: IUser;
  menus: IMenus;
  collapsed: boolean;
  onSignOut: () => void;
  onEdit: () => void;
  notifications: any[];
  onCollapseChange: (collapsed: boolean) => void;
  onAllNotificationsRead: () => void;
}

class Header extends PureComponent<IHeaderProps> {
  handleClickMenu = (e: any) => {
    e.key === 'SignOut' && this.props.onSignOut();
  };
  render() {
    const { fixed, user, collapsed, notifications, onCollapseChange, onAllNotificationsRead } = this.props;

    const rightContent = [
      <Menu key="user" mode="horizontal" onClick={this.handleClickMenu}>
        <SubMenu
          title={
            <Fragment>
              <span style={{ color: '#999', marginRight: 4 }}>Hi,</span>
              <span>{user?.name}</span>
              <Avatar style={{ marginLeft: 8 }} src={user?.avatar} />
            </Fragment>
          }
        >
          <Menu.Item key="SignOut">Sign out</Menu.Item>
        </SubMenu>
      </Menu>,
    ];

    rightContent.unshift(
      <Popover
        placement="bottomRight"
        trigger="click"
        key="notifications"
        overlayClassName={styles.notificationPopover}
        getPopupContainer={() => document.querySelector('#primaryLayout')}
        content={
          <div className={styles.notification}>
            <List
              itemLayout="horizontal"
              dataSource={notifications}
              locale={{
                emptyText: 'You have viewed all notifications.',
              }}
              renderItem={(item: any) => (
                <List.Item className={styles.notificationItem}>
                  <List.Item.Meta
                    title={
                      <Ellipsis tooltip lines={1}>
                        {item.title}
                      </Ellipsis>
                    }
                    description={moment(item.date).fromNow()}
                  />
                  <RightOutlined style={{ fontSize: 10, color: '#ccc' }} />
                </List.Item>
              )}
            />
            {notifications.length ? (
              <div onClick={onAllNotificationsRead} className={styles.clearButton}>
                Clear notifications
              </div>
            ) : null}
          </div>
        }
      >
        <Badge count={notifications.length} dot offset={[-10, 10]} className={styles.iconButton}>
          <BellOutlined className={styles.iconFont} />
        </Badge>
      </Popover>
    );

    return (
      <Layout.Header
        className={classnames(styles.header, {
          [styles.fixed]: fixed,
          [styles.collapsed]: collapsed,
        })}
        id="layoutHeader"
      >
        <div className={styles.button} onClick={onCollapseChange.bind(this, !collapsed)}>
          <LegacyIcon
            type={classnames({
              'menu-unfold': collapsed,
              'menu-fold': !collapsed,
            })}
          />
        </div>

        <div className={styles.rightContainer}>{rightContent}</div>
      </Layout.Header>
    );
  }
}

export default Header;
