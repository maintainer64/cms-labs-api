/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  AuthRenewManagerCredentialsRequest,
  AuthSwaggerSSOTokenResponse,
  CoreJsonRpcPath,
  transportWithAuth, transportWithoutAuth
} from '@/helpers/api';
import { TFormikData, TMutationCustomOptions } from '@/helpers/queries/types';
import { CamelCasedPropertiesDeep } from 'type-fest';
import queryClient from '@/helpers/queries/base';

type Params = CamelCasedPropertiesDeep<AuthRenewManagerCredentialsRequest['params']>;
type Response = CamelCasedPropertiesDeep<AuthSwaggerSSOTokenResponse['result']>;
type FormData = TFormikData<Params>;

export const useMutationUserLogin = (options: TMutationCustomOptions<Response, FormData> = {}) => {
  return useMutation<Response, unknown, FormData>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ values }: FormData) => {
      if (values === null) return;
      return transportWithoutAuth.t.rpc(CoreJsonRpcPath, {
        method: 'user.login',
        params: values
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
