// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FarmFactory

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

// FarmFactoryFarm is an auto generated low-level Go binding around an user-defined struct.
type FarmFactoryFarm struct {
	Name        string
	Owner       common.Address
	CreatedTime uint32
	Exists      bool
}

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

// FarmFactoryMetaData contains all meta data concerning the FarmFactory contract.
var FarmFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gridFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_userContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_droneFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"FarmCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"}],\"name\":\"GridAddedToFarm\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recordManager\",\"type\":\"address\"}],\"name\":\"RecordManagerSet\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"operationTypeCode\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"detail\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"evidenceIPFSHash\",\"type\":\"bytes32\"}],\"name\":\"addGridMaintenance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"enumGridObjectFactory.GridTerrainType\",\"name\":\"terrainType\",\"type\":\"uint8\"}],\"name\":\"addGridToFarm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"createFarm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"droneFactory\",\"outputs\":[{\"internalType\":\"contractDroneObjectFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"}],\"name\":\"farmExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"farmGridCodeToId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"farmGridList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"farms\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"createdTime\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"}],\"name\":\"getFarm\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"createdTime\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structFarmFactory.Farm\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"}],\"name\":\"getFarmGrids\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"}],\"name\":\"getGridIdByCode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"}],\"name\":\"getGridInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"enumGridObjectFactory.GridTerrainType\",\"name\":\"terrainType\",\"type\":\"uint8\"},{\"internalType\":\"enumGridObjectFactory.GridStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"enumGridObjectFactory.CitrusDiseaseType\",\"name\":\"diseaseType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"lastStatusUpdateTime\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structGridObjectFactory.GridInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"}],\"name\":\"gridExistsInFarm\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gridFactory\",\"outputs\":[{\"internalType\":\"contractGridObjectFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"isFarmOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recordManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_recordManager\",\"type\":\"address\"}],\"name\":\"setRecordManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"enumGridObjectFactory.CitrusDiseaseType\",\"name\":\"newType\",\"type\":\"uint8\"}],\"name\":\"updateGridDiseaseType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"enumGridObjectFactory.CitrusDiseaseType\",\"name\":\"newType\",\"type\":\"uint8\"}],\"name\":\"updateGridDiseaseTypeFromRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"enumGridObjectFactory.GridStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"updateGridStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"enumGridObjectFactory.GridStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"updateGridStatusFromRecord\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userContract\",\"outputs\":[{\"internalType\":\"contractUser\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FarmFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use FarmFactoryMetaData.ABI instead.
var FarmFactoryABI = FarmFactoryMetaData.ABI

// FarmFactory is an auto generated Go binding around an Ethereum contract.
type FarmFactory struct {
	FarmFactoryCaller     // Read-only binding to the contract
	FarmFactoryTransactor // Write-only binding to the contract
	FarmFactoryFilterer   // Log filterer for contract events
}

// FarmFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type FarmFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FarmFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FarmFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FarmFactorySession struct {
	Contract     *FarmFactory      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FarmFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FarmFactoryCallerSession struct {
	Contract *FarmFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FarmFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FarmFactoryTransactorSession struct {
	Contract     *FarmFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FarmFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type FarmFactoryRaw struct {
	Contract *FarmFactory // Generic contract binding to access the raw methods on
}

// FarmFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FarmFactoryCallerRaw struct {
	Contract *FarmFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// FarmFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FarmFactoryTransactorRaw struct {
	Contract *FarmFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFarmFactory creates a new instance of FarmFactory, bound to a specific deployed contract.
func NewFarmFactory(address common.Address, backend bind.ContractBackend) (*FarmFactory, error) {
	contract, err := bindFarmFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FarmFactory{FarmFactoryCaller: FarmFactoryCaller{contract: contract}, FarmFactoryTransactor: FarmFactoryTransactor{contract: contract}, FarmFactoryFilterer: FarmFactoryFilterer{contract: contract}}, nil
}

// NewFarmFactoryCaller creates a new read-only instance of FarmFactory, bound to a specific deployed contract.
func NewFarmFactoryCaller(address common.Address, caller bind.ContractCaller) (*FarmFactoryCaller, error) {
	contract, err := bindFarmFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FarmFactoryCaller{contract: contract}, nil
}

// NewFarmFactoryTransactor creates a new write-only instance of FarmFactory, bound to a specific deployed contract.
func NewFarmFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*FarmFactoryTransactor, error) {
	contract, err := bindFarmFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FarmFactoryTransactor{contract: contract}, nil
}

// NewFarmFactoryFilterer creates a new log filterer instance of FarmFactory, bound to a specific deployed contract.
func NewFarmFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*FarmFactoryFilterer, error) {
	contract, err := bindFarmFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FarmFactoryFilterer{contract: contract}, nil
}

// bindFarmFactory binds a generic wrapper to an already deployed contract.
func bindFarmFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FarmFactoryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FarmFactory *FarmFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FarmFactory.Contract.FarmFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FarmFactory *FarmFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarmFactory.Contract.FarmFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FarmFactory *FarmFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FarmFactory.Contract.FarmFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FarmFactory *FarmFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FarmFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FarmFactory *FarmFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarmFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FarmFactory *FarmFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FarmFactory.Contract.contract.Transact(opts, method, params...)
}

// DroneFactory is a free data retrieval call binding the contract method 0xc7b26142.
//
// Solidity: function droneFactory() view returns(address)
func (_FarmFactory *FarmFactoryCaller) DroneFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "droneFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DroneFactory is a free data retrieval call binding the contract method 0xc7b26142.
//
// Solidity: function droneFactory() view returns(address)
func (_FarmFactory *FarmFactorySession) DroneFactory() (common.Address, error) {
	return _FarmFactory.Contract.DroneFactory(&_FarmFactory.CallOpts)
}

// DroneFactory is a free data retrieval call binding the contract method 0xc7b26142.
//
// Solidity: function droneFactory() view returns(address)
func (_FarmFactory *FarmFactoryCallerSession) DroneFactory() (common.Address, error) {
	return _FarmFactory.Contract.DroneFactory(&_FarmFactory.CallOpts)
}

// FarmExists is a free data retrieval call binding the contract method 0x3460c114.
//
// Solidity: function farmExists(address farmId) view returns(bool)
func (_FarmFactory *FarmFactoryCaller) FarmExists(opts *bind.CallOpts, farmId common.Address) (bool, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "farmExists", farmId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FarmExists is a free data retrieval call binding the contract method 0x3460c114.
//
// Solidity: function farmExists(address farmId) view returns(bool)
func (_FarmFactory *FarmFactorySession) FarmExists(farmId common.Address) (bool, error) {
	return _FarmFactory.Contract.FarmExists(&_FarmFactory.CallOpts, farmId)
}

// FarmExists is a free data retrieval call binding the contract method 0x3460c114.
//
// Solidity: function farmExists(address farmId) view returns(bool)
func (_FarmFactory *FarmFactoryCallerSession) FarmExists(farmId common.Address) (bool, error) {
	return _FarmFactory.Contract.FarmExists(&_FarmFactory.CallOpts, farmId)
}

// FarmGridCodeToId is a free data retrieval call binding the contract method 0x48ac1571.
//
// Solidity: function farmGridCodeToId(address , string ) view returns(address)
func (_FarmFactory *FarmFactoryCaller) FarmGridCodeToId(opts *bind.CallOpts, arg0 common.Address, arg1 string) (common.Address, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "farmGridCodeToId", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FarmGridCodeToId is a free data retrieval call binding the contract method 0x48ac1571.
//
// Solidity: function farmGridCodeToId(address , string ) view returns(address)
func (_FarmFactory *FarmFactorySession) FarmGridCodeToId(arg0 common.Address, arg1 string) (common.Address, error) {
	return _FarmFactory.Contract.FarmGridCodeToId(&_FarmFactory.CallOpts, arg0, arg1)
}

// FarmGridCodeToId is a free data retrieval call binding the contract method 0x48ac1571.
//
// Solidity: function farmGridCodeToId(address , string ) view returns(address)
func (_FarmFactory *FarmFactoryCallerSession) FarmGridCodeToId(arg0 common.Address, arg1 string) (common.Address, error) {
	return _FarmFactory.Contract.FarmGridCodeToId(&_FarmFactory.CallOpts, arg0, arg1)
}

// FarmGridList is a free data retrieval call binding the contract method 0x7a0e4699.
//
// Solidity: function farmGridList(address , uint256 ) view returns(address)
func (_FarmFactory *FarmFactoryCaller) FarmGridList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "farmGridList", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FarmGridList is a free data retrieval call binding the contract method 0x7a0e4699.
//
// Solidity: function farmGridList(address , uint256 ) view returns(address)
func (_FarmFactory *FarmFactorySession) FarmGridList(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _FarmFactory.Contract.FarmGridList(&_FarmFactory.CallOpts, arg0, arg1)
}

// FarmGridList is a free data retrieval call binding the contract method 0x7a0e4699.
//
// Solidity: function farmGridList(address , uint256 ) view returns(address)
func (_FarmFactory *FarmFactoryCallerSession) FarmGridList(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _FarmFactory.Contract.FarmGridList(&_FarmFactory.CallOpts, arg0, arg1)
}

// Farms is a free data retrieval call binding the contract method 0x421adfa0.
//
// Solidity: function farms(address ) view returns(string name, address owner, uint32 createdTime, bool exists)
func (_FarmFactory *FarmFactoryCaller) Farms(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name        string
	Owner       common.Address
	CreatedTime uint32
	Exists      bool
}, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "farms", arg0)

	outstruct := new(struct {
		Name        string
		Owner       common.Address
		CreatedTime uint32
		Exists      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.CreatedTime = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.Exists = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// Farms is a free data retrieval call binding the contract method 0x421adfa0.
//
// Solidity: function farms(address ) view returns(string name, address owner, uint32 createdTime, bool exists)
func (_FarmFactory *FarmFactorySession) Farms(arg0 common.Address) (struct {
	Name        string
	Owner       common.Address
	CreatedTime uint32
	Exists      bool
}, error) {
	return _FarmFactory.Contract.Farms(&_FarmFactory.CallOpts, arg0)
}

// Farms is a free data retrieval call binding the contract method 0x421adfa0.
//
// Solidity: function farms(address ) view returns(string name, address owner, uint32 createdTime, bool exists)
func (_FarmFactory *FarmFactoryCallerSession) Farms(arg0 common.Address) (struct {
	Name        string
	Owner       common.Address
	CreatedTime uint32
	Exists      bool
}, error) {
	return _FarmFactory.Contract.Farms(&_FarmFactory.CallOpts, arg0)
}

// GetFarm is a free data retrieval call binding the contract method 0x3bfc187c.
//
// Solidity: function getFarm(address farmId) view returns((string,address,uint32,bool))
func (_FarmFactory *FarmFactoryCaller) GetFarm(opts *bind.CallOpts, farmId common.Address) (FarmFactoryFarm, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "getFarm", farmId)

	if err != nil {
		return *new(FarmFactoryFarm), err
	}

	out0 := *abi.ConvertType(out[0], new(FarmFactoryFarm)).(*FarmFactoryFarm)

	return out0, err

}

// GetFarm is a free data retrieval call binding the contract method 0x3bfc187c.
//
// Solidity: function getFarm(address farmId) view returns((string,address,uint32,bool))
func (_FarmFactory *FarmFactorySession) GetFarm(farmId common.Address) (FarmFactoryFarm, error) {
	return _FarmFactory.Contract.GetFarm(&_FarmFactory.CallOpts, farmId)
}

// GetFarm is a free data retrieval call binding the contract method 0x3bfc187c.
//
// Solidity: function getFarm(address farmId) view returns((string,address,uint32,bool))
func (_FarmFactory *FarmFactoryCallerSession) GetFarm(farmId common.Address) (FarmFactoryFarm, error) {
	return _FarmFactory.Contract.GetFarm(&_FarmFactory.CallOpts, farmId)
}

// GetFarmGrids is a free data retrieval call binding the contract method 0x72bc2dc5.
//
// Solidity: function getFarmGrids(address farmId) view returns(address[])
func (_FarmFactory *FarmFactoryCaller) GetFarmGrids(opts *bind.CallOpts, farmId common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "getFarmGrids", farmId)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFarmGrids is a free data retrieval call binding the contract method 0x72bc2dc5.
//
// Solidity: function getFarmGrids(address farmId) view returns(address[])
func (_FarmFactory *FarmFactorySession) GetFarmGrids(farmId common.Address) ([]common.Address, error) {
	return _FarmFactory.Contract.GetFarmGrids(&_FarmFactory.CallOpts, farmId)
}

// GetFarmGrids is a free data retrieval call binding the contract method 0x72bc2dc5.
//
// Solidity: function getFarmGrids(address farmId) view returns(address[])
func (_FarmFactory *FarmFactoryCallerSession) GetFarmGrids(farmId common.Address) ([]common.Address, error) {
	return _FarmFactory.Contract.GetFarmGrids(&_FarmFactory.CallOpts, farmId)
}

// GetGridIdByCode is a free data retrieval call binding the contract method 0xa2cf5b92.
//
// Solidity: function getGridIdByCode(address farmId, string gridCode) view returns(address)
func (_FarmFactory *FarmFactoryCaller) GetGridIdByCode(opts *bind.CallOpts, farmId common.Address, gridCode string) (common.Address, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "getGridIdByCode", farmId, gridCode)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetGridIdByCode is a free data retrieval call binding the contract method 0xa2cf5b92.
//
// Solidity: function getGridIdByCode(address farmId, string gridCode) view returns(address)
func (_FarmFactory *FarmFactorySession) GetGridIdByCode(farmId common.Address, gridCode string) (common.Address, error) {
	return _FarmFactory.Contract.GetGridIdByCode(&_FarmFactory.CallOpts, farmId, gridCode)
}

// GetGridIdByCode is a free data retrieval call binding the contract method 0xa2cf5b92.
//
// Solidity: function getGridIdByCode(address farmId, string gridCode) view returns(address)
func (_FarmFactory *FarmFactoryCallerSession) GetGridIdByCode(farmId common.Address, gridCode string) (common.Address, error) {
	return _FarmFactory.Contract.GetGridIdByCode(&_FarmFactory.CallOpts, farmId, gridCode)
}

// GetGridInfo is a free data retrieval call binding the contract method 0x98c0a3e5.
//
// Solidity: function getGridInfo(address farmId, string gridCode) view returns((string,address,uint8,uint8,uint8,uint32,bool))
func (_FarmFactory *FarmFactoryCaller) GetGridInfo(opts *bind.CallOpts, farmId common.Address, gridCode string) (GridObjectFactoryGridInfo, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "getGridInfo", farmId, gridCode)

	if err != nil {
		return *new(GridObjectFactoryGridInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(GridObjectFactoryGridInfo)).(*GridObjectFactoryGridInfo)

	return out0, err

}

// GetGridInfo is a free data retrieval call binding the contract method 0x98c0a3e5.
//
// Solidity: function getGridInfo(address farmId, string gridCode) view returns((string,address,uint8,uint8,uint8,uint32,bool))
func (_FarmFactory *FarmFactorySession) GetGridInfo(farmId common.Address, gridCode string) (GridObjectFactoryGridInfo, error) {
	return _FarmFactory.Contract.GetGridInfo(&_FarmFactory.CallOpts, farmId, gridCode)
}

// GetGridInfo is a free data retrieval call binding the contract method 0x98c0a3e5.
//
// Solidity: function getGridInfo(address farmId, string gridCode) view returns((string,address,uint8,uint8,uint8,uint32,bool))
func (_FarmFactory *FarmFactoryCallerSession) GetGridInfo(farmId common.Address, gridCode string) (GridObjectFactoryGridInfo, error) {
	return _FarmFactory.Contract.GetGridInfo(&_FarmFactory.CallOpts, farmId, gridCode)
}

// GridExistsInFarm is a free data retrieval call binding the contract method 0x2cf449d0.
//
// Solidity: function gridExistsInFarm(address farmId, string gridCode) view returns(bool)
func (_FarmFactory *FarmFactoryCaller) GridExistsInFarm(opts *bind.CallOpts, farmId common.Address, gridCode string) (bool, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "gridExistsInFarm", farmId, gridCode)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GridExistsInFarm is a free data retrieval call binding the contract method 0x2cf449d0.
//
// Solidity: function gridExistsInFarm(address farmId, string gridCode) view returns(bool)
func (_FarmFactory *FarmFactorySession) GridExistsInFarm(farmId common.Address, gridCode string) (bool, error) {
	return _FarmFactory.Contract.GridExistsInFarm(&_FarmFactory.CallOpts, farmId, gridCode)
}

// GridExistsInFarm is a free data retrieval call binding the contract method 0x2cf449d0.
//
// Solidity: function gridExistsInFarm(address farmId, string gridCode) view returns(bool)
func (_FarmFactory *FarmFactoryCallerSession) GridExistsInFarm(farmId common.Address, gridCode string) (bool, error) {
	return _FarmFactory.Contract.GridExistsInFarm(&_FarmFactory.CallOpts, farmId, gridCode)
}

// GridFactory is a free data retrieval call binding the contract method 0x03a7dcdc.
//
// Solidity: function gridFactory() view returns(address)
func (_FarmFactory *FarmFactoryCaller) GridFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "gridFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GridFactory is a free data retrieval call binding the contract method 0x03a7dcdc.
//
// Solidity: function gridFactory() view returns(address)
func (_FarmFactory *FarmFactorySession) GridFactory() (common.Address, error) {
	return _FarmFactory.Contract.GridFactory(&_FarmFactory.CallOpts)
}

// GridFactory is a free data retrieval call binding the contract method 0x03a7dcdc.
//
// Solidity: function gridFactory() view returns(address)
func (_FarmFactory *FarmFactoryCallerSession) GridFactory() (common.Address, error) {
	return _FarmFactory.Contract.GridFactory(&_FarmFactory.CallOpts)
}

// IsFarmOwner is a free data retrieval call binding the contract method 0xebb1115d.
//
// Solidity: function isFarmOwner(address farmId, address account) view returns(bool)
func (_FarmFactory *FarmFactoryCaller) IsFarmOwner(opts *bind.CallOpts, farmId common.Address, account common.Address) (bool, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "isFarmOwner", farmId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFarmOwner is a free data retrieval call binding the contract method 0xebb1115d.
//
// Solidity: function isFarmOwner(address farmId, address account) view returns(bool)
func (_FarmFactory *FarmFactorySession) IsFarmOwner(farmId common.Address, account common.Address) (bool, error) {
	return _FarmFactory.Contract.IsFarmOwner(&_FarmFactory.CallOpts, farmId, account)
}

// IsFarmOwner is a free data retrieval call binding the contract method 0xebb1115d.
//
// Solidity: function isFarmOwner(address farmId, address account) view returns(bool)
func (_FarmFactory *FarmFactoryCallerSession) IsFarmOwner(farmId common.Address, account common.Address) (bool, error) {
	return _FarmFactory.Contract.IsFarmOwner(&_FarmFactory.CallOpts, farmId, account)
}

// RecordManager is a free data retrieval call binding the contract method 0x7884bc58.
//
// Solidity: function recordManager() view returns(address)
func (_FarmFactory *FarmFactoryCaller) RecordManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "recordManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecordManager is a free data retrieval call binding the contract method 0x7884bc58.
//
// Solidity: function recordManager() view returns(address)
func (_FarmFactory *FarmFactorySession) RecordManager() (common.Address, error) {
	return _FarmFactory.Contract.RecordManager(&_FarmFactory.CallOpts)
}

// RecordManager is a free data retrieval call binding the contract method 0x7884bc58.
//
// Solidity: function recordManager() view returns(address)
func (_FarmFactory *FarmFactoryCallerSession) RecordManager() (common.Address, error) {
	return _FarmFactory.Contract.RecordManager(&_FarmFactory.CallOpts)
}

// UserContract is a free data retrieval call binding the contract method 0x5000596b.
//
// Solidity: function userContract() view returns(address)
func (_FarmFactory *FarmFactoryCaller) UserContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "userContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserContract is a free data retrieval call binding the contract method 0x5000596b.
//
// Solidity: function userContract() view returns(address)
func (_FarmFactory *FarmFactorySession) UserContract() (common.Address, error) {
	return _FarmFactory.Contract.UserContract(&_FarmFactory.CallOpts)
}

// UserContract is a free data retrieval call binding the contract method 0x5000596b.
//
// Solidity: function userContract() view returns(address)
func (_FarmFactory *FarmFactoryCallerSession) UserContract() (common.Address, error) {
	return _FarmFactory.Contract.UserContract(&_FarmFactory.CallOpts)
}

// AddGridMaintenance is a paid mutator transaction binding the contract method 0x76bfa790.
//
// Solidity: function addGridMaintenance(address farmId, string gridCode, uint8 operationTypeCode, string detail, bytes32 evidenceIPFSHash) returns()
func (_FarmFactory *FarmFactoryTransactor) AddGridMaintenance(opts *bind.TransactOpts, farmId common.Address, gridCode string, operationTypeCode uint8, detail string, evidenceIPFSHash [32]byte) (*types.Transaction, error) {
	return _FarmFactory.contract.Transact(opts, "addGridMaintenance", farmId, gridCode, operationTypeCode, detail, evidenceIPFSHash)
}

// AddGridMaintenance is a paid mutator transaction binding the contract method 0x76bfa790.
//
// Solidity: function addGridMaintenance(address farmId, string gridCode, uint8 operationTypeCode, string detail, bytes32 evidenceIPFSHash) returns()
func (_FarmFactory *FarmFactorySession) AddGridMaintenance(farmId common.Address, gridCode string, operationTypeCode uint8, detail string, evidenceIPFSHash [32]byte) (*types.Transaction, error) {
	return _FarmFactory.Contract.AddGridMaintenance(&_FarmFactory.TransactOpts, farmId, gridCode, operationTypeCode, detail, evidenceIPFSHash)
}

// AddGridMaintenance is a paid mutator transaction binding the contract method 0x76bfa790.
//
// Solidity: function addGridMaintenance(address farmId, string gridCode, uint8 operationTypeCode, string detail, bytes32 evidenceIPFSHash) returns()
func (_FarmFactory *FarmFactoryTransactorSession) AddGridMaintenance(farmId common.Address, gridCode string, operationTypeCode uint8, detail string, evidenceIPFSHash [32]byte) (*types.Transaction, error) {
	return _FarmFactory.Contract.AddGridMaintenance(&_FarmFactory.TransactOpts, farmId, gridCode, operationTypeCode, detail, evidenceIPFSHash)
}

// AddGridToFarm is a paid mutator transaction binding the contract method 0x1c239d09.
//
// Solidity: function addGridToFarm(address farmId, address gridId, string gridCode, uint8 terrainType) returns()
func (_FarmFactory *FarmFactoryTransactor) AddGridToFarm(opts *bind.TransactOpts, farmId common.Address, gridId common.Address, gridCode string, terrainType uint8) (*types.Transaction, error) {
	return _FarmFactory.contract.Transact(opts, "addGridToFarm", farmId, gridId, gridCode, terrainType)
}

// AddGridToFarm is a paid mutator transaction binding the contract method 0x1c239d09.
//
// Solidity: function addGridToFarm(address farmId, address gridId, string gridCode, uint8 terrainType) returns()
func (_FarmFactory *FarmFactorySession) AddGridToFarm(farmId common.Address, gridId common.Address, gridCode string, terrainType uint8) (*types.Transaction, error) {
	return _FarmFactory.Contract.AddGridToFarm(&_FarmFactory.TransactOpts, farmId, gridId, gridCode, terrainType)
}

// AddGridToFarm is a paid mutator transaction binding the contract method 0x1c239d09.
//
// Solidity: function addGridToFarm(address farmId, address gridId, string gridCode, uint8 terrainType) returns()
func (_FarmFactory *FarmFactoryTransactorSession) AddGridToFarm(farmId common.Address, gridId common.Address, gridCode string, terrainType uint8) (*types.Transaction, error) {
	return _FarmFactory.Contract.AddGridToFarm(&_FarmFactory.TransactOpts, farmId, gridId, gridCode, terrainType)
}

// CreateFarm is a paid mutator transaction binding the contract method 0x8b2809e6.
//
// Solidity: function createFarm(address farmId, string name) returns()
func (_FarmFactory *FarmFactoryTransactor) CreateFarm(opts *bind.TransactOpts, farmId common.Address, name string) (*types.Transaction, error) {
	return _FarmFactory.contract.Transact(opts, "createFarm", farmId, name)
}

// CreateFarm is a paid mutator transaction binding the contract method 0x8b2809e6.
//
// Solidity: function createFarm(address farmId, string name) returns()
func (_FarmFactory *FarmFactorySession) CreateFarm(farmId common.Address, name string) (*types.Transaction, error) {
	return _FarmFactory.Contract.CreateFarm(&_FarmFactory.TransactOpts, farmId, name)
}

// CreateFarm is a paid mutator transaction binding the contract method 0x8b2809e6.
//
// Solidity: function createFarm(address farmId, string name) returns()
func (_FarmFactory *FarmFactoryTransactorSession) CreateFarm(farmId common.Address, name string) (*types.Transaction, error) {
	return _FarmFactory.Contract.CreateFarm(&_FarmFactory.TransactOpts, farmId, name)
}

// SetRecordManager is a paid mutator transaction binding the contract method 0x874b00d5.
//
// Solidity: function setRecordManager(address _recordManager) returns()
func (_FarmFactory *FarmFactoryTransactor) SetRecordManager(opts *bind.TransactOpts, _recordManager common.Address) (*types.Transaction, error) {
	return _FarmFactory.contract.Transact(opts, "setRecordManager", _recordManager)
}

// SetRecordManager is a paid mutator transaction binding the contract method 0x874b00d5.
//
// Solidity: function setRecordManager(address _recordManager) returns()
func (_FarmFactory *FarmFactorySession) SetRecordManager(_recordManager common.Address) (*types.Transaction, error) {
	return _FarmFactory.Contract.SetRecordManager(&_FarmFactory.TransactOpts, _recordManager)
}

// SetRecordManager is a paid mutator transaction binding the contract method 0x874b00d5.
//
// Solidity: function setRecordManager(address _recordManager) returns()
func (_FarmFactory *FarmFactoryTransactorSession) SetRecordManager(_recordManager common.Address) (*types.Transaction, error) {
	return _FarmFactory.Contract.SetRecordManager(&_FarmFactory.TransactOpts, _recordManager)
}

// UpdateGridDiseaseType is a paid mutator transaction binding the contract method 0x2c5951a6.
//
// Solidity: function updateGridDiseaseType(address farmId, string gridCode, uint8 newType) returns()
func (_FarmFactory *FarmFactoryTransactor) UpdateGridDiseaseType(opts *bind.TransactOpts, farmId common.Address, gridCode string, newType uint8) (*types.Transaction, error) {
	return _FarmFactory.contract.Transact(opts, "updateGridDiseaseType", farmId, gridCode, newType)
}

// UpdateGridDiseaseType is a paid mutator transaction binding the contract method 0x2c5951a6.
//
// Solidity: function updateGridDiseaseType(address farmId, string gridCode, uint8 newType) returns()
func (_FarmFactory *FarmFactorySession) UpdateGridDiseaseType(farmId common.Address, gridCode string, newType uint8) (*types.Transaction, error) {
	return _FarmFactory.Contract.UpdateGridDiseaseType(&_FarmFactory.TransactOpts, farmId, gridCode, newType)
}

// UpdateGridDiseaseType is a paid mutator transaction binding the contract method 0x2c5951a6.
//
// Solidity: function updateGridDiseaseType(address farmId, string gridCode, uint8 newType) returns()
func (_FarmFactory *FarmFactoryTransactorSession) UpdateGridDiseaseType(farmId common.Address, gridCode string, newType uint8) (*types.Transaction, error) {
	return _FarmFactory.Contract.UpdateGridDiseaseType(&_FarmFactory.TransactOpts, farmId, gridCode, newType)
}

// UpdateGridDiseaseTypeFromRecord is a paid mutator transaction binding the contract method 0x09c76654.
//
// Solidity: function updateGridDiseaseTypeFromRecord(address farmId, string gridCode, uint8 newType) returns()
func (_FarmFactory *FarmFactoryTransactor) UpdateGridDiseaseTypeFromRecord(opts *bind.TransactOpts, farmId common.Address, gridCode string, newType uint8) (*types.Transaction, error) {
	return _FarmFactory.contract.Transact(opts, "updateGridDiseaseTypeFromRecord", farmId, gridCode, newType)
}

// UpdateGridDiseaseTypeFromRecord is a paid mutator transaction binding the contract method 0x09c76654.
//
// Solidity: function updateGridDiseaseTypeFromRecord(address farmId, string gridCode, uint8 newType) returns()
func (_FarmFactory *FarmFactorySession) UpdateGridDiseaseTypeFromRecord(farmId common.Address, gridCode string, newType uint8) (*types.Transaction, error) {
	return _FarmFactory.Contract.UpdateGridDiseaseTypeFromRecord(&_FarmFactory.TransactOpts, farmId, gridCode, newType)
}

// UpdateGridDiseaseTypeFromRecord is a paid mutator transaction binding the contract method 0x09c76654.
//
// Solidity: function updateGridDiseaseTypeFromRecord(address farmId, string gridCode, uint8 newType) returns()
func (_FarmFactory *FarmFactoryTransactorSession) UpdateGridDiseaseTypeFromRecord(farmId common.Address, gridCode string, newType uint8) (*types.Transaction, error) {
	return _FarmFactory.Contract.UpdateGridDiseaseTypeFromRecord(&_FarmFactory.TransactOpts, farmId, gridCode, newType)
}

// UpdateGridStatus is a paid mutator transaction binding the contract method 0x6b02fb2e.
//
// Solidity: function updateGridStatus(address farmId, string gridCode, uint8 newStatus) returns()
func (_FarmFactory *FarmFactoryTransactor) UpdateGridStatus(opts *bind.TransactOpts, farmId common.Address, gridCode string, newStatus uint8) (*types.Transaction, error) {
	return _FarmFactory.contract.Transact(opts, "updateGridStatus", farmId, gridCode, newStatus)
}

// UpdateGridStatus is a paid mutator transaction binding the contract method 0x6b02fb2e.
//
// Solidity: function updateGridStatus(address farmId, string gridCode, uint8 newStatus) returns()
func (_FarmFactory *FarmFactorySession) UpdateGridStatus(farmId common.Address, gridCode string, newStatus uint8) (*types.Transaction, error) {
	return _FarmFactory.Contract.UpdateGridStatus(&_FarmFactory.TransactOpts, farmId, gridCode, newStatus)
}

// UpdateGridStatus is a paid mutator transaction binding the contract method 0x6b02fb2e.
//
// Solidity: function updateGridStatus(address farmId, string gridCode, uint8 newStatus) returns()
func (_FarmFactory *FarmFactoryTransactorSession) UpdateGridStatus(farmId common.Address, gridCode string, newStatus uint8) (*types.Transaction, error) {
	return _FarmFactory.Contract.UpdateGridStatus(&_FarmFactory.TransactOpts, farmId, gridCode, newStatus)
}

// UpdateGridStatusFromRecord is a paid mutator transaction binding the contract method 0x38ff2222.
//
// Solidity: function updateGridStatusFromRecord(address farmId, string gridCode, uint8 newStatus) returns()
func (_FarmFactory *FarmFactoryTransactor) UpdateGridStatusFromRecord(opts *bind.TransactOpts, farmId common.Address, gridCode string, newStatus uint8) (*types.Transaction, error) {
	return _FarmFactory.contract.Transact(opts, "updateGridStatusFromRecord", farmId, gridCode, newStatus)
}

// UpdateGridStatusFromRecord is a paid mutator transaction binding the contract method 0x38ff2222.
//
// Solidity: function updateGridStatusFromRecord(address farmId, string gridCode, uint8 newStatus) returns()
func (_FarmFactory *FarmFactorySession) UpdateGridStatusFromRecord(farmId common.Address, gridCode string, newStatus uint8) (*types.Transaction, error) {
	return _FarmFactory.Contract.UpdateGridStatusFromRecord(&_FarmFactory.TransactOpts, farmId, gridCode, newStatus)
}

// UpdateGridStatusFromRecord is a paid mutator transaction binding the contract method 0x38ff2222.
//
// Solidity: function updateGridStatusFromRecord(address farmId, string gridCode, uint8 newStatus) returns()
func (_FarmFactory *FarmFactoryTransactorSession) UpdateGridStatusFromRecord(farmId common.Address, gridCode string, newStatus uint8) (*types.Transaction, error) {
	return _FarmFactory.Contract.UpdateGridStatusFromRecord(&_FarmFactory.TransactOpts, farmId, gridCode, newStatus)
}

// FarmFactoryFarmCreatedIterator is returned from FilterFarmCreated and is used to iterate over the raw logs and unpacked data for FarmCreated events raised by the FarmFactory contract.
type FarmFactoryFarmCreatedIterator struct {
	Event *FarmFactoryFarmCreated // Event containing the contract specifics and raw log

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
func (it *FarmFactoryFarmCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmFactoryFarmCreated)
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
		it.Event = new(FarmFactoryFarmCreated)
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
func (it *FarmFactoryFarmCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmFactoryFarmCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmFactoryFarmCreated represents a FarmCreated event raised by the FarmFactory contract.
type FarmFactoryFarmCreated struct {
	FarmId common.Address
	Name   string
	Owner  common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFarmCreated is a free log retrieval operation binding the contract event 0xcfb9790749b5d99475809cf22fdf8e7b1992ea93e7bd0513964e36ef06a1e60d.
//
// Solidity: event FarmCreated(address indexed farmId, string name, address indexed owner)
func (_FarmFactory *FarmFactoryFilterer) FilterFarmCreated(opts *bind.FilterOpts, farmId []common.Address, owner []common.Address) (*FarmFactoryFarmCreatedIterator, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FarmFactory.contract.FilterLogs(opts, "FarmCreated", farmIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &FarmFactoryFarmCreatedIterator{contract: _FarmFactory.contract, event: "FarmCreated", logs: logs, sub: sub}, nil
}

// WatchFarmCreated is a free log subscription operation binding the contract event 0xcfb9790749b5d99475809cf22fdf8e7b1992ea93e7bd0513964e36ef06a1e60d.
//
// Solidity: event FarmCreated(address indexed farmId, string name, address indexed owner)
func (_FarmFactory *FarmFactoryFilterer) WatchFarmCreated(opts *bind.WatchOpts, sink chan<- *FarmFactoryFarmCreated, farmId []common.Address, owner []common.Address) (event.Subscription, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FarmFactory.contract.WatchLogs(opts, "FarmCreated", farmIdRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmFactoryFarmCreated)
				if err := _FarmFactory.contract.UnpackLog(event, "FarmCreated", log); err != nil {
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

// ParseFarmCreated is a log parse operation binding the contract event 0xcfb9790749b5d99475809cf22fdf8e7b1992ea93e7bd0513964e36ef06a1e60d.
//
// Solidity: event FarmCreated(address indexed farmId, string name, address indexed owner)
func (_FarmFactory *FarmFactoryFilterer) ParseFarmCreated(log types.Log) (*FarmFactoryFarmCreated, error) {
	event := new(FarmFactoryFarmCreated)
	if err := _FarmFactory.contract.UnpackLog(event, "FarmCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmFactoryGridAddedToFarmIterator is returned from FilterGridAddedToFarm and is used to iterate over the raw logs and unpacked data for GridAddedToFarm events raised by the FarmFactory contract.
type FarmFactoryGridAddedToFarmIterator struct {
	Event *FarmFactoryGridAddedToFarm // Event containing the contract specifics and raw log

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
func (it *FarmFactoryGridAddedToFarmIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmFactoryGridAddedToFarm)
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
		it.Event = new(FarmFactoryGridAddedToFarm)
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
func (it *FarmFactoryGridAddedToFarmIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmFactoryGridAddedToFarmIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmFactoryGridAddedToFarm represents a GridAddedToFarm event raised by the FarmFactory contract.
type FarmFactoryGridAddedToFarm struct {
	FarmId   common.Address
	GridId   common.Address
	GridCode string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGridAddedToFarm is a free log retrieval operation binding the contract event 0x0e9a1d62fce5c0e7e933ff9e079fa2f4176aba4ced923b8f70c643f643a0e53e.
//
// Solidity: event GridAddedToFarm(address indexed farmId, address indexed gridId, string gridCode)
func (_FarmFactory *FarmFactoryFilterer) FilterGridAddedToFarm(opts *bind.FilterOpts, farmId []common.Address, gridId []common.Address) (*FarmFactoryGridAddedToFarmIterator, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}
	var gridIdRule []interface{}
	for _, gridIdItem := range gridId {
		gridIdRule = append(gridIdRule, gridIdItem)
	}

	logs, sub, err := _FarmFactory.contract.FilterLogs(opts, "GridAddedToFarm", farmIdRule, gridIdRule)
	if err != nil {
		return nil, err
	}
	return &FarmFactoryGridAddedToFarmIterator{contract: _FarmFactory.contract, event: "GridAddedToFarm", logs: logs, sub: sub}, nil
}

// WatchGridAddedToFarm is a free log subscription operation binding the contract event 0x0e9a1d62fce5c0e7e933ff9e079fa2f4176aba4ced923b8f70c643f643a0e53e.
//
// Solidity: event GridAddedToFarm(address indexed farmId, address indexed gridId, string gridCode)
func (_FarmFactory *FarmFactoryFilterer) WatchGridAddedToFarm(opts *bind.WatchOpts, sink chan<- *FarmFactoryGridAddedToFarm, farmId []common.Address, gridId []common.Address) (event.Subscription, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}
	var gridIdRule []interface{}
	for _, gridIdItem := range gridId {
		gridIdRule = append(gridIdRule, gridIdItem)
	}

	logs, sub, err := _FarmFactory.contract.WatchLogs(opts, "GridAddedToFarm", farmIdRule, gridIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmFactoryGridAddedToFarm)
				if err := _FarmFactory.contract.UnpackLog(event, "GridAddedToFarm", log); err != nil {
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

// ParseGridAddedToFarm is a log parse operation binding the contract event 0x0e9a1d62fce5c0e7e933ff9e079fa2f4176aba4ced923b8f70c643f643a0e53e.
//
// Solidity: event GridAddedToFarm(address indexed farmId, address indexed gridId, string gridCode)
func (_FarmFactory *FarmFactoryFilterer) ParseGridAddedToFarm(log types.Log) (*FarmFactoryGridAddedToFarm, error) {
	event := new(FarmFactoryGridAddedToFarm)
	if err := _FarmFactory.contract.UnpackLog(event, "GridAddedToFarm", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmFactoryRecordManagerSetIterator is returned from FilterRecordManagerSet and is used to iterate over the raw logs and unpacked data for RecordManagerSet events raised by the FarmFactory contract.
type FarmFactoryRecordManagerSetIterator struct {
	Event *FarmFactoryRecordManagerSet // Event containing the contract specifics and raw log

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
func (it *FarmFactoryRecordManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmFactoryRecordManagerSet)
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
		it.Event = new(FarmFactoryRecordManagerSet)
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
func (it *FarmFactoryRecordManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmFactoryRecordManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmFactoryRecordManagerSet represents a RecordManagerSet event raised by the FarmFactory contract.
type FarmFactoryRecordManagerSet struct {
	RecordManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRecordManagerSet is a free log retrieval operation binding the contract event 0x6084bd037a8f41ccdc5ab463ecf8ffad957713993d48659aa5947b245ade8070.
//
// Solidity: event RecordManagerSet(address indexed recordManager)
func (_FarmFactory *FarmFactoryFilterer) FilterRecordManagerSet(opts *bind.FilterOpts, recordManager []common.Address) (*FarmFactoryRecordManagerSetIterator, error) {

	var recordManagerRule []interface{}
	for _, recordManagerItem := range recordManager {
		recordManagerRule = append(recordManagerRule, recordManagerItem)
	}

	logs, sub, err := _FarmFactory.contract.FilterLogs(opts, "RecordManagerSet", recordManagerRule)
	if err != nil {
		return nil, err
	}
	return &FarmFactoryRecordManagerSetIterator{contract: _FarmFactory.contract, event: "RecordManagerSet", logs: logs, sub: sub}, nil
}

// WatchRecordManagerSet is a free log subscription operation binding the contract event 0x6084bd037a8f41ccdc5ab463ecf8ffad957713993d48659aa5947b245ade8070.
//
// Solidity: event RecordManagerSet(address indexed recordManager)
func (_FarmFactory *FarmFactoryFilterer) WatchRecordManagerSet(opts *bind.WatchOpts, sink chan<- *FarmFactoryRecordManagerSet, recordManager []common.Address) (event.Subscription, error) {

	var recordManagerRule []interface{}
	for _, recordManagerItem := range recordManager {
		recordManagerRule = append(recordManagerRule, recordManagerItem)
	}

	logs, sub, err := _FarmFactory.contract.WatchLogs(opts, "RecordManagerSet", recordManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmFactoryRecordManagerSet)
				if err := _FarmFactory.contract.UnpackLog(event, "RecordManagerSet", log); err != nil {
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

// ParseRecordManagerSet is a log parse operation binding the contract event 0x6084bd037a8f41ccdc5ab463ecf8ffad957713993d48659aa5947b245ade8070.
//
// Solidity: event RecordManagerSet(address indexed recordManager)
func (_FarmFactory *FarmFactoryFilterer) ParseRecordManagerSet(log types.Log) (*FarmFactoryRecordManagerSet, error) {
	event := new(FarmFactoryRecordManagerSet)
	if err := _FarmFactory.contract.UnpackLog(event, "RecordManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
