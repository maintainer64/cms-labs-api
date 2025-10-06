'use client';
import React from 'react';
import { addToast, Button, Checkbox, Input } from '@heroui/react';
import { Formik } from 'formik';
import useLanguageBrowser from '@/helpers/locale';
import { Link, useNavigate } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import dayjs from 'dayjs';
import { Loading } from '@/components/scroll/loader';
import { RolesSelector } from '@/components/base-forms/roles';
import { MapUserItem, UserItem } from '@/helpers/queries/user/use-infinity-user-list';
import { useQueryUserGet } from '@/helpers/queries/user/use-query-user-get';
import { useMutationUserUpsert } from '@/helpers/queries/user/use-mutation-user-upsert';

interface EditFormProps {
  id?: number;
}

const defaultValues: UserItem = {
  createdAt: '',
  deletedAt: '',
  email: '',
  groupName: '',
  id: undefined,
  lastLaunchId: '',
  ltiUserId: '',
  name: '',
  updatedAt: '',
  roles: []
};

export const AccountsEditForm = ({ id }: EditFormProps) => {
  const {
    locale: { UserForm, Forms, Sidebar }
  } = useLanguageBrowser();
  const navigate = useNavigate();
  const queryUser = useQueryUserGet({ id });
  const initialValues = MapUserItem(queryUser.data?.model, queryUser.data?.roles) ?? defaultValues;

  const { mutate } = useMutationUserUpsert({
    onSuccess: (data) => {
      navigate(RoutesLocation.accountsEdit(data?.id?.toString() || ''), { replace: true });
      addToast({
        title: Forms.SaveSuccess,
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
  if (queryUser.isLoading) return <Loading size='md' />;
  return (
    <Formik
      initialValues={initialValues}
      validationSchema={undefined}
      onSubmit={(values, formikHelpers) => {
        if (values === null) return;
        mutate({
          email: values.email ?? '',
          groupName: values.groupName,
          id: values.id,
          isActive: !values.deletedAt,
          ltiUserId: values.ltiUserId,
          name: values.name ?? '',
          roles: values.roles?.map((roleId) => parseInt(roleId.toString())),
          store: values.store || {}
        });
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
            <RolesSelector
              label={UserForm.FieldUserRole}
              selectedKeys={values.roles ?? []}
              onSelectionChange={(keys) => setFieldValue('roles', Array.from(keys))}
            />
            <Input
              variant='bordered'
              label={UserForm.FieldGroupName}
              type='text'
              value={values.groupName ?? ''}
              onChange={handleChange('groupName')}
            />
            <Input
              variant='bordered'
              label={UserForm.FieldExternalLTIID}
              type='text'
              value={values.ltiUserId ?? ''}
              onChange={handleChange('ltiUserId')}
            />
            <Checkbox type='checkbox' defaultSelected={Boolean(values.deletedAt)} onChange={handleChange('deletedAt')}>
              {UserForm.FieldIsDeactivated}
            </Checkbox>
            <Input
              variant='bordered'
              label={UserForm.FieldCreatedAt}
              type='datetime-local'
              value={dayjs(initialValues.createdAt ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={UserForm.FieldUpdatedAt}
              type='datetime-local'
              value={dayjs(initialValues.updatedAt ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Button onPress={() => handleSubmit()} variant='flat' color='primary'>
              {Sidebar.Save}
            </Button>
            <Button
              as={Link}
              variant='flat'
              color='secondary'
              to={RoutesLocation.ltiAttemptUser(values.id?.toString())}
            >
              {UserForm.FieldRelationAttempts}
            </Button>
          </div>
        </>
      )}
    </Formik>
  );
};
