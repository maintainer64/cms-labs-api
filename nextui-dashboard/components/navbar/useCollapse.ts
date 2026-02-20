import { useQueryUserGlobalStoreGet } from '@/helpers/queries/user/use-query-user-global-store-get';
import { useMutationUserGlobalStoreSet } from '@/helpers/queries/user/use-mutation-user-global-store-set';

const useCollapsedBrowser = () => {
  const globalStoreQuery = useQueryUserGlobalStoreGet({});
  const { mutate } = useMutationUserGlobalStoreSet();
  // @ts-ignore
  const collapsed = globalStoreQuery?.data?.['collapsed'] as boolean;
  return {
    collapsed: collapsed,
    setCollapsed: (collapsed: boolean) => {
      mutate({ ...(globalStoreQuery?.data || {}), collapsed: collapsed });
    }
  };
};
export default useCollapsedBrowser;
