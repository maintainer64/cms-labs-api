import { ShellFrame } from '@/components/topology/terminal/service/context';

export interface IMiddleware {
  onOpen: (socket?: WebSocket | null, type?: string) => void;
  onData: (socket?: WebSocket | null, shellFrame?: ShellFrame, type?: string) => void;
}
