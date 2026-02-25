'use client';
import React from 'react';
import { House, UserRoundKey, UsersRound } from 'lucide-react';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { useParams } from 'react-router-dom';
import { RolesEditForm } from '@/components/pages/roles/edit/form';
import { RoleBasedAccess } from '@/components/layout/roleBasedAccess';
import { UserRoleBase } from '@/helpers/queries/sso/auth';

export const RolesFormsEdit = () => {
  const { id } = useParams();
  const { locale } = useLanguageBrowser();
  const {
    locale: {
      Tables: { RoleTable }
    }
  } = useLanguageBrowser();
  const crumbs = [
    {
      icon: <House className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.Home,
      href: RoutesLocation.home()
    },
    {
      icon: <UsersRound className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.Users,
      href: RoutesLocation.accounts()
    },
    {
      icon: <UserRoundKey className='w-5 h-5 stroke-[#969696]' />,
      name: locale.Sidebar.Roles,
      href: RoutesLocation.roles()
    },
    {
      icon: undefined,
      name: locale.Sidebar.AnyList,
      href: '#'
    }
  ];
  return (
    <CrumbsLayout name={RoleTable.Title} crumbs={crumbs}>
      <div className='max-w-[95rem] mx-auto w-full'>
        <RoleBasedAccess allowedRoles={[UserRoleBase.Admin]}>
          <RolesEditForm id={parseInt(id ?? '', 10)} />
        </RoleBasedAccess>
      </div>
    </CrumbsLayout>
  );
};
