'use client';
import React from 'react';
import {
  Accordion,
  AccordionItem,
  addToast,
  Button,
  Dropdown,
  DropdownItem,
  DropdownMenu,
  DropdownTrigger,
  Input
} from '@heroui/react';
import { Formik } from 'formik';
import { ModelsLTIForm } from '@/helpers/api';
import useLanguageBrowser from '@/helpers/locale';
import { useNavigate } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import dayjs from 'dayjs';
import { Loading } from '@/components/scroll/loader';
import { Textarea } from '@heroui/input';
import { LtiFormURILTIMoodle } from '@/components/pages/lti-forms/edit/lti-forms-popup';
import { useConfirmPopup } from '@/components/hooks/useDeletePopup';
import { CamelCasedPropertiesDeep } from 'type-fest';
import { useQueryLtiFormGet } from '@/helpers/queries/lti_form/use-query-lti-form-get';
import { useMutationLtiFormUpsert } from '@/helpers/queries/lti_form/use-mutation-lti-form-upsert';
import { useMutationLtiFormDelete } from '@/helpers/queries/lti_form/use-mutation-lti-form-delete';

interface EditFormProps {
  id?: number;
}

const defaultValues: CamelCasedPropertiesDeep<ModelsLTIForm> = {
  baseUri: '',
  createdAt: '',
  keySetUri: '',
  ltiAuthLoginUri: '',
  ltiAuthTokenUri: '',
  ltiClientId: '',
  ltiDeploymentId: '',
  name: '',
  privateKey: '',
  publicKey: '',
  targetLinkUri: '',
  updatedAt: ''
};

const extractUrlWithPath = (url: string, path: string) => {
  try {
    return `${new URL(url).origin}${path}`;
  } catch (e) {
    return '';
  }
};

export const LtiIntegrationsEditForm = ({ id }: EditFormProps) => {
  const {
    locale: { LTIForm, Forms, Sidebar }
  } = useLanguageBrowser();
  const navigate = useNavigate();
  const response = useQueryLtiFormGet({ id });
  const initialValues = response.data?.model ?? defaultValues;
  const { mutate } = useMutationLtiFormUpsert({
    onSuccess: (data) => {
      navigate(RoutesLocation.ltiFormsEdit(data?.id?.toString() || ''), { replace: true });
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
  const onDeleteMutation = useMutationLtiFormDelete({
    onSuccess: () => {
      navigate(RoutesLocation.ltiForms(), { replace: true });
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
  const ltiMoodleSettings = LtiFormURILTIMoodle();
  const ltiFormDeletePopup = useConfirmPopup({
    title: LTIForm.DeletePopup.Title,
    description: LTIForm.DeletePopup.Description,
    onConfirm: onDeleteMutation.mutate.bind(onDeleteMutation.mutate, { id })
  });
  if (response.isLoading) return <Loading size='md' />;
  return (
    <Formik
      initialValues={initialValues}
      validationSchema={undefined}
      onSubmit={(values, formikHelpers) => {
        mutate({
          authLoginUri: values.ltiAuthLoginUri,
          authTokenUri: values.ltiAuthTokenUri,
          baseUri: values.baseUri,
          clientId: values.ltiClientId,
          deploymentId: values.ltiDeploymentId,
          id: values.id,
          keySetUri: values.keySetUri,
          name: values.name,
          targetLinkUri: values.targetLinkUri,
          ssoUrl: values.ssoUrl ? values.ssoUrl : undefined
        });
      }}
    >
      {({ values, handleChange, setFieldValue, handleSubmit }) => (
        <>
          {ltiMoodleSettings.component((url) => {
            const baseURI = url.replace(/\/$/, '');
            setFieldValue('baseUri', baseURI);
            setFieldValue('ltiAuthLoginUri', `${baseURI}/mod/lti/auth.php`);
            setFieldValue('ltiAuthTokenUri', `${baseURI}/mod/lti/token.php`);
            setFieldValue('keySetUri', `${baseURI}/mod/lti/certs.php`);
          })}
          {ltiFormDeletePopup.component({})}
          <div className='flex flex-col gap-4 mb-4'>
            <Input
              variant='bordered'
              label={LTIForm.FieldID}
              type='number'
              value={(initialValues.id ?? 0).toString()}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={LTIForm.FieldName}
              type='text'
              value={values.name ?? ''}
              onChange={handleChange('name')}
            />
            <div>
              <Dropdown>
                <DropdownTrigger>
                  <Button variant='bordered'>{LTIForm.ButtonBaseURI}</Button>
                </DropdownTrigger>
                <DropdownMenu aria-label='Static Actions'>
                  <DropdownItem key='ButtonBaseURIMoodle' onPress={ltiMoodleSettings.onOpen}>
                    {LTIForm.ButtonBaseURIMoodle}
                  </DropdownItem>
                </DropdownMenu>
              </Dropdown>
            </div>
            <Input
              variant='bordered'
              label={LTIForm.FieldBaseURI}
              description={LTIForm.DescriptionBaseURI}
              type='url'
              value={values.baseUri ?? ''}
              onChange={handleChange('baseUri')}
            />
            <Input
              variant='bordered'
              label={LTIForm.FieldLTIAuthLoginUri}
              description={LTIForm.DescriptionLTIAuthLoginUri}
              type='url'
              value={values.ltiAuthLoginUri ?? ''}
              onChange={handleChange('ltiAuthLoginUri')}
            />
            <Input
              variant='bordered'
              label={LTIForm.FieldLTIAuthTokenUri}
              description={LTIForm.DescriptionLTIAuthTokenUri}
              type='url'
              value={values.ltiAuthTokenUri ?? ''}
              onChange={handleChange('ltiAuthTokenUri')}
            />
            <Input
              variant='bordered'
              label={LTIForm.FieldTargetLinkUri}
              description={LTIForm.DescriptionTargetLinkUri}
              type='url'
              value={values.targetLinkUri ?? ''}
              onChange={handleChange('targetLinkUri')}
            />
            <Input
              variant='bordered'
              label={LTIForm.FieldKeySetURI}
              description={LTIForm.DescriptionKeySetURI}
              type='url'
              value={values.keySetUri ?? ''}
              onChange={handleChange('keySetUri')}
            />
            <Input
              variant='bordered'
              label={LTIForm.FieldLTIClientID}
              type='text'
              value={values.ltiClientId ?? ''}
              onChange={handleChange('ltiClientId')}
            />
            <Input
              variant='bordered'
              label={LTIForm.FieldLTIDeployment}
              type='text'
              value={values.ltiDeploymentId ?? ''}
              onChange={handleChange('ltiDeploymentId')}
            />
            <Accordion>
              <AccordionItem
                key='1'
                aria-label={LTIForm.MoodleProviderParams.SectionTitle}
                title={LTIForm.MoodleProviderParams.SectionTitle}
              >
                <div className='flex flex-col gap-4'>
                  <Input
                    variant='bordered'
                    label={LTIForm.MoodleProviderParams.ToolURL}
                    type='text'
                    value={extractUrlWithPath(initialValues.targetLinkUri, '')}
                    readOnly
                  />
                  <Input
                    variant='bordered'
                    label={LTIForm.MoodleProviderParams.LTIVersion}
                    type='text'
                    value={LTIForm.MoodleProviderParams.LTIVersionValue}
                    readOnly
                  />
                  <Input
                    variant='bordered'
                    label={LTIForm.MoodleProviderParams.PublicKeyType}
                    type='text'
                    value={LTIForm.MoodleProviderParams.PublicKeyTypeValue}
                    readOnly
                  />
                  <Textarea
                    variant='bordered'
                    label={LTIForm.MoodleProviderParams.PublicKey}
                    type='text'
                    value={initialValues.publicKey ?? ''}
                    readOnly
                  />
                  <Input
                    variant='bordered'
                    label={LTIForm.MoodleProviderParams.InitiateLoginURL}
                    type='text'
                    value={extractUrlWithPath(initialValues.targetLinkUri, '/api/v2/lti/login')}
                    readOnly
                  />
                  <Input
                    variant='bordered'
                    label={LTIForm.MoodleProviderParams.RedirectionURI}
                    type='text'
                    value={extractUrlWithPath(initialValues.targetLinkUri, '/api/v2/lti/launch')}
                    readOnly
                  />
                  <Input
                    variant='bordered'
                    label={LTIForm.MoodleProviderParams.DefaultLaunchContainer}
                    type='text'
                    value={LTIForm.MoodleProviderParams.DefaultLaunchContainerValue}
                    readOnly
                  />
                  <Input
                    variant='bordered'
                    label={LTIForm.MoodleProviderParams.IMSLTIAssignmentGradeServices}
                    type='text'
                    value={LTIForm.MoodleProviderParams.IMSLTIAssignmentGradeServicesValue}
                    readOnly
                  />
                  <Input
                    variant='bordered'
                    label={LTIForm.MoodleProviderParams.IMSLTINamesRoleProvisioning}
                    type='text'
                    value={LTIForm.MoodleProviderParams.IMSLTINamesRoleProvisioningValue}
                    readOnly
                  />
                  <Input
                    variant='bordered'
                    label={LTIForm.MoodleProviderParams.ToolSettings}
                    type='text'
                    value={LTIForm.MoodleProviderParams.ToolSettingsValue}
                    readOnly
                  />
                  <Input
                    variant='bordered'
                    label={LTIForm.MoodleProviderParams.ShareLauncherNameWithTool}
                    type='text'
                    value={LTIForm.MoodleProviderParams.ShareLauncherNameWithToolValue}
                    readOnly
                  />
                  <Input
                    variant='bordered'
                    label={LTIForm.MoodleProviderParams.ShareLauncherEmailWithTool}
                    type='text'
                    value={LTIForm.MoodleProviderParams.ShareLauncherEmailWithToolValue}
                    readOnly
                  />
                  <Input
                    variant='bordered'
                    label={LTIForm.MoodleProviderParams.AcceptGradesTool}
                    type='text'
                    value={LTIForm.MoodleProviderParams.AcceptGradesToolValue}
                    readOnly
                  />
                </div>
              </AccordionItem>
            </Accordion>
            <Input
              variant='bordered'
              label={LTIForm.FieldSSOURL}
              description={LTIForm.DescriptionSSOURL}
              type='url'
              value={values.ssoUrl ?? ''}
              onChange={handleChange('ssoUrl')}
            />
            <Input
              variant='bordered'
              label={LTIForm.FieldCreatedAt}
              type='datetime-local'
              value={dayjs(initialValues.createdAt ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={LTIForm.FieldUpdatedAt}
              type='datetime-local'
              value={dayjs(initialValues.updatedAt ?? '').format('YYYY-MM-DDTHH:mm')}
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
