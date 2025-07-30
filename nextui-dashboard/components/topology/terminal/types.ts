type Color = 'default' | 'primary' | 'secondary' | 'success' | 'warning' | 'danger';

export function terminalStatusToColor(color?: string): Color {
  if (color === 'running') return 'success';
  if (color === 'waiting') return 'warning';
  if (color === 'terminated') return 'danger';
  return 'default';
}
