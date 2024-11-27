import { UsersTableWrapper } from '@/components/pages/accounts/table/table';
import React from 'react';
import { useUsersList } from '@/helpers/queries/users/get';
import { ContentCardWrapperMain } from '@/components/home/card-wrapper';
import useLanguageBrowser from '@/helpers/locale';
import { RoutesLocation } from '@/components/routes';

const HomeUsersWidget = () => {
  const {
    locale: {
      Tables: { UsersTable }
    }
  } = useLanguageBrowser();
  const response = useUsersList({ limit: 5 });
  const users = response?.data?.pages.flatMap((p) => p.result?.model ?? []) || [];
  return (
    <ContentCardWrapperMain title={UsersTable.TitleWidgetHome} link={RoutesLocation.accounts()} wrapChildren={false}>
      <UsersTableWrapper users={users} isLoading={response.isLoading} />
    </ContentCardWrapperMain>
  );
};

export default HomeUsersWidget;
