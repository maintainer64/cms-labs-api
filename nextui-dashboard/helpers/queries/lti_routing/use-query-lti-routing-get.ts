/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesLTIRoutingGetRequest,
  type UsecasesLTIRoutingGetResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesLTIRoutingGetRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesLTIRoutingGetResponse['result']>;

export const useQueryLtiRoutingGet = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'lti_routing.get', params, { retry: 3 })
  );
};
