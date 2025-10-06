/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesLTIFormListRequest,
  type UsecasesLTIFormListResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesLTIFormListRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesLTIFormListResponse['result']>;

export const useQueryLtiFormList = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'lti_form.list', params, { retry: 3 })
  );
};
