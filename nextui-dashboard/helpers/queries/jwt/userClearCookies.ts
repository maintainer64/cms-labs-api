import { postV1TokenLogout } from '@/helpers/api';
import queryClient from '@/helpers/queries/base';
import TokenManager from '@/helpers/api/axios-refresh';

export const userClearCookies = async () => {
  await postV1TokenLogout();
  await queryClient.invalidateQueries({ queryKey: ['userGetCookies'] });
  TokenManager.clearToken();
};
