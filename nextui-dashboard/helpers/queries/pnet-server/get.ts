import { useInfiniteQuery, useQueries, useQuery } from '@tanstack/react-query';
import { postV1PnetServerGet, postV1PnetServerList, usecases_PNETServerListInputDTO } from '@/helpers/api';
import { InputProps } from '@heroui/input/dist/input';

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

export const usePnetServerByID = (id?: number) => {
  return useQuery({
    queryKey: ['postV1PnetServerGet', id ?? 0],
    queryFn: () => {
      return id ? postV1PnetServerGet({ form: { id: id } }) : undefined;
    },
    retry: 3
  });
};

export const usePnetServerAutocompleteData = (search: string, props: InputProps) => {
  const _id = parseInt(props.value?.toString() || '');
  const results = useQueries({
    queries: [
      {
        queryKey: ['postV1PnetServerList', search],
        queryFn: () =>
          postV1PnetServerList({
            form: {
              limit: 20,
              offset: 0,
              search: search
            }
          }),
        retry: 3
      },
      {
        queryKey: ['postV1PnetServerGet', _id],
        queryFn: () => (_id ? postV1PnetServerGet({ form: { id: _id } }) : undefined),
        retry: 3
      }
    ]
  });

  const entitiesSearch = results?.[0]?.data?.result?.model || [];
  const entityCurrent = results?.[1]?.data?.result?.model;
  const entities =
    entityCurrent && !entitiesSearch.find((entity) => entity.id === entityCurrent.id)
      ? [...entitiesSearch, entityCurrent]
      : entitiesSearch;

  // Объединяем данные и статусы загрузки
  const isLoading = results.some((result) => result.isLoading);
  const items = entities.map((entity) => {
    return { key: entity.id || 0, value: entity.name || '' };
  });

  return { isLoading, items };
};
