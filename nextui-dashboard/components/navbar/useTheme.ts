import { useTheme as useThemeNext } from 'next-themes';
import { useEffect } from 'react';
import { useQueryUserGlobalStoreGet } from '@/helpers/queries/user/use-query-user-global-store-get';
import { useMutationUserGlobalStoreSet } from '@/helpers/queries/user/use-mutation-user-global-store-set';

export type ThemeType = 'dark' | 'light';

const useThemeBrowser = () => {
  const globalStoreQuery = useQueryUserGlobalStoreGet({});
  const { mutate } = useMutationUserGlobalStoreSet();
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
