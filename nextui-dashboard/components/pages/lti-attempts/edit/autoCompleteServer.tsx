import React from 'react';
import { InputProps } from '@heroui/input/dist/input';
import { AutoCompleteFull } from '@/components/base-forms/autocomplete';
import { usePnetServerAutocompleteData } from '@/helpers/queries/pnet-server/get';

export const ServerInput = (props: InputProps) => {
  return <AutoCompleteFull {...props} fetchData={usePnetServerAutocompleteData} />;
};
