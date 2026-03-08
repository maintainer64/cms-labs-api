'use client';
import React, { useEffect } from 'react';
import { Background, Controls, ReactFlow } from '@xyflow/react';
import { edgeTypes, nodeTypes } from './objectTypes';
import { getLayoutElements } from './autoLayout';
import { RFEdgeTopology, RFNodeTopology } from '@/components/topology/objectTypes/types';
import { HorizontalInfiniteLoader } from '@/components/scroll/loader';
import { ErrorModal } from '@/components/pages/auth/error';
import { Button } from '@heroui/react';
import { useParamsConnectTopology } from '@/components/topology/utils';
import { TerminalActionFunc } from '@/components/topology/terminal/service/context';
import useLanguageBrowser from '@/helpers/locale';
import { InfoModalBlock } from '@/components/layout/infoModalBlock';
import useThemeBrowser from '@/components/navbar/useTheme';
import { useQueryTopologyGet } from '@/helpers/queries/topology/use-query-topology-get';

interface TopologyFlowVisualizationProps {
  dispatch?: TerminalActionFunc;
}

export const TopologyFlowVisualization = ({ dispatch }: TopologyFlowVisualizationProps) => {
  const { theme } = useThemeBrowser();
  const {
    locale: {
      Topology: { Connect }
    }
  } = useLanguageBrowser();
  const { namespace } = useParamsConnectTopology();
  const queryTopology = useQueryTopologyGet({ namespace });
  useEffect(() => {
    window.document.title = namespace;
    return () => {
      window.document.title = 'CMS LABS';
    };
  }, []);
  const initNodes = (queryTopology.data?.topology?.nodes ?? []) as RFNodeTopology[];
  const initEdges = (queryTopology.data?.topology?.edges ?? []) as RFEdgeTopology[];
  const object = getLayoutElements(
    initNodes,
    initEdges,
    (queryTopology.data?.topology?.direction || 'TB') as 'TB' | 'LR'
  );
  if (queryTopology.isLoading) return <HorizontalInfiniteLoader />;
  if (queryTopology.error) {
    // @ts-ignore
    const errMsg = queryTopology?.error?.data?.message || Connect.Error;
    return (
      <ErrorModal title={Connect.ErrorModalViewTitle} description={errMsg}>
        <Button onPress={() => queryTopology.refetch()} href='#' variant='light' color='primary'>
          {Connect.ErrorModalRetry}
        </Button>
      </ErrorModal>
    );
  }
  if (!queryTopology.data?.topology) {
    return (
      <>
        <HorizontalInfiniteLoader />
        <InfoModalBlock
          title={Connect.WaitModalTitle}
          href={queryTopology.data?.webUrl}
          description={Connect.WaitModalDescription}
          buttonText={Connect.WaitModalButtonText}
        />
      </>
    );
  }
  return (
    <ReactFlow
      colorMode={theme}
      nodes={object.nodes}
      edges={object.edges}
      nodeTypes={nodeTypes}
      edgeTypes={edgeTypes}
      onNodeClick={(_, node: RFNodeTopology) => {
        dispatch?.({
          type: 'ADD_CLIENT',
          payload: {
            id: node.id,
            namespace: namespace,
            name: node.data.label
          }
        });
        dispatch?.({
          type: 'TOGGLE_VISIBILITY',
          payload: {
            id: node.id,
            // Если открывается терминал с общей топологии, нужно открыть поверх всего
            onChange: (visibility?: boolean) => {
              if (visibility) {
                dispatch?.({
                  type: 'BRING_TO_FRONT',
                  payload: {
                    id: node.id
                  }
                });
              }
            }
          }
        });
      }}
      fitView={true}
    >
      <Background />
      <Controls />
    </ReactFlow>
  );
};
