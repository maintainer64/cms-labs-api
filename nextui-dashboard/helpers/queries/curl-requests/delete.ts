import { useMutation } from '@tanstack/react-query';
import {
  postV1CurlRequestDelete,
  type PostV1CurlRequestDeleteResponse,
  usecases_CurlRequestDeleteInputDTO
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '../base';

export const useCurlRequestDelete = (
  options: TMutationCustomOptions<PostV1CurlRequestDeleteResponse, unknown, usecases_CurlRequestDeleteInputDTO> = {}
) => {
  return useMutation<PostV1CurlRequestDeleteResponse, unknown, usecases_CurlRequestDeleteInputDTO>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ id }: usecases_LTIRoutingDeleteInputDTO) => {
      if (!id) return null;
      return postV1CurlRequestDelete({
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
      await queryClient.invalidateQueries({ queryKey: ['postV1CurlRequestList'] });
      await queryClient.invalidateQueries({ queryKey: ['postV1CurlRequestGet'] });
    }
  });
};
