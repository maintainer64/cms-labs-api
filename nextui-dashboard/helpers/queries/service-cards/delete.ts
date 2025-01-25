import { useMutation } from '@tanstack/react-query';
import {
  type PostV1PnetServerDeleteResponse,
  postV1ServiceCardDelete,
  usecases_ServiceCardDeleteInputDTO
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '../base';

export const useServiceCardDelete = (
  options: TMutationCustomOptions<PostV1PnetServerDeleteResponse, unknown, usecases_ServiceCardDeleteInputDTO> = {}
) => {
  return useMutation<PostV1PnetServerDeleteResponse, unknown, usecases_ServiceCardDeleteInputDTO>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ id }: usecases_ServiceCardDeleteInputDTO) => {
      if (!id) return null;
      return postV1ServiceCardDelete({
        form: {
          id: id
        }
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: ['postV1ServiceCardList'] });
      await queryClient.invalidateQueries({ queryKey: ['postV1ServiceCardGet'] });
    }
  });
};
