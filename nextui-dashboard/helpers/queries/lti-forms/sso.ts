import { useQuery } from '@tanstack/react-query';
import { getV1LtiFormSso } from '@/helpers/api';

export const useLTIFormsSSOList = () => {
  return useQuery({
    queryKey: ['getV1LtiFormSso'],
    queryFn: () => {
      return getV1LtiFormSso();
    },
    gcTime: Infinity,
    retry: 3
  });
};
