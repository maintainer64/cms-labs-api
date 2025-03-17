import React from 'react';
import { AutoCompleteFull, AutoCompleteFullProps } from '@/components/base-forms/autocomplete';
import { usePnetServerAutocompleteData } from '@/helpers/queries/pnet-server/get';

export const ServerInput = (props: AutoCompleteFullProps) => {
  return <AutoCompleteFull {...props} fetchData={usePnetServerAutocompleteData} />;
};
