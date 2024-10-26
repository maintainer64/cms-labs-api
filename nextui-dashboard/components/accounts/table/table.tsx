import {Table, TableBody, TableCell, TableColumn, TableHeader, TableRow,} from "@nextui-org/react";
import React from "react";
import {RenderCell} from "./render-cell";
import useLanguageBrowser from "@/helpers/locale";
import {Loading} from "@/components/auth/loader";
import {models_UserListItem} from "@/helpers/api";

interface UsersTableWrapperProps {
    isLoading?: boolean
    users?: models_UserListItem[]
}

export const UsersTableWrapper = ({users, isLoading}: UsersTableWrapperProps) => {
    const {locale: {Tables: {UsersTable}}} = useLanguageBrowser();
    return (
        <div className=" w-full flex flex-col gap-4">
            <Table aria-label="Users table">
                <TableHeader columns={UsersTable.Columns}>
                    {(column) => (
                        <TableColumn
                            key={column.uid}
                            hideHeader={column.uid === "actions"}
                            align={column.uid === "actions" ? "center" : "start"}
                        >
                            {column.name}
                        </TableColumn>
                    )}
                </TableHeader>
                <TableBody
                    isLoading={isLoading}
                    loadingContent={<Loading/>}
                    items={users ?? []}
                >
                    {(item) => (
                        <TableRow>
                            {(columnKey) => (
                                <TableCell>
                                    {RenderCell({item, columnKey})}
                                </TableCell>
                            )}
                        </TableRow>
                    )}
                </TableBody>
            </Table>
        </div>
    );
};
