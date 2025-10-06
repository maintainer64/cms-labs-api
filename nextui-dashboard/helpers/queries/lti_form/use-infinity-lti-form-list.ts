import { useInfiniteQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  UsecasesLTIFormListRequest,
  UsecasesLTIFormListResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';
import { RpcTransport } from '@/helpers/api/core/request';

type Params = CamelCasedPropertiesDeep<UsecasesLTIFormListRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesLTIFormListResponse['result']>;

export const useInfinityLtiFormList = (params?: Params) => {
  return useInfiniteQuery({
    queryKey: RpcTransport.getInfiniteQueryKey(CoreJsonRpcPath, 'lti_form.list', params ?? {}),
    queryFn: ({ pageParam }) => {
      return transportWithAuth.rpc<Response>(CoreJsonRpcPath, {
        method: 'lti_form.list',
        params: {
          limit: params?.limit ?? 100,
          offset: pageParam,
          search: (params?.search?.length || '') < 3 ? '' : params?.search
        } as Params
      });
    },
    refetchOnWindowFocus: false,
    refetchInterval: false,
    gcTime: 0,
    getNextPageParam: (response, pages) => {
      const totalCount = response?.totalCount || 0;
      const count = pages.flatMap((p) => p?.model).length;
      return totalCount && totalCount > count ? count : undefined;
    },
    initialPageParam: 0,
    retry: 3
  });
};
