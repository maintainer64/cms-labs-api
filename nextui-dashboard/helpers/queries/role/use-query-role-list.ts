/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesRoleListRequest,
  type UsecasesRoleListResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesRoleListRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesRoleListResponse['result']>;

export const useQueryRoleList = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'role.list', params, { retry: 3 })
  );
};
