/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  ModelsPNETServer,
  transportWithAuth,
  type UsecasesPNETServerGetRequest,
  type UsecasesPNETServerGetResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesPNETServerGetRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesPNETServerGetResponse['result']>;

export interface PnetServerItem extends CamelCasedPropertiesDeep<ModelsPNETServer> {
  roles?: Array<number>;
}

export function MapServerItem(server?: CamelCasedPropertiesDeep<ModelsPNETServer>, roles?: Array<number>) {
  if (!server) return {} as PnetServerItem;
  const newModel = server as PnetServerItem;
  newModel.roles = roles;
  return newModel;
}

export const useQueryServerGet = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'server.get', params, {
      retry: 3,
      enabled: !!params.id
    })
  );
};
