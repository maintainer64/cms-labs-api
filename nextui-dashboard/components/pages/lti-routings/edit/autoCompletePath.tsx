import { curlRequestsData } from '@/helpers/queries/curl-requests/get';
import { InputProps } from '@heroui/input/dist/input';
import { AutoCompleteFull } from '@/components/base-forms/autocomplete';
import { Input } from '@heroui/react';
import { clabgateTaskListData } from '@/helpers/queries/tasks/get';

interface Props extends InputProps {
  labsTypeUnl: string;
}

export const LabsPathInput = (props: Props) => {
  if (props.labsTypeUnl === 'curl') {
    return <AutoCompleteFull {...props} fetchData={curlRequestsData.bind(curlRequestsData)} />;
  }
  if (props.labsTypeUnl === 'clabgate') {
    return <AutoCompleteFull {...props} fetchData={clabgateTaskListData.bind(clabgateTaskListData)} />;
  }
  return <Input {...props} />;
};

export const TestsPathInput = (props: Props) => {
  return <Input {...props} />;
};
