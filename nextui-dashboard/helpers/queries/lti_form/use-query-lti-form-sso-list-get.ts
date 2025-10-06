/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesLTIFormListSSORequest,
  type UsecasesLTIFormListSSOResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesLTIFormListSSORequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesLTIFormListSSOResponse['result']>;

export const useQueryLtiFormSsoListGet = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'lti_form.sso_list_get', params, { retry: 3 })
  );
};
