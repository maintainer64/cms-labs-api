'use client';
import React from 'react';
import { Accordion, AccordionItem, Button, Checkbox, Input, Select, SelectItem } from '@nextui-org/react';
import { Formik } from 'formik';
import { models_LTIRouting } from '@/helpers/api';
import useLanguageBrowser from '@/helpers/locale';
import { useNavigate } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import dayjs from 'dayjs';
import { useAlert } from '@/components/alerts/hooks';
import { Loading } from '@/components/scroll/loader';
import { useConfirmPopup } from '@/components/hooks/useDeletePopup';
import { useLTIRoutingDelete } from '@/helpers/queries/lti-routing/delete';
import { useLTIRoutingByID } from '@/helpers/queries/lti-routing/get';
import { useLTIRoutingUpsert } from '@/helpers/queries/lti-routing/upsert';
import { LabsPathInput, TestsPathInput } from './autoCompletePath';

interface EditFormProps {
  id?: number;
}

const defaultValues: models_LTIRouting = {
  collaboration: 0,
  created_at: '',
  id: undefined,
  lti_description: '',
  lti_params_task: '',
  lti_task_id: '',
  lti_title: '',
  name: '',
  pinned_session_minutes: 0,
  pnet_labs_type: 'default',
  pnet_labs_path: '',
  pnet_test_path: '',
  is_default: false,
  updated_at: ''
};

export const LtiRoutingLabsType = () => {
  const {
    locale: { LTIRouting }
  } = useLanguageBrowser();
  return [
    { key: 'default', label: LTIRouting.FieldPNETLabsTypeDefault },
    { key: 'file', label: LTIRouting.FieldPNETLabsTypeFile },
    {
      key: 'enumeration',
      label: LTIRouting.FieldPNETLabsTypeEnumeration
    }
  ];
};

export const LtiRoutingEditForm = ({ id }: EditFormProps) => {
  const {
    locale: { LTIRouting, Forms, Sidebar }
  } = useLanguageBrowser();
  const { showAlert } = useAlert();
  const navigate = useNavigate();
  const response = useLTIRoutingByID(id);
  const routingTypes = LtiRoutingLabsType();
  const initialValues = response.data?.result?.model ?? defaultValues;
  const { mutate } = useLTIRoutingUpsert({
    onSuccess: (data, { formikHelpers }) => {
      navigate(RoutesLocation.ltiRoutingEdit(data.result?.id?.toString() || ''), { replace: true });
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
  const onDeleteMutation = useLTIRoutingDelete({
    onSuccess: () => {
      navigate(RoutesLocation.ltiRouting(), { replace: true });
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
    title: LTIRouting.DeletePopup.Title,
    description: LTIRouting.DeletePopup.Description,
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
              label={LTIRouting.FieldID}
              type='number'
              value={(initialValues.id ?? 0).toString()}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={LTIRouting.FieldName}
              type='text'
              value={values.name ?? ''}
              onChange={handleChange('name')}
            />
            <Checkbox type='checkbox' defaultSelected={values.is_default} onChange={handleChange('is_default')}>
              {LTIRouting.FieldIsDefault}
            </Checkbox>
            <Accordion>
              <AccordionItem key='1' aria-label={LTIRouting.SectionRouteParams} title={LTIRouting.SectionRouteParams}>
                <div className='flex flex-col gap-4'>
                  <p>{LTIRouting.SectionRouteParamsDescription}</p>
                  <Input
                    variant='bordered'
                    label={LTIRouting.FieldLTITitle}
                    type='text'
                    value={values.lti_title ?? ''}
                    onChange={handleChange('lti_title')}
                  />
                  <Input
                    variant='bordered'
                    label={LTIRouting.FieldLTIDescription}
                    type='text'
                    value={values.lti_description ?? ''}
                    onChange={handleChange('lti_description')}
                  />
                  <Input
                    variant='bordered'
                    label={LTIRouting.FieldLTITaskID}
                    type='text'
                    value={values.lti_task_id ?? ''}
                    onChange={handleChange('lti_task_id')}
                  />
                  <Input
                    variant='bordered'
                    label={LTIRouting.FieldLTIParamsTask}
                    type='text'
                    value={values.lti_params_task ?? ''}
                    onChange={handleChange('lti_params_task')}
                  />
                </div>
              </AccordionItem>
              <AccordionItem key='2' aria-label={LTIRouting.SectionActionParams} title={LTIRouting.SectionActionParams}>
                <div className='flex flex-col gap-4'>
                  <Input
                    variant='bordered'
                    label={LTIRouting.FieldCollaboration}
                    type='number'
                    value={(values.collaboration || 0).toString()}
                    onChange={handleChange('collaboration')}
                  />
                  <Input
                    variant='bordered'
                    label={LTIRouting.FieldPinnedSessionMinutes}
                    type='number'
                    value={(values.pinned_session_minutes || 0).toString()}
                    onChange={handleChange('pinned_session_minutes')}
                  />
                  <Select
                    variant='bordered'
                    label={LTIRouting.FieldPNETLabsType}
                    selectedKeys={[values.pnet_labs_type ?? '']}
                    onSelectionChange={(keys) => setFieldValue('pnet_labs_type', keys.currentKey || 'default')}
                  >
                    {routingTypes.map((type) => (
                      <SelectItem key={type.key}>{type.label}</SelectItem>
                    ))}
                  </Select>
                  <LabsPathInput
                    labsTypeUnl={values.pnet_labs_type ?? 'default'}
                    variant='bordered'
                    label={LTIRouting.FieldPNETLabsPath}
                    type='url'
                    value={values.pnet_labs_path ?? ''}
                    onChange={handleChange('pnet_labs_path')}
                  />
                  <TestsPathInput
                    labsTypeUnl={values.pnet_labs_type ?? 'default'}
                    variant='bordered'
                    label={LTIRouting.FieldPNETTestPath}
                    type='url'
                    value={values.pnet_test_path ?? ''}
                    onChange={handleChange('pnet_test_path')}
                  />
                </div>
              </AccordionItem>
            </Accordion>
            <Input
              variant='bordered'
              label={LTIRouting.FieldCreatedAt}
              type='datetime-local'
              value={dayjs(initialValues.created_at ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={LTIRouting.FieldUpdatedAt}
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
