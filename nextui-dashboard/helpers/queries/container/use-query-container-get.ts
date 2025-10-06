/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  ClabgateJsonRpcPath,
  transportWithAuth,
  type UsecasesContainerGetRequest,
  type UsecasesContainersGetResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesContainerGetRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesContainersGetResponse['result']>;

export const useQueryContainerGet = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(ClabgateJsonRpcPath, 'container.get', params, { retry: 3 })
  );
};
