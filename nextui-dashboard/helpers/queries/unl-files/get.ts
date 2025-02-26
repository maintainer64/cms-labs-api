import { useQueries } from '@tanstack/react-query';
import { postV1UnlFileGet, postV1UnlFileList } from '@/helpers/api';

export const useUNLFileData = (search?: string, type?: string[], id?: number) => {
  const results = useQueries({
    queries: [
      {
        queryKey: ['postV1UnlFileList', search, ...(type ?? [])],
        queryFn: () => postV1UnlFileList({ form: { search, type, limit: 20, offset: 0 } }),
        retry: 3
      },
      {
        queryKey: ['postV1UnlFileGet', id],
        queryFn: () => (id ? postV1UnlFileGet({ form: { id } }) : undefined),
        retry: 3
      }
    ]
  });

  const filesSearch = results?.[0]?.data?.result?.model || [];
  const fileCurrent = results?.[1]?.data?.result?.model;
  const files =
    fileCurrent && !filesSearch.find((file) => file.id === fileCurrent.id)
      ? [...filesSearch, fileCurrent]
      : filesSearch;

  // Объединяем данные и статусы загрузки
  const isLoading = results.some((result) => result.isLoading);

  return { isLoading, files };
};
