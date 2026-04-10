'use client';
import React from 'react';
import { Accordion, AccordionItem, addToast, Button, Checkbox, Input, Select, SelectItem } from '@heroui/react';
import { Formik } from 'formik';
import { ModelsLTIRouting } from '@/helpers/api';
import useLanguageBrowser from '@/helpers/locale';
import { useNavigate } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import dayjs from 'dayjs';
import { Loading } from '@/components/scroll/loader';
import { useConfirmPopup } from '@/components/hooks/useDeletePopup';
import { LabsPathInput, TestsPathInput } from './autoCompletePath';
import { ServerInput } from '@/components/pages/lti-attempts/edit/autoCompleteServer';
import { CamelCasedPropertiesDeep } from 'type-fest';
import { useMutationLtiRoutingUpsert } from '@/helpers/queries/lti_routing/use-mutation-lti-routing-upsert';
import { useQueryLtiRoutingGet } from '@/helpers/queries/lti_routing/use-query-lti-routing-get';
import { useMutationLtiRoutingDelete } from '@/helpers/queries/lti_routing/use-mutation-lti-routing-delete';

interface EditFormProps {
  id?: number;
}

const defaultValues: CamelCasedPropertiesDeep<ModelsLTIRouting> = {
  collaboration: 0,
  createdAt: '',
  id: undefined,
  ltiDescription: '',
  ltiParamsTask: '',
  ltiTitle: '',
  name: '',
  pinnedSessionMinutes: 0,
  pnetLabsType: 'default',
  pnetLabsPath: '',
  pnetTestPath: '',
  pnetServerId: 0,
  isDefault: false,
  updatedAt: '',
  ltiTaskId: '',
  ltiCourseId: '',
  ltiSubId: ''
};

export const LtiRoutingLabsType = () => {
  const {
    locale: { LTIRouting }
  } = useLanguageBrowser();
  return [
    {
      key: 'default',
      label: LTIRouting.FieldPNETLabsTypeDefault,
      description: LTIRouting.FieldPNETLabsTypeDefaultDescription
    },
    {
      key: 'curl',
      label: LTIRouting.FieldPNETLabsTypeCurl,
      description: LTIRouting.FieldPNETLabsTypeCurlDescription
    },
    {
      key: 'sso',
      label: LTIRouting.FieldPNETLabsTypeSSO,
      description: LTIRouting.FieldPNETLabsTypeSSODescription
    },
    {
      key: 'clabgate',
      label: LTIRouting.FieldPNETLabsTypeClabgate,
      description: LTIRouting.FieldPNETLabsTypeClabgateDescription
    }
  ];
};

export const LtiRoutingEditForm = ({ id }: EditFormProps) => {
  const {
    locale: { LTIRouting, Forms, Sidebar }
  } = useLanguageBrowser();
  const navigate = useNavigate();
  const response = useQueryLtiRoutingGet({ id });
  const routingTypes = LtiRoutingLabsType();
  const initialValues = response.data?.model ?? defaultValues;
  const { mutate } = useMutationLtiRoutingUpsert({
    onSuccess: (data) => {
      navigate(RoutesLocation.ltiRoutingEdit(data?.id?.toString() || ''), { replace: true });
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
  const onDeleteMutation = useMutationLtiRoutingDelete({
    onSuccess: () => {
      navigate(RoutesLocation.ltiRouting(), { replace: true });
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
    title: LTIRouting.DeletePopup.Title,
    description: LTIRouting.DeletePopup.Description,
    onConfirm: onDeleteMutation.mutate.bind(onDeleteMutation.mutate, { id })
  });
  if (response.isLoading) return <Loading size='md' />;
  return (
    <Formik
      initialValues={initialValues}
      validationSchema={undefined}
      onSubmit={(values) => {
        mutate({
          collaboration: values.collaboration || 0,
          id: values.id,
          ltiDescription: values.ltiDescription,
          ltiParamsTask: values.ltiParamsTask,
          ltiTaskId: values.ltiTaskId,
          ltiCourseId: values.ltiCourseId,
          ltiSubId: values.ltiSubId,
          ltiTitle: values.ltiTitle,
          name: values.name,
          pinnedSessionMinutes: values.pinnedSessionMinutes || 0,
          pnetLabsType: values.pnetLabsType,
          pnetLabsPath: values.pnetLabsPath,
          pnetTestPath: values.pnetTestPath,
          pnetServerId: parseInt(values.pnetServerId?.toString() || '0'),
          isDefault: values.isDefault
        });
      }}
    >
      {({ values, handleChange, setFieldValue, handleSubmit }) => (
        <>
          {AuthProviderDeletePopup.component({})}
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
            <Checkbox type='checkbox' defaultSelected={values.isDefault} onChange={handleChange('isDefault')}>
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
                    value={values.ltiTitle ?? ''}
                    onChange={handleChange('ltiTitle')}
                  />
                  <Input
                    variant='bordered'
                    label={LTIRouting.FieldLTIDescription}
                    type='text'
                    value={values.ltiDescription ?? ''}
                    onChange={handleChange('ltiDescription')}
                  />
                  <Input
                    variant='bordered'
                    label={LTIRouting.FieldLTIParamsTask}
                    type='text'
                    value={values.ltiParamsTask ?? ''}
                    onChange={handleChange('ltiParamsTask')}
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
                    value={(values.pinnedSessionMinutes || 0).toString()}
                    onChange={handleChange('pinnedSessionMinutes')}
                  />
                  <Select
                    variant='bordered'
                    label={LTIRouting.FieldPNETLabsType}
                    selectedKeys={[values.pnetLabsType ?? '']}
                    onSelectionChange={(keys) => setFieldValue('pnetLabsType', keys.currentKey || 'default')}
                  >
                    {routingTypes.map((type) => (
                      <SelectItem key={type.key} description={type.description}>
                        {type.label}
                      </SelectItem>
                    ))}
                  </Select>
                  <LabsPathInput
                    labsTypeUnl={values.pnetLabsType ?? 'default'}
                    variant='bordered'
                    label={LTIRouting.FieldPNETLabsPath}
                    value={values.pnetLabsPath ?? ''}
                    onChange={handleChange('pnetLabsPath')}
                  />
                  <TestsPathInput
                    labsTypeUnl={values.pnetLabsType ?? 'default'}
                    variant='bordered'
                    label={LTIRouting.FieldPNETTestPath}
                    value={values.pnetTestPath ?? ''}
                    onChange={handleChange('pnetTestPath')}
                  />
                  <ServerInput
                    variant='bordered'
                    defaultItems={[{ key: 0, value: LTIRouting.FieldPNETServerDefault }]}
                    label={LTIRouting.FieldPNETServer}
                    value={values.pnetServerId?.toString() ?? ''}
                    onChange={handleChange('pnetServerId')}
                  />
                </div>
              </AccordionItem>
            </Accordion>
            <Input
              variant='bordered'
              label={LTIRouting.FieldCreatedAt}
              type='datetime-local'
              value={dayjs(initialValues.createdAt ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={LTIRouting.FieldUpdatedAt}
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
