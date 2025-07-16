/**
 * Рассчитываем размеры на основе количества символов
 */
export const calculateDimensionsLabel = (text: string) => {
  const charWidth = 7; // Средняя ширина символа в пикселях для font-size 0.8em
  const lineHeight = 17; // Высота строки

  const charCount = text?.length || 0;
  const width = Math.max(20, charCount * charWidth + 8); // Минимальная ширина 20px + padding
  const height = lineHeight;

  return { width, height };
};
