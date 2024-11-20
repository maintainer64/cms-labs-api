import {useMutation} from "@tanstack/react-query";
import {
    models_PNETServer,
    type PostV1LtiFormUpsertResponse,
    postV1PnetServerUpsert,
    PostV1PnetServerUpsertResponse,
} from "@/helpers/api";
import {TFormikData, TMutationCustomOptions} from "@/helpers/queries/types";
import queryClient from "../base";

export const usePnetServerUpsert = (
    options: TMutationCustomOptions<PostV1PnetServerUpsertResponse, unknown, TFormikData<models_PNETServer>> = {}
) => {
    return useMutation<PostV1LtiFormUpsertResponse, unknown, TFormikData<models_PNETServer>>({
        // @ts-ignore
        mutationFn: ({values}: TFormikData<models_PNETServer>) => {
            if (values === null) return null;
            return postV1PnetServerUpsert({
                form: {
                    id: values.id,
                    is_active: values.is_active,
                    name: values.name || '',
                    minutes_for_disconnect: values.minutes_for_disconnect || 0,
                    unit_rate: values.unit_rate,
                    url: values.url || '',
                }
            })
        },
        ...options,
        async onSuccess(...args) {
            await queryClient.invalidateQueries({queryKey: ['postV1PnetServerList']});
            await queryClient.invalidateQueries({queryKey: ['postV1PnetServerGet']});
            if (options.onSuccess) {
                options.onSuccess(...args);
            }
        },
    });
}
