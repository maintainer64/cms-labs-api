import { useQueries } from '@tanstack/react-query';
import { postV1UnlFileGet, postV1UnlFileList } from '@/helpers/api';
import { InputProps } from '@heroui/input/dist/input';

export const useUNLFileData = (type: string[], search: string, props: InputProps) => {
  const _id = parseInt(props.value || '');
  const results = useQueries({
    queries: [
      {
        queryKey: ['postV1UnlFileList', search, ...(type ?? [])],
        queryFn: () => postV1UnlFileList({ form: { search, type, limit: 20, offset: 0 } }),
        retry: 3
      },
      {
        queryKey: ['postV1UnlFileGet', _id],
        queryFn: () => (_id ? postV1UnlFileGet({ form: { id: _id } }) : undefined),
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
    return { key: entity.id || 0, value: entity.path || '' };
  });

  return { isLoading, items };
};
