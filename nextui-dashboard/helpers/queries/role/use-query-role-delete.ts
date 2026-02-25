/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesRoleDeleteRequest,
  type UsecasesRoleDeleteResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesRoleDeleteRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesRoleDeleteResponse['result']>;

export const useQueryRoleDelete = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'role.delete', params, { retry: 3 })
  );
};
