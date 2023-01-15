package security

import (
	"testing"
)

const testPassword = "i_am_a_te$t_pa55wórd"

func TestHashPassword(t *testing.T) {
	hashedTestPassword := HashPassword(testPassword)
	if !PasswordMatchesHash(testPassword, hashedTestPassword) {
		t.Fatalf("Test password should match hashed test password")
	}

	wrongPassword := "wr0ng_p4ssw0rd"
	if PasswordMatchesHash(wrongPassword, hashedTestPassword) {
		t.Fatalf("Wrong password should not match hashed test password")
	}
}
