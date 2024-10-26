import React from "react";

export function Loading() {
    return (<div
        className="animate-spin rounded-full h-24 w-24 border-t-4 border-black border-opacity-50">

    </div>);
}

export default function AuthLoadingWrapper() {
    return (
        <div
            className="flex items-center justify-center h-screen">
            <Loading/>
        </div>
    );
}