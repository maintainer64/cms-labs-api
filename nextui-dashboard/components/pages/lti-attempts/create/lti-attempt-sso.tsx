'use client';

import useLanguageBrowser from '@/helpers/locale';
import { ReactNode } from 'react';
import { SSOAuthorizationGet, SSOAuthorizationReset } from '@/components/pages/auth/ssoSave';
import { useUserProfile } from '@/components/providers/auth-jwt/hooks';
import { SSOAuthorizationComplete } from '@/components/pages/auth/sso';
import { useQuerySsoAuthorize } from '@/helpers/queries/sso/use-query-sso-authorize';

interface Props {
  children: ReactNode;
}

export const LTIAttemptSSO = ({ children }: Props) => {
  const { locale } = useLanguageBrowser();

  const user = useUserProfile();
  const params = SSOAuthorizationGet();
  if (user && user.sub && params && params.redirectUri && params.clientId) {
    const authSSO = useQuerySsoAuthorize(params);
    if (authSSO.data) {
      SSOAuthorizationReset();
      window.location.href = SSOAuthorizationComplete(authSSO.data);
    }
    // @ts-ignore
    const text = authSSO.error?.message || locale.SSO.Wait;
    return <div className='text-center text-[25px] font-bold mb-6'>{text}</div>;
  }
  return <>{children}</>;
};
