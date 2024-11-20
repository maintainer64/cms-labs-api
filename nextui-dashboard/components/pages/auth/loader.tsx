import React from "react";
import {Loading} from "@/components/scroll/loader";

export default function AuthLoadingWrapper() {
    return (
        <div
            className="flex items-center justify-center h-screen">
            <Loading/>
        </div>
    );
}
