import React, { useReducer } from 'react';
import { TopologyFlowVisualization, TopologyLayout } from '@/components/topology/view';
import { TopologyConnect } from '@/components/topology/connect';
import { TerminalWindows } from '@/components/topology/terminal/window';
import { terminalInitialState, terminalReducer } from '@/components/topology/terminal/context';

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
