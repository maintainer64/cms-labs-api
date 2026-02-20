/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesTargetListRequest,
  type UsecasesTargetListResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesTargetListRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesTargetListResponse['result']>;

export const useQueryTargetList = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'target.list', params, {
      retry: 3,
      refetchOnWindowFocus: false,
      refetchInterval: false,
      gcTime: 0
    })
  );
};
