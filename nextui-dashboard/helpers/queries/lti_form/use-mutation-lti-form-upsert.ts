/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  UsecasesAuthProviderEditRequest,
  UsecasesAuthProviderEditResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesAuthProviderEditRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesAuthProviderEditResponse['result']>;

export const useMutationAuthProviderUpsert = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    mutationFn: (params: Params) => {
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'auth_provider.upsert',
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
