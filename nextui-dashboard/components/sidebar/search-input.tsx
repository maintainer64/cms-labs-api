import React, {useEffect, useState} from 'react';
import useDebounce from "@/components/sidebar/search-debounce-input";
import {Input} from "@nextui-org/react";

interface SearchInputProps {
    defaultValue?: string;
    setValue?: (value: string) => void;
    placeholder?: string;
}

const SearchInput = ({defaultValue = '', setValue, placeholder}: SearchInputProps) => {
    const [inputValue, setInputValue] = useState(defaultValue);
    const debouncedValue = useDebounce(inputValue, 300);

    useEffect(() => {
        setValue?.(debouncedValue);
    }, [debouncedValue, setValue]);

    return (
        <Input
            type="text"
            value={inputValue}
            onChange={(e) => setInputValue(e.target.value)}
            placeholder={placeholder}
            classNames={{
                input: "w-full",
                mainWrapper: "w-full",
            }}
        />
    );
};

export default SearchInput;