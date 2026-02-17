/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesTopologiesGetRequest,
  type UsecasesTopologiesGetResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesTopologiesGetRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesTopologiesGetResponse['result']>;

export const useQueryTopologyGet = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'topology.get', params, {
      retry: 3,
      enabled: !!params.namespace
    })
  );
};
