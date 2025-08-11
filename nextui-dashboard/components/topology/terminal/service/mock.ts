import { Terminal } from '@xterm/xterm';
import { FitAddon } from '@xterm/addon-fit';
import debounce from 'lodash/debounce';
import { ShellFrame } from '@/components/topology/terminal/service/context';

type Function = (...args: any[]) => any;

export class MockTerminalService {
  private term: Terminal | null = null;
  private fitAddon: FitAddon | null = null;
  private debouncedFit: Function | null = null;
  private connected = false;
  private connecting = false;
  private mockSessionId = 'mock-session-' + Math.random().toString(36).substr(2, 8);
  private onToast: (message: string) => void = () => {};
  private inputBuffer: string = '';
  private isCommandRunning = false;
  private onReconnect: () => void = () => {};

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

  async connect(sessionId?: string, url?: string, type?: string): Promise<void> {
    this.connecting = true;

    // Имитация задержки подключения
    await new Promise((resolve) => setTimeout(resolve, 500));

    this.initTerm();
    this.connected = true;
    this.connecting = false;

    // Имитация успешного подключения
    this.onSockOpen(sessionId || this.mockSessionId);

    // Вывод приветственного сообщения
    this.term?.writeln('\x1b[32mConnected to mock terminal service\x1b[0m');
    this.term?.writeln(`Session ID: ${sessionId || this.mockSessionId}`);
    this.term?.writeln('Try typing commands like "help", "date", "ping"');
    this.writePrompt();

    this.onToast('Mock terminal connected successfully');
  }

  disconnect(): void {
    if (this.connected) {
      this.term?.writeln('\r\n\x1b[31mDisconnected from mock terminal\x1b[0m');
      this.onToast('Mock terminal disconnected');
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

  private onSockOpen(sessionId: string): void {
    const bindFrame: ShellFrame = {
      Op: 'bind',
      SessionID: sessionId,
      Cols: this.term?.cols || 80,
      Rows: this.term?.rows || 24
    };

    // Имитация ответа сервера
    setTimeout(() => {
      this.handleFrame({
        Op: 'stdout',
        Data: '\x1b[32mSession bound successfully\x1b[0m\r\n'
      });
      this.writePrompt();
    }, 100);
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

  private onTerminalSendString(str: string): void {
    if (!this.connected || !this.term || this.isCommandRunning) return;

    // Обработка специальных клавиш
    if (str === '\r') {
      // Enter
      this.processCommand();
      return;
    } else if (str === '\x7f') {
      // Backspace
      if (this.inputBuffer.length > 0) {
        this.inputBuffer = this.inputBuffer.slice(0, -1);
        this.term.write('\b \b');
      }
      return;
    } else if (str.charCodeAt(0) < 32) {
      // Другие управляющие символы
      return;
    }

    // Обычные символы
    this.inputBuffer += str;
    this.term.write(str);
  }

  private processCommand(): void {
    if (!this.term) return;

    const command = this.inputBuffer.trim();
    this.inputBuffer = '';
    this.term.write('\r\n');
    this.isCommandRunning = true;

    // Обработка команд
    switch (command.toLowerCase()) {
      case '':
        break;
      case 'help':
        this.term.writeln('Available mock commands:');
        this.term.writeln('  help     - Show this help');
        this.term.writeln('  date     - Show current date/time');
        this.term.writeln('  ping     - Test connection');
        this.term.writeln('  clear    - Clear terminal');
        this.term.writeln('  error    - Generate error message');
        break;
      case 'date':
        this.term.writeln(new Date().toString());
        break;
      case 'ping':
        this.term.write('PONG');
        this.onToast('Received PONG response');
        break;
      case 'clear':
        this.term.clear();
        break;
      case 'exit':
        this.connected = false;
        this.connecting = false;
        this.onToast(`Connection closed: 'Unknown reason'`);
        this.onReconnect();
        break;
      case 'error':
        this.handleFrame({
          Op: 'stderr',
          Data: 'This is a mock error message\r\n'
        });
        this.onToast('Mock error generated');
        break;
      default:
        this.term.writeln(`\x1b[31mCommand not found: ${command}\x1b[0m`);
    }

    this.isCommandRunning = false;
    this.writePrompt();
  }

  private writePrompt(): void {
    if (!this.term) return;
    this.term.write('$ ');
  }

  private onTerminalResize(): void {
    if (!this.term) return;

    // Имитация ответа сервера на resize
    setTimeout(() => {
      this.handleFrame({
        Op: 'stdout',
        Data: `\r\n\x1b[33mTerminal resized to ${this.term?.cols}x${this.term?.rows}\x1b[0m\r\n`
      });
      this.writePrompt();
    }, 100);
  }

  public fit(): void {
    this.debouncedFit?.();
  }
}
