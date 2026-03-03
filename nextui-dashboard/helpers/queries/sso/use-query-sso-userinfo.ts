/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import { type AuthSwaggerSSOTokenPublicData, CoreJsonRpcPath, transportWithAuth } from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';
import { RpcTransport } from '@/helpers/api/core/request';
import { objectToCamel } from 'ts-case-convert';

type Response = CamelCasedPropertiesDeep<AuthSwaggerSSOTokenPublicData>;

const method = '/api/v1/sso/userinfo';

export const useQuerySsoUserInfo = () => {
  return useQuery({
    queryKey: RpcTransport.getQueryKey(method, '', {}),
    queryFn: ({ signal }: { signal?: AbortSignal }) => {
      return transportWithAuth
        .getTransport()
        .get(method, { signal })
        .then((resp) => {
          return objectToCamel(resp.data) as Response;
        });
    },
    retry: 1
  });
};
