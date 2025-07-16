import React from 'react';
import { useTopologyCreate } from '@/helpers/queries/topology/get';
import { Loading } from '@/components/scroll/loader';
import { ErrorModal } from '@/components/pages/auth/error';
import { Button } from '@heroui/react';
import { RoutesLocation } from '@/components/routes';
import { useParamsConnectTopology } from '@/components/topology/utils';

export const TopologyConnect = () => {
  const params = useParamsConnectTopology();
  const queryTopologyCreate = useTopologyCreate(params.username, params.taskId, params.redeploy);
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
  window.location.href =
    RoutesLocation.topologyView(queryTopologyCreate.data?.result?.namespace ?? '') +
    `?username=${params.username}&taskId=${params.taskId}`;
  return <></>;
};
