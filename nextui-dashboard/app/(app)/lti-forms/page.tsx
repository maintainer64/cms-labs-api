import React from "react";
import {LTIFormsList} from "../../../components/pages/lti-forms";
import {Layout} from "@/components/layout/layout";
import {LtiFormsEdit} from "@/components/pages/lti-forms/edit/lti-forms-edit";

export const LTIFormsPage = () => {
    return <Layout><LTIFormsList/></Layout>;
};

export const LTIFormsPageEdit = () => {
    return <Layout><LtiFormsEdit/></Layout>;LTIFormsPageEdit
};
