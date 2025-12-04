package tests

import (
	"testing"

	"github.com/chessclub2205-dev/versus-service/internal/payments"
)

func TestPayouts(t *testing.T) {
	stake := int64(6250) // one slice = 6250 kobo = 62.50 Naira
	w, l, c, total, err := payments.Payouts(stake)
	if err != nil {
		t.Fatal(err)
	}
	if total != stake*2 {
		t.Fatalf("total mismatch: %d", total)
	}
	if w != stake+(stake*30)/100 {
		t.Fatalf("winner mismatch: %d", w)
	}
	if l != (stake*50)/100 {
		t.Fatalf("loser mismatch: %d", l)
	}
	if (w + l + c) != total {
		t.Fatalf("sum mismatch: %d != %d", w+l+c, total)
	}
}
