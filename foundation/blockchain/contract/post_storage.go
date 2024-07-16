// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// PostStoragePost is an auto generated low-level Go binding around an user-defined struct.
type PostStoragePost struct {
	Author    common.Address
	Title     string
	Content   string
	Timestamp *big.Int
	Category  uint8
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumPostStorage.Category\",\"name\":\"category\",\"type\":\"uint8\"}],\"name\":\"PostPublished\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getAllPosts\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumPostStorage.Category\",\"name\":\"category\",\"type\":\"uint8\"}],\"internalType\":\"structPostStorage.Post[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllUsers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPostByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumPostStorage.Category\",\"name\":\"category\",\"type\":\"uint8\"}],\"internalType\":\"structPostStorage.Post\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUsersPost\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumPostStorage.Category\",\"name\":\"category\",\"type\":\"uint8\"}],\"internalType\":\"structPostStorage.Post[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"},{\"internalType\":\"enumPostStorage.Category\",\"name\":\"_category\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"}],\"name\":\"post\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userPosts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumPostStorage.Category\",\"name\":\"category\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611cc2806100206000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80637c7bf4df1161005b5780637c7bf4df146101005780638064d14914610130578063e2842d7914610164578063e8b4edb4146101825761007d565b80633056110014610082578063365b98b2146100a05780633aa5cc94146100d0575b600080fd5b61008a61019e565b6040516100979190611865565b60405180910390f35b6100ba60048036038101906100b59190611499565b6106df565b6040516100c791906117c7565b60405180910390f35b6100ea60048036038101906100e591906113a1565b61071e565b6040516100f79190611865565b60405180910390f35b61011a600480360381019061011591906113ca565b6109c6565b60405161012791906118d7565b60405180910390f35b61014a600480360381019061014591906113ca565b610cf3565b60405161015b9594939291906117e2565b60405180910390f35b61016c610e83565b6040516101799190611843565b60405180910390f35b61019c60048036038101906101979190611406565b610f11565b005b60606000805b60018054905081101561027757600080600183815481106101ee577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490508261026291906119f9565b9150808061026f90611b24565b9150506101a4565b5060008167ffffffffffffffff8111156102ba577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6040519080825280602002602001820160405280156102f357816020015b6102e06111da565b8152602001906001900390816102d85790505b5090506000805b6001805490508110156106d557600080600060018481548110610346577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b8282101561060857838290600052602060002090600502016040518060a00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461045490611af2565b80601f016020809104026020016040519081016040528092919081815260200182805461048090611af2565b80156104cd5780601f106104a2576101008083540402835291602001916104cd565b820191906000526020600020905b8154815290600101906020018083116104b057829003601f168201915b505050505081526020016002820180546104e690611af2565b80601f016020809104026020016040519081016040528092919081815260200182805461051290611af2565b801561055f5780601f106105345761010080835404028352916020019161055f565b820191906000526020600020905b81548152906001019060200180831161054257829003601f168201915b50505050508152602001600382015481526020016004820160009054906101000a900460ff1660028111156105bd577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60028111156105f5577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81525050815260200190600101906103cb565b50505050905060005b81518110156106c057818181518110610653577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151858581518110610694577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b602002602001018190525083806106aa90611b24565b94505080806106b890611b24565b915050610611565b505080806106cd90611b24565b9150506102fa565b5081935050505090565b600181815481106106ef57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156109bb57838290600052602060002090600502016040518060a00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461080790611af2565b80601f016020809104026020016040519081016040528092919081815260200182805461083390611af2565b80156108805780601f1061085557610100808354040283529160200191610880565b820191906000526020600020905b81548152906001019060200180831161086357829003601f168201915b5050505050815260200160028201805461089990611af2565b80601f01602080910402602001604051908101604052809291908181526020018280546108c590611af2565b80156109125780601f106108e757610100808354040283529160200191610912565b820191906000526020600020905b8154815290600101906020018083116108f557829003601f168201915b50505050508152602001600382015481526020016004820160009054906101000a900460ff166002811115610970577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60028111156109a8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815250508152602001906001019061077e565b505050509050919050565b6109ce6111da565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490508210610a51576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a48906118b7565b60405180910390fd5b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208281548110610ac7577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002090600502016040518060a00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600182018054610b4690611af2565b80601f0160208091040260200160405190810160405280929190818152602001828054610b7290611af2565b8015610bbf5780601f10610b9457610100808354040283529160200191610bbf565b820191906000526020600020905b815481529060010190602001808311610ba257829003601f168201915b50505050508152602001600282018054610bd890611af2565b80601f0160208091040260200160405190810160405280929190818152602001828054610c0490611af2565b8015610c515780601f10610c2657610100808354040283529160200191610c51565b820191906000526020600020905b815481529060010190602001808311610c3457829003601f168201915b50505050508152602001600382015481526020016004820160009054906101000a900460ff166002811115610caf577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115610ce7577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81525050905092915050565b60006020528160005260406000208181548110610d0f57600080fd5b9060005260206000209060050201600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001018054610d5990611af2565b80601f0160208091040260200160405190810160405280929190818152602001828054610d8590611af2565b8015610dd25780601f10610da757610100808354040283529160200191610dd2565b820191906000526020600020905b815481529060010190602001808311610db557829003601f168201915b505050505090806002018054610de790611af2565b80601f0160208091040260200160405190810160405280929190818152602001828054610e1390611af2565b8015610e605780601f10610e3557610100808354040283529160200191610e60565b820191906000526020600020905b815481529060010190602001808311610e4357829003601f168201915b5050505050908060030154908060040160009054906101000a900460ff16905085565b60606001805480602002602001604051908101604052809291908181526020018280548015610f0757602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610ebd575b5050505050905090565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490501415610fbf576001819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b60006040518060a001604052808373ffffffffffffffffffffffffffffffffffffffff168152602001868152602001858152602001428152602001846002811115611033577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81525090506000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081908060018154018082558091505060019003906000526020600020906005020160009091909190915060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506020820151816001019080519060200190611103929190611257565b506040820151816002019080519060200190611120929190611257565b506060820151816003015560808201518160040160006101000a81548160ff0219169083600281111561117c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b021790555050508173ffffffffffffffffffffffffffffffffffffffff167f4be77e23b4d9d657bb35113acb2ddeb9fabadff4c213859965d0b6d9b385738186856040516111cb929190611887565b60405180910390a25050505050565b6040518060a00160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001606081526020016000815260200160006002811115611251577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81525090565b82805461126390611af2565b90600052602060002090601f01602090048101928261128557600085556112cc565b82601f1061129e57805160ff19168380011785556112cc565b828001600101855582156112cc579182015b828111156112cb5782518255916020019190600101906112b0565b5b5090506112d991906112dd565b5090565b5b808211156112f65760008160009055506001016112de565b5090565b600061130d6113088461192a565b6118f9565b90508281526020810184848401111561132557600080fd5b611330848285611ab0565b509392505050565b60008135905061134781611c4e565b92915050565b60008135905061135c81611c65565b92915050565b600082601f83011261137357600080fd5b81356113838482602086016112fa565b91505092915050565b60008135905061139b81611c75565b92915050565b6000602082840312156113b357600080fd5b60006113c184828501611338565b91505092915050565b600080604083850312156113dd57600080fd5b60006113eb85828601611338565b92505060206113fc8582860161138c565b9150509250929050565b6000806000806080858703121561141c57600080fd5b600085013567ffffffffffffffff81111561143657600080fd5b61144287828801611362565b945050602085013567ffffffffffffffff81111561145f57600080fd5b61146b87828801611362565b935050604061147c8782880161134d565b925050606061148d87828801611338565b91505092959194509250565b6000602082840312156114ab57600080fd5b60006114b98482850161138c565b91505092915050565b60006114ce83836114ee565b60208301905092915050565b60006114e683836116af565b905092915050565b6114f781611a4f565b82525050565b61150681611a4f565b82525050565b60006115178261197a565b61152181856119b5565b935061152c8361195a565b8060005b8381101561155d57815161154488826114c2565b975061154f8361199b565b925050600181019050611530565b5085935050505092915050565b600061157582611985565b61157f81856119c6565b9350836020820285016115918561196a565b8060005b858110156115cd57848403895281516115ae85826114da565b94506115b9836119a8565b925060208a01995050600181019050611595565b50829750879550505050505092915050565b6115e881611a9e565b82525050565b6115f781611a9e565b82525050565b600061160882611990565b61161281856119d7565b9350611622818560208601611abf565b61162b81611c29565b840191505092915050565b600061164182611990565b61164b81856119e8565b935061165b818560208601611abf565b61166481611c29565b840191505092915050565b600061167c6013836119e8565b91507f506f737420646f6573206e6f74206578697374000000000000000000000000006000830152602082019050919050565b600060a0830160008301516116c760008601826114ee565b50602083015184820360208601526116df82826115fd565b915050604083015184820360408601526116f982826115fd565b915050606083015161170e60608601826117a9565b50608083015161172160808601826115df565b508091505092915050565b600060a08301600083015161174460008601826114ee565b506020830151848203602086015261175c82826115fd565b9150506040830151848203604086015261177682826115fd565b915050606083015161178b60608601826117a9565b50608083015161179e60808601826115df565b508091505092915050565b6117b281611a94565b82525050565b6117c181611a94565b82525050565b60006020820190506117dc60008301846114fd565b92915050565b600060a0820190506117f760008301886114fd565b81810360208301526118098187611636565b9050818103604083015261181d8186611636565b905061182c60608301856117b8565b61183960808301846115ee565b9695505050505050565b6000602082019050818103600083015261185d818461150c565b905092915050565b6000602082019050818103600083015261187f818461156a565b905092915050565b600060408201905081810360008301526118a18185611636565b90506118b060208301846115ee565b9392505050565b600060208201905081810360008301526118d08161166f565b9050919050565b600060208201905081810360008301526118f1818461172c565b905092915050565b6000604051905081810181811067ffffffffffffffff821117156119205761191f611bfa565b5b8060405250919050565b600067ffffffffffffffff82111561194557611944611bfa565b5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611a0482611a94565b9150611a0f83611a94565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611a4457611a43611b6d565b5b828201905092915050565b6000611a5a82611a74565b9050919050565b6000819050611a6f82611c3a565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611aa982611a61565b9050919050565b82818337600083830152505050565b60005b83811015611add578082015181840152602081019050611ac2565b83811115611aec576000848401525b50505050565b60006002820490506001821680611b0a57607f821691505b60208210811415611b1e57611b1d611bcb565b5b50919050565b6000611b2f82611a94565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611b6257611b61611b6d565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b60038110611c4b57611c4a611b9c565b5b50565b611c5781611a4f565b8114611c6257600080fd5b50565b60038110611c7257600080fd5b50565b611c7e81611a94565b8114611c8957600080fd5b5056fea2646970667358221220affa21724124c9a383d1d89278f8d327c438cfd153325fb23047e79bb1056d0564736f6c63430008000033",
}

// ContractABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractMetaData.ABI instead.
var ContractABI = ContractMetaData.ABI

// ContractBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractMetaData.Bin instead.
var ContractBin = ContractMetaData.Bin

// DeployContract deploys a new Ethereum contract, binding an instance of Contract to it.
func DeployContract(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// Contract is an auto generated Go binding around an Ethereum contract.
type Contract struct {
	ContractCaller     // Read-only binding to the contract
	ContractTransactor // Write-only binding to the contract
	ContractFilterer   // Log filterer for contract events
}

// ContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractSession struct {
	Contract     *Contract         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractCallerSession struct {
	Contract *ContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractTransactorSession struct {
	Contract     *ContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractRaw struct {
	Contract *Contract // Generic contract binding to access the raw methods on
}

// ContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractCallerRaw struct {
	Contract *ContractCaller // Generic read-only contract binding to access the raw methods on
}

// ContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractTransactorRaw struct {
	Contract *ContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContract creates a new instance of Contract, bound to a specific deployed contract.
func NewContract(address common.Address, backend bind.ContractBackend) (*Contract, error) {
	contract, err := bindContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contract{ContractCaller: ContractCaller{contract: contract}, ContractTransactor: ContractTransactor{contract: contract}, ContractFilterer: ContractFilterer{contract: contract}}, nil
}

// NewContractCaller creates a new read-only instance of Contract, bound to a specific deployed contract.
func NewContractCaller(address common.Address, caller bind.ContractCaller) (*ContractCaller, error) {
	contract, err := bindContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractCaller{contract: contract}, nil
}

// NewContractTransactor creates a new write-only instance of Contract, bound to a specific deployed contract.
func NewContractTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractTransactor, error) {
	contract, err := bindContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractTransactor{contract: contract}, nil
}

// NewContractFilterer creates a new log filterer instance of Contract, bound to a specific deployed contract.
func NewContractFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractFilterer, error) {
	contract, err := bindContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractFilterer{contract: contract}, nil
}

// bindContract binds a generic wrapper to an already deployed contract.
func bindContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.ContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.ContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contract *ContractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contract *ContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contract *ContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contract.Contract.contract.Transact(opts, method, params...)
}

// GetAllPosts is a free data retrieval call binding the contract method 0x30561100.
//
// Solidity: function getAllPosts() view returns((address,string,string,uint256,uint8)[])
func (_Contract *ContractCaller) GetAllPosts(opts *bind.CallOpts) ([]PostStoragePost, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAllPosts")

	if err != nil {
		return *new([]PostStoragePost), err
	}

	out0 := *abi.ConvertType(out[0], new([]PostStoragePost)).(*[]PostStoragePost)

	return out0, err

}

// GetAllPosts is a free data retrieval call binding the contract method 0x30561100.
//
// Solidity: function getAllPosts() view returns((address,string,string,uint256,uint8)[])
func (_Contract *ContractSession) GetAllPosts() ([]PostStoragePost, error) {
	return _Contract.Contract.GetAllPosts(&_Contract.CallOpts)
}

// GetAllPosts is a free data retrieval call binding the contract method 0x30561100.
//
// Solidity: function getAllPosts() view returns((address,string,string,uint256,uint8)[])
func (_Contract *ContractCallerSession) GetAllPosts() ([]PostStoragePost, error) {
	return _Contract.Contract.GetAllPosts(&_Contract.CallOpts)
}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() view returns(address[])
func (_Contract *ContractCaller) GetAllUsers(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getAllUsers")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() view returns(address[])
func (_Contract *ContractSession) GetAllUsers() ([]common.Address, error) {
	return _Contract.Contract.GetAllUsers(&_Contract.CallOpts)
}

// GetAllUsers is a free data retrieval call binding the contract method 0xe2842d79.
//
// Solidity: function getAllUsers() view returns(address[])
func (_Contract *ContractCallerSession) GetAllUsers() ([]common.Address, error) {
	return _Contract.Contract.GetAllUsers(&_Contract.CallOpts)
}

// GetPostByIndex is a free data retrieval call binding the contract method 0x7c7bf4df.
//
// Solidity: function getPostByIndex(address user, uint256 index) view returns((address,string,string,uint256,uint8))
func (_Contract *ContractCaller) GetPostByIndex(opts *bind.CallOpts, user common.Address, index *big.Int) (PostStoragePost, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPostByIndex", user, index)

	if err != nil {
		return *new(PostStoragePost), err
	}

	out0 := *abi.ConvertType(out[0], new(PostStoragePost)).(*PostStoragePost)

	return out0, err

}

// GetPostByIndex is a free data retrieval call binding the contract method 0x7c7bf4df.
//
// Solidity: function getPostByIndex(address user, uint256 index) view returns((address,string,string,uint256,uint8))
func (_Contract *ContractSession) GetPostByIndex(user common.Address, index *big.Int) (PostStoragePost, error) {
	return _Contract.Contract.GetPostByIndex(&_Contract.CallOpts, user, index)
}

// GetPostByIndex is a free data retrieval call binding the contract method 0x7c7bf4df.
//
// Solidity: function getPostByIndex(address user, uint256 index) view returns((address,string,string,uint256,uint8))
func (_Contract *ContractCallerSession) GetPostByIndex(user common.Address, index *big.Int) (PostStoragePost, error) {
	return _Contract.Contract.GetPostByIndex(&_Contract.CallOpts, user, index)
}

// GetUsersPost is a free data retrieval call binding the contract method 0x3aa5cc94.
//
// Solidity: function getUsersPost(address user) view returns((address,string,string,uint256,uint8)[])
func (_Contract *ContractCaller) GetUsersPost(opts *bind.CallOpts, user common.Address) ([]PostStoragePost, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getUsersPost", user)

	if err != nil {
		return *new([]PostStoragePost), err
	}

	out0 := *abi.ConvertType(out[0], new([]PostStoragePost)).(*[]PostStoragePost)

	return out0, err

}

// GetUsersPost is a free data retrieval call binding the contract method 0x3aa5cc94.
//
// Solidity: function getUsersPost(address user) view returns((address,string,string,uint256,uint8)[])
func (_Contract *ContractSession) GetUsersPost(user common.Address) ([]PostStoragePost, error) {
	return _Contract.Contract.GetUsersPost(&_Contract.CallOpts, user)
}

// GetUsersPost is a free data retrieval call binding the contract method 0x3aa5cc94.
//
// Solidity: function getUsersPost(address user) view returns((address,string,string,uint256,uint8)[])
func (_Contract *ContractCallerSession) GetUsersPost(user common.Address) ([]PostStoragePost, error) {
	return _Contract.Contract.GetUsersPost(&_Contract.CallOpts, user)
}

// UserPosts is a free data retrieval call binding the contract method 0x8064d149.
//
// Solidity: function userPosts(address , uint256 ) view returns(address author, string title, string content, uint256 timestamp, uint8 category)
func (_Contract *ContractCaller) UserPosts(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Author    common.Address
	Title     string
	Content   string
	Timestamp *big.Int
	Category  uint8
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "userPosts", arg0, arg1)

	outstruct := new(struct {
		Author    common.Address
		Title     string
		Content   string
		Timestamp *big.Int
		Category  uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Author = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Title = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Content = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Category = *abi.ConvertType(out[4], new(uint8)).(*uint8)

	return *outstruct, err

}

// UserPosts is a free data retrieval call binding the contract method 0x8064d149.
//
// Solidity: function userPosts(address , uint256 ) view returns(address author, string title, string content, uint256 timestamp, uint8 category)
func (_Contract *ContractSession) UserPosts(arg0 common.Address, arg1 *big.Int) (struct {
	Author    common.Address
	Title     string
	Content   string
	Timestamp *big.Int
	Category  uint8
}, error) {
	return _Contract.Contract.UserPosts(&_Contract.CallOpts, arg0, arg1)
}

// UserPosts is a free data retrieval call binding the contract method 0x8064d149.
//
// Solidity: function userPosts(address , uint256 ) view returns(address author, string title, string content, uint256 timestamp, uint8 category)
func (_Contract *ContractCallerSession) UserPosts(arg0 common.Address, arg1 *big.Int) (struct {
	Author    common.Address
	Title     string
	Content   string
	Timestamp *big.Int
	Category  uint8
}, error) {
	return _Contract.Contract.UserPosts(&_Contract.CallOpts, arg0, arg1)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_Contract *ContractCaller) Users(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "users", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_Contract *ContractSession) Users(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.Users(&_Contract.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0x365b98b2.
//
// Solidity: function users(uint256 ) view returns(address)
func (_Contract *ContractCallerSession) Users(arg0 *big.Int) (common.Address, error) {
	return _Contract.Contract.Users(&_Contract.CallOpts, arg0)
}

// Post is a paid mutator transaction binding the contract method 0xe8b4edb4.
//
// Solidity: function post(string _title, string _content, uint8 _category, address author) returns()
func (_Contract *ContractTransactor) Post(opts *bind.TransactOpts, _title string, _content string, _category uint8, author common.Address) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "post", _title, _content, _category, author)
}

// Post is a paid mutator transaction binding the contract method 0xe8b4edb4.
//
// Solidity: function post(string _title, string _content, uint8 _category, address author) returns()
func (_Contract *ContractSession) Post(_title string, _content string, _category uint8, author common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Post(&_Contract.TransactOpts, _title, _content, _category, author)
}

// Post is a paid mutator transaction binding the contract method 0xe8b4edb4.
//
// Solidity: function post(string _title, string _content, uint8 _category, address author) returns()
func (_Contract *ContractTransactorSession) Post(_title string, _content string, _category uint8, author common.Address) (*types.Transaction, error) {
	return _Contract.Contract.Post(&_Contract.TransactOpts, _title, _content, _category, author)
}

// ContractPostPublishedIterator is returned from FilterPostPublished and is used to iterate over the raw logs and unpacked data for PostPublished events raised by the Contract contract.
type ContractPostPublishedIterator struct {
	Event *ContractPostPublished // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *ContractPostPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractPostPublished)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(ContractPostPublished)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *ContractPostPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractPostPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractPostPublished represents a PostPublished event raised by the Contract contract.
type ContractPostPublished struct {
	Author   common.Address
	Title    string
	Category uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPostPublished is a free log retrieval operation binding the contract event 0x4be77e23b4d9d657bb35113acb2ddeb9fabadff4c213859965d0b6d9b3857381.
//
// Solidity: event PostPublished(address indexed author, string title, uint8 category)
func (_Contract *ContractFilterer) FilterPostPublished(opts *bind.FilterOpts, author []common.Address) (*ContractPostPublishedIterator, error) {

	var authorRule []interface{}
	for _, authorItem := range author {
		authorRule = append(authorRule, authorItem)
	}

	logs, sub, err := _Contract.contract.FilterLogs(opts, "PostPublished", authorRule)
	if err != nil {
		return nil, err
	}
	return &ContractPostPublishedIterator{contract: _Contract.contract, event: "PostPublished", logs: logs, sub: sub}, nil
}

// WatchPostPublished is a free log subscription operation binding the contract event 0x4be77e23b4d9d657bb35113acb2ddeb9fabadff4c213859965d0b6d9b3857381.
//
// Solidity: event PostPublished(address indexed author, string title, uint8 category)
func (_Contract *ContractFilterer) WatchPostPublished(opts *bind.WatchOpts, sink chan<- *ContractPostPublished, author []common.Address) (event.Subscription, error) {

	var authorRule []interface{}
	for _, authorItem := range author {
		authorRule = append(authorRule, authorItem)
	}

	logs, sub, err := _Contract.contract.WatchLogs(opts, "PostPublished", authorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractPostPublished)
				if err := _Contract.contract.UnpackLog(event, "PostPublished", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePostPublished is a log parse operation binding the contract event 0x4be77e23b4d9d657bb35113acb2ddeb9fabadff4c213859965d0b6d9b3857381.
//
// Solidity: event PostPublished(address indexed author, string title, uint8 category)
func (_Contract *ContractFilterer) ParsePostPublished(log types.Log) (*ContractPostPublished, error) {
	event := new(ContractPostPublished)
	if err := _Contract.contract.UnpackLog(event, "PostPublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
