'use client';
import React from 'react';
import { addToast, Button, Input } from '@heroui/react';
import { Formik } from 'formik';
import useLanguageBrowser from '@/helpers/locale';
import { useNavigate } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import dayjs from 'dayjs';
import { Loading } from '@/components/scroll/loader';
import { useConfirmPopup } from '@/components/hooks/useDeletePopup';
import { models_Role } from '@/helpers/api';
import { useRoleDelete } from '@/helpers/queries/roles/delete';
import { useRoleUpsert } from '@/helpers/queries/roles/upsert';
import { useRolesList } from '@/helpers/queries/roles/get';

interface EditFormProps {
  id?: number;
}

const defaultValues: models_Role = {
  code: '',
  created_at: '',
  name: '',
  updated_at: ''
};

export const RolesEditForm = ({ id }: EditFormProps) => {
  const {
    locale: { Roles, Forms, Sidebar }
  } = useLanguageBrowser();
  const navigate = useNavigate();
  const queryRoles = useRolesList();
  const initialValues = queryRoles.data?.result?.model.filter((role) => role.id === id)?.[0] ?? defaultValues;
  const { mutate } = useRoleUpsert({
    onSuccess: (data, { formikHelpers }) => {
      navigate(RoutesLocation.rolesEdit(data.result?.id?.toString() || ''), { replace: true });
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
  const onDeleteMutation = useRoleDelete({
    onSuccess: () => {
      navigate(RoutesLocation.roles(), { replace: true });
      addToast({
        title: Forms.DeleteSuccess,
        color: 'success'
      });
    },
    onError: (error: any) => {
      addToast({
        title: Forms.DeleteError,
        description: error.body.msg,
        color: 'danger'
      });
    }
  });
  const roleDeletePopup = useConfirmPopup({
    title: Roles.DeletePopup.Title,
    description: Roles.DeletePopup.Description,
    onConfirm: onDeleteMutation.mutate.bind(onDeleteMutation.mutate, { id: id || 0 })
  });
  if (queryRoles.isLoading) return <Loading size='md' />;
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
          {roleDeletePopup.component({})}
          <div className='flex flex-col gap-4 mb-4'>
            <Input
              variant='bordered'
              label={Roles.FieldID}
              type='number'
              value={(initialValues.id ?? 0).toString()}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={Roles.FieldName}
              type='text'
              value={values.name ?? ''}
              onChange={handleChange('name')}
            />
            <Input
              variant='bordered'
              label={Roles.FieldCode}
              type='text'
              value={values.code ?? ''}
              onChange={handleChange('code')}
            />
            <Input
              variant='bordered'
              label={Roles.FieldCreatedAt}
              type='datetime-local'
              value={dayjs(initialValues.created_at ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={Roles.FieldUpdatedAt}
              type='datetime-local'
              value={dayjs(initialValues.updated_at ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Button onPress={() => handleSubmit()} variant='flat' color='primary'>
              {Sidebar.Save}
            </Button>
            <Button onPress={roleDeletePopup.onOpen} variant='flat' color='danger'>
              {Sidebar.Delete}
            </Button>
          </div>
        </>
      )}
    </Formik>
  );
};
