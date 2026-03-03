/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type UsecasesLTIAttemptEditRequest,
  type UsecasesLTIAttemptEditResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesLTIAttemptEditRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesLTIAttemptEditResponse['result']>;

export const useQueryLtiAttemptUpdate = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'lti_attempt.update', params, { retry: 3 })
  );
};
