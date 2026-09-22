import { useMutation } from '@tanstack/react-query';
import {
  ClabgateJsonRpcPath,
  transportWithAuth,
  type UsecasesSessionOpenInputDTO,
  type UsecasesSessionOpenOutputDTO
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesSessionOpenInputDTO>;
type Response = CamelCasedPropertiesDeep<UsecasesSessionOpenOutputDTO>;

export const useMutationSessionOpen = () =>
  useMutation<Response, unknown, Params>({
    mutationFn: (params) => transportWithAuth.rpc(ClabgateJsonRpcPath, { method: 'session.open', params })
  });
