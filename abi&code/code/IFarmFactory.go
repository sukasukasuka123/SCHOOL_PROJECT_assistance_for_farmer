// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package IFarmFactory

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

// IFarmFactoryMetaData contains all meta data concerning the IFarmFactory contract.
var IFarmFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"}],\"name\":\"farmExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"}],\"name\":\"gridExistsInFarm\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isFarmOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"enumGridObjectFactory.CitrusDiseaseType\",\"name\":\"newType\",\"type\":\"uint8\"}],\"name\":\"updateGridDiseaseTypeFromRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"enumGridObjectFactory.GridStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"updateGridStatusFromRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IFarmFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use IFarmFactoryMetaData.ABI instead.
var IFarmFactoryABI = IFarmFactoryMetaData.ABI

// IFarmFactory is an auto generated Go binding around an Ethereum contract.
type IFarmFactory struct {
	IFarmFactoryCaller     // Read-only binding to the contract
	IFarmFactoryTransactor // Write-only binding to the contract
	IFarmFactoryFilterer   // Log filterer for contract events
}

// IFarmFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type IFarmFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFarmFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IFarmFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFarmFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IFarmFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IFarmFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IFarmFactorySession struct {
	Contract     *IFarmFactory     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IFarmFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IFarmFactoryCallerSession struct {
	Contract *IFarmFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// IFarmFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IFarmFactoryTransactorSession struct {
	Contract     *IFarmFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// IFarmFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type IFarmFactoryRaw struct {
	Contract *IFarmFactory // Generic contract binding to access the raw methods on
}

// IFarmFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IFarmFactoryCallerRaw struct {
	Contract *IFarmFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// IFarmFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IFarmFactoryTransactorRaw struct {
	Contract *IFarmFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIFarmFactory creates a new instance of IFarmFactory, bound to a specific deployed contract.
func NewIFarmFactory(address common.Address, backend bind.ContractBackend) (*IFarmFactory, error) {
	contract, err := bindIFarmFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IFarmFactory{IFarmFactoryCaller: IFarmFactoryCaller{contract: contract}, IFarmFactoryTransactor: IFarmFactoryTransactor{contract: contract}, IFarmFactoryFilterer: IFarmFactoryFilterer{contract: contract}}, nil
}

// NewIFarmFactoryCaller creates a new read-only instance of IFarmFactory, bound to a specific deployed contract.
func NewIFarmFactoryCaller(address common.Address, caller bind.ContractCaller) (*IFarmFactoryCaller, error) {
	contract, err := bindIFarmFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IFarmFactoryCaller{contract: contract}, nil
}

// NewIFarmFactoryTransactor creates a new write-only instance of IFarmFactory, bound to a specific deployed contract.
func NewIFarmFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*IFarmFactoryTransactor, error) {
	contract, err := bindIFarmFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IFarmFactoryTransactor{contract: contract}, nil
}

// NewIFarmFactoryFilterer creates a new log filterer instance of IFarmFactory, bound to a specific deployed contract.
func NewIFarmFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*IFarmFactoryFilterer, error) {
	contract, err := bindIFarmFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IFarmFactoryFilterer{contract: contract}, nil
}

// bindIFarmFactory binds a generic wrapper to an already deployed contract.
func bindIFarmFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IFarmFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFarmFactory *IFarmFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFarmFactory.Contract.IFarmFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFarmFactory *IFarmFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFarmFactory.Contract.IFarmFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFarmFactory *IFarmFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFarmFactory.Contract.IFarmFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IFarmFactory *IFarmFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IFarmFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IFarmFactory *IFarmFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IFarmFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IFarmFactory *IFarmFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IFarmFactory.Contract.contract.Transact(opts, method, params...)
}

// FarmExists is a free data retrieval call binding the contract method 0x3460c114.
//
// Solidity: function farmExists(address farmId) view returns(bool)
func (_IFarmFactory *IFarmFactoryCaller) FarmExists(opts *bind.CallOpts, farmId common.Address) (bool, error) {
	var out []interface{}
	err := _IFarmFactory.contract.Call(opts, &out, "farmExists", farmId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FarmExists is a free data retrieval call binding the contract method 0x3460c114.
//
// Solidity: function farmExists(address farmId) view returns(bool)
func (_IFarmFactory *IFarmFactorySession) FarmExists(farmId common.Address) (bool, error) {
	return _IFarmFactory.Contract.FarmExists(&_IFarmFactory.CallOpts, farmId)
}

// FarmExists is a free data retrieval call binding the contract method 0x3460c114.
//
// Solidity: function farmExists(address farmId) view returns(bool)
func (_IFarmFactory *IFarmFactoryCallerSession) FarmExists(farmId common.Address) (bool, error) {
	return _IFarmFactory.Contract.FarmExists(&_IFarmFactory.CallOpts, farmId)
}

// GridExistsInFarm is a free data retrieval call binding the contract method 0x2cf449d0.
//
// Solidity: function gridExistsInFarm(address farmId, string gridCode) view returns(bool)
func (_IFarmFactory *IFarmFactoryCaller) GridExistsInFarm(opts *bind.CallOpts, farmId common.Address, gridCode string) (bool, error) {
	var out []interface{}
	err := _IFarmFactory.contract.Call(opts, &out, "gridExistsInFarm", farmId, gridCode)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GridExistsInFarm is a free data retrieval call binding the contract method 0x2cf449d0.
//
// Solidity: function gridExistsInFarm(address farmId, string gridCode) view returns(bool)
func (_IFarmFactory *IFarmFactorySession) GridExistsInFarm(farmId common.Address, gridCode string) (bool, error) {
	return _IFarmFactory.Contract.GridExistsInFarm(&_IFarmFactory.CallOpts, farmId, gridCode)
}

// GridExistsInFarm is a free data retrieval call binding the contract method 0x2cf449d0.
//
// Solidity: function gridExistsInFarm(address farmId, string gridCode) view returns(bool)
func (_IFarmFactory *IFarmFactoryCallerSession) GridExistsInFarm(farmId common.Address, gridCode string) (bool, error) {
	return _IFarmFactory.Contract.GridExistsInFarm(&_IFarmFactory.CallOpts, farmId, gridCode)
}

// IsFarmOwner is a free data retrieval call binding the contract method 0xebb1115d.
//
// Solidity: function isFarmOwner(address farmId, address account) view returns(bool)
func (_IFarmFactory *IFarmFactoryCaller) IsFarmOwner(opts *bind.CallOpts, farmId common.Address, account common.Address) (bool, error) {
	var out []interface{}
	err := _IFarmFactory.contract.Call(opts, &out, "isFarmOwner", farmId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFarmOwner is a free data retrieval call binding the contract method 0xebb1115d.
//
// Solidity: function isFarmOwner(address farmId, address account) view returns(bool)
func (_IFarmFactory *IFarmFactorySession) IsFarmOwner(farmId common.Address, account common.Address) (bool, error) {
	return _IFarmFactory.Contract.IsFarmOwner(&_IFarmFactory.CallOpts, farmId, account)
}

// IsFarmOwner is a free data retrieval call binding the contract method 0xebb1115d.
//
// Solidity: function isFarmOwner(address farmId, address account) view returns(bool)
func (_IFarmFactory *IFarmFactoryCallerSession) IsFarmOwner(farmId common.Address, account common.Address) (bool, error) {
	return _IFarmFactory.Contract.IsFarmOwner(&_IFarmFactory.CallOpts, farmId, account)
}

// UpdateGridDiseaseTypeFromRecord is a paid mutator transaction binding the contract method 0x09c76654.
//
// Solidity: function updateGridDiseaseTypeFromRecord(address farmId, string gridCode, uint8 newType) returns()
func (_IFarmFactory *IFarmFactoryTransactor) UpdateGridDiseaseTypeFromRecord(opts *bind.TransactOpts, farmId common.Address, gridCode string, newType uint8) (*types.Transaction, error) {
	return _IFarmFactory.contract.Transact(opts, "updateGridDiseaseTypeFromRecord", farmId, gridCode, newType)
}

// UpdateGridDiseaseTypeFromRecord is a paid mutator transaction binding the contract method 0x09c76654.
//
// Solidity: function updateGridDiseaseTypeFromRecord(address farmId, string gridCode, uint8 newType) returns()
func (_IFarmFactory *IFarmFactorySession) UpdateGridDiseaseTypeFromRecord(farmId common.Address, gridCode string, newType uint8) (*types.Transaction, error) {
	return _IFarmFactory.Contract.UpdateGridDiseaseTypeFromRecord(&_IFarmFactory.TransactOpts, farmId, gridCode, newType)
}

// UpdateGridDiseaseTypeFromRecord is a paid mutator transaction binding the contract method 0x09c76654.
//
// Solidity: function updateGridDiseaseTypeFromRecord(address farmId, string gridCode, uint8 newType) returns()
func (_IFarmFactory *IFarmFactoryTransactorSession) UpdateGridDiseaseTypeFromRecord(farmId common.Address, gridCode string, newType uint8) (*types.Transaction, error) {
	return _IFarmFactory.Contract.UpdateGridDiseaseTypeFromRecord(&_IFarmFactory.TransactOpts, farmId, gridCode, newType)
}

// UpdateGridStatusFromRecord is a paid mutator transaction binding the contract method 0x38ff2222.
//
// Solidity: function updateGridStatusFromRecord(address farmId, string gridCode, uint8 newStatus) returns()
func (_IFarmFactory *IFarmFactoryTransactor) UpdateGridStatusFromRecord(opts *bind.TransactOpts, farmId common.Address, gridCode string, newStatus uint8) (*types.Transaction, error) {
	return _IFarmFactory.contract.Transact(opts, "updateGridStatusFromRecord", farmId, gridCode, newStatus)
}

// UpdateGridStatusFromRecord is a paid mutator transaction binding the contract method 0x38ff2222.
//
// Solidity: function updateGridStatusFromRecord(address farmId, string gridCode, uint8 newStatus) returns()
func (_IFarmFactory *IFarmFactorySession) UpdateGridStatusFromRecord(farmId common.Address, gridCode string, newStatus uint8) (*types.Transaction, error) {
	return _IFarmFactory.Contract.UpdateGridStatusFromRecord(&_IFarmFactory.TransactOpts, farmId, gridCode, newStatus)
}

// UpdateGridStatusFromRecord is a paid mutator transaction binding the contract method 0x38ff2222.
//
// Solidity: function updateGridStatusFromRecord(address farmId, string gridCode, uint8 newStatus) returns()
func (_IFarmFactory *IFarmFactoryTransactorSession) UpdateGridStatusFromRecord(farmId common.Address, gridCode string, newStatus uint8) (*types.Transaction, error) {
	return _IFarmFactory.Contract.UpdateGridStatusFromRecord(&_IFarmFactory.TransactOpts, farmId, gridCode, newStatus)
}
