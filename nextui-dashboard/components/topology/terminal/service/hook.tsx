import { useEffect, useRef, useState } from 'react';
import { TerminalService } from '@/components/topology/terminal/service/real';
import { addToast } from '@heroui/react';
import { TerminalClient } from '@/components/topology/terminal/service/context';
import { UsecasesContainerGetItem } from '@/helpers/api';
import useLanguageBrowser from '@/helpers/locale';
import { CamelCasedPropertiesDeep } from 'type-fest';
import { useMutationContainerGet } from '@/helpers/queries/container/use-mutation-container-get';

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
  const { mutateAsync } = useMutationContainerGet({});
  const terminalService = useRef<TerminalService | null>(null);
  const isMock = localStorage.getItem('socketJsMock') === 'true';
  const [container, setContainer] = useState<CamelCasedPropertiesDeep<UsecasesContainerGetItem>>({
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
      const container = responseContainers?.containers?.[0];
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
        terminalService.current = new TerminalService(terminalRef.current!, Terminal);
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
      await terminalService.current.connect({
        sessionId: container.sessionId || '',
        type: container.type || 'default',
        url: container.connectUrl || '',
        startup: container.startup || ''
      });
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
