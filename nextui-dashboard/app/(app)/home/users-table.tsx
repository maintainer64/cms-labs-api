import { UsersTableWrapper } from '@/components/pages/accounts/table/table';
import React from 'react';
import { ContentCardWrapperMain } from '@/components/home/card-wrapper';
import useLanguageBrowser from '@/helpers/locale';
import { RoutesLocation } from '@/components/routes';
import { MapUserItem, useInfinityUserList } from '@/helpers/queries/user/use-infinity-user-list';

const HomeUsersWidget = () => {
  const {
    locale: {
      Tables: { UsersTable }
    }
  } = useLanguageBrowser();
  const response = useInfinityUserList({ limit: 5 });
  const users =
    response?.data?.pages.flatMap((p) => p?.model?.map((item) => MapUserItem(item.model, item.roles)) ?? []) || [];
  return (
    <ContentCardWrapperMain title={UsersTable.TitleWidgetHome} link={RoutesLocation.accounts()} wrapChildren={false}>
      <UsersTableWrapper users={users} isLoading={response.isLoading} />
    </ContentCardWrapperMain>
  );
};

export default HomeUsersWidget;
