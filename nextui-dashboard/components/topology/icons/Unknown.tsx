import React from 'react';

export const UnknownIcon = () => {
  return (
    <svg viewBox='0 0 33 48' fill='none' xmlns='http://www.w3.org/2000/svg'>
      <g clipPath='url(#clip0_undef_01)'>
        {/* Передняя грань */}
        <path fillRule='evenodd' clipRule='evenodd' d='M24.495 8.36694H0.5V46.9159H24.495V8.36694Z' fill='#0e6f9c' />
        <path d='M24.495 8.36694H0.5V46.9159H24.495' stroke='white' strokeWidth='0.787' strokeLinecap='square' />
        {/* Верхняя/боковая грань */}
        <path
          fillRule='evenodd'
          clipRule='evenodd'
          d='M24.495 46.916L32.362 39.442V0.5H8.367L0.5 8.367H24.495V46.916Z'
          fill='#0e6f9c'
          stroke='white'
          strokeWidth='0.787'
          strokeLinecap='round'
        />
        {/* Диагональное ребро */}
        <path d='M24.494 8.367L32.361 0.5' stroke='white' strokeWidth='0.787' strokeLinecap='square' />
        {/* Знак вопроса — дуга */}
        <path
          d='M9.8 22C9.8 19.2 11.5 17.2 13.8 17.2C16.1 17.2 17.8 19 17.8 21.2C17.8 23.5 15.6 24.8 13.8 26.4V29.8'
          stroke='white'
          strokeWidth='1.1'
          strokeLinecap='round'
          strokeLinejoin='round'
        />
        {/* Знак вопроса — точка */}
        <circle cx='13.8' cy='33' r='0.95' fill='white' />
      </g>
      <defs>
        <clipPath id='clip0_undef_01'>
          <rect width='33' height='48' fill='white' />
        </clipPath>
      </defs>
    </svg>
  );
};
