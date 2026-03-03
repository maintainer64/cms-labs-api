import { RpcTransport } from './request';
import { transportWithoutAuth } from '@/helpers/api/core/transportWithoutAuth';

export const transportWithAuth = new RpcTransport();
transportWithAuth.setupAuth({
  getAuthData: transportWithoutAuth.userTokenAccess.bind(transportWithoutAuth),
  updateAuthData: transportWithoutAuth.userTokenRefresh.bind(transportWithoutAuth)
});
