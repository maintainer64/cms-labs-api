import { useInfiniteQuery, useQueries, useQuery } from '@tanstack/react-query';
import { postV1CurlRequestGet, postV1CurlRequestList, usecases_CurlRequestListInputDTO } from '@/helpers/api';
import { InputProps } from '@heroui/input/dist/input';

export const curlRequestsData = (search: string, props: InputProps) => {
  const _id = parseInt(props.value || '');
  const results = useQueries({
    queries: [
      {
        queryKey: ['postV1UnlFileList', search, null],
        queryFn: () => postV1CurlRequestList({ form: { search, limit: 20, offset: 0 } }),
        retry: 3
      },
      {
        queryKey: ['postV1CurlRequestList', null, _id],
        queryFn: () => (_id ? postV1CurlRequestList({ form: { ids: [_id], limit: 1, offset: 0 } }) : undefined),
        retry: 3
      }
    ]
  });

  const entitiesSearch = results?.[0]?.data?.result?.model || [];
  const entityCurrent = (results?.[1]?.data?.result?.model || [])?.[0]; // Получаем первый элемент из списка по ID
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

export const useCurlRequestList = (params?: usecases_CurlRequestListInputDTO) => {
  return useInfiniteQuery({
    queryKey: ['postV1CurlRequestList', params?.search],
    queryFn: ({ pageParam }) => {
      return postV1CurlRequestList({
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

export const useCurlRequestByID = (id?: number) => {
  return useQuery({
    queryKey: ['postV1CurlRequestGet', id ?? ''],
    queryFn: () => {
      return id ? postV1CurlRequestGet({ form: { id: id } }) : undefined;
    },
    retry: 3
  });
};
