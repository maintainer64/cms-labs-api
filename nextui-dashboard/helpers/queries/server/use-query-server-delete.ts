/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesServerDeleteRequest,
  type UsecasesServerDeleteResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesServerDeleteRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesServerDeleteResponse['result']>;

export const useQueryServerDelete = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'server.delete', params, { retry: 3 })
  );
};
