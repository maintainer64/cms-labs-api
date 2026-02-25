/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  ClabgateJsonRpcPath,
  transportWithAuth,
  UsecasesTaskListRequest,
  UsecasesTaskListResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import { CamelCasedPropertiesDeep } from 'type-fest';
import queryClient from '@/helpers/queries/base';

type Params = CamelCasedPropertiesDeep<UsecasesTaskListRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesTaskListResponse['result']>;

export const useMutationTaskList = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    mutationFn: (params: Params) => {
      return transportWithAuth.rpc(ClabgateJsonRpcPath, {
        method: 'task.list',
        params: params
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: [ClabgateJsonRpcPath, 'task.list'] });
    }
  });
};
