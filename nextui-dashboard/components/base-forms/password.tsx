import {Input} from "@nextui-org/react";
import React, {useId, useState} from "react";
import {UseInputProps} from "@nextui-org/input/dist/use-input";


interface PasswordInputProps extends Omit<UseInputProps, "isMultiline"> {
}


const EyeIconCrossed = () => {
    return (
        <svg className="shrink-0 size-3.5" width="24" height="24" viewBox="0 0 24 24" fill="none"
             stroke="currentColor"
             strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
            <path d="M9.88 9.88a3 3 0 1 0 4.24 4.24"></path>
            <path
                d="M10.73 5.08A10.43 10.43 0 0 1 12 5c7 0 10 7 10 7a13.16 13.16 0 0 1-1.67 2.68"></path>
            <path
                d="M6.61 6.61A13.526 13.526 0 0 0 2 12s3 7 10 7a9.74 9.74 0 0 0 5.39-1.61"></path>
            <line x1="2" x2="22" y1="2" y2="22"></line>
        </svg>
    )
}

const EyeIcon = () => {
    return (
        <svg className="shrink-0 size-3.5" width="24" height="24" viewBox="0 0 24 24" fill="none"
             stroke="currentColor"
             strokeWidth="2" strokeLinecap="round" strokeLinejoin="round">
            <path d="M2 12s3-7 10-7 10 7 10 7-3 7-10 7-10-7-10-7Z"></path>
            <circle cx="12" cy="12" r="3"></circle>
        </svg>
    )
}

export const PasswordInput = (props: PasswordInputProps) => {
    const [masked, setMasked] = useState<boolean>(true);
    return (
        <div className="relative">
            <Input
                {...props}
                type={masked ? 'password' : 'text'}
            />
            <button
                type="button"
                className="absolute inset-y-0 end-0 flex items-center z-20 px-3 cursor-pointer text-gray-400 rounded-e-md focus:outline-none focus:text-blue-600 dark:text-neutral-600 dark:focus:text-blue-500"
                onClick={setMasked.bind(setMasked, !masked)}
            >
                {masked ? <EyeIconCrossed/> : <EyeIcon/>}
            </button>
        </div>
    )
}