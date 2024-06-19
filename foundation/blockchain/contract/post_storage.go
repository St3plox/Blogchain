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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumPostStorage.Category\",\"name\":\"category\",\"type\":\"uint8\"}],\"name\":\"PostPublished\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getMyPosts\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumPostStorage.Category\",\"name\":\"category\",\"type\":\"uint8\"}],\"internalType\":\"structPostStorage.Post[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUsersPost\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumPostStorage.Category\",\"name\":\"category\",\"type\":\"uint8\"}],\"internalType\":\"structPostStorage.Post[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"},{\"internalType\":\"enumPostStorage.Category\",\"name\":\"_category\",\"type\":\"uint8\"}],\"name\":\"post\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userPosts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumPostStorage.Category\",\"name\":\"category\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f80fd5b506111f78061001c5f395ff3fe608060405234801561000f575f80fd5b506004361061004a575f3560e01c80631998ab801461004e5780633aa5cc941461006a5780635950eb5b1461009a5780638064d149146100b8575b5f80fd5b61006860048036038101906100639190610a32565b6100ec565b005b610084600480360381019061007f9190610b14565b6102a2565b6040516100919190610d6e565b60405180910390f35b6100a26104f2565b6040516100af9190610d6e565b60405180910390f35b6100d260048036038101906100cd9190610db8565b610740565b6040516100e3959493929190610e6b565b60405180910390f35b5f6040518060a001604052803373ffffffffffffffffffffffffffffffffffffffff16815260200185815260200184815260200142815260200183600281111561013957610138610bef565b5b81525090505f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f2081908060018154018082558091505060019003905f5260205f2090600502015f909190919091505f820151815f015f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010190816101fa91906110c4565b50604082015181600201908161021091906110c4565b50606082015181600301556080820151816004015f6101000a81548160ff0219169083600281111561024557610244610bef565b5b021790555050503373ffffffffffffffffffffffffffffffffffffffff167f4be77e23b4d9d657bb35113acb2ddeb9fabadff4c213859965d0b6d9b38573818584604051610294929190611193565b60405180910390a250505050565b60605f808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b828210156104e7578382905f5260205f2090600502016040518060a00160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201805461038490610ef7565b80601f01602080910402602001604051908101604052809291908181526020018280546103b090610ef7565b80156103fb5780601f106103d2576101008083540402835291602001916103fb565b820191905f5260205f20905b8154815290600101906020018083116103de57829003601f168201915b5050505050815260200160028201805461041490610ef7565b80601f016020809104026020016040519081016040528092919081815260200182805461044090610ef7565b801561048b5780601f106104625761010080835404028352916020019161048b565b820191905f5260205f20905b81548152906001019060200180831161046e57829003601f168201915b5050505050815260200160038201548152602001600482015f9054906101000a900460ff1660028111156104c2576104c1610bef565b5b60028111156104d4576104d3610bef565b5b81525050815260200190600101906102ff565b505050509050919050565b60605f803373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20805480602002602001604051908101604052809291908181526020015f905b82821015610737578382905f5260205f2090600502016040518060a00160405290815f82015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820180546105d490610ef7565b80601f016020809104026020016040519081016040528092919081815260200182805461060090610ef7565b801561064b5780601f106106225761010080835404028352916020019161064b565b820191905f5260205f20905b81548152906001019060200180831161062e57829003601f168201915b5050505050815260200160028201805461066490610ef7565b80601f016020809104026020016040519081016040528092919081815260200182805461069090610ef7565b80156106db5780601f106106b2576101008083540402835291602001916106db565b820191905f5260205f20905b8154815290600101906020018083116106be57829003601f168201915b5050505050815260200160038201548152602001600482015f9054906101000a900460ff16600281111561071257610711610bef565b5b600281111561072457610723610bef565b5b815250508152602001906001019061054f565b50505050905090565b5f602052815f5260405f208181548110610758575f80fd5b905f5260205f2090600502015f9150915050805f015f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169080600101805461079d90610ef7565b80601f01602080910402602001604051908101604052809291908181526020018280546107c990610ef7565b80156108145780601f106107eb57610100808354040283529160200191610814565b820191905f5260205f20905b8154815290600101906020018083116107f757829003601f168201915b50505050509080600201805461082990610ef7565b80601f016020809104026020016040519081016040528092919081815260200182805461085590610ef7565b80156108a05780601f10610877576101008083540402835291602001916108a0565b820191905f5260205f20905b81548152906001019060200180831161088357829003601f168201915b505050505090806003015490806004015f9054906101000a900460ff16905085565b5f604051905090565b5f80fd5b5f80fd5b5f80fd5b5f80fd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610921826108db565b810181811067ffffffffffffffff821117156109405761093f6108eb565b5b80604052505050565b5f6109526108c2565b905061095e8282610918565b919050565b5f67ffffffffffffffff82111561097d5761097c6108eb565b5b610986826108db565b9050602081019050919050565b828183375f83830152505050565b5f6109b36109ae84610963565b610949565b9050828152602081018484840111156109cf576109ce6108d7565b5b6109da848285610993565b509392505050565b5f82601f8301126109f6576109f56108d3565b5b8135610a068482602086016109a1565b91505092915050565b60038110610a1b575f80fd5b50565b5f81359050610a2c81610a0f565b92915050565b5f805f60608486031215610a4957610a486108cb565b5b5f84013567ffffffffffffffff811115610a6657610a656108cf565b5b610a72868287016109e2565b935050602084013567ffffffffffffffff811115610a9357610a926108cf565b5b610a9f868287016109e2565b9250506040610ab086828701610a1e565b9150509250925092565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f610ae382610aba565b9050919050565b610af381610ad9565b8114610afd575f80fd5b50565b5f81359050610b0e81610aea565b92915050565b5f60208284031215610b2957610b286108cb565b5b5f610b3684828501610b00565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b610b7181610ad9565b82525050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f610ba982610b77565b610bb38185610b81565b9350610bc3818560208601610b91565b610bcc816108db565b840191505092915050565b5f819050919050565b610be981610bd7565b82525050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602160045260245ffd5b60038110610c2d57610c2c610bef565b5b50565b5f819050610c3d82610c1c565b919050565b5f610c4c82610c30565b9050919050565b610c5c81610c42565b82525050565b5f60a083015f830151610c775f860182610b68565b5060208301518482036020860152610c8f8282610b9f565b91505060408301518482036040860152610ca98282610b9f565b9150506060830151610cbe6060860182610be0565b506080830151610cd16080860182610c53565b508091505092915050565b5f610ce78383610c62565b905092915050565b5f602082019050919050565b5f610d0582610b3f565b610d0f8185610b49565b935083602082028501610d2185610b59565b805f5b85811015610d5c5784840389528151610d3d8582610cdc565b9450610d4883610cef565b925060208a01995050600181019050610d24565b50829750879550505050505092915050565b5f6020820190508181035f830152610d868184610cfb565b905092915050565b610d9781610bd7565b8114610da1575f80fd5b50565b5f81359050610db281610d8e565b92915050565b5f8060408385031215610dce57610dcd6108cb565b5b5f610ddb85828601610b00565b9250506020610dec85828601610da4565b9150509250929050565b610dff81610ad9565b82525050565b5f82825260208201905092915050565b5f610e1f82610b77565b610e298185610e05565b9350610e39818560208601610b91565b610e42816108db565b840191505092915050565b610e5681610bd7565b82525050565b610e6581610c42565b82525050565b5f60a082019050610e7e5f830188610df6565b8181036020830152610e908187610e15565b90508181036040830152610ea48186610e15565b9050610eb36060830185610e4d565b610ec06080830184610e5c565b9695505050505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f6002820490506001821680610f0e57607f821691505b602082108103610f2157610f20610eca565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f60088302610f837fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610f48565b610f8d8683610f48565b95508019841693508086168417925050509392505050565b5f819050919050565b5f610fc8610fc3610fbe84610bd7565b610fa5565b610bd7565b9050919050565b5f819050919050565b610fe183610fae565b610ff5610fed82610fcf565b848454610f54565b825550505050565b5f90565b611009610ffd565b611014818484610fd8565b505050565b5b818110156110375761102c5f82611001565b60018101905061101a565b5050565b601f82111561107c5761104d81610f27565b61105684610f39565b81016020851015611065578190505b61107961107185610f39565b830182611019565b50505b505050565b5f82821c905092915050565b5f61109c5f1984600802611081565b1980831691505092915050565b5f6110b4838361108d565b9150826002028217905092915050565b6110cd82610b77565b67ffffffffffffffff8111156110e6576110e56108eb565b5b6110f08254610ef7565b6110fb82828561103b565b5f60209050601f83116001811461112c575f841561111a578287015190505b61112485826110a9565b86555061118b565b601f19841661113a86610f27565b5f5b828110156111615784890151825560018201915060208501945060208101905061113c565b8683101561117e578489015161117a601f89168261108d565b8355505b6001600288020188555050505b505050505050565b5f6040820190508181035f8301526111ab8185610e15565b90506111ba6020830184610e5c565b939250505056fea2646970667358221220431b5d8095d8084e86e522f0bff60dffc527a0d9b0f41e2c3afb7ce862ccbadd64736f6c634300081a0033",
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

// GetMyPosts is a free data retrieval call binding the contract method 0x5950eb5b.
//
// Solidity: function getMyPosts() view returns((address,string,string,uint256,uint8)[])
func (_Contract *ContractCaller) GetMyPosts(opts *bind.CallOpts) ([]PostStoragePost, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getMyPosts")

	if err != nil {
		return *new([]PostStoragePost), err
	}

	out0 := *abi.ConvertType(out[0], new([]PostStoragePost)).(*[]PostStoragePost)

	return out0, err

}

// GetMyPosts is a free data retrieval call binding the contract method 0x5950eb5b.
//
// Solidity: function getMyPosts() view returns((address,string,string,uint256,uint8)[])
func (_Contract *ContractSession) GetMyPosts() ([]PostStoragePost, error) {
	return _Contract.Contract.GetMyPosts(&_Contract.CallOpts)
}

// GetMyPosts is a free data retrieval call binding the contract method 0x5950eb5b.
//
// Solidity: function getMyPosts() view returns((address,string,string,uint256,uint8)[])
func (_Contract *ContractCallerSession) GetMyPosts() ([]PostStoragePost, error) {
	return _Contract.Contract.GetMyPosts(&_Contract.CallOpts)
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

// Post is a paid mutator transaction binding the contract method 0x1998ab80.
//
// Solidity: function post(string _title, string _content, uint8 _category) returns()
func (_Contract *ContractTransactor) Post(opts *bind.TransactOpts, _title string, _content string, _category uint8) (*types.Transaction, error) {
	return _Contract.contract.Transact(opts, "post", _title, _content, _category)
}

// Post is a paid mutator transaction binding the contract method 0x1998ab80.
//
// Solidity: function post(string _title, string _content, uint8 _category) returns()
func (_Contract *ContractSession) Post(_title string, _content string, _category uint8) (*types.Transaction, error) {
	return _Contract.Contract.Post(&_Contract.TransactOpts, _title, _content, _category)
}

// Post is a paid mutator transaction binding the contract method 0x1998ab80.
//
// Solidity: function post(string _title, string _content, uint8 _category) returns()
func (_Contract *ContractTransactorSession) Post(_title string, _content string, _category uint8) (*types.Transaction, error) {
	return _Contract.Contract.Post(&_Contract.TransactOpts, _title, _content, _category)
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
