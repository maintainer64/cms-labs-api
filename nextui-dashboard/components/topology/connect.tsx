import React from 'react';
import { useTopologyCreate } from '@/helpers/queries/topology/get';
import { useLocation } from 'react-router-dom';
import { Loading } from '@/components/scroll/loader';
import { ErrorModal } from '@/components/pages/auth/error';
import { Button } from '@heroui/react';
import { RoutesLocation } from '@/components/routes';
import { useUserProfile } from '@/components/providers/auth-jwt/hooks';

export const TopologyConnect = () => {
  const { search } = useLocation();
  const user = useUserProfile();

  const query = new URLSearchParams(search);
  const queryTopologyCreate = useTopologyCreate(
    query.get('username') || user.username,
    query.get('taskId') || '',
    false
  );
  if (queryTopologyCreate.isLoading) return <Loading size='md' />;
  if (queryTopologyCreate.error) {
    // @ts-ignore
    const errMsg = queryTopologyCreate?.error?.body?.msg || 'Внутрянняя ошибка';
    return (
      <ErrorModal title={'Подключение к топологии'} description={errMsg}>
        <Button onPress={() => queryTopologyCreate.refetch()} href='#' variant='light' color='primary'>
          Попробовать снова
        </Button>
      </ErrorModal>
    );
  }
  window.location.href = RoutesLocation.topologyView(queryTopologyCreate.data?.result?.namespace ?? '');
  return <></>;
};
