import { useQuery } from '@tanstack/react-query';
import { postV1PnetServerQueueList } from '@/helpers/api';

export const usePnetServerQueueList = () => {
  return useQuery({
    queryKey: ['postV1PnetServerQueueList'],
    queryFn: () => {
      return postV1PnetServerQueueList();
    },
    retry: 1
  });
};
