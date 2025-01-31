import { useQuery } from '@tanstack/react-query';
import { postV1SsoUserinfo } from '@/helpers/api';

export const userGetCookies = () => {
  return useQuery({
    queryKey: ['userGetCookies'],
    queryFn: () => {
      return postV1SsoUserinfo({ authorization: '' });
    },
    retry: 0
  });
};
