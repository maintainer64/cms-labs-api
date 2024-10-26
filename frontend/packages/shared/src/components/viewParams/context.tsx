import {createContext, FC, ReactNode, useState} from 'react';

type VkAdminViewContextProps = {
    data?: any
    setData?: (data: any) => void
};
export const VkAdminViewContext = createContext<VkAdminViewContextProps>({
    data: undefined,
    setData: undefined,
});

type VkAdminViewProviderProps = {
    children: ReactNode;
}

export const VkAdminViewProvider: FC<VkAdminViewProviderProps> = ({children}) => {
    const [data, setData] = useState<any>(
        undefined
    );
    const contextValue: VkAdminViewContextProps = {
        data,
        setData,
    };
    return (
        <VkAdminViewContext.Provider value={contextValue}>
            {children}
        </VkAdminViewContext.Provider>
    );
};
