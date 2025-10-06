import React from 'react';
import { AutoCompleteFull, AutoCompleteFullProps } from '@/components/base-forms/autocomplete';
import { useQueryServerAutoComplete } from '@/helpers/queries/server/use-query-server-auto-complete';

export const ServerInput = (props: AutoCompleteFullProps) => {
  return <AutoCompleteFull {...props} fetchData={useQueryServerAutoComplete} />;
};
