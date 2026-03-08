'use client';
import React from 'react';
import { addToast, Button, Input } from '@heroui/react';
import { Formik } from 'formik';
import useLanguageBrowser from '@/helpers/locale';
import { useMutationUserPasswordChange } from '@/helpers/queries/user/use-mutation-user-password-change';
import { CamelCasedPropertiesDeep } from 'type-fest';
import { AuthUserPasswordChangeRequest } from '@/helpers/api';

const defaultValues: CamelCasedPropertiesDeep<AuthUserPasswordChangeRequest['params']> = {
  oldPassword: '',
  newPassword: '',
  againPassword: ''
};

export const ProfilePasswordChangeForm = () => {
  const {
    locale: { UserFormPasswordChange, Forms, Sidebar }
  } = useLanguageBrowser();
  const { mutate } = useMutationUserPasswordChange({
    onSuccess: (data, { formikHelpers }) => {
      formikHelpers.resetForm();
      addToast({
        title: Forms.SaveSuccess,
        description: '',
        color: 'success'
      });
    },
    onError: (error: any) => {
      addToast({
        title: Forms.SaveError,
        description: error.data.message,
        color: 'danger'
      });
    }
  });
  return (
    <Formik
      initialValues={defaultValues}
      validationSchema={undefined}
      onSubmit={(values, formikHelpers) => {
        // @ts-ignore
        mutate({ values, formikHelpers });
      }}
    >
      {({ values, handleChange, handleSubmit }) => (
        <>
          <div className='flex flex-col w-1/2 gap-4 mb-4'>
            <Input
              autoComplete='off'
              variant='bordered'
              label={UserFormPasswordChange.FieldOldPassword}
              type='password'
              value={values.oldPassword ?? ''}
              onChange={handleChange('oldPassword')}
            />
            <Input
              autoComplete='off'
              variant='bordered'
              label={UserFormPasswordChange.FieldNewPassword}
              type='password'
              value={values.newPassword ?? ''}
              onChange={handleChange('newPassword')}
            />
            <Input
              autoComplete='off'
              variant='bordered'
              label={UserFormPasswordChange.FieldAgainPassword}
              type='password'
              value={values.againPassword ?? ''}
              onChange={handleChange('againPassword')}
            />
            <Button onPress={() => handleSubmit()} variant='flat' color='primary'>
              {Sidebar.Save}
            </Button>
          </div>
        </>
      )}
    </Formik>
  );
};
