// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package User

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

// UserUserInfo is an auto generated low-level Go binding around an user-defined struct.
type UserUserInfo struct {
	Name           string
	Role           uint8
	ContactInfo    string
	RegisteredTime uint32
	Exists         bool
}

// UserMetaData contains all meta data concerning the User contract.
var UserMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"removedAdmin\",\"type\":\"address\"}],\"name\":\"AdminRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"enumUser.UserRole\",\"name\":\"role\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"name\":\"UserRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"userId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumUser.UserRole\",\"name\":\"newRole\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"}],\"name\":\"UserRoleUpdated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"adminCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userId\",\"type\":\"address\"}],\"name\":\"getUserInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"enumUser.UserRole\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"contactInfo\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"registeredTime\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"internalType\":\"structUser.UserInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userId\",\"type\":\"address\"}],\"name\":\"getUserRole\",\"outputs\":[{\"internalType\":\"enumUser.UserRole\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"enumUser.UserRole\",\"name\":\"_role\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"_contactInfo\",\"type\":\"string\"}],\"name\":\"registerUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userId\",\"type\":\"address\"},{\"internalType\":\"enumUser.UserRole\",\"name\":\"_newRole\",\"type\":\"uint8\"}],\"name\":\"updateUserRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_userId\",\"type\":\"address\"}],\"name\":\"userExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"users\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"enumUser.UserRole\",\"name\":\"role\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"contactInfo\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"registeredTime\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// UserABI is the input ABI used to generate the binding from.
// Deprecated: Use UserMetaData.ABI instead.
var UserABI = UserMetaData.ABI

// User is an auto generated Go binding around an Ethereum contract.
type User struct {
	UserCaller     // Read-only binding to the contract
	UserTransactor // Write-only binding to the contract
	UserFilterer   // Log filterer for contract events
}

// UserCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserSession struct {
	Contract     *User             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserCallerSession struct {
	Contract *UserCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// UserTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserTransactorSession struct {
	Contract     *UserTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UserRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserRaw struct {
	Contract *User // Generic contract binding to access the raw methods on
}

// UserCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserCallerRaw struct {
	Contract *UserCaller // Generic read-only contract binding to access the raw methods on
}

// UserTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserTransactorRaw struct {
	Contract *UserTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUser creates a new instance of User, bound to a specific deployed contract.
func NewUser(address common.Address, backend bind.ContractBackend) (*User, error) {
	contract, err := bindUser(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &User{UserCaller: UserCaller{contract: contract}, UserTransactor: UserTransactor{contract: contract}, UserFilterer: UserFilterer{contract: contract}}, nil
}

// NewUserCaller creates a new read-only instance of User, bound to a specific deployed contract.
func NewUserCaller(address common.Address, caller bind.ContractCaller) (*UserCaller, error) {
	contract, err := bindUser(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserCaller{contract: contract}, nil
}

// NewUserTransactor creates a new write-only instance of User, bound to a specific deployed contract.
func NewUserTransactor(address common.Address, transactor bind.ContractTransactor) (*UserTransactor, error) {
	contract, err := bindUser(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserTransactor{contract: contract}, nil
}

// NewUserFilterer creates a new log filterer instance of User, bound to a specific deployed contract.
func NewUserFilterer(address common.Address, filterer bind.ContractFilterer) (*UserFilterer, error) {
	contract, err := bindUser(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserFilterer{contract: contract}, nil
}

// bindUser binds a generic wrapper to an already deployed contract.
func bindUser(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := UserMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _User.Contract.UserCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.UserTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_User *UserCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _User.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_User *UserTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _User.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_User *UserTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _User.Contract.contract.Transact(opts, method, params...)
}

// AdminCount is a free data retrieval call binding the contract method 0x2b7832b3.
//
// Solidity: function adminCount() view returns(uint256)
func (_User *UserCaller) AdminCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "adminCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AdminCount is a free data retrieval call binding the contract method 0x2b7832b3.
//
// Solidity: function adminCount() view returns(uint256)
func (_User *UserSession) AdminCount() (*big.Int, error) {
	return _User.Contract.AdminCount(&_User.CallOpts)
}

// AdminCount is a free data retrieval call binding the contract method 0x2b7832b3.
//
// Solidity: function adminCount() view returns(uint256)
func (_User *UserCallerSession) AdminCount() (*big.Int, error) {
	return _User.Contract.AdminCount(&_User.CallOpts)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address _userId) view returns((string,uint8,string,uint32,bool))
func (_User *UserCaller) GetUserInfo(opts *bind.CallOpts, _userId common.Address) (UserUserInfo, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "getUserInfo", _userId)

	if err != nil {
		return *new(UserUserInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(UserUserInfo)).(*UserUserInfo)

	return out0, err

}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address _userId) view returns((string,uint8,string,uint32,bool))
func (_User *UserSession) GetUserInfo(_userId common.Address) (UserUserInfo, error) {
	return _User.Contract.GetUserInfo(&_User.CallOpts, _userId)
}

// GetUserInfo is a free data retrieval call binding the contract method 0x6386c1c7.
//
// Solidity: function getUserInfo(address _userId) view returns((string,uint8,string,uint32,bool))
func (_User *UserCallerSession) GetUserInfo(_userId common.Address) (UserUserInfo, error) {
	return _User.Contract.GetUserInfo(&_User.CallOpts, _userId)
}

// GetUserRole is a free data retrieval call binding the contract method 0x27820851.
//
// Solidity: function getUserRole(address _userId) view returns(uint8)
func (_User *UserCaller) GetUserRole(opts *bind.CallOpts, _userId common.Address) (uint8, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "getUserRole", _userId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetUserRole is a free data retrieval call binding the contract method 0x27820851.
//
// Solidity: function getUserRole(address _userId) view returns(uint8)
func (_User *UserSession) GetUserRole(_userId common.Address) (uint8, error) {
	return _User.Contract.GetUserRole(&_User.CallOpts, _userId)
}

// GetUserRole is a free data retrieval call binding the contract method 0x27820851.
//
// Solidity: function getUserRole(address _userId) view returns(uint8)
func (_User *UserCallerSession) GetUserRole(_userId common.Address) (uint8, error) {
	return _User.Contract.GetUserRole(&_User.CallOpts, _userId)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_User *UserCaller) IsAdmin(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "isAdmin", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_User *UserSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _User.Contract.IsAdmin(&_User.CallOpts, arg0)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address ) view returns(bool)
func (_User *UserCallerSession) IsAdmin(arg0 common.Address) (bool, error) {
	return _User.Contract.IsAdmin(&_User.CallOpts, arg0)
}

// UserExists is a free data retrieval call binding the contract method 0x0e666e49.
//
// Solidity: function userExists(address _userId) view returns(bool)
func (_User *UserCaller) UserExists(opts *bind.CallOpts, _userId common.Address) (bool, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "userExists", _userId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UserExists is a free data retrieval call binding the contract method 0x0e666e49.
//
// Solidity: function userExists(address _userId) view returns(bool)
func (_User *UserSession) UserExists(_userId common.Address) (bool, error) {
	return _User.Contract.UserExists(&_User.CallOpts, _userId)
}

// UserExists is a free data retrieval call binding the contract method 0x0e666e49.
//
// Solidity: function userExists(address _userId) view returns(bool)
func (_User *UserCallerSession) UserExists(_userId common.Address) (bool, error) {
	return _User.Contract.UserExists(&_User.CallOpts, _userId)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(string name, uint8 role, string contactInfo, uint32 registeredTime, bool exists)
func (_User *UserCaller) Users(opts *bind.CallOpts, arg0 common.Address) (struct {
	Name           string
	Role           uint8
	ContactInfo    string
	RegisteredTime uint32
	Exists         bool
}, error) {
	var out []interface{}
	err := _User.contract.Call(opts, &out, "users", arg0)

	outstruct := new(struct {
		Name           string
		Role           uint8
		ContactInfo    string
		RegisteredTime uint32
		Exists         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Role = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.ContactInfo = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.RegisteredTime = *abi.ConvertType(out[3], new(uint32)).(*uint32)
	outstruct.Exists = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(string name, uint8 role, string contactInfo, uint32 registeredTime, bool exists)
func (_User *UserSession) Users(arg0 common.Address) (struct {
	Name           string
	Role           uint8
	ContactInfo    string
	RegisteredTime uint32
	Exists         bool
}, error) {
	return _User.Contract.Users(&_User.CallOpts, arg0)
}

// Users is a free data retrieval call binding the contract method 0xa87430ba.
//
// Solidity: function users(address ) view returns(string name, uint8 role, string contactInfo, uint32 registeredTime, bool exists)
func (_User *UserCallerSession) Users(arg0 common.Address) (struct {
	Name           string
	Role           uint8
	ContactInfo    string
	RegisteredTime uint32
	Exists         bool
}, error) {
	return _User.Contract.Users(&_User.CallOpts, arg0)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x3d28a301.
//
// Solidity: function registerUser(address _userId, string _name, uint8 _role, string _contactInfo) returns()
func (_User *UserTransactor) RegisterUser(opts *bind.TransactOpts, _userId common.Address, _name string, _role uint8, _contactInfo string) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "registerUser", _userId, _name, _role, _contactInfo)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x3d28a301.
//
// Solidity: function registerUser(address _userId, string _name, uint8 _role, string _contactInfo) returns()
func (_User *UserSession) RegisterUser(_userId common.Address, _name string, _role uint8, _contactInfo string) (*types.Transaction, error) {
	return _User.Contract.RegisterUser(&_User.TransactOpts, _userId, _name, _role, _contactInfo)
}

// RegisterUser is a paid mutator transaction binding the contract method 0x3d28a301.
//
// Solidity: function registerUser(address _userId, string _name, uint8 _role, string _contactInfo) returns()
func (_User *UserTransactorSession) RegisterUser(_userId common.Address, _name string, _role uint8, _contactInfo string) (*types.Transaction, error) {
	return _User.Contract.RegisterUser(&_User.TransactOpts, _userId, _name, _role, _contactInfo)
}

// UpdateUserRole is a paid mutator transaction binding the contract method 0xc291fe62.
//
// Solidity: function updateUserRole(address _userId, uint8 _newRole) returns()
func (_User *UserTransactor) UpdateUserRole(opts *bind.TransactOpts, _userId common.Address, _newRole uint8) (*types.Transaction, error) {
	return _User.contract.Transact(opts, "updateUserRole", _userId, _newRole)
}

// UpdateUserRole is a paid mutator transaction binding the contract method 0xc291fe62.
//
// Solidity: function updateUserRole(address _userId, uint8 _newRole) returns()
func (_User *UserSession) UpdateUserRole(_userId common.Address, _newRole uint8) (*types.Transaction, error) {
	return _User.Contract.UpdateUserRole(&_User.TransactOpts, _userId, _newRole)
}

// UpdateUserRole is a paid mutator transaction binding the contract method 0xc291fe62.
//
// Solidity: function updateUserRole(address _userId, uint8 _newRole) returns()
func (_User *UserTransactorSession) UpdateUserRole(_userId common.Address, _newRole uint8) (*types.Transaction, error) {
	return _User.Contract.UpdateUserRole(&_User.TransactOpts, _userId, _newRole)
}

// UserAdminAddedIterator is returned from FilterAdminAdded and is used to iterate over the raw logs and unpacked data for AdminAdded events raised by the User contract.
type UserAdminAddedIterator struct {
	Event *UserAdminAdded // Event containing the contract specifics and raw log

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
func (it *UserAdminAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserAdminAdded)
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
		it.Event = new(UserAdminAdded)
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
func (it *UserAdminAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserAdminAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserAdminAdded represents a AdminAdded event raised by the User contract.
type UserAdminAdded struct {
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminAdded is a free log retrieval operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed newAdmin)
func (_User *UserFilterer) FilterAdminAdded(opts *bind.FilterOpts, newAdmin []common.Address) (*UserAdminAddedIterator, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _User.contract.FilterLogs(opts, "AdminAdded", newAdminRule)
	if err != nil {
		return nil, err
	}
	return &UserAdminAddedIterator{contract: _User.contract, event: "AdminAdded", logs: logs, sub: sub}, nil
}

// WatchAdminAdded is a free log subscription operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed newAdmin)
func (_User *UserFilterer) WatchAdminAdded(opts *bind.WatchOpts, sink chan<- *UserAdminAdded, newAdmin []common.Address) (event.Subscription, error) {

	var newAdminRule []interface{}
	for _, newAdminItem := range newAdmin {
		newAdminRule = append(newAdminRule, newAdminItem)
	}

	logs, sub, err := _User.contract.WatchLogs(opts, "AdminAdded", newAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserAdminAdded)
				if err := _User.contract.UnpackLog(event, "AdminAdded", log); err != nil {
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

// ParseAdminAdded is a log parse operation binding the contract event 0x44d6d25963f097ad14f29f06854a01f575648a1ef82f30e562ccd3889717e339.
//
// Solidity: event AdminAdded(address indexed newAdmin)
func (_User *UserFilterer) ParseAdminAdded(log types.Log) (*UserAdminAdded, error) {
	event := new(UserAdminAdded)
	if err := _User.contract.UnpackLog(event, "AdminAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserAdminRemovedIterator is returned from FilterAdminRemoved and is used to iterate over the raw logs and unpacked data for AdminRemoved events raised by the User contract.
type UserAdminRemovedIterator struct {
	Event *UserAdminRemoved // Event containing the contract specifics and raw log

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
func (it *UserAdminRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserAdminRemoved)
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
		it.Event = new(UserAdminRemoved)
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
func (it *UserAdminRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserAdminRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserAdminRemoved represents a AdminRemoved event raised by the User contract.
type UserAdminRemoved struct {
	RemovedAdmin common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterAdminRemoved is a free log retrieval operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed removedAdmin)
func (_User *UserFilterer) FilterAdminRemoved(opts *bind.FilterOpts, removedAdmin []common.Address) (*UserAdminRemovedIterator, error) {

	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}

	logs, sub, err := _User.contract.FilterLogs(opts, "AdminRemoved", removedAdminRule)
	if err != nil {
		return nil, err
	}
	return &UserAdminRemovedIterator{contract: _User.contract, event: "AdminRemoved", logs: logs, sub: sub}, nil
}

// WatchAdminRemoved is a free log subscription operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed removedAdmin)
func (_User *UserFilterer) WatchAdminRemoved(opts *bind.WatchOpts, sink chan<- *UserAdminRemoved, removedAdmin []common.Address) (event.Subscription, error) {

	var removedAdminRule []interface{}
	for _, removedAdminItem := range removedAdmin {
		removedAdminRule = append(removedAdminRule, removedAdminItem)
	}

	logs, sub, err := _User.contract.WatchLogs(opts, "AdminRemoved", removedAdminRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserAdminRemoved)
				if err := _User.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
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

// ParseAdminRemoved is a log parse operation binding the contract event 0xa3b62bc36326052d97ea62d63c3d60308ed4c3ea8ac079dd8499f1e9c4f80c0f.
//
// Solidity: event AdminRemoved(address indexed removedAdmin)
func (_User *UserFilterer) ParseAdminRemoved(log types.Log) (*UserAdminRemoved, error) {
	event := new(UserAdminRemoved)
	if err := _User.contract.UnpackLog(event, "AdminRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserUserRegisteredIterator is returned from FilterUserRegistered and is used to iterate over the raw logs and unpacked data for UserRegistered events raised by the User contract.
type UserUserRegisteredIterator struct {
	Event *UserUserRegistered // Event containing the contract specifics and raw log

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
func (it *UserUserRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserUserRegistered)
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
		it.Event = new(UserUserRegistered)
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
func (it *UserUserRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserUserRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserUserRegistered represents a UserRegistered event raised by the User contract.
type UserUserRegistered struct {
	UserId    common.Address
	Name      string
	Role      uint8
	Timestamp uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUserRegistered is a free log retrieval operation binding the contract event 0x34f38c02d99abe0fa9a0ff44d588516965741625980bc7663e6405d8afa8df40.
//
// Solidity: event UserRegistered(address indexed userId, string name, uint8 indexed role, uint32 timestamp)
func (_User *UserFilterer) FilterUserRegistered(opts *bind.FilterOpts, userId []common.Address, role []uint8) (*UserUserRegisteredIterator, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _User.contract.FilterLogs(opts, "UserRegistered", userIdRule, roleRule)
	if err != nil {
		return nil, err
	}
	return &UserUserRegisteredIterator{contract: _User.contract, event: "UserRegistered", logs: logs, sub: sub}, nil
}

// WatchUserRegistered is a free log subscription operation binding the contract event 0x34f38c02d99abe0fa9a0ff44d588516965741625980bc7663e6405d8afa8df40.
//
// Solidity: event UserRegistered(address indexed userId, string name, uint8 indexed role, uint32 timestamp)
func (_User *UserFilterer) WatchUserRegistered(opts *bind.WatchOpts, sink chan<- *UserUserRegistered, userId []common.Address, role []uint8) (event.Subscription, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}

	logs, sub, err := _User.contract.WatchLogs(opts, "UserRegistered", userIdRule, roleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserUserRegistered)
				if err := _User.contract.UnpackLog(event, "UserRegistered", log); err != nil {
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

// ParseUserRegistered is a log parse operation binding the contract event 0x34f38c02d99abe0fa9a0ff44d588516965741625980bc7663e6405d8afa8df40.
//
// Solidity: event UserRegistered(address indexed userId, string name, uint8 indexed role, uint32 timestamp)
func (_User *UserFilterer) ParseUserRegistered(log types.Log) (*UserUserRegistered, error) {
	event := new(UserUserRegistered)
	if err := _User.contract.UnpackLog(event, "UserRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UserUserRoleUpdatedIterator is returned from FilterUserRoleUpdated and is used to iterate over the raw logs and unpacked data for UserRoleUpdated events raised by the User contract.
type UserUserRoleUpdatedIterator struct {
	Event *UserUserRoleUpdated // Event containing the contract specifics and raw log

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
func (it *UserUserRoleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UserUserRoleUpdated)
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
		it.Event = new(UserUserRoleUpdated)
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
func (it *UserUserRoleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UserUserRoleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UserUserRoleUpdated represents a UserRoleUpdated event raised by the User contract.
type UserUserRoleUpdated struct {
	UserId    common.Address
	NewRole   uint8
	Timestamp uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUserRoleUpdated is a free log retrieval operation binding the contract event 0xa1c1fdda1b3bc17fdb2a03c36c0bba716b0175888d93d15f335270635129bad9.
//
// Solidity: event UserRoleUpdated(address indexed userId, uint8 indexed newRole, uint32 timestamp)
func (_User *UserFilterer) FilterUserRoleUpdated(opts *bind.FilterOpts, userId []common.Address, newRole []uint8) (*UserUserRoleUpdatedIterator, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}
	var newRoleRule []interface{}
	for _, newRoleItem := range newRole {
		newRoleRule = append(newRoleRule, newRoleItem)
	}

	logs, sub, err := _User.contract.FilterLogs(opts, "UserRoleUpdated", userIdRule, newRoleRule)
	if err != nil {
		return nil, err
	}
	return &UserUserRoleUpdatedIterator{contract: _User.contract, event: "UserRoleUpdated", logs: logs, sub: sub}, nil
}

// WatchUserRoleUpdated is a free log subscription operation binding the contract event 0xa1c1fdda1b3bc17fdb2a03c36c0bba716b0175888d93d15f335270635129bad9.
//
// Solidity: event UserRoleUpdated(address indexed userId, uint8 indexed newRole, uint32 timestamp)
func (_User *UserFilterer) WatchUserRoleUpdated(opts *bind.WatchOpts, sink chan<- *UserUserRoleUpdated, userId []common.Address, newRole []uint8) (event.Subscription, error) {

	var userIdRule []interface{}
	for _, userIdItem := range userId {
		userIdRule = append(userIdRule, userIdItem)
	}
	var newRoleRule []interface{}
	for _, newRoleItem := range newRole {
		newRoleRule = append(newRoleRule, newRoleItem)
	}

	logs, sub, err := _User.contract.WatchLogs(opts, "UserRoleUpdated", userIdRule, newRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UserUserRoleUpdated)
				if err := _User.contract.UnpackLog(event, "UserRoleUpdated", log); err != nil {
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

// ParseUserRoleUpdated is a log parse operation binding the contract event 0xa1c1fdda1b3bc17fdb2a03c36c0bba716b0175888d93d15f335270635129bad9.
//
// Solidity: event UserRoleUpdated(address indexed userId, uint8 indexed newRole, uint32 timestamp)
func (_User *UserFilterer) ParseUserRoleUpdated(log types.Log) (*UserUserRoleUpdated, error) {
	event := new(UserUserRoleUpdated)
	if err := _User.contract.UnpackLog(event, "UserRoleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
