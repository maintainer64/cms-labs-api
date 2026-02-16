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
        method: 'lti_form.delete',
        params: params
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'lti_form.delete'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'lti_form.get'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'lti_form.list'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'lti_form.sso_list_get'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'lti_form.upsert'] });
    }
  });
};
