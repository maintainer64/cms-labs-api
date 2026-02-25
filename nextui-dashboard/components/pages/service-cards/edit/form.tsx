'use client';
import React from 'react';
import { addToast, Button, Checkbox, Input } from '@heroui/react';
import { Formik } from 'formik';
import { ModelsServiceCard } from '@/helpers/api';
import useLanguageBrowser from '@/helpers/locale';
import { useNavigate } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import dayjs from 'dayjs';
import { Loading } from '@/components/scroll/loader';
import { useConfirmPopup } from '@/components/hooks/useDeletePopup';
import { CamelCasedPropertiesDeep } from 'type-fest';
import { useQueryServiceCardGet } from '@/helpers/queries/service_card/use-query-service-card-get';
import { useMutationServiceCardUpsert } from '@/helpers/queries/service_card/use-mutation-service-card-upsert';
import { useMutationServiceCardDelete } from '@/helpers/queries/service_card/use-mutation-service-card-delete';

interface EditFormProps {
  id?: number;
}

const defaultValues: CamelCasedPropertiesDeep<ModelsServiceCard> = {
  createdAt: '',
  description: '',
  imageUrl: '',
  url: '',
  isActive: true,
  name: '',
  order: 0,
  updatedAt: ''
};

export const ServiceCardEditForm = ({ id }: EditFormProps) => {
  const {
    locale: { ServiceCards, Forms, Sidebar }
  } = useLanguageBrowser();
  const navigate = useNavigate();
  const response = useQueryServiceCardGet({ id });
  const initialValues = response.data?.model ?? defaultValues;
  const { mutate } = useMutationServiceCardUpsert({
    onSuccess: (data) => {
      navigate(RoutesLocation.serviceCardsEdit(data?.id?.toString() || ''), { replace: true });
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
  const onDeleteMutation = useMutationServiceCardDelete({
    onSuccess: () => {
      navigate(RoutesLocation.serviceCards(), { replace: true });
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
  const serviceCardDeletePopup = useConfirmPopup({
    title: ServiceCards.DeletePopup.Title,
    description: ServiceCards.DeletePopup.Description,
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
          isActive: values.isActive ?? true,
          name: values.name || '',
          description: values.description || '',
          order: values.order || 0,
          url: values.url || '',
          imageUrl: values.imageUrl || ''
        });
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
              value={values.imageUrl ?? ''}
              onChange={handleChange('imageUrl')}
            />
            <Checkbox type='checkbox' defaultSelected={!!values.isActive} onChange={handleChange('isActive')}>
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
              value={dayjs(initialValues.createdAt ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={ServiceCards.FieldUpdatedAt}
              type='datetime-local'
              value={dayjs(initialValues.updatedAt ?? '').format('YYYY-MM-DDTHH:mm')}
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
