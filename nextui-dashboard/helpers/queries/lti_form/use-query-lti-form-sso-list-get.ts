/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithoutAuth,
  type UsecasesAuthProviderListSSORequest,
  type UsecasesAuthProviderListSSOResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesAuthProviderListSSORequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesAuthProviderListSSOResponse['result']>;

export const useQueryAuthProviderSsoListGet = (params: Params) => {
  return useQuery(
      transportWithoutAuth.t.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'lti_form.sso_list_get', params, { retry: 3 })
  );
};
