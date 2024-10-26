import {useMutation} from "@tanstack/react-query";
import {TMutationCustomOptions} from "../types";
import {
    changeReleaseHandlerApiV1ReleasesPost,
    type ChangeReleaseHandlerApiV1ReleasesPostResponse,
    ReleasesSchemaItem
} from "@colday/api";
import queryClient from "../base";

export const editRelease = (
    options: TMutationCustomOptions<ChangeReleaseHandlerApiV1ReleasesPostResponse, unknown, ReleasesSchemaItem> = {}
) => {
    return useMutation<ChangeReleaseHandlerApiV1ReleasesPostResponse, unknown, ReleasesSchemaItem>({
        mutationFn: (values) => {
            return changeReleaseHandlerApiV1ReleasesPost({requestBody: values})
        },
        ...options,
        async onSuccess(...args) {
            await queryClient.invalidateQueries({queryKey: ['releasesList']});

            if (options.onSuccess) {
                options.onSuccess(...args);
            }
        },
    });
}
