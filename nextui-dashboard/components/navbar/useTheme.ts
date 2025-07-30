import { useGlobalStoreGet, useGlobalStoreSet } from '@/helpers/queries/users/store';
import { useTheme as useThemeNext } from 'next-themes';
import { useEffect } from 'react';

export type ThemeType = 'dark' | 'light';

const useThemeBrowser = () => {
  const globalStoreQuery = useGlobalStoreGet();
  const { mutate } = useGlobalStoreSet();
  const { setTheme } = useThemeNext();
  // @ts-ignore
  const theme = (globalStoreQuery?.data?.['theme'] || 'light') as ThemeType;
  useEffect(() => {
    setTheme(theme);
  }, [theme]);
  return {
    theme: theme as ThemeType,
    setTheme: (theme: ThemeType) => {
      mutate({ ...(globalStoreQuery?.data || {}), theme: theme });
    }
  };
};
export default useThemeBrowser;
