function templateUseQuery(params) {
    return `/*
Template auto generated with params from openapi.json
*/
import {useQuery} from '@tanstack/react-query';
import {
    ${params.path},
    ${params.transport},
    type ${params.requestRPC},
    type ${params.responseRPC},
} from '@/helpers/api';
import {CamelCasedPropertiesDeep} from "type-fest";

type Params = CamelCasedPropertiesDeep<${params.requestRPC}["params"]> & object;
type Response = CamelCasedPropertiesDeep<${params.responseRPC}["result"]>;

export const useQuery${params.methodJs} = (params: Params) => {
    return useQuery(transportWithAuth.getQueryOptions<Response, Params>(
        ${params.path}, "${params.methodName}", params, {retry: 3}
    ))
};
`;
}

export default templateUseQuery;
