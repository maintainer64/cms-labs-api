import { useQuery } from '@tanstack/react-query';
import { transportWithAuth } from '@/helpers/api';

interface Props {
  path: string;
}

const staticDocsGetByPath = async ({ path }: Props) => {
  const response = await transportWithAuth.getTransport().get('/static/(docs)' + path);
  return response.data as string;
};

export const useQueryStaticDocsGet = ({ path }: Props) => {
  path = path.startsWith('/') ? path : `/${path}`;
  return useQuery({
    queryKey: ['/static/(docs)/', path],
    queryFn: () => {
      return staticDocsGetByPath({ path });
    },
    retry: 1,
    gcTime: Infinity,
    staleTime: Infinity,
    refetchOnMount: false
  });
};
