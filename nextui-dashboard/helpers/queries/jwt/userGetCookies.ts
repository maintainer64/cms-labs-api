import { useQuery } from '@tanstack/react-query';
import { getV1SsoUserinfo } from '@/helpers/api';

export const userGetCookies = () => {
  return useQuery({
    queryKey: ['userGetCookies'],
    queryFn: () => {
      return getV1SsoUserinfo({ authorization: '' });
    },
    retry: 0
  });
};
