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
  Input,
  Select,
  SelectItem
} from '@heroui/react';
import { Formik } from 'formik';
import { ModelsAuthProvider } from '@/helpers/api';
import useLanguageBrowser from '@/helpers/locale';
import { useNavigate } from 'react-router-dom';
import { RoutesLocation } from '@/components/routes';
import dayjs from 'dayjs';
import { Loading } from '@/components/scroll/loader';
import { Textarea } from '@heroui/input';
import { AuthProviderURILTIMoodle } from '@/components/pages/auth-providers/edit/auth-providers-popup';
import { useConfirmPopup } from '@/components/hooks/useDeletePopup';
import { CamelCasedPropertiesDeep } from 'type-fest';
import { useQueryAuthProviderGet } from '@/helpers/queries/lti_form/use-query-lti-form-get';
import { useMutationAuthProviderUpsert } from '@/helpers/queries/lti_form/use-mutation-lti-form-upsert';
import { useMutationAuthProviderDelete } from '@/helpers/queries/lti_form/use-mutation-lti-form-delete';

interface EditFormProps {
  id?: number;
}

const defaultValues: CamelCasedPropertiesDeep<ModelsAuthProvider> = {
  type: 'lti',
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

const extractUrlWithPath = (url?: string, path?: string) => {
  try {
    return `${new URL(url || '').origin}${path || ''}`;
  } catch (e) {
    return '';
  }
};

export const AuthProviderTypes = () => {
  const {
    locale: { AuthProvider }
  } = useLanguageBrowser();
  return [
    {
      key: 'lti',
      label: AuthProvider.FieldTypeLTI,
      description: AuthProvider.FieldTypeLTIDescription
    },
    {
      key: 'ldap',
      label: AuthProvider.FieldTypeLDAP,
      description: AuthProvider.FieldTypeLDAPDescription
    }
  ];
};

export const AuthProvidersEditForm = ({ id }: EditFormProps) => {
  const {
    locale: { AuthProvider, Forms, Sidebar }
  } = useLanguageBrowser();
  const navigate = useNavigate();
  const providerTypes = AuthProviderTypes();
  const response = useQueryAuthProviderGet({ id });
  const initialValues = response?.data?.model ?? defaultValues;
  const { mutate } = useMutationAuthProviderUpsert({
    onSuccess: (data) => {
      navigate(RoutesLocation.authProvidersEdit(data?.id?.toString() || ''), { replace: true });
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
  const onDeleteMutation = useMutationAuthProviderDelete({
    onSuccess: () => {
      navigate(RoutesLocation.authProviders(), { replace: true });
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
  const ltiMoodleSettings = AuthProviderURILTIMoodle();
  const AuthProviderDeletePopup = useConfirmPopup({
    title: AuthProvider.DeletePopup.Title,
    description: AuthProvider.DeletePopup.Description,
    onConfirm: onDeleteMutation.mutate.bind(onDeleteMutation.mutate, { id })
  });
  if (response?.isLoading) return <Loading size='md' />;
  return (
    <Formik
      initialValues={initialValues}
      validationSchema={undefined}
      onSubmit={(values, formikHelpers) => {
        mutate({
          type: values.type,
          authLoginUri: values.ltiAuthLoginUri || '',
          authTokenUri: values.ltiAuthTokenUri || '',
          baseUri: values.baseUri,
          clientId: values.ltiClientId || '',
          deploymentId: values.ltiDeploymentId || '',
          id: values.id,
          keySetUri: values.keySetUri || '',
          name: values.name,
          targetLinkUri: values.targetLinkUri || '',
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
          {AuthProviderDeletePopup.component({})}
          <div className='flex flex-col gap-4 mb-4'>
            <Input
              variant='bordered'
              label={AuthProvider.FieldID}
              type='number'
              value={(initialValues.id ?? 0).toString()}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={AuthProvider.FieldName}
              type='text'
              value={values.name ?? ''}
              onChange={handleChange('name')}
            />
            <Select
              variant='bordered'
              label={AuthProvider.FieldType}
              selectedKeys={[values.type ?? '']}
              onSelectionChange={(keys) => setFieldValue('type', keys.currentKey || 'lti')}
            >
              {providerTypes.map((type) => (
                <SelectItem key={type.key} description={type.description}>
                  {type.label}
                </SelectItem>
              ))}
            </Select>
            {values.type === 'lti' && (
              <>
                <div>
                  <Dropdown>
                    <DropdownTrigger>
                      <Button variant='bordered'>{AuthProvider.ButtonBaseURI}</Button>
                    </DropdownTrigger>
                    <DropdownMenu aria-label='Static Actions'>
                      <DropdownItem key='ButtonBaseURIMoodle' onPress={ltiMoodleSettings.onOpen}>
                        {AuthProvider.ButtonBaseURIMoodle}
                      </DropdownItem>
                    </DropdownMenu>
                  </Dropdown>
                </div>
                <Input
                  variant='bordered'
                  label={AuthProvider.FieldBaseURI}
                  description={AuthProvider.DescriptionBaseURI}
                  type='url'
                  value={values.baseUri ?? ''}
                  onChange={handleChange('baseUri')}
                />
                <Input
                  variant='bordered'
                  label={AuthProvider.FieldLTIAuthLoginUri}
                  description={AuthProvider.DescriptionLTIAuthLoginUri}
                  type='url'
                  value={values.ltiAuthLoginUri ?? ''}
                  onChange={handleChange('ltiAuthLoginUri')}
                />
                <Input
                  variant='bordered'
                  label={AuthProvider.FieldLTIAuthTokenUri}
                  description={AuthProvider.DescriptionLTIAuthTokenUri}
                  type='url'
                  value={values.ltiAuthTokenUri ?? ''}
                  onChange={handleChange('ltiAuthTokenUri')}
                />
                <Input
                  variant='bordered'
                  label={AuthProvider.FieldTargetLinkUri}
                  description={AuthProvider.DescriptionTargetLinkUri}
                  type='url'
                  value={values.targetLinkUri ?? ''}
                  onChange={handleChange('targetLinkUri')}
                />
                <Input
                  variant='bordered'
                  label={AuthProvider.FieldKeySetURI}
                  description={AuthProvider.DescriptionKeySetURI}
                  type='url'
                  value={values.keySetUri ?? ''}
                  onChange={handleChange('keySetUri')}
                />
                <Input
                  variant='bordered'
                  label={AuthProvider.FieldLTIClientID}
                  type='text'
                  value={values.ltiClientId ?? ''}
                  onChange={handleChange('ltiClientId')}
                />
                <Input
                  variant='bordered'
                  label={AuthProvider.FieldLTIDeployment}
                  type='text'
                  value={values.ltiDeploymentId ?? ''}
                  onChange={handleChange('ltiDeploymentId')}
                />
                <Accordion>
                  <AccordionItem
                    key='1'
                    aria-label={AuthProvider.MoodleProviderParams.SectionTitle}
                    title={AuthProvider.MoodleProviderParams.SectionTitle}
                  >
                    <div className='flex flex-col gap-4'>
                      <Input
                        variant='bordered'
                        label={AuthProvider.MoodleProviderParams.ToolURL}
                        type='text'
                        value={extractUrlWithPath(initialValues.targetLinkUri, '')}
                        readOnly
                      />
                      <Input
                        variant='bordered'
                        label={AuthProvider.MoodleProviderParams.LTIVersion}
                        type='text'
                        value={AuthProvider.MoodleProviderParams.LTIVersionValue}
                        readOnly
                      />
                      <Input
                        variant='bordered'
                        label={AuthProvider.MoodleProviderParams.PublicKeyType}
                        type='text'
                        value={AuthProvider.MoodleProviderParams.PublicKeyTypeValue}
                        readOnly
                      />
                      <Textarea
                        variant='bordered'
                        label={AuthProvider.MoodleProviderParams.PublicKey}
                        type='text'
                        value={initialValues.publicKey ?? ''}
                        readOnly
                      />
                      <Input
                        variant='bordered'
                        label={AuthProvider.MoodleProviderParams.InitiateLoginURL}
                        type='text'
                        value={extractUrlWithPath(initialValues.targetLinkUri, '/api/v2/lti/login')}
                        readOnly
                      />
                      <Input
                        variant='bordered'
                        label={AuthProvider.MoodleProviderParams.RedirectionURI}
                        type='text'
                        value={extractUrlWithPath(initialValues.targetLinkUri, '/api/v2/lti/launch')}
                        readOnly
                      />
                      <Input
                        variant='bordered'
                        label={AuthProvider.MoodleProviderParams.DefaultLaunchContainer}
                        type='text'
                        value={AuthProvider.MoodleProviderParams.DefaultLaunchContainerValue}
                        readOnly
                      />
                      <Input
                        variant='bordered'
                        label={AuthProvider.MoodleProviderParams.IMSLTIAssignmentGradeServices}
                        type='text'
                        value={AuthProvider.MoodleProviderParams.IMSLTIAssignmentGradeServicesValue}
                        readOnly
                      />
                      <Input
                        variant='bordered'
                        label={AuthProvider.MoodleProviderParams.IMSLTINamesRoleProvisioning}
                        type='text'
                        value={AuthProvider.MoodleProviderParams.IMSLTINamesRoleProvisioningValue}
                        readOnly
                      />
                      <Input
                        variant='bordered'
                        label={AuthProvider.MoodleProviderParams.ToolSettings}
                        type='text'
                        value={AuthProvider.MoodleProviderParams.ToolSettingsValue}
                        readOnly
                      />
                      <Input
                        variant='bordered'
                        label={AuthProvider.MoodleProviderParams.ShareLauncherNameWithTool}
                        type='text'
                        value={AuthProvider.MoodleProviderParams.ShareLauncherNameWithToolValue}
                        readOnly
                      />
                      <Input
                        variant='bordered'
                        label={AuthProvider.MoodleProviderParams.ShareLauncherEmailWithTool}
                        type='text'
                        value={AuthProvider.MoodleProviderParams.ShareLauncherEmailWithToolValue}
                        readOnly
                      />
                      <Input
                        variant='bordered'
                        label={AuthProvider.MoodleProviderParams.AcceptGradesTool}
                        type='text'
                        value={AuthProvider.MoodleProviderParams.AcceptGradesToolValue}
                        readOnly
                      />
                    </div>
                  </AccordionItem>
                </Accordion>
                <Input
                  variant='bordered'
                  label={AuthProvider.FieldSSOURL}
                  description={AuthProvider.DescriptionSSOURL}
                  type='url'
                  value={values.ssoUrl ?? ''}
                  onChange={handleChange('ssoUrl')}
                />
              </>
            )}
            {values.type === 'ldap' && (
              <>
                <Input
                  variant='bordered'
                  label={AuthProvider.FieldBaseURI}
                  description={AuthProvider.DescriptionBaseURI}
                  type='url'
                  value={values.baseUri ?? ''}
                  onChange={handleChange('baseUri')}
                />
                <Input
                  variant='bordered'
                  label={AuthProvider.FieldDN}
                  value={values.keySetUri ?? ''}
                  onChange={handleChange('keySetUri')}
                />
              </>
            )}
            <Input
              variant='bordered'
              label={AuthProvider.FieldCreatedAt}
              type='datetime-local'
              value={dayjs(initialValues.createdAt ?? '').format('YYYY-MM-DDTHH:mm')}
              isReadOnly
            />
            <Input
              variant='bordered'
              label={AuthProvider.FieldUpdatedAt}
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
