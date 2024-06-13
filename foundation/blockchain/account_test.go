package blockchain

import (
	"testing"
)

const netUrl = "http://localhost:8545"

const hardhatUsedAddr = "0xdf57089febbacf7ba0bc227dafbffa9fc08a93fdc68e1e42411a14efcf23656e"

//Tests are made to run in hardhat net

func Test_isAvailable(t *testing.T) {
	type args struct {
		addressHex string
		netUrl     string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "Address with no code and no transactions",
			args: args{
				addressHex: "0xUnusedAddressHere",
				netUrl:     netUrl,
			},
			want:    true,
			wantErr: false,
		},
		/* 		{
		   			name: "Address with contract code",
		   			args: args{
		   				addressHex: hardhatUsedAddr, // Replace with an address known to have contract code
		   				netUrl:     netUrl,
		   			},
		   			want:    false,
		   			wantErr: false,
		   		},
		   		{
		   			name: "Address with transactions",
		   			args: args{
		   				addressHex: hardhatUsedAddr,
		   				netUrl:     netUrl,
		   			},
		   			want:    false,
		   			wantErr: false,
		   		}, */
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isAvailable(tt.args.addressHex, tt.args.netUrl)
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
	type args struct {
		netUrl string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Successful account creation",
			args: args{
				netUrl: netUrl,
			},
			wantErr: false,
		},
		/* 		{
			name: "Unsuccessful account creation due to address being unavailable",
			args: args{
				netUrl: netUrl,
			},
			wantErr: true,
		}, */
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateEthAccount(tt.args.netUrl)
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
