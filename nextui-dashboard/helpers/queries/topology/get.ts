import { useMutation, useQuery } from '@tanstack/react-query';
import {
  getV1TokensJson,
  postV1ContainersGet,
  PostV1ContainersGetResponse,
  PostV1RoleUpsertResponse,
  postV1TopologiesCreate,
  postV1TopologiesGet,
  usecases_ContainersGetInputDTO
} from '@/helpers/api';
import { TMutationCustomOptions } from '@/helpers/queries/types';

export const useTopologyCreate = (username?: string, taskId?: string, redeploy = false) => {
  return useQuery({
    queryKey: ['postV1TopologiesCreate', username ?? '', taskId ?? ''],
    queryFn: () => {
      return taskId && username
        ? postV1TopologiesCreate({
            form: {
              username: username,
              task_id: taskId,
              redeploy: redeploy
            }
          })
        : undefined;
    },
    retry: 0
  });
};

export const useTopologyGet = (namespace?: string) => {
  return useQuery({
    queryKey: ['postV1TopologiesGet', namespace ?? ''],
    queryFn: () => {
      return namespace
        ? postV1TopologiesGet({
            form: {
              namespace: namespace
            }
          })
        : undefined;
    },
    retry: 10,
    retryDelay: 5000
  });
};

export const useContainersGet = (
  options: TMutationCustomOptions<PostV1ContainersGetResponse, unknown, usecases_ContainersGetInputDTO> = {}
) => {
  return useMutation<PostV1ContainersGetResponse, unknown, usecases_ContainersGetInputDTO>({
    // @ts-expect-error: return nullable value
    mutationFn: (params: usecases_ContainersGetInputDTO) => {
      if (params === null) return null;
      return postV1ContainersGet({
        form: params
      });
    },
    ...options
  });
};
