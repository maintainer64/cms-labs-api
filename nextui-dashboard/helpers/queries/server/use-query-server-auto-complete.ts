import { InputProps } from '@heroui/input/dist/input';
import { MapServerItem, useQueryServerGet } from '@/helpers/queries/server/use-query-server-get';
import { useQueryServerList } from '@/helpers/queries/server/use-query-server-list';

export const useQueryServerAutoComplete = (search: string, props: InputProps) => {
  const _id = parseInt(props.value || '');
  const queryList = useQueryServerList({
    search: search,
    limit: 20,
    offset: 0
  });
  const queryItem = useQueryServerGet({ id: _id });
  const entitiesSearch = (queryList?.data?.model || []).map((item) => MapServerItem(item.model, item.roles));
  const entityCurrent = queryItem?.data?.model;
  const entities =
    entityCurrent && !entitiesSearch.find((entity) => entity.id === entityCurrent.id)
      ? [...entitiesSearch, entityCurrent]
      : entitiesSearch;

  // Объединяем данные и статусы загрузки
  const isLoading = queryList.isLoading || queryItem.isLoading;
  const items = entities.map((entity) => {
    return { key: entity.id || 0, value: entity.name || '' };
  });

  return { isLoading, items };
};
