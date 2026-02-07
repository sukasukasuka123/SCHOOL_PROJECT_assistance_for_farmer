// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package DroneObjectFactory

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

// DroneObjectFactoryDrone is an auto generated low-level Go binding around an user-defined struct.
type DroneObjectFactoryDrone struct {
	DroneType       uint8
	FirmwareVersion string
	State           uint8
	RegisteredTime  uint32
	Exists          bool
}

// DroneObjectFactoryMaintenanceLog is an auto generated low-level Go binding around an user-defined struct.
type DroneObjectFactoryMaintenanceLog struct {
	Timestamp           uint32
	MaintenanceTypeCode uint8
	SensorCalibrated    bool
	FirmwareVersion     string
	Technician          common.Address
}

// DroneObjectFactoryMetaData contains all meta data concerning the DroneObjectFactory contract.
var DroneObjectFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"droneId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDroneObjectFactory.DroneType\",\"name\":\"droneType\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"name\":\"DroneRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"droneId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumDroneObjectFactory.DroneState\",\"name\":\"newState\",\"type\":\"uint8\"}],\"name\":\"DroneStateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"droneId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newVersion\",\"type\":\"string\"}],\"name\":\"FirmwareUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"droneId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"logId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"maintenanceType\",\"type\":\"uint8\"}],\"name\":\"MaintenanceLogAdded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_maintenanceTypeCode\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"_sensorCalibrated\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_firmwareVersion\",\"type\":\"string\"}],\"name\":\"addMaintenanceLog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"}],\"name\":\"droneExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"drones\",\"outputs\":[{\"internalType\":\"enumDroneObjectFactory.DroneType\",\"name\":\"droneType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"firmwareVersion\",\"type\":\"string\"},{\"internalType\":\"enumDroneObjectFactory.DroneState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"registeredTime\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"}],\"name\":\"getDrone\",\"outputs\":[{\"components\":[{\"internalType\":\"enumDroneObjectFactory.DroneType\",\"name\":\"droneType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"firmwareVersion\",\"type\":\"string\"},{\"internalType\":\"enumDroneObjectFactory.DroneState\",\"name\":\"state\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"registeredTime\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structDroneObjectFactory.Drone\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"}],\"name\":\"getDroneType\",\"outputs\":[{\"internalType\":\"enumDroneObjectFactory.DroneType\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_logId\",\"type\":\"uint32\"}],\"name\":\"getMaintenanceLog\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maintenanceTypeCode\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"sensorCalibrated\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"firmwareVersion\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"technician\",\"type\":\"address\"}],\"internalType\":\"structDroneObjectFactory.MaintenanceLog\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_count\",\"type\":\"uint32\"}],\"name\":\"getRecentMaintenanceLogs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maintenanceTypeCode\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"sensorCalibrated\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"firmwareVersion\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"technician\",\"type\":\"address\"}],\"internalType\":\"structDroneObjectFactory.MaintenanceLog[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"maintenanceLogCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"maintenanceLogs\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"maintenanceTypeCode\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"sensorCalibrated\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"firmwareVersion\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"technician\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"},{\"internalType\":\"enumDroneObjectFactory.DroneType\",\"name\":\"_droneType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_firmwareVersion\",\"type\":\"string\"},{\"internalType\":\"enumDroneObjectFactory.DroneState\",\"name\":\"_droneState\",\"type\":\"uint8\"}],\"name\":\"registerDrone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"},{\"internalType\":\"enumDroneObjectFactory.DroneState\",\"name\":\"_newState\",\"type\":\"uint8\"}],\"name\":\"updateDroneState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_newFirmwareVersion\",\"type\":\"string\"}],\"name\":\"updateFirmwareVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userContract\",\"outputs\":[{\"internalType\":\"contractUser\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// DroneObjectFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use DroneObjectFactoryMetaData.ABI instead.
var DroneObjectFactoryABI = DroneObjectFactoryMetaData.ABI

// DroneObjectFactory is an auto generated Go binding around an Ethereum contract.
type DroneObjectFactory struct {
	DroneObjectFactoryCaller     // Read-only binding to the contract
	DroneObjectFactoryTransactor // Write-only binding to the contract
	DroneObjectFactoryFilterer   // Log filterer for contract events
}

// DroneObjectFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type DroneObjectFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DroneObjectFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DroneObjectFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DroneObjectFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DroneObjectFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DroneObjectFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DroneObjectFactorySession struct {
	Contract     *DroneObjectFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// DroneObjectFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DroneObjectFactoryCallerSession struct {
	Contract *DroneObjectFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// DroneObjectFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DroneObjectFactoryTransactorSession struct {
	Contract     *DroneObjectFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// DroneObjectFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type DroneObjectFactoryRaw struct {
	Contract *DroneObjectFactory // Generic contract binding to access the raw methods on
}

// DroneObjectFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DroneObjectFactoryCallerRaw struct {
	Contract *DroneObjectFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// DroneObjectFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DroneObjectFactoryTransactorRaw struct {
	Contract *DroneObjectFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDroneObjectFactory creates a new instance of DroneObjectFactory, bound to a specific deployed contract.
func NewDroneObjectFactory(address common.Address, backend bind.ContractBackend) (*DroneObjectFactory, error) {
	contract, err := bindDroneObjectFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DroneObjectFactory{DroneObjectFactoryCaller: DroneObjectFactoryCaller{contract: contract}, DroneObjectFactoryTransactor: DroneObjectFactoryTransactor{contract: contract}, DroneObjectFactoryFilterer: DroneObjectFactoryFilterer{contract: contract}}, nil
}

// NewDroneObjectFactoryCaller creates a new read-only instance of DroneObjectFactory, bound to a specific deployed contract.
func NewDroneObjectFactoryCaller(address common.Address, caller bind.ContractCaller) (*DroneObjectFactoryCaller, error) {
	contract, err := bindDroneObjectFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DroneObjectFactoryCaller{contract: contract}, nil
}

// NewDroneObjectFactoryTransactor creates a new write-only instance of DroneObjectFactory, bound to a specific deployed contract.
func NewDroneObjectFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*DroneObjectFactoryTransactor, error) {
	contract, err := bindDroneObjectFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DroneObjectFactoryTransactor{contract: contract}, nil
}

// NewDroneObjectFactoryFilterer creates a new log filterer instance of DroneObjectFactory, bound to a specific deployed contract.
func NewDroneObjectFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*DroneObjectFactoryFilterer, error) {
	contract, err := bindDroneObjectFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DroneObjectFactoryFilterer{contract: contract}, nil
}

// bindDroneObjectFactory binds a generic wrapper to an already deployed contract.
func bindDroneObjectFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := DroneObjectFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DroneObjectFactory *DroneObjectFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DroneObjectFactory.Contract.DroneObjectFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DroneObjectFactory *DroneObjectFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DroneObjectFactory.Contract.DroneObjectFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DroneObjectFactory *DroneObjectFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DroneObjectFactory.Contract.DroneObjectFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DroneObjectFactory *DroneObjectFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DroneObjectFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DroneObjectFactory *DroneObjectFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DroneObjectFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DroneObjectFactory *DroneObjectFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DroneObjectFactory.Contract.contract.Transact(opts, method, params...)
}

// DroneExists is a free data retrieval call binding the contract method 0xe73fc130.
//
// Solidity: function droneExists(address _droneId) view returns(bool)
func (_DroneObjectFactory *DroneObjectFactoryCaller) DroneExists(opts *bind.CallOpts, _droneId common.Address) (bool, error) {
	var out []interface{}
	err := _DroneObjectFactory.contract.Call(opts, &out, "droneExists", _droneId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DroneExists is a free data retrieval call binding the contract method 0xe73fc130.
//
// Solidity: function droneExists(address _droneId) view returns(bool)
func (_DroneObjectFactory *DroneObjectFactorySession) DroneExists(_droneId common.Address) (bool, error) {
	return _DroneObjectFactory.Contract.DroneExists(&_DroneObjectFactory.CallOpts, _droneId)
}

// DroneExists is a free data retrieval call binding the contract method 0xe73fc130.
//
// Solidity: function droneExists(address _droneId) view returns(bool)
func (_DroneObjectFactory *DroneObjectFactoryCallerSession) DroneExists(_droneId common.Address) (bool, error) {
	return _DroneObjectFactory.Contract.DroneExists(&_DroneObjectFactory.CallOpts, _droneId)
}

// Drones is a free data retrieval call binding the contract method 0x9d5d6696.
//
// Solidity: function drones(address ) view returns(uint8 droneType, string firmwareVersion, uint8 state, uint32 registeredTime, bool exists)
func (_DroneObjectFactory *DroneObjectFactoryCaller) Drones(opts *bind.CallOpts, arg0 common.Address) (struct {
	DroneType       uint8
	FirmwareVersion string
	State           uint8
	RegisteredTime  uint32
	Exists          bool
}, error) {
	var out []interface{}
	err := _DroneObjectFactory.contract.Call(opts, &out, "drones", arg0)

	outstruct := new(struct {
		DroneType       uint8
		FirmwareVersion string
		State           uint8
		RegisteredTime  uint32
		Exists          bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DroneType = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.FirmwareVersion = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.State = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.RegisteredTime = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.Exists = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Drones is a free data retrieval call binding the contract method 0x9d5d6696.
//
// Solidity: function drones(address ) view returns(uint8 droneType, string firmwareVersion, uint8 state, uint32 registeredTime, bool exists)
func (_DroneObjectFactory *DroneObjectFactorySession) Drones(arg0 common.Address) (struct {
	DroneType       uint8
	FirmwareVersion string
	State           uint8
	RegisteredTime  uint32
	Exists          bool
}, error) {
	return _DroneObjectFactory.Contract.Drones(&_DroneObjectFactory.CallOpts, arg0)
}

// Drones is a free data retrieval call binding the contract method 0x9d5d6696.
//
// Solidity: function drones(address ) view returns(uint8 droneType, string firmwareVersion, uint8 state, uint32 registeredTime, bool exists)
func (_DroneObjectFactory *DroneObjectFactoryCallerSession) Drones(arg0 common.Address) (struct {
	DroneType       uint8
	FirmwareVersion string
	State           uint8
	RegisteredTime  uint32
	Exists          bool
}, error) {
	return _DroneObjectFactory.Contract.Drones(&_DroneObjectFactory.CallOpts, arg0)
}

// GetDrone is a free data retrieval call binding the contract method 0x87c835db.
//
// Solidity: function getDrone(address _droneId) view returns((uint8,string,uint8,uint32,bool))
func (_DroneObjectFactory *DroneObjectFactoryCaller) GetDrone(opts *bind.CallOpts, _droneId common.Address) (DroneObjectFactoryDrone, error) {
	var out []interface{}
	err := _DroneObjectFactory.contract.Call(opts, &out, "getDrone", _droneId)

	if err != nil {
		return *new(DroneObjectFactoryDrone), err
	}

	out0 := *abi.ConvertType(out[0], new(DroneObjectFactoryDrone)).(*DroneObjectFactoryDrone)

	return out0, err

}

// GetDrone is a free data retrieval call binding the contract method 0x87c835db.
//
// Solidity: function getDrone(address _droneId) view returns((uint8,string,uint8,uint32,bool))
func (_DroneObjectFactory *DroneObjectFactorySession) GetDrone(_droneId common.Address) (DroneObjectFactoryDrone, error) {
	return _DroneObjectFactory.Contract.GetDrone(&_DroneObjectFactory.CallOpts, _droneId)
}

// GetDrone is a free data retrieval call binding the contract method 0x87c835db.
//
// Solidity: function getDrone(address _droneId) view returns((uint8,string,uint8,uint32,bool))
func (_DroneObjectFactory *DroneObjectFactoryCallerSession) GetDrone(_droneId common.Address) (DroneObjectFactoryDrone, error) {
	return _DroneObjectFactory.Contract.GetDrone(&_DroneObjectFactory.CallOpts, _droneId)
}

// GetDroneType is a free data retrieval call binding the contract method 0xa8399a4c.
//
// Solidity: function getDroneType(address _droneId) view returns(uint8)
func (_DroneObjectFactory *DroneObjectFactoryCaller) GetDroneType(opts *bind.CallOpts, _droneId common.Address) (uint8, error) {
	var out []interface{}
	err := _DroneObjectFactory.contract.Call(opts, &out, "getDroneType", _droneId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetDroneType is a free data retrieval call binding the contract method 0xa8399a4c.
//
// Solidity: function getDroneType(address _droneId) view returns(uint8)
func (_DroneObjectFactory *DroneObjectFactorySession) GetDroneType(_droneId common.Address) (uint8, error) {
	return _DroneObjectFactory.Contract.GetDroneType(&_DroneObjectFactory.CallOpts, _droneId)
}

// GetDroneType is a free data retrieval call binding the contract method 0xa8399a4c.
//
// Solidity: function getDroneType(address _droneId) view returns(uint8)
func (_DroneObjectFactory *DroneObjectFactoryCallerSession) GetDroneType(_droneId common.Address) (uint8, error) {
	return _DroneObjectFactory.Contract.GetDroneType(&_DroneObjectFactory.CallOpts, _droneId)
}

// GetMaintenanceLog is a free data retrieval call binding the contract method 0x65412d24.
//
// Solidity: function getMaintenanceLog(address _droneId, uint32 _logId) view returns((uint32,uint8,bool,string,address))
func (_DroneObjectFactory *DroneObjectFactoryCaller) GetMaintenanceLog(opts *bind.CallOpts, _droneId common.Address, _logId uint32) (DroneObjectFactoryMaintenanceLog, error) {
	var out []interface{}
	err := _DroneObjectFactory.contract.Call(opts, &out, "getMaintenanceLog", _droneId, _logId)

	if err != nil {
		return *new(DroneObjectFactoryMaintenanceLog), err
	}

	out0 := *abi.ConvertType(out[0], new(DroneObjectFactoryMaintenanceLog)).(*DroneObjectFactoryMaintenanceLog)

	return out0, err

}

// GetMaintenanceLog is a free data retrieval call binding the contract method 0x65412d24.
//
// Solidity: function getMaintenanceLog(address _droneId, uint32 _logId) view returns((uint32,uint8,bool,string,address))
func (_DroneObjectFactory *DroneObjectFactorySession) GetMaintenanceLog(_droneId common.Address, _logId uint32) (DroneObjectFactoryMaintenanceLog, error) {
	return _DroneObjectFactory.Contract.GetMaintenanceLog(&_DroneObjectFactory.CallOpts, _droneId, _logId)
}

// GetMaintenanceLog is a free data retrieval call binding the contract method 0x65412d24.
//
// Solidity: function getMaintenanceLog(address _droneId, uint32 _logId) view returns((uint32,uint8,bool,string,address))
func (_DroneObjectFactory *DroneObjectFactoryCallerSession) GetMaintenanceLog(_droneId common.Address, _logId uint32) (DroneObjectFactoryMaintenanceLog, error) {
	return _DroneObjectFactory.Contract.GetMaintenanceLog(&_DroneObjectFactory.CallOpts, _droneId, _logId)
}

// GetRecentMaintenanceLogs is a free data retrieval call binding the contract method 0xae737fd9.
//
// Solidity: function getRecentMaintenanceLogs(address _droneId, uint32 _count) view returns((uint32,uint8,bool,string,address)[])
func (_DroneObjectFactory *DroneObjectFactoryCaller) GetRecentMaintenanceLogs(opts *bind.CallOpts, _droneId common.Address, _count uint32) ([]DroneObjectFactoryMaintenanceLog, error) {
	var out []interface{}
	err := _DroneObjectFactory.contract.Call(opts, &out, "getRecentMaintenanceLogs", _droneId, _count)

	if err != nil {
		return *new([]DroneObjectFactoryMaintenanceLog), err
	}

	out0 := *abi.ConvertType(out[0], new([]DroneObjectFactoryMaintenanceLog)).(*[]DroneObjectFactoryMaintenanceLog)

	return out0, err

}

// GetRecentMaintenanceLogs is a free data retrieval call binding the contract method 0xae737fd9.
//
// Solidity: function getRecentMaintenanceLogs(address _droneId, uint32 _count) view returns((uint32,uint8,bool,string,address)[])
func (_DroneObjectFactory *DroneObjectFactorySession) GetRecentMaintenanceLogs(_droneId common.Address, _count uint32) ([]DroneObjectFactoryMaintenanceLog, error) {
	return _DroneObjectFactory.Contract.GetRecentMaintenanceLogs(&_DroneObjectFactory.CallOpts, _droneId, _count)
}

// GetRecentMaintenanceLogs is a free data retrieval call binding the contract method 0xae737fd9.
//
// Solidity: function getRecentMaintenanceLogs(address _droneId, uint32 _count) view returns((uint32,uint8,bool,string,address)[])
func (_DroneObjectFactory *DroneObjectFactoryCallerSession) GetRecentMaintenanceLogs(_droneId common.Address, _count uint32) ([]DroneObjectFactoryMaintenanceLog, error) {
	return _DroneObjectFactory.Contract.GetRecentMaintenanceLogs(&_DroneObjectFactory.CallOpts, _droneId, _count)
}

// MaintenanceLogCount is a free data retrieval call binding the contract method 0xccea9068.
//
// Solidity: function maintenanceLogCount(address ) view returns(uint32)
func (_DroneObjectFactory *DroneObjectFactoryCaller) MaintenanceLogCount(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _DroneObjectFactory.contract.Call(opts, &out, "maintenanceLogCount", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaintenanceLogCount is a free data retrieval call binding the contract method 0xccea9068.
//
// Solidity: function maintenanceLogCount(address ) view returns(uint32)
func (_DroneObjectFactory *DroneObjectFactorySession) MaintenanceLogCount(arg0 common.Address) (uint32, error) {
	return _DroneObjectFactory.Contract.MaintenanceLogCount(&_DroneObjectFactory.CallOpts, arg0)
}

// MaintenanceLogCount is a free data retrieval call binding the contract method 0xccea9068.
//
// Solidity: function maintenanceLogCount(address ) view returns(uint32)
func (_DroneObjectFactory *DroneObjectFactoryCallerSession) MaintenanceLogCount(arg0 common.Address) (uint32, error) {
	return _DroneObjectFactory.Contract.MaintenanceLogCount(&_DroneObjectFactory.CallOpts, arg0)
}

// MaintenanceLogs is a free data retrieval call binding the contract method 0xfae6fe6c.
//
// Solidity: function maintenanceLogs(address , uint32 ) view returns(uint32 timestamp, uint8 maintenanceTypeCode, bool sensorCalibrated, string firmwareVersion, address technician)
func (_DroneObjectFactory *DroneObjectFactoryCaller) MaintenanceLogs(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	Timestamp           uint32
	MaintenanceTypeCode uint8
	SensorCalibrated    bool
	FirmwareVersion     string
	Technician          common.Address
}, error) {
	var out []interface{}
	err := _DroneObjectFactory.contract.Call(opts, &out, "maintenanceLogs", arg0, arg1)

	outstruct := new(struct {
		Timestamp           uint32
		MaintenanceTypeCode uint8
		SensorCalibrated    bool
		FirmwareVersion     string
		Technician          common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.MaintenanceTypeCode = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.SensorCalibrated = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.FirmwareVersion = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Technician = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// MaintenanceLogs is a free data retrieval call binding the contract method 0xfae6fe6c.
//
// Solidity: function maintenanceLogs(address , uint32 ) view returns(uint32 timestamp, uint8 maintenanceTypeCode, bool sensorCalibrated, string firmwareVersion, address technician)
func (_DroneObjectFactory *DroneObjectFactorySession) MaintenanceLogs(arg0 common.Address, arg1 uint32) (struct {
	Timestamp           uint32
	MaintenanceTypeCode uint8
	SensorCalibrated    bool
	FirmwareVersion     string
	Technician          common.Address
}, error) {
	return _DroneObjectFactory.Contract.MaintenanceLogs(&_DroneObjectFactory.CallOpts, arg0, arg1)
}

// MaintenanceLogs is a free data retrieval call binding the contract method 0xfae6fe6c.
//
// Solidity: function maintenanceLogs(address , uint32 ) view returns(uint32 timestamp, uint8 maintenanceTypeCode, bool sensorCalibrated, string firmwareVersion, address technician)
func (_DroneObjectFactory *DroneObjectFactoryCallerSession) MaintenanceLogs(arg0 common.Address, arg1 uint32) (struct {
	Timestamp           uint32
	MaintenanceTypeCode uint8
	SensorCalibrated    bool
	FirmwareVersion     string
	Technician          common.Address
}, error) {
	return _DroneObjectFactory.Contract.MaintenanceLogs(&_DroneObjectFactory.CallOpts, arg0, arg1)
}

// UserContract is a free data retrieval call binding the contract method 0x5000596b.
//
// Solidity: function userContract() view returns(address)
func (_DroneObjectFactory *DroneObjectFactoryCaller) UserContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _DroneObjectFactory.contract.Call(opts, &out, "userContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserContract is a free data retrieval call binding the contract method 0x5000596b.
//
// Solidity: function userContract() view returns(address)
func (_DroneObjectFactory *DroneObjectFactorySession) UserContract() (common.Address, error) {
	return _DroneObjectFactory.Contract.UserContract(&_DroneObjectFactory.CallOpts)
}

// UserContract is a free data retrieval call binding the contract method 0x5000596b.
//
// Solidity: function userContract() view returns(address)
func (_DroneObjectFactory *DroneObjectFactoryCallerSession) UserContract() (common.Address, error) {
	return _DroneObjectFactory.Contract.UserContract(&_DroneObjectFactory.CallOpts)
}

// AddMaintenanceLog is a paid mutator transaction binding the contract method 0x07cd1e37.
//
// Solidity: function addMaintenanceLog(address _droneId, uint8 _maintenanceTypeCode, bool _sensorCalibrated, string _firmwareVersion) returns()
func (_DroneObjectFactory *DroneObjectFactoryTransactor) AddMaintenanceLog(opts *bind.TransactOpts, _droneId common.Address, _maintenanceTypeCode uint8, _sensorCalibrated bool, _firmwareVersion string) (*types.Transaction, error) {
	return _DroneObjectFactory.contract.Transact(opts, "addMaintenanceLog", _droneId, _maintenanceTypeCode, _sensorCalibrated, _firmwareVersion)
}

// AddMaintenanceLog is a paid mutator transaction binding the contract method 0x07cd1e37.
//
// Solidity: function addMaintenanceLog(address _droneId, uint8 _maintenanceTypeCode, bool _sensorCalibrated, string _firmwareVersion) returns()
func (_DroneObjectFactory *DroneObjectFactorySession) AddMaintenanceLog(_droneId common.Address, _maintenanceTypeCode uint8, _sensorCalibrated bool, _firmwareVersion string) (*types.Transaction, error) {
	return _DroneObjectFactory.Contract.AddMaintenanceLog(&_DroneObjectFactory.TransactOpts, _droneId, _maintenanceTypeCode, _sensorCalibrated, _firmwareVersion)
}

// AddMaintenanceLog is a paid mutator transaction binding the contract method 0x07cd1e37.
//
// Solidity: function addMaintenanceLog(address _droneId, uint8 _maintenanceTypeCode, bool _sensorCalibrated, string _firmwareVersion) returns()
func (_DroneObjectFactory *DroneObjectFactoryTransactorSession) AddMaintenanceLog(_droneId common.Address, _maintenanceTypeCode uint8, _sensorCalibrated bool, _firmwareVersion string) (*types.Transaction, error) {
	return _DroneObjectFactory.Contract.AddMaintenanceLog(&_DroneObjectFactory.TransactOpts, _droneId, _maintenanceTypeCode, _sensorCalibrated, _firmwareVersion)
}

// RegisterDrone is a paid mutator transaction binding the contract method 0x749cb8f2.
//
// Solidity: function registerDrone(address _droneId, uint8 _droneType, string _firmwareVersion, uint8 _droneState) returns()
func (_DroneObjectFactory *DroneObjectFactoryTransactor) RegisterDrone(opts *bind.TransactOpts, _droneId common.Address, _droneType uint8, _firmwareVersion string, _droneState uint8) (*types.Transaction, error) {
	return _DroneObjectFactory.contract.Transact(opts, "registerDrone", _droneId, _droneType, _firmwareVersion, _droneState)
}

// RegisterDrone is a paid mutator transaction binding the contract method 0x749cb8f2.
//
// Solidity: function registerDrone(address _droneId, uint8 _droneType, string _firmwareVersion, uint8 _droneState) returns()
func (_DroneObjectFactory *DroneObjectFactorySession) RegisterDrone(_droneId common.Address, _droneType uint8, _firmwareVersion string, _droneState uint8) (*types.Transaction, error) {
	return _DroneObjectFactory.Contract.RegisterDrone(&_DroneObjectFactory.TransactOpts, _droneId, _droneType, _firmwareVersion, _droneState)
}

// RegisterDrone is a paid mutator transaction binding the contract method 0x749cb8f2.
//
// Solidity: function registerDrone(address _droneId, uint8 _droneType, string _firmwareVersion, uint8 _droneState) returns()
func (_DroneObjectFactory *DroneObjectFactoryTransactorSession) RegisterDrone(_droneId common.Address, _droneType uint8, _firmwareVersion string, _droneState uint8) (*types.Transaction, error) {
	return _DroneObjectFactory.Contract.RegisterDrone(&_DroneObjectFactory.TransactOpts, _droneId, _droneType, _firmwareVersion, _droneState)
}

// UpdateDroneState is a paid mutator transaction binding the contract method 0x9461dfc1.
//
// Solidity: function updateDroneState(address _droneId, uint8 _newState) returns()
func (_DroneObjectFactory *DroneObjectFactoryTransactor) UpdateDroneState(opts *bind.TransactOpts, _droneId common.Address, _newState uint8) (*types.Transaction, error) {
	return _DroneObjectFactory.contract.Transact(opts, "updateDroneState", _droneId, _newState)
}

// UpdateDroneState is a paid mutator transaction binding the contract method 0x9461dfc1.
//
// Solidity: function updateDroneState(address _droneId, uint8 _newState) returns()
func (_DroneObjectFactory *DroneObjectFactorySession) UpdateDroneState(_droneId common.Address, _newState uint8) (*types.Transaction, error) {
	return _DroneObjectFactory.Contract.UpdateDroneState(&_DroneObjectFactory.TransactOpts, _droneId, _newState)
}

// UpdateDroneState is a paid mutator transaction binding the contract method 0x9461dfc1.
//
// Solidity: function updateDroneState(address _droneId, uint8 _newState) returns()
func (_DroneObjectFactory *DroneObjectFactoryTransactorSession) UpdateDroneState(_droneId common.Address, _newState uint8) (*types.Transaction, error) {
	return _DroneObjectFactory.Contract.UpdateDroneState(&_DroneObjectFactory.TransactOpts, _droneId, _newState)
}

// UpdateFirmwareVersion is a paid mutator transaction binding the contract method 0xe8ad3423.
//
// Solidity: function updateFirmwareVersion(address _droneId, string _newFirmwareVersion) returns()
func (_DroneObjectFactory *DroneObjectFactoryTransactor) UpdateFirmwareVersion(opts *bind.TransactOpts, _droneId common.Address, _newFirmwareVersion string) (*types.Transaction, error) {
	return _DroneObjectFactory.contract.Transact(opts, "updateFirmwareVersion", _droneId, _newFirmwareVersion)
}

// UpdateFirmwareVersion is a paid mutator transaction binding the contract method 0xe8ad3423.
//
// Solidity: function updateFirmwareVersion(address _droneId, string _newFirmwareVersion) returns()
func (_DroneObjectFactory *DroneObjectFactorySession) UpdateFirmwareVersion(_droneId common.Address, _newFirmwareVersion string) (*types.Transaction, error) {
	return _DroneObjectFactory.Contract.UpdateFirmwareVersion(&_DroneObjectFactory.TransactOpts, _droneId, _newFirmwareVersion)
}

// UpdateFirmwareVersion is a paid mutator transaction binding the contract method 0xe8ad3423.
//
// Solidity: function updateFirmwareVersion(address _droneId, string _newFirmwareVersion) returns()
func (_DroneObjectFactory *DroneObjectFactoryTransactorSession) UpdateFirmwareVersion(_droneId common.Address, _newFirmwareVersion string) (*types.Transaction, error) {
	return _DroneObjectFactory.Contract.UpdateFirmwareVersion(&_DroneObjectFactory.TransactOpts, _droneId, _newFirmwareVersion)
}

// DroneObjectFactoryDroneRegisteredIterator is returned from FilterDroneRegistered and is used to iterate over the raw logs and unpacked data for DroneRegistered events raised by the DroneObjectFactory contract.
type DroneObjectFactoryDroneRegisteredIterator struct {
	Event *DroneObjectFactoryDroneRegistered // Event containing the contract specifics and raw log

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
func (it *DroneObjectFactoryDroneRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DroneObjectFactoryDroneRegistered)
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
		it.Event = new(DroneObjectFactoryDroneRegistered)
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
func (it *DroneObjectFactoryDroneRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DroneObjectFactoryDroneRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DroneObjectFactoryDroneRegistered represents a DroneRegistered event raised by the DroneObjectFactory contract.
type DroneObjectFactoryDroneRegistered struct {
	DroneId   common.Address
	DroneType uint8
	Timestamp uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDroneRegistered is a free log retrieval operation binding the contract event 0xb21fb076b25ed1855520f881a12cb2f43fe537c16db2b65eb67f1c28943ed0da.
//
// Solidity: event DroneRegistered(address indexed droneId, uint8 indexed droneType, uint32 timestamp)
func (_DroneObjectFactory *DroneObjectFactoryFilterer) FilterDroneRegistered(opts *bind.FilterOpts, droneId []common.Address, droneType []uint8) (*DroneObjectFactoryDroneRegisteredIterator, error) {

	var droneIdRule []interface{}
	for _, droneIdItem := range droneId {
		droneIdRule = append(droneIdRule, droneIdItem)
	}
	var droneTypeRule []interface{}
	for _, droneTypeItem := range droneType {
		droneTypeRule = append(droneTypeRule, droneTypeItem)
	}

	logs, sub, err := _DroneObjectFactory.contract.FilterLogs(opts, "DroneRegistered", droneIdRule, droneTypeRule)
	if err != nil {
		return nil, err
	}
	return &DroneObjectFactoryDroneRegisteredIterator{contract: _DroneObjectFactory.contract, event: "DroneRegistered", logs: logs, sub: sub}, nil
}

// WatchDroneRegistered is a free log subscription operation binding the contract event 0xb21fb076b25ed1855520f881a12cb2f43fe537c16db2b65eb67f1c28943ed0da.
//
// Solidity: event DroneRegistered(address indexed droneId, uint8 indexed droneType, uint32 timestamp)
func (_DroneObjectFactory *DroneObjectFactoryFilterer) WatchDroneRegistered(opts *bind.WatchOpts, sink chan<- *DroneObjectFactoryDroneRegistered, droneId []common.Address, droneType []uint8) (event.Subscription, error) {

	var droneIdRule []interface{}
	for _, droneIdItem := range droneId {
		droneIdRule = append(droneIdRule, droneIdItem)
	}
	var droneTypeRule []interface{}
	for _, droneTypeItem := range droneType {
		droneTypeRule = append(droneTypeRule, droneTypeItem)
	}

	logs, sub, err := _DroneObjectFactory.contract.WatchLogs(opts, "DroneRegistered", droneIdRule, droneTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DroneObjectFactoryDroneRegistered)
				if err := _DroneObjectFactory.contract.UnpackLog(event, "DroneRegistered", log); err != nil {
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

// ParseDroneRegistered is a log parse operation binding the contract event 0xb21fb076b25ed1855520f881a12cb2f43fe537c16db2b65eb67f1c28943ed0da.
//
// Solidity: event DroneRegistered(address indexed droneId, uint8 indexed droneType, uint32 timestamp)
func (_DroneObjectFactory *DroneObjectFactoryFilterer) ParseDroneRegistered(log types.Log) (*DroneObjectFactoryDroneRegistered, error) {
	event := new(DroneObjectFactoryDroneRegistered)
	if err := _DroneObjectFactory.contract.UnpackLog(event, "DroneRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DroneObjectFactoryDroneStateUpdatedIterator is returned from FilterDroneStateUpdated and is used to iterate over the raw logs and unpacked data for DroneStateUpdated events raised by the DroneObjectFactory contract.
type DroneObjectFactoryDroneStateUpdatedIterator struct {
	Event *DroneObjectFactoryDroneStateUpdated // Event containing the contract specifics and raw log

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
func (it *DroneObjectFactoryDroneStateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DroneObjectFactoryDroneStateUpdated)
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
		it.Event = new(DroneObjectFactoryDroneStateUpdated)
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
func (it *DroneObjectFactoryDroneStateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DroneObjectFactoryDroneStateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DroneObjectFactoryDroneStateUpdated represents a DroneStateUpdated event raised by the DroneObjectFactory contract.
type DroneObjectFactoryDroneStateUpdated struct {
	DroneId  common.Address
	NewState uint8
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDroneStateUpdated is a free log retrieval operation binding the contract event 0xad440e673a8dbd5082542a599644d4b8c7e40964a7a32a22a53631958bbf5320.
//
// Solidity: event DroneStateUpdated(address indexed droneId, uint8 indexed newState)
func (_DroneObjectFactory *DroneObjectFactoryFilterer) FilterDroneStateUpdated(opts *bind.FilterOpts, droneId []common.Address, newState []uint8) (*DroneObjectFactoryDroneStateUpdatedIterator, error) {

	var droneIdRule []interface{}
	for _, droneIdItem := range droneId {
		droneIdRule = append(droneIdRule, droneIdItem)
	}
	var newStateRule []interface{}
	for _, newStateItem := range newState {
		newStateRule = append(newStateRule, newStateItem)
	}

	logs, sub, err := _DroneObjectFactory.contract.FilterLogs(opts, "DroneStateUpdated", droneIdRule, newStateRule)
	if err != nil {
		return nil, err
	}
	return &DroneObjectFactoryDroneStateUpdatedIterator{contract: _DroneObjectFactory.contract, event: "DroneStateUpdated", logs: logs, sub: sub}, nil
}

// WatchDroneStateUpdated is a free log subscription operation binding the contract event 0xad440e673a8dbd5082542a599644d4b8c7e40964a7a32a22a53631958bbf5320.
//
// Solidity: event DroneStateUpdated(address indexed droneId, uint8 indexed newState)
func (_DroneObjectFactory *DroneObjectFactoryFilterer) WatchDroneStateUpdated(opts *bind.WatchOpts, sink chan<- *DroneObjectFactoryDroneStateUpdated, droneId []common.Address, newState []uint8) (event.Subscription, error) {

	var droneIdRule []interface{}
	for _, droneIdItem := range droneId {
		droneIdRule = append(droneIdRule, droneIdItem)
	}
	var newStateRule []interface{}
	for _, newStateItem := range newState {
		newStateRule = append(newStateRule, newStateItem)
	}

	logs, sub, err := _DroneObjectFactory.contract.WatchLogs(opts, "DroneStateUpdated", droneIdRule, newStateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DroneObjectFactoryDroneStateUpdated)
				if err := _DroneObjectFactory.contract.UnpackLog(event, "DroneStateUpdated", log); err != nil {
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

// ParseDroneStateUpdated is a log parse operation binding the contract event 0xad440e673a8dbd5082542a599644d4b8c7e40964a7a32a22a53631958bbf5320.
//
// Solidity: event DroneStateUpdated(address indexed droneId, uint8 indexed newState)
func (_DroneObjectFactory *DroneObjectFactoryFilterer) ParseDroneStateUpdated(log types.Log) (*DroneObjectFactoryDroneStateUpdated, error) {
	event := new(DroneObjectFactoryDroneStateUpdated)
	if err := _DroneObjectFactory.contract.UnpackLog(event, "DroneStateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DroneObjectFactoryFirmwareUpdatedIterator is returned from FilterFirmwareUpdated and is used to iterate over the raw logs and unpacked data for FirmwareUpdated events raised by the DroneObjectFactory contract.
type DroneObjectFactoryFirmwareUpdatedIterator struct {
	Event *DroneObjectFactoryFirmwareUpdated // Event containing the contract specifics and raw log

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
func (it *DroneObjectFactoryFirmwareUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DroneObjectFactoryFirmwareUpdated)
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
		it.Event = new(DroneObjectFactoryFirmwareUpdated)
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
func (it *DroneObjectFactoryFirmwareUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DroneObjectFactoryFirmwareUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DroneObjectFactoryFirmwareUpdated represents a FirmwareUpdated event raised by the DroneObjectFactory contract.
type DroneObjectFactoryFirmwareUpdated struct {
	DroneId    common.Address
	NewVersion string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFirmwareUpdated is a free log retrieval operation binding the contract event 0x77035d4457088e0a3b8a1779932e3091c2f6385591643c8978f6f3c67fc5d8e1.
//
// Solidity: event FirmwareUpdated(address indexed droneId, string newVersion)
func (_DroneObjectFactory *DroneObjectFactoryFilterer) FilterFirmwareUpdated(opts *bind.FilterOpts, droneId []common.Address) (*DroneObjectFactoryFirmwareUpdatedIterator, error) {

	var droneIdRule []interface{}
	for _, droneIdItem := range droneId {
		droneIdRule = append(droneIdRule, droneIdItem)
	}

	logs, sub, err := _DroneObjectFactory.contract.FilterLogs(opts, "FirmwareUpdated", droneIdRule)
	if err != nil {
		return nil, err
	}
	return &DroneObjectFactoryFirmwareUpdatedIterator{contract: _DroneObjectFactory.contract, event: "FirmwareUpdated", logs: logs, sub: sub}, nil
}

// WatchFirmwareUpdated is a free log subscription operation binding the contract event 0x77035d4457088e0a3b8a1779932e3091c2f6385591643c8978f6f3c67fc5d8e1.
//
// Solidity: event FirmwareUpdated(address indexed droneId, string newVersion)
func (_DroneObjectFactory *DroneObjectFactoryFilterer) WatchFirmwareUpdated(opts *bind.WatchOpts, sink chan<- *DroneObjectFactoryFirmwareUpdated, droneId []common.Address) (event.Subscription, error) {

	var droneIdRule []interface{}
	for _, droneIdItem := range droneId {
		droneIdRule = append(droneIdRule, droneIdItem)
	}

	logs, sub, err := _DroneObjectFactory.contract.WatchLogs(opts, "FirmwareUpdated", droneIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DroneObjectFactoryFirmwareUpdated)
				if err := _DroneObjectFactory.contract.UnpackLog(event, "FirmwareUpdated", log); err != nil {
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

// ParseFirmwareUpdated is a log parse operation binding the contract event 0x77035d4457088e0a3b8a1779932e3091c2f6385591643c8978f6f3c67fc5d8e1.
//
// Solidity: event FirmwareUpdated(address indexed droneId, string newVersion)
func (_DroneObjectFactory *DroneObjectFactoryFilterer) ParseFirmwareUpdated(log types.Log) (*DroneObjectFactoryFirmwareUpdated, error) {
	event := new(DroneObjectFactoryFirmwareUpdated)
	if err := _DroneObjectFactory.contract.UnpackLog(event, "FirmwareUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// DroneObjectFactoryMaintenanceLogAddedIterator is returned from FilterMaintenanceLogAdded and is used to iterate over the raw logs and unpacked data for MaintenanceLogAdded events raised by the DroneObjectFactory contract.
type DroneObjectFactoryMaintenanceLogAddedIterator struct {
	Event *DroneObjectFactoryMaintenanceLogAdded // Event containing the contract specifics and raw log

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
func (it *DroneObjectFactoryMaintenanceLogAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DroneObjectFactoryMaintenanceLogAdded)
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
		it.Event = new(DroneObjectFactoryMaintenanceLogAdded)
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
func (it *DroneObjectFactoryMaintenanceLogAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DroneObjectFactoryMaintenanceLogAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DroneObjectFactoryMaintenanceLogAdded represents a MaintenanceLogAdded event raised by the DroneObjectFactory contract.
type DroneObjectFactoryMaintenanceLogAdded struct {
	DroneId         common.Address
	LogId           uint32
	MaintenanceType uint8
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMaintenanceLogAdded is a free log retrieval operation binding the contract event 0x4a2b3e3e097f8d105c4c8f27d112da4f6ef371edc6efeac4d375c78f2e047d1b.
//
// Solidity: event MaintenanceLogAdded(address indexed droneId, uint32 indexed logId, uint8 maintenanceType)
func (_DroneObjectFactory *DroneObjectFactoryFilterer) FilterMaintenanceLogAdded(opts *bind.FilterOpts, droneId []common.Address, logId []uint32) (*DroneObjectFactoryMaintenanceLogAddedIterator, error) {

	var droneIdRule []interface{}
	for _, droneIdItem := range droneId {
		droneIdRule = append(droneIdRule, droneIdItem)
	}
	var logIdRule []interface{}
	for _, logIdItem := range logId {
		logIdRule = append(logIdRule, logIdItem)
	}

	logs, sub, err := _DroneObjectFactory.contract.FilterLogs(opts, "MaintenanceLogAdded", droneIdRule, logIdRule)
	if err != nil {
		return nil, err
	}
	return &DroneObjectFactoryMaintenanceLogAddedIterator{contract: _DroneObjectFactory.contract, event: "MaintenanceLogAdded", logs: logs, sub: sub}, nil
}

// WatchMaintenanceLogAdded is a free log subscription operation binding the contract event 0x4a2b3e3e097f8d105c4c8f27d112da4f6ef371edc6efeac4d375c78f2e047d1b.
//
// Solidity: event MaintenanceLogAdded(address indexed droneId, uint32 indexed logId, uint8 maintenanceType)
func (_DroneObjectFactory *DroneObjectFactoryFilterer) WatchMaintenanceLogAdded(opts *bind.WatchOpts, sink chan<- *DroneObjectFactoryMaintenanceLogAdded, droneId []common.Address, logId []uint32) (event.Subscription, error) {

	var droneIdRule []interface{}
	for _, droneIdItem := range droneId {
		droneIdRule = append(droneIdRule, droneIdItem)
	}
	var logIdRule []interface{}
	for _, logIdItem := range logId {
		logIdRule = append(logIdRule, logIdItem)
	}

	logs, sub, err := _DroneObjectFactory.contract.WatchLogs(opts, "MaintenanceLogAdded", droneIdRule, logIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DroneObjectFactoryMaintenanceLogAdded)
				if err := _DroneObjectFactory.contract.UnpackLog(event, "MaintenanceLogAdded", log); err != nil {
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

// ParseMaintenanceLogAdded is a log parse operation binding the contract event 0x4a2b3e3e097f8d105c4c8f27d112da4f6ef371edc6efeac4d375c78f2e047d1b.
//
// Solidity: event MaintenanceLogAdded(address indexed droneId, uint32 indexed logId, uint8 maintenanceType)
func (_DroneObjectFactory *DroneObjectFactoryFilterer) ParseMaintenanceLogAdded(log types.Log) (*DroneObjectFactoryMaintenanceLogAdded, error) {
	event := new(DroneObjectFactoryMaintenanceLogAdded)
	if err := _DroneObjectFactory.contract.UnpackLog(event, "MaintenanceLogAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
