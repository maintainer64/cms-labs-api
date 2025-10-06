/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesTopologyDeleteRequest,
  type UsecasesTopologyDeleteResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesTopologyDeleteRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesTopologyDeleteResponse['result']>;

export const useQueryTopologyDelete = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'topology.delete', params, { retry: 3 })
  );
};
