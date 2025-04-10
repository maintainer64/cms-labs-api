import { models_PNETServer } from '@/helpers/api';

export interface PnetServerItem extends models_PNETServer {
  roles?: Array<number>;
}

export function MapServerItem(server?: models_PNETServer, roles?: Array<number>) {
  if (!server) return {} as PnetServerItem;
  const newModel = server as PnetServerItem;
  newModel.roles = roles;
  return newModel;
}
