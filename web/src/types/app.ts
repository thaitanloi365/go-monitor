export interface IMenuItem {
  id: string;
  icon?: string;
  name?: string;
  route: string;
  breadcrumbParentId?: string;
  menuParentId?: string;
}

export type IMenus = IMenuItem[];

export interface INotificationItem {
  title: string;
  date: Date;
}

export interface IUser {
  id: number;
  updated_at: number;
  created_at: number;
  deleted_at: number;
  email: string;
  name: string;
  first_name: string;
  role: string;
  last_name: string;
  last_login: number;
  phone: string;
  avatar: string;

  token: string;
  logged_out_at: number;
}
