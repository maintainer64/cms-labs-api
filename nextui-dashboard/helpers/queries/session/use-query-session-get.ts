import { useQuery } from '@tanstack/react-query';
import {
  ClabgateJsonRpcPath,
  transportWithAuth,
  type UsecasesSessionGetInputDTO,
  type UsecasesSessionOutputDTO
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesSessionGetInputDTO> & object;
type Response = CamelCasedPropertiesDeep<UsecasesSessionOutputDTO>;

export const useQuerySessionGet = (params: Params, enabled = true, poll = false) =>
  useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(ClabgateJsonRpcPath, 'session.get', params, {
      enabled: enabled && !!params.sessionId,
      retry: 3,
      refetchOnWindowFocus: true,
      refetchInterval: (query: { state: { data?: Response } }) => {
        if (poll) return 2000;
        const phase = query.state.data?.session?.phase;
        return phase === 'ready' || phase === 'failed' || phase === 'degraded' ? false : 2000;
      }
    })
  );
