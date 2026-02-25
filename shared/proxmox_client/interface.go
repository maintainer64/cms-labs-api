package proxmox_client

import "context"

// Interface определяет интерфейс для взаимодействия с Proxmox API.
type Interface interface {
	GetNodes(ctx context.Context) ([]Node, error)
	GetQEMUVMs(ctx context.Context, node string) ([]QEMUVM, error)
	GetVMNetworkInterfaces(ctx context.Context, node string, vmid int) ([]NetworkInterface, error)
}
