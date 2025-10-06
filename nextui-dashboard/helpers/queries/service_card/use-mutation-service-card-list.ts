/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  UsecasesServiceCardListRequest,
  UsecasesServiceCardListResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesServiceCardListRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesServiceCardListResponse['result']>;

export const useMutationServiceCardList = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'service_card.list',
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
