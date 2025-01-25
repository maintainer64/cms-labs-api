'use client';
import React from 'react';
import { Button, Checkbox, Input, Select, SelectItem } from '@nextui-org/react';
import { Formik } from 'formik';
import { models_User } from '@/helpers/api';
import useLanguageBrowser from '@/helpers/locale';
import { useUserUpsert } from '@/helpers/queries/users/upsert';
import { useNavigate } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import dayjs from 'dayjs';
import { useAlert } from '@/components/alerts/hooks';
import { useUserByID } from '@/helpers/queries/users/get';
import { Loading } from '@/components/scroll/loader';

interface EditFormProps {
  id?: number;
}

const defaultValues: models_User = {
  created_at: '',
  deleted_at: '',
  email: '',
  group_name: '',
  id: undefined,
  last_launch_id: '',
  lti_user_id: '',
  name: '',
  updated_at: '',
  user_role: 'student'
};

export const UserRoles = () => {
  const {
    locale: { UserForm }
  } = useLanguageBrowser();
  return [
    { key: 'student', label: UserForm.FieldUserRoleStudent },
    { key: 'instructor', label: UserForm.FieldUserRoleInstructor },
    { key: 'admin', label: UserForm.FieldUserRoleAdmin }
  ];
};

export const AccountsEditForm = ({ id }: EditFormProps) => {
  const {
    locale: { UserForm, Forms, Sidebar }
  } = useLanguageBrowser();
  const { showAlert } = useAlert();
  const navigate = useNavigate();
  const response = useUserByID(id);
  const roles = UserRoles();
  const initialValues = response.data?.result?.model ?? defaultValues;
  const { mutate } = useUserUpsert({
    onSuccess: (data, { formikHelpers }) => {
      navigate(RoutesLocation.accountsEdit(data.result?.id?.toString() || ''), { replace: true });
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
  if (response.isLoading) return <Loading />;
  return (
    <Formik
      initialValues={initialValues}
      validationSchema={undefined}
      onSubmit={(values, formikHelpers) => {
        mutate({ values, formikHelpers });
      }}
    >
      {({ values, handleChange, setFieldValue, handleSubmit }) => (
        <>
          <div className='flex flex-col gap-4 mb-4'>
            <Input
              variant='bordered'
              label={UserForm.FieldID}
              type='number'
              value={(initialValues.id ?? 0).toString()}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={UserForm.FieldName}
              type='text'
              value={values.name ?? ''}
              onChange={handleChange('name')}
            />
            <Input
              variant='bordered'
              label={UserForm.FieldEmail}
              type='email'
              value={values.email ?? ''}
              onChange={handleChange('email')}
            />
            <Select
              variant='bordered'
              label={UserForm.FieldUserRole}
              selectedKeys={[values.user_role ?? '']}
              onSelectionChange={(keys) => setFieldValue('user_role', keys.currentKey || 'student')}
            >
              {roles.map((role) => (
                <SelectItem key={role.key}>{role.label}</SelectItem>
              ))}
            </Select>
            <Input
              variant='bordered'
              label={UserForm.FieldGroupName}
              type='text'
              value={values.group_name ?? ''}
              onChange={handleChange('group_name')}
            />
            <Input
              variant='bordered'
              label={UserForm.FieldExternalLTIID}
              type='text'
              value={values.lti_user_id ?? ''}
              onChange={handleChange('lti_user_id')}
            />
            <Checkbox type='checkbox' defaultSelected={!values.deleted_at} onChange={handleChange('deleted_at')}>
              {UserForm.FieldIsActive}
            </Checkbox>
            <Input
              variant='bordered'
              label={UserForm.FieldCreatedAt}
              type='datetime-local'
              value={dayjs(initialValues.created_at ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={UserForm.FieldUpdatedAt}
              type='datetime-local'
              value={dayjs(initialValues.updated_at ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
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
