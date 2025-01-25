import { useQuery } from '@tanstack/react-query';
import { postV1LtiAttemptCreate } from '@/helpers/api';


export const useLTIAttemptCreate = () => {
  return useQuery({
    queryKey: ['postV1LtiRoutingGet'],
    queryFn: () => {
      return postV1LtiAttemptCreate({ form: {} });
    },
    retry: 2,
    refetchInterval: 30000  // 30 seconds
  });
};
