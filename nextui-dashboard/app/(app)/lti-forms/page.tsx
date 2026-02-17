import React from 'react';
import { Layout } from '@/components/layout/layout';
import { AuthProvidersEdit } from '@/components/pages/auth-providers/edit/auth-providers-edit';
import { AuthProvidersList } from '@/components/pages/auth-providers';

export const AuthProvidersPage = () => {
  return (
    <Layout>
      <AuthProvidersList />
    </Layout>
  );
};

export const AuthProvidersPageEdit = () => {
  return (
    <Layout>
      <AuthProvidersEdit />
    </Layout>
  );
};
