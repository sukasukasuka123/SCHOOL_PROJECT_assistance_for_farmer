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
	DroneId         common.Address
	DroneTypeValue  uint8
	FirmwareVersion string
	DroneState      uint8
}

// DroneObjectFactoryDroneMaintenanceLog is an auto generated low-level Go binding around an user-defined struct.
type DroneObjectFactoryDroneMaintenanceLog struct {
	LogId             uint32
	DroneId           common.Address
	Timestamp         uint32
	DroneType         string
	MaintenanceType   string
	SensorCalibrated  bool
	FirmwareVersion   string
	TechnicianAddress common.Address
}

// DroneObjectFactoryMetaData contains all meta data concerning the DroneObjectFactory contract.
var DroneObjectFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_logId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"_timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"_droneType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_maintenanceType\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_sensorCalibrated\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"_firmwareVersion\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_technicianAddress\",\"type\":\"address\"}],\"name\":\"addDroneMaintenanceLog\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"droneMaintenanceLogs\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"logId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"droneId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"droneType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"maintenanceType\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"sensorCalibrated\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"firmwareVersion\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"technicianAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"drones\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"droneId\",\"type\":\"address\"},{\"internalType\":\"enumDroneObjectFactory.DroneType\",\"name\":\"droneTypeValue\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"firmwareVersion\",\"type\":\"string\"},{\"internalType\":\"enumDroneObjectFactory.state\",\"name\":\"droneState\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_logId\",\"type\":\"uint32\"}],\"name\":\"getDronMaintenanceLogById\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"logId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"droneId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"droneType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"maintenanceType\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"sensorCalibrated\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"firmwareVersion\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"technicianAddress\",\"type\":\"address\"}],\"internalType\":\"structDroneObjectFactory.DroneMaintenanceLog\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"page\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"pageSize\",\"type\":\"uint256\"}],\"name\":\"getDronMaintenanceLogInPage\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"logId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"droneId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"droneType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"maintenanceType\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"sensorCalibrated\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"firmwareVersion\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"technicianAddress\",\"type\":\"address\"}],\"internalType\":\"structDroneObjectFactory.DroneMaintenanceLog[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"}],\"name\":\"getDrone\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"droneId\",\"type\":\"address\"},{\"internalType\":\"enumDroneObjectFactory.DroneType\",\"name\":\"droneTypeValue\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"firmwareVersion\",\"type\":\"string\"},{\"internalType\":\"enumDroneObjectFactory.state\",\"name\":\"droneState\",\"type\":\"uint8\"}],\"internalType\":\"structDroneObjectFactory.Drone\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getDroneMaintenanceLogByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"logId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"droneId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"droneType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"maintenanceType\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"sensorCalibrated\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"firmwareVersion\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"technicianAddress\",\"type\":\"address\"}],\"internalType\":\"structDroneObjectFactory.DroneMaintenanceLog\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"}],\"name\":\"getDroneMaintenanceLogCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"}],\"name\":\"getDroneMaintenanceLogs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"logId\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"droneId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"droneType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"maintenanceType\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"sensorCalibrated\",\"type\":\"bool\"},{\"internalType\":\"string\",\"name\":\"firmwareVersion\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"technicianAddress\",\"type\":\"address\"}],\"internalType\":\"structDroneObjectFactory.DroneMaintenanceLog[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"},{\"internalType\":\"enumDroneObjectFactory.DroneType\",\"name\":\"_droneType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_firmwareVersion\",\"type\":\"string\"},{\"internalType\":\"enumDroneObjectFactory.state\",\"name\":\"_droneState\",\"type\":\"uint8\"}],\"name\":\"registerDrone\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"},{\"internalType\":\"enumDroneObjectFactory.state\",\"name\":\"_newState\",\"type\":\"uint8\"}],\"name\":\"updateDroneState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"},{\"internalType\":\"enumDroneObjectFactory.DroneType\",\"name\":\"_newDroneType\",\"type\":\"uint8\"}],\"name\":\"updateDroneType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_droneId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_newFirmwareVersion\",\"type\":\"string\"}],\"name\":\"updateFirmwareVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// DroneMaintenanceLogs is a free data retrieval call binding the contract method 0xd1451ea3.
//
// Solidity: function droneMaintenanceLogs(address , uint256 ) view returns(uint32 logId, address droneId, uint32 timestamp, string droneType, string maintenanceType, bool sensorCalibrated, string firmwareVersion, address technicianAddress)
func (_DroneObjectFactory *DroneObjectFactoryCaller) DroneMaintenanceLogs(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	LogId             uint32
	DroneId           common.Address
	Timestamp         uint32
	DroneType         string
	MaintenanceType   string
	SensorCalibrated  bool
	FirmwareVersion   string
	TechnicianAddress common.Address
}, error) {
	var out []interface{}
	err := _DroneObjectFactory.contract.Call(opts, &out, "droneMaintenanceLogs", arg0, arg1)

	outstruct := new(struct {
		LogId             uint32
		DroneId           common.Address
		Timestamp         uint32
		DroneType         string
		MaintenanceType   string
		SensorCalibrated  bool
		FirmwareVersion   string
		TechnicianAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LogId = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.DroneId = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.DroneType = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.MaintenanceType = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.SensorCalibrated = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.FirmwareVersion = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.TechnicianAddress = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// DroneMaintenanceLogs is a free data retrieval call binding the contract method 0xd1451ea3.
//
// Solidity: function droneMaintenanceLogs(address , uint256 ) view returns(uint32 logId, address droneId, uint32 timestamp, string droneType, string maintenanceType, bool sensorCalibrated, string firmwareVersion, address technicianAddress)
func (_DroneObjectFactory *DroneObjectFactorySession) DroneMaintenanceLogs(arg0 common.Address, arg1 *big.Int) (struct {
	LogId             uint32
	DroneId           common.Address
	Timestamp         uint32
	DroneType         string
	MaintenanceType   string
	SensorCalibrated  bool
	FirmwareVersion   string
	TechnicianAddress common.Address
}, error) {
	return _DroneObjectFactory.Contract.DroneMaintenanceLogs(&_DroneObjectFactory.CallOpts, arg0, arg1)
}

// DroneMaintenanceLogs is a free data retrieval call binding the contract method 0xd1451ea3.
//
// Solidity: function droneMaintenanceLogs(address , uint256 ) view returns(uint32 logId, address droneId, uint32 timestamp, string droneType, string maintenanceType, bool sensorCalibrated, string firmwareVersion, address technicianAddress)
func (_DroneObjectFactory *DroneObjectFactoryCallerSession) DroneMaintenanceLogs(arg0 common.Address, arg1 *big.Int) (struct {
	LogId             uint32
	DroneId           common.Address
	Timestamp         uint32
	DroneType         string
	MaintenanceType   string
	SensorCalibrated  bool
	FirmwareVersion   string
	TechnicianAddress common.Address
}, error) {
	return _DroneObjectFactory.Contract.DroneMaintenanceLogs(&_DroneObjectFactory.CallOpts, arg0, arg1)
}

// Drones is a free data retrieval call binding the contract method 0x9d5d6696.
//
// Solidity: function drones(address ) view returns(address droneId, uint8 droneTypeValue, string firmwareVersion, uint8 droneState)
func (_DroneObjectFactory *DroneObjectFactoryCaller) Drones(opts *bind.CallOpts, arg0 common.Address) (struct {
	DroneId         common.Address
	DroneTypeValue  uint8
	FirmwareVersion string
	DroneState      uint8
}, error) {
	var out []interface{}
	err := _DroneObjectFactory.contract.Call(opts, &out, "drones", arg0)

	outstruct := new(struct {
		DroneId         common.Address
		DroneTypeValue  uint8
		FirmwareVersion string
		DroneState      uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DroneId = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.DroneTypeValue = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.FirmwareVersion = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.DroneState = *abi.ConvertType(out[3], new(uint8)).(*uint8)

	return *outstruct, err

}

// Drones is a free data retrieval call binding the contract method 0x9d5d6696.
//
// Solidity: function drones(address ) view returns(address droneId, uint8 droneTypeValue, string firmwareVersion, uint8 droneState)
func (_DroneObjectFactory *DroneObjectFactorySession) Drones(arg0 common.Address) (struct {
	DroneId         common.Address
	DroneTypeValue  uint8
	FirmwareVersion string
	DroneState      uint8
}, error) {
	return _DroneObjectFactory.Contract.Drones(&_DroneObjectFactory.CallOpts, arg0)
}

// Drones is a free data retrieval call binding the contract method 0x9d5d6696.
//
// Solidity: function drones(address ) view returns(address droneId, uint8 droneTypeValue, string firmwareVersion, uint8 droneState)
func (_DroneObjectFactory *DroneObjectFactoryCallerSession) Drones(arg0 common.Address) (struct {
	DroneId         common.Address
	DroneTypeValue  uint8
	FirmwareVersion string
	DroneState      uint8
}, error) {
	return _DroneObjectFactory.Contract.Drones(&_DroneObjectFactory.CallOpts, arg0)
}

// GetDronMaintenanceLogById is a free data retrieval call binding the contract method 0xc6b89de6.
//
// Solidity: function getDronMaintenanceLogById(address _droneId, uint32 _logId) view returns((uint32,address,uint32,string,string,bool,string,address))
func (_DroneObjectFactory *DroneObjectFactoryCaller) GetDronMaintenanceLogById(opts *bind.CallOpts, _droneId common.Address, _logId uint32) (DroneObjectFactoryDroneMaintenanceLog, error) {
	var out []interface{}
	err := _DroneObjectFactory.contract.Call(opts, &out, "getDronMaintenanceLogById", _droneId, _logId)

	if err != nil {
		return *new(DroneObjectFactoryDroneMaintenanceLog), err
	}

	out0 := *abi.ConvertType(out[0], new(DroneObjectFactoryDroneMaintenanceLog)).(*DroneObjectFactoryDroneMaintenanceLog)

	return out0, err

}

// GetDronMaintenanceLogById is a free data retrieval call binding the contract method 0xc6b89de6.
//
// Solidity: function getDronMaintenanceLogById(address _droneId, uint32 _logId) view returns((uint32,address,uint32,string,string,bool,string,address))
func (_DroneObjectFactory *DroneObjectFactorySession) GetDronMaintenanceLogById(_droneId common.Address, _logId uint32) (DroneObjectFactoryDroneMaintenanceLog, error) {
	return _DroneObjectFactory.Contract.GetDronMaintenanceLogById(&_DroneObjectFactory.CallOpts, _droneId, _logId)
}

// GetDronMaintenanceLogById is a free data retrieval call binding the contract method 0xc6b89de6.
//
// Solidity: function getDronMaintenanceLogById(address _droneId, uint32 _logId) view returns((uint32,address,uint32,string,string,bool,string,address))
func (_DroneObjectFactory *DroneObjectFactoryCallerSession) GetDronMaintenanceLogById(_droneId common.Address, _logId uint32) (DroneObjectFactoryDroneMaintenanceLog, error) {
	return _DroneObjectFactory.Contract.GetDronMaintenanceLogById(&_DroneObjectFactory.CallOpts, _droneId, _logId)
}

// GetDronMaintenanceLogInPage is a free data retrieval call binding the contract method 0xd204e1dd.
//
// Solidity: function getDronMaintenanceLogInPage(address _droneId, uint256 page, uint256 pageSize) view returns((uint32,address,uint32,string,string,bool,string,address)[])
func (_DroneObjectFactory *DroneObjectFactoryCaller) GetDronMaintenanceLogInPage(opts *bind.CallOpts, _droneId common.Address, page *big.Int, pageSize *big.Int) ([]DroneObjectFactoryDroneMaintenanceLog, error) {
	var out []interface{}
	err := _DroneObjectFactory.contract.Call(opts, &out, "getDronMaintenanceLogInPage", _droneId, page, pageSize)

	if err != nil {
		return *new([]DroneObjectFactoryDroneMaintenanceLog), err
	}

	out0 := *abi.ConvertType(out[0], new([]DroneObjectFactoryDroneMaintenanceLog)).(*[]DroneObjectFactoryDroneMaintenanceLog)

	return out0, err

}

// GetDronMaintenanceLogInPage is a free data retrieval call binding the contract method 0xd204e1dd.
//
// Solidity: function getDronMaintenanceLogInPage(address _droneId, uint256 page, uint256 pageSize) view returns((uint32,address,uint32,string,string,bool,string,address)[])
func (_DroneObjectFactory *DroneObjectFactorySession) GetDronMaintenanceLogInPage(_droneId common.Address, page *big.Int, pageSize *big.Int) ([]DroneObjectFactoryDroneMaintenanceLog, error) {
	return _DroneObjectFactory.Contract.GetDronMaintenanceLogInPage(&_DroneObjectFactory.CallOpts, _droneId, page, pageSize)
}

// GetDronMaintenanceLogInPage is a free data retrieval call binding the contract method 0xd204e1dd.
//
// Solidity: function getDronMaintenanceLogInPage(address _droneId, uint256 page, uint256 pageSize) view returns((uint32,address,uint32,string,string,bool,string,address)[])
func (_DroneObjectFactory *DroneObjectFactoryCallerSession) GetDronMaintenanceLogInPage(_droneId common.Address, page *big.Int, pageSize *big.Int) ([]DroneObjectFactoryDroneMaintenanceLog, error) {
	return _DroneObjectFactory.Contract.GetDronMaintenanceLogInPage(&_DroneObjectFactory.CallOpts, _droneId, page, pageSize)
}

// GetDrone is a free data retrieval call binding the contract method 0x87c835db.
//
// Solidity: function getDrone(address _droneId) view returns((address,uint8,string,uint8))
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
// Solidity: function getDrone(address _droneId) view returns((address,uint8,string,uint8))
func (_DroneObjectFactory *DroneObjectFactorySession) GetDrone(_droneId common.Address) (DroneObjectFactoryDrone, error) {
	return _DroneObjectFactory.Contract.GetDrone(&_DroneObjectFactory.CallOpts, _droneId)
}

// GetDrone is a free data retrieval call binding the contract method 0x87c835db.
//
// Solidity: function getDrone(address _droneId) view returns((address,uint8,string,uint8))
func (_DroneObjectFactory *DroneObjectFactoryCallerSession) GetDrone(_droneId common.Address) (DroneObjectFactoryDrone, error) {
	return _DroneObjectFactory.Contract.GetDrone(&_DroneObjectFactory.CallOpts, _droneId)
}

// GetDroneMaintenanceLogByIndex is a free data retrieval call binding the contract method 0x6e5c405e.
//
// Solidity: function getDroneMaintenanceLogByIndex(address _droneId, uint256 index) view returns((uint32,address,uint32,string,string,bool,string,address))
func (_DroneObjectFactory *DroneObjectFactoryCaller) GetDroneMaintenanceLogByIndex(opts *bind.CallOpts, _droneId common.Address, index *big.Int) (DroneObjectFactoryDroneMaintenanceLog, error) {
	var out []interface{}
	err := _DroneObjectFactory.contract.Call(opts, &out, "getDroneMaintenanceLogByIndex", _droneId, index)

	if err != nil {
		return *new(DroneObjectFactoryDroneMaintenanceLog), err
	}

	out0 := *abi.ConvertType(out[0], new(DroneObjectFactoryDroneMaintenanceLog)).(*DroneObjectFactoryDroneMaintenanceLog)

	return out0, err

}

// GetDroneMaintenanceLogByIndex is a free data retrieval call binding the contract method 0x6e5c405e.
//
// Solidity: function getDroneMaintenanceLogByIndex(address _droneId, uint256 index) view returns((uint32,address,uint32,string,string,bool,string,address))
func (_DroneObjectFactory *DroneObjectFactorySession) GetDroneMaintenanceLogByIndex(_droneId common.Address, index *big.Int) (DroneObjectFactoryDroneMaintenanceLog, error) {
	return _DroneObjectFactory.Contract.GetDroneMaintenanceLogByIndex(&_DroneObjectFactory.CallOpts, _droneId, index)
}

// GetDroneMaintenanceLogByIndex is a free data retrieval call binding the contract method 0x6e5c405e.
//
// Solidity: function getDroneMaintenanceLogByIndex(address _droneId, uint256 index) view returns((uint32,address,uint32,string,string,bool,string,address))
func (_DroneObjectFactory *DroneObjectFactoryCallerSession) GetDroneMaintenanceLogByIndex(_droneId common.Address, index *big.Int) (DroneObjectFactoryDroneMaintenanceLog, error) {
	return _DroneObjectFactory.Contract.GetDroneMaintenanceLogByIndex(&_DroneObjectFactory.CallOpts, _droneId, index)
}

// GetDroneMaintenanceLogCount is a free data retrieval call binding the contract method 0xdf98662d.
//
// Solidity: function getDroneMaintenanceLogCount(address _droneId) view returns(uint256)
func (_DroneObjectFactory *DroneObjectFactoryCaller) GetDroneMaintenanceLogCount(opts *bind.CallOpts, _droneId common.Address) (*big.Int, error) {
	var out []interface{}
	err := _DroneObjectFactory.contract.Call(opts, &out, "getDroneMaintenanceLogCount", _droneId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDroneMaintenanceLogCount is a free data retrieval call binding the contract method 0xdf98662d.
//
// Solidity: function getDroneMaintenanceLogCount(address _droneId) view returns(uint256)
func (_DroneObjectFactory *DroneObjectFactorySession) GetDroneMaintenanceLogCount(_droneId common.Address) (*big.Int, error) {
	return _DroneObjectFactory.Contract.GetDroneMaintenanceLogCount(&_DroneObjectFactory.CallOpts, _droneId)
}

// GetDroneMaintenanceLogCount is a free data retrieval call binding the contract method 0xdf98662d.
//
// Solidity: function getDroneMaintenanceLogCount(address _droneId) view returns(uint256)
func (_DroneObjectFactory *DroneObjectFactoryCallerSession) GetDroneMaintenanceLogCount(_droneId common.Address) (*big.Int, error) {
	return _DroneObjectFactory.Contract.GetDroneMaintenanceLogCount(&_DroneObjectFactory.CallOpts, _droneId)
}

// GetDroneMaintenanceLogs is a free data retrieval call binding the contract method 0xfc85828c.
//
// Solidity: function getDroneMaintenanceLogs(address _droneId) view returns((uint32,address,uint32,string,string,bool,string,address)[])
func (_DroneObjectFactory *DroneObjectFactoryCaller) GetDroneMaintenanceLogs(opts *bind.CallOpts, _droneId common.Address) ([]DroneObjectFactoryDroneMaintenanceLog, error) {
	var out []interface{}
	err := _DroneObjectFactory.contract.Call(opts, &out, "getDroneMaintenanceLogs", _droneId)

	if err != nil {
		return *new([]DroneObjectFactoryDroneMaintenanceLog), err
	}

	out0 := *abi.ConvertType(out[0], new([]DroneObjectFactoryDroneMaintenanceLog)).(*[]DroneObjectFactoryDroneMaintenanceLog)

	return out0, err

}

// GetDroneMaintenanceLogs is a free data retrieval call binding the contract method 0xfc85828c.
//
// Solidity: function getDroneMaintenanceLogs(address _droneId) view returns((uint32,address,uint32,string,string,bool,string,address)[])
func (_DroneObjectFactory *DroneObjectFactorySession) GetDroneMaintenanceLogs(_droneId common.Address) ([]DroneObjectFactoryDroneMaintenanceLog, error) {
	return _DroneObjectFactory.Contract.GetDroneMaintenanceLogs(&_DroneObjectFactory.CallOpts, _droneId)
}

// GetDroneMaintenanceLogs is a free data retrieval call binding the contract method 0xfc85828c.
//
// Solidity: function getDroneMaintenanceLogs(address _droneId) view returns((uint32,address,uint32,string,string,bool,string,address)[])
func (_DroneObjectFactory *DroneObjectFactoryCallerSession) GetDroneMaintenanceLogs(_droneId common.Address) ([]DroneObjectFactoryDroneMaintenanceLog, error) {
	return _DroneObjectFactory.Contract.GetDroneMaintenanceLogs(&_DroneObjectFactory.CallOpts, _droneId)
}

// AddDroneMaintenanceLog is a paid mutator transaction binding the contract method 0x868946c1.
//
// Solidity: function addDroneMaintenanceLog(address _droneId, uint32 _logId, uint32 _timestamp, string _droneType, string _maintenanceType, bool _sensorCalibrated, string _firmwareVersion, address _technicianAddress) returns()
func (_DroneObjectFactory *DroneObjectFactoryTransactor) AddDroneMaintenanceLog(opts *bind.TransactOpts, _droneId common.Address, _logId uint32, _timestamp uint32, _droneType string, _maintenanceType string, _sensorCalibrated bool, _firmwareVersion string, _technicianAddress common.Address) (*types.Transaction, error) {
	return _DroneObjectFactory.contract.Transact(opts, "addDroneMaintenanceLog", _droneId, _logId, _timestamp, _droneType, _maintenanceType, _sensorCalibrated, _firmwareVersion, _technicianAddress)
}

// AddDroneMaintenanceLog is a paid mutator transaction binding the contract method 0x868946c1.
//
// Solidity: function addDroneMaintenanceLog(address _droneId, uint32 _logId, uint32 _timestamp, string _droneType, string _maintenanceType, bool _sensorCalibrated, string _firmwareVersion, address _technicianAddress) returns()
func (_DroneObjectFactory *DroneObjectFactorySession) AddDroneMaintenanceLog(_droneId common.Address, _logId uint32, _timestamp uint32, _droneType string, _maintenanceType string, _sensorCalibrated bool, _firmwareVersion string, _technicianAddress common.Address) (*types.Transaction, error) {
	return _DroneObjectFactory.Contract.AddDroneMaintenanceLog(&_DroneObjectFactory.TransactOpts, _droneId, _logId, _timestamp, _droneType, _maintenanceType, _sensorCalibrated, _firmwareVersion, _technicianAddress)
}

// AddDroneMaintenanceLog is a paid mutator transaction binding the contract method 0x868946c1.
//
// Solidity: function addDroneMaintenanceLog(address _droneId, uint32 _logId, uint32 _timestamp, string _droneType, string _maintenanceType, bool _sensorCalibrated, string _firmwareVersion, address _technicianAddress) returns()
func (_DroneObjectFactory *DroneObjectFactoryTransactorSession) AddDroneMaintenanceLog(_droneId common.Address, _logId uint32, _timestamp uint32, _droneType string, _maintenanceType string, _sensorCalibrated bool, _firmwareVersion string, _technicianAddress common.Address) (*types.Transaction, error) {
	return _DroneObjectFactory.Contract.AddDroneMaintenanceLog(&_DroneObjectFactory.TransactOpts, _droneId, _logId, _timestamp, _droneType, _maintenanceType, _sensorCalibrated, _firmwareVersion, _technicianAddress)
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

// UpdateDroneType is a paid mutator transaction binding the contract method 0x6763d8e7.
//
// Solidity: function updateDroneType(address _droneId, uint8 _newDroneType) returns()
func (_DroneObjectFactory *DroneObjectFactoryTransactor) UpdateDroneType(opts *bind.TransactOpts, _droneId common.Address, _newDroneType uint8) (*types.Transaction, error) {
	return _DroneObjectFactory.contract.Transact(opts, "updateDroneType", _droneId, _newDroneType)
}

// UpdateDroneType is a paid mutator transaction binding the contract method 0x6763d8e7.
//
// Solidity: function updateDroneType(address _droneId, uint8 _newDroneType) returns()
func (_DroneObjectFactory *DroneObjectFactorySession) UpdateDroneType(_droneId common.Address, _newDroneType uint8) (*types.Transaction, error) {
	return _DroneObjectFactory.Contract.UpdateDroneType(&_DroneObjectFactory.TransactOpts, _droneId, _newDroneType)
}

// UpdateDroneType is a paid mutator transaction binding the contract method 0x6763d8e7.
//
// Solidity: function updateDroneType(address _droneId, uint8 _newDroneType) returns()
func (_DroneObjectFactory *DroneObjectFactoryTransactorSession) UpdateDroneType(_droneId common.Address, _newDroneType uint8) (*types.Transaction, error) {
	return _DroneObjectFactory.Contract.UpdateDroneType(&_DroneObjectFactory.TransactOpts, _droneId, _newDroneType)
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
