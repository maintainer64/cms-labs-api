import { auth_SSOAuthorizeInputDTO, auth_SSOAuthorizeOutputDTO } from '@/helpers/api';

export const SSOAuthorizationParams = () => {
  const currentURI = new URL(window.location.href);
  const params: auth_SSOAuthorizeInputDTO = {
    client_id: currentURI.searchParams.get('client_id') || '',
    path: currentURI.searchParams.get('path') || '',
    redirect_uri: currentURI.searchParams.get('redirect_uri') || '',
    response_type: currentURI.searchParams.get('response_type') || '',
    scope: currentURI.searchParams.get('scope') || '',
    state: currentURI.searchParams.get('state') || '',
    nonce: currentURI.searchParams.get('nonce') || '',
    extra: currentURI.searchParams.get('extra') || ''
  };
  return params;
};

export const SSOAuthorizationComplete = (params?: auth_SSOAuthorizeOutputDTO) => {
  if (!params) return '';
  return `${params.redirect_uri || '/'}?state=${params.state || ''}&path=${params.path || ''}&code=${params.code || ''}&application=${params.application || ''}&extra=${params.extra || ''}&nonce=${params.nonce || ''}&client_id=${params.client_id || ''}&scope=${params.scope}`;
};
