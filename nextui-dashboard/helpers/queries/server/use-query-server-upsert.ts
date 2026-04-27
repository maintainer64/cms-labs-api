/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesServerEditRequest,
  type UsecasesServerEditResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesServerEditRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesServerEditResponse['result']>;

export const useQueryServerUpsert = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'server.upsert', params, { retry: 3 })
  );
};
