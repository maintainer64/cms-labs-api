import { useQuery, useInfiniteQuery } from '@tanstack/react-query';
import { postV1LtiFormList, postV1LtiFormGet, usecases_LTIFormListInputDTO } from '@/helpers/api';

export const useLTIFormsList = (params?: usecases_LTIFormListInputDTO) => {
  return useInfiniteQuery({
    queryKey: ['postV1LtiFormList', params?.search],
    queryFn: ({ pageParam }) => {
      return postV1LtiFormList({
        form: {
          limit: params?.limit ?? 100,
          offset: pageParam,
          search: (params?.search?.length || '') < 3 ? '' : params?.search
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
    retry: 3
  });
};

export const useLtiFormsByID = (id?: number) => {
  return useQuery({
    queryKey: ['postV1LtiFormGet', id],
    queryFn: () => {
      return id ? postV1LtiFormGet({ form: { id: id } }) : undefined;
    },
    retry: 3
  });
};
