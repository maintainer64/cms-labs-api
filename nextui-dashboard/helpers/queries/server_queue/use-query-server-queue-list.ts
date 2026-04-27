/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type ServerQueueServerQueueListRequest,
  type ServerQueueServerQueueListResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<ServerQueueServerQueueListRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<ServerQueueServerQueueListResponse['result']>;

export const useQueryServerQueueList = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'server_queue.list', params, { retry: 3 })
  );
};
