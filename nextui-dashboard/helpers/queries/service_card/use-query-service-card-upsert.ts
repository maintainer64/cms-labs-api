/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesServiceCardEditRequest,
  type UsecasesServiceCardEditResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesServiceCardEditRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesServiceCardEditResponse['result']>;

export const useQueryServiceCardUpsert = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'service_card.upsert', params, { retry: 3 })
  );
};
