'use client';

import { LoginSchema } from '@/helpers/schemas';
import { LoginFormType } from '@/helpers/types';
import { Button, Input } from '@heroui/react';
import { Formik } from 'formik';
import { useCallback } from 'react';
import useLanguageBrowser from '@/helpers/locale';
import { postV1TokenLogin } from '@/helpers/api';
import { FormikHelpers } from 'formik/dist/types';
import { RoutesLocation } from '@/components/routes';
import { useLTIFormsSSOList } from '@/helpers/queries/lti-forms/sso';
import { SSOAuthorizationGet } from '@/components/pages/auth/ssoSave';
import { SecurityIcon } from '@/components/icons/sso';

export const Login = () => {
  const { locale } = useLanguageBrowser();

  const initialValues: LoginFormType = {
    email: '',
    password: ''
  };
  const ssoLinks = useLTIFormsSSOList();

  const ssoButtons = ssoLinks.data?.result?.model.map((service, key) => (
    <>
      <div className='inline-flex items-center justify-center w-full'>
        <hr className='w-32 h-px my-4 border-0 dark:bg-gray-100/10 bg-gray-700/10' />
        <span className='px-3 text-sm font-medium text-gray-900 dark:text-white bg-transparent'>{locale.SSO.OR}</span>
        <hr className='w-32 h-px my-4 border-0 dark:bg-gray-100/10 bg-gray-700/10' />
      </div>
      <Button
        startContent={<SecurityIcon />}
        className='w-full'
        key={key}
        onPress={() => {
          window.location.href = service.sso_url || '';
        }}
        variant='flat'
        color='default'
      >
        {service.name}
      </Button>
    </>
  ));

  // Обработка внутренней авторизации
  const handleLogin = useCallback(async (values: LoginFormType, formikHelpers: FormikHelpers<LoginFormType>) => {
    try {
      await postV1TokenLogin({ form: { email: values.email, password: values.password } });
      const params = SSOAuthorizationGet();
      if (params === null) {
        // Default redirect
        window.location.href = RoutesLocation.home();
      }
    } catch (error: any) {
      formikHelpers.setErrors({});
      formikHelpers.setErrors({ password: error.body.msg });
    }
  }, []);

  return (
    <>
      <div className='text-center text-[25px] font-bold mb-6'>{locale.Login.PageName}</div>

      <Formik initialValues={initialValues} validationSchema={LoginSchema()} onSubmit={handleLogin}>
        {({ values, errors, touched, handleChange, handleSubmit }) => (
          <>
            <div className='flex flex-col w-1/2 gap-4 mb-4'>
              <Input
                className='w-full'
                variant='bordered'
                label={locale.Login.FieldEmail}
                type='email'
                value={values.email}
                isInvalid={!!errors.email && !!touched.email}
                errorMessage={errors.email}
                onChange={handleChange('email')}
              />
              <Input
                className='w-full'
                variant='bordered'
                label={locale.Login.FieldPassword}
                type='password'
                value={values.password}
                isInvalid={!!errors.password && !!touched.password}
                errorMessage={errors.password}
                onChange={handleChange('password')}
              />
              <Button className='w-full' onPress={() => handleSubmit()} variant='flat' color='default'>
                {locale.Login.Submit}
              </Button>
              {ssoButtons}
            </div>
          </>
        )}
      </Formik>
    </>
  );
};
