import { useQuery } from '@tanstack/react-query';
import { postV1RoleList } from '@/helpers/api';

export const useRolesList = () => {
  return useQuery({
    queryKey: ['postV1RoleList'],
    queryFn: () => {
      return postV1RoleList({ form: {} });
    },
    refetchOnWindowFocus: true,
    refetchOnMount: true,
    gcTime: 0,
    retry: 3
  });
};
