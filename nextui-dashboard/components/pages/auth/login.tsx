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
import queryClient from '@/helpers/queries/base';
import { useLTIFormsSSOList } from '@/helpers/queries/lti-forms/sso';
import { SSOAuthorizationGet } from '@/components/pages/auth/ssoSave';

export const Login = () => {
  const { locale } = useLanguageBrowser();

  const initialValues: LoginFormType = {
    email: '',
    password: ''
  };
  const ssoLinks = useLTIFormsSSOList();

  const ssoButtons = ssoLinks.data?.result?.model.map((service, key) => (
    <Button
      className='mt-2'
      key={key}
      onPress={() => {
        window.location.href = service.sso_url || '';
      }}
      variant='flat'
      color='secondary'
    >
      {service.name}
    </Button>
  ));

  // Обработка внутренней авторизации
  const handleLogin = useCallback(async (values: LoginFormType, formikHelpers: FormikHelpers<LoginFormType>) => {
    try {
      await postV1TokenLogin({ form: { email: values.email, password: values.password } });
      await queryClient.invalidateQueries({ queryKey: ['userGetCookies'] });
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
                variant='bordered'
                label={locale.Login.FieldEmail}
                type='email'
                value={values.email}
                isInvalid={!!errors.email && !!touched.email}
                errorMessage={errors.email}
                onChange={handleChange('email')}
              />
              <Input
                variant='bordered'
                label={locale.Login.FieldPassword}
                type='password'
                value={values.password}
                isInvalid={!!errors.password && !!touched.password}
                errorMessage={errors.password}
                onChange={handleChange('password')}
              />
            </div>

            <Button onPress={() => handleSubmit()} variant='flat' color='primary'>
              {locale.Login.Submit}
            </Button>
          </>
        )}
      </Formik>
      {ssoButtons}
    </>
  );
};
