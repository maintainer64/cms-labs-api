import { UseMutationOptions } from '@tanstack/react-query';
import { FormikHelpers } from 'formik';

export type TMutationCustomOptions<Response, Params> = Omit<
  UseMutationOptions<Response, unknown, Params>,
  'mutationFn'
>;

export type TFormikData<TData = unknown> = {
  values: TData;
  formikHelpers: FormikHelpers<TData>;
};
