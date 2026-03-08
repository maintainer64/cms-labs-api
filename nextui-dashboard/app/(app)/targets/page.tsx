import React from 'react';
import { Layout } from '@/components/layout/layout';
import { TargetsList } from '@/components/pages/targets';
import { TargetEdit } from '@/components/pages/targets/edit/targets-edit';

export const TargetsPage = () => {
  return (
    <Layout>
      <TargetsList />
    </Layout>
  );
};

export const TargetsPageEdit = () => {
  return (
    <Layout>
      <TargetEdit />
    </Layout>
  );
};
