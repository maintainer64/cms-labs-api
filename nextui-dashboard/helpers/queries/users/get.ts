import { useQuery, useInfiniteQuery } from '@tanstack/react-query';
import { postV1UserList, postV1UserGet, usecases_UserListInputDTO } from '@/helpers/api';

export const useUsersList = (params?: usecases_UserListInputDTO) => {
  return useInfiniteQuery({
    queryKey: ['postV1UserList', params?.search],
    queryFn: ({ pageParam }) => {
      return postV1UserList({
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

export const useUserByID = (id?: number) => {
  return useQuery({
    queryKey: ['postV1UserGet', id ?? 0],
    queryFn: () => {
      return id ? postV1UserGet({ form: { id: id } }) : undefined;
    },
    retry: 3
  });
};
