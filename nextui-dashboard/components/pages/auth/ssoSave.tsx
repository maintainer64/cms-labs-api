import { auth_SSOAuthorizeInputDTO } from '@/helpers/api';
import dayjs from 'dayjs';
import { SSOAuthorizationParams } from '@/components/pages/auth/sso';

interface SSORedirectSave {
  params: auth_SSOAuthorizeInputDTO;
  createdAt: string;
}

const SSO_REDIRECT_PARAMS_KEY = 'sso_redirect_params';

export const SSOAuthorizationSave = (params: auth_SSOAuthorizeInputDTO) => {
  if (params?.redirect_uri && params?.client_id) {
    const payload: SSORedirectSave = {
      params: params,
      createdAt: dayjs().toString()
    };
    localStorage.setItem(SSO_REDIRECT_PARAMS_KEY, JSON.stringify(payload));
  }
};

export const SSOAuthorizationReset = () => {
  localStorage.setItem(SSO_REDIRECT_PARAMS_KEY, '');
};

export const SSOAuthorizationGet = (): auth_SSOAuthorizeInputDTO | null => {
  SSOAuthorizationSave(SSOAuthorizationParams());
  const payloadString = localStorage.getItem(SSO_REDIRECT_PARAMS_KEY);
  try {
    const payload = JSON.parse(payloadString || '') as SSORedirectSave;
    if (
      payload.params.redirect_uri &&
      payload.params.client_id &&
      dayjs().diff(dayjs(payload.createdAt), 'minute') <= 3
    ) {
      return payload.params;
    }
  } catch (err) {
    return null;
  }
  return null;
};
