import React from 'react';
import {GetStraightPathParams} from '@xyflow/react';

/**
 * Функция, которая генерирует строгие горизонтальные линии
 */
export const getFlowchartPath = ({
                                     sourceX,
                                     sourceY,
                                     targetX,
                                     targetY
                                 }: GetStraightPathParams): [string, number, number, number, number] => {
    // Рассчитываем разницу между координатами
    const dx = targetX - sourceX;
    const dy = targetY - sourceY;

    // Определяем направление
    const horizontalFirst = Math.abs(dx) > Math.abs(dy);

    // Рассчитываем путь и контрольные точки
    let path: string;
    let labelX: number, labelY: number;
    const offset = 10; // Смещение метки от линии

    if (horizontalFirst) {
        // Горизонтальное соединение → вертикальное
        const midX = sourceX + dx * 0.5;
        path = `M${sourceX},${sourceY} H${midX} V${targetY} H${targetX}`;

        // Центральная точка пути
        const centerY = sourceY + dy * 0.5;
        labelX = midX;
        labelY = centerY;
    } else {
        // Вертикальное соединение → горизонтальное
        const midY = sourceY + dy * 0.5;
        path = `M${sourceX},${sourceY} V${midY} H${targetX} V${targetY}`;

        // Центральная точка пути
        const centerX = sourceX + dx * 0.5;
        labelX = centerX;
        labelY = midY;
    }

    // Рассчитываем смещение метки (перпендикулярно линии)
    const angle = Math.atan2(dy, dx);
    const offsetX = -Math.sin(angle) * offset;
    const offsetY = Math.cos(angle) * offset;

    return [path, labelX, labelY, offsetX, offsetY];
};

/**
 * Получить наконечник
 * */
export const getSym = (color?: string, type?: string, lineId?: string) => {
    const fill = color || 'inherit';
    const id = getSymId(type, lineId);
    switch (type) {
        case 'dot':
            return (
                <marker
                    style={{fill, strokeWidth: 0}}
                    id={id}
                    viewBox='0 0 10 10'
                    refX='5'
                    refY='5'
                    markerWidth='4'
                    markerHeight='4'
                    orient='auto-start-reverse'
                >
                    <circle cx='5' cy='5' r='5'/>
                </marker>
            );
        case 'square':
            return (
                <marker
                    style={{fill, strokeWidth: 0}}
                    id={id}
                    viewBox='0 0 10 10'
                    refX='5'
                    refY='5'
                    markerWidth='4'
                    markerHeight='4'
                    orient='auto-start-reverse'
                >
                    <rect width='10' height='10'/>
                </marker>
            );
        case 'arrow':
            return (
                <marker
                    style={{fill, strokeWidth: 0}}
                    id={id}
                    viewBox='0 0 10 10'
                    refX='5'
                    refY='5'
                    markerWidth='5'
                    markerHeight='5'
                    orient='auto-start-reverse'
                >
                    <path d='M 0 10 L 10 5 L 0 0 L 3 5 z'/>
                </marker>
            );
        default:
            return '';
    }
};

/**
 * Получить ID наконечника
 * */
export const getSymId = (type?: string, lineId?: string) => {
    if (!type || !lineId) return '';
    return `${type}-${lineId}`;
};

/**
 * Получить URL для стиля наконечника
 * */
export const getSymUrl = (type?: string, lineId?: string) => {
    if (!type || !lineId) return '';
    return `url(#${getSymId(type, lineId)})`;
};

/**
 * Получить PATH SVG для линии
 * */

interface PathLineParams {
    d: string;
    labelX: number;
    labelY: number;
    offsetX: number;
    offsetY: number;
    reverse: boolean;
}

/**
 * Параметры Label в линии
 */
interface LineObjectLabelProps {
    x?: number;
    y?: number;
    color?: string;
    label?: string;
}

/**
 * Рассчитываем размеры на основе количества символов
 */
export const calculateDimensionsLabel = (text: string) => {
    const charWidth = 7; // Средняя ширина символа в пикселях для font-size 0.8em
    const lineHeight = 17; // Высота строки

    const charCount = text?.length || 0;
    const width = Math.max(20, charCount * charWidth + 8); // Минимальная ширина 20px + padding
    const height = lineHeight;

    return {width, height};
};

/**
 * Отобразить Label в линии
 */
export const NodeTypeLineObjectLabel = ({x, y, color, label}: LineObjectLabelProps) => {
    if (!label) return <></>;
    const dimensions = calculateDimensionsLabel(label);

    return (
        <g transform={`translate(${x}, ${y})`}>
            {/* Фон с рассчитанными размерами */}
            <rect
                x={-dimensions.width / 2}
                y={-dimensions.height / 2}
                width={dimensions.width}
                height={dimensions.height}
                fill='#e6f5ff'
                stroke='#346789'
                strokeWidth='2'
                rx='2'
                ry='2'
            />

            {/* Видимый текст */}
            <text
                fill={color || '#0066aa'}
                textAnchor='middle'
                dominantBaseline='middle'
                className='text-xs font-bold font-sans'
                style={{
                    fontSize: '0.8em',
                    fontWeight: 'bold'
                }}
            >
                {label}
            </text>
        </g>
    );
};

/**
 * Получить размеры SVG
 * */
export const getDimensionSVG = (x1: number, y1: number, x2: number, y2: number) => {
    const minX = Math.min(x1, x2);
    const minY = Math.min(y1, y2);
    const width = Math.abs(x2 - x1) + 6;
    const height = Math.abs(y2 - y1) + 6;
    return {minX, minY, width, height};
};
