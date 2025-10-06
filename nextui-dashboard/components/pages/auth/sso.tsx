import { AuthSSOAuthorizeInputDTO, AuthSSOAuthorizeOutputDTO } from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

export const SSOAuthorizationParams = () => {
  const currentURI = new URL(window.location.href);
  const params: CamelCasedPropertiesDeep<AuthSSOAuthorizeInputDTO> = {
    clientId: currentURI.searchParams.get('client_id') || '',
    path: currentURI.searchParams.get('path') || '',
    redirectUri: currentURI.searchParams.get('redirect_uri') || '',
    responseType: currentURI.searchParams.get('response_type') || '',
    scope: currentURI.searchParams.get('scope') || '',
    state: currentURI.searchParams.get('state') || '',
    nonce: currentURI.searchParams.get('nonce') || '',
    extra: currentURI.searchParams.get('extra') || ''
  };
  return params;
};

export const SSOAuthorizationComplete = (params?: CamelCasedPropertiesDeep<AuthSSOAuthorizeOutputDTO>) => {
  if (!params) return '';
  return `${params.redirectUri || '/'}?state=${params.state || ''}&path=${params.path || ''}&code=${params.code || ''}&application=${params.application || ''}&extra=${params.extra || ''}&nonce=${params.nonce || ''}&client_id=${params.clientId || ''}&scope=${params.scope}`;
};
