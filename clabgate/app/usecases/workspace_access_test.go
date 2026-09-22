package usecases

import (
	"strings"
	"testing"
	"time"
)

func TestWorkspaceGrantExchangeAndCookieScope(t *testing.T) {
	secret := strings.Repeat("s", 32)
	sessionID := "550e8400-e29b-41d4-a716-446655440000"
	destination := "/clabgate/workspace/" + sessionID + "/lab?path=Lab.ipynb"
	grant, err := IssueWorkspaceGrant(secret, sessionID, destination, time.Minute)
	if err != nil {
		t.Fatal(err)
	}
	cookie, gotDestination, gotSessionID, err := ExchangeWorkspaceGrant(secret, grant, time.Hour)
	if err != nil {
		t.Fatal(err)
	}
	if gotDestination != destination || gotSessionID != sessionID {
		t.Fatalf("unexpected exchange result: %q %q", gotDestination, gotSessionID)
	}
	if err := VerifyWorkspaceCookie(secret, cookie, destination); err != nil {
		t.Fatalf("valid cookie rejected: %v", err)
	}
	if err := VerifyWorkspaceCookie(secret, cookie, "/clabgate/workspace/aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa/lab"); err == nil {
		t.Fatal("cookie authorized another session")
	}
}

func TestWorkspaceGrantRejectsOpenRedirectAndWeakSecret(t *testing.T) {
	if _, err := IssueWorkspaceGrant("short", "session", "/clabgate/workspace/session/lab", time.Minute); err == nil {
		t.Fatal("weak secret was accepted")
	}
	if _, err := IssueWorkspaceGrant(strings.Repeat("s", 32), "session", "https://example.test/steal", time.Minute); err == nil {
		t.Fatal("absolute redirect was accepted")
	}
}
