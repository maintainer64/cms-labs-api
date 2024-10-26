import {useMutation} from "@tanstack/react-query";
import {
    postV1TokenPasswordChange,
    PostV1TokenPasswordChangeResponse,
    usecases_UserPasswordChangeInputDTO
} from "@/helpers/api";
import {TFormikData, TMutationCustomOptions} from "@/helpers/queries/types";

export const useUserPasswordChange = (
    options: TMutationCustomOptions<PostV1TokenPasswordChangeResponse, unknown, TFormikData<usecases_UserPasswordChangeInputDTO>> = {}
) => {
    return useMutation<PostV1TokenPasswordChangeResponse, unknown, TFormikData<usecases_UserPasswordChangeInputDTO>>({
        // @ts-ignore
        mutationFn: ({values}: TFormikData<usecases_UserPasswordChangeInputDTO>) => {
            if (values === null) return null;
            return postV1TokenPasswordChange({
                form: {
                    old_password: values.old_password ?? '',
                    new_password: values.new_password ?? '',
                    again_password: values.again_password ?? '',
                }
            })
        },
        ...options,
    });
}
