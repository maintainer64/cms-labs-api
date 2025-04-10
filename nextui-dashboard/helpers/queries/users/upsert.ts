import { useMutation } from '@tanstack/react-query';
import { models_User, postV1UserUpsert, PostV1UserUpsertResponse } from '@/helpers/api';
import { TFormikData, TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '../base';
import { UserItem } from '@/helpers/queries/users/model';

export const useUserUpsert = (
  options: TMutationCustomOptions<PostV1UserUpsertResponse, unknown, TFormikData<UserItem>> = {}
) => {
  return useMutation<PostV1UserUpsertResponse, unknown, TFormikData<UserItem>>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ values }: TFormikData<UserItem>) => {
      if (values === null) return null;
      return postV1UserUpsert({
        form: {
          email: values.email ?? '',
          group_name: values.group_name,
          id: values.id,
          is_active: !values.deleted_at,
          lti_user_id: values.lti_user_id,
          name: values.name ?? '',
          roles: values.roles?.map((roleId) => parseInt(roleId.toString()))
        }
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: ['postV1UserList'] });
      await queryClient.invalidateQueries({ queryKey: ['postV1UserGet'] });
    }
  });
};
