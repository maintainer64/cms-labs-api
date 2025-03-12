import { useMutation } from '@tanstack/react-query';
import { models_LTIForm, postV1LtiFormUpsert, type PostV1LtiFormUpsertResponse } from '@/helpers/api';
import { TFormikData, TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '../base';

export const useLTIFormsUpsert = (
  options: TMutationCustomOptions<PostV1LtiFormUpsertResponse, unknown, TFormikData<models_LTIForm>> = {}
) => {
  return useMutation<PostV1LtiFormUpsertResponse, unknown, TFormikData<models_LTIForm>>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ values }: TFormikData<models_LTIForm>) => {
      if (values === null) return null;
      return postV1LtiFormUpsert({
        form: {
          auth_login_uri: values.lti_auth_login_uri,
          auth_token_uri: values.lti_auth_token_uri,
          base_uri: values.base_uri,
          client_id: values.lti_client_id,
          deployment_id: values.lti_deployment_id,
          id: values.id,
          key_set_uri: values.key_set_uri,
          name: values.name,
          target_link_uri: values.target_link_uri
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
