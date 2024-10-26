import {Route, Routes} from "react-router-dom";
import LoginPage, {LoginError} from "@/app/(auth)/layout";
import LanguagePage from "@/app/(app)/lang/page";
import * as React from "react";
import {RoutesLocation} from "@/components/routes";

const RoutesUnknown = () => {
    return <Routes>
        <Route path={RoutesLocation.login()} element={<LoginPage/>}>

        </Route>
        <Route path={RoutesLocation.language()} element={<LanguagePage/>}>

        </Route>
        <Route path={RoutesLocation.home()} element={<LoginError/>}>

        </Route>
    </Routes>
}

export default RoutesUnknown;