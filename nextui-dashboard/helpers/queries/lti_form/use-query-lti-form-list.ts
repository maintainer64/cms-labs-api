/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesAuthProviderListRequest,
  type UsecasesAuthProviderListResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesAuthProviderListRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesAuthProviderListResponse['result']>;

export const useQueryAuthProviderList = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'lti_form.list', params, { retry: 3 })
  );
};
