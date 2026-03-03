/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  UsecasesServiceCardGetRequest,
  UsecasesServiceCardGetResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesServiceCardGetRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesServiceCardGetResponse['result']>;

export const useMutationServiceCardGet = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params?.id) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'service_card.get',
        params: params
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'service_card.delete'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'service_card.get'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'service_card.list'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'service_card.upsert'] });
    }
  });
};
