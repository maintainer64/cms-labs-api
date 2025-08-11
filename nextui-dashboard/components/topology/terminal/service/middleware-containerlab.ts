import { ShellFrame } from '@/components/topology/terminal/service/context';
import { IMiddleware } from '@/components/topology/terminal/service/middleware-type';

export class ContainerLabMiddleware implements IMiddleware {
  onOpen(socket?: WebSocket | null, type?: string) {
    if (type !== 'containerlab') return;
    const bindFrame: ShellFrame = {
      Op: 'stdin',
      Data: 'shellin\r\n'
    };
    socket?.send(JSON.stringify(bindFrame));
    return;
  }

  onData(socket?: WebSocket | null, shellFrame?: ShellFrame, type?: string) {
    if (type !== 'containerlab') return;
    if (shellFrame?.Op !== 'stdout') return;
    if (
      !shellFrame?.Data?.includes('Container not found. Maybe the lab is still deploying. Try again in a few seconds.')
    )
      return;
    socket?.close(4999, 'Container not found');
  }
}
