/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesPNETServerListRequest,
  type UsecasesPNETServerListResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesPNETServerListRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesPNETServerListResponse['result']>;

export const useQueryServerList = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'server.list', params, { retry: 3 })
  );
};
