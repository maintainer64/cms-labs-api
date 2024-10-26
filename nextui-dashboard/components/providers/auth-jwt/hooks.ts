import {useContext} from "react";
import {UserProfileContext} from "./context";
import {auth_TokenPublicData} from "@/helpers/api";

export const useUserProfile = (): auth_TokenPublicData => {
    const context = useContext(UserProfileContext);
    if (!context || !context.profile) {
        console.error("useUserProfile must be used within a UserProfileProvider");
        return {};
    }
    return context.profile;
}
