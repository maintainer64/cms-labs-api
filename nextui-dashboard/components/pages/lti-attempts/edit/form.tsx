'use client';
import React from 'react';
import { Button, Input } from '@heroui/react';
import { Formik } from 'formik';
import { models_LTIAttempt } from '@/helpers/api';
import useLanguageBrowser from '@/helpers/locale';
import { useNavigate } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import dayjs from 'dayjs';
import { useAlert } from '@/components/alerts/hooks';
import { Loading } from '@/components/scroll/loader';
import { useConfirmPopup } from '@/components/hooks/useDeletePopup';
import { useLTIAttemptById } from '@/helpers/queries/lti-attempt/get';
import { useLTIAttemptDelete } from '@/helpers/queries/lti-attempt/delete';
import { useLTIAttemptUpsert } from '@/helpers/queries/lti-attempt/upsert';
import { useUserByID } from '@/helpers/queries/users/get';
import { useLTIRoutingByID } from '@/helpers/queries/lti-routing/get';

interface EditFormProps {
  id?: number;
}

const defaultValues: models_LTIAttempt = {
  created_at: '',
  expired_at: '',
  lti_routing_id: 0,
  pnet_server_id: 0,
  room_number: 0,
  user_id: 0,
  updated_at: ''
};

export const LtiAttemptEditForm = ({ id }: EditFormProps) => {
  const {
    locale: { LTIFormAttempt, Forms, Sidebar }
  } = useLanguageBrowser();
  const { showAlert } = useAlert();
  const navigate = useNavigate();
  const response = useLTIAttemptById(id);
  const initialValues = response.data?.result?.model ?? defaultValues;
  const responseUser = useUserByID(initialValues.user_id);
  const user = responseUser.data?.result?.model;
  const responseRoute = useLTIRoutingByID(initialValues.lti_routing_id);
  const route = responseRoute.data?.result?.model;
  const isLoading = response.isLoading || responseUser.isLoading || responseRoute.isLoading;
  const { mutate } = useLTIAttemptUpsert({
    onSuccess: (data, { formikHelpers }) => {
      navigate(RoutesLocation.ltiFormsEdit(data.result?.id?.toString() || ''), { replace: true });
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
  const onDeleteMutation = useLTIAttemptDelete({
    onSuccess: () => {
      navigate(RoutesLocation.ltiAttemptUser(initialValues.user_id?.toString()), { replace: true });
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
  const ltiAttemptDeletePopup = useConfirmPopup({
    title: LTIFormAttempt.DeletePopup.Title,
    description: LTIFormAttempt.DeletePopup.Description,
    onConfirm: onDeleteMutation.mutate.bind(onDeleteMutation.mutate, { id })
  });
  if (isLoading) return <Loading size={8} />;
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
          {ltiAttemptDeletePopup.component({})}
          <div className='flex flex-col gap-4 mb-4'>
            <Input
              variant='bordered'
              label={LTIFormAttempt.FieldID}
              type='number'
              value={(initialValues.id ?? 0).toString()}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={LTIFormAttempt.FieldUserId}
              type='number'
              value={(initialValues.user_id ?? 0).toString()}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={LTIFormAttempt.FieldUserName}
              type='text'
              value={user?.name ?? ''}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={LTIFormAttempt.FieldUserEmail}
              type='text'
              value={user?.email ?? ''}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={LTIFormAttempt.FieldLTIRoutingID}
              type='number'
              value={(initialValues.lti_routing_id ?? 0).toString()}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={LTIFormAttempt.FieldLTIRoutingName}
              type='text'
              value={route?.name ?? ''}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={LTIFormAttempt.FieldExpiredAt}
              type='datetime-local'
              value={dayjs(values.expired_at ?? '').format('YYYY-MM-DDTHH:mm')}
              onChange={handleChange('expired_at')}
            />
            <Input
              variant='bordered'
              label={LTIFormAttempt.FieldCreatedAt}
              type='datetime-local'
              value={dayjs(initialValues.created_at ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={LTIFormAttempt.FieldUpdatedAt}
              type='datetime-local'
              value={dayjs(initialValues.updated_at ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Button onPress={() => handleSubmit()} variant='flat' color='primary'>
              {Sidebar.Save}
            </Button>
            <Button onPress={ltiAttemptDeletePopup.onOpen} variant='flat' color='danger'>
              {Sidebar.Delete}
            </Button>
          </div>
        </>
      )}
    </Formik>
  );
};
