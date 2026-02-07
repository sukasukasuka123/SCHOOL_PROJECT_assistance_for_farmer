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
	GridCode             string
	FarmId               common.Address
	TerrainType          uint8
	Status               uint8
	DiseaseType          uint8
	LastStatusUpdateTime uint32
	Exists               bool
}

// GridObjectFactoryMaintenanceRecord is an auto generated low-level Go binding around an user-defined struct.
type GridObjectFactoryMaintenanceRecord struct {
	Timestamp         uint32
	OperationTypeCode uint8
	Detail            string
	EvidenceIPFSHash  [32]byte
	Operator          common.Address
}

// GridObjectFactoryMetaData contains all meta data concerning the GridObjectFactory contract.
var GridObjectFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_droneFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerAuthorized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"}],\"name\":\"CallerRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumGridObjectFactory.CitrusDiseaseType\",\"name\":\"oldType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"enumGridObjectFactory.CitrusDiseaseType\",\"name\":\"newType\",\"type\":\"uint8\"}],\"name\":\"GridDiseaseTypeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"}],\"name\":\"GridRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumGridObjectFactory.GridStatus\",\"name\":\"oldStatus\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"enumGridObjectFactory.GridStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"GridStatusUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"recordId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"operationType\",\"type\":\"uint8\"}],\"name\":\"MaintenanceRecorded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"operationTypeCode\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"detail\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"evidenceIPFSHash\",\"type\":\"bytes32\"}],\"name\":\"addMaintenanceRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"authorizeCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"authorizedCallers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"droneFactory\",\"outputs\":[{\"internalType\":\"contractDroneObjectFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"}],\"name\":\"getGrid\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"enumGridObjectFactory.GridTerrainType\",\"name\":\"terrainType\",\"type\":\"uint8\"},{\"internalType\":\"enumGridObjectFactory.GridStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"enumGridObjectFactory.CitrusDiseaseType\",\"name\":\"diseaseType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"lastStatusUpdateTime\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structGridObjectFactory.GridInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"recordId\",\"type\":\"uint32\"}],\"name\":\"getMaintenanceRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"operationTypeCode\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"detail\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"evidenceIPFSHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"internalType\":\"structGridObjectFactory.MaintenanceRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"count\",\"type\":\"uint32\"}],\"name\":\"getRecentMaintenanceRecords\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"operationTypeCode\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"detail\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"evidenceIPFSHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"internalType\":\"structGridObjectFactory.MaintenanceRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"}],\"name\":\"gridExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"grids\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"enumGridObjectFactory.GridTerrainType\",\"name\":\"terrainType\",\"type\":\"uint8\"},{\"internalType\":\"enumGridObjectFactory.GridStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"enumGridObjectFactory.CitrusDiseaseType\",\"name\":\"diseaseType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"lastStatusUpdateTime\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maintenanceRecordCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"maintenanceRecords\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"operationTypeCode\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"detail\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"evidenceIPFSHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"enumGridObjectFactory.GridTerrainType\",\"name\":\"terrainType\",\"type\":\"uint8\"}],\"name\":\"registerGrid\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_caller\",\"type\":\"address\"}],\"name\":\"revokeCaller\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"internalType\":\"enumGridObjectFactory.CitrusDiseaseType\",\"name\":\"newType\",\"type\":\"uint8\"}],\"name\":\"updateDiseaseType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"internalType\":\"enumGridObjectFactory.GridStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"updateGridStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userContract\",\"outputs\":[{\"internalType\":\"contractUser\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// AuthorizedCallers is a free data retrieval call binding the contract method 0x536fff6c.
//
// Solidity: function authorizedCallers(address ) view returns(bool)
func (_GridObjectFactory *GridObjectFactoryCaller) AuthorizedCallers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GridObjectFactory.contract.Call(opts, &out, "authorizedCallers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AuthorizedCallers is a free data retrieval call binding the contract method 0x536fff6c.
//
// Solidity: function authorizedCallers(address ) view returns(bool)
func (_GridObjectFactory *GridObjectFactorySession) AuthorizedCallers(arg0 common.Address) (bool, error) {
	return _GridObjectFactory.Contract.AuthorizedCallers(&_GridObjectFactory.CallOpts, arg0)
}

// AuthorizedCallers is a free data retrieval call binding the contract method 0x536fff6c.
//
// Solidity: function authorizedCallers(address ) view returns(bool)
func (_GridObjectFactory *GridObjectFactoryCallerSession) AuthorizedCallers(arg0 common.Address) (bool, error) {
	return _GridObjectFactory.Contract.AuthorizedCallers(&_GridObjectFactory.CallOpts, arg0)
}

// DroneFactory is a free data retrieval call binding the contract method 0xc7b26142.
//
// Solidity: function droneFactory() view returns(address)
func (_GridObjectFactory *GridObjectFactoryCaller) DroneFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GridObjectFactory.contract.Call(opts, &out, "droneFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DroneFactory is a free data retrieval call binding the contract method 0xc7b26142.
//
// Solidity: function droneFactory() view returns(address)
func (_GridObjectFactory *GridObjectFactorySession) DroneFactory() (common.Address, error) {
	return _GridObjectFactory.Contract.DroneFactory(&_GridObjectFactory.CallOpts)
}

// DroneFactory is a free data retrieval call binding the contract method 0xc7b26142.
//
// Solidity: function droneFactory() view returns(address)
func (_GridObjectFactory *GridObjectFactoryCallerSession) DroneFactory() (common.Address, error) {
	return _GridObjectFactory.Contract.DroneFactory(&_GridObjectFactory.CallOpts)
}

// GetGrid is a free data retrieval call binding the contract method 0x4e3fb04a.
//
// Solidity: function getGrid(address gridId) view returns((string,address,uint8,uint8,uint8,uint32,bool))
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
// Solidity: function getGrid(address gridId) view returns((string,address,uint8,uint8,uint8,uint32,bool))
func (_GridObjectFactory *GridObjectFactorySession) GetGrid(gridId common.Address) (GridObjectFactoryGridInfo, error) {
	return _GridObjectFactory.Contract.GetGrid(&_GridObjectFactory.CallOpts, gridId)
}

// GetGrid is a free data retrieval call binding the contract method 0x4e3fb04a.
//
// Solidity: function getGrid(address gridId) view returns((string,address,uint8,uint8,uint8,uint32,bool))
func (_GridObjectFactory *GridObjectFactoryCallerSession) GetGrid(gridId common.Address) (GridObjectFactoryGridInfo, error) {
	return _GridObjectFactory.Contract.GetGrid(&_GridObjectFactory.CallOpts, gridId)
}

// GetMaintenanceRecord is a free data retrieval call binding the contract method 0x18ff6b88.
//
// Solidity: function getMaintenanceRecord(address gridId, uint32 recordId) view returns((uint32,uint8,string,bytes32,address))
func (_GridObjectFactory *GridObjectFactoryCaller) GetMaintenanceRecord(opts *bind.CallOpts, gridId common.Address, recordId uint32) (GridObjectFactoryMaintenanceRecord, error) {
	var out []interface{}
	err := _GridObjectFactory.contract.Call(opts, &out, "getMaintenanceRecord", gridId, recordId)

	if err != nil {
		return *new(GridObjectFactoryMaintenanceRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(GridObjectFactoryMaintenanceRecord)).(*GridObjectFactoryMaintenanceRecord)

	return out0, err

}

// GetMaintenanceRecord is a free data retrieval call binding the contract method 0x18ff6b88.
//
// Solidity: function getMaintenanceRecord(address gridId, uint32 recordId) view returns((uint32,uint8,string,bytes32,address))
func (_GridObjectFactory *GridObjectFactorySession) GetMaintenanceRecord(gridId common.Address, recordId uint32) (GridObjectFactoryMaintenanceRecord, error) {
	return _GridObjectFactory.Contract.GetMaintenanceRecord(&_GridObjectFactory.CallOpts, gridId, recordId)
}

// GetMaintenanceRecord is a free data retrieval call binding the contract method 0x18ff6b88.
//
// Solidity: function getMaintenanceRecord(address gridId, uint32 recordId) view returns((uint32,uint8,string,bytes32,address))
func (_GridObjectFactory *GridObjectFactoryCallerSession) GetMaintenanceRecord(gridId common.Address, recordId uint32) (GridObjectFactoryMaintenanceRecord, error) {
	return _GridObjectFactory.Contract.GetMaintenanceRecord(&_GridObjectFactory.CallOpts, gridId, recordId)
}

// GetRecentMaintenanceRecords is a free data retrieval call binding the contract method 0x13fcfc2b.
//
// Solidity: function getRecentMaintenanceRecords(address gridId, uint32 count) view returns((uint32,uint8,string,bytes32,address)[])
func (_GridObjectFactory *GridObjectFactoryCaller) GetRecentMaintenanceRecords(opts *bind.CallOpts, gridId common.Address, count uint32) ([]GridObjectFactoryMaintenanceRecord, error) {
	var out []interface{}
	err := _GridObjectFactory.contract.Call(opts, &out, "getRecentMaintenanceRecords", gridId, count)

	if err != nil {
		return *new([]GridObjectFactoryMaintenanceRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]GridObjectFactoryMaintenanceRecord)).(*[]GridObjectFactoryMaintenanceRecord)

	return out0, err

}

// GetRecentMaintenanceRecords is a free data retrieval call binding the contract method 0x13fcfc2b.
//
// Solidity: function getRecentMaintenanceRecords(address gridId, uint32 count) view returns((uint32,uint8,string,bytes32,address)[])
func (_GridObjectFactory *GridObjectFactorySession) GetRecentMaintenanceRecords(gridId common.Address, count uint32) ([]GridObjectFactoryMaintenanceRecord, error) {
	return _GridObjectFactory.Contract.GetRecentMaintenanceRecords(&_GridObjectFactory.CallOpts, gridId, count)
}

// GetRecentMaintenanceRecords is a free data retrieval call binding the contract method 0x13fcfc2b.
//
// Solidity: function getRecentMaintenanceRecords(address gridId, uint32 count) view returns((uint32,uint8,string,bytes32,address)[])
func (_GridObjectFactory *GridObjectFactoryCallerSession) GetRecentMaintenanceRecords(gridId common.Address, count uint32) ([]GridObjectFactoryMaintenanceRecord, error) {
	return _GridObjectFactory.Contract.GetRecentMaintenanceRecords(&_GridObjectFactory.CallOpts, gridId, count)
}

// GridExists is a free data retrieval call binding the contract method 0xcbf62e42.
//
// Solidity: function gridExists(address gridId) view returns(bool)
func (_GridObjectFactory *GridObjectFactoryCaller) GridExists(opts *bind.CallOpts, gridId common.Address) (bool, error) {
	var out []interface{}
	err := _GridObjectFactory.contract.Call(opts, &out, "gridExists", gridId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GridExists is a free data retrieval call binding the contract method 0xcbf62e42.
//
// Solidity: function gridExists(address gridId) view returns(bool)
func (_GridObjectFactory *GridObjectFactorySession) GridExists(gridId common.Address) (bool, error) {
	return _GridObjectFactory.Contract.GridExists(&_GridObjectFactory.CallOpts, gridId)
}

// GridExists is a free data retrieval call binding the contract method 0xcbf62e42.
//
// Solidity: function gridExists(address gridId) view returns(bool)
func (_GridObjectFactory *GridObjectFactoryCallerSession) GridExists(gridId common.Address) (bool, error) {
	return _GridObjectFactory.Contract.GridExists(&_GridObjectFactory.CallOpts, gridId)
}

// Grids is a free data retrieval call binding the contract method 0xadbb79cd.
//
// Solidity: function grids(address ) view returns(string gridCode, address farmId, uint8 terrainType, uint8 status, uint8 diseaseType, uint32 lastStatusUpdateTime, bool exists)
func (_GridObjectFactory *GridObjectFactoryCaller) Grids(opts *bind.CallOpts, arg0 common.Address) (struct {
	GridCode             string
	FarmId               common.Address
	TerrainType          uint8
	Status               uint8
	DiseaseType          uint8
	LastStatusUpdateTime uint32
	Exists               bool
}, error) {
	var out []interface{}
	err := _GridObjectFactory.contract.Call(opts, &out, "grids", arg0)

	outstruct := new(struct {
		GridCode             string
		FarmId               common.Address
		TerrainType          uint8
		Status               uint8
		DiseaseType          uint8
		LastStatusUpdateTime uint32
		Exists               bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.GridCode = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.FarmId = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TerrainType = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.Status = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.DiseaseType = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.LastStatusUpdateTime = *abi.ConvertType(out[5], new(uint32)).(*uint32)
	outstruct.Exists = *abi.ConvertType(out[6], new(bool)).(*bool)

	return *outstruct, err

}

// Grids is a free data retrieval call binding the contract method 0xadbb79cd.
//
// Solidity: function grids(address ) view returns(string gridCode, address farmId, uint8 terrainType, uint8 status, uint8 diseaseType, uint32 lastStatusUpdateTime, bool exists)
func (_GridObjectFactory *GridObjectFactorySession) Grids(arg0 common.Address) (struct {
	GridCode             string
	FarmId               common.Address
	TerrainType          uint8
	Status               uint8
	DiseaseType          uint8
	LastStatusUpdateTime uint32
	Exists               bool
}, error) {
	return _GridObjectFactory.Contract.Grids(&_GridObjectFactory.CallOpts, arg0)
}

// Grids is a free data retrieval call binding the contract method 0xadbb79cd.
//
// Solidity: function grids(address ) view returns(string gridCode, address farmId, uint8 terrainType, uint8 status, uint8 diseaseType, uint32 lastStatusUpdateTime, bool exists)
func (_GridObjectFactory *GridObjectFactoryCallerSession) Grids(arg0 common.Address) (struct {
	GridCode             string
	FarmId               common.Address
	TerrainType          uint8
	Status               uint8
	DiseaseType          uint8
	LastStatusUpdateTime uint32
	Exists               bool
}, error) {
	return _GridObjectFactory.Contract.Grids(&_GridObjectFactory.CallOpts, arg0)
}

// MaintenanceRecordCount is a free data retrieval call binding the contract method 0xbb318b4f.
//
// Solidity: function maintenanceRecordCount(address ) view returns(uint32)
func (_GridObjectFactory *GridObjectFactoryCaller) MaintenanceRecordCount(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _GridObjectFactory.contract.Call(opts, &out, "maintenanceRecordCount", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaintenanceRecordCount is a free data retrieval call binding the contract method 0xbb318b4f.
//
// Solidity: function maintenanceRecordCount(address ) view returns(uint32)
func (_GridObjectFactory *GridObjectFactorySession) MaintenanceRecordCount(arg0 common.Address) (uint32, error) {
	return _GridObjectFactory.Contract.MaintenanceRecordCount(&_GridObjectFactory.CallOpts, arg0)
}

// MaintenanceRecordCount is a free data retrieval call binding the contract method 0xbb318b4f.
//
// Solidity: function maintenanceRecordCount(address ) view returns(uint32)
func (_GridObjectFactory *GridObjectFactoryCallerSession) MaintenanceRecordCount(arg0 common.Address) (uint32, error) {
	return _GridObjectFactory.Contract.MaintenanceRecordCount(&_GridObjectFactory.CallOpts, arg0)
}

// MaintenanceRecords is a free data retrieval call binding the contract method 0xe93725f9.
//
// Solidity: function maintenanceRecords(address , uint32 ) view returns(uint32 timestamp, uint8 operationTypeCode, string detail, bytes32 evidenceIPFSHash, address operator)
func (_GridObjectFactory *GridObjectFactoryCaller) MaintenanceRecords(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	Timestamp         uint32
	OperationTypeCode uint8
	Detail            string
	EvidenceIPFSHash  [32]byte
	Operator          common.Address
}, error) {
	var out []interface{}
	err := _GridObjectFactory.contract.Call(opts, &out, "maintenanceRecords", arg0, arg1)

	outstruct := new(struct {
		Timestamp         uint32
		OperationTypeCode uint8
		Detail            string
		EvidenceIPFSHash  [32]byte
		Operator          common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.OperationTypeCode = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.Detail = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.EvidenceIPFSHash = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.Operator = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// MaintenanceRecords is a free data retrieval call binding the contract method 0xe93725f9.
//
// Solidity: function maintenanceRecords(address , uint32 ) view returns(uint32 timestamp, uint8 operationTypeCode, string detail, bytes32 evidenceIPFSHash, address operator)
func (_GridObjectFactory *GridObjectFactorySession) MaintenanceRecords(arg0 common.Address, arg1 uint32) (struct {
	Timestamp         uint32
	OperationTypeCode uint8
	Detail            string
	EvidenceIPFSHash  [32]byte
	Operator          common.Address
}, error) {
	return _GridObjectFactory.Contract.MaintenanceRecords(&_GridObjectFactory.CallOpts, arg0, arg1)
}

// MaintenanceRecords is a free data retrieval call binding the contract method 0xe93725f9.
//
// Solidity: function maintenanceRecords(address , uint32 ) view returns(uint32 timestamp, uint8 operationTypeCode, string detail, bytes32 evidenceIPFSHash, address operator)
func (_GridObjectFactory *GridObjectFactoryCallerSession) MaintenanceRecords(arg0 common.Address, arg1 uint32) (struct {
	Timestamp         uint32
	OperationTypeCode uint8
	Detail            string
	EvidenceIPFSHash  [32]byte
	Operator          common.Address
}, error) {
	return _GridObjectFactory.Contract.MaintenanceRecords(&_GridObjectFactory.CallOpts, arg0, arg1)
}

// UserContract is a free data retrieval call binding the contract method 0x5000596b.
//
// Solidity: function userContract() view returns(address)
func (_GridObjectFactory *GridObjectFactoryCaller) UserContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GridObjectFactory.contract.Call(opts, &out, "userContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserContract is a free data retrieval call binding the contract method 0x5000596b.
//
// Solidity: function userContract() view returns(address)
func (_GridObjectFactory *GridObjectFactorySession) UserContract() (common.Address, error) {
	return _GridObjectFactory.Contract.UserContract(&_GridObjectFactory.CallOpts)
}

// UserContract is a free data retrieval call binding the contract method 0x5000596b.
//
// Solidity: function userContract() view returns(address)
func (_GridObjectFactory *GridObjectFactoryCallerSession) UserContract() (common.Address, error) {
	return _GridObjectFactory.Contract.UserContract(&_GridObjectFactory.CallOpts)
}

// AddMaintenanceRecord is a paid mutator transaction binding the contract method 0x8c4f7db5.
//
// Solidity: function addMaintenanceRecord(address gridId, uint8 operationTypeCode, string detail, bytes32 evidenceIPFSHash) returns()
func (_GridObjectFactory *GridObjectFactoryTransactor) AddMaintenanceRecord(opts *bind.TransactOpts, gridId common.Address, operationTypeCode uint8, detail string, evidenceIPFSHash [32]byte) (*types.Transaction, error) {
	return _GridObjectFactory.contract.Transact(opts, "addMaintenanceRecord", gridId, operationTypeCode, detail, evidenceIPFSHash)
}

// AddMaintenanceRecord is a paid mutator transaction binding the contract method 0x8c4f7db5.
//
// Solidity: function addMaintenanceRecord(address gridId, uint8 operationTypeCode, string detail, bytes32 evidenceIPFSHash) returns()
func (_GridObjectFactory *GridObjectFactorySession) AddMaintenanceRecord(gridId common.Address, operationTypeCode uint8, detail string, evidenceIPFSHash [32]byte) (*types.Transaction, error) {
	return _GridObjectFactory.Contract.AddMaintenanceRecord(&_GridObjectFactory.TransactOpts, gridId, operationTypeCode, detail, evidenceIPFSHash)
}

// AddMaintenanceRecord is a paid mutator transaction binding the contract method 0x8c4f7db5.
//
// Solidity: function addMaintenanceRecord(address gridId, uint8 operationTypeCode, string detail, bytes32 evidenceIPFSHash) returns()
func (_GridObjectFactory *GridObjectFactoryTransactorSession) AddMaintenanceRecord(gridId common.Address, operationTypeCode uint8, detail string, evidenceIPFSHash [32]byte) (*types.Transaction, error) {
	return _GridObjectFactory.Contract.AddMaintenanceRecord(&_GridObjectFactory.TransactOpts, gridId, operationTypeCode, detail, evidenceIPFSHash)
}

// AuthorizeCaller is a paid mutator transaction binding the contract method 0x2c388d5d.
//
// Solidity: function authorizeCaller(address _caller) returns()
func (_GridObjectFactory *GridObjectFactoryTransactor) AuthorizeCaller(opts *bind.TransactOpts, _caller common.Address) (*types.Transaction, error) {
	return _GridObjectFactory.contract.Transact(opts, "authorizeCaller", _caller)
}

// AuthorizeCaller is a paid mutator transaction binding the contract method 0x2c388d5d.
//
// Solidity: function authorizeCaller(address _caller) returns()
func (_GridObjectFactory *GridObjectFactorySession) AuthorizeCaller(_caller common.Address) (*types.Transaction, error) {
	return _GridObjectFactory.Contract.AuthorizeCaller(&_GridObjectFactory.TransactOpts, _caller)
}

// AuthorizeCaller is a paid mutator transaction binding the contract method 0x2c388d5d.
//
// Solidity: function authorizeCaller(address _caller) returns()
func (_GridObjectFactory *GridObjectFactoryTransactorSession) AuthorizeCaller(_caller common.Address) (*types.Transaction, error) {
	return _GridObjectFactory.Contract.AuthorizeCaller(&_GridObjectFactory.TransactOpts, _caller)
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

// RevokeCaller is a paid mutator transaction binding the contract method 0xb7743530.
//
// Solidity: function revokeCaller(address _caller) returns()
func (_GridObjectFactory *GridObjectFactoryTransactor) RevokeCaller(opts *bind.TransactOpts, _caller common.Address) (*types.Transaction, error) {
	return _GridObjectFactory.contract.Transact(opts, "revokeCaller", _caller)
}

// RevokeCaller is a paid mutator transaction binding the contract method 0xb7743530.
//
// Solidity: function revokeCaller(address _caller) returns()
func (_GridObjectFactory *GridObjectFactorySession) RevokeCaller(_caller common.Address) (*types.Transaction, error) {
	return _GridObjectFactory.Contract.RevokeCaller(&_GridObjectFactory.TransactOpts, _caller)
}

// RevokeCaller is a paid mutator transaction binding the contract method 0xb7743530.
//
// Solidity: function revokeCaller(address _caller) returns()
func (_GridObjectFactory *GridObjectFactoryTransactorSession) RevokeCaller(_caller common.Address) (*types.Transaction, error) {
	return _GridObjectFactory.Contract.RevokeCaller(&_GridObjectFactory.TransactOpts, _caller)
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

// GridObjectFactoryCallerAuthorizedIterator is returned from FilterCallerAuthorized and is used to iterate over the raw logs and unpacked data for CallerAuthorized events raised by the GridObjectFactory contract.
type GridObjectFactoryCallerAuthorizedIterator struct {
	Event *GridObjectFactoryCallerAuthorized // Event containing the contract specifics and raw log

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
func (it *GridObjectFactoryCallerAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GridObjectFactoryCallerAuthorized)
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
		it.Event = new(GridObjectFactoryCallerAuthorized)
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
func (it *GridObjectFactoryCallerAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GridObjectFactoryCallerAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GridObjectFactoryCallerAuthorized represents a CallerAuthorized event raised by the GridObjectFactory contract.
type GridObjectFactoryCallerAuthorized struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallerAuthorized is a free log retrieval operation binding the contract event 0x6acfd92212f0ec670af78f8ba2273e4a85ac17023475650c25be0440602bff2d.
//
// Solidity: event CallerAuthorized(address indexed caller)
func (_GridObjectFactory *GridObjectFactoryFilterer) FilterCallerAuthorized(opts *bind.FilterOpts, caller []common.Address) (*GridObjectFactoryCallerAuthorizedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _GridObjectFactory.contract.FilterLogs(opts, "CallerAuthorized", callerRule)
	if err != nil {
		return nil, err
	}
	return &GridObjectFactoryCallerAuthorizedIterator{contract: _GridObjectFactory.contract, event: "CallerAuthorized", logs: logs, sub: sub}, nil
}

// WatchCallerAuthorized is a free log subscription operation binding the contract event 0x6acfd92212f0ec670af78f8ba2273e4a85ac17023475650c25be0440602bff2d.
//
// Solidity: event CallerAuthorized(address indexed caller)
func (_GridObjectFactory *GridObjectFactoryFilterer) WatchCallerAuthorized(opts *bind.WatchOpts, sink chan<- *GridObjectFactoryCallerAuthorized, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _GridObjectFactory.contract.WatchLogs(opts, "CallerAuthorized", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GridObjectFactoryCallerAuthorized)
				if err := _GridObjectFactory.contract.UnpackLog(event, "CallerAuthorized", log); err != nil {
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

// ParseCallerAuthorized is a log parse operation binding the contract event 0x6acfd92212f0ec670af78f8ba2273e4a85ac17023475650c25be0440602bff2d.
//
// Solidity: event CallerAuthorized(address indexed caller)
func (_GridObjectFactory *GridObjectFactoryFilterer) ParseCallerAuthorized(log types.Log) (*GridObjectFactoryCallerAuthorized, error) {
	event := new(GridObjectFactoryCallerAuthorized)
	if err := _GridObjectFactory.contract.UnpackLog(event, "CallerAuthorized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GridObjectFactoryCallerRevokedIterator is returned from FilterCallerRevoked and is used to iterate over the raw logs and unpacked data for CallerRevoked events raised by the GridObjectFactory contract.
type GridObjectFactoryCallerRevokedIterator struct {
	Event *GridObjectFactoryCallerRevoked // Event containing the contract specifics and raw log

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
func (it *GridObjectFactoryCallerRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GridObjectFactoryCallerRevoked)
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
		it.Event = new(GridObjectFactoryCallerRevoked)
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
func (it *GridObjectFactoryCallerRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GridObjectFactoryCallerRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GridObjectFactoryCallerRevoked represents a CallerRevoked event raised by the GridObjectFactory contract.
type GridObjectFactoryCallerRevoked struct {
	Caller common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterCallerRevoked is a free log retrieval operation binding the contract event 0x0491b0192bae7692618bfa4eff3f4942d2d8ec3300ef2e63d325b45e937c4ff1.
//
// Solidity: event CallerRevoked(address indexed caller)
func (_GridObjectFactory *GridObjectFactoryFilterer) FilterCallerRevoked(opts *bind.FilterOpts, caller []common.Address) (*GridObjectFactoryCallerRevokedIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _GridObjectFactory.contract.FilterLogs(opts, "CallerRevoked", callerRule)
	if err != nil {
		return nil, err
	}
	return &GridObjectFactoryCallerRevokedIterator{contract: _GridObjectFactory.contract, event: "CallerRevoked", logs: logs, sub: sub}, nil
}

// WatchCallerRevoked is a free log subscription operation binding the contract event 0x0491b0192bae7692618bfa4eff3f4942d2d8ec3300ef2e63d325b45e937c4ff1.
//
// Solidity: event CallerRevoked(address indexed caller)
func (_GridObjectFactory *GridObjectFactoryFilterer) WatchCallerRevoked(opts *bind.WatchOpts, sink chan<- *GridObjectFactoryCallerRevoked, caller []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}

	logs, sub, err := _GridObjectFactory.contract.WatchLogs(opts, "CallerRevoked", callerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GridObjectFactoryCallerRevoked)
				if err := _GridObjectFactory.contract.UnpackLog(event, "CallerRevoked", log); err != nil {
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

// ParseCallerRevoked is a log parse operation binding the contract event 0x0491b0192bae7692618bfa4eff3f4942d2d8ec3300ef2e63d325b45e937c4ff1.
//
// Solidity: event CallerRevoked(address indexed caller)
func (_GridObjectFactory *GridObjectFactoryFilterer) ParseCallerRevoked(log types.Log) (*GridObjectFactoryCallerRevoked, error) {
	event := new(GridObjectFactoryCallerRevoked)
	if err := _GridObjectFactory.contract.UnpackLog(event, "CallerRevoked", log); err != nil {
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

// GridObjectFactoryGridRegisteredIterator is returned from FilterGridRegistered and is used to iterate over the raw logs and unpacked data for GridRegistered events raised by the GridObjectFactory contract.
type GridObjectFactoryGridRegisteredIterator struct {
	Event *GridObjectFactoryGridRegistered // Event containing the contract specifics and raw log

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
func (it *GridObjectFactoryGridRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GridObjectFactoryGridRegistered)
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
		it.Event = new(GridObjectFactoryGridRegistered)
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
func (it *GridObjectFactoryGridRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GridObjectFactoryGridRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GridObjectFactoryGridRegistered represents a GridRegistered event raised by the GridObjectFactory contract.
type GridObjectFactoryGridRegistered struct {
	FarmId   common.Address
	GridId   common.Address
	GridCode string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGridRegistered is a free log retrieval operation binding the contract event 0x9023782af8a470114d8ec3d466f6315216c9ace02192cca1160639bfbd260675.
//
// Solidity: event GridRegistered(address indexed farmId, address indexed gridId, string gridCode)
func (_GridObjectFactory *GridObjectFactoryFilterer) FilterGridRegistered(opts *bind.FilterOpts, farmId []common.Address, gridId []common.Address) (*GridObjectFactoryGridRegisteredIterator, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}
	var gridIdRule []interface{}
	for _, gridIdItem := range gridId {
		gridIdRule = append(gridIdRule, gridIdItem)
	}

	logs, sub, err := _GridObjectFactory.contract.FilterLogs(opts, "GridRegistered", farmIdRule, gridIdRule)
	if err != nil {
		return nil, err
	}
	return &GridObjectFactoryGridRegisteredIterator{contract: _GridObjectFactory.contract, event: "GridRegistered", logs: logs, sub: sub}, nil
}

// WatchGridRegistered is a free log subscription operation binding the contract event 0x9023782af8a470114d8ec3d466f6315216c9ace02192cca1160639bfbd260675.
//
// Solidity: event GridRegistered(address indexed farmId, address indexed gridId, string gridCode)
func (_GridObjectFactory *GridObjectFactoryFilterer) WatchGridRegistered(opts *bind.WatchOpts, sink chan<- *GridObjectFactoryGridRegistered, farmId []common.Address, gridId []common.Address) (event.Subscription, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}
	var gridIdRule []interface{}
	for _, gridIdItem := range gridId {
		gridIdRule = append(gridIdRule, gridIdItem)
	}

	logs, sub, err := _GridObjectFactory.contract.WatchLogs(opts, "GridRegistered", farmIdRule, gridIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GridObjectFactoryGridRegistered)
				if err := _GridObjectFactory.contract.UnpackLog(event, "GridRegistered", log); err != nil {
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

// ParseGridRegistered is a log parse operation binding the contract event 0x9023782af8a470114d8ec3d466f6315216c9ace02192cca1160639bfbd260675.
//
// Solidity: event GridRegistered(address indexed farmId, address indexed gridId, string gridCode)
func (_GridObjectFactory *GridObjectFactoryFilterer) ParseGridRegistered(log types.Log) (*GridObjectFactoryGridRegistered, error) {
	event := new(GridObjectFactoryGridRegistered)
	if err := _GridObjectFactory.contract.UnpackLog(event, "GridRegistered", log); err != nil {
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
// Solidity: event GridStatusUpdated(address indexed gridId, uint8 indexed oldStatus, uint8 indexed newStatus)
func (_GridObjectFactory *GridObjectFactoryFilterer) FilterGridStatusUpdated(opts *bind.FilterOpts, gridId []common.Address, oldStatus []uint8, newStatus []uint8) (*GridObjectFactoryGridStatusUpdatedIterator, error) {

	var gridIdRule []interface{}
	for _, gridIdItem := range gridId {
		gridIdRule = append(gridIdRule, gridIdItem)
	}
	var oldStatusRule []interface{}
	for _, oldStatusItem := range oldStatus {
		oldStatusRule = append(oldStatusRule, oldStatusItem)
	}
	var newStatusRule []interface{}
	for _, newStatusItem := range newStatus {
		newStatusRule = append(newStatusRule, newStatusItem)
	}

	logs, sub, err := _GridObjectFactory.contract.FilterLogs(opts, "GridStatusUpdated", gridIdRule, oldStatusRule, newStatusRule)
	if err != nil {
		return nil, err
	}
	return &GridObjectFactoryGridStatusUpdatedIterator{contract: _GridObjectFactory.contract, event: "GridStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchGridStatusUpdated is a free log subscription operation binding the contract event 0xde93df79c81c16af41ba78f153dbf7f25dd33d3e21527e15b46c51d9d5aaee0d.
//
// Solidity: event GridStatusUpdated(address indexed gridId, uint8 indexed oldStatus, uint8 indexed newStatus)
func (_GridObjectFactory *GridObjectFactoryFilterer) WatchGridStatusUpdated(opts *bind.WatchOpts, sink chan<- *GridObjectFactoryGridStatusUpdated, gridId []common.Address, oldStatus []uint8, newStatus []uint8) (event.Subscription, error) {

	var gridIdRule []interface{}
	for _, gridIdItem := range gridId {
		gridIdRule = append(gridIdRule, gridIdItem)
	}
	var oldStatusRule []interface{}
	for _, oldStatusItem := range oldStatus {
		oldStatusRule = append(oldStatusRule, oldStatusItem)
	}
	var newStatusRule []interface{}
	for _, newStatusItem := range newStatus {
		newStatusRule = append(newStatusRule, newStatusItem)
	}

	logs, sub, err := _GridObjectFactory.contract.WatchLogs(opts, "GridStatusUpdated", gridIdRule, oldStatusRule, newStatusRule)
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
// Solidity: event GridStatusUpdated(address indexed gridId, uint8 indexed oldStatus, uint8 indexed newStatus)
func (_GridObjectFactory *GridObjectFactoryFilterer) ParseGridStatusUpdated(log types.Log) (*GridObjectFactoryGridStatusUpdated, error) {
	event := new(GridObjectFactoryGridStatusUpdated)
	if err := _GridObjectFactory.contract.UnpackLog(event, "GridStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GridObjectFactoryMaintenanceRecordedIterator is returned from FilterMaintenanceRecorded and is used to iterate over the raw logs and unpacked data for MaintenanceRecorded events raised by the GridObjectFactory contract.
type GridObjectFactoryMaintenanceRecordedIterator struct {
	Event *GridObjectFactoryMaintenanceRecorded // Event containing the contract specifics and raw log

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
func (it *GridObjectFactoryMaintenanceRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GridObjectFactoryMaintenanceRecorded)
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
		it.Event = new(GridObjectFactoryMaintenanceRecorded)
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
func (it *GridObjectFactoryMaintenanceRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GridObjectFactoryMaintenanceRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GridObjectFactoryMaintenanceRecorded represents a MaintenanceRecorded event raised by the GridObjectFactory contract.
type GridObjectFactoryMaintenanceRecorded struct {
	GridId        common.Address
	RecordId      uint32
	OperationType uint8
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterMaintenanceRecorded is a free log retrieval operation binding the contract event 0x67717799ec1ff866286ec717a03156f2bb4418e3dfbf93c290498a23b846897b.
//
// Solidity: event MaintenanceRecorded(address indexed gridId, uint32 indexed recordId, uint8 operationType)
func (_GridObjectFactory *GridObjectFactoryFilterer) FilterMaintenanceRecorded(opts *bind.FilterOpts, gridId []common.Address, recordId []uint32) (*GridObjectFactoryMaintenanceRecordedIterator, error) {

	var gridIdRule []interface{}
	for _, gridIdItem := range gridId {
		gridIdRule = append(gridIdRule, gridIdItem)
	}
	var recordIdRule []interface{}
	for _, recordIdItem := range recordId {
		recordIdRule = append(recordIdRule, recordIdItem)
	}

	logs, sub, err := _GridObjectFactory.contract.FilterLogs(opts, "MaintenanceRecorded", gridIdRule, recordIdRule)
	if err != nil {
		return nil, err
	}
	return &GridObjectFactoryMaintenanceRecordedIterator{contract: _GridObjectFactory.contract, event: "MaintenanceRecorded", logs: logs, sub: sub}, nil
}

// WatchMaintenanceRecorded is a free log subscription operation binding the contract event 0x67717799ec1ff866286ec717a03156f2bb4418e3dfbf93c290498a23b846897b.
//
// Solidity: event MaintenanceRecorded(address indexed gridId, uint32 indexed recordId, uint8 operationType)
func (_GridObjectFactory *GridObjectFactoryFilterer) WatchMaintenanceRecorded(opts *bind.WatchOpts, sink chan<- *GridObjectFactoryMaintenanceRecorded, gridId []common.Address, recordId []uint32) (event.Subscription, error) {

	var gridIdRule []interface{}
	for _, gridIdItem := range gridId {
		gridIdRule = append(gridIdRule, gridIdItem)
	}
	var recordIdRule []interface{}
	for _, recordIdItem := range recordId {
		recordIdRule = append(recordIdRule, recordIdItem)
	}

	logs, sub, err := _GridObjectFactory.contract.WatchLogs(opts, "MaintenanceRecorded", gridIdRule, recordIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GridObjectFactoryMaintenanceRecorded)
				if err := _GridObjectFactory.contract.UnpackLog(event, "MaintenanceRecorded", log); err != nil {
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

// ParseMaintenanceRecorded is a log parse operation binding the contract event 0x67717799ec1ff866286ec717a03156f2bb4418e3dfbf93c290498a23b846897b.
//
// Solidity: event MaintenanceRecorded(address indexed gridId, uint32 indexed recordId, uint8 operationType)
func (_GridObjectFactory *GridObjectFactoryFilterer) ParseMaintenanceRecorded(log types.Log) (*GridObjectFactoryMaintenanceRecorded, error) {
	event := new(GridObjectFactoryMaintenanceRecorded)
	if err := _GridObjectFactory.contract.UnpackLog(event, "MaintenanceRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
