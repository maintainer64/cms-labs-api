import { useEffect, useRef } from 'react';
import '@xterm/xterm/css/xterm.css';
import { useTerminal } from '@/components/topology/terminal/service/hook';

interface KubernetesTerminalFullScreenProps {
  namespace?: string;
  node?: string;
}

export const KubernetesTerminalFullScreen = ({ namespace, node }: KubernetesTerminalFullScreenProps) => {
  const terminalRef = useRef<HTMLDivElement>(null);
  // Инициализация терминала и подключение
  const { terminalService, container } = useTerminal({
    terminalRef: terminalRef,
    node: {
      id: node || '',
      namespace: namespace || '',
      name: node || '',
      index: 1,
      isVisible: true,
      isOpacity: false,
      zIndex: 40
    }
  });
  useEffect(() => {
    window.document.title = `${container.label} (${container.name})`;
    terminalService.current?.fit();
    return () => {
      window.document.title = 'CMS LABS';
    };
  }, [container, terminalService]);
  return (
    <div className='w-screen h-screen'>
      <div ref={terminalRef} style={{ height: `100%`, backgroundColor: 'black' }} />
    </div>
  );
};
