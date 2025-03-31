import { auth_SwaggerSSOToken, postV1TokenRenew } from '@/helpers/api/src';
import { ApiRequestOptions } from '@/helpers/api/src/core/ApiRequestOptions'; // Импортируем библиотеку для работы с JWT
const TokenManager = {
  cachedTokens: null as auth_SwaggerSSOToken | null,
  tokenExpiration: null as number | null,

  // Функция для получения токена
  async getToken(options: ApiRequestOptions): Promise<string> {
    if (options.method == 'POST' && options.url === '/v1/token/renew') return '';
    if (options.method == 'POST' && options.url === '/v1/token/login') return '';
    if (options.method == 'POST' && options.url === '/v1/token/logout') return '';
    if (options.method == 'GET' && options.url === '/v1/lti-form/sso') return '';
    const currentTime = Math.floor(Date.now() / 1000); // Текущее время в секундах
    if (this.cachedTokens && this.tokenExpiration && currentTime < this.tokenExpiration) {
      return this.cachedTokens.access_token || '';
    }
    const response = await postV1TokenRenew({ form: {} });
    const accessToken = response.result?.access_token || '';
    const expires = response.result?.expires_in || 0;
    if (accessToken && expires) {
      this.tokenExpiration = expires;
      this.cachedTokens = response.result ?? null;
      return accessToken;
    }
    return '';
  },

  clearToken() {
    this.cachedTokens = null;
    this.tokenExpiration = null;
  }
};
export default TokenManager;
