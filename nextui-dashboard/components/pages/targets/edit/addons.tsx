'use client';
import React, {useState} from 'react';
import {
    Button,
    Select,
    SelectItem,
    Table,
    TableBody,
    TableCell,
    TableColumn,
    TableHeader,
    TableRow,
} from '@heroui/react';
import useLanguageBrowser from "@/helpers/locale";


interface AddonsTabProps {
    targetId: string;
    availableAddons: Array<{ id?: string; name?: string; type?: string }>;
    connectedAddons: Array<{ addonId?: string; type?: string }>;
    onConnectAddon: any;
    onDisconnectAddon: any;
}

export const AddonsTab = ({
                              targetId,
                              availableAddons,
                              connectedAddons,
                              onConnectAddon,
                              onDisconnectAddon,
                          }: AddonsTabProps) => {
    const {locale: {Target: TargetLocale, Sidebar}} = useLanguageBrowser();
    const [selectedAddon, setSelectedAddon] = useState('');

    const connectedAddonIds = connectedAddons.map(a => a.addonId);
    const availableToConnect = availableAddons.filter(a => !connectedAddonIds.includes(a.id));

    const handleConnect = () => {
        if (selectedAddon && targetId) {
            onConnectAddon({targetId, addonId: selectedAddon});
            setSelectedAddon('');
        }
    };

    return (
        <div className='flex flex-col gap-4 p-4 border rounded-lg'>
            <div className='flex gap-2 items-center'>
                <Select
                    variant='bordered'
                    label={TargetLocale.Addon.SelectAddon || 'Выберите аддон'}
                    selectedKeys={selectedAddon ? [selectedAddon] : []}
                    onChange={(e) => setSelectedAddon(e.target.value)}
                    className='flex-1'
                >
                    {availableToConnect.map((addon) => (
                        <SelectItem key={addon.id}>
                            {addon.name} ({addon.type})
                        </SelectItem>
                    ))}
                </Select>
                <Button color='primary' onPress={handleConnect} isDisabled={!selectedAddon}>
                    {TargetLocale.Addon.Connect || 'Подключить'}
                </Button>
            </div>

            <Table aria-label='Connected addons'>
                <TableHeader>
                    <TableColumn key='name'>Название</TableColumn>
                    <TableColumn key='type'>Тип</TableColumn>
                    <TableColumn key='actions'>{''}</TableColumn>
                </TableHeader>
                <TableBody>
                    {connectedAddons.length === 0 ? (
                        <TableRow key='empty'>
                            <TableCell>{'-'}</TableCell>
                            <TableCell>{'-'}</TableCell>
                            <TableCell>{''}</TableCell>
                        </TableRow>
                    ) : (
                        connectedAddons.map((addon, idx) => {
                            const addonInfo = availableAddons.find(a => a.id === addon.addonId);
                            return (
                                <TableRow key={idx}>
                                    <TableCell>{addonInfo?.name || addon.addonId}</TableCell>
                                    <TableCell>{addon.type}</TableCell>
                                    <TableCell>
                                        <Button size='sm' color='danger' variant='flat'
                                                onPress={() => onDisconnectAddon({targetId, addonId: addon.addonId})}>
                                            {TargetLocale.Addon.Disconnect || 'Отключить'}
                                        </Button>
                                    </TableCell>
                                </TableRow>
                            );
                        })
                    )}
                </TableBody>
            </Table>
        </div>
    );
};