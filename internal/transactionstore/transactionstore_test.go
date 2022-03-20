package transactionstore

import (
	"testing"
	"time"
)

func TestCreateAndGet(t *testing.T) {
	// Create and store a single Transaction
	ts := New()
	req := TransactionCreateRequest{
		Payee:  "Payee 1",
		Amount: 54.31,
		Date:   time.Now(),
	}
	id := ts.CreateTransaction(req)

	// We should be able to retrieve this trans by ID, but nothing with other
	// IDs.
	trans, err := ts.GetTransaction(id)
	if err != nil {
		t.Fatal(err)
	}

	if trans.ID != id {
		t.Errorf("got trans id=%d, expecting id=%d", trans.ID, id)
	}
	if trans.Payee != req.Payee {
		t.Errorf("got Payee=%v, want %v", trans.Payee, req.Payee)
	}
}
