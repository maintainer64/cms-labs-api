'use client';
import React from 'react';
import { Button, Input } from '@heroui/react';
import { Formik } from 'formik';
import { usecases_UserPasswordChangeInputDTO } from '@/helpers/api';
import useLanguageBrowser from '@/helpers/locale';
import { useAlert } from '@/components/alerts/hooks';
import { useUserPasswordChange } from '@/helpers/queries/users/passwordChange';

const defaultValues: usecases_UserPasswordChangeInputDTO = {
  old_password: '',
  new_password: '',
  again_password: ''
};

export const ProfilePasswordChangeForm = () => {
  const {
    locale: { UserFormPasswordChange, Forms, Sidebar }
  } = useLanguageBrowser();
  const { showAlert } = useAlert();
  const { mutate } = useUserPasswordChange({
    onSuccess: (data, { formikHelpers }) => {
      formikHelpers.resetForm();
      showAlert({
        type: 'success',
        message: Forms.SaveSuccess
      });
    },
    onError: (error: any) => {
      showAlert({
        type: 'danger',
        message: Forms.SaveError,
        description: error.body.msg
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
