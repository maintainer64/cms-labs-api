'use client';

import { LoginSchema } from '@/helpers/schemas';
import { Button, Input } from '@heroui/react';
import { Formik } from 'formik';
import useLanguageBrowser from '@/helpers/locale';
import { RoutesLocation } from '@/components/routes';
import { SSOAuthorizationGet } from '@/components/pages/auth/ssoSave';
import { SecurityIcon } from '@/components/icons/sso';
import { useQueryAuthProviderSsoListGet } from '@/helpers/queries/lti_form/use-query-lti-form-sso-list-get';
import { useMutationUserLogin } from '@/helpers/queries/user/use-mutation-user-login';
import { CamelCasedPropertiesDeep } from 'type-fest';
import { AuthRenewManagerCredentialsInputDTO } from '@/helpers/api';

const defaultValues: CamelCasedPropertiesDeep<AuthRenewManagerCredentialsInputDTO> = {
  email: '',
  password: ''
};

export const Login = () => {
  const { locale } = useLanguageBrowser();
  const ssoLinks = useQueryAuthProviderSsoListGet({});

  const ssoButtons = ssoLinks.data?.model.map((service, key) => (
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
          window.location.href = service.ssoUrl || '';
        }}
        variant='flat'
        color='default'
      >
        {service.name}
      </Button>
    </>
  ));

  const { mutate } = useMutationUserLogin({
    onSuccess: (data, { formikHelpers }) => {
      formikHelpers.resetForm();
      const params = SSOAuthorizationGet();
      if (params === null) {
        // Default redirect
        window.location.href = RoutesLocation.home();
      }
    },
    onError: (error: any, { formikHelpers }) => {
      formikHelpers.setErrors({});
      formikHelpers.setErrors({ password: error.body.msg });
    }
  });

  return (
    <>
      <div className='text-center text-[25px] font-bold mb-6'>{locale.Login.PageName}</div>

      <Formik
        initialValues={defaultValues}
        validationSchema={LoginSchema()}
        onSubmit={(values, formikHelpers) => {
          // @ts-ignore
          mutate({ values, formikHelpers });
        }}
      >
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
