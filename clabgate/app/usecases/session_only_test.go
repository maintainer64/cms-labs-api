package usecases

import (
	"testing"

	"github.com/maintainer64/cms-labs-api/shared/cms_client"
	"github.com/maintainer64/cms-labs-api/shared/jsonrpc"
)

func TestTopologyGetRequiresSessionID(t *testing.T) {
	_, err := (&TopologiesGetUC{}).
		SetContext(&cms_client.SSOTokenPublicData{Sub: "42"}).
		Execute(TopologiesGetInputDTO{})
	rpcErr, ok := err.(jsonrpc.RpcError)
	if !ok || rpcErr.Code != "session_required" {
		t.Fatalf("expected session_required, got %#v", err)
	}
}

func TestNodeActionRequiresSessionID(t *testing.T) {
	_, err := (&NodeActionsUC{}).
		SetContext(&cms_client.SSOTokenPublicData{Sub: "42"}).
		Execute(NodeActionInputDTO{Actions: []NodeActionItem{{Node: "r1", Action: "restart"}}})
	if err == nil || err.Error() != "session_id is required" {
		t.Fatalf("expected session_id validation error, got %#v", err)
	}
}
