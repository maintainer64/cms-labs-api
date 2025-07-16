import { Terminal } from '@xterm/xterm';
import { FitAddon } from '@xterm/addon-fit';
import debounce from 'lodash/debounce';
import { ShellFrame } from './types';

type Function = (...args: any[]) => any;

export class MockTerminalService {
  private term: Terminal | null = null;
  private fitAddon: FitAddon | null = null;
  private debouncedFit: Function | null = null;
  private connected = false;
  private mockSessionId = 'mock-session';
  private onToast: (message: string) => void = () => {};
  private mockBuffer: string = '';

  constructor(private container: HTMLElement) {}

  setOnToast(callback: (message: string) => void) {
    this.onToast = callback;
  }

  async connect(namespace: string, podName: string, container: string): Promise<void> {
    this.initTerm();
    this.connected = true;

    // Симуляция onConnectionOpen
    const startData: ShellFrame = { Op: 'bind', SessionID: this.mockSessionId };
    this.handleFrame(startData);

    this.term?.writeln(`MOCK Connected to ${namespace}/${podName}/${container}`);
    this.term?.writeln('$ ');
    this.onToast('MOCK Connection established');

    // Симуляция resize
    this.onTerminalResize();
  }

  disconnect(): void {
    this.term?.writeln('\r\nMOCK Connection closed');
    if (this.term) {
      this.term.dispose();
      this.term = null;
    }
    this.connected = false;
    this.onToast('MOCK Disconnected');
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
  }

  private handleFrame(frame: ShellFrame): void {
    if (frame.Op === 'stdout') {
      this.term?.write(frame.Data || '');
    }
    if (frame.Op === 'toast') {
      this.onToast(frame.Data || '');
    }
  }

  private onTerminalSendString(str: string): void {
    if (!this.connected) return;

    // Симуляция stdin: эхо ввода + случайный toast
    this.mockBuffer += str;
    if (str.includes('\r')) {
      // Enter
      this.term?.writeln(`\r\nMOCK Echo: ${this.mockBuffer.trim()}`);
      this.term?.writeln('$ ');
      if (Math.random() > 0.5) {
        this.onToast(`MOCK Toast: Command "${this.mockBuffer.trim()}" processed`);
      }
      this.mockBuffer = '';
    } else {
      this.term?.write(str);
    }
  }

  private onTerminalResize(): void {
    // Симуляция resize
    this.term?.writeln(`\r\nMOCK Resized to ${this.term?.cols}x${this.term?.rows}`);
  }
}
