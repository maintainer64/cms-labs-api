/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  ServerQueueServerQueueUpsertRequest,
  ServerQueueServerQueueUpsertResponse,
  transportWithAuth
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<ServerQueueServerQueueUpsertRequest['params']>;
type Response = CamelCasedPropertiesDeep<ServerQueueServerQueueUpsertResponse['result']>;

export const useMutationServerQueueUpsert = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params?.id) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'server_queue.upsert',
        params: params
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'server_queue.list'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'server_queue.upsert'] });
    }
  });
};
