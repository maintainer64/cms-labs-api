/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesServiceCardListRequest,
  type UsecasesServiceCardListResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesServiceCardListRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesServiceCardListResponse['result']>;

export const useQueryServiceCardList = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'service_card.list', params, { retry: 3 })
  );
};
