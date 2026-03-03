/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesCurlRequestListRequest,
  type UsecasesCurlRequestListResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesCurlRequestListRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesCurlRequestListResponse['result']>;

export const useQueryCurlRequestList = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'curl_request.list', params, { retry: 3 })
  );
};
