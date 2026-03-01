'use client';
import React, {useMemo, useState} from 'react';
import {
    addToast,
    Button,
    Chip,
    Divider,
    Input,
    Modal,
    ModalBody,
    ModalContent,
    ModalFooter,
    ModalHeader,
    Select,
    SelectItem,
    Tab,
    Table,
    TableBody,
    TableCell,
    TableColumn,
    TableHeader,
    TableRow,
    Tabs,
    Textarea,
    useDisclosure,
} from '@heroui/react';
import {Formik} from 'formik';
import useLanguageBrowser from '@/helpers/locale';
import {useNavigate} from 'react-router-dom';
import {RoutesLocation} from '@/components/routes';
import dayjs from 'dayjs';
import {Loading} from '@/components/scroll/loader';
import {useConfirmPopup} from '@/components/hooks/useDeletePopup';
import {useQueryTargetGet} from '@/helpers/queries/target/use-query-target-get';
import {useMutationTargetUpsert} from '@/helpers/queries/target/use-mutation-target-upsert';
import {useMutationTargetDelete} from '@/helpers/queries/target/use-mutation-target-delete';
import {useQueryTargetList} from '@/helpers/queries/target/use-query-target-list';
import {useMutationTargetRelationCreate} from '@/helpers/queries/target/use-mutation-target-relation-create';
import {useMutationTargetRelationDelete} from '@/helpers/queries/target/use-mutation-target-relation-delete';
import {useMutationTargetUserUpsert} from '@/helpers/queries/target/use-mutation-target-user-upsert';
import {useMutationTargetUserDelete} from '@/helpers/queries/target/use-mutation-target-user-delete';
import {useMutationTargetAddonCreate} from '@/helpers/queries/target/use-mutation-target-addon-create';
import {useMutationTargetAddonDelete} from '@/helpers/queries/target/use-mutation-target-addon-delete';
import {useQueryUserList} from '@/helpers/queries/user/use-query-user-list';
import {AutoCompleteFull} from '@/components/base-forms/autocomplete';
import {Cpu, Globe, HardDrive, MemoryStick, Network, Plus, Tag} from 'lucide-react';
import {UsersTab} from "@/components/pages/targets/edit/users";
import {AddonsTab} from "@/components/pages/targets/edit/addons";

interface EditFormProps {
    id?: string;
}

const TARGET_TYPES = [
    {key: 'server', value: 'server'},
    {key: 'virtual', value: 'virtual'},
    {key: 'service', value: 'service'},
    {key: 'module', value: 'module'}
];

const LINK_TYPES = [
    {key: 'tag', value: 'tag', icon: Tag},
    {key: 'disk', value: 'disk', icon: HardDrive},
    {key: 'ip', value: 'ip', icon: Network},
    {key: 'ram', value: 'ram', icon: MemoryStick},
    {key: 'cpu', value: 'cpu', icon: Cpu},
    {key: 'url', value: 'url', icon: Globe}
];

interface LinkItem {
    type: string;
    value: string;
}

interface TargetFullResponse {
    id?: string;
    name?: string;
    description?: string;
    type?: string;
    tags?: string[];
    links?: LinkItem[];
    internalTags?: string[];
    internalLinks?: LinkItem[];
    createdAt?: string;
    updatedAt?: string;
    targetUsers?: Array<{ userId: number; roles: string[] }>;
    availableAddons?: Array<{ id: string; name: string; type: string }>;
    connectedAddons?: Array<{ addonId: string; type: string }>;
}

const defaultValues: TargetFullResponse = {
    id: '',
    name: '',
    description: '',
    type: 'server',
    tags: [],
    links: [],
    internalTags: [],
    internalLinks: [],
    createdAt: '',
    updatedAt: '',
    targetUsers: [],
    availableAddons: [],
    connectedAddons: []
};

const getLinkIcon = (type: string) => {
    const linkType = LINK_TYPES.find(lt => lt.value === type);
    if (linkType) {
        const Icon = linkType.icon;
        return <Icon className='w-4 h-4'/>;
    }
    return <Globe className='w-4 h-4'/>;
};

// Вспомогательная функция для проверки, является ли строка URL
const isHttpUrl = (value: string) => {
    return value.startsWith('http://') || value.startsWith('https://');
};

export const TargetEditForm = ({id}: EditFormProps) => {
    const {
        locale: {Target: TargetLocale, Forms, Sidebar}
    } = useLanguageBrowser();
    const navigate = useNavigate();
    const isEdit = !!id;
    const response = useQueryTargetGet({id: id || ''});

    const initialValues = useMemo(() => {
        const data = response.data;
        if (!data) return defaultValues;
        return {
            ...defaultValues,
            ...data,
            tags: data.tags || [],
            links: data.links || [],
            internalTags: data.internalTags || [],
            internalLinks: data.internalLinks || [],
            targetUsers: data.targetUsers || [],
            availableAddons: data.availableAddons || [],
            connectedAddons: data.connectedAddons || []
        };
    }, [response.data, response.isLoading]);

    // Состояния для модального окна добавления тега/ссылки
    const {isOpen, onOpen, onClose} = useDisclosure();
    const [modalItemType, setModalItemType] = useState('tag'); // 'tag' или один из LINK_TYPES
    const [modalItemValue, setModalItemValue] = useState('');

    const [selectedParentId, setSelectedParentId] = useState('');

    const listResponse = useQueryTargetList({});
    const parentTargets = useMemo(() => {
        const targets = listResponse.data?.model?.targets || [];
        return targets
            .filter((t) => t?.taget?.id !== id)
            .map((t) => ({
                key: t?.taget?.id || '',
                value: t?.taget?.name || ''
            }));
    }, [listResponse.data, id]);

    const existingRelations = useMemo(() => {
        const relations = listResponse.data?.model?.relations || [];
        return relations
            .filter((r) => r?.relation?.toTargetId === id && r?.relation?.relationType === 'depends_on')
            .map((r) => r?.relation?.fromTargetId || '');
    }, [listResponse.data, id]);

    const {mutate: upsertMutate} = useMutationTargetUpsert({
        onSuccess: (data) => {
            if (!isEdit) {
                navigate(RoutesLocation.targetsEdit(data?.id?.toString() || ''), {replace: true});
            }
            addToast({title: Forms.SaveSuccess, color: 'success'});
        },
        onError: (error: any) => {
            addToast({title: Forms.SaveError, description: error?.data?.message, color: 'danger'});
        }
    });

    const {mutate: deleteMutate} = useMutationTargetDelete({
        onSuccess: () => {
            navigate(RoutesLocation.targets(), {replace: true});
            addToast({title: Forms.DeleteSuccess, color: 'success'});
        },
        onError: (error: any) => {
            addToast({title: Forms.DeleteError, description: error?.data?.message, color: 'danger'});
        }
    });

    const {mutate: relationCreateMutate} = useMutationTargetRelationCreate({
        onSuccess: () => {
            addToast({title: Forms.SaveSuccess, color: 'success'});
            listResponse.refetch();
        },
        onError: (error: any) => {
            addToast({title: Forms.SaveError, description: error?.data?.message, color: 'danger'});
        }
    });

    const {mutate: relationDeleteMutate} = useMutationTargetRelationDelete({
        onSuccess: () => {
            addToast({title: Forms.SaveSuccess, color: 'success'});
            listResponse.refetch();
        },
        onError: (error: any) => {
            addToast({title: Forms.SaveError, description: error?.data?.message, color: 'danger'});
        }
    });

    const {mutate: userUpsertMutate} = useMutationTargetUserUpsert({
        onSuccess: () => {
            addToast({title: Forms.SaveSuccess, color: 'success'});
            response.refetch();
        },
        onError: (error: any) => {
            addToast({title: Forms.SaveError, description: error?.data?.message, color: 'danger'});
        }
    });

    const {mutate: userDeleteMutate} = useMutationTargetUserDelete({
        onSuccess: () => {
            addToast({title: Forms.DeleteSuccess, color: 'success'});
            response.refetch();
        },
        onError: (error: any) => {
            addToast({title: Forms.DeleteError, description: error?.data?.message, color: 'danger'});
        }
    });

    const {mutate: addonCreateMutate} = useMutationTargetAddonCreate({
        onSuccess: () => {
            addToast({title: Forms.SaveSuccess, color: 'success'});
            response.refetch();
        },
        onError: (error: any) => {
            addToast({title: Forms.SaveError, description: error?.data?.message, color: 'danger'});
        }
    });

    const {mutate: addonDeleteMutate} = useMutationTargetAddonDelete({
        onSuccess: () => {
            addToast({title: Forms.DeleteSuccess, color: 'success'});
            response.refetch();
        },
        onError: (error: any) => {
            addToast({title: Forms.DeleteError, description: error?.data?.message, color: 'danger'});
        }
    });

    const deletePopup = useConfirmPopup({
        title: TargetLocale.DeletePopup.Title,
        description: TargetLocale.DeletePopup.Description,
        onConfirm: deleteMutate.bind(deleteMutate, {id: id || ''})
    });

    const handleAddParent = () => {
        if (selectedParentId && id) {
            relationCreateMutate({
                fromTargetId: selectedParentId,
                toTargetId: id,
                relationType: 'depends_on'
            });
        }
    };

    const handleRemoveParent = (parentId: string) => {
        if (parentId && id) {
            relationDeleteMutate({
                fromTargetId: parentId,
                toTargetId: id,
                relationType: 'depends_on'
            });
        }
    };

    // Функции добавления через модалку
    const handleModalAdd = (values: any, setFieldValue: any) => {
        if (!modalItemValue.trim()) return;

        if (modalItemType === 'tag') {
            // Добавляем тег
            const tags = [...(values.tags || []), modalItemValue.trim()];
            setFieldValue('tags', tags);
            handleSubmit(values, {}, {...values, tags});
        } else {
            // Добавляем ссылку с выбранным типом
            const links = [...(values.links || []), {type: modalItemType, value: modalItemValue.trim()}];
            setFieldValue('links', links);
            handleSubmit(values, {}, {...values, links});
        }

        // Сброс модалки
        setModalItemType('tag');
        setModalItemValue('');
        onClose();
    };

    const handleDeleteTag = (index: number, values: any, setFieldValue: any, isInternal: boolean) => {
        if (isInternal) return;
        const tags = [...(values.tags || [])];
        tags.splice(index, 1);
        setFieldValue('tags', tags);
        handleSubmit({...values, tags}, {}, {...values, tags});
    };

    const handleDeleteLink = (index: number, values: any, setFieldValue: any, isInternal: boolean) => {
        if (isInternal) return;
        const links = [...(values.links || [])];
        links.splice(index, 1);
        setFieldValue('links', links);
        handleSubmit({...values, links}, {}, {...values, links});
    };

    const handleSubmit = (values: any, formikHelpers: any, extraData?: any) => {
        const dataToSave = extraData || values;
        const links = (dataToSave.links || []) as any[];
        upsertMutate({
            id: values.id,
            name: values.name || '',
            description: values.description,
            type: values.type || 'server',
            tags: dataToSave.tags as string[],
            links: links.map((l) => ({type: l?.type, value: l?.value}))
        });
    };

    // Обработчик клика по ссылке
    const handleLinkClick = (linkValue: string) => {
        if (isHttpUrl(linkValue)) {
            window.open(linkValue, '_blank', 'noopener,noreferrer');
        }
    };

    if (response.isLoading) return <Loading size='md'/>;

    return (
        <Formik
            initialValues={initialValues}
            validationSchema={undefined}
            onSubmit={handleSubmit}
        >
            {({values, handleChange, handleSubmit, setFieldValue}) => (
                <>
                    {deletePopup.component({})}

                    {/* Модальное окно добавления тега/ссылки */}
                    <Modal isOpen={isOpen} onClose={onClose}>
                        <ModalContent>
                            <ModalHeader>{'Добавить тег или ссылку'}</ModalHeader>
                            <ModalBody>
                                <div className='flex flex-col gap-4'>
                                    <Select
                                        label='Тип'
                                        variant='bordered'
                                        selectedKeys={[modalItemType]}
                                        onChange={(e) => setModalItemType(e.target.value)}
                                    >
                                        {/* Разделитель не обязателен, но можно визуально отделить */}
                                        {LINK_TYPES.map((lt) => (
                                            <SelectItem key={lt.value} startContent={getLinkIcon(lt.value)}>
                                                {lt.value}
                                            </SelectItem>
                                        ))}
                                    </Select>
                                    <Input
                                        label='Значение'
                                        variant='bordered'
                                        value={modalItemValue}
                                        onChange={(e) => setModalItemValue(e.target.value)}
                                        placeholder='Введите текст или URL'
                                    />
                                </div>
                            </ModalBody>
                            <ModalFooter>
                                <Button color='danger' variant='light' onPress={onClose}>
                                    Отмена
                                </Button>
                                <Button color='primary' onPress={() => handleModalAdd(values, setFieldValue)}>
                                    Добавить
                                </Button>
                            </ModalFooter>
                        </ModalContent>
                    </Modal>

                    <Tabs aria-label='Target options'>
                        <Tab key='main' title={TargetLocale.Relation?.Title === 'Связи' ? 'Основное' : 'Main'}>
                            <div className='flex flex-col gap-4 mb-4 p-4 border rounded-lg'>
                                {isEdit && (
                                    <Input
                                        variant='bordered'
                                        label={TargetLocale.FieldID}
                                        value={initialValues.id ?? ''}
                                        isReadOnly
                                    />
                                )}
                                <Input
                                    variant='bordered'
                                    label={TargetLocale.FieldName}
                                    type='text'
                                    value={values.name ?? ''}
                                    onChange={handleChange('name')}
                                    isReadOnly={isEdit}
                                />
                                <Textarea
                                    variant='bordered'
                                    label={TargetLocale.FieldDescription}
                                    value={values.description ?? ''}
                                    onChange={handleChange('description')}
                                    minRows={3}
                                />
                                <Select
                                    variant='bordered'
                                    label={TargetLocale.FieldType}
                                    selectedKeys={values.type ? [values.type] : []}
                                    onChange={handleChange('type')}
                                    isDisabled={isEdit} // при редактировании нельзя менять тип
                                >
                                    {TARGET_TYPES.map((t) => (
                                        <SelectItem key={t.value}>
                                            {TargetLocale.Types[t.value as keyof typeof TargetLocale.Types]}
                                        </SelectItem>
                                    ))}
                                </Select>

                                <Divider className='my-2'/>
                                <div className='flex justify-between items-center'>
                                    <h4 className='font-semibold'>{TargetLocale.FieldTags || 'Теги и ссылки'}</h4>
                                    <Button size='sm' color='primary' variant='flat' onPress={onOpen}
                                            startContent={<Plus className='w-4 h-4'/>}>
                                        Добавить
                                    </Button>
                                </div>

                                {/* Объединённый блок тегов и ссылок */}
                                <div className='flex flex-wrap gap-2'>
                                    {/* Обычные теги */}
                                    {((values.tags as string[]) || []).map((tag, idx) => (
                                        <Chip
                                            key={`tag-${idx}`}
                                            variant='flat'
                                            onClose={() => handleDeleteTag(idx, values, setFieldValue, false)}
                                            className='cursor-pointer'
                                        >
                                            {tag}
                                        </Chip>
                                    ))}

                                    {/* Обычные ссылки */}
                                    {((values.links as LinkItem[]) || []).map((link, idx) => {
                                        const isLink = link.type === 'url' && isHttpUrl(link.value);
                                        return (
                                            <Chip
                                                key={`link-${idx}`}
                                                variant='flat'
                                                onClose={() => handleDeleteLink(idx, values, setFieldValue, false)}
                                                className={isLink ? 'cursor-pointer' : ''}
                                                startContent={getLinkIcon(link.type)}
                                                onClick={isLink ? () => handleLinkClick(link.value) : undefined}
                                            >
                                                {link.value}
                                            </Chip>
                                        );
                                    })}

                                    {/* Internal теги (без кнопки удаления) */}
                                    {((values.internalTags as string[]) || []).map((tag, idx) => (
                                        <Chip
                                            key={`internal-tag-${idx}`}
                                            variant='flat'
                                            className='opacity-70 cursor-default'
                                        >
                                            {tag}
                                        </Chip>
                                    ))}

                                    {/* Internal ссылки (без кнопки удаления, но кликабельны если http) */}
                                    {((values.internalLinks as LinkItem[]) || []).map((link, idx) => {
                                        const isLink = link.type === 'url' && isHttpUrl(link.value);
                                        return (
                                            <Chip
                                                key={`internal-link-${idx}`}
                                                variant='flat'
                                                className={isLink ? 'cursor-pointer opacity-70' : 'opacity-70 cursor-default'}
                                                startContent={getLinkIcon(link.type)}
                                                onClick={isLink ? () => handleLinkClick(link.value) : undefined}
                                            >
                                                {link.value}
                                            </Chip>
                                        );
                                    })}

                                    {/* Если ничего нет */}
                                    {(!values.tags?.length && !values.links?.length && !values.internalTags?.length && !values.internalLinks?.length) && (
                                        <span className='text-gray-400 text-sm'>Нет тегов или ссылок</span>
                                    )}
                                </div>

                                <Divider className='my-2'/>
                                <h4 className='font-semibold'>{TargetLocale.Relation?.Title}</h4>
                                {existingRelations.length > 0 ? (
                                    <div className='flex items-center gap-2'>
                                        <Chip>{parentTargets.find((p) => p.key === existingRelations[0])?.value || existingRelations[0]}</Chip>
                                        <Button
                                            size='sm'
                                            color='danger'
                                            variant='flat'
                                            onPress={() => handleRemoveParent(existingRelations[0])}
                                        >
                                            {TargetLocale.Relation?.RemoveParent}
                                        </Button>
                                    </div>
                                ) : (
                                    <div className='flex gap-2 items-center'>
                                        <div className='flex-1'>
                                            <AutoCompleteFull
                                                label={TargetLocale.Relation?.SelectParent}
                                                placeholder={TargetLocale.Relation?.SelectParent}
                                                defaultItems={parentTargets}
                                                value={selectedParentId}
                                                onChange={(e: any) => setSelectedParentId(e.target.value)}
                                            />
                                        </div>
                                        <Button color='primary' onPress={handleAddParent}>
                                            {TargetLocale.Relation?.AddParent}
                                        </Button>
                                    </div>
                                )}

                                {isEdit && (
                                    <>
                                        <Divider className='my-2'/>
                                        <Input
                                            variant='bordered'
                                            label={TargetLocale.FieldCreatedAt}
                                            type='datetime-local'
                                            value={dayjs(initialValues.createdAt ?? '').format('YYYY-MM-DDTHH:mm')}
                                            isReadOnly
                                        />
                                        <Input
                                            variant='bordered'
                                            label={TargetLocale.FieldUpdatedAt}
                                            type='datetime-local'
                                            value={dayjs(initialValues.updatedAt ?? '').format('YYYY-MM-DDTHH:mm')}
                                            isReadOnly
                                        />
                                    </>
                                )}
                                <Button onPress={() => handleSubmit()} variant='flat' color='primary'>
                                    {Sidebar.Save}
                                </Button>
                                {isEdit && (
                                    <Button onPress={deletePopup.onOpen} variant='flat' color='danger'>
                                        {Sidebar.Delete}
                                    </Button>
                                )}
                            </div>
                        </Tab>

                        <Tab key='users' title={TargetLocale.User?.Title || 'Пользователи'}>
                            <UsersTab
                                targetId={id || ''}
                                targetUsers={initialValues.targetUsers || []}
                                onUpsertUser={userUpsertMutate}
                                onRemoveUser={userDeleteMutate}
                            />
                        </Tab>

                        <Tab key='addons' title={TargetLocale.Addon?.Title || 'Аддоны'}>
                            <AddonsTab
                                targetId={id || ''}
                                availableAddons={initialValues.availableAddons || []}
                                connectedAddons={initialValues.connectedAddons || []}
                                onConnectAddon={addonCreateMutate}
                                onDisconnectAddon={addonDeleteMutate}
                            />
                        </Tab>
                    </Tabs>
                </>
            )}
        </Formik>
    );
};

