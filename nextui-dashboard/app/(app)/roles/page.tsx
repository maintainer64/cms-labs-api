import React from 'react';
import { Layout } from '@/components/layout/layout';
import { RolesList } from '@/components/pages/roles';
import { PnetFormsEdit } from '@/components/pages/roles/edit/pnet-forms-edit';

export const RolesListPage = () => {
  return (
    <Layout>
      <RolesList />
    </Layout>
  );
};

export const RolesPageEdit = () => {
  return (
    <Layout>
      <PnetFormsEdit />
    </Layout>
  );
};
