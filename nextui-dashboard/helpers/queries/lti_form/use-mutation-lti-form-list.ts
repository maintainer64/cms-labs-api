/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  UsecasesLTIFormListRequest,
  UsecasesLTIFormListResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesLTIFormListRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesLTIFormListResponse['result']>;

export const useMutationLtiFormList = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'lti_form.list',
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
