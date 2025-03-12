import { useMutation } from '@tanstack/react-query';
import { models_ServiceCard, postV1ServiceCardUpsert, type PostV1ServiceCardUpsertResponse } from '@/helpers/api';
import { TFormikData, TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '../base';

export const useServiceCardUpsert = (
  options: TMutationCustomOptions<PostV1ServiceCardUpsertResponse, unknown, TFormikData<models_ServiceCard>> = {}
) => {
  return useMutation<PostV1ServiceCardUpsertResponse, unknown, TFormikData<models_ServiceCard>>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ values }: TFormikData<models_ServiceCard>) => {
      if (values === null) return null;
      return postV1ServiceCardUpsert({
        form: {
          id: values.id,
          is_active: values.is_active ?? true,
          name: values.name || '',
          description: values.description || '',
          order: values.order || 0,
          url: values.url || '',
          image_url: values.image_url || ''
        }
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: ['postV1ServiceCardList'] });
      await queryClient.invalidateQueries({ queryKey: ['postV1ServiceCardGet'] });
    }
  });
};
