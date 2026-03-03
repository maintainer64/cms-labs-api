/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  ClabgateJsonRpcPath,
  transportWithAuth,
  UsecasesTopologyCreateRequest,
  UsecasesTopologyCreateResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesTopologyCreateRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesTopologyCreateResponse['result']>;

export const useMutationTopologyCreate = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params) return null;
      return transportWithAuth.rpc(ClabgateJsonRpcPath, {
        method: 'topology.create',
        params: params
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: [ClabgateJsonRpcPath, 'topology.create'] });
      await queryClient.invalidateQueries({ queryKey: [ClabgateJsonRpcPath, 'topology.delete'] });
      await queryClient.invalidateQueries({ queryKey: [ClabgateJsonRpcPath, 'topology.get'] });
    }
  });
};
