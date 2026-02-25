/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesUserEditRequest,
  type UsecasesUserEditResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesUserEditRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesUserEditResponse['result']>;

export const useQueryUserUpsert = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'user.upsert', params, { retry: 3 })
  );
};
