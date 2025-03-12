import React, { ChangeEvent, useState } from 'react';
import useDebounce from '@/components/sidebar/search-debounce-input';
import { InputProps } from '@heroui/input/dist/input';
import { Autocomplete, AutocompleteItem } from '@heroui/react';

interface CallbackProps {
  isLoading: boolean;
  items: { key: string | number; value: string }[];
}

interface Props extends InputProps {
  fetchData: (search: string, props: InputProps) => CallbackProps;
}

export const AutoCompleteFull = (props: Props) => {
  const [inputChangeValue, onInputChange] = useState('');
  const debounceChangeValue = useDebounce(inputChangeValue, 300);
  const responseSearch = props.fetchData(debounceChangeValue, props);

  return (
    <Autocomplete
      variant={props.variant}
      className={props.className}
      items={responseSearch?.items}
      selectedKey={props.value}
      label={props.label}
      labelPlacement='inside'
      description={props.description}
      placeholder={props.placeholder}
      onInputChange={onInputChange}
      onSelectionChange={(key) => {
        console.log('Key');
        const event = {
          target: {
            value: key
          }
        } as ChangeEvent<HTMLInputElement>;
        key && props.onChange?.(event);
        console.log(key);
      }}
      isLoading={responseSearch.isLoading}
      aria-label={props.label?.toString()}
    >
      {(item) => <AutocompleteItem key={item.key}>{item.value}</AutocompleteItem>}
    </Autocomplete>
  );
};
