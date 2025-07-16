import React from 'react';
import { Background, Controls, ReactFlow } from '@xyflow/react';
import { edgeTypes, nodeTypes } from './objectTypes';
import { getLayoutElements } from './autoLayout';
import { useTopologyGet } from '@/helpers/queries/topology/get';
import { RFEdgeTopology, RFNodeTopology } from '@/components/topology/objectTypes/types';
import { Loading } from '@/components/scroll/loader';
import { ErrorModal } from '@/components/pages/auth/error';
import { Button } from '@heroui/react';
import TopologyMenu from '@/components/topology/menu/menu';
import { useParamsConnectTopology } from '@/components/topology/utils';
import { TerminalActionFunc } from '@/components/topology/terminal/context';

interface Props {
  children: React.ReactNode;
}

export const TopologyLayout = ({ children }: Props) => {
  return (
    <div className='flex flex-col h-screen'>
      <TopologyMenu />
      <div className='flex-1 w-full'>{children}</div>
    </div>
  );
};

interface TopologyFlowVisualizationProps {
  dispatch?: TerminalActionFunc;
}

export const TopologyFlowVisualization = ({ dispatch }: TopologyFlowVisualizationProps) => {
  const { namespace } = useParamsConnectTopology();
  const queryTopology = useTopologyGet(namespace);
  const initNodes = (queryTopology.data?.result?.topology?.nodes ?? []) as RFNodeTopology[];
  const initEdges = (queryTopology.data?.result?.topology?.edges ?? []) as RFEdgeTopology[];
  const object = getLayoutElements(
    initNodes,
    initEdges,
    (queryTopology.data?.result?.topology?.direction || 'TB') as 'TB' | 'LR'
  );
  if (queryTopology.isLoading) return <Loading size='md' />;
  if (queryTopology.error) {
    // @ts-ignore
    const errMsg = queryTopology?.error?.body?.msg || 'Внутрянняя ошибка';
    return (
      <ErrorModal title={'Отображение топологии'} description={errMsg}>
        <Button onPress={() => queryTopology.refetch()} href='#' variant='light' color='primary'>
          Попробовать снова
        </Button>
      </ErrorModal>
    );
  }
  return (
    <ReactFlow
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
      fitView
    >
      <Background />
      <Controls />
    </ReactFlow>
  );
};
