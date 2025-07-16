import { useQuery } from '@tanstack/react-query';
import { getV1TokensJson, postV1TopologiesCreate, postV1TopologiesGet } from '@/helpers/api';

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

export const useTopologyTokenJson = () => {
  return useQuery({
    queryKey: ['getV1TokensJson'],
    queryFn: () => {
      return getV1TokensJson({ form: {} });
    },
    retry: 0
  });
};
