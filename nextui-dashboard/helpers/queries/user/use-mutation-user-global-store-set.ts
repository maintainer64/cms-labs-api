/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import { CoreJsonRpcPath, transportWithAuth, TypesUserStoreSetRequest, TypesUserStoreSetResponse } from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<TypesUserStoreSetRequest['params']>;
type Response = CamelCasedPropertiesDeep<TypesUserStoreSetResponse['result']>;

export const useMutationUserGlobalStoreSet = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'user.global_store_set',
        params: params
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'user.global_store_get'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'user.global_store_set'] });
    }
  });
};
