/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesCurlRequestGetRequest,
  type UsecasesCurlRequestGetResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesCurlRequestGetRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesCurlRequestGetResponse['result']>;

export const useQueryCurlRequestGet = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'curl_request.get', params, {
      retry: 3,
      enabled: !!params.id
    })
  );
};
