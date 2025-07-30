import React, { useReducer } from 'react';
import { TopologyFlowVisualization } from '@/components/topology/view';
import { TopologyConnect } from '@/components/topology/connect';
import { TerminalWindows } from '@/components/topology/terminal/window';
import { terminalInitialState, terminalReducer } from '@/components/topology/terminal/service/context';
import { TopologyLayout } from '@/components/topology/layout';
import { KubernetesTerminalFullScreen } from '@/components/topology/terminal/terminalFullScreen';
import { useParams } from 'react-router-dom';

export const TopologyPageConnect = () => {
  return (
    <TopologyLayout>
      <TopologyConnect />
    </TopologyLayout>
  );
};

export const TopologyPageView = () => {
  const [state, dispatch] = useReducer(terminalReducer, terminalInitialState);
  return (
    <TopologyLayout>
      <TerminalWindows nodes={state.clients} dispatch={dispatch} />
      <TopologyFlowVisualization dispatch={dispatch} />
    </TopologyLayout>
  );
};

export const TopologyDevicePageView = () => {
  const { namespace, device } = useParams();
  return <KubernetesTerminalFullScreen namespace={namespace} node={device} />;
};
