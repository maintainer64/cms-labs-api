import React from "react";
import {AuthLayoutWrapper} from "@/components/auth/authLayout";
import "@/styles/globals.css";
import {Login} from "@/components/auth/login";
import AuthError from "@/components/auth/error";

export default function LoginPage() {
    return <AuthLayoutWrapper>
        <Login/>
    </AuthLayoutWrapper>;
}

export function LoginError() {
    return (
        <AuthLayoutWrapper>
            <AuthError/>
        </AuthLayoutWrapper>
    );
}
