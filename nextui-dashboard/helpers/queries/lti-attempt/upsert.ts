import { useMutation } from '@tanstack/react-query';
import { models_LTIAttempt, postV1LtiAttemptEdit, type PostV1LtiAttemptEditResponse } from '@/helpers/api';
import { TFormikData, TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '../base';

export const useLTIAttemptUpsert = (
  options: TMutationCustomOptions<PostV1LtiAttemptEditResponse, unknown, TFormikData<models_LTIAttempt>> = {}
) => {
  return useMutation<PostV1LtiAttemptEditResponse, unknown, TFormikData<models_LTIAttempt>>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ values }: TFormikData<models_LTIAttempt>) => {
      if (values === null) return null;
      return postV1LtiAttemptEdit({
        form: {
          expired_at: values.expired_at,
          id: values.id,
          pnet_server_id: parseInt(values.pnet_server_id?.toString() || '')
        }
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: ['postV1LtiAttemptGet'] });
      await queryClient.invalidateQueries({ queryKey: ['postV1LtiRoutingGet'] });
      await queryClient.invalidateQueries({ queryKey: ['postV1LtiAttemptList'] });
    }
  });
};
