package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := Product{Name: "capuccino", Price: 10, SKU: "abs-sd-dsfs"}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}

}
