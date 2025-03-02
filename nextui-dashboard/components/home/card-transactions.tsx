import React from 'react';
import { useLTIAttemptList } from '@/helpers/queries/lti-attempt/get';
import { LTIAttemptTableWrapper } from '@/components/pages/lti-attempts/table/table';
import useLanguageBrowser from '@/helpers/locale';
import { ContentCardWrapperMain } from '@/components/home/card-wrapper';

export const CardTransactions = () => {
  const {
    locale: {
      Tables: { LTIAttemptsTable }
    }
  } = useLanguageBrowser();
  const response = useLTIAttemptList({ limit: 10 });
  const items = response.data?.result?.model || [];
  return (
    <ContentCardWrapperMain title={LTIAttemptsTable.TitleWidgetHome} wrapChildren={false}>
      <LTIAttemptTableWrapper isLoading={response.isLoading} rows={items} />
    </ContentCardWrapperMain>
  );
};
