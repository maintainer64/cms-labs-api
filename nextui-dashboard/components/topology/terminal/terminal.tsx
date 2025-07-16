import { useCallback, useEffect, useRef, useState } from 'react';
import { useScreenSize } from 'use-screen-size';
import '@xterm/xterm/css/xterm.css';
import { TerminalActionFunc, TerminalClient } from './context';
import { TerminalService } from '@/components/topology/terminal/terminalService';
import { MockTerminalService } from './mockTerminalService';
import { addToast } from '@heroui/react';

interface KubernetesTerminalProps {
  isMock: boolean;
  client: TerminalClient;
  dispatch?: TerminalActionFunc;
}

export const KubernetesTerminal = ({ isMock, client, dispatch }: KubernetesTerminalProps) => {
  const size = useScreenSize();
  const [isDragging, setIsDragging] = useState(false);
  const [dragStart, setDragStart] = useState({ x: 0, y: 0 });
  const [position, setPosition] = useState({
    x: 100 + client.index * 20,
    y: 100 + client.index * 20
  });
  const [dimension, setDimension] = useState({
    width: 400,
    height: 300
  });

  const terminalRef = useRef<HTMLDivElement>(null);
  const headerRef = useRef<HTMLDivElement>(null);
  const terminalService = useRef<TerminalService | MockTerminalService | null>(null);
  // Инициализация терминала и подключение
  useEffect(() => {
    if (!terminalRef.current || !client.isVisible) return;

    const initTerminal = async () => {
      if (terminalService.current) {
        terminalService.current.disconnect();
      }

      if (isMock) {
        terminalService.current = new MockTerminalService(terminalRef.current!);
      } else {
        terminalService.current = new TerminalService(terminalRef.current!);
      }

      // Установка toast callback
      terminalService.current.setOnToast((msg) =>
        addToast({
          title: `${client.name}(${client.id})`,
          description: msg,
          color: 'default'
        })
      );

      await terminalService.current.connect(client.namespace || 'default', client.id, '#');
    };

    initTerminal();

    return () => {
      terminalService.current?.disconnect();
      terminalService.current = null;
    };
  }, [client.isVisible, client.namespace, client.id, isMock]);

  // Обработчики drag
  const handleMouseDown = useCallback(
    (e: React.MouseEvent) => {
      if (e.target !== headerRef.current) return;
      setIsDragging(true);
      setDragStart({
        x: e.clientX - position.x,
        y: e.clientY - position.y
      });
      dispatch?.({ type: 'BRING_TO_FRONT', payload: { id: client.id } });
    },
    [position, client.id, dispatch]
  );

  const handleMouseMove = useCallback(
    (e: MouseEvent) => {
      if (!isDragging) return;
      const [minX, minY, maxX, maxY] = [
        10.0,
        70.0,
        size.width - dimension.width - 5.0,
        size.height - dimension.height - 5.0
      ];
      setPosition({
        x: Math.min(Math.max(minX, e.clientX - dragStart.x), maxX),
        y: Math.min(Math.max(minY, e.clientY - dragStart.y), maxY)
      });
    },
    [isDragging, dragStart]
  );

  const handleMouseUp = useCallback(() => {
    setIsDragging(false);
  }, []);

  // Подписка на события drag
  useEffect(() => {
    if (isDragging) {
      window.addEventListener('mousemove', handleMouseMove);
      window.addEventListener('mouseup', handleMouseUp);
      return () => {
        window.removeEventListener('mousemove', handleMouseMove);
        window.removeEventListener('mouseup', handleMouseUp);
      };
    }
  }, [isDragging, handleMouseMove, handleMouseUp]);

  // Обработчик изменения размеров
  const handleResize = useCallback(
    (e: React.MouseEvent) => {
      e.stopPropagation();
      const startWidth = dimension.width;
      const startHeight = dimension.height;
      const startX = e.clientX;
      const startY = e.clientY;

      const onMouseMove = (e: MouseEvent) => {
        setDimension({
          width: Math.max(300, startWidth + (e.clientX - startX)),
          height: Math.max(200, startHeight + (e.clientY - startY))
        });
      };

      const onMouseUp = () => {
        window.removeEventListener('mousemove', onMouseMove);
        window.removeEventListener('mouseup', onMouseUp);
      };

      window.addEventListener('mousemove', onMouseMove);
      window.addEventListener('mouseup', onMouseUp);
    },
    [dimension]
  );

  if (!client.isVisible) return null;

  return (
    <div
      className='fixed rounded-lg overflow-hidden shadow-xl border border-slate-600'
      style={{
        width: `${dimension.width}px`,
        height: `${dimension.height}px`,
        left: `${position.x}px`,
        top: `${position.y}px`,
        zIndex: client.zIndex,
        opacity: client.isOpacity ? 0.6 : 1,
        backdropFilter: 'blur(4px)',
        backgroundColor: 'rgba(30, 41, 59, 0.7)',
        transform: 'translate3d(0,0,0)'
      }}
      onMouseDown={() => dispatch?.({ type: 'BRING_TO_FRONT', payload: { id: client.id } })}
    >
      <div
        ref={headerRef}
        className='flex items-center justify-between bg-slate-800 bg-opacity-70 px-3 py-2 cursor-move'
        onMouseDown={handleMouseDown}
      >
        <div className='text-slate-200 text-sm truncate max-w-[200px]'>
          {client.name} ({client.id})
        </div>
        <div className='flex space-x-2'>
          <button
            onClick={() => dispatch?.({ type: 'TOGGLE_OPACITY', payload: { id: client.id } })}
            className='text-slate-300 hover:text-blue-400'
            title='Toggle opacity'
          >
            {client.isOpacity ? '◉' : '◎'}
          </button>
          <button
            onClick={() => console.log('Open in new window')}
            className='text-slate-300 hover:text-blue-400'
            title='Open in new window'
          >
            ↗
          </button>
          <button
            onClick={() => dispatch?.({ type: 'TOGGLE_VISIBILITY', payload: { id: client.id } })}
            className='text-slate-300 hover:text-red-400'
            title='Close'
          >
            ✕
          </button>
        </div>
      </div>

      <div
        className='absolute bottom-0 right-0 w-4 h-4 cursor-se-resize bg-blue-500 opacity-0 hover:opacity-100'
        onMouseDown={handleResize}
      />

      <div
        ref={terminalRef}
        className='w-full h-full'
        style={{ height: `calc(100% - 40px)`, backgroundColor: 'black' }}
      />
    </div>
  );
};
