import {Handle, Position} from '@xyflow/react';
import React from 'react';
import {topology_TopologiesNode} from "@/helpers/api";
import ImageIcon from "@/components/topology/icons";

export const nodeTypeDefault = ({data}: topology_TopologiesNode) => {
    return (
        <div className='
      min-w-[4rem] w-auto h-auto
      relative
      flex flex-col items-center
      group
    '>
            {/* Контейнер для иконки с фиксированным размером */}
            <div className='w-12 h-12 flex items-center justify-center'>
                <ImageIcon icon={data?.icon} label={data?.label}/>
            </div>

            {/* Текст с переносом строк */}
            <div className='
        text-sm font-medium
        text-gray-800 dark:text-gray-200
        text-center              // Центрирование текста
        break-words              // Перенос длинных слов
        max-w-[8rem]             // Макс. ширина текста
      '>
                {data?.label || 'Unnamed'}
            </div>

            {data?.type && (
                <div className='
                    absolute -bottom-8 left-1/2 transform -translate-x-1/2
                    px-2 py-1
                    bg-gray-700 dark:bg-gray-200
                    text-white dark:text-gray-800
                    text-xs rounded
                    opacity-0 group-hover:opacity-100
                    transition-opacity
                    whitespace-nowrap
                    z-10
                '>
                    {data.type}
                    <div
                        className='absolute -top-1 left-1/2 transform -translate-x-1/2 w-2 h-2 bg-gray-700 dark:bg-gray-200 rotate-45'/>
                </div>
            )}

            <Handle position={Position.Bottom} type='target' hidden={true}/>
            <Handle position={Position.Bottom} type='source' hidden={true}/>
        </div>
    );
};