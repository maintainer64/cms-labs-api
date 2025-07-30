import { useEffect, useRef, useState } from 'react';
import { TerminalService } from '@/components/topology/terminal/service/real';
import { MockTerminalService } from '@/components/topology/terminal/service/mock';
import { addToast } from '@heroui/react';
import { TerminalClient } from '@/components/topology/terminal/service/context';
import { useContainersGet } from '@/helpers/queries/topology/get';
import { usecases_ContainersGetItem } from '@/helpers/api';
import useLanguageBrowser from '@/helpers/locale';

interface UseTerminalParams {
  terminalRef: React.RefObject<any>;
  node: TerminalClient;
}

const DELAY_MIN = 5000; // 5c
const DELAY_MAX = 15000; // 15с

export const useTerminal = ({ terminalRef, node }: UseTerminalParams) => {
  const {
    locale: {
      Topology: { Terminal }
    }
  } = useLanguageBrowser();
  const { mutateAsync } = useContainersGet({});
  const terminalService = useRef<TerminalService | MockTerminalService | null>(null);
  const isMock = localStorage.getItem('socketJsMock') === 'true';
  const [container, setContainer] = useState<usecases_ContainersGetItem>({
    name: node.id,
    namespace: node.namespace,
    status: 'unknown'
  });

  useEffect(() => {
    if (!terminalRef.current) return;

    // Переменные для управления задержкой
    let timeoutId: NodeJS.Timeout | null = null;

    const initTerminal = async () => {
      // Очистка предыдущего подключения
      terminalService.current?.disconnect();
      const responseContainers = await mutateAsync({ namespace: node.namespace, deployment: node.id });
      const container = responseContainers.result?.containers?.[0];
      if (!container) {
        scheduleInitTerminal();
        return;
      }
      setContainer(container);
      if (!container.ready) {
        scheduleInitTerminal();
        return;
      }
      // Инициализация сервиса
      if (!terminalService.current) {
        terminalService.current = isMock
          ? new MockTerminalService(terminalRef.current!, Terminal)
          : new TerminalService(terminalRef.current!, Terminal);
      }
      // Настройка callback для уведомлений
      terminalService.current.setOnToast((msg) =>
        addToast({
          title: `${container.label} (${container.name})`,
          description: msg,
          color: 'default',
          timeout: 1
        })
      );
      terminalService.current.setOnReconnect(() => {
        scheduleInitTerminal();
      });
      await terminalService.current.connect(container.session_id, container.connect_url);
    };

    // Штука для автоматического перезапроса транспорта если отвалилось подключение
    const scheduleInitTerminal = () => {
      if (timeoutId) return;
      // Откладываем выполнение
      timeoutId = setTimeout(
        () => {
          initTerminal();
          timeoutId = null;
        },
        Math.floor(Math.random() * DELAY_MAX) + DELAY_MIN
      );
    };

    initTerminal();

    return () => {
      timeoutId ? clearTimeout(timeoutId) : undefined;
      terminalService.current?.disconnect();
      terminalService.current = null;
    };
  }, [isMock, node.name, node.id, node.namespace, mutateAsync]);

  return {
    terminalService,
    container
  };
};
