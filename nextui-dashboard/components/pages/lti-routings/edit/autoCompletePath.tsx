import { curlRequestsData } from '@/helpers/queries/curl-requests/get';
import { InputProps } from '@heroui/input/dist/input';
import { AutoCompleteFull } from '@/components/base-forms/autocomplete';
import { Input } from '@heroui/react';

interface Props extends InputProps {
  labsTypeUnl: string;
}

export const LabsPathInput = (props: Props) => {
  if (props.labsTypeUnl !== 'curl') {
    return <Input {...props} />;
  }
  return <AutoCompleteFull {...props} fetchData={curlRequestsData.bind(curlRequestsData)} />;
};

export const TestsPathInput = (props: Props) => {
  return <Input {...props} />;
};
