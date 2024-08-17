package blockchain

import (
	"testing"
)

const netUrl = "http://localhost:8545"

const hardhatUsedAddr = "0xdf57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e"

// Tests are made to run in hardhat net

func Test_isAvailable(t *testing.T) {
	client, err := NewClient(netUrl)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	tests := []struct {
		name    string
		address string
		want    bool
		wantErr bool
	}{
		{
			name:    "Address with no code and no transactions",
			address: "0xUnusedAddressHere",
			want:    true,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.isAvailable(tt.address)
			if (err != nil) != tt.wantErr {
				t.Errorf("isAvailable() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isAvailable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateEthAccount(t *testing.T) {
	client, err := NewClient(netUrl)
	if err != nil {
		t.Fatalf("Failed to create client: %v", err)
	}

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Successful account creation",
			wantErr: false,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := client.CreateEthAccount()
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateEthAccount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				// Check if the returned AccountData is valid
				if len(got.AddressHex) == 0 || len(got.PrivateKey) == 0 || len(got.PublicKey) == 0 {
					t.Errorf("CreateEthAccount() returned invalid account data: %v", got)
				}
			}
		})
	}
}
