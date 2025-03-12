package models

import (
	"fmt"
	"strconv"
	"strings"
)

// Wireshark модель для таблицы wiresharks
type Wireshark struct {
	WsID       int64  `gorm:"primary_key;column:ws_id;AUTO_INCREMENT"`
	WsTenant   int    `gorm:"column:ws_tenant"`
	WsLab      string `gorm:"column:ws_lab;size:200"`
	WsNode     int    `gorm:"column:ws_node"`
	WsIf       int    `gorm:"column:ws_if"`
	WsNet      int    `gorm:"column:ws_net"`
	WsNodeName string `gorm:"column:ws_node_name;size:150"`
	WsIfName   string `gorm:"column:ws_if_name;size:150"`
	WsDcName   string `gorm:"column:ws_dc_name;size:150"`
	WsPort     int    `gorm:"column:ws_port"`
	WsIP       string `gorm:"column:ws_ip;size:150"`
}

// TableName возвращает имя таблицы для модели Wireshark
func (Wireshark) TableName() string {
	return "wiresharks"
}

// User модель для таблицы users
type User struct {
	Pod           int     `gorm:"primary_key;column:pod;AUTO_INCREMENT"`
	Username      string  `gorm:"column:username;type:text"`
	Cookie        string  `gorm:"column:cookie;type:text"`
	Email         string  `gorm:"column:email;size:150;unique"`
	Expiration    int     `gorm:"column:expiration;default:-1"`
	Name          string  `gorm:"column:name;type:text"`
	Password      string  `gorm:"column:password;type:text"`
	Session       *int64  `gorm:"column:session"`
	IP            *string `gorm:"column:ip;type:text"`
	Role          string  `gorm:"column:role;type:text"`
	Folder        *string `gorm:"column:folder;type:text"`
	LabSession    *int    `gorm:"column:lab_session"`
	HTML5         bool    `gorm:"column:html5"`
	License       *string `gorm:"column:license;type:text"`
	OnlineTime    int64   `gorm:"column:online_time"`
	Note          string  `gorm:"column:note;type:text"`
	Offline       *int    `gorm:"column:offline"`
	ActiveTime    *int    `gorm:"column:active_time"`
	ExpiredTime   *int    `gorm:"column:expired_time"`
	UserStatus    int     `gorm:"column:user_status;default:1"`
	UserWorkspace *string `gorm:"column:user_workspace;type:text"`
	MaxNode       *int    `gorm:"column:max_node"`
	MaxNodeLab    *int    `gorm:"column:max_node_lab"`
}

// TableName возвращает имя таблицы для модели User
func (User) TableName() string {
	return "users"
}

// UserRole модель для таблицы user_roles
type UserRole struct {
	UserRoleID        int     `gorm:"primary_key;column:user_role_id;AUTO_INCREMENT"`
	UserRoleName      string  `gorm:"column:user_role_name;size:150"`
	UserRoleWorkspace string  `gorm:"column:user_role_workspace;type:text"`
	UserRoleNote      string  `gorm:"column:user_role_note;type:text"`
	UserRoleRAM       float64 `gorm:"column:user_role_ram"`
	UserRoleCPU       float64 `gorm:"column:user_role_cpu"`
	UserRoleHDD       float64 `gorm:"column:user_role_hdd"`
}

// TableName возвращает имя таблицы для модели UserRole
func (UserRole) TableName() string {
	return "user_roles"
}

// UserPermission модель для таблицы user_permission
type UserPermission struct {
	UserPerID   int    `gorm:"primary_key;column:user_per_id;AUTO_INCREMENT"`
	UserPerRole int    `gorm:"column:user_per_role"`
	UserPerName string `gorm:"column:user_per_name;size:150"`
}

// TableName возвращает имя таблицы для модели UserPermission
func (UserPermission) TableName() string {
	return "user_permission"
}

// ProcessDevice модель для таблицы process_device
type ProcessDevice struct {
	ProcessDeviceID     string `gorm:"primary_key;column:process_device_id;size:150"`
	ProcessDeviceDTotal int    `gorm:"column:process_device_dtotal"`
	ProcessDeviceDNow   int    `gorm:"column:process_device_dnow"`
	ProcessDeviceUTotal int    `gorm:"column:process_device_utotal"`
	ProcessDeviceUNow   int    `gorm:"column:process_device_unow"`
	ProcessDeviceLog    string `gorm:"column:process_device_log;type:text"`
}

// TableName возвращает имя таблицы для модели ProcessDevice
func (ProcessDevice) TableName() string {
	return "process_device"
}

// Process модель для таблицы process
type Process struct {
	ProcessID     string `gorm:"primary_key;column:process_id;size:200"`
	ProcessDTotal int    `gorm:"column:process_dtotal"`
	ProcessDNow   int    `gorm:"column:process_dnow"`
	ProcessUTotal int    `gorm:"column:process_utotal"`
	ProcessUNow   int    `gorm:"column:process_unow"`
	ProcessFinish int    `gorm:"column:process_finish"`
}

// TableName возвращает имя таблицы для модели Process
func (Process) TableName() string {
	return "process"
}

// NodeSession модель для таблицы node_sessions
type NodeSession struct {
	NodeSessionID        int     `gorm:"primary_key;column:node_session_id"`
	NodeSessionNID       int     `gorm:"column:node_session_nid"`
	NodeSessionLab       int     `gorm:"column:node_session_lab"`
	NodeSessionPort      int     `gorm:"column:node_session_port"`
	NodeSessionType      string  `gorm:"column:node_session_type;size:150"`
	NodeSessionWorkspace string  `gorm:"column:node_session_workspace;type:text"`
	NodeSessionRAM       float64 `gorm:"column:node_session_ram"`
	NodeSessionCPU       float64 `gorm:"column:node_session_cpu"`
	NodeSessionHDD       float64 `gorm:"column:node_session_hdd"`
	NodeSessionRunning   int     `gorm:"column:node_session_running"`
	NodeSessionPod       int     `gorm:"column:node_session_pod"`
	NodeSessionIOL       int     `gorm:"column:node_session_iol"`
}

// TableName возвращает имя таблицы для модели NodeSession
func (NodeSession) TableName() string {
	return "node_sessions"
}

// LabSession модель для таблицы lab_sessions
type LabSession struct {
	LabSessionID      int    `gorm:"primary_key;column:lab_session_id;AUTO_INCREMENT"`
	LabSessionLID     string `gorm:"column:lab_session_lid;size:150"`
	LabSessionPod     int    `gorm:"column:lab_session_pod"`
	LabSessionJoined  string `gorm:"column:lab_session_joined;type:text"`
	LabSessionPath    string `gorm:"column:lab_session_path;type:text"`
	LabSessionRunning int    `gorm:"column:lab_session_running"`
}

func (l *LabSession) AddJoinedUser(userPod int) {
	strNumbers := strings.Split(l.LabSessionJoined, ",")
	// Create a set of integers
	set := make(map[int]bool)
	for _, strNum := range strNumbers {
		num, err := strconv.Atoi(strNum)
		if err != nil {
			continue
		}
		set[num] = true
	}
	set[userPod] = true

	// Convert the set back to a comma-separated string
	var result []string
	for num := range set {
		result = append(result, fmt.Sprintf("%d", num))
	}
	l.LabSessionJoined = strings.Join(result, ",")
}

// TableName возвращает имя таблицы для модели LabSession
func (LabSession) TableName() string {
	return "lab_sessions"
}

// IfSession модель для таблицы if_sessions
type IfSession struct {
	IfSessionID      int64  `gorm:"primary_key;column:if_session_id;AUTO_INCREMENT"`
	IfSessionLab     int    `gorm:"column:if_session_lab"`
	IfSessionNode    int    `gorm:"column:if_session_node"`
	IfSessionIfID    int    `gorm:"column:if_session_ifid"`
	IfSessionType    string `gorm:"column:if_session_type;size:150"`
	IfSessionQuality string `gorm:"column:if_session_quality;type:text"`
	IfSessionSuspend int    `gorm:"column:if_session_suspend"`
}

// TableName возвращает имя таблицы для модели IfSession
func (IfSession) TableName() string {
	return "if_sessions"
}

// HTML5 модель для таблицы html5
type HTML5 struct {
	Username string `gorm:"column:username;type:text"`
	Pod      int    `gorm:"column:pod"`
	Token    string `gorm:"column:token;type:text"`
}

// TableName возвращает имя таблицы для модели HTML5
func (HTML5) TableName() string {
	return "html5"
}
