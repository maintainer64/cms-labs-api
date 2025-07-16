import { Terminal } from '@xterm/xterm';
import { FitAddon } from '@xterm/addon-fit';
import debounce from 'lodash/debounce';
import { ShellFrame, SJSCloseEvent, SJSMessageEvent, TerminalResponse } from './types';

type Function = (...args: any[]) => any;

// Плейсхолдеры для API (как раньше)
const fetchUtilityShell = async (url: string): Promise<TerminalResponse> => {
  const response = await fetch(url);
  return response.json();
};

export class TerminalService {
  private term: Terminal | null = null;
  private fitAddon: FitAddon | null = null;
  private conn: WebSocket | null = null; // Теперь native WebSocket
  private debouncedFit: Function | null = null;
  private connected = false;
  private connecting = false;
  private connectionClosed = false;
  private onToast: (message: string) => void = () => {};

  constructor(private container: HTMLElement) {}

  setOnToast(callback: (message: string) => void) {
    this.onToast = callback;
  }

  async connect(namespace: string, podName: string, container: string): Promise<void> {
    if (!container || !podName || !namespace || this.connecting) return;

    this.connecting = true;
    this.connectionClosed = false;

    const terminalSessionUrl = `your-api-endpoint/shell/${namespace}/${podName}/${container}`;
    const { id } = await fetchUtilityShell(terminalSessionUrl);

    // Используем ws/wss URL (адаптируйте под ваш сервер!)
    const wsUrl = `ws://your-api/sockjs?${id}`; // Или wss:// для secure
    this.conn = new WebSocket(wsUrl);

    this.conn.onopen = () => this.onConnectionOpen(id);
    this.conn.onmessage = (evt) => this.onConnectionMessage({ data: evt.data } as SJSMessageEvent);
    this.conn.onclose = (evt) => this.onConnectionClose({ reason: evt.reason } as SJSCloseEvent);
    this.conn.onerror = (err) => this.onToast(`WebSocket error: ${err}`);

    this.initTerm();
  }

  disconnect(): void {
    if (this.conn) {
      this.conn.close();
      this.conn = null;
    }
    if (this.term) {
      this.term.dispose();
      this.term = null;
    }
    this.connected = false;
    this.connecting = false;
    this.connectionClosed = true;
  }

  private initTerm(): void {
    if (this.term) {
      this.term.dispose();
    }

    this.term = new Terminal({
      fontSize: 14,
      fontFamily: 'Consolas, "Courier New", monospace',
      cursorBlink: true
    });

    this.fitAddon = new FitAddon();
    this.term.loadAddon(this.fitAddon);
    this.term.open(this.container);
    this.debouncedFit = debounce(() => {
      this.fitAddon?.fit();
    }, 100);
    this.debouncedFit();
    window.addEventListener('resize', () => this.debouncedFit!());

    this.term.onData(this.onTerminalSendString.bind(this));
    this.term.onResize(this.onTerminalResize.bind(this));

    // Focus on connection
    this.term.focus();
  }

  private onConnectionOpen(sessionId: string): void {
    const startData: ShellFrame = { Op: 'bind', SessionID: sessionId };
    this?.conn?.send(JSON.stringify(startData));
    this.connected = true;
    this.connecting = false;
    this.connectionClosed = false;

    // Make sure the terminal is with correct display size.
    this.onTerminalResize();

    this.term?.focus();
  }

  private handleFrame(frame: ShellFrame): void {
    if (frame.Op === 'stdout') {
      this.term?.write(frame.Data || '');
    }
    if (frame.Op === 'toast') {
      this.onToast(frame.Data || '');
    }
  }

  private onConnectionMessage(evt: SJSMessageEvent): void {
    const msg = JSON.parse(evt.data) as ShellFrame;
    this.handleFrame(msg);
  }

  private onConnectionClose(evt?: SJSCloseEvent): void {
    if (!this.connected) return;
    this.connected = false;
    this.connecting = false;
    this.connectionClosed = true;
    this.onToast(evt?.reason || 'Connection closed');
  }

  private onTerminalSendString(str: string): void {
    if (this.connected && this.conn) {
      this.conn.send(
        JSON.stringify({
          Op: 'stdin',
          Data: str,
          Cols: this.term?.cols,
          Rows: this.term?.rows
        })
      );
    }
  }

  private onTerminalResize(): void {
    if (this.connected && this.conn) {
      this.conn.send(
        JSON.stringify({
          Op: 'resize',
          Cols: this.term?.cols,
          Rows: this.term?.rows
        })
      );
    }
  }
}
