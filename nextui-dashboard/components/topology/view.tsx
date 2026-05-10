'use client';
import React, { useEffect } from 'react';
import { Background, Controls, ReactFlow } from '@xyflow/react';
import { edgeTypes, nodeTypes } from './objectTypes';
import { getLayoutElements } from './autoLayout';
import { RFNodeTopology } from '@/components/topology/objectTypes/types';
import { HorizontalInfiniteLoader } from '@/components/scroll/loader';
import { ErrorModal } from '@/components/pages/auth/error';
import { Button } from '@heroui/react';
import { TerminalActionFunc } from '@/components/topology/terminal/service/context';
import useLanguageBrowser from '@/helpers/locale';
import useThemeBrowser from '@/components/navbar/useTheme';
import { useQueryTopologyGet } from '@/helpers/queries/topology/use-query-topology-get';
import { useParams } from 'react-router-dom';
import { parseTopology } from '@/components/topology/objectTypes/parse';

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
  const { namespace } = useParams();
  const queryTopology = useQueryTopologyGet({ namespace });
  useEffect(() => {
    // @ts-ignore
    window.document.title = namespace;
    return () => {
      window.document.title = 'CMS LABS';
    };
  }, []);
  if (queryTopology.isLoading) return <HorizontalInfiniteLoader />;
  if (queryTopology.error || queryTopology.data === undefined) {
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
  const parsedTopology = parseTopology(queryTopology.data);
  const object = getLayoutElements(parsedTopology.nodes, parsedTopology.edges, parsedTopology.direction);
  return (
    <ReactFlow
      colorMode={theme}
      nodes={object.nodes}
      // @ts-ignore
      edges={object.edges}
      nodeTypes={nodeTypes}
      edgeTypes={edgeTypes}
      onNodeClick={(_, node: RFNodeTopology) => {
        dispatch?.({
          type: 'ADD_CLIENT',
          payload: {
            id: node.id
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
