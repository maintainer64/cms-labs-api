import React from 'react';
import { LTIAttemptCreate } from '@/components/pages/lti-attempts/create/lti-attempt-create';
import { Layout } from '@/components/layout/layout';
import { LtiAttemptEdit } from '@/components/pages/lti-attempts/edit/lti-attempt-edit';
import { LTIAttemptConfirm } from '@/components/pages/lti-attempts/create/lti-attempt-confirm';
import { LTIAttemptSSO } from '@/components/pages/lti-attempts/create/lti-attempt-sso';
import { LTIAttemptsListPage } from '@/components/pages/lti-attempts/list/lti-attempts-list';

export const LTIAttemptsPageCreate = () => {
  return (
    <div className='flex items-center justify-center h-screen bg-background'>
      <div className='relative p-6 max-w-md mx-auto'>
        <LTIAttemptSSO>
          <LTIAttemptCreate />
        </LTIAttemptSSO>
      </div>
    </div>
  );
};

export const LTIAttemptPageConfirm = () => {
  return (
    <div className='flex items-center justify-center h-screen bg-background'>
      <div className='relative p-6 max-w-md mx-auto'>
        <LTIAttemptSSO>
          <LTIAttemptConfirm />
        </LTIAttemptSSO>
      </div>
    </div>
  );
};

export const LTIAttemptsPageEdit = () => {
  return (
    <Layout>
      <LtiAttemptEdit />
    </Layout>
  );
};

export const LTIAttemptsPage = () => {
  return (
    <Layout>
      <LTIAttemptsListPage />
    </Layout>
  );
};
