import { useMutation } from '@tanstack/react-query';
import {
  ClabgateJsonRpcPath,
  transportWithAuth,
  type UsecasesSessionCheckInputDTO,
  type UsecasesSessionCheckOutputDTO
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesSessionCheckInputDTO>;
type Response = CamelCasedPropertiesDeep<UsecasesSessionCheckOutputDTO>;

export const useMutationSessionCheck = () =>
  useMutation<Response, unknown, Params>({
    mutationFn: (params) => transportWithAuth.rpc(ClabgateJsonRpcPath, { method: 'session.check', params })
  });
