import React from 'react';
import { Accounts } from '../../../components/pages/accounts';
import { Layout } from '@/components/layout/layout';
import { AccountsEdit } from '@/components/pages/accounts/edit/accounts-edit';
import { ProfilePasswordChange } from '@/components/pages/accounts/edit/user-password-change';

export const AccountsPage = () => {
  return (
    <Layout>
      <Accounts />
    </Layout>
  );
};

export const AccountsPageEdit = () => {
  return (
    <Layout>
      <AccountsEdit />
    </Layout>
  );
};

export const ProfilePagePasswordChange = () => {
  return (
    <Layout>
      <ProfilePasswordChange />
    </Layout>
  );
};
