import { Autocomplete, AutocompleteItem, Input } from '@nextui-org/react';
import { useUNLFileData } from '@/helpers/queries/unl-files/get';
import React, { useState } from 'react';
import useDebounce from '@/components/sidebar/search-debounce-input';
import { InputProps } from '@nextui-org/input/dist/input';

interface AutoCompletePathProps extends InputProps {
  unlFileSearchTypes?: string[];
}

export const AutoCompletePath = (props: AutoCompletePathProps) => {
  const [inputChangeValue, onInputChange] = useState('');
  const debounceChangeValue = useDebounce(inputChangeValue, 300);
  const responseSearch = useUNLFileData(debounceChangeValue, props.unlFileSearchTypes, Number(props.value));

  return (
    <Autocomplete
      variant={props.variant}
      className={props.className}
      items={responseSearch.files}
      selectedKey={props.value}
      label={props.label}
      placeholder={props.placeholder}
      onInputChange={onInputChange}
      onSelectionChange={(key) => {
        // @ts-ignore
        key && props.onChange(key);
      }}
      isLoading={responseSearch.isLoading}
    >
      {(file) => <AutocompleteItem key={file.id}>{file.path}</AutocompleteItem>}
    </Autocomplete>
  );
};

interface Props extends InputProps {
  labsTypeUnl: string;
}

export const LabsPathInput = (props: Props) => {
  if (props.labsTypeUnl === 'file') {
    return <AutoCompletePath unlFileSearchTypes={['unl']} {...props} />;
  }
  return <Input {...props} />;
};

export const TestsPathInput = (props: Props) => {
  if (props.labsTypeUnl === 'file') {
    return <AutoCompletePath unlFileSearchTypes={['yml', 'yaml']} {...props} />;
  }
  return <Input {...props} />;
};
