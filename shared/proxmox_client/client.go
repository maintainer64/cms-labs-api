package proxmox_client

import (
	"context"
	"errors"
	"net"
	"strings"

	"github.com/Telmate/proxmox-api-go/proxmox"
)

// GetNodes получает список узлов кластера и преобразует их в наш тип Node.
func (p *ProxmoxAPI) GetNodes(ctx context.Context) ([]Node, error) {
	// Получаем список узлов. Возвращаемый тип — map[string]interface{}, где в "data" лежит []interface{}.
	nodesMap, err := p.Client.GetNodeList(ctx)
	if err != nil {
		return nil, err
	}
	data, ok := nodesMap["data"].([]interface{})
	if !ok {
		return nil, errors.New("неожиданный формат ответа от GetNodeList")
	}

	var nodes []Node
	for _, item := range data {
		nodeData, ok := item.(map[string]interface{})
		if !ok {
			continue
		}

		node := Node{
			Node: nodeData["node"].(string),
		}
		// Безопасно извлекаем числовые значения
		if v, ok := nodeData["maxcpu"]; ok {
			if f, ok := v.(float64); ok {
				node.MaxCPU = int(f)
			}
		}
		if v, ok := nodeData["maxmem"]; ok {
			if f, ok := v.(float64); ok {
				node.MaxMem = int64(f)
			}
		}
		if v, ok := nodeData["maxdisk"]; ok {
			if f, ok := v.(float64); ok {
				node.MaxDisk = int64(f)
			}
		}
		nodes = append(nodes, node)
	}
	return nodes, nil
}

// GetQEMUVMs получает все QEMU-виртуальные машины на указанном узле.
// Теперь мы получаем полную конфигурацию каждой ВМ, чтобы извлечь теги.
func (p *ProxmoxAPI) GetQEMUVMs(ctx context.Context, node string) ([]QEMUVM, error) {
	// Сначала получаем список всех ВМ в ресурсах кластера для быстрого доступа к базовой информации.
	resources, err := p.Client.GetResourceList(ctx, "vm")
	if err != nil {
		return nil, err
	}

	var vms []QEMUVM
	for _, res := range resources {
		vmMap, ok := res.(map[string]interface{})
		if !ok {
			continue
		}
		// Фильтруем по узлу и типу "qemu"
		if nodeVal, ok := vmMap["node"]; !ok || nodeVal != node {
			continue
		}
		if typeVal, ok := vmMap["type"]; !ok || typeVal != "qemu" {
			continue
		}

		vm := QEMUVM{
			VMID: int(vmMap["vmid"].(float64)),
		}
		if name, ok := vmMap["name"]; ok && name != nil {
			vm.Name = name.(string)
		}
		if v, ok := vmMap["maxcpu"]; ok {
			if f, ok := v.(float64); ok {
				vm.CPUs = int(f)
			}
		}
		if v, ok := vmMap["maxmem"]; ok {
			if f, ok := v.(float64); ok {
				vm.MaxMem = int64(f)
			}
		}
		if v, ok := vmMap["maxdisk"]; ok {
			if f, ok := v.(float64); ok {
				vm.MaxDisk = int64(f)
			}
		}
		if v, ok := vmMap["tags"]; ok {
			if f, ok := v.(string); ok {
				vm.Tags = strings.Split(f, ";")
			}
		}
		vms = append(vms, vm)
	}
	return vms, nil
}

func getTypeIp(ip net.IP) string {
	if ip == nil {
		return "unknown"
	}
	if ip.To4() != nil {
		return "ipv4"
	}
	return "ipv6"
}

// GetVMNetworkInterfaces запрашивает сетевые интерфейсы ВМ через гостевой агент.
func (p *ProxmoxAPI) GetVMNetworkInterfaces(ctx context.Context, node string, vmid int) ([]NetworkInterface, error) {
	vmr := proxmox.NewVmRef(proxmox.GuestID(vmid))
	vmr.SetNode(node)
	vmr.SetVmType(proxmox.GuestQemu)

	agentInfoRaw, _, err := vmr.GetAgentInformation(ctx, p.Client)
	if err != nil {
		return nil, err
	}
	agentInfo := agentInfoRaw.Get()

	var ifaces []NetworkInterface
	for _, iface := range agentInfo {
		var ips []IPAddress
		for _, addr := range iface.IpAddresses {
			ips = append(ips, IPAddress{
				IPAddressType: getTypeIp(addr),
				IPAddress:     addr.String(),
			})
		}
		ifaces = append(ifaces, NetworkInterface{
			Name:        iface.Name,
			IPAddresses: ips,
		})
	}
	return ifaces, nil
}

func parseTokenString(s string) (token proxmox.ApiTokenID, secret proxmox.ApiTokenSecret, err error) {
	// 1. Делим по '@'
	beforeAt, afterAt, found := strings.Cut(s, "@")
	if !found {
		return token, secret, errors.New("missing '@' separator")
	}
	token.User.Name = beforeAt

	// 2. Делим оставшуюся часть по '!'
	beforeBang, afterBang, found := strings.Cut(afterAt, "!")
	if !found {
		return token, secret, errors.New("missing '!' separator")
	}
	token.User.Realm = beforeBang

	// 3. Делим оставшуюся часть по '='
	name, secretValue, found := strings.Cut(afterBang, "=")
	if !found {
		return token, secret, errors.New("missing '=' separator")
	}
	token.TokenName = proxmox.ApiTokenName(name)
	secret = proxmox.ApiTokenSecret(secretValue)
	return token, secret, nil
}

func NewProxmoxAPI(debug bool, url string, token string) (*ProxmoxAPI, error) {
	client, err := proxmox.NewClient(
		url,
		nil,
		"",
		nil,
		"",
		0,
		debug,
	)
	if err != nil {
		return nil, err
	}
	tokenAPI, secretAPI, err := parseTokenString(token)
	if err != nil {
		return nil, err
	}
	client.SetAPIToken(tokenAPI, secretAPI)
	return &ProxmoxAPI{client}, nil
}
