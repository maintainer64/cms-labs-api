import { InputProps } from '@heroui/input/dist/input';
import { useQueryTaskList } from '@/helpers/queries/task/use-query-task-list';

export const useQueryTaskAutoComplete = (search: string, props: InputProps) => {
  const _id = props.value || '';
  const query = useQueryTaskList({});

  const allItems = query?.data?.model || [];

  const entityCurrent = allItems.find((entity) => entity.id === _id);
  const entitiesSearch =
    search.length > 3
      ? allItems.filter(
          (entity) => entity.title?.includes(search) || entity.fullPath?.includes(search) || entity.id === search
        )
      : allItems;
  const entities =
    entityCurrent && !entitiesSearch.find((entity) => entity.id === entityCurrent.id)
      ? [...entitiesSearch, entityCurrent]
      : entitiesSearch;
  const items = entities.map((entity) => {
    return { key: entity.id || 0, value: entity.title || entity.fullPath || entity.id || '' };
  });

  return {
    isLoading: query.isLoading,
    items
  };
};
