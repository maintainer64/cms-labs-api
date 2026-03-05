/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  UsecasesCurlRequestEditRequest,
  UsecasesCurlRequestEditResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesCurlRequestEditRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesCurlRequestEditResponse['result']>;

export const useMutationCurlRequestUpsert = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    mutationFn: (params: Params) => {
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'curl_request.upsert',
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
