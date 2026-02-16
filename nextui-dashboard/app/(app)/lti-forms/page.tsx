import React from 'react';
import { Layout } from '@/components/layout/layout';
import { AuthProvidersEdit } from '@/components/pages/lti-forms/edit/lti-forms-edit';
import { AuthProvidersList } from '@/components/pages/lti-forms';

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
