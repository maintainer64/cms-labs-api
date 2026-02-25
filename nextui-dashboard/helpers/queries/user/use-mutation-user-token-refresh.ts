/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  AuthRenewManagerRefreshRequest,
  AuthSwaggerSSOTokenResponse,
  CoreJsonRpcPath,
  transportWithAuth
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<AuthRenewManagerRefreshRequest['params']>;
type Response = CamelCasedPropertiesDeep<AuthSwaggerSSOTokenResponse['result']>;

export const useMutationUserTokenRefresh = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'user.token_refresh',
        params: params
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
    }
  });
};
