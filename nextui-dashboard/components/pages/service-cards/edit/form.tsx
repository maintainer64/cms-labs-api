'use client';
import React from 'react';
import { Button, Checkbox, Input } from '@nextui-org/react';
import { Formik } from 'formik';
import { models_ServiceCard } from '@/helpers/api';
import useLanguageBrowser from '@/helpers/locale';
import { useNavigate } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import dayjs from 'dayjs';
import { useAlert } from '@/components/alerts/hooks';
import { Loading } from '@/components/scroll/loader';
import { useConfirmPopup } from '@/components/hooks/useDeletePopup';
import { useServiceCardDelete } from '@/helpers/queries/service-cards/delete';
import { useServiceCardById } from '@/helpers/queries/service-cards/get';
import { useServiceCardUpsert } from '@/helpers/queries/service-cards/upsert';

interface EditFormProps {
  id?: number;
}

const defaultValues: models_ServiceCard = {
  created_at: '',
  description: '',
  image_url: '',
  url: '',
  is_active: true,
  name: '',
  order: 0,
  updated_at: ''
};

export const ServiceCardEditForm = ({ id }: EditFormProps) => {
  const {
    locale: { ServiceCards, Forms, Sidebar }
  } = useLanguageBrowser();
  const { showAlert } = useAlert();
  const navigate = useNavigate();
  const response = useServiceCardById(id);
  const initialValues = response.data?.result?.model ?? defaultValues;
  const { mutate } = useServiceCardUpsert({
    onSuccess: (data) => {
      navigate(RoutesLocation.serviceCardsEdit(data.result?.id?.toString() || ''), { replace: true });
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
  const onDeleteMutation = useServiceCardDelete({
    onSuccess: () => {
      navigate(RoutesLocation.serviceCards(), { replace: true });
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
  const serviceCardDeletePopup = useConfirmPopup({
    title: ServiceCards.DeletePopup.Title,
    description: ServiceCards.DeletePopup.Description,
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
      {({ values, handleChange, handleSubmit }) => (
        <>
          {serviceCardDeletePopup.component({})}
          <div className='flex flex-col gap-4 mb-4'>
            <Input
              variant='bordered'
              label={ServiceCards.FieldID}
              type='number'
              value={(initialValues.id ?? 0).toString()}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={ServiceCards.FieldName}
              type='text'
              value={values.name ?? ''}
              onChange={handleChange('name')}
            />
            <Input
              variant='bordered'
              label={ServiceCards.FieldDescription}
              type='text'
              value={values.description ?? ''}
              onChange={handleChange('description')}
            />
            <Input
              variant='bordered'
              label={ServiceCards.FieldURL}
              type='url'
              value={values.url ?? ''}
              onChange={handleChange('url')}
            />
            <Input
              variant='bordered'
              label={ServiceCards.FieldImageURL}
              type='url'
              value={values.image_url ?? ''}
              onChange={handleChange('image_url')}
            />
            <Checkbox type='checkbox' defaultSelected={!!values.is_active} onChange={handleChange('is_active')}>
              {ServiceCards.FieldIsActive}
            </Checkbox>
            <Input
              variant='bordered'
              label={ServiceCards.FieldOrder}
              type='number'
              value={(values.order ?? '').toString()}
              onChange={handleChange('order')}
            />
            <Input
              variant='bordered'
              label={ServiceCards.FieldCreatedAt}
              type='datetime-local'
              value={dayjs(initialValues.created_at ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={ServiceCards.FieldUpdatedAt}
              type='datetime-local'
              value={dayjs(initialValues.updated_at ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Button onPress={() => handleSubmit()} variant='flat' color='primary'>
              {Sidebar.Save}
            </Button>
            <Button onPress={serviceCardDeletePopup.onOpen} variant='flat' color='danger'>
              {Sidebar.Delete}
            </Button>
          </div>
        </>
      )}
    </Formik>
  );
};
