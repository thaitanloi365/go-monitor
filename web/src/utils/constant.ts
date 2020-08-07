import { IMenus } from 'types/app';

export const ROLE_TYPE = {
  ADMIN: 'admin',
  DEFAULT: 'admin',
  DEVELOPER: 'developer',
};

export const CANCEL_REQUEST_MESSAGE = 'cancel request';

const menus: IMenus = [
  {
    id: '1',
    icon: 'dashboard',
    name: 'Dashboard',
    route: '/dashboard',
  },
  {
    id: '11',
    menuParentId: '-1',
    breadcrumbParentId: '1',
    name: 'Websocket',
    route: '/ws',
  },
  {
    id: '2',
    breadcrumbParentId: '1',
    name: 'Users',
    icon: 'user',
    route: '/users',
  },
  {
    id: '21',
    menuParentId: '-1',
    breadcrumbParentId: '2',
    name: 'User Detail',
    route: '/users/:id',
  },
  {
    id: '3',
    breadcrumbParentId: '1',
    name: 'Merchants',
    icon: 'solution',
    route: '/merchants',
  },
  {
    id: '31',
    menuParentId: '-1',
    breadcrumbParentId: '3',
    name: 'Merchant Details',
    route: '/merchants/:id',
  },
  {
    id: '4',
    breadcrumbParentId: '1',
    name: 'Locations',
    icon: 'home',
    route: '/locations',
  },
  {
    id: '41',
    menuParentId: '-1',
    breadcrumbParentId: '4',
    name: 'Location Details',
    route: '/locations/:id',
  },
  {
    id: '411',
    menuParentId: '-1',
    breadcrumbParentId: '4',
    name: 'Omnivore Location Details',
    route: '/locations/:id/omnivore',
  },
  {
    id: '5',
    breadcrumbParentId: '1',
    name: 'Tabs',
    icon: 'book',
    route: '/tabs',
  },
  {
    id: '51',
    menuParentId: '-1',
    breadcrumbParentId: '5',
    name: 'Tab Details',
    route: '/tabs/:id',
  },
  {
    id: '6',
    breadcrumbParentId: '1',
    name: 'Orders',
    icon: 'red-envelope',
    route: '/orders',
  },
  {
    id: '61',
    menuParentId: '-1',
    breadcrumbParentId: '6',
    name: 'Order Details',
    route: '/orders/:id',
  },
  {
    id: '7',
    breadcrumbParentId: '1',
    name: 'Subscriptions',
    icon: 'audit',
    route: '/subscriptions',
  },
];

export default { menus };
