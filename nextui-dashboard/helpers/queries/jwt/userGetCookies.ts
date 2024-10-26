import {useQuery} from "@tanstack/react-query";
import {postV1TokenCheck} from "@/helpers/api";

export const userGetCookies = () => {
    return useQuery({
        queryKey: ["userGetCookies"],
        queryFn: () => {
            return postV1TokenCheck({form: {}});
        },
        retry: 0,
    });
}
