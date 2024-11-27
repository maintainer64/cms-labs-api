import { ReactNode, createContext, useEffect, useState } from 'react';

export type AlertType = 'info' | 'danger' | 'success' | 'warning' | 'dark';

export interface AlertContextProps {
  type: AlertType;
  message: string;
  description?: string;
}

type AlertContext = {
  showAlert: (alert: AlertContextProps) => void;
  hideAlert: (index?: number) => void;
  alertMessages: AlertContextProps[];
};

type AlertContextProvider = {
  children: ReactNode;
};

// Create a new context for the Alert
export const AlertContext = createContext<AlertContext>({
  showAlert: () => {},
  hideAlert: () => {},
  alertMessages: []
});

export const AlertProvider = ({ children }: AlertContextProvider) => {
  const [alertMessage, setAlertMessage] = useState<AlertContextProps[]>([]);

  // UseEffect hook to remove the first alert message after 5 seconds
  useEffect(() => {
    const timeout = setTimeout(() => {
      setAlertMessage((prevItems) => {
        if (prevItems.length > 0) {
          return prevItems.filter((_, i) => i != 0);
        }
        return prevItems;
      });
    }, 5 * 1000);
    return () => {
      clearTimeout(timeout);
    };
  }, [alertMessage]);

  // Context value containing the showAlert function
  const contextValue: AlertContext = {
    showAlert: (alertMessage) => {
      setAlertMessage((prev) => [...prev, alertMessage]);
    },
    hideAlert: (index?: number) => {
      setAlertMessage((prev) => prev.filter((_, i) => i != index && index !== undefined));
    },
    alertMessages: alertMessage
  };

  return <AlertContext.Provider value={contextValue}>{children}</AlertContext.Provider>;
};
