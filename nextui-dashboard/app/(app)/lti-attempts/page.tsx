import React from 'react';
import { LTIAttemptCreate } from '@/components/pages/lti-attempts/create/lti-attempt-create';
import { Layout } from '@/components/layout/layout';
import { LtiAttemptEdit } from '@/components/pages/lti-attempts/edit/lti-attempt-edit';
import { LTIAttemptList } from '@/components/pages/lti-attempts';

export const LTIAttemptsPageCreate = () => {
  return (
    <div className='flex items-center justify-center h-screen bg-gray-100'>
      <div className='relative p-6 max-w-md mx-auto'>
        <LTIAttemptCreate />
      </div>
    </div>
  );
};

export const LTIAttemptsPageUser = () => {
  return (
    <Layout>
      <LTIAttemptList />
    </Layout>
  );
};

export const LTIAttemptsPageEdit = () => {
  return (
    <Layout>
      <LtiAttemptEdit />
    </Layout>
  );
};
