import React from 'react';
import { Layout } from '@/components/layout/layout';
import { CurlRequestList } from '@/components/pages/curl-requests';
import { CurlRequestEdit } from '@/components/pages/curl-requests/edit/lti-forms-edit';

export const CurlRequestPage = () => {
  return (
    <Layout>
      <CurlRequestList />
    </Layout>
  );
};

export const CurlRequestPageEdit = () => {
  return (
    <Layout>
      <CurlRequestEdit />
    </Layout>
  );
};
