'use client';
import React from 'react';
import { House, Server } from 'lucide-react';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { useParams } from 'react-router-dom';
import { ServersEditForm } from '@/components/pages/servers/edit/form';
import { UserRoleBase } from '@/helpers/queries/sso/auth';
import { RoleBasedAccess } from '@/components/layout/roleBasedAccess';

export const ServerFormsEdit = () => {
  const { id } = useParams();
  const { locale } = useLanguageBrowser();
  const {
    locale: {
      Tables: { ServersTable }
    }
  } = useLanguageBrowser();
  const crumbs = [
    {
      icon: <House className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.Home,
      href: RoutesLocation.home()
    },
    {
      icon: <Server className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.Servers,
      href: RoutesLocation.servers()
    },
    {
      icon: undefined,
      name: locale.Sidebar.Edit,
      href: '#'
    }
  ];
  return (
    <CrumbsLayout name={ServersTable.Title} crumbs={crumbs}>
      <div className='max-w-[95rem] mx-auto w-full'>
        <RoleBasedAccess allowedRoles={[UserRoleBase.Admin]}>
          <ServersEditForm id={parseInt(id ?? '', 10)} />
        </RoleBasedAccess>
      </div>
    </CrumbsLayout>
  );
};
