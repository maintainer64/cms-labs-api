import { createContext, ReactNode } from 'react';
import { AuthSwaggerSSOTokenPublicData } from '@/helpers/api';
import AuthLoadingWrapper from '@/components/pages/auth/loader';
import { useQuerySsoUserInfo } from '@/helpers/queries/sso/use-query-sso-userinfo';
import { CamelCasedPropertiesDeep } from 'type-fest';

type UserProfileProfile = {
  profile?: CamelCasedPropertiesDeep<AuthSwaggerSSOTokenPublicData>;
};

type UserProfileProvider = {
  children: ReactNode;
};

export const UserProfileContext = createContext<UserProfileProfile>({});

export const UserProfileProvider = ({ children }: UserProfileProvider) => {
  const { data, isLoading } = useQuerySsoUserInfo();
  if (isLoading) {
    return <AuthLoadingWrapper />;
  }
  return <UserProfileContext.Provider value={{ profile: data }}>{children}</UserProfileContext.Provider>;
};
