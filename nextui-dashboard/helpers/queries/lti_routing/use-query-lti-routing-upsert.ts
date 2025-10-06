/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesLTIRoutingEditRequest,
  type UsecasesLTIRoutingEditResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesLTIRoutingEditRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesLTIRoutingEditResponse['result']>;

export const useQueryLtiRoutingUpsert = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'lti_routing.upsert', params, { retry: 3 })
  );
};
