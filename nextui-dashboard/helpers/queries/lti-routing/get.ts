import { useQuery, useInfiniteQuery } from '@tanstack/react-query';
import { postV1LtiRoutingGet, postV1LtiRoutingList, usecases_LTIRoutingListInputDTO } from '@/helpers/api';

export const useLTIRoutingList = (params?: usecases_LTIRoutingListInputDTO) => {
  return useInfiniteQuery({
    queryKey: ['postV1LtiRoutingList', params?.search],
    queryFn: ({ pageParam }) => {
      return postV1LtiRoutingList({
        form: {
          limit: params?.limit ?? 100,
          offset: pageParam,
          search: (params?.search?.length || '') < 3 ? '' : params?.search
        }
      });
    },
    refetchOnWindowFocus: false,
    refetchInterval: false,
    gcTime: 0,
    getNextPageParam: (response, pages) => {
      const totalCount = response.result?.total_count || 0;
      const count = pages.flatMap((p) => p.result?.model).length;
      return totalCount && totalCount > count ? count : undefined;
    },
    initialPageParam: 0,
    retry: 3
  });
};

export const useLTIRoutingByID = (id?: number) => {
  return useQuery({
    queryKey: ['postV1LtiRoutingGet', id ?? ''],
    queryFn: () => {
      return id ? postV1LtiRoutingGet({ form: { id: id } }) : undefined;
    },
    retry: 3
  });
};
