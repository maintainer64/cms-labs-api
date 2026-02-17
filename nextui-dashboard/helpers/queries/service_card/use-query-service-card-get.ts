/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesServiceCardGetRequest,
  type UsecasesServiceCardGetResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesServiceCardGetRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesServiceCardGetResponse['result']>;

export const useQueryServiceCardGet = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'service_card.get', params, {
      retry: 3,
      enabled: !!params.id
    })
  );
};
