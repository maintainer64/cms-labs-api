"use client";
import React from "react";
import {
    Accordion,
    AccordionItem,
    Button,
    Dropdown,
    DropdownItem,
    DropdownMenu,
    DropdownTrigger,
    Input
} from "@nextui-org/react";
import {Formik} from "formik";
import {models_LTIForm} from "@/helpers/api";
import useLanguageBrowser from "@/helpers/locale";
import {useNavigate} from "react-router-dom";
import {RoutesLocation} from "@/components/routes";
import dayjs from "dayjs";
import {useAlert} from "@/components/alerts/hooks";
import {Loading} from "@/components/scroll/loader";
import {useLtiFormsByID} from "@/helpers/queries/lti-forms/get";
import {useLTIFormsUpsert} from "@/helpers/queries/lti-forms/upsert";
import {Textarea} from "@nextui-org/input";
import {LtiFormURILTIMoodle} from "@/components/pages/lti-forms/edit/lti-forms-popup";
import {useConfirmPopup} from "@/components/hooks/useDeletePopup";
import {useLTIFormsDelete} from "@/helpers/queries/lti-forms/delete";


interface EditFormProps {
    id?: number
}

const defaultValues: models_LTIForm = {
    base_uri: '',
    created_at: '',
    key_set_uri: '',
    lti_auth_login_uri: '',
    lti_auth_token_uri: '',
    lti_client_id: '',
    lti_deployment_id: '',
    name: '',
    private_key: '',
    public_key: '',
    target_link_uri: '',
    updated_at: '',
}

const extractUrlWithPath = (url: string, path: string) => {
    try {
        return `${(new URL(url)).origin}${path}`
    } catch (e) {
        return ""
    }
}


export const LtiIntegrationsEditForm = ({id}: EditFormProps) => {
    const {locale: {LTIForm, Forms, Sidebar}} = useLanguageBrowser();
    const {showAlert} = useAlert();
    const navigate = useNavigate();
    const response = useLtiFormsByID(id);
    const initialValues = response.data?.result?.model ?? defaultValues;
    const {mutate} = useLTIFormsUpsert({
        onSuccess: (data, {formikHelpers},) => {
            navigate(RoutesLocation.ltiFormsEdit(data.result?.id?.toString() || ''), {replace: true});
            showAlert({
                type: "success",
                message: Forms.SaveSuccess,
            });
        },
        onError: (error: any) => {
            showAlert({
                type: "danger",
                message: Forms.SaveError,
                description: error.body.msg
            });
        }
    });
    const onDeleteMutation = useLTIFormsDelete({
        onSuccess: () => {
            navigate(RoutesLocation.ltiForms(), {replace: true});
            showAlert({
                type: "success",
                message: Forms.DeleteSuccess,
            });
        },
        onError: (error: any) => {
            showAlert({
                type: "danger",
                message: Forms.DeleteError,
                description: error.body.msg
            });
        }
    });
    const ltiMoodleSettings = LtiFormURILTIMoodle();
    const ltiFormDeletePopup = useConfirmPopup({
        title: LTIForm.DeletePopup.Title,
        description: LTIForm.DeletePopup.Description,
        onConfirm: onDeleteMutation.mutate.bind(onDeleteMutation.mutate, {id})
    });
    if (response.isLoading) return <Loading size={8}/>;
    return (
        <Formik
            initialValues={initialValues}
            validationSchema={undefined}
            onSubmit={(values, formikHelpers) => {
                mutate({values, formikHelpers})
            }}>
            {({values, handleChange, setFieldValue, handleSubmit}) => (
                <>
                    {ltiMoodleSettings.component((url) => {
                        const baseURI = url.replace(/\/$/, "");
                        setFieldValue('base_uri', baseURI);
                        setFieldValue('lti_auth_login_uri', `${baseURI}/mod/lti/auth.php`);
                        setFieldValue('lti_auth_token_uri', `${baseURI}/mod/lti/token.php`);
                        setFieldValue('key_set_uri', `${baseURI}/mod/lti/certs.php`);
                    })}
                    {ltiFormDeletePopup.component({})}
                    <div className='flex flex-col w-1/2 gap-4 mb-4'>
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
                            value={values.name ?? ""}
                            onChange={handleChange('name')}
                        />
                        <div>
                            <Dropdown>
                                <DropdownTrigger>
                                    <Button variant="bordered">{LTIForm.ButtonBaseURI}</Button>
                                </DropdownTrigger>
                                <DropdownMenu aria-label="Static Actions">
                                    <DropdownItem
                                        key="ButtonBaseURIMoodle"
                                        onClick={ltiMoodleSettings.onOpen}
                                    >
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
                            value={values.base_uri ?? ""}
                            onChange={handleChange('base_uri')}
                        />
                        <Input
                            variant='bordered'
                            label={LTIForm.FieldLTIAuthLoginUri}
                            description={LTIForm.DescriptionLTIAuthLoginUri}
                            type='url'
                            value={values.lti_auth_login_uri ?? ""}
                            onChange={handleChange('lti_auth_login_uri')}
                        />
                        <Input
                            variant='bordered'
                            label={LTIForm.FieldLTIAuthTokenUri}
                            description={LTIForm.DescriptionLTIAuthTokenUri}
                            type='url'
                            value={values.lti_auth_token_uri ?? ""}
                            onChange={handleChange('lti_auth_token_uri')}
                        />
                        <Input
                            variant='bordered'
                            label={LTIForm.FieldTargetLinkUri}
                            description={LTIForm.DescriptionTargetLinkUri}
                            type='url'
                            value={values.target_link_uri ?? ""}
                            onChange={handleChange('target_link_uri')}
                        />
                        <Input
                            variant='bordered'
                            label={LTIForm.FieldKeySetURI}
                            description={LTIForm.DescriptionKeySetURI}
                            type='url'
                            value={values.key_set_uri ?? ""}
                            onChange={handleChange('key_set_uri')}
                        />
                        <Input
                            variant='bordered'
                            label={LTIForm.FieldLTIClientID}
                            type='text'
                            value={values.lti_client_id ?? ""}
                            onChange={handleChange('lti_client_id')}
                        />
                        <Input
                            variant='bordered'
                            label={LTIForm.FieldLTIDeployment}
                            type='text'
                            value={values.lti_deployment_id ?? ""}
                            onChange={handleChange('lti_deployment_id')}
                        />
                        <Accordion>
                            <AccordionItem key="1" aria-label="Параметры для Moodle" title="Параметры для Moodle">
                                <div className="flex flex-col gap-4">
                                    <Input
                                        variant='bordered'
                                        label={LTIForm.MoodleProviderParams.ToolURL}
                                        type='text'
                                        value={
                                            extractUrlWithPath(initialValues.target_link_uri, '')
                                        }
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
                                        value={initialValues.public_key ?? ""}
                                        readOnly
                                    />
                                    <Input
                                        variant='bordered'
                                        label={LTIForm.MoodleProviderParams.InitiateLoginURL}
                                        type='text'
                                        value={
                                            extractUrlWithPath(initialValues.target_link_uri, "/api/v2/lti/login")
                                        }
                                        readOnly
                                    />
                                    <Input
                                        variant='bordered'
                                        label={LTIForm.MoodleProviderParams.RedirectionURI}
                                        type='text'
                                        value={
                                            extractUrlWithPath(initialValues.target_link_uri, "/api/v2/lti/launch")
                                        }
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
                            label={LTIForm.FieldCreatedAt}
                            type='datetime-local'
                            value={dayjs(initialValues.created_at ?? "").format('YYYY-MM-DDTHH:mm')}
                            isReadOnly
                        />
                        <Input
                            variant='bordered'
                            label={LTIForm.FieldUpdatedAt}
                            type='datetime-local'
                            value={dayjs(initialValues.updated_at ?? "").format('YYYY-MM-DDTHH:mm')}
                            isReadOnly
                        />
                        <Button
                            onPress={() => handleSubmit()}
                            variant='flat'
                            color='primary'>
                            {Sidebar.Save}
                        </Button>
                        <Button
                            onPress={ltiFormDeletePopup.onOpen}
                            variant='flat'
                            color='danger'>
                            {Sidebar.Delete}
                        </Button>
                    </div>
                </>
            )}
        </Formik>
    );
}
