"use client";
import React from "react";
import {HouseIcon} from "@/components/icons/breadcrumb/house-icon";
import {UsersIcon} from "@/components/icons/breadcrumb/users-icon";
import {RoutesLocation} from "@/components/routes";
import useLanguageBrowser from "@/helpers/locale";
import {CrumbsLayout} from "@/components/layout/crumbs";
import {AccountsEditForm} from "@/components/pages/accounts/edit/form";
import {useParams} from "react-router-dom";


export const AccountsEdit = () => {
    let {id} = useParams();
    const {locale} = useLanguageBrowser();
    const {locale: {Tables: {UsersTable}}} = useLanguageBrowser();
    const crumbs = [
        {
            icon: <HouseIcon/>,
            name: locale.Sidebar.Home,
            href: RoutesLocation.home(),
        },
        {
            icon: <UsersIcon/>,
            name: locale.Sidebar.Users,
            href: RoutesLocation.accounts(),
        },
        {
            icon: undefined,
            name: locale.Sidebar.Edit,
            href: "#",
        },
    ];
    return (
        <CrumbsLayout name={UsersTable.Title} crumbs={crumbs}>
            <div className="max-w-[95rem] mx-auto w-full">
                <AccountsEditForm id={parseInt(id ?? "", 10)}/>
            </div>
        </CrumbsLayout>
    );
};
