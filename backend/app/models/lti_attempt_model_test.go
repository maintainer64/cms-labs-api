package models

import (
	"encoding/json"
	"testing"
)

func TestLTIAttemptResultPreservesStructuredCheckerTasks(t *testing.T) {
	payload := []byte(`{
  "max_score": 2,
  "current_score": 1,
  "result_display": "1/2 checks passed",
  "tasks": [{
    "title": "SSH",
    "description": "router accepts SSH",
    "logs": [{"node": "r1", "namespace": "lab-one", "message": "connected"}],
    "complete": true
  }]
}`)

	result := LTIAttemptResult{}
	if err := json.Unmarshal(payload, &result); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}
	if len(result.Tasks) != 1 || len(result.Tasks[0].Logs) != 1 {
		t.Fatalf("structured checker tasks were lost: %+v", result.Tasks)
	}
	if result.Tasks[0].Logs[0].Namespace != "lab-one" {
		t.Fatalf("unexpected structured log: %+v", result.Tasks[0].Logs[0])
	}
}
