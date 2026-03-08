/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import { CoreJsonRpcPath, transportWithAuth } from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';

type Params = {
  fromTargetId?: string;
  toTargetId?: string;
  relationType?: string;
};
type Response = object;

export const useMutationTargetRelationDelete = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    mutationFn: (params: Params) => {
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'target.relation_delete',
        params: params
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'target.list'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'target.get'] });
    }
  });
};
