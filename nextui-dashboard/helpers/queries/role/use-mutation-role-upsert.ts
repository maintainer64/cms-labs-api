/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import { CoreJsonRpcPath, transportWithAuth, UsecasesRoleEditRequest, UsecasesRoleEditResponse } from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesRoleEditRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesRoleEditResponse['result']>;

export const useMutationRoleUpsert = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params?.id) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'role.upsert',
        params: params
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'role.delete'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'role.list'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'role.upsert'] });
    }
  });
};
