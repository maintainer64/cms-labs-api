import React from 'react';
import { LTIAttemptCreate } from '@/components/pages/lti-attempts/create/lti-attempt-create';

export const LTIAttemptsPageCreate = () => {
  return (
    <div className='flex items-center justify-center h-screen bg-gray-100'>
      <div className='relative p-6 max-w-md mx-auto'>
        <LTIAttemptCreate />
      </div>
    </div>
  );
};
