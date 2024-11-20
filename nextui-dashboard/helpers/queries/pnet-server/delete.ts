import {useMutation} from "@tanstack/react-query";
import {
    postV1LtiFormDelete,
    type PostV1PnetServerDeleteResponse,
    usecases_PNETServerDeleteInputDTO,
} from "@/helpers/api";
import {TMutationCustomOptions} from "@/helpers/queries/types";
import queryClient from "../base";

export const usePnetServerDelete = (
    options: TMutationCustomOptions<PostV1PnetServerDeleteResponse, unknown, usecases_PNETServerDeleteInputDTO> = {}
) => {
    return useMutation<PostV1PnetServerDeleteResponse, unknown, usecases_PNETServerDeleteInputDTO>({
        // @ts-ignore
        mutationFn: ({id}: usecases_PNETServerDeleteInputDTO) => {
            if (!id) return null;
            return postV1LtiFormDelete({
                form: {
                    id: id
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
