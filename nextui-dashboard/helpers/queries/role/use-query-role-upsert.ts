/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesRoleEditRequest,
  type UsecasesRoleEditResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesRoleEditRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesRoleEditResponse['result']>;

export const useQueryRoleUpsert = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'role.upsert', params, { retry: 3 })
  );
};
