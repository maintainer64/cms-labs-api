import { useMutation } from '@tanstack/react-query';
import {
  postV1LtiRoutingDelete,
  type PostV1LtiRoutingDeleteResponse,
  usecases_LTIRoutingDeleteInputDTO
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '../base';

export const useLTIRoutingDelete = (
  options: TMutationCustomOptions<PostV1LtiRoutingDeleteResponse, unknown, usecases_LTIRoutingDeleteInputDTO> = {}
) => {
  return useMutation<PostV1LtiRoutingDeleteResponse, unknown, usecases_LTIRoutingDeleteInputDTO>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ id }: usecases_LTIRoutingDeleteInputDTO) => {
      if (!id) return null;
      return postV1LtiRoutingDelete({
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
      await queryClient.invalidateQueries({ queryKey: ['postV1LtiRoutingList'] });
      await queryClient.invalidateQueries({ queryKey: ['postV1LtiRoutingGet'] });
    }
  });
};
