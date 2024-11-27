import React from 'react';
import { Layout } from '@/components/layout/layout';
import { LtiFormsEdit } from '@/components/pages/lti-forms/edit/lti-forms-edit';
import { LTIFormsList } from '@/components/pages/lti-forms';

export const LTIFormsPage = () => {
  return (
    <Layout>
      <LTIFormsList />
    </Layout>
  );
};

export const LTIFormsPageEdit = () => {
  return (
    <Layout>
      <LtiFormsEdit />
    </Layout>
  );
};
