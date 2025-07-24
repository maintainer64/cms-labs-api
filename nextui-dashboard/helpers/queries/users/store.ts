import { useMutation, useQuery } from '@tanstack/react-query';
import { getV1GlobalStore, postV1GlobalStore } from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '@/helpers/queries/base';

export const useGlobalStoreGet = () => {
  return useQuery({
    queryKey: ['globalStore'],
    queryFn: getV1GlobalStore,
    retry: 3,
    refetchOnMount: false
  });
};

export const useGlobalStoreSet = (options: TMutationCustomOptions<unknown, unknown, any> = {}) => {
  return useMutation<unknown, unknown, any>({
    // @ts-ignore
    mutationFn: (values: any) => {
      if (values === null) return null;
      queryClient.setQueryData(['globalStore'], () => values);
      return postV1GlobalStore({ form: values });
    },
    ...options
  });
};
