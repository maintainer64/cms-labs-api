import { ShellFrame } from '@/components/topology/terminal/service/context';
import { ContainerConnectParams, IMiddleware } from '@/components/topology/terminal/service/middleware-type';

export class ContainerLabMiddleware implements IMiddleware {
  onOpen(socket?: WebSocket | null, connect?: ContainerConnectParams) {
    if (connect?.type !== 'containerlab') return;
    if (!connect?.startup) return;
    const bindFrame: ShellFrame = {
      Op: 'stdin',
      Data: `${connect?.startup}\r\n`
    };
    socket?.send(JSON.stringify(bindFrame));
    return;
  }

  onData(socket?: WebSocket | null, shellFrame?: ShellFrame, connect?: ContainerConnectParams) {
    if (connect?.type !== 'containerlab') return;
    if (shellFrame?.Op !== 'stdout') return;
    if (
      !shellFrame?.Data?.includes('Container not found. Maybe the lab is still deploying. Try again in a few seconds.')
    )
      return;
    socket?.close(4999, 'Container not found');
  }
}
