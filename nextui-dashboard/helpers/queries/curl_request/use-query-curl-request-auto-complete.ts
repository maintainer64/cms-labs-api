import { useQueries } from '@tanstack/react-query';
import type { UsecasesCurlRequestListRequest, UsecasesCurlRequestListResponse } from '@/helpers/api';
import { CoreJsonRpcPath, transportWithAuth } from '@/helpers/api';
import { InputProps } from '@heroui/input/dist/input';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesCurlRequestListRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesCurlRequestListResponse['result']>;

export const useQueryCurlRequestAutoComplete = (search: string, props: InputProps) => {
  const _id = parseInt(props.value || '');
  const results = useQueries({
    queries: [
      transportWithAuth.getQueryOptions<Response, Params>(
        CoreJsonRpcPath,
        'curl_request.list',
        {
          search: search,
          limit: 20,
          offset: 0
        },
        {
          retry: 3
        }
      ),
      transportWithAuth.getQueryOptions<Response, Params>(
        CoreJsonRpcPath,
        'curl_request.list',
        {
          ids: [_id],
          limit: 1,
          offset: 0
        },
        {
          retry: 3
        }
      )
    ]
  });

  const entitiesSearch = results?.[0]?.data?.model || [];
  const entityCurrent = (results?.[1]?.data?.model || [])?.[0]; // Получаем первый элемент из списка по ID
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
