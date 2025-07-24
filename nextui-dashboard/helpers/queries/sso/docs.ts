import { useQuery } from '@tanstack/react-query';
import { request } from '@/helpers/api/backend/core/request';
import { CancelablePromise, OpenAPI } from '@/helpers/api/backend';

interface Props {
  path: string;
}

export const getV1Docs = ({ path }: Props): CancelablePromise<any> => {
  return request(
    { ...OpenAPI, BASE: '/static/(docs)' },
    {
      method: 'GET',
      url: path
    }
  );
};

export const useDocs = ({ path }: Props) => {
  path = path.startsWith('/') ? path : `/${path}`;
  return useQuery({
    queryKey: ['getV1Docs', path],
    queryFn: () => {
      return getV1Docs({ path });
    },
    retry: 1,
    gcTime: Infinity,
    staleTime: Infinity,
    refetchOnMount: false
  });
};
