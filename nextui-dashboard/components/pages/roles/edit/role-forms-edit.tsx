'use client';
import React from 'react';
import { HouseIcon } from '@/components/icons/breadcrumb/house-icon';
import { RoutesLocation } from '@/components/routes';
import useLanguageBrowser from '@/helpers/locale';
import { CrumbsLayout } from '@/components/layout/crumbs';
import { useParams } from 'react-router-dom';
import { RolesEditForm } from '@/components/pages/roles/edit/form';
import { UsersIcon } from '@/components/icons/breadcrumb/users-icon';
import { RolesIcon } from '@/components/icons/breadcrumb/roles-icon';
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
      icon: <HouseIcon />,
      name: locale.Sidebar.Home,
      href: RoutesLocation.home()
    },
    {
      icon: <UsersIcon />,
      name: locale.Sidebar.Users,
      href: RoutesLocation.accounts()
    },
    {
      icon: <RolesIcon />,
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
