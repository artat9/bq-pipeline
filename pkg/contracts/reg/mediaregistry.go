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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nameRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"adPoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"mediaEOA\",\"type\":\"address\"}],\"name\":\"addMedia\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"allAccounts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mediaEOA\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"created\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eventEmitterAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mediaFactoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mediaRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nameRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mediaEOA\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"updateMedia\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// AdPoolAddress is a free data retrieval call binding the contract method 0xc9cd71b5.
//
// Solidity: function adPoolAddress() view returns(address)
func (_Contracts *ContractsCaller) AdPoolAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "adPoolAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AdPoolAddress is a free data retrieval call binding the contract method 0xc9cd71b5.
//
// Solidity: function adPoolAddress() view returns(address)
func (_Contracts *ContractsSession) AdPoolAddress() (common.Address, error) {
	return _Contracts.Contract.AdPoolAddress(&_Contracts.CallOpts)
}

// AdPoolAddress is a free data retrieval call binding the contract method 0xc9cd71b5.
//
// Solidity: function adPoolAddress() view returns(address)
func (_Contracts *ContractsCallerSession) AdPoolAddress() (common.Address, error) {
	return _Contracts.Contract.AdPoolAddress(&_Contracts.CallOpts)
}

// AllAccounts is a free data retrieval call binding the contract method 0x0df8e9fe.
//
// Solidity: function allAccounts(address ) view returns(address proxy, address mediaEOA, string metadata)
func (_Contracts *ContractsCaller) AllAccounts(opts *bind.CallOpts, arg0 common.Address) (struct {
	Proxy    common.Address
	MediaEOA common.Address
	Metadata string
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "allAccounts", arg0)

	outstruct := new(struct {
		Proxy    common.Address
		MediaEOA common.Address
		Metadata string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Proxy = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.MediaEOA = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Metadata = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// AllAccounts is a free data retrieval call binding the contract method 0x0df8e9fe.
//
// Solidity: function allAccounts(address ) view returns(address proxy, address mediaEOA, string metadata)
func (_Contracts *ContractsSession) AllAccounts(arg0 common.Address) (struct {
	Proxy    common.Address
	MediaEOA common.Address
	Metadata string
}, error) {
	return _Contracts.Contract.AllAccounts(&_Contracts.CallOpts, arg0)
}

// AllAccounts is a free data retrieval call binding the contract method 0x0df8e9fe.
//
// Solidity: function allAccounts(address ) view returns(address proxy, address mediaEOA, string metadata)
func (_Contracts *ContractsCallerSession) AllAccounts(arg0 common.Address) (struct {
	Proxy    common.Address
	MediaEOA common.Address
	Metadata string
}, error) {
	return _Contracts.Contract.AllAccounts(&_Contracts.CallOpts, arg0)
}

// Created is a free data retrieval call binding the contract method 0xd42efd83.
//
// Solidity: function created(address proxy) view returns(bool)
func (_Contracts *ContractsCaller) Created(opts *bind.CallOpts, proxy common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "created", proxy)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Created is a free data retrieval call binding the contract method 0xd42efd83.
//
// Solidity: function created(address proxy) view returns(bool)
func (_Contracts *ContractsSession) Created(proxy common.Address) (bool, error) {
	return _Contracts.Contract.Created(&_Contracts.CallOpts, proxy)
}

// Created is a free data retrieval call binding the contract method 0xd42efd83.
//
// Solidity: function created(address proxy) view returns(bool)
func (_Contracts *ContractsCallerSession) Created(proxy common.Address) (bool, error) {
	return _Contracts.Contract.Created(&_Contracts.CallOpts, proxy)
}

// EventEmitterAddress is a free data retrieval call binding the contract method 0xf0726291.
//
// Solidity: function eventEmitterAddress() view returns(address)
func (_Contracts *ContractsCaller) EventEmitterAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "eventEmitterAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EventEmitterAddress is a free data retrieval call binding the contract method 0xf0726291.
//
// Solidity: function eventEmitterAddress() view returns(address)
func (_Contracts *ContractsSession) EventEmitterAddress() (common.Address, error) {
	return _Contracts.Contract.EventEmitterAddress(&_Contracts.CallOpts)
}

// EventEmitterAddress is a free data retrieval call binding the contract method 0xf0726291.
//
// Solidity: function eventEmitterAddress() view returns(address)
func (_Contracts *ContractsCallerSession) EventEmitterAddress() (common.Address, error) {
	return _Contracts.Contract.EventEmitterAddress(&_Contracts.CallOpts)
}

// MediaFactoryAddress is a free data retrieval call binding the contract method 0x8f6059d6.
//
// Solidity: function mediaFactoryAddress() view returns(address)
func (_Contracts *ContractsCaller) MediaFactoryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "mediaFactoryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MediaFactoryAddress is a free data retrieval call binding the contract method 0x8f6059d6.
//
// Solidity: function mediaFactoryAddress() view returns(address)
func (_Contracts *ContractsSession) MediaFactoryAddress() (common.Address, error) {
	return _Contracts.Contract.MediaFactoryAddress(&_Contracts.CallOpts)
}

// MediaFactoryAddress is a free data retrieval call binding the contract method 0x8f6059d6.
//
// Solidity: function mediaFactoryAddress() view returns(address)
func (_Contracts *ContractsCallerSession) MediaFactoryAddress() (common.Address, error) {
	return _Contracts.Contract.MediaFactoryAddress(&_Contracts.CallOpts)
}

// MediaRegistryAddress is a free data retrieval call binding the contract method 0x6a58767e.
//
// Solidity: function mediaRegistryAddress() view returns(address)
func (_Contracts *ContractsCaller) MediaRegistryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "mediaRegistryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MediaRegistryAddress is a free data retrieval call binding the contract method 0x6a58767e.
//
// Solidity: function mediaRegistryAddress() view returns(address)
func (_Contracts *ContractsSession) MediaRegistryAddress() (common.Address, error) {
	return _Contracts.Contract.MediaRegistryAddress(&_Contracts.CallOpts)
}

// MediaRegistryAddress is a free data retrieval call binding the contract method 0x6a58767e.
//
// Solidity: function mediaRegistryAddress() view returns(address)
func (_Contracts *ContractsCallerSession) MediaRegistryAddress() (common.Address, error) {
	return _Contracts.Contract.MediaRegistryAddress(&_Contracts.CallOpts)
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

// OwnerOf is a free data retrieval call binding the contract method 0x14afd79e.
//
// Solidity: function ownerOf(address proxy) view returns(address)
func (_Contracts *ContractsCaller) OwnerOf(opts *bind.CallOpts, proxy common.Address) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "ownerOf", proxy)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x14afd79e.
//
// Solidity: function ownerOf(address proxy) view returns(address)
func (_Contracts *ContractsSession) OwnerOf(proxy common.Address) (common.Address, error) {
	return _Contracts.Contract.OwnerOf(&_Contracts.CallOpts, proxy)
}

// OwnerOf is a free data retrieval call binding the contract method 0x14afd79e.
//
// Solidity: function ownerOf(address proxy) view returns(address)
func (_Contracts *ContractsCallerSession) OwnerOf(proxy common.Address) (common.Address, error) {
	return _Contracts.Contract.OwnerOf(&_Contracts.CallOpts, proxy)
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

// AddMedia is a paid mutator transaction binding the contract method 0xc1756787.
//
// Solidity: function addMedia(address proxy, string metadata, address mediaEOA) returns()
func (_Contracts *ContractsTransactor) AddMedia(opts *bind.TransactOpts, proxy common.Address, metadata string, mediaEOA common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addMedia", proxy, metadata, mediaEOA)
}

// AddMedia is a paid mutator transaction binding the contract method 0xc1756787.
//
// Solidity: function addMedia(address proxy, string metadata, address mediaEOA) returns()
func (_Contracts *ContractsSession) AddMedia(proxy common.Address, metadata string, mediaEOA common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddMedia(&_Contracts.TransactOpts, proxy, metadata, mediaEOA)
}

// AddMedia is a paid mutator transaction binding the contract method 0xc1756787.
//
// Solidity: function addMedia(address proxy, string metadata, address mediaEOA) returns()
func (_Contracts *ContractsTransactorSession) AddMedia(proxy common.Address, metadata string, mediaEOA common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddMedia(&_Contracts.TransactOpts, proxy, metadata, mediaEOA)
}

// UpdateMedia is a paid mutator transaction binding the contract method 0x179bcff7.
//
// Solidity: function updateMedia(address mediaEOA, string metadata) returns()
func (_Contracts *ContractsTransactor) UpdateMedia(opts *bind.TransactOpts, mediaEOA common.Address, metadata string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateMedia", mediaEOA, metadata)
}

// UpdateMedia is a paid mutator transaction binding the contract method 0x179bcff7.
//
// Solidity: function updateMedia(address mediaEOA, string metadata) returns()
func (_Contracts *ContractsSession) UpdateMedia(mediaEOA common.Address, metadata string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateMedia(&_Contracts.TransactOpts, mediaEOA, metadata)
}

// UpdateMedia is a paid mutator transaction binding the contract method 0x179bcff7.
//
// Solidity: function updateMedia(address mediaEOA, string metadata) returns()
func (_Contracts *ContractsTransactorSession) UpdateMedia(mediaEOA common.Address, metadata string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateMedia(&_Contracts.TransactOpts, mediaEOA, metadata)
}
