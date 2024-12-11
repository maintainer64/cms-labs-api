import { useMutation } from '@tanstack/react-query';
import { models_LTIRouting, postV1LtiRoutingUpsert, PostV1LtiRoutingUpsertResponse } from '@/helpers/api';
import { TFormikData, TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '../base';

export const useLTIRoutingUpsert = (
  options: TMutationCustomOptions<PostV1LtiRoutingUpsertResponse, unknown, TFormikData<models_LTIRouting>> = {}
) => {
  return useMutation<PostV1LtiRoutingUpsertResponse, unknown, TFormikData<models_LTIRouting>>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ values }: TFormikData<models_LTIRouting>) => {
      if (values === null) return null;
      return postV1LtiRoutingUpsert({
        form: {
          collaboration: values.collaboration || 0,
          id: values.id,
          lti_description: values.lti_description,
          lti_params_task: values.lti_params_task,
          lti_task_id: values.lti_task_id,
          lti_title: values.lti_title,
          name: values.name,
          pinned_session_minutes: values.pinned_session_minutes || 0,
          pnet_labs_path: values.pnet_labs_path,
          pnet_test_path: values.pnet_test_path
        }
      });
    },
    ...options,
    async onSuccess(...args) {
      await queryClient.invalidateQueries({ queryKey: ['postV1LtiRoutingList'] });
      await queryClient.invalidateQueries({ queryKey: ['postV1LtiRoutingGet'] });
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
    }
  });
};
