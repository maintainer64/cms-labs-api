export interface TerminalClient {
  id: string;
  label?: string;
  index: number;
  isVisible: boolean;
  isOpacity: boolean;
  zIndex: number;
}

type TerminalAction =
  | { type: 'ADD_CLIENT'; payload: { id: string; label?: string } }
  | { type: 'REMOVE_CLIENT'; payload: { id: string } }
  | { type: 'TOGGLE_VISIBILITY'; payload: { id: string; onChange?: (visibility?: boolean) => void } }
  | { type: 'TOGGLE_OPACITY'; payload: { id: string } }
  | { type: 'BRING_TO_FRONT'; payload: { id: string } }
  | { type: 'SHOW_NODE_ACTIONS_MODAL'; payload: { id: string; label?: string; x: number; y: number } }
  | { type: 'HIDE_NODE_ACTIONS_MODAL' };

export type TerminalActionFunc = (action: TerminalAction) => void;

interface TerminalState {
  clients: TerminalClient[];
  selectedNodeId: string | null;
  selectedNodeLabel: string | null;
  selectedNodeX: number;
  selectedNodeY: number;
}

export const terminalInitialState: TerminalState = {
  clients: [],
  selectedNodeId: null,
  selectedNodeLabel: null,
  selectedNodeX: 0,
  selectedNodeY: 0
};

export function terminalReducer(state: TerminalState, action: TerminalAction): TerminalState {
  switch (action.type) {
    case 'ADD_CLIENT': {
      const { id, label } = action.payload;
      if (state.clients.some((c) => c.id === id)) return state;

      return {
        ...state,
        clients: [
          ...state.clients,
          {
            id,
            label,
            index: state.clients.length,
            isVisible: false,
            isOpacity: false,
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

      const activeIndex = currentClients.findIndex((c) => c.id === id);
      if (activeIndex === -1) return state;

      const sortedClients = [...currentClients].sort((a, b) => a.zIndex - b.zIndex);

      const activeZIndex = currentClients[activeIndex].zIndex;

      if (activeZIndex === sortedClients[sortedClients.length - 1]?.zIndex) {
        return state;
      }

      let currentZ = 1;
      const updatedClients = sortedClients.map((client) => {
        if (client.id === id) {
          return { ...client, zIndex: sortedClients.length };
        }
        return { ...client, zIndex: currentZ++ };
      });

      return {
        ...state,
        clients: updatedClients.sort((a, b) => {
          const aIndex = currentClients.findIndex((c) => c.id === a.id);
          const bIndex = currentClients.findIndex((c) => c.id === b.id);
          return aIndex - bIndex;
        })
      };
    }

    case 'SHOW_NODE_ACTIONS_MODAL': {
      return {
        ...state,
        selectedNodeId: action.payload.id,
        selectedNodeLabel: action.payload.label || null,
        selectedNodeX: action.payload.x,
        selectedNodeY: action.payload.y
      };
    }

    case 'HIDE_NODE_ACTIONS_MODAL': {
      return {
        ...state,
        selectedNodeId: null,
        selectedNodeLabel: null,
        selectedNodeX: 0,
        selectedNodeY: 0
      };
    }

    default:
      return state;
  }
}
