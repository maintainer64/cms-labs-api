import { useMutation } from '@tanstack/react-query';
import {
  postV1PnetServerDelete,
  type PostV1PnetServerDeleteResponse,
  usecases_PNETServerDeleteInputDTO
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '../base';

export const usePnetServerDelete = (
  options: TMutationCustomOptions<PostV1PnetServerDeleteResponse, unknown, usecases_PNETServerDeleteInputDTO> = {}
) => {
  return useMutation<PostV1PnetServerDeleteResponse, unknown, usecases_PNETServerDeleteInputDTO>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ id }: usecases_PNETServerDeleteInputDTO) => {
      if (!id) return null;
      return postV1PnetServerDelete({
        form: {
          id: id
        }
      });
    },
    ...options,
    async onSuccess(...args) {
      await queryClient.invalidateQueries({ queryKey: ['postV1PnetServerList'] });
      await queryClient.invalidateQueries({ queryKey: ['postV1PnetServerGet'] });
      await queryClient.invalidateQueries({ queryKey: ['postV1PnetServerQueueList'] });
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
    }
  });
};
