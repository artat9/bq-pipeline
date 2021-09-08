// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nameRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"}],\"name\":\"Accept\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"Bid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Book\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Call\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bitId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"successfulBidder\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"Close\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"Deny\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromTimestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTimestamp\",\"type\":\"uint256\"}],\"name\":\"NewPost\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"Propose\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"bitId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Refund\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"accept\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPosts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"fromTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"successfulBidId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bidderInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"enumAdManager.DraftStatus\",\"name\":\"status\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"bidderList\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bidders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"book\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bookedBidIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"}],\"name\":\"call\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"}],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"displayByMetadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributionRightAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"inventories\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mediaMetadata\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"metadataList\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nameRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"fromTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTimestamp\",\"type\":\"uint256\"}],\"name\":\"newPost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextBidId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextPostId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postOwnerPoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"bidId\",\"type\":\"uint256\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservedRightAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// AllPosts is a free data retrieval call binding the contract method 0x718ce2bc.
//
// Solidity: function allPosts(uint256 ) view returns(uint256 postId, address owner, string metadata, uint256 fromTimestamp, uint256 toTimestamp, uint256 successfulBidId)
func (_Contracts *ContractsCaller) AllPosts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PostId          *big.Int
	Owner           common.Address
	Metadata        string
	FromTimestamp   *big.Int
	ToTimestamp     *big.Int
	SuccessfulBidId *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "allPosts", arg0)

	outstruct := new(struct {
		PostId          *big.Int
		Owner           common.Address
		Metadata        string
		FromTimestamp   *big.Int
		ToTimestamp     *big.Int
		SuccessfulBidId *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PostId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Metadata = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.FromTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ToTimestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.SuccessfulBidId = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AllPosts is a free data retrieval call binding the contract method 0x718ce2bc.
//
// Solidity: function allPosts(uint256 ) view returns(uint256 postId, address owner, string metadata, uint256 fromTimestamp, uint256 toTimestamp, uint256 successfulBidId)
func (_Contracts *ContractsSession) AllPosts(arg0 *big.Int) (struct {
	PostId          *big.Int
	Owner           common.Address
	Metadata        string
	FromTimestamp   *big.Int
	ToTimestamp     *big.Int
	SuccessfulBidId *big.Int
}, error) {
	return _Contracts.Contract.AllPosts(&_Contracts.CallOpts, arg0)
}

// AllPosts is a free data retrieval call binding the contract method 0x718ce2bc.
//
// Solidity: function allPosts(uint256 ) view returns(uint256 postId, address owner, string metadata, uint256 fromTimestamp, uint256 toTimestamp, uint256 successfulBidId)
func (_Contracts *ContractsCallerSession) AllPosts(arg0 *big.Int) (struct {
	PostId          *big.Int
	Owner           common.Address
	Metadata        string
	FromTimestamp   *big.Int
	ToTimestamp     *big.Int
	SuccessfulBidId *big.Int
}, error) {
	return _Contracts.Contract.AllPosts(&_Contracts.CallOpts, arg0)
}

// BidderInfo is a free data retrieval call binding the contract method 0x0d5daf3b.
//
// Solidity: function bidderInfo(uint256 ) view returns(uint256 bidId, uint256 postId, address sender, uint256 price, string metadata, uint8 status)
func (_Contracts *ContractsCaller) BidderInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	BidId    *big.Int
	PostId   *big.Int
	Sender   common.Address
	Price    *big.Int
	Metadata string
	Status   uint8
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "bidderInfo", arg0)

	outstruct := new(struct {
		BidId    *big.Int
		PostId   *big.Int
		Sender   common.Address
		Price    *big.Int
		Metadata string
		Status   uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BidId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PostId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Sender = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Metadata = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Status = *abi.ConvertType(out[5], new(uint8)).(*uint8)

	return *outstruct, err

}

// BidderInfo is a free data retrieval call binding the contract method 0x0d5daf3b.
//
// Solidity: function bidderInfo(uint256 ) view returns(uint256 bidId, uint256 postId, address sender, uint256 price, string metadata, uint8 status)
func (_Contracts *ContractsSession) BidderInfo(arg0 *big.Int) (struct {
	BidId    *big.Int
	PostId   *big.Int
	Sender   common.Address
	Price    *big.Int
	Metadata string
	Status   uint8
}, error) {
	return _Contracts.Contract.BidderInfo(&_Contracts.CallOpts, arg0)
}

// BidderInfo is a free data retrieval call binding the contract method 0x0d5daf3b.
//
// Solidity: function bidderInfo(uint256 ) view returns(uint256 bidId, uint256 postId, address sender, uint256 price, string metadata, uint8 status)
func (_Contracts *ContractsCallerSession) BidderInfo(arg0 *big.Int) (struct {
	BidId    *big.Int
	PostId   *big.Int
	Sender   common.Address
	Price    *big.Int
	Metadata string
	Status   uint8
}, error) {
	return _Contracts.Contract.BidderInfo(&_Contracts.CallOpts, arg0)
}

// BidderList is a free data retrieval call binding the contract method 0xd3584e84.
//
// Solidity: function bidderList(uint256 postId) view returns(uint256[])
func (_Contracts *ContractsCaller) BidderList(opts *bind.CallOpts, postId *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "bidderList", postId)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BidderList is a free data retrieval call binding the contract method 0xd3584e84.
//
// Solidity: function bidderList(uint256 postId) view returns(uint256[])
func (_Contracts *ContractsSession) BidderList(postId *big.Int) ([]*big.Int, error) {
	return _Contracts.Contract.BidderList(&_Contracts.CallOpts, postId)
}

// BidderList is a free data retrieval call binding the contract method 0xd3584e84.
//
// Solidity: function bidderList(uint256 postId) view returns(uint256[])
func (_Contracts *ContractsCallerSession) BidderList(postId *big.Int) ([]*big.Int, error) {
	return _Contracts.Contract.BidderList(&_Contracts.CallOpts, postId)
}

// Bidders is a free data retrieval call binding the contract method 0xd1c11205.
//
// Solidity: function bidders(uint256 , uint256 ) view returns(uint256)
func (_Contracts *ContractsCaller) Bidders(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "bidders", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Bidders is a free data retrieval call binding the contract method 0xd1c11205.
//
// Solidity: function bidders(uint256 , uint256 ) view returns(uint256)
func (_Contracts *ContractsSession) Bidders(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Contracts.Contract.Bidders(&_Contracts.CallOpts, arg0, arg1)
}

// Bidders is a free data retrieval call binding the contract method 0xd1c11205.
//
// Solidity: function bidders(uint256 , uint256 ) view returns(uint256)
func (_Contracts *ContractsCallerSession) Bidders(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Contracts.Contract.Bidders(&_Contracts.CallOpts, arg0, arg1)
}

// BookedBidIds is a free data retrieval call binding the contract method 0x0e376fda.
//
// Solidity: function bookedBidIds(uint256 ) view returns(uint256)
func (_Contracts *ContractsCaller) BookedBidIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "bookedBidIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BookedBidIds is a free data retrieval call binding the contract method 0x0e376fda.
//
// Solidity: function bookedBidIds(uint256 ) view returns(uint256)
func (_Contracts *ContractsSession) BookedBidIds(arg0 *big.Int) (*big.Int, error) {
	return _Contracts.Contract.BookedBidIds(&_Contracts.CallOpts, arg0)
}

// BookedBidIds is a free data retrieval call binding the contract method 0x0e376fda.
//
// Solidity: function bookedBidIds(uint256 ) view returns(uint256)
func (_Contracts *ContractsCallerSession) BookedBidIds(arg0 *big.Int) (*big.Int, error) {
	return _Contracts.Contract.BookedBidIds(&_Contracts.CallOpts, arg0)
}

// DisplayByMetadata is a free data retrieval call binding the contract method 0x938baab8.
//
// Solidity: function displayByMetadata(address account, string metadata) view returns(string)
func (_Contracts *ContractsCaller) DisplayByMetadata(opts *bind.CallOpts, account common.Address, metadata string) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "displayByMetadata", account, metadata)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DisplayByMetadata is a free data retrieval call binding the contract method 0x938baab8.
//
// Solidity: function displayByMetadata(address account, string metadata) view returns(string)
func (_Contracts *ContractsSession) DisplayByMetadata(account common.Address, metadata string) (string, error) {
	return _Contracts.Contract.DisplayByMetadata(&_Contracts.CallOpts, account, metadata)
}

// DisplayByMetadata is a free data retrieval call binding the contract method 0x938baab8.
//
// Solidity: function displayByMetadata(address account, string metadata) view returns(string)
func (_Contracts *ContractsCallerSession) DisplayByMetadata(account common.Address, metadata string) (string, error) {
	return _Contracts.Contract.DisplayByMetadata(&_Contracts.CallOpts, account, metadata)
}

// DistributionRightAddress is a free data retrieval call binding the contract method 0xac6919a5.
//
// Solidity: function distributionRightAddress() view returns(address)
func (_Contracts *ContractsCaller) DistributionRightAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "distributionRightAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DistributionRightAddress is a free data retrieval call binding the contract method 0xac6919a5.
//
// Solidity: function distributionRightAddress() view returns(address)
func (_Contracts *ContractsSession) DistributionRightAddress() (common.Address, error) {
	return _Contracts.Contract.DistributionRightAddress(&_Contracts.CallOpts)
}

// DistributionRightAddress is a free data retrieval call binding the contract method 0xac6919a5.
//
// Solidity: function distributionRightAddress() view returns(address)
func (_Contracts *ContractsCallerSession) DistributionRightAddress() (common.Address, error) {
	return _Contracts.Contract.DistributionRightAddress(&_Contracts.CallOpts)
}

// Inventories is a free data retrieval call binding the contract method 0x971e178d.
//
// Solidity: function inventories(address , string , uint256 ) view returns(uint256)
func (_Contracts *ContractsCaller) Inventories(opts *bind.CallOpts, arg0 common.Address, arg1 string, arg2 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "inventories", arg0, arg1, arg2)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Inventories is a free data retrieval call binding the contract method 0x971e178d.
//
// Solidity: function inventories(address , string , uint256 ) view returns(uint256)
func (_Contracts *ContractsSession) Inventories(arg0 common.Address, arg1 string, arg2 *big.Int) (*big.Int, error) {
	return _Contracts.Contract.Inventories(&_Contracts.CallOpts, arg0, arg1, arg2)
}

// Inventories is a free data retrieval call binding the contract method 0x971e178d.
//
// Solidity: function inventories(address , string , uint256 ) view returns(uint256)
func (_Contracts *ContractsCallerSession) Inventories(arg0 common.Address, arg1 string, arg2 *big.Int) (*big.Int, error) {
	return _Contracts.Contract.Inventories(&_Contracts.CallOpts, arg0, arg1, arg2)
}

// MediaMetadata is a free data retrieval call binding the contract method 0x04caebc0.
//
// Solidity: function mediaMetadata(address , uint256 ) view returns(string)
func (_Contracts *ContractsCaller) MediaMetadata(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "mediaMetadata", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MediaMetadata is a free data retrieval call binding the contract method 0x04caebc0.
//
// Solidity: function mediaMetadata(address , uint256 ) view returns(string)
func (_Contracts *ContractsSession) MediaMetadata(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _Contracts.Contract.MediaMetadata(&_Contracts.CallOpts, arg0, arg1)
}

// MediaMetadata is a free data retrieval call binding the contract method 0x04caebc0.
//
// Solidity: function mediaMetadata(address , uint256 ) view returns(string)
func (_Contracts *ContractsCallerSession) MediaMetadata(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _Contracts.Contract.MediaMetadata(&_Contracts.CallOpts, arg0, arg1)
}

// MetadataList is a free data retrieval call binding the contract method 0x7859427a.
//
// Solidity: function metadataList() view returns(string[])
func (_Contracts *ContractsCaller) MetadataList(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "metadataList")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// MetadataList is a free data retrieval call binding the contract method 0x7859427a.
//
// Solidity: function metadataList() view returns(string[])
func (_Contracts *ContractsSession) MetadataList() ([]string, error) {
	return _Contracts.Contract.MetadataList(&_Contracts.CallOpts)
}

// MetadataList is a free data retrieval call binding the contract method 0x7859427a.
//
// Solidity: function metadataList() view returns(string[])
func (_Contracts *ContractsCallerSession) MetadataList() ([]string, error) {
	return _Contracts.Contract.MetadataList(&_Contracts.CallOpts)
}

// NameRegistryAddress is a free data retrieval call binding the contract method 0x27b7a2f5.
//
// Solidity: function nameRegistryAddress() view returns(address)
func (_Contracts *ContractsCaller) NameRegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "nameRegistryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NameRegistryAddress is a free data retrieval call binding the contract method 0x27b7a2f5.
//
// Solidity: function nameRegistryAddress() view returns(address)
func (_Contracts *ContractsSession) NameRegistryAddress() (common.Address, error) {
	return _Contracts.Contract.NameRegistryAddress(&_Contracts.CallOpts)
}

// NameRegistryAddress is a free data retrieval call binding the contract method 0x27b7a2f5.
//
// Solidity: function nameRegistryAddress() view returns(address)
func (_Contracts *ContractsCallerSession) NameRegistryAddress() (common.Address, error) {
	return _Contracts.Contract.NameRegistryAddress(&_Contracts.CallOpts)
}

// NextBidId is a free data retrieval call binding the contract method 0xdc269049.
//
// Solidity: function nextBidId() view returns(uint256)
func (_Contracts *ContractsCaller) NextBidId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "nextBidId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextBidId is a free data retrieval call binding the contract method 0xdc269049.
//
// Solidity: function nextBidId() view returns(uint256)
func (_Contracts *ContractsSession) NextBidId() (*big.Int, error) {
	return _Contracts.Contract.NextBidId(&_Contracts.CallOpts)
}

// NextBidId is a free data retrieval call binding the contract method 0xdc269049.
//
// Solidity: function nextBidId() view returns(uint256)
func (_Contracts *ContractsCallerSession) NextBidId() (*big.Int, error) {
	return _Contracts.Contract.NextBidId(&_Contracts.CallOpts)
}

// NextPostId is a free data retrieval call binding the contract method 0x5c09bd98.
//
// Solidity: function nextPostId() view returns(uint256)
func (_Contracts *ContractsCaller) NextPostId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "nextPostId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextPostId is a free data retrieval call binding the contract method 0x5c09bd98.
//
// Solidity: function nextPostId() view returns(uint256)
func (_Contracts *ContractsSession) NextPostId() (*big.Int, error) {
	return _Contracts.Contract.NextPostId(&_Contracts.CallOpts)
}

// NextPostId is a free data retrieval call binding the contract method 0x5c09bd98.
//
// Solidity: function nextPostId() view returns(uint256)
func (_Contracts *ContractsCallerSession) NextPostId() (*big.Int, error) {
	return _Contracts.Contract.NextPostId(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Contracts *ContractsCallerSession) Owner() (common.Address, error) {
	return _Contracts.Contract.Owner(&_Contracts.CallOpts)
}

// PostOwnerPoolAddress is a free data retrieval call binding the contract method 0xef95e90f.
//
// Solidity: function postOwnerPoolAddress() view returns(address)
func (_Contracts *ContractsCaller) PostOwnerPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "postOwnerPoolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PostOwnerPoolAddress is a free data retrieval call binding the contract method 0xef95e90f.
//
// Solidity: function postOwnerPoolAddress() view returns(address)
func (_Contracts *ContractsSession) PostOwnerPoolAddress() (common.Address, error) {
	return _Contracts.Contract.PostOwnerPoolAddress(&_Contracts.CallOpts)
}

// PostOwnerPoolAddress is a free data retrieval call binding the contract method 0xef95e90f.
//
// Solidity: function postOwnerPoolAddress() view returns(address)
func (_Contracts *ContractsCallerSession) PostOwnerPoolAddress() (common.Address, error) {
	return _Contracts.Contract.PostOwnerPoolAddress(&_Contracts.CallOpts)
}

// ReservedRightAddress is a free data retrieval call binding the contract method 0x29a594b5.
//
// Solidity: function reservedRightAddress() view returns(address)
func (_Contracts *ContractsCaller) ReservedRightAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "reservedRightAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReservedRightAddress is a free data retrieval call binding the contract method 0x29a594b5.
//
// Solidity: function reservedRightAddress() view returns(address)
func (_Contracts *ContractsSession) ReservedRightAddress() (common.Address, error) {
	return _Contracts.Contract.ReservedRightAddress(&_Contracts.CallOpts)
}

// ReservedRightAddress is a free data retrieval call binding the contract method 0x29a594b5.
//
// Solidity: function reservedRightAddress() view returns(address)
func (_Contracts *ContractsCallerSession) ReservedRightAddress() (common.Address, error) {
	return _Contracts.Contract.ReservedRightAddress(&_Contracts.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_Contracts *ContractsCaller) VaultAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "vaultAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_Contracts *ContractsSession) VaultAddress() (common.Address, error) {
	return _Contracts.Contract.VaultAddress(&_Contracts.CallOpts)
}

// VaultAddress is a free data retrieval call binding the contract method 0x430bf08a.
//
// Solidity: function vaultAddress() view returns(address)
func (_Contracts *ContractsCallerSession) VaultAddress() (common.Address, error) {
	return _Contracts.Contract.VaultAddress(&_Contracts.CallOpts)
}

// Accept is a paid mutator transaction binding the contract method 0x19b05f49.
//
// Solidity: function accept(uint256 postId) returns()
func (_Contracts *ContractsTransactor) Accept(opts *bind.TransactOpts, postId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "accept", postId)
}

// Accept is a paid mutator transaction binding the contract method 0x19b05f49.
//
// Solidity: function accept(uint256 postId) returns()
func (_Contracts *ContractsSession) Accept(postId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Accept(&_Contracts.TransactOpts, postId)
}

// Accept is a paid mutator transaction binding the contract method 0x19b05f49.
//
// Solidity: function accept(uint256 postId) returns()
func (_Contracts *ContractsTransactorSession) Accept(postId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Accept(&_Contracts.TransactOpts, postId)
}

// Bid1 is a paid mutator transaction binding the contract method 0x0fc55fd1.
//
// Solidity: function bid(uint256 postId, string metadata) payable returns()
func (_Contracts *ContractsTransactor) Bid1(opts *bind.TransactOpts, postId *big.Int, metadata string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "bid", postId, metadata)
}

// Bid1 is a paid mutator transaction binding the contract method 0x0fc55fd1.
//
// Solidity: function bid(uint256 postId, string metadata) payable returns()
func (_Contracts *ContractsSession) Bid1(postId *big.Int, metadata string) (*types.Transaction, error) {
	return _Contracts.Contract.Bid1(&_Contracts.TransactOpts, postId, metadata)
}

// Bid1 is a paid mutator transaction binding the contract method 0x0fc55fd1.
//
// Solidity: function bid(uint256 postId, string metadata) payable returns()
func (_Contracts *ContractsTransactorSession) Bid1(postId *big.Int, metadata string) (*types.Transaction, error) {
	return _Contracts.Contract.Bid1(&_Contracts.TransactOpts, postId, metadata)
}

// Book is a paid mutator transaction binding the contract method 0x1116fd04.
//
// Solidity: function book(uint256 postId) payable returns()
func (_Contracts *ContractsTransactor) Book(opts *bind.TransactOpts, postId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "book", postId)
}

// Book is a paid mutator transaction binding the contract method 0x1116fd04.
//
// Solidity: function book(uint256 postId) payable returns()
func (_Contracts *ContractsSession) Book(postId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Book(&_Contracts.TransactOpts, postId)
}

// Book is a paid mutator transaction binding the contract method 0x1116fd04.
//
// Solidity: function book(uint256 postId) payable returns()
func (_Contracts *ContractsTransactorSession) Book(postId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Book(&_Contracts.TransactOpts, postId)
}

// Call is a paid mutator transaction binding the contract method 0x2b096926.
//
// Solidity: function call(uint256 bidId) returns()
func (_Contracts *ContractsTransactor) Call(opts *bind.TransactOpts, bidId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "call", bidId)
}

// Call is a paid mutator transaction binding the contract method 0x2b096926.
//
// Solidity: function call(uint256 bidId) returns()
func (_Contracts *ContractsSession) Call(bidId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Call(&_Contracts.TransactOpts, bidId)
}

// Call is a paid mutator transaction binding the contract method 0x2b096926.
//
// Solidity: function call(uint256 bidId) returns()
func (_Contracts *ContractsTransactorSession) Call(bidId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Call(&_Contracts.TransactOpts, bidId)
}

// Close is a paid mutator transaction binding the contract method 0x0aebeb4e.
//
// Solidity: function close(uint256 bidId) returns()
func (_Contracts *ContractsTransactor) Close(opts *bind.TransactOpts, bidId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "close", bidId)
}

// Close is a paid mutator transaction binding the contract method 0x0aebeb4e.
//
// Solidity: function close(uint256 bidId) returns()
func (_Contracts *ContractsSession) Close(bidId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Close(&_Contracts.TransactOpts, bidId)
}

// Close is a paid mutator transaction binding the contract method 0x0aebeb4e.
//
// Solidity: function close(uint256 bidId) returns()
func (_Contracts *ContractsTransactorSession) Close(bidId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Close(&_Contracts.TransactOpts, bidId)
}

// Deny is a paid mutator transaction binding the contract method 0x56f66282.
//
// Solidity: function deny(uint256 postId) returns()
func (_Contracts *ContractsTransactor) Deny(opts *bind.TransactOpts, postId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "deny", postId)
}

// Deny is a paid mutator transaction binding the contract method 0x56f66282.
//
// Solidity: function deny(uint256 postId) returns()
func (_Contracts *ContractsSession) Deny(postId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Deny(&_Contracts.TransactOpts, postId)
}

// Deny is a paid mutator transaction binding the contract method 0x56f66282.
//
// Solidity: function deny(uint256 postId) returns()
func (_Contracts *ContractsTransactorSession) Deny(postId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Deny(&_Contracts.TransactOpts, postId)
}

// NewPost is a paid mutator transaction binding the contract method 0x8f50b143.
//
// Solidity: function newPost(string metadata, uint256 fromTimestamp, uint256 toTimestamp) returns()
func (_Contracts *ContractsTransactor) NewPost(opts *bind.TransactOpts, metadata string, fromTimestamp *big.Int, toTimestamp *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "newPost", metadata, fromTimestamp, toTimestamp)
}

// NewPost is a paid mutator transaction binding the contract method 0x8f50b143.
//
// Solidity: function newPost(string metadata, uint256 fromTimestamp, uint256 toTimestamp) returns()
func (_Contracts *ContractsSession) NewPost(metadata string, fromTimestamp *big.Int, toTimestamp *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.NewPost(&_Contracts.TransactOpts, metadata, fromTimestamp, toTimestamp)
}

// NewPost is a paid mutator transaction binding the contract method 0x8f50b143.
//
// Solidity: function newPost(string metadata, uint256 fromTimestamp, uint256 toTimestamp) returns()
func (_Contracts *ContractsTransactorSession) NewPost(metadata string, fromTimestamp *big.Int, toTimestamp *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.NewPost(&_Contracts.TransactOpts, metadata, fromTimestamp, toTimestamp)
}

// Propose is a paid mutator transaction binding the contract method 0xd4f6b5ec.
//
// Solidity: function propose(uint256 postId, string metadata) returns()
func (_Contracts *ContractsTransactor) Propose(opts *bind.TransactOpts, postId *big.Int, metadata string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "propose", postId, metadata)
}

// Propose is a paid mutator transaction binding the contract method 0xd4f6b5ec.
//
// Solidity: function propose(uint256 postId, string metadata) returns()
func (_Contracts *ContractsSession) Propose(postId *big.Int, metadata string) (*types.Transaction, error) {
	return _Contracts.Contract.Propose(&_Contracts.TransactOpts, postId, metadata)
}

// Propose is a paid mutator transaction binding the contract method 0xd4f6b5ec.
//
// Solidity: function propose(uint256 postId, string metadata) returns()
func (_Contracts *ContractsTransactorSession) Propose(postId *big.Int, metadata string) (*types.Transaction, error) {
	return _Contracts.Contract.Propose(&_Contracts.TransactOpts, postId, metadata)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 bidId) returns()
func (_Contracts *ContractsTransactor) Refund(opts *bind.TransactOpts, bidId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "refund", bidId)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 bidId) returns()
func (_Contracts *ContractsSession) Refund(bidId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Refund(&_Contracts.TransactOpts, bidId)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 bidId) returns()
func (_Contracts *ContractsTransactorSession) Refund(bidId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Refund(&_Contracts.TransactOpts, bidId)
}

// ContractsAcceptIterator is returned from FilterAccept and is used to iterate over the raw logs and unpacked data for Accept events raised by the Contracts contract.
type ContractsAcceptIterator struct {
	Event *ContractsAccept // Event containing the contract specifics and raw log

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
func (it *ContractsAcceptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsAccept)
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
		it.Event = new(ContractsAccept)
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
func (it *ContractsAcceptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsAcceptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsAccept represents a Accept event raised by the Contracts contract.
type ContractsAccept struct {
	PostId *big.Int
	BidId  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAccept is a free log retrieval operation binding the contract event 0x66b40296db692bd2f393bda18cf3f9d917966186767279128a62afeb90e43b39.
//
// Solidity: event Accept(uint256 postId, uint256 bidId)
func (_Contracts *ContractsFilterer) FilterAccept(opts *bind.FilterOpts) (*ContractsAcceptIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Accept")
	if err != nil {
		return nil, err
	}
	return &ContractsAcceptIterator{contract: _Contracts.contract, event: "Accept", logs: logs, sub: sub}, nil
}

// WatchAccept is a free log subscription operation binding the contract event 0x66b40296db692bd2f393bda18cf3f9d917966186767279128a62afeb90e43b39.
//
// Solidity: event Accept(uint256 postId, uint256 bidId)
func (_Contracts *ContractsFilterer) WatchAccept(opts *bind.WatchOpts, sink chan<- *ContractsAccept) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Accept")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsAccept)
				if err := _Contracts.contract.UnpackLog(event, "Accept", log); err != nil {
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

// ParseAccept is a log parse operation binding the contract event 0x66b40296db692bd2f393bda18cf3f9d917966186767279128a62afeb90e43b39.
//
// Solidity: event Accept(uint256 postId, uint256 bidId)
func (_Contracts *ContractsFilterer) ParseAccept(log types.Log) (*ContractsAccept, error) {
	event := new(ContractsAccept)
	if err := _Contracts.contract.UnpackLog(event, "Accept", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsBidIterator is returned from FilterBid and is used to iterate over the raw logs and unpacked data for Bid events raised by the Contracts contract.
type ContractsBidIterator struct {
	Event *ContractsBid // Event containing the contract specifics and raw log

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
func (it *ContractsBidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsBid)
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
		it.Event = new(ContractsBid)
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
func (it *ContractsBidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsBidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsBid represents a Bid event raised by the Contracts contract.
type ContractsBid struct {
	BidId    *big.Int
	PostId   *big.Int
	Sender   common.Address
	Price    *big.Int
	Metadata string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBid is a free log retrieval operation binding the contract event 0x01662ff7392d9961759cf40cab7a64e5ad660ba64a9df944343ce23b2b94c447.
//
// Solidity: event Bid(uint256 bidId, uint256 postId, address sender, uint256 price, string metadata)
func (_Contracts *ContractsFilterer) FilterBid(opts *bind.FilterOpts) (*ContractsBidIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Bid")
	if err != nil {
		return nil, err
	}
	return &ContractsBidIterator{contract: _Contracts.contract, event: "Bid", logs: logs, sub: sub}, nil
}

// WatchBid is a free log subscription operation binding the contract event 0x01662ff7392d9961759cf40cab7a64e5ad660ba64a9df944343ce23b2b94c447.
//
// Solidity: event Bid(uint256 bidId, uint256 postId, address sender, uint256 price, string metadata)
func (_Contracts *ContractsFilterer) WatchBid(opts *bind.WatchOpts, sink chan<- *ContractsBid) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Bid")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsBid)
				if err := _Contracts.contract.UnpackLog(event, "Bid", log); err != nil {
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

// ParseBid is a log parse operation binding the contract event 0x01662ff7392d9961759cf40cab7a64e5ad660ba64a9df944343ce23b2b94c447.
//
// Solidity: event Bid(uint256 bidId, uint256 postId, address sender, uint256 price, string metadata)
func (_Contracts *ContractsFilterer) ParseBid(log types.Log) (*ContractsBid, error) {
	event := new(ContractsBid)
	if err := _Contracts.contract.UnpackLog(event, "Bid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsBookIterator is returned from FilterBook and is used to iterate over the raw logs and unpacked data for Book events raised by the Contracts contract.
type ContractsBookIterator struct {
	Event *ContractsBook // Event containing the contract specifics and raw log

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
func (it *ContractsBookIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsBook)
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
		it.Event = new(ContractsBook)
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
func (it *ContractsBookIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsBookIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsBook represents a Book event raised by the Contracts contract.
type ContractsBook struct {
	BidId  *big.Int
	PostId *big.Int
	Sender common.Address
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBook is a free log retrieval operation binding the contract event 0xa9dfceace4f13fa67be6499ce107d4e132c418d2cdeb93219a08c7f9c7fa634a.
//
// Solidity: event Book(uint256 bidId, uint256 postId, address sender, uint256 price)
func (_Contracts *ContractsFilterer) FilterBook(opts *bind.FilterOpts) (*ContractsBookIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Book")
	if err != nil {
		return nil, err
	}
	return &ContractsBookIterator{contract: _Contracts.contract, event: "Book", logs: logs, sub: sub}, nil
}

// WatchBook is a free log subscription operation binding the contract event 0xa9dfceace4f13fa67be6499ce107d4e132c418d2cdeb93219a08c7f9c7fa634a.
//
// Solidity: event Book(uint256 bidId, uint256 postId, address sender, uint256 price)
func (_Contracts *ContractsFilterer) WatchBook(opts *bind.WatchOpts, sink chan<- *ContractsBook) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Book")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsBook)
				if err := _Contracts.contract.UnpackLog(event, "Book", log); err != nil {
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

// ParseBook is a log parse operation binding the contract event 0xa9dfceace4f13fa67be6499ce107d4e132c418d2cdeb93219a08c7f9c7fa634a.
//
// Solidity: event Book(uint256 bidId, uint256 postId, address sender, uint256 price)
func (_Contracts *ContractsFilterer) ParseBook(log types.Log) (*ContractsBook, error) {
	event := new(ContractsBook)
	if err := _Contracts.contract.UnpackLog(event, "Book", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsCallIterator is returned from FilterCall and is used to iterate over the raw logs and unpacked data for Call events raised by the Contracts contract.
type ContractsCallIterator struct {
	Event *ContractsCall // Event containing the contract specifics and raw log

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
func (it *ContractsCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCall)
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
		it.Event = new(ContractsCall)
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
func (it *ContractsCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCall represents a Call event raised by the Contracts contract.
type ContractsCall struct {
	BidId  *big.Int
	PostId *big.Int
	Sender common.Address
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCall is a free log retrieval operation binding the contract event 0x70c21df4a109699ac5cbee04e25054d865a14a532529a731a11afcd25ed1914a.
//
// Solidity: event Call(uint256 bidId, uint256 postId, address sender, uint256 price)
func (_Contracts *ContractsFilterer) FilterCall(opts *bind.FilterOpts) (*ContractsCallIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Call")
	if err != nil {
		return nil, err
	}
	return &ContractsCallIterator{contract: _Contracts.contract, event: "Call", logs: logs, sub: sub}, nil
}

// WatchCall is a free log subscription operation binding the contract event 0x70c21df4a109699ac5cbee04e25054d865a14a532529a731a11afcd25ed1914a.
//
// Solidity: event Call(uint256 bidId, uint256 postId, address sender, uint256 price)
func (_Contracts *ContractsFilterer) WatchCall(opts *bind.WatchOpts, sink chan<- *ContractsCall) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Call")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCall)
				if err := _Contracts.contract.UnpackLog(event, "Call", log); err != nil {
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

// ParseCall is a log parse operation binding the contract event 0x70c21df4a109699ac5cbee04e25054d865a14a532529a731a11afcd25ed1914a.
//
// Solidity: event Call(uint256 bidId, uint256 postId, address sender, uint256 price)
func (_Contracts *ContractsFilterer) ParseCall(log types.Log) (*ContractsCall, error) {
	event := new(ContractsCall)
	if err := _Contracts.contract.UnpackLog(event, "Call", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsCloseIterator is returned from FilterClose and is used to iterate over the raw logs and unpacked data for Close events raised by the Contracts contract.
type ContractsCloseIterator struct {
	Event *ContractsClose // Event containing the contract specifics and raw log

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
func (it *ContractsCloseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsClose)
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
		it.Event = new(ContractsClose)
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
func (it *ContractsCloseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCloseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsClose represents a Close event raised by the Contracts contract.
type ContractsClose struct {
	BitId            *big.Int
	PostId           *big.Int
	SuccessfulBidder common.Address
	Price            *big.Int
	Metadata         string
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterClose is a free log retrieval operation binding the contract event 0xf763d6ab2a4f1accf9c4fca103d30b90f90db83afeaad6b67481601184b1690d.
//
// Solidity: event Close(uint256 bitId, uint256 postId, address successfulBidder, uint256 price, string metadata)
func (_Contracts *ContractsFilterer) FilterClose(opts *bind.FilterOpts) (*ContractsCloseIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Close")
	if err != nil {
		return nil, err
	}
	return &ContractsCloseIterator{contract: _Contracts.contract, event: "Close", logs: logs, sub: sub}, nil
}

// WatchClose is a free log subscription operation binding the contract event 0xf763d6ab2a4f1accf9c4fca103d30b90f90db83afeaad6b67481601184b1690d.
//
// Solidity: event Close(uint256 bitId, uint256 postId, address successfulBidder, uint256 price, string metadata)
func (_Contracts *ContractsFilterer) WatchClose(opts *bind.WatchOpts, sink chan<- *ContractsClose) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Close")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsClose)
				if err := _Contracts.contract.UnpackLog(event, "Close", log); err != nil {
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

// ParseClose is a log parse operation binding the contract event 0xf763d6ab2a4f1accf9c4fca103d30b90f90db83afeaad6b67481601184b1690d.
//
// Solidity: event Close(uint256 bitId, uint256 postId, address successfulBidder, uint256 price, string metadata)
func (_Contracts *ContractsFilterer) ParseClose(log types.Log) (*ContractsClose, error) {
	event := new(ContractsClose)
	if err := _Contracts.contract.UnpackLog(event, "Close", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsDenyIterator is returned from FilterDeny and is used to iterate over the raw logs and unpacked data for Deny events raised by the Contracts contract.
type ContractsDenyIterator struct {
	Event *ContractsDeny // Event containing the contract specifics and raw log

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
func (it *ContractsDenyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsDeny)
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
		it.Event = new(ContractsDeny)
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
func (it *ContractsDenyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsDenyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsDeny represents a Deny event raised by the Contracts contract.
type ContractsDeny struct {
	BidId  *big.Int
	PostId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeny is a free log retrieval operation binding the contract event 0xb4f2cd71a07f3c432393e99783b21b73c1ccf8f528fd95885166c9f8a83669d6.
//
// Solidity: event Deny(uint256 bidId, uint256 postId)
func (_Contracts *ContractsFilterer) FilterDeny(opts *bind.FilterOpts) (*ContractsDenyIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Deny")
	if err != nil {
		return nil, err
	}
	return &ContractsDenyIterator{contract: _Contracts.contract, event: "Deny", logs: logs, sub: sub}, nil
}

// WatchDeny is a free log subscription operation binding the contract event 0xb4f2cd71a07f3c432393e99783b21b73c1ccf8f528fd95885166c9f8a83669d6.
//
// Solidity: event Deny(uint256 bidId, uint256 postId)
func (_Contracts *ContractsFilterer) WatchDeny(opts *bind.WatchOpts, sink chan<- *ContractsDeny) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Deny")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsDeny)
				if err := _Contracts.contract.UnpackLog(event, "Deny", log); err != nil {
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

// ParseDeny is a log parse operation binding the contract event 0xb4f2cd71a07f3c432393e99783b21b73c1ccf8f528fd95885166c9f8a83669d6.
//
// Solidity: event Deny(uint256 bidId, uint256 postId)
func (_Contracts *ContractsFilterer) ParseDeny(log types.Log) (*ContractsDeny, error) {
	event := new(ContractsDeny)
	if err := _Contracts.contract.UnpackLog(event, "Deny", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsNewPostIterator is returned from FilterNewPost and is used to iterate over the raw logs and unpacked data for NewPost events raised by the Contracts contract.
type ContractsNewPostIterator struct {
	Event *ContractsNewPost // Event containing the contract specifics and raw log

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
func (it *ContractsNewPostIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsNewPost)
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
		it.Event = new(ContractsNewPost)
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
func (it *ContractsNewPostIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsNewPostIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsNewPost represents a NewPost event raised by the Contracts contract.
type ContractsNewPost struct {
	PostId        *big.Int
	Owner         common.Address
	Metadata      string
	FromTimestamp *big.Int
	ToTimestamp   *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewPost is a free log retrieval operation binding the contract event 0xe535a02a01920b38061e7785d5f3d4688d895a51cc21bb9e09b59a11d3befcfb.
//
// Solidity: event NewPost(uint256 postId, address owner, string metadata, uint256 fromTimestamp, uint256 toTimestamp)
func (_Contracts *ContractsFilterer) FilterNewPost(opts *bind.FilterOpts) (*ContractsNewPostIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "NewPost")
	if err != nil {
		return nil, err
	}
	return &ContractsNewPostIterator{contract: _Contracts.contract, event: "NewPost", logs: logs, sub: sub}, nil
}

// WatchNewPost is a free log subscription operation binding the contract event 0xe535a02a01920b38061e7785d5f3d4688d895a51cc21bb9e09b59a11d3befcfb.
//
// Solidity: event NewPost(uint256 postId, address owner, string metadata, uint256 fromTimestamp, uint256 toTimestamp)
func (_Contracts *ContractsFilterer) WatchNewPost(opts *bind.WatchOpts, sink chan<- *ContractsNewPost) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "NewPost")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsNewPost)
				if err := _Contracts.contract.UnpackLog(event, "NewPost", log); err != nil {
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

// ParseNewPost is a log parse operation binding the contract event 0xe535a02a01920b38061e7785d5f3d4688d895a51cc21bb9e09b59a11d3befcfb.
//
// Solidity: event NewPost(uint256 postId, address owner, string metadata, uint256 fromTimestamp, uint256 toTimestamp)
func (_Contracts *ContractsFilterer) ParseNewPost(log types.Log) (*ContractsNewPost, error) {
	event := new(ContractsNewPost)
	if err := _Contracts.contract.UnpackLog(event, "NewPost", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsProposeIterator is returned from FilterPropose and is used to iterate over the raw logs and unpacked data for Propose events raised by the Contracts contract.
type ContractsProposeIterator struct {
	Event *ContractsPropose // Event containing the contract specifics and raw log

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
func (it *ContractsProposeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsPropose)
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
		it.Event = new(ContractsPropose)
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
func (it *ContractsProposeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsProposeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsPropose represents a Propose event raised by the Contracts contract.
type ContractsPropose struct {
	BidId    *big.Int
	PostId   *big.Int
	Metadata string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPropose is a free log retrieval operation binding the contract event 0x3210cee2deb4838dde3ae606fd6b2e36543e05f2629f4d6fe336fc572c06a335.
//
// Solidity: event Propose(uint256 bidId, uint256 postId, string metadata)
func (_Contracts *ContractsFilterer) FilterPropose(opts *bind.FilterOpts) (*ContractsProposeIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Propose")
	if err != nil {
		return nil, err
	}
	return &ContractsProposeIterator{contract: _Contracts.contract, event: "Propose", logs: logs, sub: sub}, nil
}

// WatchPropose is a free log subscription operation binding the contract event 0x3210cee2deb4838dde3ae606fd6b2e36543e05f2629f4d6fe336fc572c06a335.
//
// Solidity: event Propose(uint256 bidId, uint256 postId, string metadata)
func (_Contracts *ContractsFilterer) WatchPropose(opts *bind.WatchOpts, sink chan<- *ContractsPropose) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Propose")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsPropose)
				if err := _Contracts.contract.UnpackLog(event, "Propose", log); err != nil {
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

// ParsePropose is a log parse operation binding the contract event 0x3210cee2deb4838dde3ae606fd6b2e36543e05f2629f4d6fe336fc572c06a335.
//
// Solidity: event Propose(uint256 bidId, uint256 postId, string metadata)
func (_Contracts *ContractsFilterer) ParsePropose(log types.Log) (*ContractsPropose, error) {
	event := new(ContractsPropose)
	if err := _Contracts.contract.UnpackLog(event, "Propose", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsRefundIterator is returned from FilterRefund and is used to iterate over the raw logs and unpacked data for Refund events raised by the Contracts contract.
type ContractsRefundIterator struct {
	Event *ContractsRefund // Event containing the contract specifics and raw log

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
func (it *ContractsRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsRefund)
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
		it.Event = new(ContractsRefund)
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
func (it *ContractsRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsRefund represents a Refund event raised by the Contracts contract.
type ContractsRefund struct {
	BitId  *big.Int
	PostId *big.Int
	Sender common.Address
	Price  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRefund is a free log retrieval operation binding the contract event 0x303a7a5ed564cf1b7e8ec302e0d290e0c69d9785d1e5d43db651d7995f0bc019.
//
// Solidity: event Refund(uint256 bitId, uint256 postId, address sender, uint256 price)
func (_Contracts *ContractsFilterer) FilterRefund(opts *bind.FilterOpts) (*ContractsRefundIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return &ContractsRefundIterator{contract: _Contracts.contract, event: "Refund", logs: logs, sub: sub}, nil
}

// WatchRefund is a free log subscription operation binding the contract event 0x303a7a5ed564cf1b7e8ec302e0d290e0c69d9785d1e5d43db651d7995f0bc019.
//
// Solidity: event Refund(uint256 bitId, uint256 postId, address sender, uint256 price)
func (_Contracts *ContractsFilterer) WatchRefund(opts *bind.WatchOpts, sink chan<- *ContractsRefund) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Refund")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsRefund)
				if err := _Contracts.contract.UnpackLog(event, "Refund", log); err != nil {
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

// ParseRefund is a log parse operation binding the contract event 0x303a7a5ed564cf1b7e8ec302e0d290e0c69d9785d1e5d43db651d7995f0bc019.
//
// Solidity: event Refund(uint256 bitId, uint256 postId, address sender, uint256 price)
func (_Contracts *ContractsFilterer) ParseRefund(log types.Log) (*ContractsRefund, error) {
	event := new(ContractsRefund)
	if err := _Contracts.contract.UnpackLog(event, "Refund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
