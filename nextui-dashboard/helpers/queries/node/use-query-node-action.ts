/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  ClabgateJsonRpcPath,
  transportWithAuth,
  type UsecasesNodeActionRequest,
  type UsecasesNodeActionResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesNodeActionRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesNodeActionResponse['result']>;

export const useQueryNodeAction = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(ClabgateJsonRpcPath, 'node.action', params, { retry: 3 })
  );
};
