import { useQuery } from '@tanstack/react-query';
import { postV1ServiceCardGet, postV1ServiceCardList } from '@/helpers/api';

export const useServiceCardList = () => {
  return useQuery({
    queryKey: ['postV1ServiceCardList'],
    queryFn: () => {
      return postV1ServiceCardList({
        form: {
          limit: 100,
          offset: 0
        }
      });
    },
    retry: 3
  });
};

export const useServiceCardById = (id?: number) => {
  return useQuery({
    queryKey: ['postV1ServiceCardGet', id],
    queryFn: () => {
      return id ? postV1ServiceCardGet({ form: { id: id } }) : undefined;
    },
    retry: 3
  });
};
