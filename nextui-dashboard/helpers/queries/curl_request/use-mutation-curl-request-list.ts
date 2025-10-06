/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  UsecasesCurlRequestListRequest,
  UsecasesCurlRequestListResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesCurlRequestListRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesCurlRequestListResponse['result']>;

export const useMutationCurlRequestList = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'curl_request.list',
        params: params
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'curl_request.delete'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'curl_request.get'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'curl_request.list'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'curl_request.upsert'] });
    }
  });
};
