/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesLTIRoutingDeleteRequest,
  type UsecasesLTIRoutingDeleteResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesLTIRoutingDeleteRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesLTIRoutingDeleteResponse['result']>;

export const useQueryLtiRoutingDelete = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'lti_routing.delete', params, { retry: 3 })
  );
};
