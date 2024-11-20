import React from "react";
import {EditIcon} from "../../../icons/table/edit-icon";
import {models_PNETServerListItem} from "@/helpers/api";
import {Link} from "react-router-dom";
import {RoutesLocation} from "@/components/routes";
import {Chip} from "@nextui-org/react";
import dayjs from "dayjs";

interface Props {
    item: models_PNETServerListItem;
    columnKey: string | React.Key;
}

const isConnectDistribution = (lastOnlineStatus?: string, minutesForDisconnect?: number): boolean => {
    if (!minutesForDisconnect) return true;
    if (!lastOnlineStatus) return false;
    const now = dayjs();
    const lastOnline = dayjs(lastOnlineStatus);
    return lastOnline.add(minutesForDisconnect, 'minute').isAfter(now);
}

export const RenderCellWithLocale = (locale: any, {item, columnKey}: Props) => {
    const {locale: {Tables: {PnetServersTable: {ColumnStatus}}}} = locale;
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
                        <span>{item.url}</span>
                    </div>
                </div>
            );
        case "unitRate":
            return (
                <div>
                    <div>
                        <span>{item.unit_rate}</span>
                    </div>
                </div>
            );
        case "status":
            return (
                <div className="flex flex-col gap-2">
                    <div>
                        {item.is_active ? (<Chip size="sm" color="success">{ColumnStatus.Activated}</Chip>) : (
                            <Chip size="sm" color="danger">{ColumnStatus.Deactivated}</Chip>)}
                    </div>
                    <div>
                        {isConnectDistribution(item.last_online_status, item.minutes_for_disconnect) ? (
                            <Chip size="sm" color="success">{ColumnStatus.ConnectDistribution}</Chip>) : (
                            <Chip size="sm" color="danger">{ColumnStatus.DisconnectDistribution}</Chip>)}
                    </div>
                </div>
            );
        case "actions":
            return (
                <div className="flex items-center gap-4 ">
                    <div>
                        <Link to={RoutesLocation.pnetServersEdit(item.id?.toString())}>
                            <EditIcon size={20} fill="#979797"/>
                        </Link>
                    </div>
                </div>
            );
        default:
            return "";
    }
};
