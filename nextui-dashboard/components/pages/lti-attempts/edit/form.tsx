'use client';
import React from 'react';
import { addToast, Button, Input } from '@heroui/react';
import { Formik } from 'formik';
import { ModelsLTIAttempt } from '@/helpers/api';
import useLanguageBrowser from '@/helpers/locale';
import { useNavigate } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import dayjs from 'dayjs';
import { Loading } from '@/components/scroll/loader';
import { useConfirmPopup } from '@/components/hooks/useDeletePopup';
import { ServerInput } from '@/components/pages/lti-attempts/edit/autoCompleteServer';
import { CamelCasedPropertiesDeep } from 'type-fest';
import { useQueryLtiAttemptGet } from '@/helpers/queries/lti_attempt/use-query-lti-attempt-get';
import { useQueryUserGet } from '@/helpers/queries/user/use-query-user-get';
import { useQueryLtiRoutingGet } from '@/helpers/queries/lti_routing/use-query-lti-routing-get';
import { useMutationLtiAttemptUpdate } from '@/helpers/queries/lti_attempt/use-mutation-lti-attempt-update';
import { useMutationLtiAttemptDelete } from '@/helpers/queries/lti_attempt/use-mutation-lti-attempt-delete';

interface EditFormProps {
  id?: number;
}

const defaultValues: CamelCasedPropertiesDeep<ModelsLTIAttempt> = {
  attemptId: '',
  createdAt: '',
  expiredAt: '',
  ltiRoutingId: 0,
  pnetServerId: 0,
  roomId: 0,
  userId: 0,
  updatedAt: ''
};

export const LtiAttemptEditForm = ({ id }: EditFormProps) => {
  const {
    locale: { AuthProviderAttempt, Forms, Sidebar }
  } = useLanguageBrowser();
  const navigate = useNavigate();
  const response = useQueryLtiAttemptGet({ id });
  const initialValues = response.data?.model ?? defaultValues;
  const responseUser = useQueryUserGet({ id: initialValues.userId });
  const user = responseUser.data?.model;
  const responseRoute = useQueryLtiRoutingGet({ id: initialValues.ltiRoutingId });
  const route = responseRoute.data?.model;
  const isLoading = response.isLoading || responseUser.isLoading || responseRoute.isLoading;
  const { mutate } = useMutationLtiAttemptUpdate({
    onSuccess: (data) => {
      navigate(RoutesLocation.ltiAttemptEdit(data?.id?.toString() || ''), { replace: true });
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
  const onDeleteMutation = useMutationLtiAttemptDelete({
    onSuccess: () => {
      navigate(RoutesLocation.ltiAttemptUser(initialValues.userId?.toString()), { replace: true });
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
  const ltiAttemptDeletePopup = useConfirmPopup({
    title: AuthProviderAttempt.DeletePopup.Title,
    description: AuthProviderAttempt.DeletePopup.Description,
    onConfirm: onDeleteMutation.mutate.bind(onDeleteMutation.mutate, { id })
  });
  if (isLoading) return <Loading size='md' />;
  return (
    <Formik
      initialValues={initialValues}
      validationSchema={undefined}
      onSubmit={(values, formikHelpers) => {
        mutate({
          expiredAt: dayjs(values.expiredAt).format(),
          id: values.id,
          pnetServerId: parseInt(values.pnetServerId?.toString() || '')
        });
      }}
    >
      {({ values, handleChange, setFieldValue, handleSubmit }) => (
        <>
          {ltiAttemptDeletePopup.component({})}
          <div className='flex flex-col gap-4 mb-4'>
            <Input
              variant='bordered'
              label={AuthProviderAttempt.FieldID}
              value={(initialValues.attemptId ?? '').toString()}
              isReadOnly
            />
            <ServerInput
              variant='bordered'
              label={AuthProviderAttempt.FieldPNETServer}
              value={(values.pnetServerId ?? '').toString()}
              onChange={handleChange('pnetServerId')}
            />
            <Input
              variant='bordered'
              label={AuthProviderAttempt.FieldExpiredAt}
              type='datetime-local'
              value={dayjs(values.expiredAt ?? '').format('YYYY-MM-DDTHH:mm')}
              onChange={handleChange('expiredAt')}
            />
            {initialValues.roomId && (
              <Input
                variant='bordered'
                label={AuthProviderAttempt.FieldRoomNumber}
                type='number'
                value={(initialValues.roomId ?? 0).toString()}
                isReadOnly
              />
            )}
            <Input
              variant='bordered'
              label={AuthProviderAttempt.FieldUserId}
              type='number'
              value={(initialValues.userId ?? 0).toString()}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={AuthProviderAttempt.FieldUserName}
              type='text'
              value={user?.name ?? ''}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={AuthProviderAttempt.FieldUserEmail}
              type='text'
              value={user?.email ?? ''}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={AuthProviderAttempt.FieldLTIRoutingID}
              type='number'
              value={(initialValues.ltiRoutingId ?? 0).toString()}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={AuthProviderAttempt.FieldLTIRoutingName}
              type='text'
              value={route?.name ?? ''}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={AuthProviderAttempt.FieldCreatedAt}
              type='datetime-local'
              value={dayjs(initialValues.createdAt ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={AuthProviderAttempt.FieldUpdatedAt}
              type='datetime-local'
              value={dayjs(initialValues.updatedAt ?? '').format('YYYY-MM-DDTHH:mm')}
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
