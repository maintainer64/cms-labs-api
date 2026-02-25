/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  AuthUserPasswordChangeRequest,
  AuthUserPasswordChangeResponse,
  CoreJsonRpcPath,
  transportWithAuth
} from '@/helpers/api';
import { TFormikData, TMutationCustomOptions } from '@/helpers/queries/types';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Response = CamelCasedPropertiesDeep<AuthUserPasswordChangeResponse['result']>;
type Params = CamelCasedPropertiesDeep<AuthUserPasswordChangeRequest['params']>;
type FormData = TFormikData<Params>;

export const useMutationUserPasswordChange = (options: TMutationCustomOptions<Response, FormData> = {}) => {
  return useMutation<Response, unknown, FormData>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ values }: FormData) => {
      if (values === null) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'user.password_change',
        params: values
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
