import { Terminal } from '@xterm/xterm';
import { FitAddon } from '@xterm/addon-fit';
import SockJS from 'sockjs-client';
import debounce from 'lodash/debounce';
import { ShellFrame } from '@/components/topology/terminal/service/context';
import { ContainerLabMiddleware } from '@/components/topology/terminal/service/middleware-containerlab';
import { ContainerConnectParams } from '@/components/topology/terminal/service/middleware-type';

type Function = (...args: any[]) => any;

export class TerminalService {
  private term: Terminal | null = null;
  private fitAddon: FitAddon | null = null;
  private sock: WebSocket | null = null;
  private debouncedFit: Function | null = null;
  private connected = false;
  private connecting = false;
  private onToast: (message: string) => void = () => {};
  private onReconnect: () => void = () => {};
  private middlewares = [new ContainerLabMiddleware()];

  constructor(
    private container: HTMLElement,
    private locale: any
  ) {}

  setOnToast(callback: (message: string) => void) {
    this.onToast = callback;
  }

  setOnReconnect(callback: () => void) {
    this.onReconnect = callback;
  }

  async connect(connectParams?: ContainerConnectParams): Promise<void> {
    if (!connectParams?.sessionId) {
      this.onReconnect();
    }
    this.connecting = true;

    try {
      this.sock = new SockJS(connectParams?.url || '', null, {
        sessionId: () => connectParams?.sessionId ?? ''
      });

      this.sock.onopen = () => {
        this.onSockOpen(connectParams);
      };

      this.sock.onmessage = (e) => {
        this.onSockMessage(e.data, connectParams);
      };

      this.sock.onclose = (e) => {
        this.onSockClose({ reason: e.reason || 'Connection closed' });
      };

      this.sock.onerror = (error) => {
        this.onToast(`Socket error: ${error}`);
        this.onSockClose({ reason: 'Socket error' });
      };

      this.initTerm();
    } catch (error) {
      this.connecting = false;
      this.onToast(`${this.locale.ConnectionError}: ${error instanceof Error ? error.message : String(error)}`);
      throw error;
    }
  }

  disconnect(): void {
    if (this.sock) {
      this.sock.close();
      this.sock = null;
    }
    if (this.term) {
      this.term.dispose();
      this.term = null;
    }
    this.connected = false;
    this.connecting = false;
  }

  private initTerm(): void {
    if (this.term) {
      this.term.dispose();
    }

    this.term = new Terminal({
      fontSize: 14,
      fontFamily: 'Consolas, "Courier New", monospace',
      cursorBlink: true,
      allowProposedApi: true
    });

    this.fitAddon = new FitAddon();
    this.term.loadAddon(this.fitAddon);
    this.term.open(this.container);

    this.debouncedFit = debounce(() => {
      this.fitAddon?.fit();
      this.onTerminalResize();
    }, 100);

    this.debouncedFit();
    window.addEventListener('resize', () => this.debouncedFit!());

    this.term.onData(this.onTerminalSendString.bind(this));
    this.term.onResize(this.onTerminalResize.bind(this));
  }

  private onSockOpen(connectParams?: ContainerConnectParams): void {
    if (!this.sock) return;
    if (!connectParams?.sessionId) return;

    const bindFrame: ShellFrame = {
      Op: 'bind',
      SessionID: connectParams.sessionId,
      Cols: this.term?.cols || 80,
      Rows: this.term?.rows || 24
    };

    this.sock.send(JSON.stringify(bindFrame));
    this.connected = true;
    this.connecting = false;
    this.onToast(this.locale.TerminalConnected);
    for (const middleware of this.middlewares) {
      middleware.onOpen(this.sock, connectParams);
    }
  }

  private onSockMessage(data: string, connectParams?: ContainerConnectParams): void {
    try {
      const frame: ShellFrame = JSON.parse(data);
      this.handleFrame(frame);
      for (const middleware of this.middlewares) {
        middleware.onData(this.sock, frame, connectParams);
      }
    } catch (error) {
      console.error('Error parsing message:', error);
    }
  }

  private handleFrame(frame: ShellFrame): void {
    if (!this.term) return;

    switch (frame.Op) {
      case 'stdout':
        this.term.write(frame.Data || '');
        break;
      case 'stderr':
        this.term.write('\x1b[31m' + (frame.Data || '') + '\x1b[0m');
        break;
      case 'toast':
        this.onToast(frame.Data || '');
        break;
      case 'resize':
        if (frame.Cols && frame.Rows) {
          this.term.resize(frame.Cols, frame.Rows);
        }
        break;
    }
  }

  private onSockClose(e: { reason: string }): void {
    this.connected = false;
    this.connecting = false;
    this.onToast(`${this.locale.ConnectionError}: ${e.reason || this.locale.UnknownReason}`);
    this.onReconnect();
  }

  private onTerminalSendString(str: string): void {
    if (!this.connected || !this.sock || !this.term) return;

    const frame: ShellFrame = {
      Op: 'stdin',
      Data: str,
      Cols: this.term.cols,
      Rows: this.term.rows
    };

    this.sock.send(JSON.stringify(frame));
  }

  private onTerminalResize(): void {
    if (!this.connected || !this.sock || !this.term) return;

    const frame: ShellFrame = {
      Op: 'resize',
      Cols: this.term.cols,
      Rows: this.term.rows
    };

    this.sock.send(JSON.stringify(frame));
  }

  public fit(): void {
    this.debouncedFit?.();
  }
}
