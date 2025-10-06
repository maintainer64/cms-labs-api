/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  UsecasesLTIRoutingListRequest,
  UsecasesLTIRoutingListResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesLTIRoutingListRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesLTIRoutingListResponse['result']>;

export const useMutationLtiRoutingList = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: Params) => {
      if (!params) return null;
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'lti_routing.list',
        params: params
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'lti_routing.delete'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'lti_routing.get'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'lti_routing.list'] });
      await queryClient.invalidateQueries({ queryKey: [CoreJsonRpcPath, 'lti_routing.upsert'] });
    }
  });
};
