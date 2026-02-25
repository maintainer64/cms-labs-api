import { useContext } from 'react';
import { UserProfileContext } from './context';
import { AuthSwaggerSSOTokenPublicData } from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

export const useUserProfile = (): CamelCasedPropertiesDeep<AuthSwaggerSSOTokenPublicData> => {
  const context = useContext(UserProfileContext);
  if (!context || !context.profile) {
    console.error('useUserProfile must be used within a UserProfileProvider');
    return {};
  }
  return context.profile;
};
