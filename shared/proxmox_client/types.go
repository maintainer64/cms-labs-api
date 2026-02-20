package proxmox_client

import (
	"github.com/Telmate/proxmox-api-go/proxmox"
)

// Node представляет узел Proxmox (физический сервер).
type Node struct {
	Node    string
	MaxCPU  int
	MaxMem  int64 // байты
	MaxDisk int64 // байты
}

// QEMUVM представляет виртуальную машину QEMU.
type QEMUVM struct {
	VMID    int
	Name    string
	CPUs    int
	MaxMem  int64    // байты
	MaxDisk int64    // байты
	Tags    []string // теги как срез строк
}

// NetworkInterface представляет сетевой интерфейс ВМ (от гостевого агента).
type NetworkInterface struct {
	Name        string
	IPAddresses []IPAddress
}

type IPAddress struct {
	IPAddressType string // "ipv4" или "ipv6"
	IPAddress     string
}

// managedLinkTypes определяет типы ссылок, управляемых процессом синхронизации.
var managedLinkTypes = map[string]bool{
	"cpu":  true,
	"ram":  true,
	"disk": true,
	"ip":   true,
}

func IsManagedLinkType(typ string) bool {
	return managedLinkTypes[typ]
}

// ProxmoxAPI адаптирует *proxmox.Client к интерфейсу ProxmoxClient.
type ProxmoxAPI struct {
	*proxmox.Client
}
