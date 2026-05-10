import { InputProps } from '@heroui/input/dist/input';
import { Input } from '@heroui/react';

interface Props extends InputProps {
  labsTypeUnl: string;
}

export const LabsPathInput = (props: Props) => {
  return <Input {...props} />;
};

export const TestsPathInput = (props: Props) => {
  return <Input {...props} />;
};
