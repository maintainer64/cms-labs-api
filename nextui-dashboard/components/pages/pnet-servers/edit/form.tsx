'use client';
import React from 'react';
import { addToast, Button, Checkbox, Input, Select, SelectItem } from '@heroui/react';
import { Formik } from 'formik';
import useLanguageBrowser from '@/helpers/locale';
import { useNavigate } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import dayjs from 'dayjs';
import { Loading } from '@/components/scroll/loader';
import { useConfirmPopup } from '@/components/hooks/useDeletePopup';
import { PasswordInput } from '@/components/base-forms/password';
import { RolesSelector } from '@/components/base-forms/roles';
import { MapServerItem, PnetServerItem, useQueryServerGet } from '@/helpers/queries/server/use-query-server-get';
import { useMutationServerUpsert } from '@/helpers/queries/server/use-mutation-server-upsert';
import { useMutationServerDelete } from '@/helpers/queries/server/use-mutation-server-delete';

interface EditFormProps {
  id?: number;
}

const defaultValues: PnetServerItem = {
  clientId: '',
  type: 'pnet',
  createdAt: '',
  isActive: true,
  lastCountUsers: 0,
  lastOnlineStatus: '',
  maxCountUsersLimit: 0,
  minutesForDisconnect: 0,
  name: '',
  token: '',
  unitRate: 0,
  updatedAt: '',
  url: '',
  roles: []
};

export const PnetServersEditForm = ({ id }: EditFormProps) => {
  const {
    locale: { PnetServers, Forms, Sidebar }
  } = useLanguageBrowser();
  const navigate = useNavigate();
  const response = useQueryServerGet({ id });
  const initialValues = MapServerItem(response.data?.model, response.data?.roles) ?? defaultValues;
  const { mutate } = useMutationServerUpsert({
    onSuccess: (data) => {
      navigate(RoutesLocation.pnetServersEdit(data?.id?.toString() || ''), { replace: true });
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
  const onDeleteMutation = useMutationServerDelete({
    onSuccess: () => {
      navigate(RoutesLocation.pnetServers(), { replace: true });
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
  const AuthProviderDeletePopup = useConfirmPopup({
    title: PnetServers.DeletePopup.Title,
    description: PnetServers.DeletePopup.Description,
    onConfirm: onDeleteMutation.mutate.bind(onDeleteMutation.mutate, { id })
  });
  if (response.isLoading) return <Loading size='md' />;
  return (
    <Formik
      initialValues={initialValues}
      validationSchema={undefined}
      onSubmit={(values, formikHelpers) => {
        mutate({
          id: values.id,
          isActive: values.isActive,
          name: values.name || '',
          minutesForDisconnect: values.minutesForDisconnect || 0,
          maxCountUsersLimit: values.maxCountUsersLimit || 0,
          unitRate: values.unitRate,
          url: values.url || '',
          type: values.type || 'pnet',
          clientId: values.clientId || '',
          token: values.token || '',
          roles: values.roles?.map((roleId) => parseInt(roleId.toString()))
        });
      }}
    >
      {({ values, handleChange, setFieldValue, handleSubmit }) => (
        <>
          {AuthProviderDeletePopup.component({})}
          <div className='flex flex-col gap-4 mb-4'>
            <Input
              variant='bordered'
              label={PnetServers.FieldID}
              type='number'
              value={(initialValues.id ?? 0).toString()}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={PnetServers.FieldName}
              type='text'
              value={values.name ?? ''}
              onChange={handleChange('name')}
            />
            <Input
              variant='bordered'
              label={PnetServers.FieldURL}
              type='url'
              value={values.url ?? ''}
              onChange={handleChange('url')}
            />
            <Input
              variant='bordered'
              label={PnetServers.FieldClientID}
              description={PnetServers.FieldClientIDDescription}
              type='text'
              value={values.clientId ?? ''}
              onChange={handleChange('clientId')}
            />
            <RolesSelector
              label={PnetServers.FieldAllowedRoles}
              description={PnetServers.FieldAllowedRolesDescription}
              selectedKeys={values.roles ?? []}
              onSelectionChange={(keys) => setFieldValue('roles', Array.from(keys))}
            />
            <Select
              variant='bordered'
              label={PnetServers.FieldType}
              selectedKeys={[values.type ?? '']}
              onSelectionChange={(keys) => setFieldValue('type', keys.currentKey || 'pnet')}
            >
              <SelectItem key='pnet'>PNET</SelectItem>
              <SelectItem key='openid'>OpenID</SelectItem>
              <SelectItem key='k8s'>K8S</SelectItem>
            </Select>
            <Checkbox type='checkbox' defaultSelected={!!values.isActive} onChange={handleChange('isActive')}>
              {PnetServers.FieldIsActive}
            </Checkbox>
            {values.type === 'pnet' && (
              <>
                <Input
                  variant='bordered'
                  label={PnetServers.FieldUnitRate}
                  type='number'
                  value={(values.unitRate ?? '').toString()}
                  onChange={handleChange('unitRate')}
                />
                <Input
                  variant='bordered'
                  label={PnetServers.FieldMinutesForDisconnect}
                  description={PnetServers.DescriptionMinutesForDisconnect}
                  type='number'
                  value={(values.minutesForDisconnect ?? '').toString()}
                  onChange={handleChange('minutesForDisconnect')}
                />
                <Input
                  variant='bordered'
                  label={PnetServers.FieldMaxCountUsers}
                  description={PnetServers.DescriptionMaxCountUsers}
                  type='number'
                  value={(values.maxCountUsersLimit ?? '').toString()}
                  onChange={handleChange('maxCountUsersLimit')}
                />
                <Input
                  variant='bordered'
                  label={PnetServers.FieldLastOnlineStatus}
                  type='datetime-local'
                  value={dayjs(initialValues.lastOnlineStatus ?? '').format('YYYY-MM-DDTHH:mm')}
                  isReadOnly
                />
                <Input
                  variant='bordered'
                  label={PnetServers.FieldLastCountUsers}
                  type='number'
                  value={(initialValues.lastCountUsers ?? '').toString()}
                  isReadOnly
                />
              </>
            )}
            <PasswordInput
              variant='bordered'
              label={PnetServers.FieldToken}
              type='password'
              value={values.token ?? ''}
              onChange={handleChange('token')}
            />
            <Input
              variant='bordered'
              label={PnetServers.FieldCreatedAt}
              type='datetime-local'
              value={dayjs(initialValues.createdAt ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={PnetServers.FieldUpdatedAt}
              type='datetime-local'
              value={dayjs(initialValues.updatedAt ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Button onPress={() => handleSubmit()} variant='flat' color='primary'>
              {Sidebar.Save}
            </Button>
            <Button onPress={AuthProviderDeletePopup.onOpen} variant='flat' color='danger'>
              {Sidebar.Delete}
            </Button>
          </div>
        </>
      )}
    </Formik>
  );
};
