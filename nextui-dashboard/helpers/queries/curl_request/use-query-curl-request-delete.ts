/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesCurlRequestDeleteRequest,
  type UsecasesCurlRequestDeleteResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesCurlRequestDeleteRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesCurlRequestDeleteResponse['result']>;

export const useQueryCurlRequestDelete = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'curl_request.delete', params, { retry: 3 })
  );
};
