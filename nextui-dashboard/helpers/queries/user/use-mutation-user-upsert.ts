/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import { CoreJsonRpcPath, transportWithAuth, UsecasesUserEditRequest, UsecasesUserEditResponse } from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesUserEditRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesUserEditResponse['result']>;

export const useMutationUserUpsert = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    mutationFn: (params: Params) => {
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'user.upsert',
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
