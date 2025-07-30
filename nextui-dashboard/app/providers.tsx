'use client';
import * as React from 'react';
import { HeroUIProvider } from '@heroui/system';
import { ThemeProvider as NextThemesProvider } from 'next-themes';
import { ThemeProviderProps } from 'next-themes/dist/types';
import { UserProfileProvider } from '@/components/providers/auth-jwt/context';
import { QueryClientProvider } from '@tanstack/react-query';
import queryClient from '@/helpers/queries/base';
import { BrowserRouter } from 'react-router-dom';
import RoutesDynamic from '@/app/routes';
import { ToastProvider } from '@heroui/toast';
import { zIndexClassToast } from '@/components/providers/const';

export interface ProvidersProps {
  themeProps?: ThemeProviderProps;
}

export function Providers({ themeProps }: ProvidersProps) {
  return (
    <QueryClientProvider client={queryClient}>
      <BrowserRouter basename='/'>
        <UserProfileProvider>
          <HeroUIProvider>
            <NextThemesProvider defaultTheme='system' attribute='class' {...themeProps}>
              <ToastProvider
                regionProps={{
                  classNames: { base: zIndexClassToast }
                }}
              />
              <RoutesDynamic />
            </NextThemesProvider>
          </HeroUIProvider>
        </UserProfileProvider>
      </BrowserRouter>
    </QueryClientProvider>
  );
}
