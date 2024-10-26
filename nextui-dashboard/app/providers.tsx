"use client";
import * as React from "react";
import {NextUIProvider} from "@nextui-org/system";
import {ThemeProvider as NextThemesProvider} from "next-themes";
import {ThemeProviderProps} from "next-themes/dist/types";
import {UserProfileProvider} from "@/components/providers/auth-jwt/context";
import {QueryClientProvider} from "@tanstack/react-query";
import queryClient from "@/helpers/queries/base";
import {BrowserRouter} from "react-router-dom";
import RoutesDynamic from "@/app/routes";
import GroupAlerts from "@/components/alerts/groupAlerts";
import {AlertProvider} from "@/components/alerts/context";

export interface ProvidersProps {
    themeProps?: ThemeProviderProps;
}

export function Providers({themeProps}: ProvidersProps) {
    return (
        <QueryClientProvider client={queryClient}>
            <BrowserRouter>
                <AlertProvider>
                    <UserProfileProvider>
                        <NextUIProvider>
                            <NextThemesProvider
                                defaultTheme='system'
                                attribute='class'
                                {...themeProps}>
                                <GroupAlerts>
                                    <RoutesDynamic/>
                                </GroupAlerts>
                            </NextThemesProvider>
                        </NextUIProvider>
                    </UserProfileProvider>
                </AlertProvider>
            </BrowserRouter>
        </QueryClientProvider>
    );
}
