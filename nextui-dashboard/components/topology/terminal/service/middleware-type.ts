import { ShellFrame } from '@/components/topology/terminal/service/context';

export interface IMiddleware {
  onOpen: (socket?: WebSocket | null, connect?: ContainerConnectParams) => void;
  onData: (socket?: WebSocket | null, shellFrame?: ShellFrame, connect?: ContainerConnectParams) => void;
}

export interface ContainerConnectParams {
  sessionId: string;
  type: string;
  url: string;
  startup: string;
}
