import { createRoot } from 'react-dom/client';
import '@/styles/globals.css';
import '@/styles/reactflow.css';
import { App } from '@/app/App';

createRoot(document.getElementById('root')!).render(<App />);
