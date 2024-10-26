import {Route, Routes} from "react-router-dom";
import LanguagePage from "@/app/(app)/lang/page";
import * as React from "react";
import LoginPage, {LoginError} from "@/app/(auth)/layout";
import {RoutesLocation} from "@/components/routes";

const RoutesStudent = () => {
    return <Routes>
        <Route path={RoutesLocation.language()} element={<LanguagePage/>}/>
        <Route path={RoutesLocation.login()} element={<LoginPage/>}/>
        <Route path={RoutesLocation.tasks()} element={<LoginError/>}/>
        <Route path={RoutesLocation.home()} element={<LoginError/>}/>
    </Routes>
}

export default RoutesStudent;