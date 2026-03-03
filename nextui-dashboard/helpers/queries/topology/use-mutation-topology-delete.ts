/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  ClabgateJsonRpcPath,
  CoreJsonRpcPath,
  transportWithAuth,
  UsecasesTopologyDeleteRequest,
  UsecasesTopologyDeleteResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesTopologyDeleteRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesTopologyDeleteResponse['result']>;

export const useMutationTopologyDelete = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params?.namespaces) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'topology.delete',
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
