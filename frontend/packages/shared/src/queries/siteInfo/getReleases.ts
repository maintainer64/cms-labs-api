import {useQuery} from "@tanstack/react-query";
import {releasesHandlerApiV1ReleasesGet} from "@colday/api";

export const useGetReleasesList = (all: boolean) => {
    const data = all ? {releases: 'all'} : undefined;
    return useQuery({
        queryKey: ["releasesList", all],
        queryFn: () => {
            return releasesHandlerApiV1ReleasesGet(data);
        },
    });
};
