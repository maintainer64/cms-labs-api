'use client';
import React from 'react';
import { House, KeyRound } from 'lucide-react';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { useParams } from 'react-router-dom';
import { AuthProvidersEditForm } from '@/components/pages/auth-providers/edit/form';

export const AuthProvidersEdit = () => {
  const { id } = useParams();
  const { locale } = useLanguageBrowser();
  const {
    locale: {
      Tables: { AuthProvidersTable }
    }
  } = useLanguageBrowser();
  const crumbs = [
    {
      icon: <House className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.Home,
      href: RoutesLocation.home()
    },
    {
      icon: <KeyRound className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.AuthProviders,
      href: RoutesLocation.authProviders()
    },
    {
      icon: undefined,
      name: locale.Sidebar.Edit,
      href: '#'
    }
  ];
  return (
    <CrumbsLayout name={AuthProvidersTable.Title} crumbs={crumbs}>
      <div className='max-w-[95rem] mx-auto w-full'>
        <AuthProvidersEditForm id={parseInt(id ?? '', 10)} />
      </div>
    </CrumbsLayout>
  );
};
