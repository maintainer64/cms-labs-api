/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  ClabgateJsonRpcPath,
  transportWithAuth,
  UsecasesTopologiesGetRequest,
  UsecasesTopologiesGetResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesTopologiesGetRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesTopologiesGetResponse['result']>;

export const useMutationTopologyGet = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params?.attemptNumber || !params?.username) return null;
      return transportWithAuth.rpc(ClabgateJsonRpcPath, {
        method: 'topology.get',
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
