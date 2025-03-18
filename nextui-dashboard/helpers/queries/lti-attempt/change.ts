import { TMutationCustomOptions } from '@/helpers/queries/types';
import {
  postV1LtiAttemptCreate,
  type PostV1LtiAttemptCreateResponse,
  usecases_LTIAttemptCreateInputDTO
} from '@/helpers/api';
import { useMutation } from '@tanstack/react-query';
import queryClient from '@/helpers/queries/base';

export const useLTIAttemptChange = (
  options: TMutationCustomOptions<PostV1LtiAttemptCreateResponse, unknown, usecases_LTIAttemptCreateInputDTO> = {}
) => {
  return useMutation<PostV1LtiAttemptCreateResponse, unknown, usecases_LTIAttemptCreateInputDTO>({
    mutationFn: ({ room_number }: usecases_LTIAttemptCreateInputDTO) => {
      return postV1LtiAttemptCreate({
        form: {
          room_number: room_number
        }
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: ['postV1LtiAttemptCreate'] });
    }
  });
};
