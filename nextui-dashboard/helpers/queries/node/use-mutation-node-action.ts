/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  ClabgateJsonRpcPath,
  transportWithAuth,
  UsecasesNodeActionRequest,
  UsecasesNodeActionResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesNodeActionRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesNodeActionResponse['result']>;

export const useMutationNodeAction = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params?.sessionId || !params?.actions) return null;
      return transportWithAuth.rpc(ClabgateJsonRpcPath, {
        method: 'node.action',
        params: params
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: [ClabgateJsonRpcPath, 'topology.get'] });
    }
  });
};
