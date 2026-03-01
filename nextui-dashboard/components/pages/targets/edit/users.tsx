'use client';
import React, {useState} from 'react';
import {
    Button,
    Chip,
    Modal,
    ModalBody,
    ModalContent,
    ModalFooter,
    ModalHeader,
    Select,
    SelectItem,
    Table,
    TableBody,
    TableCell,
    TableColumn,
    TableHeader,
    TableRow,
    useDisclosure,
} from '@heroui/react';

import {UserEmailInput} from '@/components/pages/targets/edit/autoCompleteUser';
import useLanguageBrowser from '@/helpers/locale';
import {useQueryUserList} from "@/helpers/queries/user/use-query-user-list";

interface UsersTabProps {
    targetId: string;
    targetUsers: TargetUser[];

    onUpsertUser: (data: { targetId: string; userId: number; roles: string[] }) => void;
    onRemoveUser: (data: { targetId: string; userId: number }) => void;
}

type TargetUser = { userId?: number; roles?: string[] };

type ModalForm = {
    mode: 'add' | 'edit';
    userId: number | null;
    roles: string[];
};

export const UsersTab = ({targetId, targetUsers, onUpsertUser, onRemoveUser}: UsersTabProps) => {
    const {locale: {Target: TargetLocale, Sidebar}} = useLanguageBrowser();
    const {isOpen, onOpen, onClose} = useDisclosure();

    const [form, setForm] = useState<ModalForm>({
        mode: 'add',
        userId: null,
        roles: [],
    });

    const targetUsersIds: number[] = targetUsers
        ?.map((u) => u.userId)
        ?.filter((id): id is number => id !== undefined && id !== null) || [];

    const userListQuery = targetUsersIds ? useQueryUserList({
        userIds: targetUsersIds,
        offset: 0,
        limit: targetUsersIds?.length
    }) : undefined;
    const targetUsersList = userListQuery?.data?.model || []

    const roles = TargetLocale.Roles;

    const openAdd = () => {
        setForm({mode: 'add', userId: null, roles: []});
        onOpen();
    };

    const openEdit = (u: TargetUser) => {
        setForm({mode: 'edit', userId: u.userId ?? null, roles: u.roles ?? []});
        onOpen();
    };

    const save = () => {
        if (!form.userId) return;
        onUpsertUser({targetId, userId: form.userId, roles: form.roles});
        onClose();
    };

    const remove = () => {
        if (!form.userId) return;
        onRemoveUser({targetId, userId: form.userId});
        onClose();
    };

    return (
        <div className="flex flex-col gap-4 p-4">
            <Modal isOpen={isOpen} onClose={onClose}>
                <ModalContent>
                    <ModalHeader>
                        {form.mode === 'add'
                            ? TargetLocale.Users.AddTitle
                            : TargetLocale.Users.EditTitle}
                    </ModalHeader>

                    <ModalBody className="flex flex-col gap-4">
                        <UserEmailInput
                            label={TargetLocale?.Users?.UserLabel ?? 'Пользователь'}
                            variant="bordered"
                            isReadOnly={form.mode === 'edit'}
                            value={form.userId?.toString()}
                            placeholder={TargetLocale?.Users?.UserPlaceholder ?? 'Начните вводить имя'}
                            onChange={(event) => {
                                setForm((p) => ({
                                    ...p,
                                    userId: parseInt(event.target.value),
                                }));
                            }}
                        />

                        <Select
                            label={TargetLocale?.Users?.RolesLabel ?? 'Роли'}
                            variant="bordered"
                            selectionMode="multiple"
                            selectedKeys={form.roles}
                            onSelectionChange={(keys) =>
                                setForm((p) => ({
                                    ...p,
                                    roles: Array.from(keys) as string[],
                                }))
                            }
                        >
                            {roles.map(({key, value}) => (
                                <SelectItem key={key}>{value}</SelectItem>
                            ))}
                        </Select>
                    </ModalBody>

                    <ModalFooter>
                        {form.mode === 'edit' ? (
                            <Button color="danger" variant="light" onPress={remove}>
                                {TargetLocale?.Users?.Remove ?? 'Удалить'}
                            </Button>
                        ) : (
                            <Button color="danger" variant="light" onPress={onClose}>
                                {Sidebar?.Cancel ?? 'Отмена'}
                            </Button>
                        )}

                        <Button color="primary" onPress={save} isDisabled={!form.userId}>
                            {Sidebar?.Save ?? 'Сохранить'}
                        </Button>
                    </ModalFooter>
                </ModalContent>
            </Modal>

            <div className="flex justify-end">
                <Button color="primary" onPress={openAdd}>
                    {TargetLocale?.Users?.AddButton ?? 'Добавить пользователя'}
                </Button>
            </div>

            {targetUsers.length === 0 ? (
                <div className="text-center text-gray-400 py-8">
                    {TargetLocale?.Users?.NoUsers ?? 'Нет привязанных пользователей'}
                </div>
            ) : (
                <Table aria-label="Target users" removeWrapper className="border rounded-lg">
                    <TableHeader>
                        <TableColumn>{TargetLocale?.Users?.UserLabel ?? 'Пользователь'}</TableColumn>
                        <TableColumn>{TargetLocale?.Users?.RolesLabel ?? 'Роли'}</TableColumn>
                    </TableHeader>

                    <TableBody>
                        {targetUsers.map((u) => (
                            <TableRow
                                key={String(u.userId)}
                                className="cursor-pointer hover:bg-default-100 transition-colors"
                                onClick={() => openEdit(u)}
                            >
                                <TableCell>
                                    <div className="flex flex-col">
                                        <div className="font-medium leading-5">
                                            {targetUsersList.find((user) => user.model.id === u.userId)?.model?.name}
                                        </div>
                                        <div className="text-sm text-gray-500 leading-5">
                                            {targetUsersList.find((user) => user.model.id === u.userId)?.model?.email}
                                        </div>
                                    </div>
                                </TableCell>

                                <TableCell>
                                    <div className="flex flex-wrap gap-1">
                                        {u.roles?.length ? (
                                            u.roles.map((r) => (
                                                <Chip key={r} size="sm" variant="flat">
                                                    {roles.find((role) => role.key === r)?.value || r}
                                                </Chip>
                                            ))
                                        ) : (
                                            <span className="text-sm text-gray-400">
                        {TargetLocale?.Users?.NoRoles ?? 'Нет ролей'}
                      </span>
                                        )}
                                    </div>
                                </TableCell>
                            </TableRow>
                        ))}
                    </TableBody>
                </Table>
            )}
        </div>
    );
};
