import {useContext} from "react";
import {VkAdminViewContext} from "./context";

type useDataProviderProps<TResponse> = {
    data?: TResponse,
    setData?: (data: TResponse) => void
}

export function useDataProvider<TResponse>(): useDataProviderProps<TResponse> {
    const context = useContext(VkAdminViewContext);
    if (!context) {
        throw new Error("useDataProvider must be used within a VkAdminViewProvider");
    }
    return {
        data: context.data as TResponse,
        setData: context.setData as (data: TResponse) => void
    }
}
