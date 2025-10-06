/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type TypesUserStoreGetRequest,
  type TypesUserStoreGetResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<TypesUserStoreGetRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<TypesUserStoreGetResponse['result']>;

export const useQueryUserGlobalStoreGet = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'user.global_store_get', params, { retry: 3 })
  );
};
