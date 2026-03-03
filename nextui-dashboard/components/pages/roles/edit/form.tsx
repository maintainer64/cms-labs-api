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
import { ModelsRole } from '@/helpers/api';
import { CamelCasedPropertiesDeep } from 'type-fest';
import { useQueryRoleList } from '@/helpers/queries/role/use-query-role-list';
import { useMutationRoleUpsert } from '@/helpers/queries/role/use-mutation-role-upsert';
import { useMutationRoleDelete } from '@/helpers/queries/role/use-mutation-role-delete';

interface EditFormProps {
  id?: number;
}

const defaultValues: CamelCasedPropertiesDeep<ModelsRole> = {
  code: '',
  createdAt: '',
  name: '',
  updatedAt: ''
};

export const RolesEditForm = ({ id }: EditFormProps) => {
  const {
    locale: { Roles, Forms, Sidebar }
  } = useLanguageBrowser();
  const navigate = useNavigate();
  const queryRoles = useQueryRoleList({});
  const initialValues = queryRoles.data?.model.filter((role) => role.id === id)?.[0] ?? defaultValues;
  const { mutate } = useMutationRoleUpsert({
    onSuccess: (data) => {
      navigate(RoutesLocation.rolesEdit(data?.id?.toString() || ''), { replace: true });
      addToast({
        title: Forms.SaveSuccess,
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
  const onDeleteMutation = useMutationRoleDelete({
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
        description: error.data.message,
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
        mutate({
          id: values.id,
          name: values.name,
          code: values.code
        });
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
              value={dayjs(initialValues.createdAt ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={Roles.FieldUpdatedAt}
              type='datetime-local'
              value={dayjs(initialValues.updatedAt ?? '').format('YYYY-MM-DDTHH:mm')}
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
