/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  ModelsServer,
  transportWithAuth,
  type UsecasesServerGetRequest,
  type UsecasesServerGetResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesServerGetRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesServerGetResponse['result']>;

export interface ServerItem extends CamelCasedPropertiesDeep<ModelsServer> {
  roles?: Array<number>;
  token?: string;
}

export function MapServerItem(server?: CamelCasedPropertiesDeep<ModelsServer>, roles?: Array<number>) {
  if (!server) return {} as ServerItem;
  const newModel = server as ServerItem;
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
