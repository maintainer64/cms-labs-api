import { useQuery } from '@tanstack/react-query';
import { postV1TasksList } from '@/helpers/api';
import { InputProps } from '@heroui/input/dist/input';

export const useClabgateTaskList = () => {
  return useQuery({
    queryKey: ['postV1TasksList'],
    queryFn: () => {
      return postV1TasksList({ form: {} });
    },
    retry: 1,
    gcTime: Infinity,
    staleTime: Infinity,
    refetchOnMount: false
  });
};

export const clabgateTaskListData = (search: string, props: InputProps) => {
  const _id = props.value || '';
  const query = useClabgateTaskList();

  const allItems = query?.data?.result?.model || [];

  const entityCurrent = allItems.find((entity) => entity.id === _id);
  const entitiesSearch =
    search.length > 3
      ? allItems.filter(
          (entity) => entity.title?.includes(search) || entity.full_path?.includes(search) || entity.id === search
        )
      : allItems;
  const entities =
    entityCurrent && !entitiesSearch.find((entity) => entity.id === entityCurrent.id)
      ? [...entitiesSearch, entityCurrent]
      : entitiesSearch;
  const items = entities.map((entity) => {
    return { key: entity.id || 0, value: entity.title || entity.full_path || entity.id || '' };
  });

  return {
    isLoading: query.isLoading,
    items
  };
};
