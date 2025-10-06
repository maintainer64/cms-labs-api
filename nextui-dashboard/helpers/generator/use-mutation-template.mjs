function templateUseMutation(params) {
    return `/*
Template auto generated with params from openapi.json
*/
import {useMutation} from '@tanstack/react-query';
import {
    ${params.path},
    ${params.transport},
    ${params.requestRPC},
    ${params.responseRPC},
} from '@/helpers/api';
import {TMutationCustomOptions} from '@/helpers/queries/types';
import queryClient from "@/helpers/queries/base";
import {CamelCasedPropertiesDeep} from "type-fest";

type Params = CamelCasedPropertiesDeep<${params.requestRPC}["params"]>;
type Response = CamelCasedPropertiesDeep<${params.responseRPC}["result"]>;


export const useMutation${params.methodJs} = (
    options: TMutationCustomOptions<Response, Params> = {}
) => {
    return useMutation<Response, unknown, Params>({
        // @ts-expect-error: return nullable value
        mutationFn: (params: Params) => {
            if (!params?.id) return null;
            return ${params.transport}.rpc(
                ${params.path},
                {
                    method: "${params.methodName}",
                    params: params,
                }
            );
        },
        ...options,
        async onSuccess(...args) {
            if (options.onSuccess) {
                options.onSuccess(...args);
            }
            // Add invalidateQueries
        }
    });
};
`;
}

export default templateUseMutation;
