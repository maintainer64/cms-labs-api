import React from 'react';
import { Layout } from '@/components/layout/layout';
import { ServerFormsEdit } from '@/components/pages/servers/edit/pnet-forms-edit';
import { ServersList } from '@/components/pages/servers';

export const ServersPage = () => {
  return (
    <Layout>
      <ServersList />
    </Layout>
  );
};

export const ServersPageEdit = () => {
  return (
    <Layout>
      <ServerFormsEdit />
    </Layout>
  );
};
