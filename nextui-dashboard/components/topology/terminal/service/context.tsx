export interface TerminalClient {
  id: string;
  name?: string;
  index: number;
  isVisible: boolean;
  namespace?: string;
  isOpacity: boolean;
  zIndex: number;
}

export interface ShellFrame {
  Op: string;
  SessionID?: string;
  Data?: string;
  Cols?: number;
  Rows?: number;
}

type TerminalAction =
  | { type: 'ADD_CLIENT'; payload: { id: string; namespace?: string; name?: string } }
  | { type: 'REMOVE_CLIENT'; payload: { id: string } }
  | { type: 'TOGGLE_VISIBILITY'; payload: { id: string; onChange?: (visibility?: boolean) => void } }
  | { type: 'TOGGLE_OPACITY'; payload: { id: string } }
  | { type: 'BRING_TO_FRONT'; payload: { id: string } };

export type TerminalActionFunc = (action: TerminalAction) => void;

interface TerminalState {
  clients: TerminalClient[];
}

export const terminalInitialState: TerminalState = {
  clients: []
};

export function terminalReducer(state: TerminalState, action: TerminalAction): TerminalState {
  switch (action.type) {
    case 'ADD_CLIENT': {
      const { id, namespace, name } = action.payload;
      if (state.clients.some((c) => c.id === id)) return state;

      return {
        ...state,
        clients: [
          ...state.clients,
          {
            id,
            name,
            index: state.clients.length,
            isVisible: false,
            isOpacity: false,
            namespace,
            zIndex: state.clients.length + 1
          }
        ]
      };
    }

    case 'REMOVE_CLIENT': {
      return {
        ...state,
        clients: state.clients.filter((c) => c.id !== action.payload.id)
      };
    }

    case 'TOGGLE_VISIBILITY': {
      const client = state.clients.find((c) => c.id === action.payload.id);
      action.payload?.onChange?.(!client?.isVisible);
      return {
        ...state,
        clients: state.clients.map((c) => (c.id === action.payload.id ? { ...c, isVisible: !c.isVisible } : c))
      };
    }

    case 'TOGGLE_OPACITY': {
      return {
        ...state,
        clients: state.clients.map((c) => (c.id === action.payload.id ? { ...c, isOpacity: !c.isOpacity } : c))
      };
    }

    case 'BRING_TO_FRONT': {
      const { id } = action.payload;
      const currentClients = [...state.clients];

      // Находим индекс клиента, который нужно вывести на передний план
      const activeIndex = currentClients.findIndex((c) => c.id === id);
      if (activeIndex === -1) return state;

      // Сортируем клиентов по текущему zIndex
      const sortedClients = [...currentClients].sort((a, b) => a.zIndex - b.zIndex);

      // Находим текущий zIndex активного клиента
      const activeZIndex = currentClients[activeIndex].zIndex;

      // Если клиент уже на переднем плане, ничего не делаем
      if (activeZIndex === sortedClients[sortedClients.length - 1]?.zIndex) {
        return state;
      }

      // Перераспределяем zIndex
      let currentZ = 1;
      const updatedClients = sortedClients.map((client) => {
        if (client.id === id) {
          // Активному клиенту даем максимальный zIndex
          return { ...client, zIndex: sortedClients.length };
        }
        // Остальным распределяем по порядку
        return { ...client, zIndex: currentZ++ };
      });

      return {
        ...state,
        clients: updatedClients.sort((a, b) => {
          // Восстанавливаем исходный порядок клиентов
          const aIndex = currentClients.findIndex((c) => c.id === a.id);
          const bIndex = currentClients.findIndex((c) => c.id === b.id);
          return aIndex - bIndex;
        })
      };
    }

    default:
      return state;
  }
}
