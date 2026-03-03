import React from 'react';
import { AutoCompleteFull, AutoCompleteFullProps } from '@/components/base-forms/autocomplete';
import { useQueryServerAutoComplete } from '@/helpers/queries/server/use-query-server-auto-complete';
import { useQueryUserAutoComplete } from '@/helpers/queries/user/use-query-user-auto-complete';

export const UserEmailInput = (props: AutoCompleteFullProps) => {
  return <AutoCompleteFull {...props} fetchData={useQueryUserAutoComplete} />;
};
