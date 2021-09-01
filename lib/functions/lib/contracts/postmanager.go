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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nameRegistry\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"accumulatorAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allDonations\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serialNum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"donateTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allPosts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"periodHours\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"donatedCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"donatedSum\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"auroraAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"burnUltraRare\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"}],\"name\":\"cancel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"}],\"name\":\"cancelableAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"serialNum\",\"type\":\"uint256\"}],\"name\":\"computeDonationReceiptId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"}],\"name\":\"computeRefundRequestId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"donate\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eventAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nameRegistryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"capacity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"periodHours\",\"type\":\"uint256\"}],\"name\":\"newPost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextReceiptSeed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"postAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"receiptAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"}],\"name\":\"refund\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refundAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"refundReceipts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"refundRequests\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"receiptId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"metadataURI\",\"type\":\"string\"}],\"name\":\"requestRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requireVoucher\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setTokenRewardAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"switchVoucherRequirement\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalValueLocked\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ultraRareAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vaultAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"voucherAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"postId\",\"type\":\"uint256\"}],\"name\":\"withdrawalAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// AccumulatorAddress is a free data retrieval call binding the contract method 0x049199c0.
//
// Solidity: function accumulatorAddress() view returns(address)
func (_Contracts *ContractsCaller) AccumulatorAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "accumulatorAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccumulatorAddress is a free data retrieval call binding the contract method 0x049199c0.
//
// Solidity: function accumulatorAddress() view returns(address)
func (_Contracts *ContractsSession) AccumulatorAddress() (common.Address, error) {
	return _Contracts.Contract.AccumulatorAddress(&_Contracts.CallOpts)
}

// AccumulatorAddress is a free data retrieval call binding the contract method 0x049199c0.
//
// Solidity: function accumulatorAddress() view returns(address)
func (_Contracts *ContractsCallerSession) AccumulatorAddress() (common.Address, error) {
	return _Contracts.Contract.AccumulatorAddress(&_Contracts.CallOpts)
}

// AllDonations is a free data retrieval call binding the contract method 0x06c49713.
//
// Solidity: function allDonations(uint256 ) view returns(uint256 receiptId, uint256 postId, uint256 serialNum, address sender, uint256 amount, uint256 donateTime)
func (_Contracts *ContractsCaller) AllDonations(opts *bind.CallOpts, arg0 *big.Int) (struct {
	ReceiptId  *big.Int
	PostId     *big.Int
	SerialNum  *big.Int
	Sender     common.Address
	Amount     *big.Int
	DonateTime *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "allDonations", arg0)

	outstruct := new(struct {
		ReceiptId  *big.Int
		PostId     *big.Int
		SerialNum  *big.Int
		Sender     common.Address
		Amount     *big.Int
		DonateTime *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReceiptId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.PostId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.SerialNum = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Sender = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.DonateTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AllDonations is a free data retrieval call binding the contract method 0x06c49713.
//
// Solidity: function allDonations(uint256 ) view returns(uint256 receiptId, uint256 postId, uint256 serialNum, address sender, uint256 amount, uint256 donateTime)
func (_Contracts *ContractsSession) AllDonations(arg0 *big.Int) (struct {
	ReceiptId  *big.Int
	PostId     *big.Int
	SerialNum  *big.Int
	Sender     common.Address
	Amount     *big.Int
	DonateTime *big.Int
}, error) {
	return _Contracts.Contract.AllDonations(&_Contracts.CallOpts, arg0)
}

// AllDonations is a free data retrieval call binding the contract method 0x06c49713.
//
// Solidity: function allDonations(uint256 ) view returns(uint256 receiptId, uint256 postId, uint256 serialNum, address sender, uint256 amount, uint256 donateTime)
func (_Contracts *ContractsCallerSession) AllDonations(arg0 *big.Int) (struct {
	ReceiptId  *big.Int
	PostId     *big.Int
	SerialNum  *big.Int
	Sender     common.Address
	Amount     *big.Int
	DonateTime *big.Int
}, error) {
	return _Contracts.Contract.AllDonations(&_Contracts.CallOpts, arg0)
}

// AllPosts is a free data retrieval call binding the contract method 0x718ce2bc.
//
// Solidity: function allPosts(uint256 ) view returns(uint256 postId, string metadataURI, uint256 capacity, uint256 periodHours, uint256 startTime, uint256 endTime, uint256 donatedCount, uint256 donatedSum)
func (_Contracts *ContractsCaller) AllPosts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	PostId       *big.Int
	MetadataURI  string
	Capacity     *big.Int
	PeriodHours  *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	DonatedCount *big.Int
	DonatedSum   *big.Int
}, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "allPosts", arg0)

	outstruct := new(struct {
		PostId       *big.Int
		MetadataURI  string
		Capacity     *big.Int
		PeriodHours  *big.Int
		StartTime    *big.Int
		EndTime      *big.Int
		DonatedCount *big.Int
		DonatedSum   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.PostId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MetadataURI = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Capacity = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PeriodHours = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StartTime = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.EndTime = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.DonatedCount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.DonatedSum = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// AllPosts is a free data retrieval call binding the contract method 0x718ce2bc.
//
// Solidity: function allPosts(uint256 ) view returns(uint256 postId, string metadataURI, uint256 capacity, uint256 periodHours, uint256 startTime, uint256 endTime, uint256 donatedCount, uint256 donatedSum)
func (_Contracts *ContractsSession) AllPosts(arg0 *big.Int) (struct {
	PostId       *big.Int
	MetadataURI  string
	Capacity     *big.Int
	PeriodHours  *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	DonatedCount *big.Int
	DonatedSum   *big.Int
}, error) {
	return _Contracts.Contract.AllPosts(&_Contracts.CallOpts, arg0)
}

// AllPosts is a free data retrieval call binding the contract method 0x718ce2bc.
//
// Solidity: function allPosts(uint256 ) view returns(uint256 postId, string metadataURI, uint256 capacity, uint256 periodHours, uint256 startTime, uint256 endTime, uint256 donatedCount, uint256 donatedSum)
func (_Contracts *ContractsCallerSession) AllPosts(arg0 *big.Int) (struct {
	PostId       *big.Int
	MetadataURI  string
	Capacity     *big.Int
	PeriodHours  *big.Int
	StartTime    *big.Int
	EndTime      *big.Int
	DonatedCount *big.Int
	DonatedSum   *big.Int
}, error) {
	return _Contracts.Contract.AllPosts(&_Contracts.CallOpts, arg0)
}

// AuroraAddress is a free data retrieval call binding the contract method 0x9783c809.
//
// Solidity: function auroraAddress() view returns(address)
func (_Contracts *ContractsCaller) AuroraAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "auroraAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AuroraAddress is a free data retrieval call binding the contract method 0x9783c809.
//
// Solidity: function auroraAddress() view returns(address)
func (_Contracts *ContractsSession) AuroraAddress() (common.Address, error) {
	return _Contracts.Contract.AuroraAddress(&_Contracts.CallOpts)
}

// AuroraAddress is a free data retrieval call binding the contract method 0x9783c809.
//
// Solidity: function auroraAddress() view returns(address)
func (_Contracts *ContractsCallerSession) AuroraAddress() (common.Address, error) {
	return _Contracts.Contract.AuroraAddress(&_Contracts.CallOpts)
}

// CancelableAmount is a free data retrieval call binding the contract method 0x20c5d0fa.
//
// Solidity: function cancelableAmount(uint256 receiptId) view returns(uint256)
func (_Contracts *ContractsCaller) CancelableAmount(opts *bind.CallOpts, receiptId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "cancelableAmount", receiptId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CancelableAmount is a free data retrieval call binding the contract method 0x20c5d0fa.
//
// Solidity: function cancelableAmount(uint256 receiptId) view returns(uint256)
func (_Contracts *ContractsSession) CancelableAmount(receiptId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.CancelableAmount(&_Contracts.CallOpts, receiptId)
}

// CancelableAmount is a free data retrieval call binding the contract method 0x20c5d0fa.
//
// Solidity: function cancelableAmount(uint256 receiptId) view returns(uint256)
func (_Contracts *ContractsCallerSession) CancelableAmount(receiptId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.CancelableAmount(&_Contracts.CallOpts, receiptId)
}

// ComputeDonationReceiptId is a free data retrieval call binding the contract method 0xf40e16a3.
//
// Solidity: function computeDonationReceiptId(address account, uint256 seed, uint256 serialNum) pure returns(uint256)
func (_Contracts *ContractsCaller) ComputeDonationReceiptId(opts *bind.CallOpts, account common.Address, seed *big.Int, serialNum *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "computeDonationReceiptId", account, seed, serialNum)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeDonationReceiptId is a free data retrieval call binding the contract method 0xf40e16a3.
//
// Solidity: function computeDonationReceiptId(address account, uint256 seed, uint256 serialNum) pure returns(uint256)
func (_Contracts *ContractsSession) ComputeDonationReceiptId(account common.Address, seed *big.Int, serialNum *big.Int) (*big.Int, error) {
	return _Contracts.Contract.ComputeDonationReceiptId(&_Contracts.CallOpts, account, seed, serialNum)
}

// ComputeDonationReceiptId is a free data retrieval call binding the contract method 0xf40e16a3.
//
// Solidity: function computeDonationReceiptId(address account, uint256 seed, uint256 serialNum) pure returns(uint256)
func (_Contracts *ContractsCallerSession) ComputeDonationReceiptId(account common.Address, seed *big.Int, serialNum *big.Int) (*big.Int, error) {
	return _Contracts.Contract.ComputeDonationReceiptId(&_Contracts.CallOpts, account, seed, serialNum)
}

// ComputeRefundRequestId is a free data retrieval call binding the contract method 0xbd49e9c5.
//
// Solidity: function computeRefundRequestId(address account, uint256 receiptId) pure returns(uint256)
func (_Contracts *ContractsCaller) ComputeRefundRequestId(opts *bind.CallOpts, account common.Address, receiptId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "computeRefundRequestId", account, receiptId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeRefundRequestId is a free data retrieval call binding the contract method 0xbd49e9c5.
//
// Solidity: function computeRefundRequestId(address account, uint256 receiptId) pure returns(uint256)
func (_Contracts *ContractsSession) ComputeRefundRequestId(account common.Address, receiptId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.ComputeRefundRequestId(&_Contracts.CallOpts, account, receiptId)
}

// ComputeRefundRequestId is a free data retrieval call binding the contract method 0xbd49e9c5.
//
// Solidity: function computeRefundRequestId(address account, uint256 receiptId) pure returns(uint256)
func (_Contracts *ContractsCallerSession) ComputeRefundRequestId(account common.Address, receiptId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.ComputeRefundRequestId(&_Contracts.CallOpts, account, receiptId)
}

// EventAddress is a free data retrieval call binding the contract method 0xfad56779.
//
// Solidity: function eventAddress() view returns(address)
func (_Contracts *ContractsCaller) EventAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "eventAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EventAddress is a free data retrieval call binding the contract method 0xfad56779.
//
// Solidity: function eventAddress() view returns(address)
func (_Contracts *ContractsSession) EventAddress() (common.Address, error) {
	return _Contracts.Contract.EventAddress(&_Contracts.CallOpts)
}

// EventAddress is a free data retrieval call binding the contract method 0xfad56779.
//
// Solidity: function eventAddress() view returns(address)
func (_Contracts *ContractsCallerSession) EventAddress() (common.Address, error) {
	return _Contracts.Contract.EventAddress(&_Contracts.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x4d6861a6.
//
// Solidity: function isOpen(uint256 postId) view returns(bool)
func (_Contracts *ContractsCaller) IsOpen(opts *bind.CallOpts, postId *big.Int) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "isOpen", postId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x4d6861a6.
//
// Solidity: function isOpen(uint256 postId) view returns(bool)
func (_Contracts *ContractsSession) IsOpen(postId *big.Int) (bool, error) {
	return _Contracts.Contract.IsOpen(&_Contracts.CallOpts, postId)
}

// IsOpen is a free data retrieval call binding the contract method 0x4d6861a6.
//
// Solidity: function isOpen(uint256 postId) view returns(bool)
func (_Contracts *ContractsCallerSession) IsOpen(postId *big.Int) (bool, error) {
	return _Contracts.Contract.IsOpen(&_Contracts.CallOpts, postId)
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

// NextReceiptSeed is a free data retrieval call binding the contract method 0xb756ed38.
//
// Solidity: function nextReceiptSeed() view returns(uint256)
func (_Contracts *ContractsCaller) NextReceiptSeed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "nextReceiptSeed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextReceiptSeed is a free data retrieval call binding the contract method 0xb756ed38.
//
// Solidity: function nextReceiptSeed() view returns(uint256)
func (_Contracts *ContractsSession) NextReceiptSeed() (*big.Int, error) {
	return _Contracts.Contract.NextReceiptSeed(&_Contracts.CallOpts)
}

// NextReceiptSeed is a free data retrieval call binding the contract method 0xb756ed38.
//
// Solidity: function nextReceiptSeed() view returns(uint256)
func (_Contracts *ContractsCallerSession) NextReceiptSeed() (*big.Int, error) {
	return _Contracts.Contract.NextReceiptSeed(&_Contracts.CallOpts)
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

// PostAddress is a free data retrieval call binding the contract method 0x99de7e75.
//
// Solidity: function postAddress() view returns(address)
func (_Contracts *ContractsCaller) PostAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "postAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PostAddress is a free data retrieval call binding the contract method 0x99de7e75.
//
// Solidity: function postAddress() view returns(address)
func (_Contracts *ContractsSession) PostAddress() (common.Address, error) {
	return _Contracts.Contract.PostAddress(&_Contracts.CallOpts)
}

// PostAddress is a free data retrieval call binding the contract method 0x99de7e75.
//
// Solidity: function postAddress() view returns(address)
func (_Contracts *ContractsCallerSession) PostAddress() (common.Address, error) {
	return _Contracts.Contract.PostAddress(&_Contracts.CallOpts)
}

// ReceiptAddress is a free data retrieval call binding the contract method 0x7cf707fc.
//
// Solidity: function receiptAddress() view returns(address)
func (_Contracts *ContractsCaller) ReceiptAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "receiptAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReceiptAddress is a free data retrieval call binding the contract method 0x7cf707fc.
//
// Solidity: function receiptAddress() view returns(address)
func (_Contracts *ContractsSession) ReceiptAddress() (common.Address, error) {
	return _Contracts.Contract.ReceiptAddress(&_Contracts.CallOpts)
}

// ReceiptAddress is a free data retrieval call binding the contract method 0x7cf707fc.
//
// Solidity: function receiptAddress() view returns(address)
func (_Contracts *ContractsCallerSession) ReceiptAddress() (common.Address, error) {
	return _Contracts.Contract.ReceiptAddress(&_Contracts.CallOpts)
}

// RefundAddress is a free data retrieval call binding the contract method 0x0cb61f6c.
//
// Solidity: function refundAddress() view returns(address)
func (_Contracts *ContractsCaller) RefundAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "refundAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RefundAddress is a free data retrieval call binding the contract method 0x0cb61f6c.
//
// Solidity: function refundAddress() view returns(address)
func (_Contracts *ContractsSession) RefundAddress() (common.Address, error) {
	return _Contracts.Contract.RefundAddress(&_Contracts.CallOpts)
}

// RefundAddress is a free data retrieval call binding the contract method 0x0cb61f6c.
//
// Solidity: function refundAddress() view returns(address)
func (_Contracts *ContractsCallerSession) RefundAddress() (common.Address, error) {
	return _Contracts.Contract.RefundAddress(&_Contracts.CallOpts)
}

// RefundReceipts is a free data retrieval call binding the contract method 0x5b7f97a2.
//
// Solidity: function refundReceipts(uint256 , uint256 ) view returns(uint256)
func (_Contracts *ContractsCaller) RefundReceipts(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "refundReceipts", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RefundReceipts is a free data retrieval call binding the contract method 0x5b7f97a2.
//
// Solidity: function refundReceipts(uint256 , uint256 ) view returns(uint256)
func (_Contracts *ContractsSession) RefundReceipts(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Contracts.Contract.RefundReceipts(&_Contracts.CallOpts, arg0, arg1)
}

// RefundReceipts is a free data retrieval call binding the contract method 0x5b7f97a2.
//
// Solidity: function refundReceipts(uint256 , uint256 ) view returns(uint256)
func (_Contracts *ContractsCallerSession) RefundReceipts(arg0 *big.Int, arg1 *big.Int) (*big.Int, error) {
	return _Contracts.Contract.RefundReceipts(&_Contracts.CallOpts, arg0, arg1)
}

// RefundRequests is a free data retrieval call binding the contract method 0x7d974ade.
//
// Solidity: function refundRequests(uint256 ) view returns(uint256)
func (_Contracts *ContractsCaller) RefundRequests(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "refundRequests", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RefundRequests is a free data retrieval call binding the contract method 0x7d974ade.
//
// Solidity: function refundRequests(uint256 ) view returns(uint256)
func (_Contracts *ContractsSession) RefundRequests(arg0 *big.Int) (*big.Int, error) {
	return _Contracts.Contract.RefundRequests(&_Contracts.CallOpts, arg0)
}

// RefundRequests is a free data retrieval call binding the contract method 0x7d974ade.
//
// Solidity: function refundRequests(uint256 ) view returns(uint256)
func (_Contracts *ContractsCallerSession) RefundRequests(arg0 *big.Int) (*big.Int, error) {
	return _Contracts.Contract.RefundRequests(&_Contracts.CallOpts, arg0)
}

// RequireVoucher is a free data retrieval call binding the contract method 0x66041aa1.
//
// Solidity: function requireVoucher() view returns(bool)
func (_Contracts *ContractsCaller) RequireVoucher(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "requireVoucher")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RequireVoucher is a free data retrieval call binding the contract method 0x66041aa1.
//
// Solidity: function requireVoucher() view returns(bool)
func (_Contracts *ContractsSession) RequireVoucher() (bool, error) {
	return _Contracts.Contract.RequireVoucher(&_Contracts.CallOpts)
}

// RequireVoucher is a free data retrieval call binding the contract method 0x66041aa1.
//
// Solidity: function requireVoucher() view returns(bool)
func (_Contracts *ContractsCallerSession) RequireVoucher() (bool, error) {
	return _Contracts.Contract.RequireVoucher(&_Contracts.CallOpts)
}

// TokenRewardAmount is a free data retrieval call binding the contract method 0x9c642d5d.
//
// Solidity: function tokenRewardAmount() view returns(uint256)
func (_Contracts *ContractsCaller) TokenRewardAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "tokenRewardAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenRewardAmount is a free data retrieval call binding the contract method 0x9c642d5d.
//
// Solidity: function tokenRewardAmount() view returns(uint256)
func (_Contracts *ContractsSession) TokenRewardAmount() (*big.Int, error) {
	return _Contracts.Contract.TokenRewardAmount(&_Contracts.CallOpts)
}

// TokenRewardAmount is a free data retrieval call binding the contract method 0x9c642d5d.
//
// Solidity: function tokenRewardAmount() view returns(uint256)
func (_Contracts *ContractsCallerSession) TokenRewardAmount() (*big.Int, error) {
	return _Contracts.Contract.TokenRewardAmount(&_Contracts.CallOpts)
}

// TotalValueLocked is a free data retrieval call binding the contract method 0xec18154e.
//
// Solidity: function totalValueLocked() view returns(uint256)
func (_Contracts *ContractsCaller) TotalValueLocked(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "totalValueLocked")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalValueLocked is a free data retrieval call binding the contract method 0xec18154e.
//
// Solidity: function totalValueLocked() view returns(uint256)
func (_Contracts *ContractsSession) TotalValueLocked() (*big.Int, error) {
	return _Contracts.Contract.TotalValueLocked(&_Contracts.CallOpts)
}

// TotalValueLocked is a free data retrieval call binding the contract method 0xec18154e.
//
// Solidity: function totalValueLocked() view returns(uint256)
func (_Contracts *ContractsCallerSession) TotalValueLocked() (*big.Int, error) {
	return _Contracts.Contract.TotalValueLocked(&_Contracts.CallOpts)
}

// UltraRareAddress is a free data retrieval call binding the contract method 0x56064555.
//
// Solidity: function ultraRareAddress() view returns(address)
func (_Contracts *ContractsCaller) UltraRareAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "ultraRareAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UltraRareAddress is a free data retrieval call binding the contract method 0x56064555.
//
// Solidity: function ultraRareAddress() view returns(address)
func (_Contracts *ContractsSession) UltraRareAddress() (common.Address, error) {
	return _Contracts.Contract.UltraRareAddress(&_Contracts.CallOpts)
}

// UltraRareAddress is a free data retrieval call binding the contract method 0x56064555.
//
// Solidity: function ultraRareAddress() view returns(address)
func (_Contracts *ContractsCallerSession) UltraRareAddress() (common.Address, error) {
	return _Contracts.Contract.UltraRareAddress(&_Contracts.CallOpts)
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

// VoucherAddress is a free data retrieval call binding the contract method 0x100575b2.
//
// Solidity: function voucherAddress() view returns(address)
func (_Contracts *ContractsCaller) VoucherAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "voucherAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// VoucherAddress is a free data retrieval call binding the contract method 0x100575b2.
//
// Solidity: function voucherAddress() view returns(address)
func (_Contracts *ContractsSession) VoucherAddress() (common.Address, error) {
	return _Contracts.Contract.VoucherAddress(&_Contracts.CallOpts)
}

// VoucherAddress is a free data retrieval call binding the contract method 0x100575b2.
//
// Solidity: function voucherAddress() view returns(address)
func (_Contracts *ContractsCallerSession) VoucherAddress() (common.Address, error) {
	return _Contracts.Contract.VoucherAddress(&_Contracts.CallOpts)
}

// WithdrawalAmount is a free data retrieval call binding the contract method 0x08fc3385.
//
// Solidity: function withdrawalAmount(uint256 postId) view returns(uint256)
func (_Contracts *ContractsCaller) WithdrawalAmount(opts *bind.CallOpts, postId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "withdrawalAmount", postId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawalAmount is a free data retrieval call binding the contract method 0x08fc3385.
//
// Solidity: function withdrawalAmount(uint256 postId) view returns(uint256)
func (_Contracts *ContractsSession) WithdrawalAmount(postId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.WithdrawalAmount(&_Contracts.CallOpts, postId)
}

// WithdrawalAmount is a free data retrieval call binding the contract method 0x08fc3385.
//
// Solidity: function withdrawalAmount(uint256 postId) view returns(uint256)
func (_Contracts *ContractsCallerSession) WithdrawalAmount(postId *big.Int) (*big.Int, error) {
	return _Contracts.Contract.WithdrawalAmount(&_Contracts.CallOpts, postId)
}

// BurnUltraRare is a paid mutator transaction binding the contract method 0x6a07780c.
//
// Solidity: function burnUltraRare(uint256 receiptId, string metadataURI) returns()
func (_Contracts *ContractsTransactor) BurnUltraRare(opts *bind.TransactOpts, receiptId *big.Int, metadataURI string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "burnUltraRare", receiptId, metadataURI)
}

// BurnUltraRare is a paid mutator transaction binding the contract method 0x6a07780c.
//
// Solidity: function burnUltraRare(uint256 receiptId, string metadataURI) returns()
func (_Contracts *ContractsSession) BurnUltraRare(receiptId *big.Int, metadataURI string) (*types.Transaction, error) {
	return _Contracts.Contract.BurnUltraRare(&_Contracts.TransactOpts, receiptId, metadataURI)
}

// BurnUltraRare is a paid mutator transaction binding the contract method 0x6a07780c.
//
// Solidity: function burnUltraRare(uint256 receiptId, string metadataURI) returns()
func (_Contracts *ContractsTransactorSession) BurnUltraRare(receiptId *big.Int, metadataURI string) (*types.Transaction, error) {
	return _Contracts.Contract.BurnUltraRare(&_Contracts.TransactOpts, receiptId, metadataURI)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 receiptId) returns()
func (_Contracts *ContractsTransactor) Cancel(opts *bind.TransactOpts, receiptId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "cancel", receiptId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 receiptId) returns()
func (_Contracts *ContractsSession) Cancel(receiptId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Cancel(&_Contracts.TransactOpts, receiptId)
}

// Cancel is a paid mutator transaction binding the contract method 0x40e58ee5.
//
// Solidity: function cancel(uint256 receiptId) returns()
func (_Contracts *ContractsTransactorSession) Cancel(receiptId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Cancel(&_Contracts.TransactOpts, receiptId)
}

// Donate is a paid mutator transaction binding the contract method 0x8d59d1f1.
//
// Solidity: function donate(uint256 postId, string metadataURI) payable returns()
func (_Contracts *ContractsTransactor) Donate(opts *bind.TransactOpts, postId *big.Int, metadataURI string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "donate", postId, metadataURI)
}

// Donate is a paid mutator transaction binding the contract method 0x8d59d1f1.
//
// Solidity: function donate(uint256 postId, string metadataURI) payable returns()
func (_Contracts *ContractsSession) Donate(postId *big.Int, metadataURI string) (*types.Transaction, error) {
	return _Contracts.Contract.Donate(&_Contracts.TransactOpts, postId, metadataURI)
}

// Donate is a paid mutator transaction binding the contract method 0x8d59d1f1.
//
// Solidity: function donate(uint256 postId, string metadataURI) payable returns()
func (_Contracts *ContractsTransactorSession) Donate(postId *big.Int, metadataURI string) (*types.Transaction, error) {
	return _Contracts.Contract.Donate(&_Contracts.TransactOpts, postId, metadataURI)
}

// NewPost is a paid mutator transaction binding the contract method 0x8f50b143.
//
// Solidity: function newPost(string metadataURI, uint256 capacity, uint256 periodHours) returns()
func (_Contracts *ContractsTransactor) NewPost(opts *bind.TransactOpts, metadataURI string, capacity *big.Int, periodHours *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "newPost", metadataURI, capacity, periodHours)
}

// NewPost is a paid mutator transaction binding the contract method 0x8f50b143.
//
// Solidity: function newPost(string metadataURI, uint256 capacity, uint256 periodHours) returns()
func (_Contracts *ContractsSession) NewPost(metadataURI string, capacity *big.Int, periodHours *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.NewPost(&_Contracts.TransactOpts, metadataURI, capacity, periodHours)
}

// NewPost is a paid mutator transaction binding the contract method 0x8f50b143.
//
// Solidity: function newPost(string metadataURI, uint256 capacity, uint256 periodHours) returns()
func (_Contracts *ContractsTransactorSession) NewPost(metadataURI string, capacity *big.Int, periodHours *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.NewPost(&_Contracts.TransactOpts, metadataURI, capacity, periodHours)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 receiptId) payable returns()
func (_Contracts *ContractsTransactor) Refund(opts *bind.TransactOpts, receiptId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "refund", receiptId)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 receiptId) payable returns()
func (_Contracts *ContractsSession) Refund(receiptId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Refund(&_Contracts.TransactOpts, receiptId)
}

// Refund is a paid mutator transaction binding the contract method 0x278ecde1.
//
// Solidity: function refund(uint256 receiptId) payable returns()
func (_Contracts *ContractsTransactorSession) Refund(receiptId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Refund(&_Contracts.TransactOpts, receiptId)
}

// RequestRefund is a paid mutator transaction binding the contract method 0xc5ac3d4b.
//
// Solidity: function requestRefund(uint256 receiptId, string metadataURI) returns()
func (_Contracts *ContractsTransactor) RequestRefund(opts *bind.TransactOpts, receiptId *big.Int, metadataURI string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "requestRefund", receiptId, metadataURI)
}

// RequestRefund is a paid mutator transaction binding the contract method 0xc5ac3d4b.
//
// Solidity: function requestRefund(uint256 receiptId, string metadataURI) returns()
func (_Contracts *ContractsSession) RequestRefund(receiptId *big.Int, metadataURI string) (*types.Transaction, error) {
	return _Contracts.Contract.RequestRefund(&_Contracts.TransactOpts, receiptId, metadataURI)
}

// RequestRefund is a paid mutator transaction binding the contract method 0xc5ac3d4b.
//
// Solidity: function requestRefund(uint256 receiptId, string metadataURI) returns()
func (_Contracts *ContractsTransactorSession) RequestRefund(receiptId *big.Int, metadataURI string) (*types.Transaction, error) {
	return _Contracts.Contract.RequestRefund(&_Contracts.TransactOpts, receiptId, metadataURI)
}

// SetTokenRewardAmount is a paid mutator transaction binding the contract method 0x1dbb4379.
//
// Solidity: function setTokenRewardAmount(uint256 amount) returns()
func (_Contracts *ContractsTransactor) SetTokenRewardAmount(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setTokenRewardAmount", amount)
}

// SetTokenRewardAmount is a paid mutator transaction binding the contract method 0x1dbb4379.
//
// Solidity: function setTokenRewardAmount(uint256 amount) returns()
func (_Contracts *ContractsSession) SetTokenRewardAmount(amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetTokenRewardAmount(&_Contracts.TransactOpts, amount)
}

// SetTokenRewardAmount is a paid mutator transaction binding the contract method 0x1dbb4379.
//
// Solidity: function setTokenRewardAmount(uint256 amount) returns()
func (_Contracts *ContractsTransactorSession) SetTokenRewardAmount(amount *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.SetTokenRewardAmount(&_Contracts.TransactOpts, amount)
}

// SwitchVoucherRequirement is a paid mutator transaction binding the contract method 0x294f51ff.
//
// Solidity: function switchVoucherRequirement() returns()
func (_Contracts *ContractsTransactor) SwitchVoucherRequirement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "switchVoucherRequirement")
}

// SwitchVoucherRequirement is a paid mutator transaction binding the contract method 0x294f51ff.
//
// Solidity: function switchVoucherRequirement() returns()
func (_Contracts *ContractsSession) SwitchVoucherRequirement() (*types.Transaction, error) {
	return _Contracts.Contract.SwitchVoucherRequirement(&_Contracts.TransactOpts)
}

// SwitchVoucherRequirement is a paid mutator transaction binding the contract method 0x294f51ff.
//
// Solidity: function switchVoucherRequirement() returns()
func (_Contracts *ContractsTransactorSession) SwitchVoucherRequirement() (*types.Transaction, error) {
	return _Contracts.Contract.SwitchVoucherRequirement(&_Contracts.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 postId) returns()
func (_Contracts *ContractsTransactor) Withdraw(opts *bind.TransactOpts, postId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "withdraw", postId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 postId) returns()
func (_Contracts *ContractsSession) Withdraw(postId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts, postId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 postId) returns()
func (_Contracts *ContractsTransactorSession) Withdraw(postId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.Withdraw(&_Contracts.TransactOpts, postId)
}
