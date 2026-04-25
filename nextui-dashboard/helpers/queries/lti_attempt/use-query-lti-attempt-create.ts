/*
Template auto generated with params from openapi.json
*/
import {useQuery} from '@tanstack/react-query';
import {
    CoreJsonRpcPath,
    transportWithAuth,
    type UsecasesLTIAttemptCreateRequest,
    type UsecasesLTIAttemptCreateResponse
} from '@/helpers/api';
import {CamelCasedPropertiesDeep} from 'type-fest';

type Params = CamelCasedPropertiesDeep<UsecasesLTIAttemptCreateRequest['params']> & object;
type Response = CamelCasedPropertiesDeep<UsecasesLTIAttemptCreateResponse['result']>;

export const useQueryLtiAttemptCreate = (params: Params) => {
        return useQuery(
            transportWithAuth.getQueryOptions<Response, Params>(
                CoreJsonRpcPath,
                'lti_attempt.create',
                params, {
                    retry: 3,
                    refetchInterval: 3000 // 3 seconds
                }
            )
        );
    }
;
