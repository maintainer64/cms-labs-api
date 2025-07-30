import { TerminalActionFunc, TerminalClient } from '@/components/topology/terminal/service/context';
import { KubernetesTerminal } from '@/components/topology/terminal/terminal';

interface TerminalWindows {
  nodes: TerminalClient[];
  dispatch?: TerminalActionFunc;
}

export const TerminalWindows = ({ nodes, dispatch }: TerminalWindows) => {
  return (
    <>
      {nodes.map((node) => (
        <KubernetesTerminal key={node.id} node={node} dispatch={dispatch} />
      ))}
    </>
  );
};
