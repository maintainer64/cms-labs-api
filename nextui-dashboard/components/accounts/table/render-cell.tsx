import {Chip} from "@nextui-org/react";
import React from "react";
import {EditIcon} from "../../icons/table/edit-icon";
import {models_UserListItem} from "@/helpers/api";
import dayjs from "dayjs";
import {Link} from "react-router-dom";
import {RoutesLocation} from "@/components/routes";

interface Props {
    item: models_UserListItem;
    columnKey: string | React.Key;
}

export const RenderCell = ({item, columnKey}: Props) => {
    switch (columnKey) {
        case "id":
            return (
                <div>
                    <div>
                        <span>#{item.id ?? 0}</span>
                    </div>
                </div>
            );
        case "name":
            return (
                <div>
                    <div>
                        <span>{item.name}</span>
                    </div>
                    <div>
                        <span>{item.email}</span>
                    </div>
                </div>
            );
        case "role":
            return (
                <div>
                    <div>
                        <span>{item.user_role}</span>
                    </div>
                </div>
            );
        case "group":
            return (
                <div>
                    <div>
                        <span>{item.group_name}</span>
                    </div>
                </div>
            );
        case "status":
            return (
                <Chip
                    size="sm"
                    variant="flat"
                    color={item.deleted_at ? "danger" : "success"}
                >
                    <span className="capitalize text-xs">
                        {dayjs(item.deleted_at || item.created_at).format('DD.MM.YYYY HH:mm')}
                    </span>
                </Chip>
            );
        case "actions":
            return (
                <div className="flex items-center gap-4 ">
                    <div>
                        <Link to={RoutesLocation.accountsEdit(item.id?.toString())}>
                            <EditIcon size={20} fill="#979797"/>
                        </Link>
                    </div>
                </div>
            );
        default:
            return "";
    }
};
