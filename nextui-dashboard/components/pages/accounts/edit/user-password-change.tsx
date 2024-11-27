'use client';
import React from 'react';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { SettingsIcon } from '@/components/icons/sidebar/settings-icon';
import { HomeIcon } from '@/components/icons/sidebar/home-icon';
import { ProfilePasswordChangeForm } from '@/components/pages/accounts/edit/form-recover';

export const ProfilePasswordChange = () => {
  const { locale } = useLanguageBrowser();
  const {
    locale: {
      Tables: { UsersTable }
    }
  } = useLanguageBrowser();
  const crumbs = [
    {
      icon: <HomeIcon />,
      name: locale.Sidebar.Home,
      href: RoutesLocation.home()
    },
    {
      icon: <SettingsIcon />,
      name: locale.Sidebar.Profile,
      href: RoutesLocation.profileChangePassword()
    },
    {
      icon: undefined,
      name: locale.UserNavBar.PasswordChange,
      href: '#'
    }
  ];
  return (
    <CrumbsLayout name={locale.UserNavBar.PasswordChange} crumbs={crumbs}>
      <div className='max-w-[95rem] mx-auto w-full'>
        <ProfilePasswordChangeForm />
      </div>
    </CrumbsLayout>
  );
};
