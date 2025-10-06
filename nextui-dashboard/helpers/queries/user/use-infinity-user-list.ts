/*
Template auto generated with params from openapi.json
*/
import { useInfiniteQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  ModelsUser,
  transportWithAuth,
  type UsecasesUserListRequest,
  type UsecasesUserListResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';
import { RpcTransport } from '@/helpers/api/core/request';

type Params = CamelCasedPropertiesDeep<UsecasesUserListRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesUserListResponse['result']>;

export const useInfinityUserList = (params?: Params) => {
  return useInfiniteQuery({
    queryKey: RpcTransport.getInfiniteQueryKey(CoreJsonRpcPath, 'user.list', params ?? {}),
    queryFn: ({ pageParam }) => {
      return transportWithAuth.rpc<Response>(CoreJsonRpcPath, {
        method: 'user.list',
        params: {
          limit: params?.limit ?? 100,
          offset: pageParam,
          search: (params?.search?.length || '') < 3 ? '' : params?.search
        } as Params
      });
    },
    refetchOnWindowFocus: false,
    refetchInterval: false,
    gcTime: 0,
    getNextPageParam: (response, pages) => {
      const totalCount = response?.totalCount || 0;
      const count = pages.flatMap((p) => p?.model).length;
      return totalCount && totalCount > count ? count : undefined;
    },
    initialPageParam: 0,
    retry: 3
  });
};

export interface UserItem extends CamelCasedPropertiesDeep<ModelsUser> {
  roles?: Array<number>;
}

export function MapUserItem(user?: CamelCasedPropertiesDeep<ModelsUser>, roles?: Array<number>) {
  if (!user) return {} as UserItem;
  const newModel = user as UserItem;
  newModel.roles = roles;
  return newModel;
}
