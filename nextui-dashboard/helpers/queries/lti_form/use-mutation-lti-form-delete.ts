/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  UsecasesAuthProviderDeleteRequest,
  UsecasesAuthProviderDeleteResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesAuthProviderDeleteRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesAuthProviderDeleteResponse['result']>;

export const useMutationAuthProviderDelete = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params?.id) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'auth_provider.delete',
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
