import React from 'react';
import { Layout } from '@/components/layout/layout';
import { PnetServersList } from '../../../components/pages/pnet-servers';
import { PnetFormsEdit } from '@/components/pages/pnet-servers/edit/pnet-forms-edit';

export const PnetServersPage = () => {
  return (
    <Layout>
      <PnetServersList />
    </Layout>
  );
};

export const PnetServersPageEdit = () => {
  return (
    <Layout>
      <PnetFormsEdit />
    </Layout>
  );
};
