export interface ShellFrame {
  Op: string;
  Data?: string;
  SessionID?: string;
  Cols?: number;
  Rows?: number;
}

export interface SJSMessageEvent {
  data: string;
}

export interface SJSCloseEvent {
  reason: string;
}

export interface PodContainerList {
  containers: string[];
}

export interface TerminalResponse {
  id: string;
}
