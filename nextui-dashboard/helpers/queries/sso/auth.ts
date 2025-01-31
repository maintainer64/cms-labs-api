import { useQuery } from '@tanstack/react-query';
import { auth_SSOAuthorizeInputDTO, postV1SsoAuthorize } from '@/helpers/api';

export const useSSOAuth = (params: auth_SSOAuthorizeInputDTO) => {
  return useQuery({
    queryKey: ['postV1SsoAuthorize', params],
    queryFn: () => {
      return postV1SsoAuthorize({ form: params });
    },
    retry: 1
  });
};
