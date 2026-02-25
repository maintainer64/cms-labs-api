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
import { Textarea } from '@heroui/input';
import { useMutationCurlRequestUpsert } from '@/helpers/queries/curl_request/use-mutation-curl-request-upsert';
import { curlParsedData, curlParsedDataValidator } from '@/components/pages/curl-requests/edit/validate';
import { useQueryCurlRequestGet } from '@/helpers/queries/curl_request/use-query-curl-request-get';
import { CamelCasedPropertiesDeep } from 'type-fest';
import { ModelsCurlRequest } from '@/helpers/api';
import { useMutationCurlRequestDelete } from '@/helpers/queries/curl_request/use-mutation-curl-request-delete';

interface EditFormProps {
  id?: number;
}

const defaultValues: CamelCasedPropertiesDeep<ModelsCurlRequest> = {
  body: undefined,
  createdAt: '',
  headers: undefined,
  id: undefined,
  method: undefined,
  name: undefined,
  raw: '',
  timeout: 300,
  updatedAt: '',
  url: undefined
};

export const CurlRequestEditForm = ({ id }: EditFormProps) => {
  const {
    locale: { CurlRequest, Forms, Sidebar }
  } = useLanguageBrowser();
  const navigate = useNavigate();
  const response = useQueryCurlRequestGet({ id });
  const initialValues = response.data?.model ?? defaultValues;
  const { mutate } = useMutationCurlRequestUpsert({
    onSuccess: (data) => {
      navigate(RoutesLocation.curlRequestEdit(data?.id?.toString() || ''), { replace: true });
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
  const onDeleteMutation = useMutationCurlRequestDelete({
    onSuccess: () => {
      navigate(RoutesLocation.curlRequest(), { replace: true });
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
  const curlRequestDeletePopup = useConfirmPopup({
    title: CurlRequest.DeletePopup.Title,
    description: CurlRequest.DeletePopup.Description,
    onConfirm: onDeleteMutation.mutate.bind(onDeleteMutation.mutate, { id })
  });
  if (response.isLoading) return <Loading size='md' />;
  return (
    <Formik
      initialValues={initialValues}
      onSubmit={(values, formikHelpers) => {
        if (values === null) return;
        const curlParsed = curlParsedData(values.raw);
        const timeout = parseInt(`${values.timeout}`);
        mutate({
          body: curlParsed.body,
          headers: curlParsed.headers,
          id: values.id,
          method: curlParsed.method,
          name: curlParsed.name,
          rawRequest: values.raw,
          timeout: timeout,
          url: curlParsed.url
        });
      }}
    >
      {({ values, handleChange, setFieldValue, handleSubmit }) => (
        <>
          {curlRequestDeletePopup.component({})}
          <div className='flex flex-col gap-4 mb-4'>
            <Input
              variant='bordered'
              label={CurlRequest.FieldID}
              type='number'
              value={(initialValues.id ?? 0).toString()}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={CurlRequest.FieldName}
              type='text'
              value={values.name ?? ''}
              onChange={handleChange('name')}
            />
            <Input
              variant='bordered'
              label={CurlRequest.FieldTimeout}
              description={CurlRequest.FieldTimeoutDescription}
              type='number'
              value={values.timeout?.toString() ?? '300'}
              onChange={handleChange('timeout')}
            />
            <Textarea
              variant='bordered'
              minRows={5}
              label={CurlRequest.FieldCurl}
              description={CurlRequest.FieldCurlDescription}
              type='text'
              value={values.raw ?? ''}
              onChange={handleChange('raw')}
              validate={curlParsedDataValidator}
            />
            <Input
              variant='bordered'
              label={CurlRequest.FieldCreatedAt}
              type='datetime-local'
              value={dayjs(initialValues.createdAt ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={CurlRequest.FieldUpdatedAt}
              type='datetime-local'
              value={dayjs(initialValues.updatedAt ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Button onPress={() => handleSubmit()} variant='flat' color='primary'>
              {Sidebar.Save}
            </Button>
            <Button onPress={curlRequestDeletePopup.onOpen} variant='flat' color='danger'>
              {Sidebar.Delete}
            </Button>
          </div>
        </>
      )}
    </Formik>
  );
};
