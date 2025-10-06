/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import { CoreJsonRpcPath, transportWithAuth, UsecasesUserListRequest, UsecasesUserListResponse } from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesUserListRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesUserListResponse['result']>;

export const useMutationUserList = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'user.list',
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
