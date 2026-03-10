import { useEffect, useState } from 'react';
import {
  Button,
  Card,
  CardBody,
  Checkbox,
  Dropdown,
  DropdownItem,
  DropdownMenu,
  DropdownTrigger,
  Input,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  Select,
  SelectItem,
  useDisclosure
} from '@heroui/react';
import { Database, HardDrive, MoreVertical, RotateCcw, Server, Trash2, XCircle } from 'lucide-react';
import useLanguageBrowser from '@/helpers/locale';
import { useQueryUserGet } from '@/helpers/queries/user/use-query-user-get';
import { useUserProfile } from '@/components/providers/auth-jwt/hooks';

interface ConnectedAddon {
  addonId?: string;
  type?: string;
  requestDeletedUserId?: number;
  config?: Record<string, any>;
}

interface AddonsTabProps {
  targetId: string;
  availableAddons: Array<{
    id?: string;
    name?: string;
    type?: string;
    configFields?: Array<{ name: string; label: string; type: 'text' | 'number' | 'url'; required?: boolean }>;
  }>;
  connectedAddons: ConnectedAddon[];
  onConnectAddon: (data: { targetId: string; addonId: string; config?: Record<string, any> }) => void;
  onDisconnectAddon: (data: { targetId: string; addonId: string; revoke?: boolean }) => void;
  onResetAddon?: (data: { targetId: string; addonId: string }) => void;
}

export const AddonsTab = ({
  targetId,
  availableAddons,
  connectedAddons,
  onConnectAddon,
  onDisconnectAddon,
  onResetAddon
}: AddonsTabProps) => {
  const {
    locale: { Target: TargetLocale, Sidebar }
  } = useLanguageBrowser();
  const user = useUserProfile();

  const { isOpen: isDeleteModalOpen, onOpen: onDeleteModalOpen, onClose: onDeleteModalClose } = useDisclosure();
  const { isOpen: isConnectModalOpen, onOpen: onConnectModalOpen, onClose: onConnectModalClose } = useDisclosure();
  const { isOpen: isResetModalOpen, onOpen: onResetModalOpen, onClose: onResetModalClose } = useDisclosure();
  const { isOpen: isRevokeModalOpen, onOpen: onRevokeModalOpen, onClose: onRevokeModalClose } = useDisclosure();

  const [selectedAddonForConnect, setSelectedAddonForConnect] = useState('');
  const [configValues, setConfigValues] = useState<Record<string, any>>({});

  const [deleteModalAddon, setDeleteModalAddon] = useState<ConnectedAddon | null>(null);
  const [confirmSteps, setConfirmSteps] = useState({ step1: false, step2: false, step3: false });

  const [resetModalAddon, setResetModalAddon] = useState<ConnectedAddon | null>(null);
  const [revokeModalAddon, setRevokeModalAddon] = useState<ConnectedAddon | null>(null);

  const pendingAddon = connectedAddons.find((a) => a.requestDeletedUserId && a.requestDeletedUserId > 0);
  const pendingUserQuery = useQueryUserGet({ id: pendingAddon?.requestDeletedUserId });
  const pendingUser = pendingUserQuery?.data?.model;

  const connectedAddonIds = connectedAddons.map((a) => a.addonId);
  const availableToConnect = availableAddons.filter((a) => !connectedAddonIds.includes(a.id));

  useEffect(() => {
    setConfigValues({});
  }, [selectedAddonForConnect]);

  const selectedAddonInfo = availableAddons.find((a) => a.id === selectedAddonForConnect);

  const handleConnectConfirm = () => {
    if (selectedAddonForConnect && targetId) {
      onConnectAddon({
        targetId,
        addonId: selectedAddonForConnect,
        config: configValues
      });
      setSelectedAddonForConnect('');
      setConfigValues({});
      onConnectModalClose();
    }
  };

  const handleDeleteRequest = (addon: ConnectedAddon) => {
    setDeleteModalAddon(addon);
    setConfirmSteps({ step1: false, step2: false, step3: false });
    onDeleteModalOpen();
  };

  const handleConfirmDelete = () => {
    if (deleteModalAddon?.addonId) {
      onDisconnectAddon({ targetId, addonId: deleteModalAddon.addonId });
      onDeleteModalClose();
      setDeleteModalAddon(null);
    }
  };

  const handleResetRequest = (addon: ConnectedAddon) => {
    setResetModalAddon(addon);
    onResetModalOpen();
  };

  const handleConfirmReset = () => {
    if (resetModalAddon?.addonId && onResetAddon) {
      onResetAddon({ targetId, addonId: resetModalAddon.addonId });
      onResetModalClose();
      setResetModalAddon(null);
    }
  };

  const handleRevokeRequest = (addon: ConnectedAddon) => {
    setRevokeModalAddon(addon);
    onRevokeModalOpen();
  };

  const handleConfirmRevoke = () => {
    if (revokeModalAddon?.addonId) {
      onDisconnectAddon({ targetId, addonId: revokeModalAddon.addonId, revoke: true });
      onRevokeModalClose();
      setRevokeModalAddon(null);
    }
  };

  const canConfirmDelete = confirmSteps.step1 && confirmSteps.step2 && confirmSteps.step3;

  const getAddonDescription = (addon: ConnectedAddon) => {
    const addonInfo = availableAddons.find((a) => a.id === addon.addonId);
    if (!addonInfo || !addon.config) return null;

    switch (addonInfo.type) {
      case 'postgresql':
      case 'mysql':
      case 's3':
        return (
          <div className='flex flex-col gap-1 text-sm'>
            <div className='flex items-center gap-2'>
              <Database className='w-4 h-4 text-gray-500' />
              <span>
                База: <strong>{addon.config.name || '-'}</strong>
              </span>
            </div>
            <div className='flex items-center gap-2'>
              <HardDrive className='w-4 h-4 text-gray-500' />
              <span>
                {TargetLocale.Addon.DatabaseSize}: {addon.config.currentSizeMb || 0} из {addon.config.maxSizeMb || 1024}{' '}
                Мб
              </span>
            </div>
          </div>
        );
      case 'kubernetes': {
        const expiredDate = addon.config.expired
          ? new Date(addon.config.expired * 1000).toLocaleDateString('ru-RU')
          : null;
        return (
          <div className='flex flex-col gap-1 text-sm'>
            <div className='flex items-center gap-2'>
              <Server className='w-4 h-4 text-gray-500' />
              <span>
                {TargetLocale.Addon.Namespace}: <strong>{addon.config.name || '-'}</strong>
              </span>
            </div>
            {expiredDate && (
              <div>
                <span>
                  {TargetLocale.Addon.Expires}: <strong>{expiredDate}</strong>
                </span>
              </div>
            )}
          </div>
        );
      }
      case 'harbor':
        return (
          <div className='flex flex-col gap-1 text-sm'>
            <div>
              <span>{TargetLocale.Addon.Registry}: </span>
              <a
                href={addon.config.url}
                target='_blank'
                rel='noopener noreferrer'
                className='text-blue-600 hover:underline'
              >
                {addon.config.url}
              </a>
            </div>
            <div className='text-gray-500'>{TargetLocale.Addon.ApiKeyInVault}</div>
          </div>
        );
      case 'vault':
        return (
          <div className='text-sm'>
            <span>API ключ: </span>
            <a
              href={addon.config.url}
              target='_blank'
              rel='noopener noreferrer'
              className='text-blue-600 hover:underline'
            >
              {addon.config.url}
            </a>
          </div>
        );
      default:
        return null;
    }
  };

  return (
    <div className='flex flex-col gap-4'>
      <div className='flex justify-end'>
        <Button color='primary' onPress={onConnectModalOpen}>
          {TargetLocale.Addon.Connect}
        </Button>
      </div>

      {pendingAddon && (
        <div className='p-4 bg-warning-50 border border-warning-200 rounded-lg'>
          <p className='text-warning-700 font-medium'>{TargetLocale.Addon.DeletePendingMessage}</p>
          <p className='text-sm text-warning-600 mt-1'>
            {TargetLocale.Addon.DeleteRequestedBy}:{' '}
            {pendingUser?.name || pendingUser?.email || pendingAddon.requestDeletedUserId}
          </p>
        </div>
      )}

      {connectedAddons.length === 0 ? (
        <div className='text-center text-gray-400 py-8'>{TargetLocale.Addon.NoAddons}</div>
      ) : (
        <div className='grid gap-3'>
          {connectedAddons.map((addon, idx) => {
            const addonInfo = availableAddons.find((a) => a.id === addon.addonId);
            const isRequestDeleted = (addon.requestDeletedUserId && addon.requestDeletedUserId > 0) as boolean;
            const isRequestDeletedYou = (addon.requestDeletedUserId &&
              addon.requestDeletedUserId > 0 &&
              addon.requestDeletedUserId?.toString() === user.sub) as boolean;
            const displayName = addonInfo?.name || addon.addonId;

            return (
              <Card key={idx} className='border'>
                <CardBody className='flex flex-row items-center justify-between p-4'>
                  <div className='flex-1'>
                    <div className='flex items-center gap-2 mb-1'>
                      <span className='font-medium'>{displayName}</span>
                      <span className='text-sm text-gray-500'>({addon.type})</span>
                    </div>
                    {getAddonDescription(addon)}
                    {isRequestDeleted ? (
                      <span className='text-sm text-gray-500'>{TargetLocale.Addon.DeletePendingMessage}</span>
                    ) : undefined}
                  </div>
                  <div className='ml-4'>
                    <Dropdown>
                      <DropdownTrigger>
                        <Button isIconOnly size='sm' variant='light'>
                          <MoreVertical className='w-4 h-4' />
                        </Button>
                      </DropdownTrigger>
                      <DropdownMenu aria-label='Addon actions'>
                        <DropdownItem
                          key='reset'
                          onPress={() => handleResetRequest(addon)}
                          startContent={<RotateCcw className='w-4 h-4' />}
                        >
                          {TargetLocale.Addon.Reset}
                        </DropdownItem>
                        <DropdownItem
                          key='revokeDelete'
                          isDisabled={!isRequestDeleted}
                          onPress={() => handleRevokeRequest(addon)}
                          startContent={<XCircle className='w-4 h-4' />}
                        >
                          {TargetLocale.Addon.RevokeDeleteRequest}
                        </DropdownItem>
                        <DropdownItem
                          key='delete'
                          isDisabled={isRequestDeletedYou}
                          onPress={() => handleDeleteRequest(addon)}
                          startContent={<Trash2 className='w-4 h-4' />}
                          className='text-danger'
                          color='danger'
                        >
                          {TargetLocale.Addon.Disconnect}
                        </DropdownItem>
                      </DropdownMenu>
                    </Dropdown>
                  </div>
                </CardBody>
              </Card>
            );
          })}
        </div>
      )}

      <Modal isOpen={isConnectModalOpen} onClose={onConnectModalClose} size='lg'>
        <ModalContent>
          <ModalHeader>{TargetLocale.Addon.ConnectTitle}</ModalHeader>
          <ModalBody className='flex flex-col gap-4'>
            <Select
              label={TargetLocale.Addon.SelectAddon}
              placeholder='Выберите дополнение'
              variant='bordered'
              selectedKeys={selectedAddonForConnect ? [selectedAddonForConnect] : []}
              onSelectionChange={(keys) => {
                const selected = Array.from(keys)[0];
                if (selected) setSelectedAddonForConnect(selected as string);
              }}
              renderValue={(items) => {
                return items.map((item) => {
                  const addon = availableAddons.find((a) => a.id === item.key);
                  return (
                    <div key={item.key} className='flex items-center gap-2'>
                      <span>{addon?.name}</span>
                      <span className='text-gray-500'>({addon?.type})</span>
                    </div>
                  );
                });
              }}
            >
              {availableToConnect.map((addon) => (
                <SelectItem key={addon.id} textValue={`${addon.name} (${addon.type})`}>
                  {addon.name} ({addon.type})
                </SelectItem>
              ))}
            </Select>

            {selectedAddonInfo?.configFields?.map((field) => (
              <Input
                key={field.name}
                type={field.type}
                label={field.label}
                placeholder={`Введите ${field.label.toLowerCase()}`}
                variant='bordered'
                value={configValues[field.name] || ''}
                onChange={(e) => setConfigValues((prev) => ({ ...prev, [field.name]: e.target.value }))}
                isRequired={field.required}
              />
            ))}
          </ModalBody>
          <ModalFooter>
            <Button color='danger' variant='light' onPress={onConnectModalClose}>
              {Sidebar.Cancel}
            </Button>
            <Button color='primary' onPress={handleConnectConfirm} isDisabled={!selectedAddonForConnect}>
              {TargetLocale.Addon.Connect}
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>

      <Modal isOpen={isDeleteModalOpen} onClose={onDeleteModalClose}>
        <ModalContent>
          <ModalHeader>{TargetLocale.Addon.DeleteRequest}</ModalHeader>
          <ModalBody className='flex flex-col gap-4'>
            <p className='text-lg font-semibold'>{TargetLocale.Addon.DeleteConfirmWarning}</p>
            <div>
              <Checkbox
                isSelected={confirmSteps.step1}
                onValueChange={(v) => setConfirmSteps((p) => ({ ...p, step1: v }))}
              >
                {TargetLocale.Addon.DeleteConfirmStep1}
              </Checkbox>
              <Checkbox
                isSelected={confirmSteps.step2}
                onValueChange={(v) => setConfirmSteps((p) => ({ ...p, step2: v }))}
              >
                {TargetLocale.Addon.DeleteConfirmStep2}
              </Checkbox>
              <Checkbox
                isSelected={confirmSteps.step3}
                onValueChange={(v) => setConfirmSteps((p) => ({ ...p, step3: v }))}
              >
                {TargetLocale.Addon.DeleteConfirmStep3}
              </Checkbox>
            </div>
          </ModalBody>
          <ModalFooter>
            <Button color='danger' variant='light' onPress={onDeleteModalClose}>
              {TargetLocale.Addon.DeleteCancelButton}
            </Button>
            <Button color='primary' onPress={handleConfirmDelete} isDisabled={!canConfirmDelete}>
              {TargetLocale.Addon.DeleteConfirmButton}
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>

      {onResetAddon && (
        <Modal isOpen={isResetModalOpen} onClose={onResetModalClose}>
          <ModalContent>
            <ModalHeader>{TargetLocale.Addon.ResetTitle}</ModalHeader>
            <ModalBody>
              <p>{TargetLocale.Addon.ResetDescription}</p>
            </ModalBody>
            <ModalFooter>
              <Button color='danger' variant='light' onPress={onResetModalClose}>
                {Sidebar.Cancel}
              </Button>
              <Button color='primary' onPress={handleConfirmReset}>
                {TargetLocale.Addon.Reset}
              </Button>
            </ModalFooter>
          </ModalContent>
        </Modal>
      )}

      <Modal isOpen={isRevokeModalOpen} onClose={onRevokeModalClose}>
        <ModalContent>
          <ModalHeader>{TargetLocale.Addon.RevokeDeleteRequestTitle}</ModalHeader>
          <ModalBody>
            <p>{TargetLocale.Addon.RevokeDeleteRequestDescription}</p>
          </ModalBody>
          <ModalFooter>
            <Button color='danger' variant='light' onPress={onRevokeModalClose}>
              {Sidebar.Cancel}
            </Button>
            <Button color='primary' onPress={handleConfirmRevoke}>
              {TargetLocale.Addon.RevokeDeleteRequest}
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </div>
  );
};
