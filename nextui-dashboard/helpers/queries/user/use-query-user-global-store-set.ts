/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type TypesUserStoreSetRequest,
  type TypesUserStoreSetResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<TypesUserStoreSetRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<TypesUserStoreSetResponse['result']>;

export const useQueryUserGlobalStoreSet = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'user.global_store_set', params, { retry: 3 })
  );
};
