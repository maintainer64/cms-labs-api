import React, { useReducer } from 'react';
import { TopologyFlowVisualization } from '@/components/topology/view';
import { TopologyConnect } from '@/components/topology/connect';
import { TerminalWindows } from '@/components/topology/terminal/window';
import { terminalInitialState, terminalReducer } from '@/components/topology/terminal/context';
import { TopologyLayout } from '@/components/topology/layout';

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
      <TerminalWindows isMock={true} clients={state.clients} dispatch={dispatch} />
      <TopologyFlowVisualization dispatch={dispatch} />
    </TopologyLayout>
  );
};

export const TopologyDevicePageView = () => {
  const [state, dispatch] = useReducer(terminalReducer, terminalInitialState);
  return (
    <TopologyLayout>
      <TerminalWindows isMock={true} clients={state.clients} dispatch={dispatch} />
      <TopologyFlowVisualization dispatch={dispatch} />
    </TopologyLayout>
  );
};
