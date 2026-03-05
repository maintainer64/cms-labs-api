/*
Template auto generated with params from openapi.json
*/
import { useMutation } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  UsecasesLTIRoutingEditRequest,
  UsecasesLTIRoutingEditResponse
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesLTIRoutingEditRequest['params']>;
type Response = CamelCasedPropertiesDeep<UsecasesLTIRoutingEditResponse['result']>;

export const useMutationLtiRoutingUpsert = (options: TMutationCustomOptions<Response, Params> = {}) => {
  return useMutation<Response, unknown, Params>({
    mutationFn: (params: Params) => {
      return transportWithAuth.rpc(CoreJsonRpcPath, {
        method: 'lti_routing.upsert',
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
