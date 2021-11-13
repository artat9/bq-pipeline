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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nameRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractMediaProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"name\":\"CreateProxy\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adPoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eventEmitterAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mediaFactoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mediaRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nameRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mediaEOA\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"accountMetadata\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"initializer\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"saltNonce\",\"type\":\"uint256\"}],\"name\":\"newMedia\",\"outputs\":[{\"internalType\":\"contractMediaProxy\",\"name\":\"proxy\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// NewMedia is a paid mutator transaction binding the contract method 0x7d0f18cb.
//
// Solidity: function newMedia(address mediaEOA, string accountMetadata, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_Contracts *ContractsTransactor) NewMedia(opts *bind.TransactOpts, mediaEOA common.Address, accountMetadata string, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "newMedia", mediaEOA, accountMetadata, initializer, saltNonce)
}

// NewMedia is a paid mutator transaction binding the contract method 0x7d0f18cb.
//
// Solidity: function newMedia(address mediaEOA, string accountMetadata, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_Contracts *ContractsSession) NewMedia(mediaEOA common.Address, accountMetadata string, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.NewMedia(&_Contracts.TransactOpts, mediaEOA, accountMetadata, initializer, saltNonce)
}

// NewMedia is a paid mutator transaction binding the contract method 0x7d0f18cb.
//
// Solidity: function newMedia(address mediaEOA, string accountMetadata, bytes initializer, uint256 saltNonce) returns(address proxy)
func (_Contracts *ContractsTransactorSession) NewMedia(mediaEOA common.Address, accountMetadata string, initializer []byte, saltNonce *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.NewMedia(&_Contracts.TransactOpts, mediaEOA, accountMetadata, initializer, saltNonce)
}

// ContractsCreateProxyIterator is returned from FilterCreateProxy and is used to iterate over the raw logs and unpacked data for CreateProxy events raised by the Contracts contract.
type ContractsCreateProxyIterator struct {
	Event *ContractsCreateProxy // Event containing the contract specifics and raw log

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
func (it *ContractsCreateProxyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsCreateProxy)
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
		it.Event = new(ContractsCreateProxy)
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
func (it *ContractsCreateProxyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsCreateProxyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsCreateProxy represents a CreateProxy event raised by the Contracts contract.
type ContractsCreateProxy struct {
	Proxy common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCreateProxy is a free log retrieval operation binding the contract event 0x32340bf052adf19b6294f5a724325e4cc444978d2cc5c7e2769a2df85db66098.
//
// Solidity: event CreateProxy(address proxy)
func (_Contracts *ContractsFilterer) FilterCreateProxy(opts *bind.FilterOpts) (*ContractsCreateProxyIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "CreateProxy")
	if err != nil {
		return nil, err
	}
	return &ContractsCreateProxyIterator{contract: _Contracts.contract, event: "CreateProxy", logs: logs, sub: sub}, nil
}

// WatchCreateProxy is a free log subscription operation binding the contract event 0x32340bf052adf19b6294f5a724325e4cc444978d2cc5c7e2769a2df85db66098.
//
// Solidity: event CreateProxy(address proxy)
func (_Contracts *ContractsFilterer) WatchCreateProxy(opts *bind.WatchOpts, sink chan<- *ContractsCreateProxy) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "CreateProxy")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsCreateProxy)
				if err := _Contracts.contract.UnpackLog(event, "CreateProxy", log); err != nil {
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

// ParseCreateProxy is a log parse operation binding the contract event 0x32340bf052adf19b6294f5a724325e4cc444978d2cc5c7e2769a2df85db66098.
//
// Solidity: event CreateProxy(address proxy)
func (_Contracts *ContractsFilterer) ParseCreateProxy(log types.Log) (*ContractsCreateProxy, error) {
	event := new(ContractsCreateProxy)
	if err := _Contracts.contract.UnpackLog(event, "CreateProxy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
