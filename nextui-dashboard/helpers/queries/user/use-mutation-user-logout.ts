/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import { AuthSwaggerSSOTokenResponse, AuthUserLogoutRequest, CoreJsonRpcPath, transportWithAuth } from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<AuthUserLogoutRequest['params']>;
type Response = CamelCasedPropertiesDeep<AuthSwaggerSSOTokenResponse['result']>;

export const useMutationUserLogout = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params?.id) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'user.logout',
        params: params
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'user.list'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'user.upsert'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'user.get'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'user.upsert'] });
    }
  });
};
