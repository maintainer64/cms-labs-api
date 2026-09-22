package usecases

import "testing"

func TestDecodeCheckerGradeKeepsStructuredTasks(t *testing.T) {
	payload := `{
  "max_score": 2,
  "current_score": 1,
  "result_display": "1/2 checks passed",
  "report": "student report",
  "check_id": "untrusted-check-id",
  "logs": "untrusted logs",
  "tasks": [{
    "title": "SSH",
    "description": "router accepts SSH connections",
    "logs": [{"node": "r1", "namespace": "lab-1", "message": "connected"}],
    "complete": true
  }]
}`

	grade, err := decodeCheckerGrade(payload, "server-check-id", "pod stdout")
	if err != nil {
		t.Fatalf("decodeCheckerGrade() error = %v", err)
	}
	if grade.CheckID != "server-check-id" || grade.Logs != "pod stdout" {
		t.Fatalf("server fields were not enforced: %+v", grade)
	}
	if len(grade.Tasks) != 1 || len(grade.Tasks[0].Logs) != 1 {
		t.Fatalf("structured tasks were lost: %+v", grade.Tasks)
	}
	if grade.Tasks[0].Logs[0].Message != "connected" {
		t.Fatalf("unexpected structured log: %+v", grade.Tasks[0].Logs[0])
	}
}

func TestDecodeCheckerGradeRejectsInvalidScore(t *testing.T) {
	_, err := decodeCheckerGrade(`{"max_score":1,"current_score":2,"result_display":"invalid"}`, "check", "logs")
	if err == nil {
		t.Fatal("decodeCheckerGrade() accepted a score above max_score")
	}
}

func TestDecodeCheckerGradeSupportsLegacyResultWithoutTasks(t *testing.T) {
	grade, err := decodeCheckerGrade(`{"max_score":10,"current_score":0,"result_display":"0/10"}`, "check", "")
	if err != nil {
		t.Fatalf("decodeCheckerGrade() error = %v", err)
	}
	if grade.Tasks != nil {
		t.Fatalf("legacy result unexpectedly gained tasks: %+v", grade.Tasks)
	}
}
