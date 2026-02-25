/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import { AuthSSOAuthorizeRequest, AuthSSOAuthorizeResponse, CoreJsonRpcPath, transportWithAuth } from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import { CamelCasedPropertiesDeep } from 'type-fest';
import queryClient from '@/helpers/queries/base';

type Params = CamelCasedPropertiesDeep<AuthSSOAuthorizeRequest['params']>;
type Response = CamelCasedPropertiesDeep<AuthSSOAuthorizeResponse['result']>;

export const useMutationSsoAuthorize = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'sso.authorize',
        params: params
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'sso.authorize'] });
    }
  });
};
