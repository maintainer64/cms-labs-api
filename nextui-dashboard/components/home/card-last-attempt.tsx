import React from 'react';
import { LTIAttemptTableWrapper } from '@/components/pages/lti-attempts/table/table';
import useLanguageBrowser from '@/helpers/locale';
import { ContentCardWrapperMain } from '@/components/home/card-wrapper';
import { useQueryLtiAttemptList } from '@/helpers/queries/lti_attempt/use-query-lti-attempt-list';

export const CardLastAttempt = () => {
  const {
    locale: {
      Tables: { LTIAttemptsTable }
    }
  } = useLanguageBrowser();
  const response = useQueryLtiAttemptList({ limit: 10 });
  const items = response.data?.model || [];
  return (
    <ContentCardWrapperMain title={LTIAttemptsTable.TitleWidgetHome} wrapChildren={false}>
      <LTIAttemptTableWrapper isLoading={response.isLoading} rows={items} />
    </ContentCardWrapperMain>
  );
};
