import { UseMutationOptions } from '@tanstack/react-query';
import { FormikHelpers } from 'formik';

export type TMutationCustomOptions<TData = unknown, TError = unknown, TVariables = void, TContext = unknown> = Omit<
  UseMutationOptions<TData, TError, TVariables, TContext>,
  'mutationFn'
>;

export type TFormikData<TData = unknown> = {
  values: TData;
  formikHelpers: FormikHelpers<TData>;
};
