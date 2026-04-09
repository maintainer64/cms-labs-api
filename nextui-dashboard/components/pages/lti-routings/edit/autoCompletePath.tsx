import { InputProps } from '@heroui/input/dist/input';
import { AutoCompleteFull } from '@/components/base-forms/autocomplete';
import { Input } from '@heroui/react';
import { useQueryTaskAutoComplete } from '@/helpers/queries/task/use-query-task-auto-complete';

interface Props extends InputProps {
  labsTypeUnl: string;
}

export const LabsPathInput = (props: Props) => {
  if (props.labsTypeUnl === 'clabgate') {
    return <AutoCompleteFull {...props} fetchData={useQueryTaskAutoComplete.bind(useQueryTaskAutoComplete)} />;
  }
  return <Input {...props} />;
};

export const TestsPathInput = (props: Props) => {
  return <Input {...props} />;
};
