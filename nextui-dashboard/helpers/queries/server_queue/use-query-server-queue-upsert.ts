/*
Template auto generated with params from openapi.json
*/
import { useQuery } from '@tanstack/react-query';
import {
  CoreJsonRpcPath,
  transportWithAuth,
  type ServerQueueServerQueueUpsertRequest,
  type ServerQueueServerQueueUpsertResponse
} from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';

type Params = CamelCasedPropertiesDeep<ServerQueueServerQueueUpsertRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<ServerQueueServerQueueUpsertResponse['result']>;

export const useQueryServerQueueUpsert = (params: Params) => {
  return useQuery(
    transportWithAuth.getQueryOptions<Response, Params>(CoreJsonRpcPath, 'server_queue.upsert', params, { retry: 3 })
  );
};
