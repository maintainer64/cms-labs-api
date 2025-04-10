import { useQuery } from '@tanstack/react-query';
import {
  postV1LtiAttemptCreate,
  postV1LtiAttemptGet,
  postV1LtiAttemptList,
  usecases_LTIAttemptListInputDTO
} from '@/helpers/api';

export const useLTIAttemptCreate = () => {
  return useQuery({
    queryKey: ['postV1LtiAttemptCreate'],
    queryFn: () => {
      return postV1LtiAttemptCreate({ form: {} });
    },
    retry: 2,
    refetchInterval: 10000 // 10 seconds
  });
};

export const useLTIAttemptById = (id?: number) => {
  return useQuery({
    queryKey: ['postV1LtiAttemptGet', id ?? 0],
    queryFn: () => {
      return id ? postV1LtiAttemptGet({ form: { id: id } }) : undefined;
    },
    retry: 3
  });
};

export const useLTIAttemptList = (params?: usecases_LTIAttemptListInputDTO) => {
  return useQuery({
    queryKey: ['postV1LtiAttemptList', params?.limit, params?.offset, params?.user_ids],
    queryFn: () => {
      return postV1LtiAttemptList({
        form: {
          limit: params?.limit ?? 100,
          offset: params?.offset ?? 0,
          user_ids: params?.user_ids ?? []
        }
      });
    },
    retry: 3
  });
};
