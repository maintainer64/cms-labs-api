import { useMutation } from '@tanstack/react-query';
import { models_Role, postV1RoleUpsert, PostV1RoleUpsertResponse } from '@/helpers/api';
import { TFormikData, TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '../base';

export const useRoleUpsert = (
  options: TMutationCustomOptions<PostV1RoleUpsertResponse, unknown, TFormikData<models_Role>> = {}
) => {
  return useMutation<PostV1RoleUpsertResponse, unknown, TFormikData<models_Role>>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ values }: TFormikData<models_Role>) => {
      if (values === null) return null;
      return postV1RoleUpsert({
        form: {
          id: values.id,
          name: values.name,
          code: values.code
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
