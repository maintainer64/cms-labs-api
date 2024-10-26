import {Route, Routes} from "react-router-dom";
import LoginPage from "@/app/(auth)/layout";
import LanguagePage from "@/app/(app)/lang/page";
import * as React from "react";
import HomePage from "@/app/(app)/home/page";
import {AccountsPageEdit, AccountsPage, ProfilePagePasswordChange} from "@/app/(app)/accounts/page";
import {RoutesLocation} from "@/components/routes";

const RoutesAdmin = () => {
    return <Routes>
        <Route path={RoutesLocation.accounts()} element={<AccountsPage/>}/>
        <Route path={RoutesLocation.accountsEdit()} element={<AccountsPageEdit/>}/>
        <Route path={RoutesLocation.profileChangePassword()} element={<ProfilePagePasswordChange/>}/>
        <Route path={RoutesLocation.accountsCreate()} element={<AccountsPageEdit/>}/>
        <Route path={RoutesLocation.login()} element={<LoginPage/>}/>
        <Route path={RoutesLocation.language()} element={<LanguagePage/>}/>
        <Route path={RoutesLocation.home()} element={<HomePage/>}/>
    </Routes>
}

export default RoutesAdmin;