import { useMutation } from '@tanstack/react-query';
import { postV1LtiFormDelete, type PostV1LtiFormDeleteResponse, usecases_LTIFormDeleteInputDTO } from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '../base';

export const useLTIFormsDelete = (
  options: TMutationCustomOptions<PostV1LtiFormDeleteResponse, unknown, usecases_LTIFormDeleteInputDTO> = {}
) => {
  return useMutation<PostV1LtiFormDeleteResponse, unknown, usecases_LTIFormDeleteInputDTO>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ id }: usecases_LTIFormDeleteInputDTO) => {
      if (!id) return null;
      return postV1LtiFormDelete({
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
      await queryClient.invalidateQueries({ queryKey: ['postV1LtiFormGet'] });
      await queryClient.invalidateQueries({ queryKey: ['postV1LtiFormList'] });
    }
  });
};
