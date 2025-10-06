/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesLTIFormEditRequest,
  type UsecasesLTIFormEditResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesLTIFormEditRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesLTIFormEditResponse['result']>;

export const useQueryLtiFormUpsert = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'lti_form.upsert', params, { retry: 3 })
  );
};
