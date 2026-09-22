import { useMutation } from '@tanstack/react-query';
import {
  ClabgateJsonRpcPath,
  transportWithAuth,
  type UsecasesSessionStopInputDTO,
  type UsecasesSessionStopOutputDTO
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesSessionStopInputDTO>;
type Response = CamelCasedPropertiesDeep<UsecasesSessionStopOutputDTO>;

export const useMutationSessionStop = () =>
  useMutation<Response, unknown, Params>({
    mutationFn: (params) => transportWithAuth.rpc(ClabgateJsonRpcPath, { method: 'session.stop', params })
  });
