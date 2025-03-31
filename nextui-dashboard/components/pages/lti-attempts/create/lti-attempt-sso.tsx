'use client';

import useLanguageBrowser from '@/helpers/locale';
import { ReactNode } from 'react';
import { SSOAuthorizationGet, SSOAuthorizationReset } from '@/components/pages/auth/ssoSave';
import { useUserProfile } from '@/components/providers/auth-jwt/hooks';
import { useSSOAuth } from '@/helpers/queries/sso/auth';
import { SSOAuthorizationComplete } from '@/components/pages/auth/sso';

interface Props {
  children: ReactNode;
}

export const LTIAttemptSSO = ({ children }: Props) => {
  const { locale } = useLanguageBrowser();

  const user = useUserProfile();
  const params = SSOAuthorizationGet();
  if (user && user.sub && params && params.redirect_uri && params.client_id) {
    const authSSO = useSSOAuth(params);
    if (authSSO.data?.result) {
      SSOAuthorizationReset();
      window.location.href = SSOAuthorizationComplete(authSSO.data?.result);
    }
    // @ts-ignore
    const text = authSSO.error?.body?.msg || locale.SSO.Wait;
    return <div className='text-center text-[25px] font-bold mb-6'>{text}</div>;
  }
  return <>{children}</>;
};
