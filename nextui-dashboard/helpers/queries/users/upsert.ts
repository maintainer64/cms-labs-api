import {useMutation, useQuery} from "@tanstack/react-query";
import {models_User, postV1UserGet, postV1UserList, postV1UserUpsert, PostV1UserUpsertResponse} from "@/helpers/api";
import {TFormikData, TMutationCustomOptions} from "@/helpers/queries/types";
import queryClient from "../base";

export const useUserUpsert = (
    options: TMutationCustomOptions<PostV1UserUpsertResponse, unknown, TFormikData<models_User>> = {}
) => {
    return useMutation<PostV1UserUpsertResponse, unknown, TFormikData<models_User>>({
        // @ts-ignore
        mutationFn: ({values}: TFormikData<models_User>) => {
            if (values === null) return null;
            return postV1UserUpsert({
                form: {
                    email: values.email ?? '',
                    group_name: values.group_name,
                    id: values.id,
                    is_active: !values.deleted_at,
                    lti_user_id: values.lti_user_id,
                    name: values.name ?? '',
                    user_role: values.user_role,
                }
            })
        },
        ...options,
        async onSuccess(...args) {
            await queryClient.invalidateQueries({queryKey: ['postV1UserList']});
            await queryClient.invalidateQueries({queryKey: ['postV1UserGet']});
            if (options.onSuccess) {
                options.onSuccess(...args);
            }
        },
    });
}