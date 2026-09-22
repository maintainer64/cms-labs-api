import { useMutation } from '@tanstack/react-query';
import {
  ClabgateJsonRpcPath,
  transportWithAuth,
  type UsecasesSessionEnsureInputDTO,
  type UsecasesSessionOutputDTO
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesSessionEnsureInputDTO>;
type Response = CamelCasedPropertiesDeep<UsecasesSessionOutputDTO>;

export const useMutationSessionEnsure = () =>
  useMutation<Response, unknown, Params>({
    mutationFn: (params) => transportWithAuth.rpc(ClabgateJsonRpcPath, { method: 'session.ensure', params })
  });
