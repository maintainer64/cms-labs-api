import { TerminalActionFunc, TerminalClient } from '@/components/topology/terminal/context';
import { KubernetesTerminal } from '@/components/topology/terminal/terminal';

interface TerminalWindows {
  isMock: boolean;
  clients: TerminalClient[];
  dispatch?: TerminalActionFunc;
}

export const TerminalWindows = ({ isMock, clients, dispatch }: TerminalWindows) => {
  return (
    <>
      {clients.map((client) => (
        <KubernetesTerminal isMock={isMock} key={client.id} client={client} dispatch={dispatch} />
      ))}
    </>
  );
};
