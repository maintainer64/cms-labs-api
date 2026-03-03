/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  ClabgateJsonRpcPath,
  transportWithAuth,
  type UsecasesTaskListRequest,
  type UsecasesTaskListResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesTaskListRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesTaskListResponse['result']>;

export const useQueryTaskList = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(ClabgateJsonRpcPath, 'task.list', params, { retry: 3 })
  );
};
