import {createContext, ReactNode} from 'react';
import {userGetCookies} from "@/helpers/queries/jwt/userGetCookies";
import {auth_TokenPublicData} from "@/helpers/api";
import AuthLoadingWrapper from "@/components/pages/auth/loader";

type UserProfileProfile = {
    profile?: auth_TokenPublicData
};

type UserProfileProvider = {
    children: ReactNode;
};

export const UserProfileContext = createContext<UserProfileProfile>({});

export const UserProfileProvider: React.FC<UserProfileProvider> = ({children}) => {
    const {data, isLoading} = userGetCookies();
    if (isLoading) {
        return <AuthLoadingWrapper/>
    }
    return (
        <UserProfileContext.Provider value={{profile: data?.result}}>
            {children}
        </UserProfileContext.Provider>
    );
};
