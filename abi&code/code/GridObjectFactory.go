// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GridObjectFactory

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

// GridObjectFactoryGridInfo is an auto generated low-level Go binding around an user-defined struct.
type GridObjectFactoryGridInfo struct {
	GridId               common.Address
	GridCode             string
	FarmId               common.Address
	TerrainType          uint8
	Status               uint8
	DiseaseType          uint8
	LastStatusUpdateTime uint32
}

// GridObjectFactoryGridMaintenanceRecord is an auto generated low-level Go binding around an user-defined struct.
type GridObjectFactoryGridMaintenanceRecord struct {
	RecordId      uint32
	Timestamp     uint32
	OperationType string
	Detail        string
	EvidenceIPFS  string
	Operator      common.Address
}

// GridObjectFactoryMetaData contains all meta data concerning the GridObjectFactory contract.
var GridObjectFactoryMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"}],\"name\":\"FarmGridRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumGridObjectFactory.CitrusDiseaseType\",\"name\":\"oldType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumGridObjectFactory.CitrusDiseaseType\",\"name\":\"newType\",\"type\":\"uint8\"}],\"name\":\"GridDiseaseTypeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"recordId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"operationType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"GridMaintenanceRecorded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumGridObjectFactory.GridStatus\",\"name\":\"oldStatus\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumGridObjectFactory.GridStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"GridStatusUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"recordId\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"operationType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"detail\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"evidenceIPFS\",\"type\":\"string\"}],\"name\":\"addMaintenanceRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"}],\"name\":\"getGrid\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"enumGridObjectFactory.GridTerrainType\",\"name\":\"terrainType\",\"type\":\"uint8\"},{\"internalType\":\"enumGridObjectFactory.GridStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"enumGridObjectFactory.CitrusDiseaseType\",\"name\":\"diseaseType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"lastStatusUpdateTime\",\"type\":\"uint32\"}],\"internalType\":\"structGridObjectFactory.GridInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getMaintenanceRecordByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"recordId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"operationType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"detail\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"evidenceIPFS\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"internalType\":\"structGridObjectFactory.GridMaintenanceRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"}],\"name\":\"getMaintenanceRecordCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"}],\"name\":\"getMaintenanceRecords\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"recordId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"operationType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"detail\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"evidenceIPFS\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"internalType\":\"structGridObjectFactory.GridMaintenanceRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"grids\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"enumGridObjectFactory.GridTerrainType\",\"name\":\"terrainType\",\"type\":\"uint8\"},{\"internalType\":\"enumGridObjectFactory.GridStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"enumGridObjectFactory.CitrusDiseaseType\",\"name\":\"diseaseType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"lastStatusUpdateTime\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"maintenanceRecords\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"recordId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"operationType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"detail\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"evidenceIPFS\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"enumGridObjectFactory.GridTerrainType\",\"name\":\"terrainType\",\"type\":\"uint8\"}],\"name\":\"registerGrid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"internalType\":\"enumGridObjectFactory.CitrusDiseaseType\",\"name\":\"newType\",\"type\":\"uint8\"}],\"name\":\"updateDiseaseType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"internalType\":\"enumGridObjectFactory.GridStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"updateGridStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// GridObjectFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use GridObjectFactoryMetaData.ABI instead.
var GridObjectFactoryABI = GridObjectFactoryMetaData.ABI

// GridObjectFactory is an auto generated Go binding around an Ethereum contract.
type GridObjectFactory struct {
	GridObjectFactoryCaller     // Read-only binding to the contract
	GridObjectFactoryTransactor // Write-only binding to the contract
	GridObjectFactoryFilterer   // Log filterer for contract events
}

// GridObjectFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type GridObjectFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GridObjectFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GridObjectFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GridObjectFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GridObjectFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GridObjectFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GridObjectFactorySession struct {
	Contract     *GridObjectFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GridObjectFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GridObjectFactoryCallerSession struct {
	Contract *GridObjectFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// GridObjectFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GridObjectFactoryTransactorSession struct {
	Contract     *GridObjectFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// GridObjectFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type GridObjectFactoryRaw struct {
	Contract *GridObjectFactory // Generic contract binding to access the raw methods on
}

// GridObjectFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GridObjectFactoryCallerRaw struct {
	Contract *GridObjectFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// GridObjectFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GridObjectFactoryTransactorRaw struct {
	Contract *GridObjectFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGridObjectFactory creates a new instance of GridObjectFactory, bound to a specific deployed contract.
func NewGridObjectFactory(address common.Address, backend bind.ContractBackend) (*GridObjectFactory, error) {
	contract, err := bindGridObjectFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GridObjectFactory{GridObjectFactoryCaller: GridObjectFactoryCaller{contract: contract}, GridObjectFactoryTransactor: GridObjectFactoryTransactor{contract: contract}, GridObjectFactoryFilterer: GridObjectFactoryFilterer{contract: contract}}, nil
}

// NewGridObjectFactoryCaller creates a new read-only instance of GridObjectFactory, bound to a specific deployed contract.
func NewGridObjectFactoryCaller(address common.Address, caller bind.ContractCaller) (*GridObjectFactoryCaller, error) {
	contract, err := bindGridObjectFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GridObjectFactoryCaller{contract: contract}, nil
}

// NewGridObjectFactoryTransactor creates a new write-only instance of GridObjectFactory, bound to a specific deployed contract.
func NewGridObjectFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*GridObjectFactoryTransactor, error) {
	contract, err := bindGridObjectFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GridObjectFactoryTransactor{contract: contract}, nil
}

// NewGridObjectFactoryFilterer creates a new log filterer instance of GridObjectFactory, bound to a specific deployed contract.
func NewGridObjectFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*GridObjectFactoryFilterer, error) {
	contract, err := bindGridObjectFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GridObjectFactoryFilterer{contract: contract}, nil
}

// bindGridObjectFactory binds a generic wrapper to an already deployed contract.
func bindGridObjectFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GridObjectFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GridObjectFactory *GridObjectFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GridObjectFactory.Contract.GridObjectFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GridObjectFactory *GridObjectFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GridObjectFactory.Contract.GridObjectFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GridObjectFactory *GridObjectFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GridObjectFactory.Contract.GridObjectFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GridObjectFactory *GridObjectFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GridObjectFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GridObjectFactory *GridObjectFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GridObjectFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GridObjectFactory *GridObjectFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GridObjectFactory.Contract.contract.Transact(opts, method, params...)
}

// GetGrid is a free data retrieval call binding the contract method 0x4e3fb04a.
//
// Solidity: function getGrid(address gridId) view returns((address,string,address,uint8,uint8,uint8,uint32))
func (_GridObjectFactory *GridObjectFactoryCaller) GetGrid(opts *bind.CallOpts, gridId common.Address) (GridObjectFactoryGridInfo, error) {
	var out []interface{}
	err := _GridObjectFactory.contract.Call(opts, &out, "getGrid", gridId)

	if err != nil {
		return *new(GridObjectFactoryGridInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(GridObjectFactoryGridInfo)).(*GridObjectFactoryGridInfo)

	return out0, err

}

// GetGrid is a free data retrieval call binding the contract method 0x4e3fb04a.
//
// Solidity: function getGrid(address gridId) view returns((address,string,address,uint8,uint8,uint8,uint32))
func (_GridObjectFactory *GridObjectFactorySession) GetGrid(gridId common.Address) (GridObjectFactoryGridInfo, error) {
	return _GridObjectFactory.Contract.GetGrid(&_GridObjectFactory.CallOpts, gridId)
}

// GetGrid is a free data retrieval call binding the contract method 0x4e3fb04a.
//
// Solidity: function getGrid(address gridId) view returns((address,string,address,uint8,uint8,uint8,uint32))
func (_GridObjectFactory *GridObjectFactoryCallerSession) GetGrid(gridId common.Address) (GridObjectFactoryGridInfo, error) {
	return _GridObjectFactory.Contract.GetGrid(&_GridObjectFactory.CallOpts, gridId)
}

// GetMaintenanceRecordByIndex is a free data retrieval call binding the contract method 0xdf6e1f4c.
//
// Solidity: function getMaintenanceRecordByIndex(address gridId, uint256 index) view returns((uint32,uint32,string,string,string,address))
func (_GridObjectFactory *GridObjectFactoryCaller) GetMaintenanceRecordByIndex(opts *bind.CallOpts, gridId common.Address, index *big.Int) (GridObjectFactoryGridMaintenanceRecord, error) {
	var out []interface{}
	err := _GridObjectFactory.contract.Call(opts, &out, "getMaintenanceRecordByIndex", gridId, index)

	if err != nil {
		return *new(GridObjectFactoryGridMaintenanceRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(GridObjectFactoryGridMaintenanceRecord)).(*GridObjectFactoryGridMaintenanceRecord)

	return out0, err

}

// GetMaintenanceRecordByIndex is a free data retrieval call binding the contract method 0xdf6e1f4c.
//
// Solidity: function getMaintenanceRecordByIndex(address gridId, uint256 index) view returns((uint32,uint32,string,string,string,address))
func (_GridObjectFactory *GridObjectFactorySession) GetMaintenanceRecordByIndex(gridId common.Address, index *big.Int) (GridObjectFactoryGridMaintenanceRecord, error) {
	return _GridObjectFactory.Contract.GetMaintenanceRecordByIndex(&_GridObjectFactory.CallOpts, gridId, index)
}

// GetMaintenanceRecordByIndex is a free data retrieval call binding the contract method 0xdf6e1f4c.
//
// Solidity: function getMaintenanceRecordByIndex(address gridId, uint256 index) view returns((uint32,uint32,string,string,string,address))
func (_GridObjectFactory *GridObjectFactoryCallerSession) GetMaintenanceRecordByIndex(gridId common.Address, index *big.Int) (GridObjectFactoryGridMaintenanceRecord, error) {
	return _GridObjectFactory.Contract.GetMaintenanceRecordByIndex(&_GridObjectFactory.CallOpts, gridId, index)
}

// GetMaintenanceRecordCount is a free data retrieval call binding the contract method 0x7629699b.
//
// Solidity: function getMaintenanceRecordCount(address gridId) view returns(uint256)
func (_GridObjectFactory *GridObjectFactoryCaller) GetMaintenanceRecordCount(opts *bind.CallOpts, gridId common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GridObjectFactory.contract.Call(opts, &out, "getMaintenanceRecordCount", gridId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaintenanceRecordCount is a free data retrieval call binding the contract method 0x7629699b.
//
// Solidity: function getMaintenanceRecordCount(address gridId) view returns(uint256)
func (_GridObjectFactory *GridObjectFactorySession) GetMaintenanceRecordCount(gridId common.Address) (*big.Int, error) {
	return _GridObjectFactory.Contract.GetMaintenanceRecordCount(&_GridObjectFactory.CallOpts, gridId)
}

// GetMaintenanceRecordCount is a free data retrieval call binding the contract method 0x7629699b.
//
// Solidity: function getMaintenanceRecordCount(address gridId) view returns(uint256)
func (_GridObjectFactory *GridObjectFactoryCallerSession) GetMaintenanceRecordCount(gridId common.Address) (*big.Int, error) {
	return _GridObjectFactory.Contract.GetMaintenanceRecordCount(&_GridObjectFactory.CallOpts, gridId)
}

// GetMaintenanceRecords is a free data retrieval call binding the contract method 0xf67bc1dd.
//
// Solidity: function getMaintenanceRecords(address gridId) view returns((uint32,uint32,string,string,string,address)[])
func (_GridObjectFactory *GridObjectFactoryCaller) GetMaintenanceRecords(opts *bind.CallOpts, gridId common.Address) ([]GridObjectFactoryGridMaintenanceRecord, error) {
	var out []interface{}
	err := _GridObjectFactory.contract.Call(opts, &out, "getMaintenanceRecords", gridId)

	if err != nil {
		return *new([]GridObjectFactoryGridMaintenanceRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]GridObjectFactoryGridMaintenanceRecord)).(*[]GridObjectFactoryGridMaintenanceRecord)

	return out0, err

}

// GetMaintenanceRecords is a free data retrieval call binding the contract method 0xf67bc1dd.
//
// Solidity: function getMaintenanceRecords(address gridId) view returns((uint32,uint32,string,string,string,address)[])
func (_GridObjectFactory *GridObjectFactorySession) GetMaintenanceRecords(gridId common.Address) ([]GridObjectFactoryGridMaintenanceRecord, error) {
	return _GridObjectFactory.Contract.GetMaintenanceRecords(&_GridObjectFactory.CallOpts, gridId)
}

// GetMaintenanceRecords is a free data retrieval call binding the contract method 0xf67bc1dd.
//
// Solidity: function getMaintenanceRecords(address gridId) view returns((uint32,uint32,string,string,string,address)[])
func (_GridObjectFactory *GridObjectFactoryCallerSession) GetMaintenanceRecords(gridId common.Address) ([]GridObjectFactoryGridMaintenanceRecord, error) {
	return _GridObjectFactory.Contract.GetMaintenanceRecords(&_GridObjectFactory.CallOpts, gridId)
}

// Grids is a free data retrieval call binding the contract method 0xadbb79cd.
//
// Solidity: function grids(address ) view returns(address gridId, string gridCode, address farmId, uint8 terrainType, uint8 status, uint8 diseaseType, uint32 lastStatusUpdateTime)
func (_GridObjectFactory *GridObjectFactoryCaller) Grids(opts *bind.CallOpts, arg0 common.Address) (struct {
	GridId               common.Address
	GridCode             string
	FarmId               common.Address
	TerrainType          uint8
	Status               uint8
	DiseaseType          uint8
	LastStatusUpdateTime uint32
}, error) {
	var out []interface{}
	err := _GridObjectFactory.contract.Call(opts, &out, "grids", arg0)

	outstruct := new(struct {
		GridId               common.Address
		GridCode             string
		FarmId               common.Address
		TerrainType          uint8
		Status               uint8
		DiseaseType          uint8
		LastStatusUpdateTime uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GridId = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.GridCode = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.FarmId = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.TerrainType = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.Status = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.DiseaseType = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.LastStatusUpdateTime = *abi.ConvertType(out[6], new(uint32)).(*uint32)

	return *outstruct, err

}

// Grids is a free data retrieval call binding the contract method 0xadbb79cd.
//
// Solidity: function grids(address ) view returns(address gridId, string gridCode, address farmId, uint8 terrainType, uint8 status, uint8 diseaseType, uint32 lastStatusUpdateTime)
func (_GridObjectFactory *GridObjectFactorySession) Grids(arg0 common.Address) (struct {
	GridId               common.Address
	GridCode             string
	FarmId               common.Address
	TerrainType          uint8
	Status               uint8
	DiseaseType          uint8
	LastStatusUpdateTime uint32
}, error) {
	return _GridObjectFactory.Contract.Grids(&_GridObjectFactory.CallOpts, arg0)
}

// Grids is a free data retrieval call binding the contract method 0xadbb79cd.
//
// Solidity: function grids(address ) view returns(address gridId, string gridCode, address farmId, uint8 terrainType, uint8 status, uint8 diseaseType, uint32 lastStatusUpdateTime)
func (_GridObjectFactory *GridObjectFactoryCallerSession) Grids(arg0 common.Address) (struct {
	GridId               common.Address
	GridCode             string
	FarmId               common.Address
	TerrainType          uint8
	Status               uint8
	DiseaseType          uint8
	LastStatusUpdateTime uint32
}, error) {
	return _GridObjectFactory.Contract.Grids(&_GridObjectFactory.CallOpts, arg0)
}

// MaintenanceRecords is a free data retrieval call binding the contract method 0xa2cc6d5a.
//
// Solidity: function maintenanceRecords(address , uint256 ) view returns(uint32 recordId, uint32 timestamp, string operationType, string detail, string evidenceIPFS, address operator)
func (_GridObjectFactory *GridObjectFactoryCaller) MaintenanceRecords(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	RecordId      uint32
	Timestamp     uint32
	OperationType string
	Detail        string
	EvidenceIPFS  string
	Operator      common.Address
}, error) {
	var out []interface{}
	err := _GridObjectFactory.contract.Call(opts, &out, "maintenanceRecords", arg0, arg1)

	outstruct := new(struct {
		RecordId      uint32
		Timestamp     uint32
		OperationType string
		Detail        string
		EvidenceIPFS  string
		Operator      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RecordId = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.OperationType = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Detail = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.EvidenceIPFS = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Operator = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// MaintenanceRecords is a free data retrieval call binding the contract method 0xa2cc6d5a.
//
// Solidity: function maintenanceRecords(address , uint256 ) view returns(uint32 recordId, uint32 timestamp, string operationType, string detail, string evidenceIPFS, address operator)
func (_GridObjectFactory *GridObjectFactorySession) MaintenanceRecords(arg0 common.Address, arg1 *big.Int) (struct {
	RecordId      uint32
	Timestamp     uint32
	OperationType string
	Detail        string
	EvidenceIPFS  string
	Operator      common.Address
}, error) {
	return _GridObjectFactory.Contract.MaintenanceRecords(&_GridObjectFactory.CallOpts, arg0, arg1)
}

// MaintenanceRecords is a free data retrieval call binding the contract method 0xa2cc6d5a.
//
// Solidity: function maintenanceRecords(address , uint256 ) view returns(uint32 recordId, uint32 timestamp, string operationType, string detail, string evidenceIPFS, address operator)
func (_GridObjectFactory *GridObjectFactoryCallerSession) MaintenanceRecords(arg0 common.Address, arg1 *big.Int) (struct {
	RecordId      uint32
	Timestamp     uint32
	OperationType string
	Detail        string
	EvidenceIPFS  string
	Operator      common.Address
}, error) {
	return _GridObjectFactory.Contract.MaintenanceRecords(&_GridObjectFactory.CallOpts, arg0, arg1)
}

// AddMaintenanceRecord is a paid mutator transaction binding the contract method 0x1c426167.
//
// Solidity: function addMaintenanceRecord(address gridId, uint32 recordId, string operationType, string detail, string evidenceIPFS) returns()
func (_GridObjectFactory *GridObjectFactoryTransactor) AddMaintenanceRecord(opts *bind.TransactOpts, gridId common.Address, recordId uint32, operationType string, detail string, evidenceIPFS string) (*types.Transaction, error) {
	return _GridObjectFactory.contract.Transact(opts, "addMaintenanceRecord", gridId, recordId, operationType, detail, evidenceIPFS)
}

// AddMaintenanceRecord is a paid mutator transaction binding the contract method 0x1c426167.
//
// Solidity: function addMaintenanceRecord(address gridId, uint32 recordId, string operationType, string detail, string evidenceIPFS) returns()
func (_GridObjectFactory *GridObjectFactorySession) AddMaintenanceRecord(gridId common.Address, recordId uint32, operationType string, detail string, evidenceIPFS string) (*types.Transaction, error) {
	return _GridObjectFactory.Contract.AddMaintenanceRecord(&_GridObjectFactory.TransactOpts, gridId, recordId, operationType, detail, evidenceIPFS)
}

// AddMaintenanceRecord is a paid mutator transaction binding the contract method 0x1c426167.
//
// Solidity: function addMaintenanceRecord(address gridId, uint32 recordId, string operationType, string detail, string evidenceIPFS) returns()
func (_GridObjectFactory *GridObjectFactoryTransactorSession) AddMaintenanceRecord(gridId common.Address, recordId uint32, operationType string, detail string, evidenceIPFS string) (*types.Transaction, error) {
	return _GridObjectFactory.Contract.AddMaintenanceRecord(&_GridObjectFactory.TransactOpts, gridId, recordId, operationType, detail, evidenceIPFS)
}

// RegisterGrid is a paid mutator transaction binding the contract method 0x39b85f7f.
//
// Solidity: function registerGrid(address gridId, string gridCode, address farmId, uint8 terrainType) returns()
func (_GridObjectFactory *GridObjectFactoryTransactor) RegisterGrid(opts *bind.TransactOpts, gridId common.Address, gridCode string, farmId common.Address, terrainType uint8) (*types.Transaction, error) {
	return _GridObjectFactory.contract.Transact(opts, "registerGrid", gridId, gridCode, farmId, terrainType)
}

// RegisterGrid is a paid mutator transaction binding the contract method 0x39b85f7f.
//
// Solidity: function registerGrid(address gridId, string gridCode, address farmId, uint8 terrainType) returns()
func (_GridObjectFactory *GridObjectFactorySession) RegisterGrid(gridId common.Address, gridCode string, farmId common.Address, terrainType uint8) (*types.Transaction, error) {
	return _GridObjectFactory.Contract.RegisterGrid(&_GridObjectFactory.TransactOpts, gridId, gridCode, farmId, terrainType)
}

// RegisterGrid is a paid mutator transaction binding the contract method 0x39b85f7f.
//
// Solidity: function registerGrid(address gridId, string gridCode, address farmId, uint8 terrainType) returns()
func (_GridObjectFactory *GridObjectFactoryTransactorSession) RegisterGrid(gridId common.Address, gridCode string, farmId common.Address, terrainType uint8) (*types.Transaction, error) {
	return _GridObjectFactory.Contract.RegisterGrid(&_GridObjectFactory.TransactOpts, gridId, gridCode, farmId, terrainType)
}

// UpdateDiseaseType is a paid mutator transaction binding the contract method 0x9e888917.
//
// Solidity: function updateDiseaseType(address gridId, uint8 newType) returns()
func (_GridObjectFactory *GridObjectFactoryTransactor) UpdateDiseaseType(opts *bind.TransactOpts, gridId common.Address, newType uint8) (*types.Transaction, error) {
	return _GridObjectFactory.contract.Transact(opts, "updateDiseaseType", gridId, newType)
}

// UpdateDiseaseType is a paid mutator transaction binding the contract method 0x9e888917.
//
// Solidity: function updateDiseaseType(address gridId, uint8 newType) returns()
func (_GridObjectFactory *GridObjectFactorySession) UpdateDiseaseType(gridId common.Address, newType uint8) (*types.Transaction, error) {
	return _GridObjectFactory.Contract.UpdateDiseaseType(&_GridObjectFactory.TransactOpts, gridId, newType)
}

// UpdateDiseaseType is a paid mutator transaction binding the contract method 0x9e888917.
//
// Solidity: function updateDiseaseType(address gridId, uint8 newType) returns()
func (_GridObjectFactory *GridObjectFactoryTransactorSession) UpdateDiseaseType(gridId common.Address, newType uint8) (*types.Transaction, error) {
	return _GridObjectFactory.Contract.UpdateDiseaseType(&_GridObjectFactory.TransactOpts, gridId, newType)
}

// UpdateGridStatus is a paid mutator transaction binding the contract method 0xc829f03b.
//
// Solidity: function updateGridStatus(address gridId, uint8 newStatus) returns()
func (_GridObjectFactory *GridObjectFactoryTransactor) UpdateGridStatus(opts *bind.TransactOpts, gridId common.Address, newStatus uint8) (*types.Transaction, error) {
	return _GridObjectFactory.contract.Transact(opts, "updateGridStatus", gridId, newStatus)
}

// UpdateGridStatus is a paid mutator transaction binding the contract method 0xc829f03b.
//
// Solidity: function updateGridStatus(address gridId, uint8 newStatus) returns()
func (_GridObjectFactory *GridObjectFactorySession) UpdateGridStatus(gridId common.Address, newStatus uint8) (*types.Transaction, error) {
	return _GridObjectFactory.Contract.UpdateGridStatus(&_GridObjectFactory.TransactOpts, gridId, newStatus)
}

// UpdateGridStatus is a paid mutator transaction binding the contract method 0xc829f03b.
//
// Solidity: function updateGridStatus(address gridId, uint8 newStatus) returns()
func (_GridObjectFactory *GridObjectFactoryTransactorSession) UpdateGridStatus(gridId common.Address, newStatus uint8) (*types.Transaction, error) {
	return _GridObjectFactory.Contract.UpdateGridStatus(&_GridObjectFactory.TransactOpts, gridId, newStatus)
}

// GridObjectFactoryFarmGridRegisteredIterator is returned from FilterFarmGridRegistered and is used to iterate over the raw logs and unpacked data for FarmGridRegistered events raised by the GridObjectFactory contract.
type GridObjectFactoryFarmGridRegisteredIterator struct {
	Event *GridObjectFactoryFarmGridRegistered // Event containing the contract specifics and raw log

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
func (it *GridObjectFactoryFarmGridRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GridObjectFactoryFarmGridRegistered)
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
		it.Event = new(GridObjectFactoryFarmGridRegistered)
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
func (it *GridObjectFactoryFarmGridRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GridObjectFactoryFarmGridRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GridObjectFactoryFarmGridRegistered represents a FarmGridRegistered event raised by the GridObjectFactory contract.
type GridObjectFactoryFarmGridRegistered struct {
	FarmId   common.Address
	GridId   common.Address
	GridCode string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFarmGridRegistered is a free log retrieval operation binding the contract event 0xb4043ad9452e8039c88af7ed262dbb5e8afd1ecd023590a82eb223bf40a10668.
//
// Solidity: event FarmGridRegistered(address indexed farmId, address indexed gridId, string gridCode)
func (_GridObjectFactory *GridObjectFactoryFilterer) FilterFarmGridRegistered(opts *bind.FilterOpts, farmId []common.Address, gridId []common.Address) (*GridObjectFactoryFarmGridRegisteredIterator, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}
	var gridIdRule []interface{}
	for _, gridIdItem := range gridId {
		gridIdRule = append(gridIdRule, gridIdItem)
	}

	logs, sub, err := _GridObjectFactory.contract.FilterLogs(opts, "FarmGridRegistered", farmIdRule, gridIdRule)
	if err != nil {
		return nil, err
	}
	return &GridObjectFactoryFarmGridRegisteredIterator{contract: _GridObjectFactory.contract, event: "FarmGridRegistered", logs: logs, sub: sub}, nil
}

// WatchFarmGridRegistered is a free log subscription operation binding the contract event 0xb4043ad9452e8039c88af7ed262dbb5e8afd1ecd023590a82eb223bf40a10668.
//
// Solidity: event FarmGridRegistered(address indexed farmId, address indexed gridId, string gridCode)
func (_GridObjectFactory *GridObjectFactoryFilterer) WatchFarmGridRegistered(opts *bind.WatchOpts, sink chan<- *GridObjectFactoryFarmGridRegistered, farmId []common.Address, gridId []common.Address) (event.Subscription, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}
	var gridIdRule []interface{}
	for _, gridIdItem := range gridId {
		gridIdRule = append(gridIdRule, gridIdItem)
	}

	logs, sub, err := _GridObjectFactory.contract.WatchLogs(opts, "FarmGridRegistered", farmIdRule, gridIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GridObjectFactoryFarmGridRegistered)
				if err := _GridObjectFactory.contract.UnpackLog(event, "FarmGridRegistered", log); err != nil {
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

// ParseFarmGridRegistered is a log parse operation binding the contract event 0xb4043ad9452e8039c88af7ed262dbb5e8afd1ecd023590a82eb223bf40a10668.
//
// Solidity: event FarmGridRegistered(address indexed farmId, address indexed gridId, string gridCode)
func (_GridObjectFactory *GridObjectFactoryFilterer) ParseFarmGridRegistered(log types.Log) (*GridObjectFactoryFarmGridRegistered, error) {
	event := new(GridObjectFactoryFarmGridRegistered)
	if err := _GridObjectFactory.contract.UnpackLog(event, "FarmGridRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GridObjectFactoryGridDiseaseTypeUpdatedIterator is returned from FilterGridDiseaseTypeUpdated and is used to iterate over the raw logs and unpacked data for GridDiseaseTypeUpdated events raised by the GridObjectFactory contract.
type GridObjectFactoryGridDiseaseTypeUpdatedIterator struct {
	Event *GridObjectFactoryGridDiseaseTypeUpdated // Event containing the contract specifics and raw log

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
func (it *GridObjectFactoryGridDiseaseTypeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GridObjectFactoryGridDiseaseTypeUpdated)
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
		it.Event = new(GridObjectFactoryGridDiseaseTypeUpdated)
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
func (it *GridObjectFactoryGridDiseaseTypeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GridObjectFactoryGridDiseaseTypeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GridObjectFactoryGridDiseaseTypeUpdated represents a GridDiseaseTypeUpdated event raised by the GridObjectFactory contract.
type GridObjectFactoryGridDiseaseTypeUpdated struct {
	GridId  common.Address
	OldType uint8
	NewType uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGridDiseaseTypeUpdated is a free log retrieval operation binding the contract event 0x6a66c323c1c849a4d73387adf181d7c3e57a60b1486debb5ac11cd47c2b9d095.
//
// Solidity: event GridDiseaseTypeUpdated(address indexed gridId, uint8 oldType, uint8 newType)
func (_GridObjectFactory *GridObjectFactoryFilterer) FilterGridDiseaseTypeUpdated(opts *bind.FilterOpts, gridId []common.Address) (*GridObjectFactoryGridDiseaseTypeUpdatedIterator, error) {

	var gridIdRule []interface{}
	for _, gridIdItem := range gridId {
		gridIdRule = append(gridIdRule, gridIdItem)
	}

	logs, sub, err := _GridObjectFactory.contract.FilterLogs(opts, "GridDiseaseTypeUpdated", gridIdRule)
	if err != nil {
		return nil, err
	}
	return &GridObjectFactoryGridDiseaseTypeUpdatedIterator{contract: _GridObjectFactory.contract, event: "GridDiseaseTypeUpdated", logs: logs, sub: sub}, nil
}

// WatchGridDiseaseTypeUpdated is a free log subscription operation binding the contract event 0x6a66c323c1c849a4d73387adf181d7c3e57a60b1486debb5ac11cd47c2b9d095.
//
// Solidity: event GridDiseaseTypeUpdated(address indexed gridId, uint8 oldType, uint8 newType)
func (_GridObjectFactory *GridObjectFactoryFilterer) WatchGridDiseaseTypeUpdated(opts *bind.WatchOpts, sink chan<- *GridObjectFactoryGridDiseaseTypeUpdated, gridId []common.Address) (event.Subscription, error) {

	var gridIdRule []interface{}
	for _, gridIdItem := range gridId {
		gridIdRule = append(gridIdRule, gridIdItem)
	}

	logs, sub, err := _GridObjectFactory.contract.WatchLogs(opts, "GridDiseaseTypeUpdated", gridIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GridObjectFactoryGridDiseaseTypeUpdated)
				if err := _GridObjectFactory.contract.UnpackLog(event, "GridDiseaseTypeUpdated", log); err != nil {
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

// ParseGridDiseaseTypeUpdated is a log parse operation binding the contract event 0x6a66c323c1c849a4d73387adf181d7c3e57a60b1486debb5ac11cd47c2b9d095.
//
// Solidity: event GridDiseaseTypeUpdated(address indexed gridId, uint8 oldType, uint8 newType)
func (_GridObjectFactory *GridObjectFactoryFilterer) ParseGridDiseaseTypeUpdated(log types.Log) (*GridObjectFactoryGridDiseaseTypeUpdated, error) {
	event := new(GridObjectFactoryGridDiseaseTypeUpdated)
	if err := _GridObjectFactory.contract.UnpackLog(event, "GridDiseaseTypeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GridObjectFactoryGridMaintenanceRecordedIterator is returned from FilterGridMaintenanceRecorded and is used to iterate over the raw logs and unpacked data for GridMaintenanceRecorded events raised by the GridObjectFactory contract.
type GridObjectFactoryGridMaintenanceRecordedIterator struct {
	Event *GridObjectFactoryGridMaintenanceRecorded // Event containing the contract specifics and raw log

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
func (it *GridObjectFactoryGridMaintenanceRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GridObjectFactoryGridMaintenanceRecorded)
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
		it.Event = new(GridObjectFactoryGridMaintenanceRecorded)
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
func (it *GridObjectFactoryGridMaintenanceRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GridObjectFactoryGridMaintenanceRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GridObjectFactoryGridMaintenanceRecorded represents a GridMaintenanceRecorded event raised by the GridObjectFactory contract.
type GridObjectFactoryGridMaintenanceRecorded struct {
	GridId        common.Address
	RecordId      uint32
	OperationType string
	Operator      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterGridMaintenanceRecorded is a free log retrieval operation binding the contract event 0x90b71c895d8aac16c64f910d031d8d438c34bb27c341bbdeb3f69f5b05d9d9b7.
//
// Solidity: event GridMaintenanceRecorded(address indexed gridId, uint32 recordId, string operationType, address operator)
func (_GridObjectFactory *GridObjectFactoryFilterer) FilterGridMaintenanceRecorded(opts *bind.FilterOpts, gridId []common.Address) (*GridObjectFactoryGridMaintenanceRecordedIterator, error) {

	var gridIdRule []interface{}
	for _, gridIdItem := range gridId {
		gridIdRule = append(gridIdRule, gridIdItem)
	}

	logs, sub, err := _GridObjectFactory.contract.FilterLogs(opts, "GridMaintenanceRecorded", gridIdRule)
	if err != nil {
		return nil, err
	}
	return &GridObjectFactoryGridMaintenanceRecordedIterator{contract: _GridObjectFactory.contract, event: "GridMaintenanceRecorded", logs: logs, sub: sub}, nil
}

// WatchGridMaintenanceRecorded is a free log subscription operation binding the contract event 0x90b71c895d8aac16c64f910d031d8d438c34bb27c341bbdeb3f69f5b05d9d9b7.
//
// Solidity: event GridMaintenanceRecorded(address indexed gridId, uint32 recordId, string operationType, address operator)
func (_GridObjectFactory *GridObjectFactoryFilterer) WatchGridMaintenanceRecorded(opts *bind.WatchOpts, sink chan<- *GridObjectFactoryGridMaintenanceRecorded, gridId []common.Address) (event.Subscription, error) {

	var gridIdRule []interface{}
	for _, gridIdItem := range gridId {
		gridIdRule = append(gridIdRule, gridIdItem)
	}

	logs, sub, err := _GridObjectFactory.contract.WatchLogs(opts, "GridMaintenanceRecorded", gridIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GridObjectFactoryGridMaintenanceRecorded)
				if err := _GridObjectFactory.contract.UnpackLog(event, "GridMaintenanceRecorded", log); err != nil {
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

// ParseGridMaintenanceRecorded is a log parse operation binding the contract event 0x90b71c895d8aac16c64f910d031d8d438c34bb27c341bbdeb3f69f5b05d9d9b7.
//
// Solidity: event GridMaintenanceRecorded(address indexed gridId, uint32 recordId, string operationType, address operator)
func (_GridObjectFactory *GridObjectFactoryFilterer) ParseGridMaintenanceRecorded(log types.Log) (*GridObjectFactoryGridMaintenanceRecorded, error) {
	event := new(GridObjectFactoryGridMaintenanceRecorded)
	if err := _GridObjectFactory.contract.UnpackLog(event, "GridMaintenanceRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GridObjectFactoryGridStatusUpdatedIterator is returned from FilterGridStatusUpdated and is used to iterate over the raw logs and unpacked data for GridStatusUpdated events raised by the GridObjectFactory contract.
type GridObjectFactoryGridStatusUpdatedIterator struct {
	Event *GridObjectFactoryGridStatusUpdated // Event containing the contract specifics and raw log

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
func (it *GridObjectFactoryGridStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GridObjectFactoryGridStatusUpdated)
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
		it.Event = new(GridObjectFactoryGridStatusUpdated)
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
func (it *GridObjectFactoryGridStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GridObjectFactoryGridStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GridObjectFactoryGridStatusUpdated represents a GridStatusUpdated event raised by the GridObjectFactory contract.
type GridObjectFactoryGridStatusUpdated struct {
	GridId    common.Address
	OldStatus uint8
	NewStatus uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGridStatusUpdated is a free log retrieval operation binding the contract event 0xde93df79c81c16af41ba78f153dbf7f25dd33d3e21527e15b46c51d9d5aaee0d.
//
// Solidity: event GridStatusUpdated(address indexed gridId, uint8 oldStatus, uint8 newStatus)
func (_GridObjectFactory *GridObjectFactoryFilterer) FilterGridStatusUpdated(opts *bind.FilterOpts, gridId []common.Address) (*GridObjectFactoryGridStatusUpdatedIterator, error) {

	var gridIdRule []interface{}
	for _, gridIdItem := range gridId {
		gridIdRule = append(gridIdRule, gridIdItem)
	}

	logs, sub, err := _GridObjectFactory.contract.FilterLogs(opts, "GridStatusUpdated", gridIdRule)
	if err != nil {
		return nil, err
	}
	return &GridObjectFactoryGridStatusUpdatedIterator{contract: _GridObjectFactory.contract, event: "GridStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchGridStatusUpdated is a free log subscription operation binding the contract event 0xde93df79c81c16af41ba78f153dbf7f25dd33d3e21527e15b46c51d9d5aaee0d.
//
// Solidity: event GridStatusUpdated(address indexed gridId, uint8 oldStatus, uint8 newStatus)
func (_GridObjectFactory *GridObjectFactoryFilterer) WatchGridStatusUpdated(opts *bind.WatchOpts, sink chan<- *GridObjectFactoryGridStatusUpdated, gridId []common.Address) (event.Subscription, error) {

	var gridIdRule []interface{}
	for _, gridIdItem := range gridId {
		gridIdRule = append(gridIdRule, gridIdItem)
	}

	logs, sub, err := _GridObjectFactory.contract.WatchLogs(opts, "GridStatusUpdated", gridIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GridObjectFactoryGridStatusUpdated)
				if err := _GridObjectFactory.contract.UnpackLog(event, "GridStatusUpdated", log); err != nil {
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

// ParseGridStatusUpdated is a log parse operation binding the contract event 0xde93df79c81c16af41ba78f153dbf7f25dd33d3e21527e15b46c51d9d5aaee0d.
//
// Solidity: event GridStatusUpdated(address indexed gridId, uint8 oldStatus, uint8 newStatus)
func (_GridObjectFactory *GridObjectFactoryFilterer) ParseGridStatusUpdated(log types.Log) (*GridObjectFactoryGridStatusUpdated, error) {
	event := new(GridObjectFactoryGridStatusUpdated)
	if err := _GridObjectFactory.contract.UnpackLog(event, "GridStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
