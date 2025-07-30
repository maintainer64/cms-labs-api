import { useMutation } from '@tanstack/react-query';
import { postV1PnetServerUpsert, PostV1PnetServerUpsertResponse } from '@/helpers/api';
import { TFormikData, TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '../base';
import { PnetServerItem } from '@/helpers/queries/pnet-server/model';

export const usePnetServerUpsert = (
  options: TMutationCustomOptions<PostV1PnetServerUpsertResponse, unknown, TFormikData<PnetServerItem>> = {}
) => {
  return useMutation<PostV1PnetServerUpsertResponse, unknown, TFormikData<PnetServerItem>>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ values }: TFormikData<PnetServerItem>) => {
      if (values === null) return null;
      return postV1PnetServerUpsert({
        form: {
          id: values.id,
          is_active: values.is_active,
          name: values.name || '',
          minutes_for_disconnect: values.minutes_for_disconnect || 0,
          max_count_users_limit: values.max_count_users_limit || 0,
          unit_rate: values.unit_rate,
          url: values.url || '',
          type: values.type || 'pnet',
          client_id: values.client_id || '',
          roles: values.roles?.map((roleId) => parseInt(roleId.toString()))
        }
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: ['postV1PnetServerList'] });
      await queryClient.invalidateQueries({ queryKey: ['postV1PnetServerGet'] });
      await queryClient.invalidateQueries({ queryKey: ['postV1PnetServerQueueList'] });
    }
  });
};
