'use client';
import React from 'react';
import { addToast, Button, Input } from '@heroui/react';
import { Formik } from 'formik';
import { auth_UserPasswordChangeInputDTO } from '@/helpers/api';
import useLanguageBrowser from '@/helpers/locale';
import { useUserPasswordChange } from '@/helpers/queries/users/passwordChange';

const defaultValues: auth_UserPasswordChangeInputDTO = {
  old_password: '',
  new_password: '',
  again_password: ''
};

export const ProfilePasswordChangeForm = () => {
  const {
    locale: { UserFormPasswordChange, Forms, Sidebar }
  } = useLanguageBrowser();
  const { mutate } = useUserPasswordChange({
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
        description: error.body.msg,
        color: 'danger'
      });
    }
  });
  return (
    <Formik
      initialValues={defaultValues}
      validationSchema={undefined}
      onSubmit={(values, formikHelpers) => {
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
              value={values.old_password ?? ''}
              onChange={handleChange('old_password')}
            />
            <Input
              autoComplete='off'
              variant='bordered'
              label={UserFormPasswordChange.FieldNewPassword}
              type='password'
              value={values.new_password ?? ''}
              onChange={handleChange('new_password')}
            />
            <Input
              autoComplete='off'
              variant='bordered'
              label={UserFormPasswordChange.FieldAgainPassword}
              type='password'
              value={values.again_password ?? ''}
              onChange={handleChange('again_password')}
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
