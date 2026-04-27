/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesServerListRequest,
  type UsecasesServerListResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesServerListRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesServerListResponse['result']>;

export const useQueryServerList = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'server.list', params, { retry: 3 })
  );
};
