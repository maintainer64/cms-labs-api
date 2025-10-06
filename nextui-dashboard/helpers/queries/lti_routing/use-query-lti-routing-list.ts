/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesLTIRoutingListRequest,
  type UsecasesLTIRoutingListResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesLTIRoutingListRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesLTIRoutingListResponse['result']>;

export const useQueryLtiRoutingList = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'lti_routing.list', params, { retry: 3 })
  );
};
