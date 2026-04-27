/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  ClabgateJsonRpcPath,
  CoreJsonRpcPath,
  transportWithAuth,
  UsecasesServerGetRequest,
  UsecasesServerGetResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesServerGetRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesServerGetResponse['result']>;

export const useMutationServerGet = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params?.id) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'server.get',
        params: params
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'server.delete'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'server.get'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'server.list'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'server.ping'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'server.upsert'] });
    }
  });
};
