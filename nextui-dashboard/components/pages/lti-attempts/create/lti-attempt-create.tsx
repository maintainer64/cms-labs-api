'use client';

import { useLTIAttemptCreate } from '@/helpers/queries/lti-attempt/get';
import { Loading } from '@/components/scroll/loader';
import React from 'react';
import { ErrorModal } from '@/components/pages/auth/error';
import useLanguageBrowser from '@/helpers/locale';
import { Link } from '@heroui/react';

export const LTIAttemptCreate = () => {
  const {
    locale: {
      PnetServersQueue: { LTIAttemptRoom }
    }
  } = useLanguageBrowser();
  const response = useLTIAttemptCreate();
  if (response.isLoading) return <Loading size={8} />;
  if (response.error)
    return (
      <ErrorModal title={LTIAttemptRoom.ErrorPageTitle} description={response.error.body.msg}>
        <Link onPress={() => response.refetch()} href='#' underline='always'>
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
