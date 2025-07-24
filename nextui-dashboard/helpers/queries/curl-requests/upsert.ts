import { useMutation } from '@tanstack/react-query';
import { models_CurlRequest, postV1CurlRequestUpsert, PostV1CurlRequestUpsertResponse } from '@/helpers/api';
import { TFormikData, TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '../base';
import { curlParsedData } from '@/components/pages/curl-requests/edit/validate';

export const useCurlRequestUpsert = (
  options: TMutationCustomOptions<PostV1CurlRequestUpsertResponse, unknown, TFormikData<models_CurlRequest>> = {}
) => {
  return useMutation<PostV1CurlRequestUpsertResponse, unknown, TFormikData<models_CurlRequest>>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ values }: TFormikData<models_CurlRequest>) => {
      if (values === null) return null;
      const curlParsed = curlParsedData(values.raw);
      const timeout = parseInt(`${values.timeout}`);
      return postV1CurlRequestUpsert({
        form: {
          body: curlParsed.body || '',
          headers: curlParsed.header || {},
          id: values.id,
          method: curlParsed.method || '',
          name: values.name ?? '',
          raw_request: values.raw,
          timeout: timeout,
          url: curlParsed.url
        }
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: ['postV1CurlRequestList'] });
      await queryClient.invalidateQueries({ queryKey: ['postV1CurlRequestGet'] });
    }
  });
};
