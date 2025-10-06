'use client';
import React, { useEffect } from 'react';
import { HorizontalInfiniteLoader } from '@/components/scroll/loader';
import { ErrorModal } from '@/components/pages/auth/error';
import { Button } from '@heroui/react';
import { RoutesLocation } from '@/components/routes';
import { useParamsConnectTopology } from '@/components/topology/utils';
import useLanguageBrowser from '@/helpers/locale';
import { useQueryTopologyCreate } from '@/helpers/queries/topology/use-query-topology-create';

export const TopologyConnect = () => {
  const {
    locale: {
      Topology: { Connect }
    }
  } = useLanguageBrowser();
  const params = useParamsConnectTopology();
  const queryTopologyCreate = useQueryTopologyCreate({
    username: params.username,
    taskId: params.taskId,
    redeploy: params.redeploy
  });
  useEffect(() => {
    window.document.title = Connect.ErrorModalConnectTitle;
  }, []);
  if (queryTopologyCreate.isLoading) return <HorizontalInfiniteLoader />;
  if (queryTopologyCreate.error) {
    // @ts-ignore
    const errMsg = queryTopologyCreate?.error?.body?.msg || Connect.Error;
    return (
      <ErrorModal title={Connect.ErrorModalConnectTitle} description={errMsg}>
        <Button onPress={() => queryTopologyCreate.refetch()} href='#' variant='light' color='primary'>
          {Connect.ErrorModalRetry}
        </Button>
      </ErrorModal>
    );
  }
  window.location.href =
    RoutesLocation.topologyView(queryTopologyCreate.data?.namespace ?? '') +
    `?username=${params.username}&taskId=${params.taskId}`;
  return <></>;
};
