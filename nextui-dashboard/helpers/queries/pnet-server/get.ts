import { useQuery, useInfiniteQuery } from '@tanstack/react-query';
import { postV1PnetServerList, postV1PnetServerGet, usecases_PNETServerListInputDTO } from '@/helpers/api';

export const usePnetServerList = (params?: usecases_PNETServerListInputDTO) => {
  return useInfiniteQuery({
    queryKey: ['postV1PnetServerList', params?.search, params?.status, params?.order_by],
    queryFn: ({ pageParam }) => {
      return postV1PnetServerList({
        form: {
          limit: params?.limit ?? 100,
          offset: pageParam,
          search: (params?.search?.length || '') < 3 ? '' : params?.search,
          order_by: params?.order_by || '',
          status: params?.status || ''
        }
      });
    },
    refetchOnWindowFocus: false,
    refetchInterval: false,
    refetchOnMount: false,
    getNextPageParam: (response, pages) => {
      const totalCount = response.result?.total_count || 0;
      const count = pages.flatMap((p) => p.result?.model).length;
      return totalCount && totalCount > count ? count : undefined;
    },
    initialPageParam: 0,
    retry: 0
  });
};

export const usePnetServerByID = (id?: number) => {
  return useQuery({
    queryKey: ['postV1PnetServerGet', id],
    queryFn: () => {
      return id ? postV1PnetServerGet({ form: { id: id } }) : undefined;
    },
    retry: 0
  });
};
