import { useMutation } from '@tanstack/react-query';
import {
  postV1TopologiesDelete,
  type PostV1TopologiesDeleteResponse,
  usecases_TopologiesDeleteInputDTO
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';
import queryClient from '../base';

export const useTopologyDelete = (
  options: TMutationCustomOptions<PostV1TopologiesDeleteResponse, unknown, usecases_TopologiesDeleteInputDTO> = {}
) => {
  return useMutation<PostV1TopologiesDeleteResponse, unknown, usecases_TopologiesDeleteInputDTO>({
    // @ts-expect-error: return nullable value
    mutationFn: ({ namespaces }: usecases_TopologiesDeleteInputDTO) => {
      if (!namespaces) return null;
      return postV1TopologiesDelete({
        form: { namespaces }
      });
    },
    ...options,
    async onSuccess(...args) {
      if (options.onSuccess) {
        options.onSuccess(...args);
      }
      await queryClient.invalidateQueries({ queryKey: ['postV1TopologiesCreate'] });
      await queryClient.invalidateQueries({ queryKey: ['postV1TopologiesGet'] });
      await queryClient.invalidateQueries({ queryKey: ['getV1TokensJson'] });
    }
  });
};
