/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  UsecasesLTIAttemptListRequest,
  UsecasesLTIAttemptListResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesLTIAttemptListRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesLTIAttemptListResponse['result']>;

export const useMutationLtiAttemptList = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'lti_attempt.list',
        params: params
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'lti_attempt.create'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'lti_attempt.update'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'lti_attempt.delete'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'lti_attempt.get'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'lti_attempt.list'] });
    }
  });
};
