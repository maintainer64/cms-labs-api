import React from 'react';
import { Layout } from '@/components/layout/layout';
import { LTIRoutingList } from '@/components/pages/lti-routings';
import { LTIRoutingEdit } from '@/components/pages/lti-routings/edit/lti-forms-edit';

export const LTIRoutingPage = () => {
  return (
    <Layout>
      <LTIRoutingList />
    </Layout>
  );
};

export const LTIRoutingPageEdit = () => {
  return (
    <Layout>
      <LTIRoutingEdit />
    </Layout>
  );
};
