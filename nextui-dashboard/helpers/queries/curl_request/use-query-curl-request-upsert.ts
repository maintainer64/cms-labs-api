/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesCurlRequestEditRequest,
  type UsecasesCurlRequestEditResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesCurlRequestEditRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesCurlRequestEditResponse['result']>;

export const useQueryCurlRequestUpsert = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'curl_request.upsert', params, { retry: 3 })
  );
};
