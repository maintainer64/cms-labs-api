'use client';
import React from 'react';
import { Button, Checkbox, Input } from '@nextui-org/react';
import { Formik } from 'formik';
import { models_PNETServer } from '@/helpers/api';
import useLanguageBrowser from '@/helpers/locale';
import { useNavigate } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import dayjs from 'dayjs';
import { useAlert } from '@/components/alerts/hooks';
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
  const { showAlert } = useAlert();
  const navigate = useNavigate();
  const response = usePnetServerByID(id);
  const initialValues = response.data?.result?.model ?? defaultValues;
  const { mutate } = usePnetServerUpsert({
    onSuccess: (data, { formikHelpers }) => {
      navigate(RoutesLocation.pnetServersEdit(data.result?.id?.toString() || ''), { replace: true });
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
  const onDeleteMutation = usePnetServerDelete({
    onSuccess: () => {
      navigate(RoutesLocation.pnetServers(), { replace: true });
      showAlert({
        type: 'success',
        message: Forms.DeleteSuccess
      });
    },
    onError: (error: any) => {
      showAlert({
        type: 'danger',
        message: Forms.DeleteError,
        description: error.body.msg
      });
    }
  });
  const ltiFormDeletePopup = useConfirmPopup({
    title: PnetServers.DeletePopup.Title,
    description: PnetServers.DeletePopup.Description,
    onConfirm: onDeleteMutation.mutate.bind(onDeleteMutation.mutate, { id })
  });
  if (response.isLoading) return <Loading size={8} />;
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
            <Checkbox type='checkbox' defaultSelected={!!values.is_active} onChange={handleChange('is_active')}>
              {PnetServers.FieldIsActive}
            </Checkbox>
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
            <PasswordInput
              variant='bordered'
              label={PnetServers.FieldToken}
              type='password'
              value={initialValues.token ?? ''}
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
