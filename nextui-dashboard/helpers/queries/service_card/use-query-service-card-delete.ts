/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesServiceCardDeleteRequest,
  type UsecasesServiceCardDeleteResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesServiceCardDeleteRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesServiceCardDeleteResponse['result']>;

export const useQueryServiceCardDelete = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'service_card.delete', params, { retry: 3 })
  );
};
