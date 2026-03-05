/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  UsecasesServiceCardEditRequest,
  UsecasesServiceCardEditResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesServiceCardEditRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesServiceCardEditResponse['result']>;

export const useMutationServiceCardUpsert = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    mutationFn: (params: Params) => {
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'service_card.upsert',
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
