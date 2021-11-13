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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"accept\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tokenMetadata\",\"type\":\"string\"}],\"name\":\"acceptOffer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accepted\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"spaceMetadata\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"displayStartTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"displayEndTimestamp\",\"type\":\"uint256\"}],\"name\":\"adId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adPoolAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPeriods\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"mediaProxy\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"spaceMetadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenMetadata\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"saleStartTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"saleEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"displayStartTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"displayEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumAd.Pricing\",\"name\":\"pricing\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startPrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"sold\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"balance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"bidding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"bidder\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"buy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"buyBasedOnTime\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"currentPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"deletePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deniedReason\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"deny\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"spaceMetadata\",\"type\":\"string\"}],\"name\":\"display\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eventEmitterAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"nameRegistry\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mediaFactoryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mediaRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nameRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"spaceMetadata\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tokenMetadata\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"saleEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"displayStartTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"displayEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"enumAd.Pricing\",\"name\":\"pricing\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"minPrice\",\"type\":\"uint256\"}],\"name\":\"newPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"spaceMetadata\",\"type\":\"string\"}],\"name\":\"newSpace\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"spaceMetadata\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"displayStartTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"displayEndTimestamp\",\"type\":\"uint256\"}],\"name\":\"offerPeriod\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"offered\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"spaceMetadata\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"displayStartTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"displayEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadata\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"proposed\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"receiveToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"spaced\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"spaceMetadata\",\"type\":\"string\"}],\"name\":\"tokenIdsOf\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferToBundle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newMediaEOA\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"newMetadata\",\"type\":\"string\"}],\"name\":\"updateMedia\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// Accepted is a free data retrieval call binding the contract method 0xf19b8273.
//
// Solidity: function accepted(uint256 ) view returns(string)
func (_Contracts *ContractsCaller) Accepted(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "accepted", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Accepted is a free data retrieval call binding the contract method 0xf19b8273.
//
// Solidity: function accepted(uint256 ) view returns(string)
func (_Contracts *ContractsSession) Accepted(arg0 *big.Int) (string, error) {
	return _Contracts.Contract.Accepted(&_Contracts.CallOpts, arg0)
}

// Accepted is a free data retrieval call binding the contract method 0xf19b8273.
//
// Solidity: function accepted(uint256 ) view returns(string)
func (_Contracts *ContractsCallerSession) Accepted(arg0 *big.Int) (string, error) {
	return _Contracts.Contract.Accepted(&_Contracts.CallOpts, arg0)
}

// AdId is a free data retrieval call binding the contract method 0xc531c236.
//
// Solidity: function adId(string spaceMetadata, uint256 displayStartTimestamp, uint256 displayEndTimestamp) pure returns(uint256)
func (_Contracts *ContractsCaller) AdId(opts *bind.CallOpts, spaceMetadata string, displayStartTimestamp *big.Int, displayEndTimestamp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "adId", spaceMetadata, displayStartTimestamp, displayEndTimestamp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdId is a free data retrieval call binding the contract method 0xc531c236.
//
// Solidity: function adId(string spaceMetadata, uint256 displayStartTimestamp, uint256 displayEndTimestamp) pure returns(uint256)
func (_Contracts *ContractsSession) AdId(spaceMetadata string, displayStartTimestamp *big.Int, displayEndTimestamp *big.Int) (*big.Int, error) {
	return _Contracts.Contract.AdId(&_Contracts.CallOpts, spaceMetadata, displayStartTimestamp, displayEndTimestamp)
}

// AdId is a free data retrieval call binding the contract method 0xc531c236.
//
// Solidity: function adId(string spaceMetadata, uint256 displayStartTimestamp, uint256 displayEndTimestamp) pure returns(uint256)
func (_Contracts *ContractsCallerSession) AdId(spaceMetadata string, displayStartTimestamp *big.Int, displayEndTimestamp *big.Int) (*big.Int, error) {
	return _Contracts.Contract.AdId(&_Contracts.CallOpts, spaceMetadata, displayStartTimestamp, displayEndTimestamp)
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

// AllPeriods is a free data retrieval call binding the contract method 0xb3a0ebc9.
//
// Solidity: function allPeriods(uint256 ) view returns(address mediaProxy, string spaceMetadata, string tokenMetadata, uint256 saleStartTimestamp, uint256 saleEndTimestamp, uint256 displayStartTimestamp, uint256 displayEndTimestamp, uint8 pricing, uint256 minPrice, uint256 startPrice, bool sold)
func (_Contracts *ContractsCaller) AllPeriods(opts *bind.CallOpts, arg0 *big.Int) (struct {
	MediaProxy            common.Address
	SpaceMetadata         string
	TokenMetadata         string
	SaleStartTimestamp    *big.Int
	SaleEndTimestamp      *big.Int
	DisplayStartTimestamp *big.Int
	DisplayEndTimestamp   *big.Int
	Pricing               uint8
	MinPrice              *big.Int
	StartPrice            *big.Int
	Sold                  bool
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "allPeriods", arg0)

	outstruct := new(struct {
		MediaProxy            common.Address
		SpaceMetadata         string
		TokenMetadata         string
		SaleStartTimestamp    *big.Int
		SaleEndTimestamp      *big.Int
		DisplayStartTimestamp *big.Int
		DisplayEndTimestamp   *big.Int
		Pricing               uint8
		MinPrice              *big.Int
		StartPrice            *big.Int
		Sold                  bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MediaProxy = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.SpaceMetadata = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.TokenMetadata = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.SaleStartTimestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.SaleEndTimestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.DisplayStartTimestamp = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.DisplayEndTimestamp = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.Pricing = *abi.ConvertType(out[7], new(uint8)).(*uint8)
	outstruct.MinPrice = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.StartPrice = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.Sold = *abi.ConvertType(out[10], new(bool)).(*bool)

	return *outstruct, err

}

// AllPeriods is a free data retrieval call binding the contract method 0xb3a0ebc9.
//
// Solidity: function allPeriods(uint256 ) view returns(address mediaProxy, string spaceMetadata, string tokenMetadata, uint256 saleStartTimestamp, uint256 saleEndTimestamp, uint256 displayStartTimestamp, uint256 displayEndTimestamp, uint8 pricing, uint256 minPrice, uint256 startPrice, bool sold)
func (_Contracts *ContractsSession) AllPeriods(arg0 *big.Int) (struct {
	MediaProxy            common.Address
	SpaceMetadata         string
	TokenMetadata         string
	SaleStartTimestamp    *big.Int
	SaleEndTimestamp      *big.Int
	DisplayStartTimestamp *big.Int
	DisplayEndTimestamp   *big.Int
	Pricing               uint8
	MinPrice              *big.Int
	StartPrice            *big.Int
	Sold                  bool
}, error) {
	return _Contracts.Contract.AllPeriods(&_Contracts.CallOpts, arg0)
}

// AllPeriods is a free data retrieval call binding the contract method 0xb3a0ebc9.
//
// Solidity: function allPeriods(uint256 ) view returns(address mediaProxy, string spaceMetadata, string tokenMetadata, uint256 saleStartTimestamp, uint256 saleEndTimestamp, uint256 displayStartTimestamp, uint256 displayEndTimestamp, uint8 pricing, uint256 minPrice, uint256 startPrice, bool sold)
func (_Contracts *ContractsCallerSession) AllPeriods(arg0 *big.Int) (struct {
	MediaProxy            common.Address
	SpaceMetadata         string
	TokenMetadata         string
	SaleStartTimestamp    *big.Int
	SaleEndTimestamp      *big.Int
	DisplayStartTimestamp *big.Int
	DisplayEndTimestamp   *big.Int
	Pricing               uint8
	MinPrice              *big.Int
	StartPrice            *big.Int
	Sold                  bool
}, error) {
	return _Contracts.Contract.AllPeriods(&_Contracts.CallOpts, arg0)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Contracts *ContractsCaller) Balance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "balance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Contracts *ContractsSession) Balance() (*big.Int, error) {
	return _Contracts.Contract.Balance(&_Contracts.CallOpts)
}

// Balance is a free data retrieval call binding the contract method 0xb69ef8a8.
//
// Solidity: function balance() view returns(uint256)
func (_Contracts *ContractsCallerSession) Balance() (*big.Int, error) {
	return _Contracts.Contract.Balance(&_Contracts.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contracts *ContractsCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contracts *ContractsSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Contracts *ContractsCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Contracts.Contract.BalanceOf(&_Contracts.CallOpts, owner)
}

// Bidding is a free data retrieval call binding the contract method 0xcc889b0b.
//
// Solidity: function bidding(uint256 ) view returns(uint256 tokenId, address bidder, uint256 price)
func (_Contracts *ContractsCaller) Bidding(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenId *big.Int
	Bidder  common.Address
	Price   *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "bidding", arg0)

	outstruct := new(struct {
		TokenId *big.Int
		Bidder  common.Address
		Price   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Bidder = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Bidding is a free data retrieval call binding the contract method 0xcc889b0b.
//
// Solidity: function bidding(uint256 ) view returns(uint256 tokenId, address bidder, uint256 price)
func (_Contracts *ContractsSession) Bidding(arg0 *big.Int) (struct {
	TokenId *big.Int
	Bidder  common.Address
	Price   *big.Int
}, error) {
	return _Contracts.Contract.Bidding(&_Contracts.CallOpts, arg0)
}

// Bidding is a free data retrieval call binding the contract method 0xcc889b0b.
//
// Solidity: function bidding(uint256 ) view returns(uint256 tokenId, address bidder, uint256 price)
func (_Contracts *ContractsCallerSession) Bidding(arg0 *big.Int) (struct {
	TokenId *big.Int
	Bidder  common.Address
	Price   *big.Int
}, error) {
	return _Contracts.Contract.Bidding(&_Contracts.CallOpts, arg0)
}

// CurrentPrice is a free data retrieval call binding the contract method 0x7a3c4c17.
//
// Solidity: function currentPrice(uint256 tokenId) view returns(uint256)
func (_Contracts *ContractsCaller) CurrentPrice(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "currentPrice", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentPrice is a free data retrieval call binding the contract method 0x7a3c4c17.
//
// Solidity: function currentPrice(uint256 tokenId) view returns(uint256)
func (_Contracts *ContractsSession) CurrentPrice(tokenId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.CurrentPrice(&_Contracts.CallOpts, tokenId)
}

// CurrentPrice is a free data retrieval call binding the contract method 0x7a3c4c17.
//
// Solidity: function currentPrice(uint256 tokenId) view returns(uint256)
func (_Contracts *ContractsCallerSession) CurrentPrice(tokenId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.CurrentPrice(&_Contracts.CallOpts, tokenId)
}

// DeniedReason is a free data retrieval call binding the contract method 0xa9a86f0e.
//
// Solidity: function deniedReason(uint256 ) view returns(string)
func (_Contracts *ContractsCaller) DeniedReason(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "deniedReason", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DeniedReason is a free data retrieval call binding the contract method 0xa9a86f0e.
//
// Solidity: function deniedReason(uint256 ) view returns(string)
func (_Contracts *ContractsSession) DeniedReason(arg0 *big.Int) (string, error) {
	return _Contracts.Contract.DeniedReason(&_Contracts.CallOpts, arg0)
}

// DeniedReason is a free data retrieval call binding the contract method 0xa9a86f0e.
//
// Solidity: function deniedReason(uint256 ) view returns(string)
func (_Contracts *ContractsCallerSession) DeniedReason(arg0 *big.Int) (string, error) {
	return _Contracts.Contract.DeniedReason(&_Contracts.CallOpts, arg0)
}

// Display is a free data retrieval call binding the contract method 0x53c14776.
//
// Solidity: function display(string spaceMetadata) view returns(string)
func (_Contracts *ContractsCaller) Display(opts *bind.CallOpts, spaceMetadata string) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "display", spaceMetadata)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Display is a free data retrieval call binding the contract method 0x53c14776.
//
// Solidity: function display(string spaceMetadata) view returns(string)
func (_Contracts *ContractsSession) Display(spaceMetadata string) (string, error) {
	return _Contracts.Contract.Display(&_Contracts.CallOpts, spaceMetadata)
}

// Display is a free data retrieval call binding the contract method 0x53c14776.
//
// Solidity: function display(string spaceMetadata) view returns(string)
func (_Contracts *ContractsCallerSession) Display(spaceMetadata string) (string, error) {
	return _Contracts.Contract.Display(&_Contracts.CallOpts, spaceMetadata)
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

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contracts *ContractsCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contracts *ContractsSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetApproved(&_Contracts.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Contracts *ContractsCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Contracts.Contract.GetApproved(&_Contracts.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contracts *ContractsCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contracts *ContractsSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Contracts.Contract.IsApprovedForAll(&_Contracts.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Contracts *ContractsCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Contracts.Contract.IsApprovedForAll(&_Contracts.CallOpts, owner, operator)
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

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Contracts *ContractsCallerSession) Name() (string, error) {
	return _Contracts.Contract.Name(&_Contracts.CallOpts)
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

// Offered is a free data retrieval call binding the contract method 0x01b99e04.
//
// Solidity: function offered(uint256 ) view returns(string spaceMetadata, uint256 displayStartTimestamp, uint256 displayEndTimestamp, address sender, uint256 price)
func (_Contracts *ContractsCaller) Offered(opts *bind.CallOpts, arg0 *big.Int) (struct {
	SpaceMetadata         string
	DisplayStartTimestamp *big.Int
	DisplayEndTimestamp   *big.Int
	Sender                common.Address
	Price                 *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "offered", arg0)

	outstruct := new(struct {
		SpaceMetadata         string
		DisplayStartTimestamp *big.Int
		DisplayEndTimestamp   *big.Int
		Sender                common.Address
		Price                 *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SpaceMetadata = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.DisplayStartTimestamp = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DisplayEndTimestamp = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Sender = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Offered is a free data retrieval call binding the contract method 0x01b99e04.
//
// Solidity: function offered(uint256 ) view returns(string spaceMetadata, uint256 displayStartTimestamp, uint256 displayEndTimestamp, address sender, uint256 price)
func (_Contracts *ContractsSession) Offered(arg0 *big.Int) (struct {
	SpaceMetadata         string
	DisplayStartTimestamp *big.Int
	DisplayEndTimestamp   *big.Int
	Sender                common.Address
	Price                 *big.Int
}, error) {
	return _Contracts.Contract.Offered(&_Contracts.CallOpts, arg0)
}

// Offered is a free data retrieval call binding the contract method 0x01b99e04.
//
// Solidity: function offered(uint256 ) view returns(string spaceMetadata, uint256 displayStartTimestamp, uint256 displayEndTimestamp, address sender, uint256 price)
func (_Contracts *ContractsCallerSession) Offered(arg0 *big.Int) (struct {
	SpaceMetadata         string
	DisplayStartTimestamp *big.Int
	DisplayEndTimestamp   *big.Int
	Sender                common.Address
	Price                 *big.Int
}, error) {
	return _Contracts.Contract.Offered(&_Contracts.CallOpts, arg0)
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

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contracts *ContractsCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contracts *ContractsSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Contracts.Contract.OwnerOf(&_Contracts.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Contracts *ContractsCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Contracts.Contract.OwnerOf(&_Contracts.CallOpts, tokenId)
}

// Proposed is a free data retrieval call binding the contract method 0x58ba1bb8.
//
// Solidity: function proposed(uint256 ) view returns(string)
func (_Contracts *ContractsCaller) Proposed(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "proposed", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Proposed is a free data retrieval call binding the contract method 0x58ba1bb8.
//
// Solidity: function proposed(uint256 ) view returns(string)
func (_Contracts *ContractsSession) Proposed(arg0 *big.Int) (string, error) {
	return _Contracts.Contract.Proposed(&_Contracts.CallOpts, arg0)
}

// Proposed is a free data retrieval call binding the contract method 0x58ba1bb8.
//
// Solidity: function proposed(uint256 ) view returns(string)
func (_Contracts *ContractsCallerSession) Proposed(arg0 *big.Int) (string, error) {
	return _Contracts.Contract.Proposed(&_Contracts.CallOpts, arg0)
}

// Spaced is a free data retrieval call binding the contract method 0x65045a65.
//
// Solidity: function spaced(string ) view returns(bool)
func (_Contracts *ContractsCaller) Spaced(opts *bind.CallOpts, arg0 string) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "spaced", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Spaced is a free data retrieval call binding the contract method 0x65045a65.
//
// Solidity: function spaced(string ) view returns(bool)
func (_Contracts *ContractsSession) Spaced(arg0 string) (bool, error) {
	return _Contracts.Contract.Spaced(&_Contracts.CallOpts, arg0)
}

// Spaced is a free data retrieval call binding the contract method 0x65045a65.
//
// Solidity: function spaced(string ) view returns(bool)
func (_Contracts *ContractsCallerSession) Spaced(arg0 string) (bool, error) {
	return _Contracts.Contract.Spaced(&_Contracts.CallOpts, arg0)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Contracts *ContractsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Contracts *ContractsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contracts.Contract.SupportsInterface(&_Contracts.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Contracts *ContractsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Contracts.Contract.SupportsInterface(&_Contracts.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Contracts *ContractsCallerSession) Symbol() (string, error) {
	return _Contracts.Contract.Symbol(&_Contracts.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Contracts *ContractsCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Contracts *ContractsSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Contracts.Contract.TokenByIndex(&_Contracts.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Contracts *ContractsCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Contracts.Contract.TokenByIndex(&_Contracts.CallOpts, index)
}

// TokenIdsOf is a free data retrieval call binding the contract method 0xd028bd66.
//
// Solidity: function tokenIdsOf(string spaceMetadata) view returns(uint256[])
func (_Contracts *ContractsCaller) TokenIdsOf(opts *bind.CallOpts, spaceMetadata string) ([]*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "tokenIdsOf", spaceMetadata)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokenIdsOf is a free data retrieval call binding the contract method 0xd028bd66.
//
// Solidity: function tokenIdsOf(string spaceMetadata) view returns(uint256[])
func (_Contracts *ContractsSession) TokenIdsOf(spaceMetadata string) ([]*big.Int, error) {
	return _Contracts.Contract.TokenIdsOf(&_Contracts.CallOpts, spaceMetadata)
}

// TokenIdsOf is a free data retrieval call binding the contract method 0xd028bd66.
//
// Solidity: function tokenIdsOf(string spaceMetadata) view returns(uint256[])
func (_Contracts *ContractsCallerSession) TokenIdsOf(spaceMetadata string) ([]*big.Int, error) {
	return _Contracts.Contract.TokenIdsOf(&_Contracts.CallOpts, spaceMetadata)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Contracts *ContractsCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Contracts *ContractsSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Contracts.Contract.TokenOfOwnerByIndex(&_Contracts.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Contracts *ContractsCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Contracts.Contract.TokenOfOwnerByIndex(&_Contracts.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Contracts *ContractsCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Contracts *ContractsSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Contracts.Contract.TokenURI(&_Contracts.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Contracts *ContractsCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Contracts.Contract.TokenURI(&_Contracts.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsSession) TotalSupply() (*big.Int, error) {
	return _Contracts.Contract.TotalSupply(&_Contracts.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Contracts *ContractsCallerSession) TotalSupply() (*big.Int, error) {
	return _Contracts.Contract.TotalSupply(&_Contracts.CallOpts)
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
// Solidity: function accept(uint256 tokenId) returns()
func (_Contracts *ContractsTransactor) Accept(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "accept", tokenId)
}

// Accept is a paid mutator transaction binding the contract method 0x19b05f49.
//
// Solidity: function accept(uint256 tokenId) returns()
func (_Contracts *ContractsSession) Accept(tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Accept(&_Contracts.TransactOpts, tokenId)
}

// Accept is a paid mutator transaction binding the contract method 0x19b05f49.
//
// Solidity: function accept(uint256 tokenId) returns()
func (_Contracts *ContractsTransactorSession) Accept(tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Accept(&_Contracts.TransactOpts, tokenId)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x444115f6.
//
// Solidity: function acceptOffer(uint256 tokenId, string tokenMetadata) returns()
func (_Contracts *ContractsTransactor) AcceptOffer(opts *bind.TransactOpts, tokenId *big.Int, tokenMetadata string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "acceptOffer", tokenId, tokenMetadata)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x444115f6.
//
// Solidity: function acceptOffer(uint256 tokenId, string tokenMetadata) returns()
func (_Contracts *ContractsSession) AcceptOffer(tokenId *big.Int, tokenMetadata string) (*types.Transaction, error) {
	return _Contracts.Contract.AcceptOffer(&_Contracts.TransactOpts, tokenId, tokenMetadata)
}

// AcceptOffer is a paid mutator transaction binding the contract method 0x444115f6.
//
// Solidity: function acceptOffer(uint256 tokenId, string tokenMetadata) returns()
func (_Contracts *ContractsTransactorSession) AcceptOffer(tokenId *big.Int, tokenMetadata string) (*types.Transaction, error) {
	return _Contracts.Contract.AcceptOffer(&_Contracts.TransactOpts, tokenId, tokenMetadata)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contracts *ContractsTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contracts *ContractsSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Approve(&_Contracts.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Contracts *ContractsTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Approve(&_Contracts.TransactOpts, to, tokenId)
}

// Bid1 is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 tokenId) payable returns()
func (_Contracts *ContractsTransactor) Bid1(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "bid", tokenId)
}

// Bid1 is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 tokenId) payable returns()
func (_Contracts *ContractsSession) Bid1(tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Bid1(&_Contracts.TransactOpts, tokenId)
}

// Bid1 is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 tokenId) payable returns()
func (_Contracts *ContractsTransactorSession) Bid1(tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Bid1(&_Contracts.TransactOpts, tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 tokenId) payable returns()
func (_Contracts *ContractsTransactor) Buy(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "buy", tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 tokenId) payable returns()
func (_Contracts *ContractsSession) Buy(tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Buy(&_Contracts.TransactOpts, tokenId)
}

// Buy is a paid mutator transaction binding the contract method 0xd96a094a.
//
// Solidity: function buy(uint256 tokenId) payable returns()
func (_Contracts *ContractsTransactorSession) Buy(tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Buy(&_Contracts.TransactOpts, tokenId)
}

// BuyBasedOnTime is a paid mutator transaction binding the contract method 0xd8680069.
//
// Solidity: function buyBasedOnTime(uint256 tokenId) payable returns()
func (_Contracts *ContractsTransactor) BuyBasedOnTime(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "buyBasedOnTime", tokenId)
}

// BuyBasedOnTime is a paid mutator transaction binding the contract method 0xd8680069.
//
// Solidity: function buyBasedOnTime(uint256 tokenId) payable returns()
func (_Contracts *ContractsSession) BuyBasedOnTime(tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.BuyBasedOnTime(&_Contracts.TransactOpts, tokenId)
}

// BuyBasedOnTime is a paid mutator transaction binding the contract method 0xd8680069.
//
// Solidity: function buyBasedOnTime(uint256 tokenId) payable returns()
func (_Contracts *ContractsTransactorSession) BuyBasedOnTime(tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.BuyBasedOnTime(&_Contracts.TransactOpts, tokenId)
}

// DeletePeriod is a paid mutator transaction binding the contract method 0x4b455a51.
//
// Solidity: function deletePeriod(uint256 tokenId) returns()
func (_Contracts *ContractsTransactor) DeletePeriod(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "deletePeriod", tokenId)
}

// DeletePeriod is a paid mutator transaction binding the contract method 0x4b455a51.
//
// Solidity: function deletePeriod(uint256 tokenId) returns()
func (_Contracts *ContractsSession) DeletePeriod(tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DeletePeriod(&_Contracts.TransactOpts, tokenId)
}

// DeletePeriod is a paid mutator transaction binding the contract method 0x4b455a51.
//
// Solidity: function deletePeriod(uint256 tokenId) returns()
func (_Contracts *ContractsTransactorSession) DeletePeriod(tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.DeletePeriod(&_Contracts.TransactOpts, tokenId)
}

// Deny is a paid mutator transaction binding the contract method 0x7393f289.
//
// Solidity: function deny(uint256 tokenId, string reason) returns()
func (_Contracts *ContractsTransactor) Deny(opts *bind.TransactOpts, tokenId *big.Int, reason string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "deny", tokenId, reason)
}

// Deny is a paid mutator transaction binding the contract method 0x7393f289.
//
// Solidity: function deny(uint256 tokenId, string reason) returns()
func (_Contracts *ContractsSession) Deny(tokenId *big.Int, reason string) (*types.Transaction, error) {
	return _Contracts.Contract.Deny(&_Contracts.TransactOpts, tokenId, reason)
}

// Deny is a paid mutator transaction binding the contract method 0x7393f289.
//
// Solidity: function deny(uint256 tokenId, string reason) returns()
func (_Contracts *ContractsTransactorSession) Deny(tokenId *big.Int, reason string) (*types.Transaction, error) {
	return _Contracts.Contract.Deny(&_Contracts.TransactOpts, tokenId, reason)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string title, string baseURI, address nameRegistry) returns()
func (_Contracts *ContractsTransactor) Initialize(opts *bind.TransactOpts, title string, baseURI string, nameRegistry common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "initialize", title, baseURI, nameRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string title, string baseURI, address nameRegistry) returns()
func (_Contracts *ContractsSession) Initialize(title string, baseURI string, nameRegistry common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, title, baseURI, nameRegistry)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string title, string baseURI, address nameRegistry) returns()
func (_Contracts *ContractsTransactorSession) Initialize(title string, baseURI string, nameRegistry common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.Initialize(&_Contracts.TransactOpts, title, baseURI, nameRegistry)
}

// NewPeriod is a paid mutator transaction binding the contract method 0x899c9989.
//
// Solidity: function newPeriod(string spaceMetadata, string tokenMetadata, uint256 saleEndTimestamp, uint256 displayStartTimestamp, uint256 displayEndTimestamp, uint8 pricing, uint256 minPrice) returns()
func (_Contracts *ContractsTransactor) NewPeriod(opts *bind.TransactOpts, spaceMetadata string, tokenMetadata string, saleEndTimestamp *big.Int, displayStartTimestamp *big.Int, displayEndTimestamp *big.Int, pricing uint8, minPrice *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "newPeriod", spaceMetadata, tokenMetadata, saleEndTimestamp, displayStartTimestamp, displayEndTimestamp, pricing, minPrice)
}

// NewPeriod is a paid mutator transaction binding the contract method 0x899c9989.
//
// Solidity: function newPeriod(string spaceMetadata, string tokenMetadata, uint256 saleEndTimestamp, uint256 displayStartTimestamp, uint256 displayEndTimestamp, uint8 pricing, uint256 minPrice) returns()
func (_Contracts *ContractsSession) NewPeriod(spaceMetadata string, tokenMetadata string, saleEndTimestamp *big.Int, displayStartTimestamp *big.Int, displayEndTimestamp *big.Int, pricing uint8, minPrice *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.NewPeriod(&_Contracts.TransactOpts, spaceMetadata, tokenMetadata, saleEndTimestamp, displayStartTimestamp, displayEndTimestamp, pricing, minPrice)
}

// NewPeriod is a paid mutator transaction binding the contract method 0x899c9989.
//
// Solidity: function newPeriod(string spaceMetadata, string tokenMetadata, uint256 saleEndTimestamp, uint256 displayStartTimestamp, uint256 displayEndTimestamp, uint8 pricing, uint256 minPrice) returns()
func (_Contracts *ContractsTransactorSession) NewPeriod(spaceMetadata string, tokenMetadata string, saleEndTimestamp *big.Int, displayStartTimestamp *big.Int, displayEndTimestamp *big.Int, pricing uint8, minPrice *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.NewPeriod(&_Contracts.TransactOpts, spaceMetadata, tokenMetadata, saleEndTimestamp, displayStartTimestamp, displayEndTimestamp, pricing, minPrice)
}

// NewSpace is a paid mutator transaction binding the contract method 0x2bc4bc68.
//
// Solidity: function newSpace(string spaceMetadata) returns()
func (_Contracts *ContractsTransactor) NewSpace(opts *bind.TransactOpts, spaceMetadata string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "newSpace", spaceMetadata)
}

// NewSpace is a paid mutator transaction binding the contract method 0x2bc4bc68.
//
// Solidity: function newSpace(string spaceMetadata) returns()
func (_Contracts *ContractsSession) NewSpace(spaceMetadata string) (*types.Transaction, error) {
	return _Contracts.Contract.NewSpace(&_Contracts.TransactOpts, spaceMetadata)
}

// NewSpace is a paid mutator transaction binding the contract method 0x2bc4bc68.
//
// Solidity: function newSpace(string spaceMetadata) returns()
func (_Contracts *ContractsTransactorSession) NewSpace(spaceMetadata string) (*types.Transaction, error) {
	return _Contracts.Contract.NewSpace(&_Contracts.TransactOpts, spaceMetadata)
}

// OfferPeriod is a paid mutator transaction binding the contract method 0x1daf7ff9.
//
// Solidity: function offerPeriod(string spaceMetadata, uint256 displayStartTimestamp, uint256 displayEndTimestamp) payable returns()
func (_Contracts *ContractsTransactor) OfferPeriod(opts *bind.TransactOpts, spaceMetadata string, displayStartTimestamp *big.Int, displayEndTimestamp *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "offerPeriod", spaceMetadata, displayStartTimestamp, displayEndTimestamp)
}

// OfferPeriod is a paid mutator transaction binding the contract method 0x1daf7ff9.
//
// Solidity: function offerPeriod(string spaceMetadata, uint256 displayStartTimestamp, uint256 displayEndTimestamp) payable returns()
func (_Contracts *ContractsSession) OfferPeriod(spaceMetadata string, displayStartTimestamp *big.Int, displayEndTimestamp *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.OfferPeriod(&_Contracts.TransactOpts, spaceMetadata, displayStartTimestamp, displayEndTimestamp)
}

// OfferPeriod is a paid mutator transaction binding the contract method 0x1daf7ff9.
//
// Solidity: function offerPeriod(string spaceMetadata, uint256 displayStartTimestamp, uint256 displayEndTimestamp) payable returns()
func (_Contracts *ContractsTransactorSession) OfferPeriod(spaceMetadata string, displayStartTimestamp *big.Int, displayEndTimestamp *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.OfferPeriod(&_Contracts.TransactOpts, spaceMetadata, displayStartTimestamp, displayEndTimestamp)
}

// Propose is a paid mutator transaction binding the contract method 0xd4f6b5ec.
//
// Solidity: function propose(uint256 tokenId, string metadata) returns()
func (_Contracts *ContractsTransactor) Propose(opts *bind.TransactOpts, tokenId *big.Int, metadata string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "propose", tokenId, metadata)
}

// Propose is a paid mutator transaction binding the contract method 0xd4f6b5ec.
//
// Solidity: function propose(uint256 tokenId, string metadata) returns()
func (_Contracts *ContractsSession) Propose(tokenId *big.Int, metadata string) (*types.Transaction, error) {
	return _Contracts.Contract.Propose(&_Contracts.TransactOpts, tokenId, metadata)
}

// Propose is a paid mutator transaction binding the contract method 0xd4f6b5ec.
//
// Solidity: function propose(uint256 tokenId, string metadata) returns()
func (_Contracts *ContractsTransactorSession) Propose(tokenId *big.Int, metadata string) (*types.Transaction, error) {
	return _Contracts.Contract.Propose(&_Contracts.TransactOpts, tokenId, metadata)
}

// ReceiveToken is a paid mutator transaction binding the contract method 0x37df00c9.
//
// Solidity: function receiveToken(uint256 tokenId) payable returns()
func (_Contracts *ContractsTransactor) ReceiveToken(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "receiveToken", tokenId)
}

// ReceiveToken is a paid mutator transaction binding the contract method 0x37df00c9.
//
// Solidity: function receiveToken(uint256 tokenId) payable returns()
func (_Contracts *ContractsSession) ReceiveToken(tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ReceiveToken(&_Contracts.TransactOpts, tokenId)
}

// ReceiveToken is a paid mutator transaction binding the contract method 0x37df00c9.
//
// Solidity: function receiveToken(uint256 tokenId) payable returns()
func (_Contracts *ContractsTransactorSession) ReceiveToken(tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ReceiveToken(&_Contracts.TransactOpts, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contracts *ContractsTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contracts *ContractsSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SafeTransferFrom(&_Contracts.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Contracts *ContractsTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SafeTransferFrom(&_Contracts.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Contracts *ContractsTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Contracts *ContractsSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SafeTransferFrom0(&_Contracts.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Contracts *ContractsTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Contracts.Contract.SafeTransferFrom0(&_Contracts.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contracts *ContractsTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contracts *ContractsSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contracts.Contract.SetApprovalForAll(&_Contracts.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Contracts *ContractsTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Contracts.Contract.SetApprovalForAll(&_Contracts.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contracts *ContractsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contracts *ContractsSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferFrom(&_Contracts.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Contracts *ContractsTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferFrom(&_Contracts.TransactOpts, from, to, tokenId)
}

// TransferToBundle is a paid mutator transaction binding the contract method 0xd1b900cd.
//
// Solidity: function transferToBundle(address from, address to, uint256 tokenId) returns()
func (_Contracts *ContractsTransactor) TransferToBundle(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "transferToBundle", from, to, tokenId)
}

// TransferToBundle is a paid mutator transaction binding the contract method 0xd1b900cd.
//
// Solidity: function transferToBundle(address from, address to, uint256 tokenId) returns()
func (_Contracts *ContractsSession) TransferToBundle(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferToBundle(&_Contracts.TransactOpts, from, to, tokenId)
}

// TransferToBundle is a paid mutator transaction binding the contract method 0xd1b900cd.
//
// Solidity: function transferToBundle(address from, address to, uint256 tokenId) returns()
func (_Contracts *ContractsTransactorSession) TransferToBundle(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.TransferToBundle(&_Contracts.TransactOpts, from, to, tokenId)
}

// UpdateMedia is a paid mutator transaction binding the contract method 0x179bcff7.
//
// Solidity: function updateMedia(address newMediaEOA, string newMetadata) returns()
func (_Contracts *ContractsTransactor) UpdateMedia(opts *bind.TransactOpts, newMediaEOA common.Address, newMetadata string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "updateMedia", newMediaEOA, newMetadata)
}

// UpdateMedia is a paid mutator transaction binding the contract method 0x179bcff7.
//
// Solidity: function updateMedia(address newMediaEOA, string newMetadata) returns()
func (_Contracts *ContractsSession) UpdateMedia(newMediaEOA common.Address, newMetadata string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateMedia(&_Contracts.TransactOpts, newMediaEOA, newMetadata)
}

// UpdateMedia is a paid mutator transaction binding the contract method 0x179bcff7.
//
// Solidity: function updateMedia(address newMediaEOA, string newMetadata) returns()
func (_Contracts *ContractsTransactorSession) UpdateMedia(newMediaEOA common.Address, newMetadata string) (*types.Transaction, error) {
	return _Contracts.Contract.UpdateMedia(&_Contracts.TransactOpts, newMediaEOA, newMetadata)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contracts *ContractsTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contracts *ContractsSession) Withdraw() (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_Contracts *ContractsTransactorSession) Withdraw() (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts)
}

// ContractsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Contracts contract.
type ContractsApprovalIterator struct {
	Event *ContractsApproval // Event containing the contract specifics and raw log

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
func (it *ContractsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsApproval)
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
		it.Event = new(ContractsApproval)
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
func (it *ContractsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsApproval represents a Approval event raised by the Contracts contract.
type ContractsApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contracts *ContractsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ContractsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsApprovalIterator{contract: _Contracts.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contracts *ContractsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractsApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsApproval)
				if err := _Contracts.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Contracts *ContractsFilterer) ParseApproval(log types.Log) (*ContractsApproval, error) {
	event := new(ContractsApproval)
	if err := _Contracts.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Contracts contract.
type ContractsApprovalForAllIterator struct {
	Event *ContractsApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ContractsApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsApprovalForAll)
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
		it.Event = new(ContractsApprovalForAll)
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
func (it *ContractsApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsApprovalForAll represents a ApprovalForAll event raised by the Contracts contract.
type ContractsApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contracts *ContractsFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ContractsApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ContractsApprovalForAllIterator{contract: _Contracts.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contracts *ContractsFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ContractsApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsApprovalForAll)
				if err := _Contracts.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Contracts *ContractsFilterer) ParseApprovalForAll(log types.Log) (*ContractsApprovalForAll, error) {
	event := new(ContractsApprovalForAll)
	if err := _Contracts.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Contracts contract.
type ContractsTransferIterator struct {
	Event *ContractsTransfer // Event containing the contract specifics and raw log

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
func (it *ContractsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsTransfer)
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
		it.Event = new(ContractsTransfer)
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
func (it *ContractsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsTransfer represents a Transfer event raised by the Contracts contract.
type ContractsTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contracts *ContractsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ContractsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ContractsTransferIterator{contract: _Contracts.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contracts *ContractsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractsTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsTransfer)
				if err := _Contracts.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Contracts *ContractsFilterer) ParseTransfer(log types.Log) (*ContractsTransfer, error) {
	event := new(ContractsTransfer)
	if err := _Contracts.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
