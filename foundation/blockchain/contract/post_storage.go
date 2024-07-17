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
	Id        *big.Int
	Author    common.Address
	Title     string
	Content   string
	Timestamp *big.Int
	Category  uint8
}

// ContractMetaData contains all meta data concerning the Contract contract.
var ContractMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"enumPostStorage.Category\",\"name\":\"category\",\"type\":\"uint8\"}],\"name\":\"PostPublished\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getAllPosts\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumPostStorage.Category\",\"name\":\"category\",\"type\":\"uint8\"}],\"internalType\":\"structPostStorage.Post[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllUsers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"getPostByID\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumPostStorage.Category\",\"name\":\"category\",\"type\":\"uint8\"}],\"internalType\":\"structPostStorage.Post\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getPostByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumPostStorage.Category\",\"name\":\"category\",\"type\":\"uint8\"}],\"internalType\":\"structPostStorage.Post\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"page\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getPostsPaginated\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumPostStorage.Category\",\"name\":\"category\",\"type\":\"uint8\"}],\"internalType\":\"structPostStorage.Post[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUsersPost\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumPostStorage.Category\",\"name\":\"category\",\"type\":\"uint8\"}],\"internalType\":\"structPostStorage.Post[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"},{\"internalType\":\"enumPostStorage.Category\",\"name\":\"_category\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"}],\"name\":\"post\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postCounter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"userPosts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"author\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumPostStorage.Category\",\"name\":\"category\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612951806100206000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80635645c1e7116100665780635645c1e71461016f5780637c7bf4df1461019f5780638064d149146101cf578063e2842d7914610204578063e8b4edb4146102225761009e565b806309b3d75e146100a357806330561100146100c1578063365b98b2146100df5780633aa5cc941461010f5780633e92733e1461013f575b600080fd5b6100ab61023e565b6040516100b89190612432565b60405180910390f35b6100c9610244565b6040516100d6919061238e565b60405180910390f35b6100f960048036038101906100f49190611f41565b61078f565b6040516101069190612351565b60405180910390f35b61012960048036038101906101249190611e49565b6107ce565b604051610136919061238e565b60405180910390f35b61015960048036038101906101549190611f6a565b610a80565b604051610166919061238e565b60405180910390f35b61018960048036038101906101849190611f41565b611087565b6040516101969190612410565b60405180910390f35b6101b960048036038101906101b49190611e72565b611323565b6040516101c69190612410565b60405180910390f35b6101e960048036038101906101e49190611e72565b61165a565b6040516101fb9695949392919061244d565b60405180910390f35b61020c6117f0565b604051610219919061236c565b60405180910390f35b61023c60048036038101906102379190611eae565b61187e565b005b60035481565b60606000805b60028054905081101561031d5760008060028381548110610294577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805490508261030891906125fa565b91508080610315906127b3565b91505061024a565b5060008167ffffffffffffffff811115610360577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60405190808252806020026020018201604052801561039957816020015b610386611c7b565b81526020019060019003908161037e5790505b5090506000805b600280549050811015610785576000806000600284815481106103ec577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b828210156106b857838290600052602060002090600602016040518060c0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201805461050490612781565b80601f016020809104026020016040519081016040528092919081815260200182805461053090612781565b801561057d5780601f106105525761010080835404028352916020019161057d565b820191906000526020600020905b81548152906001019060200180831161056057829003601f168201915b5050505050815260200160038201805461059690612781565b80601f01602080910402602001604051908101604052809291908181526020018280546105c290612781565b801561060f5780601f106105e45761010080835404028352916020019161060f565b820191906000526020600020905b8154815290600101906020018083116105f257829003601f168201915b50505050508152602001600482015481526020016005820160009054906101000a900460ff16600281111561066d577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60028111156106a5577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152505081526020019060010190610471565b50505050905060005b815181101561077057818181518110610703577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151858581518110610744577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181905250838061075a906127b3565b9450508080610768906127b3565b9150506106c1565b5050808061077d906127b3565b9150506103a0565b5081935050505090565b6002818154811061079f57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60606000808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b82821015610a7557838290600052602060002090600602016040518060c0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820180546108c190612781565b80601f01602080910402602001604051908101604052809291908181526020018280546108ed90612781565b801561093a5780601f1061090f5761010080835404028352916020019161093a565b820191906000526020600020905b81548152906001019060200180831161091d57829003601f168201915b5050505050815260200160038201805461095390612781565b80601f016020809104026020016040519081016040528092919081815260200182805461097f90612781565b80156109cc5780601f106109a1576101008083540402835291602001916109cc565b820191906000526020600020905b8154815290600101906020018083116109af57829003601f168201915b50505050508152602001600482015481526020016005820160009054906101000a900460ff166002811115610a2a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115610a62577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815250508152602001906001019061082e565b505050509050919050565b60606000805b600280549050811015610b595760008060028381548110610ad0577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054905082610b4491906125fa565b91508080610b51906127b3565b915050610a86565b50808385610b679190612650565b10610ba7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b9e906123f0565b60405180910390fd5b60008385610bb59190612650565b90506000828583610bc691906125fa565b11610bdc578482610bd791906125fa565b610bde565b825b905060008282610bee91906126aa565b67ffffffffffffffff811115610c2d577f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051908082528060200260200182016040528015610c6657816020015b610c53611c7b565b815260200190600190039081610c4b5790505b50905060008060005b60028054905081101561107757600080600060028481548110610cbb577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020016000905b82821015610f8757838290600052602060002090600602016040518060c0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600282018054610dd390612781565b80601f0160208091040260200160405190810160405280929190818152602001828054610dff90612781565b8015610e4c5780601f10610e2157610100808354040283529160200191610e4c565b820191906000526020600020905b815481529060010190602001808311610e2f57829003601f168201915b50505050508152602001600382018054610e6590612781565b80601f0160208091040260200160405190810160405280929190818152602001828054610e9190612781565b8015610ede5780601f10610eb357610100808354040283529160200191610ede565b820191906000526020600020905b815481529060010190602001808311610ec157829003601f168201915b50505050508152602001600482015481526020016005820160009054906101000a900460ff166002811115610f3c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115610f74577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b8152505081526020019060010190610d40565b50505050905060005b815181101561106257878410158015610fa857508684105b1561104157818181518110610fe6577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010151868681518110611027577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6020026020010181905250848061103d906127b3565b9550505b838061104c906127b3565b945050808061105a906127b3565b915050610f90565b5050808061106f906127b3565b915050610c6f565b5082965050505050505092915050565b61108f611c7b565b6000600160008481526020019081526020016000206000015414156110e9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016110e0906123b0565b60405180910390fd5b600160008381526020019081526020016000206040518060c0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160028201805461117790612781565b80601f01602080910402602001604051908101604052809291908181526020018280546111a390612781565b80156111f05780601f106111c5576101008083540402835291602001916111f0565b820191906000526020600020905b8154815290600101906020018083116111d357829003601f168201915b5050505050815260200160038201805461120990612781565b80601f016020809104026020016040519081016040528092919081815260200182805461123590612781565b80156112825780601f1061125757610100808354040283529160200191611282565b820191906000526020600020905b81548152906001019060200180831161126557829003601f168201915b50505050508152602001600482015481526020016005820160009054906101000a900460ff1660028111156112e0577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6002811115611318577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b815250509050919050565b61132b611c7b565b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054905082106113ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016113a5906123d0565b60405180910390fd5b6000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208281548110611424577f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b90600052602060002090600602016040518060c0016040529081600082015481526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820180546114ad90612781565b80601f01602080910402602001604051908101604052809291908181526020018280546114d990612781565b80156115265780601f106114fb57610100808354040283529160200191611526565b820191906000526020600020905b81548152906001019060200180831161150957829003601f168201915b5050505050815260200160038201805461153f90612781565b80601f016020809104026020016040519081016040528092919081815260200182805461156b90612781565b80156115b85780601f1061158d576101008083540402835291602001916115b8565b820191906000526020600020905b81548152906001019060200180831161159b57829003601f168201915b50505050508152602001600482015481526020016005820160009054906101000a900460ff166002811115611616577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600281111561164e577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81525050905092915050565b6000602052816000526040600020818154811061167657600080fd5b9060005260206000209060060201600091509150508060000154908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020180546116c690612781565b80601f01602080910402602001604051908101604052809291908181526020018280546116f290612781565b801561173f5780601f106117145761010080835404028352916020019161173f565b820191906000526020600020905b81548152906001019060200180831161172257829003601f168201915b50505050509080600301805461175490612781565b80601f016020809104026020016040519081016040528092919081815260200182805461178090612781565b80156117cd5780601f106117a2576101008083540402835291602001916117cd565b820191906000526020600020905b8154815290600101906020018083116117b057829003601f168201915b5050505050908060040154908060050160009054906101000a900460ff16905086565b6060600280548060200260200160405190810160405280929190818152602001828054801561187457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161182a575b5050505050905090565b60008060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002080549050141561192c576002819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b6003600081548092919061193f906127b3565b91905055506000600354905060006040518060c001604052808381526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018781526020018681526020014281526020018560028111156119c5577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81525090506000808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190806001815401808255809150506001900390600052602060002090600602016000909190919091506000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002019080519060200190611a9f929190611cff565b506060820151816003019080519060200190611abc929190611cff565b506080820151816004015560a08201518160050160006101000a81548160ff02191690836002811115611b18577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b0217905550505080600160008481526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002019080519060200190611ba0929190611cff565b506060820151816003019080519060200190611bbd929190611cff565b506080820151816004015560a08201518160050160006101000a81548160ff02191690836002811115611c19577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b02179055509050508273ffffffffffffffffffffffffffffffffffffffff167ffbd60bad560a02cd13b1995d28d1379c3deb7bc6dd8d65adcb98d895ea4b7fa4838887604051611c6b939291906124bc565b60405180910390a2505050505050565b6040518060c0016040528060008152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001606081526020016000815260200160006002811115611cf9577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b81525090565b828054611d0b90612781565b90600052602060002090601f016020900481019282611d2d5760008555611d74565b82601f10611d4657805160ff1916838001178555611d74565b82800160010185558215611d74579182015b82811115611d73578251825591602001919060010190611d58565b5b509050611d819190611d85565b5090565b5b80821115611d9e576000816000905550600101611d86565b5090565b6000611db5611db08461252b565b6124fa565b905082815260208101848484011115611dcd57600080fd5b611dd884828561273f565b509392505050565b600081359050611def816128dd565b92915050565b600081359050611e04816128f4565b92915050565b600082601f830112611e1b57600080fd5b8135611e2b848260208601611da2565b91505092915050565b600081359050611e4381612904565b92915050565b600060208284031215611e5b57600080fd5b6000611e6984828501611de0565b91505092915050565b60008060408385031215611e8557600080fd5b6000611e9385828601611de0565b9250506020611ea485828601611e34565b9150509250929050565b60008060008060808587031215611ec457600080fd5b600085013567ffffffffffffffff811115611ede57600080fd5b611eea87828801611e0a565b945050602085013567ffffffffffffffff811115611f0757600080fd5b611f1387828801611e0a565b9350506040611f2487828801611df5565b9250506060611f3587828801611de0565b91505092959194509250565b600060208284031215611f5357600080fd5b6000611f6184828501611e34565b91505092915050565b60008060408385031215611f7d57600080fd5b6000611f8b85828601611e34565b9250506020611f9c85828601611e34565b9150509250929050565b6000611fb28383611fd2565b60208301905092915050565b6000611fca8383612213565b905092915050565b611fdb816126de565b82525050565b611fea816126de565b82525050565b6000611ffb8261257b565b61200581856125b6565b93506120108361255b565b8060005b838110156120415781516120288882611fa6565b97506120338361259c565b925050600181019050612014565b5085935050505092915050565b600061205982612586565b61206381856125c7565b9350836020820285016120758561256b565b8060005b858110156120b157848403895281516120928582611fbe565b945061209d836125a9565b925060208a01995050600181019050612079565b50829750879550505050505092915050565b6120cc8161272d565b82525050565b6120db8161272d565b82525050565b60006120ec82612591565b6120f681856125d8565b935061210681856020860161274e565b61210f816128b8565b840191505092915050565b600061212582612591565b61212f81856125e9565b935061213f81856020860161274e565b612148816128b8565b840191505092915050565b60006121606020836125e9565b91507f506f73742077697468207468697320494420646f6573206e6f742065786973746000830152602082019050919050565b60006121a06013836125e9565b91507f506f737420646f6573206e6f74206578697374000000000000000000000000006000830152602082019050919050565b60006121e06011836125e9565b91507f50616765206f7574206f662072616e67650000000000000000000000000000006000830152602082019050919050565b600060c08301600083015161222b6000860182612333565b50602083015161223e6020860182611fd2565b506040830151848203604086015261225682826120e1565b9150506060830151848203606086015261227082826120e1565b91505060808301516122856080860182612333565b5060a083015161229860a08601826120c3565b508091505092915050565b600060c0830160008301516122bb6000860182612333565b5060208301516122ce6020860182611fd2565b50604083015184820360408601526122e682826120e1565b9150506060830151848203606086015261230082826120e1565b91505060808301516123156080860182612333565b5060a083015161232860a08601826120c3565b508091505092915050565b61233c81612723565b82525050565b61234b81612723565b82525050565b60006020820190506123666000830184611fe1565b92915050565b600060208201905081810360008301526123868184611ff0565b905092915050565b600060208201905081810360008301526123a8818461204e565b905092915050565b600060208201905081810360008301526123c981612153565b9050919050565b600060208201905081810360008301526123e981612193565b9050919050565b60006020820190508181036000830152612409816121d3565b9050919050565b6000602082019050818103600083015261242a81846122a3565b905092915050565b60006020820190506124476000830184612342565b92915050565b600060c0820190506124626000830189612342565b61246f6020830188611fe1565b8181036040830152612481818761211a565b90508181036060830152612495818661211a565b90506124a46080830185612342565b6124b160a08301846120d2565b979650505050505050565b60006060820190506124d16000830186612342565b81810360208301526124e3818561211a565b90506124f260408301846120d2565b949350505050565b6000604051905081810181811067ffffffffffffffff8211171561252157612520612889565b5b8060405250919050565b600067ffffffffffffffff82111561254657612545612889565b5b601f19601f8301169050602081019050919050565b6000819050602082019050919050565b6000819050602082019050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000602082019050919050565b6000602082019050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600061260582612723565b915061261083612723565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115612645576126446127fc565b5b828201905092915050565b600061265b82612723565b915061266683612723565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561269f5761269e6127fc565b5b828202905092915050565b60006126b582612723565b91506126c083612723565b9250828210156126d3576126d26127fc565b5b828203905092915050565b60006126e982612703565b9050919050565b60008190506126fe826128c9565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000612738826126f0565b9050919050565b82818337600083830152505050565b60005b8381101561276c578082015181840152602081019050612751565b8381111561277b576000848401525b50505050565b6000600282049050600182168061279957607f821691505b602082108114156127ad576127ac61285a565b5b50919050565b60006127be82612723565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8214156127f1576127f06127fc565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b600381106128da576128d961282b565b5b50565b6128e6816126de565b81146128f157600080fd5b50565b6003811061290157600080fd5b50565b61290d81612723565b811461291857600080fd5b5056fea2646970667358221220f104cd0a344da85eef0cee2c13d34acb1d9ae7fd6ed7b0c1b92dcc0f09af41c764736f6c63430008000033",
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
// Solidity: function getAllPosts() view returns((uint256,address,string,string,uint256,uint8)[])
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
// Solidity: function getAllPosts() view returns((uint256,address,string,string,uint256,uint8)[])
func (_Contract *ContractSession) GetAllPosts() ([]PostStoragePost, error) {
	return _Contract.Contract.GetAllPosts(&_Contract.CallOpts)
}

// GetAllPosts is a free data retrieval call binding the contract method 0x30561100.
//
// Solidity: function getAllPosts() view returns((uint256,address,string,string,uint256,uint8)[])
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

// GetPostByID is a free data retrieval call binding the contract method 0x5645c1e7.
//
// Solidity: function getPostByID(uint256 id) view returns((uint256,address,string,string,uint256,uint8))
func (_Contract *ContractCaller) GetPostByID(opts *bind.CallOpts, id *big.Int) (PostStoragePost, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPostByID", id)

	if err != nil {
		return *new(PostStoragePost), err
	}

	out0 := *abi.ConvertType(out[0], new(PostStoragePost)).(*PostStoragePost)

	return out0, err

}

// GetPostByID is a free data retrieval call binding the contract method 0x5645c1e7.
//
// Solidity: function getPostByID(uint256 id) view returns((uint256,address,string,string,uint256,uint8))
func (_Contract *ContractSession) GetPostByID(id *big.Int) (PostStoragePost, error) {
	return _Contract.Contract.GetPostByID(&_Contract.CallOpts, id)
}

// GetPostByID is a free data retrieval call binding the contract method 0x5645c1e7.
//
// Solidity: function getPostByID(uint256 id) view returns((uint256,address,string,string,uint256,uint8))
func (_Contract *ContractCallerSession) GetPostByID(id *big.Int) (PostStoragePost, error) {
	return _Contract.Contract.GetPostByID(&_Contract.CallOpts, id)
}

// GetPostByIndex is a free data retrieval call binding the contract method 0x7c7bf4df.
//
// Solidity: function getPostByIndex(address user, uint256 index) view returns((uint256,address,string,string,uint256,uint8))
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
// Solidity: function getPostByIndex(address user, uint256 index) view returns((uint256,address,string,string,uint256,uint8))
func (_Contract *ContractSession) GetPostByIndex(user common.Address, index *big.Int) (PostStoragePost, error) {
	return _Contract.Contract.GetPostByIndex(&_Contract.CallOpts, user, index)
}

// GetPostByIndex is a free data retrieval call binding the contract method 0x7c7bf4df.
//
// Solidity: function getPostByIndex(address user, uint256 index) view returns((uint256,address,string,string,uint256,uint8))
func (_Contract *ContractCallerSession) GetPostByIndex(user common.Address, index *big.Int) (PostStoragePost, error) {
	return _Contract.Contract.GetPostByIndex(&_Contract.CallOpts, user, index)
}

// GetPostsPaginated is a free data retrieval call binding the contract method 0x3e92733e.
//
// Solidity: function getPostsPaginated(uint256 page, uint256 pageSize) view returns((uint256,address,string,string,uint256,uint8)[])
func (_Contract *ContractCaller) GetPostsPaginated(opts *bind.CallOpts, page *big.Int, pageSize *big.Int) ([]PostStoragePost, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "getPostsPaginated", page, pageSize)

	if err != nil {
		return *new([]PostStoragePost), err
	}

	out0 := *abi.ConvertType(out[0], new([]PostStoragePost)).(*[]PostStoragePost)

	return out0, err

}

// GetPostsPaginated is a free data retrieval call binding the contract method 0x3e92733e.
//
// Solidity: function getPostsPaginated(uint256 page, uint256 pageSize) view returns((uint256,address,string,string,uint256,uint8)[])
func (_Contract *ContractSession) GetPostsPaginated(page *big.Int, pageSize *big.Int) ([]PostStoragePost, error) {
	return _Contract.Contract.GetPostsPaginated(&_Contract.CallOpts, page, pageSize)
}

// GetPostsPaginated is a free data retrieval call binding the contract method 0x3e92733e.
//
// Solidity: function getPostsPaginated(uint256 page, uint256 pageSize) view returns((uint256,address,string,string,uint256,uint8)[])
func (_Contract *ContractCallerSession) GetPostsPaginated(page *big.Int, pageSize *big.Int) ([]PostStoragePost, error) {
	return _Contract.Contract.GetPostsPaginated(&_Contract.CallOpts, page, pageSize)
}

// GetUsersPost is a free data retrieval call binding the contract method 0x3aa5cc94.
//
// Solidity: function getUsersPost(address user) view returns((uint256,address,string,string,uint256,uint8)[])
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
// Solidity: function getUsersPost(address user) view returns((uint256,address,string,string,uint256,uint8)[])
func (_Contract *ContractSession) GetUsersPost(user common.Address) ([]PostStoragePost, error) {
	return _Contract.Contract.GetUsersPost(&_Contract.CallOpts, user)
}

// GetUsersPost is a free data retrieval call binding the contract method 0x3aa5cc94.
//
// Solidity: function getUsersPost(address user) view returns((uint256,address,string,string,uint256,uint8)[])
func (_Contract *ContractCallerSession) GetUsersPost(user common.Address) ([]PostStoragePost, error) {
	return _Contract.Contract.GetUsersPost(&_Contract.CallOpts, user)
}

// PostCounter is a free data retrieval call binding the contract method 0x09b3d75e.
//
// Solidity: function postCounter() view returns(uint256)
func (_Contract *ContractCaller) PostCounter(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "postCounter")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PostCounter is a free data retrieval call binding the contract method 0x09b3d75e.
//
// Solidity: function postCounter() view returns(uint256)
func (_Contract *ContractSession) PostCounter() (*big.Int, error) {
	return _Contract.Contract.PostCounter(&_Contract.CallOpts)
}

// PostCounter is a free data retrieval call binding the contract method 0x09b3d75e.
//
// Solidity: function postCounter() view returns(uint256)
func (_Contract *ContractCallerSession) PostCounter() (*big.Int, error) {
	return _Contract.Contract.PostCounter(&_Contract.CallOpts)
}

// UserPosts is a free data retrieval call binding the contract method 0x8064d149.
//
// Solidity: function userPosts(address , uint256 ) view returns(uint256 id, address author, string title, string content, uint256 timestamp, uint8 category)
func (_Contract *ContractCaller) UserPosts(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Id        *big.Int
	Author    common.Address
	Title     string
	Content   string
	Timestamp *big.Int
	Category  uint8
}, error) {
	var out []interface{}
	err := _Contract.contract.Call(opts, &out, "userPosts", arg0, arg1)

	outstruct := new(struct {
		Id        *big.Int
		Author    common.Address
		Title     string
		Content   string
		Timestamp *big.Int
		Category  uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Author = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Title = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Content = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Timestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Category = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// UserPosts is a free data retrieval call binding the contract method 0x8064d149.
//
// Solidity: function userPosts(address , uint256 ) view returns(uint256 id, address author, string title, string content, uint256 timestamp, uint8 category)
func (_Contract *ContractSession) UserPosts(arg0 common.Address, arg1 *big.Int) (struct {
	Id        *big.Int
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
// Solidity: function userPosts(address , uint256 ) view returns(uint256 id, address author, string title, string content, uint256 timestamp, uint8 category)
func (_Contract *ContractCallerSession) UserPosts(arg0 common.Address, arg1 *big.Int) (struct {
	Id        *big.Int
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
	Id       *big.Int
	Author   common.Address
	Title    string
	Category uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPostPublished is a free log retrieval operation binding the contract event 0xfbd60bad560a02cd13b1995d28d1379c3deb7bc6dd8d65adcb98d895ea4b7fa4.
//
// Solidity: event PostPublished(uint256 id, address indexed author, string title, uint8 category)
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

// WatchPostPublished is a free log subscription operation binding the contract event 0xfbd60bad560a02cd13b1995d28d1379c3deb7bc6dd8d65adcb98d895ea4b7fa4.
//
// Solidity: event PostPublished(uint256 id, address indexed author, string title, uint8 category)
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

// ParsePostPublished is a log parse operation binding the contract event 0xfbd60bad560a02cd13b1995d28d1379c3deb7bc6dd8d65adcb98d895ea4b7fa4.
//
// Solidity: event PostPublished(uint256 id, address indexed author, string title, uint8 category)
func (_Contract *ContractFilterer) ParsePostPublished(log types.Log) (*ContractPostPublished, error) {
	event := new(ContractPostPublished)
	if err := _Contract.contract.UnpackLog(event, "PostPublished", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
