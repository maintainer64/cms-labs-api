/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  UsecasesAuthProviderListRequest,
  UsecasesAuthProviderListResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesAuthProviderListRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesAuthProviderListResponse['result']>;

export const useMutationAuthProviderList = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'auth_provider.list',
        params: params
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'auth_provider.delete'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'auth_provider.get'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'auth_provider.list'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'auth_provider.upsert'] });
    }
  });
};
