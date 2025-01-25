'use client';

import { useLTIAttemptCreate } from '@/helpers/queries/lti-attempt/get';
import { Loading } from '@/components/scroll/loader';
import React from 'react';
import { ErrorModal } from '@/components/pages/auth/error';
import useLanguageBrowser from '@/helpers/locale';
import { Link } from '@nextui-org/react';
import { MyComponent } from '@/components/pages/lti-attempts/create/components';

export const LTIAttemptCreate = () => {
  const {
    locale: {
      PnetServersQueue: { LTIAttemptRoom }
    }
  } = useLanguageBrowser();
  return <MyComponent />;
  const response = useLTIAttemptCreate();
  if (response.isLoading) return <Loading size={8} />;
  if (response.error)
    return (
      <ErrorModal title={LTIAttemptRoom.ErrorPageTitle} description={response.error.body.msg}>
        <Link onClick={() => response.refetch()} href='#' underline='always'>
          {LTIAttemptRoom.ErrorPageRefresh}
        </Link>
      </ErrorModal>
    );
  const result = response.data?.result;
  if (result?.auto_redirect === true) {
    window.location.href = result.next_url || '#';
    return;
  }
  return 'kdjkjscdd';
};
