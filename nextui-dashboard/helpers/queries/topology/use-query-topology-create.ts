/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesTopologyCreateRequest,
  type UsecasesTopologyCreateResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesTopologyCreateRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesTopologyCreateResponse['result']>;

export const useQueryTopologyCreate = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'topology.create', params, { retry: 3 })
  );
};
