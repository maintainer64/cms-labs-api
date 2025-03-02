import { useMutation } from '@tanstack/react-query';
import {
  postV1LtiAttemptDelete,
  type PostV1LtiAttemptDeleteResponse,
  usecases_LTIAttemptDeleteInputDTO
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '../base';

export const useLTIAttemptDelete = (
  options: TMutationCustomOptions<PostV1LtiAttemptDeleteResponse, unknown, usecases_LTIAttemptDeleteInputDTO> = {}
) => {
  return useMutation<PostV1LtiAttemptDeleteResponse, unknown, usecases_LTIAttemptDeleteInputDTO>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ id }: usecases_LTIAttemptDeleteInputDTO) => {
      if (!id) return null;
      return postV1LtiAttemptDelete({
        form: {
          id: id
        }
      });
    },
    ...options,
    async onSuccess(...args) {
      await queryClient.invalidateQueries({ queryKey: ['postV1LtiAttemptGet'] });
      await queryClient.invalidateQueries({ queryKey: ['postV1LtiRoutingGet'] });
      await queryClient.invalidateQueries({ queryKey: ['postV1LtiAttemptList'] });
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
    }
  });
};
