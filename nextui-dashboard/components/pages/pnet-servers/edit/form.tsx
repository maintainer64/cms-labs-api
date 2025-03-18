'use client';
import React from 'react';
import { addToast, Button, Checkbox, Input, Select, SelectItem } from '@heroui/react';
import { Formik } from 'formik';
import { models_PNETServer } from '@/helpers/api';
import useLanguageBrowser from '@/helpers/locale';
import { useNavigate } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import dayjs from 'dayjs';
import { Loading } from '@/components/scroll/loader';
import { useConfirmPopup } from '@/components/hooks/useDeletePopup';
import { usePnetServerByID } from '@/helpers/queries/pnet-server/get';
import { usePnetServerDelete } from '@/helpers/queries/pnet-server/delete';
import { usePnetServerUpsert } from '@/helpers/queries/pnet-server/upsert';
import { PasswordInput } from '@/components/base-forms/password';

interface EditFormProps {
  id?: number;
}

const defaultValues: models_PNETServer = {
  client_id: '',
  type: 'pnet',
  created_at: '',
  is_active: true,
  last_count_users: 0,
  last_online_status: '',
  max_count_users_limit: 0,
  minutes_for_disconnect: 0,
  name: '',
  token: '',
  unit_rate: 0,
  updated_at: '',
  url: ''
};

export const PnetServersEditForm = ({ id }: EditFormProps) => {
  const {
    locale: { PnetServers, Forms, Sidebar }
  } = useLanguageBrowser();
  const navigate = useNavigate();
  const response = usePnetServerByID(id);
  const initialValues = response.data?.result?.model ?? defaultValues;
  const { mutate } = usePnetServerUpsert({
    onSuccess: (data, { formikHelpers }) => {
      navigate(RoutesLocation.pnetServersEdit(data.result?.id?.toString() || ''), { replace: true });
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
  const onDeleteMutation = usePnetServerDelete({
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
        description: error.body.msg,
        color: 'danger'
      });
    }
  });
  const ltiFormDeletePopup = useConfirmPopup({
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
        mutate({ values, formikHelpers });
      }}
    >
      {({ values, handleChange, setFieldValue, handleSubmit }) => (
        <>
          {ltiFormDeletePopup.component({})}
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
              value={values.client_id ?? ''}
              onChange={handleChange('client_id')}
            />
            <Select
              variant='bordered'
              label={PnetServers.FieldType}
              selectedKeys={[values.type ?? '']}
              onSelectionChange={(keys) => setFieldValue('type', keys.currentKey || 'pnet')}
            >
              <SelectItem key='pnet'>PNET</SelectItem>
              <SelectItem key='openid'>OpenID</SelectItem>
            </Select>
            <Checkbox type='checkbox' defaultSelected={!!values.is_active} onChange={handleChange('is_active')}>
              {PnetServers.FieldIsActive}
            </Checkbox>
            {values.type === 'pnet' && (
              <>
                <Input
                  variant='bordered'
                  label={PnetServers.FieldUnitRate}
                  type='number'
                  value={(values.unit_rate ?? '').toString()}
                  onChange={handleChange('unit_rate')}
                />
                <Input
                  variant='bordered'
                  label={PnetServers.FieldMinutesForDisconnect}
                  description={PnetServers.DescriptionMinutesForDisconnect}
                  type='number'
                  value={(values.minutes_for_disconnect ?? '').toString()}
                  onChange={handleChange('minutes_for_disconnect')}
                />
                <Input
                  variant='bordered'
                  label={PnetServers.FieldMaxCountUsers}
                  description={PnetServers.DescriptionMaxCountUsers}
                  type='number'
                  value={(values.max_count_users_limit ?? '').toString()}
                  onChange={handleChange('max_count_users_limit')}
                />
                <Input
                  variant='bordered'
                  label={PnetServers.FieldLastOnlineStatus}
                  type='datetime-local'
                  value={dayjs(initialValues.last_online_status ?? '').format('YYYY-MM-DDTHH:mm')}
                  isReadOnly
                />
                <Input
                  variant='bordered'
                  label={PnetServers.FieldLastCountUsers}
                  type='number'
                  value={(initialValues.last_count_users ?? '').toString()}
                  isReadOnly
                />
              </>
            )}
            <PasswordInput
              variant='bordered'
              label={PnetServers.FieldToken}
              type='password'
              value={initialValues.token ?? ''}
            />
            <Input
              variant='bordered'
              label={PnetServers.FieldCreatedAt}
              type='datetime-local'
              value={dayjs(initialValues.created_at ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={PnetServers.FieldUpdatedAt}
              type='datetime-local'
              value={dayjs(initialValues.updated_at ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Button onPress={() => handleSubmit()} variant='flat' color='primary'>
              {Sidebar.Save}
            </Button>
            <Button onPress={ltiFormDeletePopup.onOpen} variant='flat' color='danger'>
              {Sidebar.Delete}
            </Button>
          </div>
        </>
      )}
    </Formik>
  );
};
