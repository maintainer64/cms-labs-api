import { useMutation } from '@tanstack/react-query';
import { postV1RoleDelete, type PostV1RoleDeleteResponse, usecases_RoleDeleteInputDTO } from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '../base';

export const useRoleDelete = (
  options: TMutationCustomOptions<PostV1RoleDeleteResponse, unknown, usecases_RoleDeleteInputDTO> = {}
) => {
  return useMutation<PostV1RoleDeleteResponse, unknown, usecases_RoleDeleteInputDTO>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ id }: usecases_RoleDeleteInputDTO) => {
      if (!id) return null;
      return postV1RoleDelete({
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
      await queryClient.invalidateQueries({ queryKey: ['postV1RoleList'] });
    }
  });
};
