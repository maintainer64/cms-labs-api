/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithoutAuth,
  type UsecasesAuthProviderListRequest,
  type UsecasesAuthProviderListResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesAuthProviderListRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesAuthProviderListResponse['result']>;

export const useQueryAuthProviderList = (params: Params) => {
  return useQuery(
    transportWithoutAuth.t.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'auth_provider.list', params, {
      retry: 3
    })
  );
};
