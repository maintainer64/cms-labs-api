import { useCallback, useEffect, useRef, useState } from 'react';
import { useScreenSize } from 'use-screen-size';
import { TerminalActionFunc, TerminalClient } from './service/context';
import { SmartLink } from '@/components/navbar/smartLink';
import { Chip, Tooltip } from '@heroui/react';
import useLanguageBrowser from '@/helpers/locale';
import { useReactFlow } from '@xyflow/react';
import { RFNodeTopology } from '@/components/topology/objectTypes/types';
import copy from 'copy-to-clipboard';

interface KubernetesTerminalProps {
  node: TerminalClient;
  dispatch?: TerminalActionFunc;
}

export const KubernetesTerminal = ({ node, dispatch }: KubernetesTerminalProps) => {
  const {
    locale: {
      Topology: { Terminal }
    }
  } = useLanguageBrowser();
  const size = useScreenSize();
  const [isDragging, setIsDragging] = useState(false);
  const [dragStart, setDragStart] = useState({ x: 0, y: 0 });
  const [position, setPosition] = useState({
    x: 100 + node.index * 20,
    y: 100 + node.index * 20
  });
  const [dimension, setDimension] = useState({
    width: 400,
    height: 300
  });

  const headerRef = useRef<HTMLDivElement>(null);
  // Инициализация терминала и подключение
  const { getNode } = useReactFlow();
  // @ts-ignore
  const nodeRF: RFNodeTopology = getNode(node.id);

  // Обработчики drag
  const handleMouseDown = useCallback(
    (e: React.MouseEvent) => {
      if (e.target !== headerRef.current) return;
      setIsDragging(true);
      setDragStart({
        x: e.clientX - position.x,
        y: e.clientY - position.y
      });
      dispatch?.({ type: 'BRING_TO_FRONT', payload: { id: node.id } });
    },
    [position, node.id, dispatch]
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

  if (!node.isVisible) return null;

  const externalIp = nodeRF?.data?.data?.serviceExternalIp?.[0];
  const shellUrl = nodeRF?.data?.data?.shellUrl;

  return (
    <div
      className='fixed rounded-lg overflow-hidden shadow-xl border border-slate-600'
      style={{
        width: `${dimension.width}px`,
        height: `${dimension.height}px`,
        left: `${position.x}px`,
        top: `${position.y}px`,
        zIndex: node.zIndex,
        opacity: node.isOpacity ? 0.6 : 1,
        backdropFilter: 'blur(4px)',
        backgroundColor: 'rgba(30, 41, 59, 0.7)',
        transform: 'translate3d(0,0,0)'
      }}
      onMouseDown={() => dispatch?.({ type: 'BRING_TO_FRONT', payload: { id: node.id } })}
    >
      <div
        ref={headerRef}
        className='flex items-center justify-between bg-slate-800 bg-opacity-70 px-3 py-2 cursor-move'
        onMouseDown={handleMouseDown}
      >
        <Tooltip content={<p>{nodeRF?.data.data?.serviceHealthy ? Terminal.ReadyStatus : Terminal.NotReadyStatus}</p>}>
          <Chip
            color={nodeRF.data.data?.serviceHealthy ? 'success' : 'warning'}
            variant='dot'
            radius='none'
            size='sm'
            classNames={{
              base: 'border-none text-white'
            }}
            onClick={() => {
              externalIp && copy(externalIp);
            }}
          >
            {nodeRF.data.label}&nbsp;{externalIp}
          </Chip>
        </Tooltip>
        <div className='flex space-x-2'>
          <button
            onClick={() => dispatch?.({ type: 'TOGGLE_OPACITY', payload: { id: node.id } })}
            className='text-slate-300 hover:text-blue-400'
            title='Toggle opacity'
          >
            {node.isOpacity ? '◉' : '◎'}
          </button>
          <SmartLink
            to={shellUrl}
            className='text-slate-300 hover:text-blue-400 cursor-pointer'
            title='Open in new window'
          >
            ↗
          </SmartLink>
          <button
            onClick={() => dispatch?.({ type: 'TOGGLE_VISIBILITY', payload: { id: node.id } })}
            className='text-slate-300 hover:text-red-400'
            title='Close'
          >
            ✕
          </button>
        </div>
      </div>
      {shellUrl && (
        <iframe
          className='w-full h-full'
          style={{ height: `calc(100% - 40px)`, backgroundColor: 'black' }}
          src={shellUrl}
        ></iframe>
      )}
      <div
        className='absolute bottom-0 right-0 w-4 h-4 cursor-se-resize bg-blue-500 opacity-0 hover:opacity-100'
        onMouseDown={handleResize}
      />
    </div>
  );
};
