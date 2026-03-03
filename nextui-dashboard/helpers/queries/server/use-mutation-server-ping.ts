/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  ExternalPNETServerPingRequest,
  ExternalPNETServerPingResponse,
  ClabgateJsonRpcPath
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<ExternalPNETServerPingRequest['params']>;
type Response = CamelCasedPropertiesDeep<ExternalPNETServerPingResponse['result']>;

export const useMutationServerPing = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'server.ping',
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
