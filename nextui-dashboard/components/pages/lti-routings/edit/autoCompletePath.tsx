import { useUNLFileData } from '@/helpers/queries/unl-files/get';
import { InputProps } from '@heroui/input/dist/input';
import { AutoCompleteFull } from '@/components/base-forms/autocomplete';
import { Input } from '@heroui/react';

interface Props extends InputProps {
  labsTypeUnl: string;
}

export const LabsPathInput = (props: Props) => {
  if (props.labsTypeUnl !== 'file') {
    return <Input {...props} />;
  }
  return <AutoCompleteFull {...props} fetchData={useUNLFileData.bind(useUNLFileData, ['unl'])} />;
};

export const TestsPathInput = (props: Props) => {
  if (props.labsTypeUnl !== 'file') {
    return <Input {...props} />;
  }
  return <AutoCompleteFull {...props} fetchData={useUNLFileData.bind(useUNLFileData, ['yml', 'yaml'])} />;
};
