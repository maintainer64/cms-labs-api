/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesTargetGetRequest,
  type UsecasesTargetGetResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesTargetGetRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesTargetGetResponse['result']>;

export const useQueryTargetGet = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'target.get', params, {
      retry: 3,
      enabled: !!params.id
    })
  );
};
