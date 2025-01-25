import { useMutation } from '@tanstack/react-query';
import { models_PNETServer, postV1PnetServerUpsert, PostV1PnetServerUpsertResponse } from '@/helpers/api';
import { TFormikData, TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '../base';

export const usePnetServerUpsert = (
  options: TMutationCustomOptions<PostV1PnetServerUpsertResponse, unknown, TFormikData<models_PNETServer>> = {}
) => {
  return useMutation<PostV1PnetServerUpsertResponse, unknown, TFormikData<models_PNETServer>>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ values }: TFormikData<models_PNETServer>) => {
      if (values === null) return null;
      return postV1PnetServerUpsert({
        form: {
          id: values.id,
          is_active: values.is_active,
          name: values.name || '',
          minutes_for_disconnect: values.minutes_for_disconnect || 0,
          max_count_users_limit: values.max_count_users_limit || 0,
          unit_rate: values.unit_rate,
          url: values.url || ''
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
