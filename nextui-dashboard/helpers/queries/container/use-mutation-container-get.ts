/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  ClabgateJsonRpcPath,
  transportWithAuth,
  UsecasesContainerGetRequest,
  UsecasesContainersGetResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import { CamelCasedPropertiesDeep } from 'type-fest';
import queryClient from '@/helpers/queries/base';

type Params = CamelCasedPropertiesDeep<UsecasesContainerGetRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesContainersGetResponse['result']>;

export const useMutationContainerGet = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params?.deployment) return null;
      return transportWithAuth.rpc(ClabgateJsonRpcPath, {
        method: 'container.get',
        params: params
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: [ClabgateJsonRpcPath, 'container.get'] });
    }
  });
};
