import { InputProps } from '@heroui/input/dist/input';
import { useQueryUserList } from '@/helpers/queries/user/use-query-user-list';

export const useQueryUserAutoComplete = (search: string, props: InputProps) => {
  const _id = parseInt(props.value || '');
  const queryList = useQueryUserList({
    search: search,
    limit: 20,
    offset: 0
  });
  const queryListItem = useQueryUserList({
    userIds: [_id],
    limit: 1,
    offset: 0
  });
  const entitiesSearch = queryList?.data?.model || [];
  const entityCurrent = queryListItem?.data?.model?.[0];
  const entities =
    entityCurrent && !entitiesSearch.find((entity) => entity.model.id === entityCurrent.model.id)
      ? [...entitiesSearch, entityCurrent]
      : entitiesSearch;
  // Объединяем данные и статусы загрузки
  const isLoading = queryList.isLoading || queryListItem?.isLoading;
  const items = entities.map((entity) => {
    return { key: entity.model.id || 0, value: entity.model.name || '' };
  });

  return { isLoading, items };
};
