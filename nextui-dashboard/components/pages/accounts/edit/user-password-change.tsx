'use client';
import React from 'react';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { ProfilePasswordChangeForm } from '@/components/pages/accounts/edit/form-recover';
import { Cog, House } from 'lucide-react';

export const ProfilePasswordChange = () => {
  const { locale } = useLanguageBrowser();
  const {
    locale: {
      Tables: { UsersTable }
    }
  } = useLanguageBrowser();
  const crumbs = [
    {
      icon: <House className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.Home,
      href: RoutesLocation.home()
    },
    {
      icon: <Cog className='w-5 h-5 stroke-[#969696]' />,
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
