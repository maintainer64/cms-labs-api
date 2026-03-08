/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesUserGetRequest,
  type UsecasesUserGetResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesUserGetRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesUserGetResponse['result']>;

export const useQueryUserGet = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'user.get', params, {
      retry: 3,
      enabled: !!params.id
    })
  );
};
