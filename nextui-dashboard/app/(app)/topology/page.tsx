import React, { useReducer } from 'react';
import { TopologyFlowVisualization } from '@/components/topology/view';
import { TerminalWindows } from '@/components/topology/terminal/window';
import { terminalInitialState, terminalReducer } from '@/components/topology/terminal/service/context';
import { TopologyLayout } from '@/components/topology/layout';

export const TopologyPageView = () => {
  const [state, dispatch] = useReducer(terminalReducer, terminalInitialState);
  return (
    <TopologyLayout>
      <TerminalWindows nodes={state.clients} dispatch={dispatch} />
      <TopologyFlowVisualization dispatch={dispatch} />
    </TopologyLayout>
  );
};

export default TopologyPageView;
