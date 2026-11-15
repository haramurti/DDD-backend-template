package service

import "testing"

// TODO: test aturan bisnis & akses, kasus DIIZINKAN dan DITOLAK.
// Contoh table-driven:
//
//	cases := []struct{ name string; actorID uint; wantErr error }{
//		{"owner boleh edit", 1, nil},
//		{"bukan owner ditolak", 2, response.ErrForbidden},
//	}
func TestPlaceholder(t *testing.T) {
	t.Skip("TODO: tulis test aturan bisnis")
}
