import { models_User } from '@/helpers/api';

export interface UserItem extends models_User {
  roles?: Array<number>;
}

export function MapUserItem(user?: models_User, roles?: Array<number>) {
  if (!user) return {} as UserItem;
  const newModel = user as UserItem;
  newModel.roles = roles;
  return newModel;
}
