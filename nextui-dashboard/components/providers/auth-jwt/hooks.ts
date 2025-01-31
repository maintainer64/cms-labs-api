import { useContext } from 'react';
import { UserProfileContext } from './context';
import { auth_SSOTokenPublicData } from '@/helpers/api';

export const useUserProfile = (): auth_SSOTokenPublicData => {
  const context = useContext(UserProfileContext);
  if (!context || !context.profile) {
    console.error('useUserProfile must be used within a UserProfileProvider');
    return {};
  }
  return context.profile;
};
