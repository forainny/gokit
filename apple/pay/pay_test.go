package pay

import "testing"

func TestGetRecentOrder(t *testing.T) {
	pay := NewPay("https://sandbox.itunes.apple.com/verifyReceipt", "bundleID")
	iap, err := pay.GetRecentOrder("")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(iap)
}
