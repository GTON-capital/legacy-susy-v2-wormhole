// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package wormhole

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// WormholeGuardianSet is an auto generated low-level Go binding around an user-defined struct.
type WormholeGuardianSet struct {
	Keys           []common.Address
	ExpirationTime uint32
}

// WormholeParsedVAA is an auto generated low-level Go binding around an user-defined struct.
type WormholeParsedVAA struct {
	Version          uint8
	Hash             [32]byte
	GuardianSetIndex uint32
	Timestamp        uint32
	Action           uint8
	Payload          []byte
}

// AddressABI is the input ABI used to generate the binding from.
const AddressABI = "[]"

// AddressBin is the compiled bytecode used for deploying new contracts.
var AddressBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220413da8f481acaadbe6fa822f1508a34333ae148421b59f44aab5808cc361cd6264736f6c634300060a0033"

// DeployAddress deploys a new Ethereum contract, binding an instance of Address to it.
func DeployAddress(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Address, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AddressBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// Address is an auto generated Go binding around an Ethereum contract.
type Address struct {
	AddressCaller     // Read-only binding to the contract
	AddressTransactor // Write-only binding to the contract
	AddressFilterer   // Log filterer for contract events
}

// AddressCaller is an auto generated read-only Go binding around an Ethereum contract.
type AddressCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AddressTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AddressFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AddressSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AddressSession struct {
	Contract     *Address          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AddressCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AddressCallerSession struct {
	Contract *AddressCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// AddressTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AddressTransactorSession struct {
	Contract     *AddressTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AddressRaw is an auto generated low-level Go binding around an Ethereum contract.
type AddressRaw struct {
	Contract *Address // Generic contract binding to access the raw methods on
}

// AddressCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AddressCallerRaw struct {
	Contract *AddressCaller // Generic read-only contract binding to access the raw methods on
}

// AddressTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AddressTransactorRaw struct {
	Contract *AddressTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAddress creates a new instance of Address, bound to a specific deployed contract.
func NewAddress(address common.Address, backend bind.ContractBackend) (*Address, error) {
	contract, err := bindAddress(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Address{AddressCaller: AddressCaller{contract: contract}, AddressTransactor: AddressTransactor{contract: contract}, AddressFilterer: AddressFilterer{contract: contract}}, nil
}

// NewAddressCaller creates a new read-only instance of Address, bound to a specific deployed contract.
func NewAddressCaller(address common.Address, caller bind.ContractCaller) (*AddressCaller, error) {
	contract, err := bindAddress(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AddressCaller{contract: contract}, nil
}

// NewAddressTransactor creates a new write-only instance of Address, bound to a specific deployed contract.
func NewAddressTransactor(address common.Address, transactor bind.ContractTransactor) (*AddressTransactor, error) {
	contract, err := bindAddress(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AddressTransactor{contract: contract}, nil
}

// NewAddressFilterer creates a new log filterer instance of Address, bound to a specific deployed contract.
func NewAddressFilterer(address common.Address, filterer bind.ContractFilterer) (*AddressFilterer, error) {
	contract, err := bindAddress(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AddressFilterer{contract: contract}, nil
}

// bindAddress binds a generic wrapper to an already deployed contract.
func bindAddress(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AddressABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.AddressCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.AddressTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Address *AddressCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Address.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Address *AddressTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Address.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Address *AddressTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Address.Contract.contract.Transact(opts, method, params...)
}

// BytesLibABI is the input ABI used to generate the binding from.
const BytesLibABI = "[]"

// BytesLibBin is the compiled bytecode used for deploying new contracts.
var BytesLibBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212202c7dac7eff89eeea0d54f2b6a3ca80a6836bb98ebe599dd20a021093bc7b94ff64736f6c634300060a0033"

// DeployBytesLib deploys a new Ethereum contract, binding an instance of BytesLib to it.
func DeployBytesLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BytesLib, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BytesLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// BytesLib is an auto generated Go binding around an Ethereum contract.
type BytesLib struct {
	BytesLibCaller     // Read-only binding to the contract
	BytesLibTransactor // Write-only binding to the contract
	BytesLibFilterer   // Log filterer for contract events
}

// BytesLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type BytesLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BytesLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BytesLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BytesLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BytesLibSession struct {
	Contract     *BytesLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BytesLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BytesLibCallerSession struct {
	Contract *BytesLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BytesLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BytesLibTransactorSession struct {
	Contract     *BytesLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BytesLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type BytesLibRaw struct {
	Contract *BytesLib // Generic contract binding to access the raw methods on
}

// BytesLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BytesLibCallerRaw struct {
	Contract *BytesLibCaller // Generic read-only contract binding to access the raw methods on
}

// BytesLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BytesLibTransactorRaw struct {
	Contract *BytesLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBytesLib creates a new instance of BytesLib, bound to a specific deployed contract.
func NewBytesLib(address common.Address, backend bind.ContractBackend) (*BytesLib, error) {
	contract, err := bindBytesLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BytesLib{BytesLibCaller: BytesLibCaller{contract: contract}, BytesLibTransactor: BytesLibTransactor{contract: contract}, BytesLibFilterer: BytesLibFilterer{contract: contract}}, nil
}

// NewBytesLibCaller creates a new read-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibCaller(address common.Address, caller bind.ContractCaller) (*BytesLibCaller, error) {
	contract, err := bindBytesLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibCaller{contract: contract}, nil
}

// NewBytesLibTransactor creates a new write-only instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibTransactor(address common.Address, transactor bind.ContractTransactor) (*BytesLibTransactor, error) {
	contract, err := bindBytesLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BytesLibTransactor{contract: contract}, nil
}

// NewBytesLibFilterer creates a new log filterer instance of BytesLib, bound to a specific deployed contract.
func NewBytesLibFilterer(address common.Address, filterer bind.ContractFilterer) (*BytesLibFilterer, error) {
	contract, err := bindBytesLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BytesLibFilterer{contract: contract}, nil
}

// bindBytesLib binds a generic wrapper to an already deployed contract.
func bindBytesLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BytesLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.BytesLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.BytesLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BytesLib *BytesLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BytesLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BytesLib *BytesLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BytesLib *BytesLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BytesLib.Contract.contract.Transact(opts, method, params...)
}

// ContextABI is the input ABI used to generate the binding from.
const ContextABI = "[]"

// Context is an auto generated Go binding around an Ethereum contract.
type Context struct {
	ContextCaller     // Read-only binding to the contract
	ContextTransactor // Write-only binding to the contract
	ContextFilterer   // Log filterer for contract events
}

// ContextCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContextCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContextTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContextFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContextSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContextSession struct {
	Contract     *Context          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContextCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContextCallerSession struct {
	Contract *ContextCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ContextTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContextTransactorSession struct {
	Contract     *ContextTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ContextRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContextRaw struct {
	Contract *Context // Generic contract binding to access the raw methods on
}

// ContextCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContextCallerRaw struct {
	Contract *ContextCaller // Generic read-only contract binding to access the raw methods on
}

// ContextTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContextTransactorRaw struct {
	Contract *ContextTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContext creates a new instance of Context, bound to a specific deployed contract.
func NewContext(address common.Address, backend bind.ContractBackend) (*Context, error) {
	contract, err := bindContext(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Context{ContextCaller: ContextCaller{contract: contract}, ContextTransactor: ContextTransactor{contract: contract}, ContextFilterer: ContextFilterer{contract: contract}}, nil
}

// NewContextCaller creates a new read-only instance of Context, bound to a specific deployed contract.
func NewContextCaller(address common.Address, caller bind.ContractCaller) (*ContextCaller, error) {
	contract, err := bindContext(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContextCaller{contract: contract}, nil
}

// NewContextTransactor creates a new write-only instance of Context, bound to a specific deployed contract.
func NewContextTransactor(address common.Address, transactor bind.ContractTransactor) (*ContextTransactor, error) {
	contract, err := bindContext(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContextTransactor{contract: contract}, nil
}

// NewContextFilterer creates a new log filterer instance of Context, bound to a specific deployed contract.
func NewContextFilterer(address common.Address, filterer bind.ContractFilterer) (*ContextFilterer, error) {
	contract, err := bindContext(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContextFilterer{contract: contract}, nil
}

// bindContext binds a generic wrapper to an already deployed contract.
func bindContext(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContextABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.ContextCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.ContextTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Context *ContextCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Context.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Context *ContextTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Context.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Context *ContextTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Context.Contract.contract.Transact(opts, method, params...)
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20Bin is the compiled bytecode used for deploying new contracts.
var ERC20Bin = "0x60806040523480156200001157600080fd5b5060405162000c1f38038062000c1f8339810160408190526200003491620001c1565b81516200004990600390602085019062000075565b5080516200005f90600490602084019062000075565b50506005805460ff191660121790555062000228565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620000b857805160ff1916838001178555620000e8565b82800160010185558215620000e8579182015b82811115620000e8578251825591602001919060010190620000cb565b50620000f6929150620000fa565b5090565b6200011791905b80821115620000f6576000815560010162000101565b90565b600082601f8301126200012b578081fd5b81516001600160401b038082111562000142578283fd5b6040516020601f8401601f191682018101838111838210171562000164578586fd5b806040525081945083825286818588010111156200018157600080fd5b600092505b83831015620001a5578583018101518284018201529182019162000186565b83831115620001b75760008185840101525b5050505092915050565b60008060408385031215620001d4578182fd5b82516001600160401b0380821115620001eb578384fd5b620001f9868387016200011a565b935060208501519150808211156200020f578283fd5b506200021e858286016200011a565b9150509250929050565b6109e780620002386000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c80633950935111610071578063395093511461012957806370a082311461013c57806395d89b411461014f578063a457c2d714610157578063a9059cbb1461016a578063dd62ed3e1461017d576100a9565b806306fdde03146100ae578063095ea7b3146100cc57806318160ddd146100ec57806323b872dd14610101578063313ce56714610114575b600080fd5b6100b6610190565b6040516100c39190610777565b60405180910390f35b6100df6100da366004610742565b610226565b6040516100c3919061076c565b6100f4610244565b6040516100c3919061090f565b6100df61010f366004610702565b61024a565b61011c6102d7565b6040516100c39190610918565b6100df610137366004610742565b6102e0565b6100f461014a3660046106b3565b610334565b6100b661034f565b6100df610165366004610742565b6103b0565b6100df610178366004610742565b61041e565b6100f461018b3660046106ce565b610432565b60038054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561021c5780601f106101f15761010080835404028352916020019161021c565b820191906000526020600020905b8154815290600101906020018083116101ff57829003601f168201915b5050505050905090565b600061023a61023361045d565b8484610461565b5060015b92915050565b60025490565b600061025784848461051e565b6102cd8461026361045d565b6102c885604051806060016040528060288152602001610965602891396001600160a01b038a166000908152600160205260408120906102a161045d565b6001600160a01b03168152602081019190915260400160002054919063ffffffff61063f16565b610461565b5060019392505050565b60055460ff1690565b600061023a6102ed61045d565b846102c885600160006102fe61045d565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff61066b16565b6001600160a01b031660009081526020819052604090205490565b60048054604080516020601f600260001961010060018816150201909516949094049384018190048102820181019092528281526060939092909183018282801561021c5780601f106101f15761010080835404028352916020019161021c565b600061023a6103bd61045d565b846102c88560405180606001604052806025815260200161098d60259139600160006103e761045d565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff61063f16565b600061023a61042b61045d565b848461051e565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b3390565b6001600160a01b0383166104905760405162461bcd60e51b8152600401610487906108cb565b60405180910390fd5b6001600160a01b0382166104b65760405162461bcd60e51b81526004016104879061080d565b6001600160a01b0380841660008181526001602090815260408083209487168084529490915290819020849055517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259061051190859061090f565b60405180910390a3505050565b6001600160a01b0383166105445760405162461bcd60e51b815260040161048790610886565b6001600160a01b03821661056a5760405162461bcd60e51b8152600401610487906107ca565b610575838383610697565b6105b88160405180606001604052806026815260200161093f602691396001600160a01b038616600090815260208190526040902054919063ffffffff61063f16565b6001600160a01b0380851660009081526020819052604080822093909355908416815220546105ed908263ffffffff61066b16565b6001600160a01b0380841660008181526020819052604090819020939093559151908516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9061051190859061090f565b600081848411156106635760405162461bcd60e51b81526004016104879190610777565b505050900390565b6000828201838110156106905760405162461bcd60e51b81526004016104879061084f565b9392505050565b505050565b80356001600160a01b038116811461023e57600080fd5b6000602082840312156106c4578081fd5b610690838361069c565b600080604083850312156106e0578081fd5b6106ea848461069c565b91506106f9846020850161069c565b90509250929050565b600080600060608486031215610716578081fd5b833561072181610926565b9250602084013561073181610926565b929592945050506040919091013590565b60008060408385031215610754578182fd5b61075e848461069c565b946020939093013593505050565b901515815260200190565b6000602080835283518082850152825b818110156107a357858101830151858201604001528201610787565b818111156107b45783604083870101525b50601f01601f1916929092016040019392505050565b60208082526023908201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260408201526265737360e81b606082015260800190565b60208082526022908201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604082015261737360f01b606082015260800190565b6020808252601b908201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604082015260600190565b60208082526025908201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604082015264647265737360d81b606082015260800190565b60208082526024908201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646040820152637265737360e01b606082015260800190565b90815260200190565b60ff91909116815260200190565b6001600160a01b038116811461093b57600080fd5b5056fe45524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa26469706673582212201f135f5f9ac34ef7ef0142599c9d5f0ebab54f063833dca1b3518daf87b92dd164736f6c634300060a0033"

// DeployERC20 deploys a new Ethereum contract, binding an instance of ERC20 to it.
func DeployERC20(auth *bind.TransactOpts, backend bind.ContractBackend, name string, symbol string) (common.Address, *types.Transaction, *ERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20Bin), backend, name, symbol)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20Session) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_ERC20 *ERC20CallerSession) Decimals() (uint8, error) {
	return _ERC20.Contract.Decimals(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20Session) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ERC20 *ERC20CallerSession) Name() (string, error) {
	return _ERC20.Contract.Name(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20Session) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ERC20 *ERC20CallerSession) Symbol() (string, error) {
	return _ERC20.Contract.Symbol(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20ABI is the input ABI used to generate the binding from.
const IERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IERC20FuncSigs maps the 4-byte function signature to its string representation.
var IERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReentrancyGuardABI is the input ABI used to generate the binding from.
const ReentrancyGuardABI = "[]"

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
}

// SafeERC20ABI is the input ABI used to generate the binding from.
const SafeERC20ABI = "[]"

// SafeERC20Bin is the compiled bytecode used for deploying new contracts.
var SafeERC20Bin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212204844ee880e73bfa5c3fd56a89cefdd26f6eab2d3a951f94c90cff0a6e1f81c7a64736f6c634300060a0033"

// DeploySafeERC20 deploys a new Ethereum contract, binding an instance of SafeERC20 to it.
func DeploySafeERC20(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeERC20, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeERC20Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// SafeERC20 is an auto generated Go binding around an Ethereum contract.
type SafeERC20 struct {
	SafeERC20Caller     // Read-only binding to the contract
	SafeERC20Transactor // Write-only binding to the contract
	SafeERC20Filterer   // Log filterer for contract events
}

// SafeERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type SafeERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeERC20Session struct {
	Contract     *SafeERC20        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeERC20CallerSession struct {
	Contract *SafeERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// SafeERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeERC20TransactorSession struct {
	Contract     *SafeERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// SafeERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type SafeERC20Raw struct {
	Contract *SafeERC20 // Generic contract binding to access the raw methods on
}

// SafeERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeERC20CallerRaw struct {
	Contract *SafeERC20Caller // Generic read-only contract binding to access the raw methods on
}

// SafeERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeERC20TransactorRaw struct {
	Contract *SafeERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeERC20 creates a new instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20(address common.Address, backend bind.ContractBackend) (*SafeERC20, error) {
	contract, err := bindSafeERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeERC20{SafeERC20Caller: SafeERC20Caller{contract: contract}, SafeERC20Transactor: SafeERC20Transactor{contract: contract}, SafeERC20Filterer: SafeERC20Filterer{contract: contract}}, nil
}

// NewSafeERC20Caller creates a new read-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Caller(address common.Address, caller bind.ContractCaller) (*SafeERC20Caller, error) {
	contract, err := bindSafeERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Caller{contract: contract}, nil
}

// NewSafeERC20Transactor creates a new write-only instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*SafeERC20Transactor, error) {
	contract, err := bindSafeERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Transactor{contract: contract}, nil
}

// NewSafeERC20Filterer creates a new log filterer instance of SafeERC20, bound to a specific deployed contract.
func NewSafeERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*SafeERC20Filterer, error) {
	contract, err := bindSafeERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeERC20Filterer{contract: contract}, nil
}

// bindSafeERC20 binds a generic wrapper to an already deployed contract.
func bindSafeERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.SafeERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.SafeERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeERC20 *SafeERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeERC20 *SafeERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeERC20 *SafeERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeERC20.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212209bf7862da5bf11451c308e3c0cd6f9436da3652e202b357029c9eca905258f2364736f6c634300060a0033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// WETHABI is the input ABI used to generate the binding from.
const WETHABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// WETHFuncSigs maps the 4-byte function signature to its string representation.
var WETHFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"d0e30db0": "deposit()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"2e1a7d4d": "withdraw(uint256)",
}

// WETH is an auto generated Go binding around an Ethereum contract.
type WETH struct {
	WETHCaller     // Read-only binding to the contract
	WETHTransactor // Write-only binding to the contract
	WETHFilterer   // Log filterer for contract events
}

// WETHCaller is an auto generated read-only Go binding around an Ethereum contract.
type WETHCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WETHTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WETHFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WETHSession struct {
	Contract     *WETH             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WETHCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WETHCallerSession struct {
	Contract *WETHCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// WETHTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WETHTransactorSession struct {
	Contract     *WETHTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WETHRaw is an auto generated low-level Go binding around an Ethereum contract.
type WETHRaw struct {
	Contract *WETH // Generic contract binding to access the raw methods on
}

// WETHCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WETHCallerRaw struct {
	Contract *WETHCaller // Generic read-only contract binding to access the raw methods on
}

// WETHTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WETHTransactorRaw struct {
	Contract *WETHTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWETH creates a new instance of WETH, bound to a specific deployed contract.
func NewWETH(address common.Address, backend bind.ContractBackend) (*WETH, error) {
	contract, err := bindWETH(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WETH{WETHCaller: WETHCaller{contract: contract}, WETHTransactor: WETHTransactor{contract: contract}, WETHFilterer: WETHFilterer{contract: contract}}, nil
}

// NewWETHCaller creates a new read-only instance of WETH, bound to a specific deployed contract.
func NewWETHCaller(address common.Address, caller bind.ContractCaller) (*WETHCaller, error) {
	contract, err := bindWETH(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WETHCaller{contract: contract}, nil
}

// NewWETHTransactor creates a new write-only instance of WETH, bound to a specific deployed contract.
func NewWETHTransactor(address common.Address, transactor bind.ContractTransactor) (*WETHTransactor, error) {
	contract, err := bindWETH(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WETHTransactor{contract: contract}, nil
}

// NewWETHFilterer creates a new log filterer instance of WETH, bound to a specific deployed contract.
func NewWETHFilterer(address common.Address, filterer bind.ContractFilterer) (*WETHFilterer, error) {
	contract, err := bindWETH(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WETHFilterer{contract: contract}, nil
}

// bindWETH binds a generic wrapper to an already deployed contract.
func bindWETH(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WETHABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETH *WETHRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETH.Contract.WETHCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETH *WETHRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH.Contract.WETHTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETH *WETHRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETH.Contract.WETHTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETH *WETHCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETH.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETH *WETHTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETH *WETHTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETH.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WETH *WETHCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WETH.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WETH *WETHSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WETH.Contract.Allowance(&_WETH.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WETH *WETHCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WETH.Contract.Allowance(&_WETH.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WETH *WETHCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WETH.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WETH *WETHSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WETH.Contract.BalanceOf(&_WETH.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WETH *WETHCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WETH.Contract.BalanceOf(&_WETH.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WETH *WETHCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WETH.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WETH *WETHSession) TotalSupply() (*big.Int, error) {
	return _WETH.Contract.TotalSupply(&_WETH.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WETH *WETHCallerSession) TotalSupply() (*big.Int, error) {
	return _WETH.Contract.TotalSupply(&_WETH.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WETH *WETHTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETH.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WETH *WETHSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETH.Contract.Approve(&_WETH.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WETH *WETHTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETH.Contract.Approve(&_WETH.TransactOpts, spender, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WETH *WETHTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETH.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WETH *WETHSession) Deposit() (*types.Transaction, error) {
	return _WETH.Contract.Deposit(&_WETH.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WETH *WETHTransactorSession) Deposit() (*types.Transaction, error) {
	return _WETH.Contract.Deposit(&_WETH.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WETH *WETHTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETH.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WETH *WETHSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETH.Contract.Transfer(&_WETH.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WETH *WETHTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETH.Contract.Transfer(&_WETH.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WETH *WETHTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETH.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WETH *WETHSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETH.Contract.TransferFrom(&_WETH.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WETH *WETHTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETH.Contract.TransferFrom(&_WETH.TransactOpts, sender, recipient, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WETH *WETHTransactor) Withdraw(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _WETH.contract.Transact(opts, "withdraw", amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WETH *WETHSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _WETH.Contract.Withdraw(&_WETH.TransactOpts, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 amount) returns()
func (_WETH *WETHTransactorSession) Withdraw(amount *big.Int) (*types.Transaction, error) {
	return _WETH.Contract.Withdraw(&_WETH.TransactOpts, amount)
}

// WETHApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WETH contract.
type WETHApprovalIterator struct {
	Event *WETHApproval // Event containing the contract specifics and raw log

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
func (it *WETHApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETHApproval)
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
		it.Event = new(WETHApproval)
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
func (it *WETHApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WETHApproval represents a Approval event raised by the WETH contract.
type WETHApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WETH *WETHFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WETHApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WETH.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WETHApprovalIterator{contract: _WETH.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WETH *WETHFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WETHApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WETH.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHApproval)
				if err := _WETH.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WETH *WETHFilterer) ParseApproval(log types.Log) (*WETHApproval, error) {
	event := new(WETHApproval)
	if err := _WETH.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WETHTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WETH contract.
type WETHTransferIterator struct {
	Event *WETHTransfer // Event containing the contract specifics and raw log

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
func (it *WETHTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETHTransfer)
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
		it.Event = new(WETHTransfer)
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
func (it *WETHTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WETHTransfer represents a Transfer event raised by the WETH contract.
type WETHTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WETH *WETHFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WETHTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WETH.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WETHTransferIterator{contract: _WETH.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WETH *WETHFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WETHTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WETH.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHTransfer)
				if err := _WETH.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WETH *WETHFilterer) ParseTransfer(log types.Log) (*WETHTransfer, error) {
	event := new(WETHTransfer)
	if err := _WETH.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WormholeABI is the input ABI used to generate the binding from.
const WormholeABI = "[{\"inputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"keys\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"expiration_time\",\"type\":\"uint32\"}],\"internalType\":\"structWormhole.GuardianSet\",\"name\":\"initial_guardian_set\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"wrapped_asset_master\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"_guardian_set_expirity\",\"type\":\"uint32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"oldGuardianIndex\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newGuardianIndex\",\"type\":\"uint32\"}],\"name\":\"LogGuardianSetChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"target_chain\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"token_chain\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"token_decimals\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"token\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sender\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"}],\"name\":\"LogTokensLocked\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"consumedVAAs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"idx\",\"type\":\"uint32\"}],\"name\":\"getGuardianSet\",\"outputs\":[{\"components\":[{\"internalType\":\"address[]\",\"name\":\"keys\",\"type\":\"address[]\"},{\"internalType\":\"uint32\",\"name\":\"expiration_time\",\"type\":\"uint32\"}],\"internalType\":\"structWormhole.GuardianSet\",\"name\":\"gs\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardian_set_expirity\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"guardian_set_index\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"guardian_sets\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"expiration_time\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"isWrappedAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"target_chain\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"refund_dust\",\"type\":\"bool\"}],\"name\":\"lockAssets\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"recipient\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"target_chain\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"nonce\",\"type\":\"uint32\"}],\"name\":\"lockETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"vaa\",\"type\":\"bytes\"}],\"name\":\"parseAndVerifyVAA\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"guardian_set_index\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"action\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"}],\"internalType\":\"structWormhole.ParsedVAA\",\"name\":\"parsed_vaa\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"vaa\",\"type\":\"bytes\"}],\"name\":\"submitVAA\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wrappedAssetMaster\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"wrappedAssets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// WormholeFuncSigs maps the 4-byte function signature to its string representation.
var WormholeFuncSigs = map[string]string{
	"a31fe409": "consumedVAAs(bytes32)",
	"f951975a": "getGuardianSet(uint32)",
	"4db47840": "guardian_set_expirity()",
	"822d82b3": "guardian_set_index()",
	"42b0aefa": "guardian_sets(uint32)",
	"1a2be4da": "isWrappedAsset(address)",
	"70713960": "lockAssets(address,uint256,bytes32,uint8,uint32,bool)",
	"58d62e46": "lockETH(bytes32,uint8,uint32)",
	"600b9aa6": "parseAndVerifyVAA(bytes)",
	"3bc0aee6": "submitVAA(bytes)",
	"99da1d3c": "wrappedAssetMaster()",
	"b6694c2a": "wrappedAssets(bytes32)",
}

// WormholeBin is the compiled bytecode used for deploying new contracts.
var WormholeBin = "0x60806040526001805460ff60a01b1916600160a11b1790553480156200002457600080fd5b50604051620029a7380380620029a78339810160408190526200004791620001be565b600160009081558052600260209081528351805185927fac33ff75c19e70fe83507db0d683fd3465c996598dc972688b7ace676c89077b92620000919284929190910190620000f7565b50602091909101516001918201805463ffffffff191663ffffffff928316179055600380546001600160401b03191664010000000094909216939093021790915580546001600160a01b0319166001600160a01b03929092169190911790555062000322565b8280548282559060005260206000209081019282156200014f579160200282015b828111156200014f57825182546001600160a01b0319166001600160a01b0390911617825560209092019160019091019062000118565b506200015d92915062000161565b5090565b6200018891905b808211156200015d5780546001600160a01b031916815560010162000168565b90565b80516001600160a01b0381168114620001a357600080fd5b92915050565b805163ffffffff81168114620001a357600080fd5b600080600060608486031215620001d3578283fd5b83516001600160401b0380821115620001ea578485fd5b81860160408189031215620001fd578586fd5b620002096040620002db565b92508051828111156200021a578687fd5b81019150601f820188136200022d578586fd5b8151620002446200023e8262000302565b620002db565b80828252602080830192508086018c82838702890101111562000265578a8bfd5b8a96505b8487101562000293576200027e8d826200018b565b84526001969096019592810192810162000269565b50818752620002a58c828701620001a9565b81880152869950620002ba8c828d016200018b565b985050505050505050620002d28560408601620001a9565b90509250925092565b6040518181016001600160401b0381118282101715620002fa57600080fd5b604052919050565b60006001600160401b0382111562000318578081fd5b5060209081020190565b61267580620003326000396000f3fe6080604052600436106100ab5760003560e01c8063707139601161006457806370713960146101c3578063822d82b3146101e357806399da1d3c146101f8578063a31fe4091461021a578063b6694c2a1461023a578063f951975a1461025a576100d1565b80631a2be4da146100e95780633bc0aee61461011f57806342b0aefa146101415780634db478401461016e57806358d62e4614610183578063600b9aa614610196576100d1565b366100d15760405162461bcd60e51b81526004016100c890612075565b60405180910390fd5b60405162461bcd60e51b81526004016100c890612075565b3480156100f557600080fd5b50610109610104366004611d21565b610287565b6040516101169190611f93565b60405180910390f35b34801561012b57600080fd5b5061013f61013a366004611e33565b61029c565b005b34801561014d57600080fd5b5061016161015c366004611e9d565b610386565b6040516101169190612576565b34801561017a57600080fd5b506101616103a1565b61013f610191366004611df5565b6103b5565b3480156101a257600080fd5b506101b66101b1366004611e33565b610565565b6040516101169190612517565b3480156101cf57600080fd5b5061013f6101de366004611d3c565b610b91565b3480156101ef57600080fd5b506101616110b7565b34801561020457600080fd5b5061020d6110c3565b6040516101169190611f42565b34801561022657600080fd5b50610109610235366004611dc5565b6110d2565b34801561024657600080fd5b5061020d610255366004611dc5565b6110e7565b34801561026657600080fd5b5061027a610275366004611e9d565b611102565b60405161011691906124ac565b60066020526000908152604090205460ff1681565b600260005414156102bf5760405162461bcd60e51b81526004016100c890612447565b60026000556102cc611c1b565b6102d68383610565565b9050806080015160ff166001141561032957600354604082015163ffffffff9081169116146103175760405162461bcd60e51b81526004016100c8906120c1565b6103248160a0015161119a565b61035e565b806080015160ff1660101415610346576103248160a0015161138d565b60405162461bcd60e51b81526004016100c890612012565b6020908101516000908152600490915260408120805460ff1916600190811790915590555050565b60026020526000908152604090206001015463ffffffff1681565b600354640100000000900463ffffffff1681565b600260005414156103d85760405162461bcd60e51b81526004016100c890612447565b600260005560015460ff838116600160a01b90920416141561040c5760405162461bcd60e51b81526004016100c890612181565b600061042234633b9aca0063ffffffff61162516565b9050600061043a34633b9aca0063ffffffff61167016565b9050806104595760405162461bcd60e51b81526004016100c890612234565b604051339083156108fc029084906000818181858888f19350505050158015610486573d6000803e3d6000fd5b5073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc26001600160a01b031663d0e30db08334036040518263ffffffff1660e01b81526004016000604051808303818588803b1580156104d857600080fd5b505af11580156104ec573d6000803e3d6000fd5b505060015460405133945073c02aaa39b223fe8d0a0e5c4f27ead9083c756cc293507f6bbd554ad75919f71fd91bf917ca6e4f41c10f03ab25751596a22253bb39aab89250610551918991600160a01b90910460ff16906009908c9089908c906125bb565b60405180910390a350506001600055505050565b61056d611c1b565b6105b7600084848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6116b2169050565b60ff168082526001146105dc5760405162461bcd60e51b81526004016100c89061214a565b610626600184848080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6116e1169050565b63ffffffff1660408083019190915280516020601f850181900481028201810190925283815260009161067f916005918790879081908401838280828437600092019190915250929392505063ffffffff6116b2169050565b60ff16905060008160420260060190506106d88186868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6116e1169050565b63ffffffff166060840152604080516020601f870181900481028201810190925285815261073091839182880391899089908190840183828082843760009201919091525092949392505063ffffffff611710169050565b805160209182012084820181905260009081526004909152604090205460ff161561076d5760405162461bcd60e51b81526004016100c89061203e565b610775611c52565b60408085015163ffffffff166000908152600260209081529082902082518154606093810282018401855293810184815290939192849284918401828280156107e757602002820191906000526020600020905b81546001600160a01b031681526001909101906020018083116107c9575b50505091835250506001919091015463ffffffff166020909101528051519091506108245760405162461bcd60e51b81526004016100c89061247e565b602081015163ffffffff161580610844575042816020015163ffffffff16115b6108605760405162461bcd60e51b81526004016100c890612306565b82600a6003836000015151600a028161087557fe5b046002028161088057fe5b0460010111156108a25760405162461bcd60e51b81526004016100c8906122ac565b60001960005b84811015610ad8576000610901826042026006018a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6116b2169050565b90508260010b8160ff16136109285760405162461bcd60e51b81526004016100c890612374565b8060ff169250600061097f836042026007018b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6117a3169050565b905060006109d2846042026027018c8c8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6117a3169050565b90506000610a25856042026047018d8d8080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6116b2169050565b87518051601b90920192509060ff8616908110610a3e57fe5b60200260200101516001600160a01b031660018b6020015183868660405160008152602001604052604051610a769493929190611f9e565b6020604051602081039080840390855afa158015610a98573d6000803e3d6000fd5b505050602060405103516001600160a01b031614610ac85760405162461bcd60e51b81526004016100c8906121c4565b5050600190920191506108a89050565b50610b258360040188888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250929392505063ffffffff6116b2169050565b60ff166080860152604080516020601f8901819004810282018101909252878152610b81916005860191600419878b0301918b908b908190840183828082843760009201919091525092949392505063ffffffff611710169050565b60a0860152509295945050505050565b60026000541415610bb45760405162461bcd60e51b81526004016100c890612447565b600260005560015460ff848116600160a01b909204161415610be85760405162461bcd60e51b81526004016100c890612181565b6000600160149054906101000a900460ff169050600080886001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b158015610c3857600080fd5b505afa158015610c4c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c709190611ec1565b6001600160a01b038a1660009081526006602052604090205490915060ff1615610ddf57604051632770a7eb60e21b81526001600160a01b038a1690639dc29fac90610cc29033908c90600401611f56565b600060405180830381600087803b158015610cdc57600080fd5b505af1158015610cf0573d6000803e3d6000fd5b50505050886001600160a01b031663026b05396040518163ffffffff1660e01b815260040160206040518083038186803b158015610d2d57600080fd5b505afa158015610d41573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d659190611ec1565b9250886001600160a01b0316631ba46cfd6040518163ffffffff1660e01b815260040160206040518083038186803b158015610da057600080fd5b505afa158015610db4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dd89190611ddd565b915061103b565b6040516370a0823160e01b81526000906001600160a01b038b16906370a0823190610e0e903090600401611f42565b60206040518083038186803b158015610e2657600080fd5b505afa158015610e3a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e5e9190611ddd565b9050610e7b6001600160a01b038b1633308c63ffffffff6117d216565b6040516370a0823160e01b81526000906001600160a01b038c16906370a0823190610eaa903090600401611f42565b60206040518083038186803b158015610ec257600080fd5b505afa158015610ed6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610efa9190611ddd565b9050610f0c818363ffffffff61183016565b995060098360ff161115610f785789610f358160ff600819870116600a0a63ffffffff61167016565b9a508615610f7257610f7233610f5b8360ff600819890116600a0a63ffffffff61162516565b6001600160a01b038f16919063ffffffff61187216565b60099350505b67ffffffffffffffff801661100e60098d6001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b158015610fc157600080fd5b505afa158015610fd5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ff99190611ec1565b849160ff910316600a0a63ffffffff61167016565b111561102c5760405162461bcd60e51b81526004016100c890611fcf565b50506001600160a01b03891691505b876110585760405162461bcd60e51b81526004016100c890612234565b336001600160a01b031660001b827f6bbd554ad75919f71fd91bf917ca6e4f41c10f03ab25751596a22253bb39aab88886858c8e8c60405161109f969594939291906125bb565b60405180910390a35050600160005550505050505050565b60035463ffffffff1681565b6001546001600160a01b031681565b60046020526000908152604090205460ff1681565b6005602052600090815260409020546001600160a01b031681565b61110a611c52565b63ffffffff8216600090815260026020908152604091829020825181546060938102820184018552938101848152909391928492849184018282801561117957602002820191906000526020600020905b81546001600160a01b0316815260019091019060200180831161115b575b50505091835250506001919091015463ffffffff1660209091015292915050565b60006111ac828263ffffffff6116e116565b60035490915063ffffffff80831691811660010116146111de5760405162461bcd60e51b81526004016100c89061226b565b60006111f183600463ffffffff6116b216565b905060608160ff1667ffffffffffffffff8111801561120f57600080fd5b50604051908082528060200260200182016040528015611239578160200160208202803683370190505b50905060005b8260ff16811015611292576000611262866005601485020163ffffffff61189616565b90508083838151811061127157fe5b6001600160a01b03909216602092830291909101909101525060010161123f565b506003805463ffffffff85811663ffffffff19831617909255166112b4611c52565b506040805180820182528381526000602080830182905260035463ffffffff1682526002815292902081518051929384936112f29284920190611c6a565b506020918201516001918201805463ffffffff1990811663ffffffff9384161790915560038054878416600090815260029096526040958690209094018054909216640100000000909404831642018316939093179055905491517fdfb80683934199683861bf00b64ecdf0984bbaf661bf27983dba382e99297a629261137d928692911690612587565b60405180910390a1505050505050565b60006113a082600463ffffffff6116b216565b905060006113b583600563ffffffff6116b216565b905060006113ca84603263ffffffff61189616565b905060006113df85604663ffffffff6116b216565b905060006113f486606863ffffffff6117a316565b90508360ff168560ff16141561141c5760405162461bcd60e51b81526004016100c890612401565b60015460ff858116600160a01b909204161461144a5760405162461bcd60e51b81526004016100c8906122cf565b60015460ff838116600160a01b909204161461154d57600061147387604763ffffffff6117a316565b90506000838260405160200161148a929190611f25565b60408051601f198184030181529181528151602092830120600081815260059093529120549091506001600160a01b0316806114e55760006114d38a606763ffffffff6116b216565b90506114e1838786846118cc565b9150505b6040516340c10f1960e01b81526001600160a01b038216906340c10f19906115139089908890600401611f56565b600060405180830381600087803b15801561152d57600080fd5b505af1158015611541573d6000803e3d6000fd5b5050505050505061161d565b600061156087605363ffffffff61189616565b90506000816001600160a01b031663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b15801561159d57600080fd5b505afa1580156115b1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115d59190611ec1565b905060098160ff161115611600576115fd8360ff600819840116600a0a63ffffffff6119da16565b92505b61161a6001600160a01b038316868563ffffffff61187216565b50505b505050505050565b600061166783836040518060400160405280601881526020017f536166654d6174683a206d6f64756c6f206279207a65726f0000000000000000815250611a14565b90505b92915050565b600061166783836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815250611a48565b600081600101835110156116d85760405162461bcd60e51b81526004016100c89061211e565b50016001015190565b600081600401835110156117075760405162461bcd60e51b81526004016100c89061211e565b50016004015190565b6060818301845110156117355760405162461bcd60e51b81526004016100c89061211e565b6060821580156117505760405191506020820160405261179a565b6040519150601f8416801560200281840101858101878315602002848b0101015b81831015611789578051835260209283019201611771565b5050858452601f01601f1916604052505b50949350505050565b600081602001835110156117c95760405162461bcd60e51b81526004016100c89061211e565b50016020015190565b61182a846323b872dd60e01b8585856040516024016117f393929190611f6f565b60408051601f198184030181529190526020810180516001600160e01b03166001600160e01b031990931692909217909152611a7f565b50505050565b600061166783836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250611b0e565b6118918363a9059cbb60e01b84846040516024016117f3929190611f56565b505050565b600081601401835110156118bc5760405162461bcd60e51b81526004016100c89061211e565b500160200151600160601b900490565b600154604051733d602d80600a3d3981f3363d3d373d3d3d363d7360601b815260609190911b6bffffffffffffffffffffffff1916601482018190526e5af43d82803e903d91602b57fd5bf360881b60288301526000918660378285f560405163a7a2d3fb60e01b81529093506001600160a01b038416915063a7a2d3fb9061195d9088908890889060040161259e565b600060405180830381600087803b15801561197757600080fd5b505af115801561198b573d6000803e3d6000fd5b5050506000968752505060056020908152604080872080546001600160a01b0319166001600160a01b03851690811790915587526006909152909420805460ff19166001179055509192915050565b6000826119e95750600061166a565b828202828482816119f657fe5b04146116675760405162461bcd60e51b81526004016100c8906121f3565b60008183611a355760405162461bcd60e51b81526004016100c89190611fbc565b50828481611a3f57fe5b06949350505050565b60008183611a695760405162461bcd60e51b81526004016100c89190611fbc565b506000838581611a7557fe5b0495945050505050565b6060611ad4826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b0316611b3a9092919063ffffffff16565b8051909150156118915780806020019051810190611af29190611da9565b6118915760405162461bcd60e51b81526004016100c8906123b7565b60008184841115611b325760405162461bcd60e51b81526004016100c89190611fbc565b505050900390565b6060611b498484600085611b51565b949350505050565b6060611b5c85611c15565b611b785760405162461bcd60e51b81526004016100c89061233d565b60006060866001600160a01b03168587604051611b959190611f09565b60006040518083038185875af1925050503d8060008114611bd2576040519150601f19603f3d011682016040523d82523d6000602084013e611bd7565b606091505b50915091508115611beb579150611b499050565b805115611bfb5780518082602001fd5b8360405162461bcd60e51b81526004016100c89190611fbc565b3b151590565b6040805160c0810182526000808252602082018190529181018290526060808201839052608082019290925260a081019190915290565b60408051808201909152606081526000602082015290565b828054828255906000526020600020908101928215611cbf579160200282015b82811115611cbf57825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611c8a565b50611ccb929150611ccf565b5090565b611cf391905b80821115611ccb5780546001600160a01b0319168155600101611cd5565b90565b80356001600160a01b038116811461166a57600080fd5b803563ffffffff8116811461166a57600080fd5b600060208284031215611d32578081fd5b6116678383611cf6565b60008060008060008060c08789031215611d54578182fd5b611d5e8888611cf6565b955060208701359450604087013593506060870135611d7c81612630565b9250611d8b8860808901611d0d565b915060a0870135611d9b8161261f565b809150509295509295509295565b600060208284031215611dba578081fd5b81516116678161261f565b600060208284031215611dd6578081fd5b5035919050565b600060208284031215611dee578081fd5b5051919050565b600080600060608486031215611e09578283fd5b833592506020840135611e1b81612630565b9150611e2a8560408601611d0d565b90509250925092565b60008060208385031215611e45578182fd5b823567ffffffffffffffff80821115611e5c578384fd5b81850186601f820112611e6d578485fd5b8035925081831115611e7d578485fd5b866020848301011115611e8e578485fd5b60200196919550909350505050565b600060208284031215611eae578081fd5b813563ffffffff81168114611667578182fd5b600060208284031215611ed2578081fd5b815161166781612630565b60008151808452611ef58160208601602086016125f3565b601f01601f19169290920160200192915050565b60008251611f1b8184602087016125f3565b9190910192915050565b60f89290921b6001600160f81b0319168252600182015260210190565b6001600160a01b0391909116815260200190565b6001600160a01b03929092168252602082015260400190565b6001600160a01b039384168152919092166020820152604081019190915260600190565b901515815260200190565b93845260ff9290921660208401526040830152606082015260800190565b6000602082526116676020830184611edd565b60208082526023908201527f6272696467652062616c616e636520776f756c6420657863656564206d6178696040820152626d756d60e81b606082015260800190565b60208082526012908201527134b73b30b634b2102b20a09030b1ba34b7b760711b604082015260600190565b60208082526018908201527f5641412077617320616c72656164792065786563757465640000000000000000604082015260600190565b6020808252602c908201527f706c6561736520757365206c6f636b45544820746f207472616e73666572204560408201526b544820746f20536f6c616e6160a01b606082015260800190565b60208082526039908201527f6f6e6c79207468652063757272656e7420677561726469616e2073657420636160408201527f6e206368616e67652074686520677561726469616e2073657400000000000000606082015260800190565b60208082526012908201527152656164206f7574206f6620626f756e647360701b604082015260600190565b60208082526018908201527f5641412076657273696f6e20696e636f6d70617469626c650000000000000000604082015260600190565b60208082526023908201527f6d757374206e6f74207472616e7366657220746f207468652073616d6520636860408201526230b4b760e91b606082015260800190565b602080825260159082015274159050481cda59db985d1d5c99481a5b9d985b1a59605a1b604082015260600190565b60208082526021908201527f536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f6040820152607760f81b606082015260800190565b6020808252601e908201527f7472756e636174656420616d6f756e74206d757374206e6f7420626520300000604082015260600190565b60208082526021908201527f696e646578206d75737420696e63726561736520696e207374657073206f66206040820152603160f81b606082015260800190565b6020808252600990820152686e6f2071756f72756d60b81b604082015260600190565b60208082526019908201527f7472616e73666572206d75737420626520696e636f6d696e6700000000000000604082015260600190565b60208082526018908201527f677561726469616e207365742068617320657870697265640000000000000000604082015260600190565b6020808252601d908201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e7472616374000000604082015260600190565b60208082526023908201527f7369676e617475726520696e6469636573206d75737420626520617363656e64604082015262696e6760e81b606082015260800190565b6020808252602a908201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6040820152691bdd081cdd58d8d9595960b21b606082015260800190565b60208082526026908201527f73616d6520636861696e207472616e736665727320617265206e6f74207375706040820152651c1bdc9d195960d21b606082015260800190565b6020808252601f908201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c00604082015260600190565b6020808252601490820152731a5b9d985b1a590819dd585c991a585b881cd95d60621b604082015260600190565b6020808252825160408383015280516060840181905260009291820190839060808601905b808310156124fa5783516001600160a01b031682529284019260019290920191908401906124d1565b5063ffffffff848801511660408701528094505050505092915050565b60006020825260ff835116602083015260208301516040830152604083015163ffffffff8082166060850152806060860151166080850152505060ff60808401511660a083015260a083015160c080840152611b4960e0840182611edd565b63ffffffff91909116815260200190565b63ffffffff92831681529116602082015260400190565b60ff93841681526020810192909252909116604082015260600190565b60ff968716815294861660208601529290941660408401526060830152608082019290925263ffffffff90911660a082015260c00190565b60005b8381101561260e5781810151838201526020016125f6565b8381111561182a5750506000910152565b801515811461262d57600080fd5b50565b60ff8116811461262d57600080fdfea264697066735822122040ac91cebe957be8dc113db8990f566f36e09f9f166886055adf3a558f815a5764736f6c634300060a0033"

// DeployWormhole deploys a new Ethereum contract, binding an instance of Wormhole to it.
func DeployWormhole(auth *bind.TransactOpts, backend bind.ContractBackend, initial_guardian_set WormholeGuardianSet, wrapped_asset_master common.Address, _guardian_set_expirity uint32) (common.Address, *types.Transaction, *Wormhole, error) {
	parsed, err := abi.JSON(strings.NewReader(WormholeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WormholeBin), backend, initial_guardian_set, wrapped_asset_master, _guardian_set_expirity)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Wormhole{WormholeCaller: WormholeCaller{contract: contract}, WormholeTransactor: WormholeTransactor{contract: contract}, WormholeFilterer: WormholeFilterer{contract: contract}}, nil
}

// Wormhole is an auto generated Go binding around an Ethereum contract.
type Wormhole struct {
	WormholeCaller     // Read-only binding to the contract
	WormholeTransactor // Write-only binding to the contract
	WormholeFilterer   // Log filterer for contract events
}

// WormholeCaller is an auto generated read-only Go binding around an Ethereum contract.
type WormholeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WormholeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WormholeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WormholeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WormholeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WormholeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WormholeSession struct {
	Contract     *Wormhole         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WormholeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WormholeCallerSession struct {
	Contract *WormholeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// WormholeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WormholeTransactorSession struct {
	Contract     *WormholeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// WormholeRaw is an auto generated low-level Go binding around an Ethereum contract.
type WormholeRaw struct {
	Contract *Wormhole // Generic contract binding to access the raw methods on
}

// WormholeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WormholeCallerRaw struct {
	Contract *WormholeCaller // Generic read-only contract binding to access the raw methods on
}

// WormholeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WormholeTransactorRaw struct {
	Contract *WormholeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWormhole creates a new instance of Wormhole, bound to a specific deployed contract.
func NewWormhole(address common.Address, backend bind.ContractBackend) (*Wormhole, error) {
	contract, err := bindWormhole(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Wormhole{WormholeCaller: WormholeCaller{contract: contract}, WormholeTransactor: WormholeTransactor{contract: contract}, WormholeFilterer: WormholeFilterer{contract: contract}}, nil
}

// NewWormholeCaller creates a new read-only instance of Wormhole, bound to a specific deployed contract.
func NewWormholeCaller(address common.Address, caller bind.ContractCaller) (*WormholeCaller, error) {
	contract, err := bindWormhole(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WormholeCaller{contract: contract}, nil
}

// NewWormholeTransactor creates a new write-only instance of Wormhole, bound to a specific deployed contract.
func NewWormholeTransactor(address common.Address, transactor bind.ContractTransactor) (*WormholeTransactor, error) {
	contract, err := bindWormhole(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WormholeTransactor{contract: contract}, nil
}

// NewWormholeFilterer creates a new log filterer instance of Wormhole, bound to a specific deployed contract.
func NewWormholeFilterer(address common.Address, filterer bind.ContractFilterer) (*WormholeFilterer, error) {
	contract, err := bindWormhole(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WormholeFilterer{contract: contract}, nil
}

// bindWormhole binds a generic wrapper to an already deployed contract.
func bindWormhole(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WormholeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wormhole *WormholeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wormhole.Contract.WormholeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wormhole *WormholeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wormhole.Contract.WormholeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wormhole *WormholeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wormhole.Contract.WormholeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Wormhole *WormholeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Wormhole.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Wormhole *WormholeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wormhole.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Wormhole *WormholeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Wormhole.Contract.contract.Transact(opts, method, params...)
}

// ConsumedVAAs is a free data retrieval call binding the contract method 0xa31fe409.
//
// Solidity: function consumedVAAs(bytes32 ) view returns(bool)
func (_Wormhole *WormholeCaller) ConsumedVAAs(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Wormhole.contract.Call(opts, &out, "consumedVAAs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ConsumedVAAs is a free data retrieval call binding the contract method 0xa31fe409.
//
// Solidity: function consumedVAAs(bytes32 ) view returns(bool)
func (_Wormhole *WormholeSession) ConsumedVAAs(arg0 [32]byte) (bool, error) {
	return _Wormhole.Contract.ConsumedVAAs(&_Wormhole.CallOpts, arg0)
}

// ConsumedVAAs is a free data retrieval call binding the contract method 0xa31fe409.
//
// Solidity: function consumedVAAs(bytes32 ) view returns(bool)
func (_Wormhole *WormholeCallerSession) ConsumedVAAs(arg0 [32]byte) (bool, error) {
	return _Wormhole.Contract.ConsumedVAAs(&_Wormhole.CallOpts, arg0)
}

// GetGuardianSet is a free data retrieval call binding the contract method 0xf951975a.
//
// Solidity: function getGuardianSet(uint32 idx) view returns((address[],uint32) gs)
func (_Wormhole *WormholeCaller) GetGuardianSet(opts *bind.CallOpts, idx uint32) (WormholeGuardianSet, error) {
	var out []interface{}
	err := _Wormhole.contract.Call(opts, &out, "getGuardianSet", idx)

	if err != nil {
		return *new(WormholeGuardianSet), err
	}

	out0 := *abi.ConvertType(out[0], new(WormholeGuardianSet)).(*WormholeGuardianSet)

	return out0, err

}

// GetGuardianSet is a free data retrieval call binding the contract method 0xf951975a.
//
// Solidity: function getGuardianSet(uint32 idx) view returns((address[],uint32) gs)
func (_Wormhole *WormholeSession) GetGuardianSet(idx uint32) (WormholeGuardianSet, error) {
	return _Wormhole.Contract.GetGuardianSet(&_Wormhole.CallOpts, idx)
}

// GetGuardianSet is a free data retrieval call binding the contract method 0xf951975a.
//
// Solidity: function getGuardianSet(uint32 idx) view returns((address[],uint32) gs)
func (_Wormhole *WormholeCallerSession) GetGuardianSet(idx uint32) (WormholeGuardianSet, error) {
	return _Wormhole.Contract.GetGuardianSet(&_Wormhole.CallOpts, idx)
}

// GuardianSetExpirity is a free data retrieval call binding the contract method 0x4db47840.
//
// Solidity: function guardian_set_expirity() view returns(uint32)
func (_Wormhole *WormholeCaller) GuardianSetExpirity(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Wormhole.contract.Call(opts, &out, "guardian_set_expirity")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GuardianSetExpirity is a free data retrieval call binding the contract method 0x4db47840.
//
// Solidity: function guardian_set_expirity() view returns(uint32)
func (_Wormhole *WormholeSession) GuardianSetExpirity() (uint32, error) {
	return _Wormhole.Contract.GuardianSetExpirity(&_Wormhole.CallOpts)
}

// GuardianSetExpirity is a free data retrieval call binding the contract method 0x4db47840.
//
// Solidity: function guardian_set_expirity() view returns(uint32)
func (_Wormhole *WormholeCallerSession) GuardianSetExpirity() (uint32, error) {
	return _Wormhole.Contract.GuardianSetExpirity(&_Wormhole.CallOpts)
}

// GuardianSetIndex is a free data retrieval call binding the contract method 0x822d82b3.
//
// Solidity: function guardian_set_index() view returns(uint32)
func (_Wormhole *WormholeCaller) GuardianSetIndex(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _Wormhole.contract.Call(opts, &out, "guardian_set_index")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GuardianSetIndex is a free data retrieval call binding the contract method 0x822d82b3.
//
// Solidity: function guardian_set_index() view returns(uint32)
func (_Wormhole *WormholeSession) GuardianSetIndex() (uint32, error) {
	return _Wormhole.Contract.GuardianSetIndex(&_Wormhole.CallOpts)
}

// GuardianSetIndex is a free data retrieval call binding the contract method 0x822d82b3.
//
// Solidity: function guardian_set_index() view returns(uint32)
func (_Wormhole *WormholeCallerSession) GuardianSetIndex() (uint32, error) {
	return _Wormhole.Contract.GuardianSetIndex(&_Wormhole.CallOpts)
}

// GuardianSets is a free data retrieval call binding the contract method 0x42b0aefa.
//
// Solidity: function guardian_sets(uint32 ) view returns(uint32 expiration_time)
func (_Wormhole *WormholeCaller) GuardianSets(opts *bind.CallOpts, arg0 uint32) (uint32, error) {
	var out []interface{}
	err := _Wormhole.contract.Call(opts, &out, "guardian_sets", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GuardianSets is a free data retrieval call binding the contract method 0x42b0aefa.
//
// Solidity: function guardian_sets(uint32 ) view returns(uint32 expiration_time)
func (_Wormhole *WormholeSession) GuardianSets(arg0 uint32) (uint32, error) {
	return _Wormhole.Contract.GuardianSets(&_Wormhole.CallOpts, arg0)
}

// GuardianSets is a free data retrieval call binding the contract method 0x42b0aefa.
//
// Solidity: function guardian_sets(uint32 ) view returns(uint32 expiration_time)
func (_Wormhole *WormholeCallerSession) GuardianSets(arg0 uint32) (uint32, error) {
	return _Wormhole.Contract.GuardianSets(&_Wormhole.CallOpts, arg0)
}

// IsWrappedAsset is a free data retrieval call binding the contract method 0x1a2be4da.
//
// Solidity: function isWrappedAsset(address ) view returns(bool)
func (_Wormhole *WormholeCaller) IsWrappedAsset(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Wormhole.contract.Call(opts, &out, "isWrappedAsset", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsWrappedAsset is a free data retrieval call binding the contract method 0x1a2be4da.
//
// Solidity: function isWrappedAsset(address ) view returns(bool)
func (_Wormhole *WormholeSession) IsWrappedAsset(arg0 common.Address) (bool, error) {
	return _Wormhole.Contract.IsWrappedAsset(&_Wormhole.CallOpts, arg0)
}

// IsWrappedAsset is a free data retrieval call binding the contract method 0x1a2be4da.
//
// Solidity: function isWrappedAsset(address ) view returns(bool)
func (_Wormhole *WormholeCallerSession) IsWrappedAsset(arg0 common.Address) (bool, error) {
	return _Wormhole.Contract.IsWrappedAsset(&_Wormhole.CallOpts, arg0)
}

// ParseAndVerifyVAA is a free data retrieval call binding the contract method 0x600b9aa6.
//
// Solidity: function parseAndVerifyVAA(bytes vaa) view returns((uint8,bytes32,uint32,uint32,uint8,bytes) parsed_vaa)
func (_Wormhole *WormholeCaller) ParseAndVerifyVAA(opts *bind.CallOpts, vaa []byte) (WormholeParsedVAA, error) {
	var out []interface{}
	err := _Wormhole.contract.Call(opts, &out, "parseAndVerifyVAA", vaa)

	if err != nil {
		return *new(WormholeParsedVAA), err
	}

	out0 := *abi.ConvertType(out[0], new(WormholeParsedVAA)).(*WormholeParsedVAA)

	return out0, err

}

// ParseAndVerifyVAA is a free data retrieval call binding the contract method 0x600b9aa6.
//
// Solidity: function parseAndVerifyVAA(bytes vaa) view returns((uint8,bytes32,uint32,uint32,uint8,bytes) parsed_vaa)
func (_Wormhole *WormholeSession) ParseAndVerifyVAA(vaa []byte) (WormholeParsedVAA, error) {
	return _Wormhole.Contract.ParseAndVerifyVAA(&_Wormhole.CallOpts, vaa)
}

// ParseAndVerifyVAA is a free data retrieval call binding the contract method 0x600b9aa6.
//
// Solidity: function parseAndVerifyVAA(bytes vaa) view returns((uint8,bytes32,uint32,uint32,uint8,bytes) parsed_vaa)
func (_Wormhole *WormholeCallerSession) ParseAndVerifyVAA(vaa []byte) (WormholeParsedVAA, error) {
	return _Wormhole.Contract.ParseAndVerifyVAA(&_Wormhole.CallOpts, vaa)
}

// WrappedAssetMaster is a free data retrieval call binding the contract method 0x99da1d3c.
//
// Solidity: function wrappedAssetMaster() view returns(address)
func (_Wormhole *WormholeCaller) WrappedAssetMaster(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Wormhole.contract.Call(opts, &out, "wrappedAssetMaster")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedAssetMaster is a free data retrieval call binding the contract method 0x99da1d3c.
//
// Solidity: function wrappedAssetMaster() view returns(address)
func (_Wormhole *WormholeSession) WrappedAssetMaster() (common.Address, error) {
	return _Wormhole.Contract.WrappedAssetMaster(&_Wormhole.CallOpts)
}

// WrappedAssetMaster is a free data retrieval call binding the contract method 0x99da1d3c.
//
// Solidity: function wrappedAssetMaster() view returns(address)
func (_Wormhole *WormholeCallerSession) WrappedAssetMaster() (common.Address, error) {
	return _Wormhole.Contract.WrappedAssetMaster(&_Wormhole.CallOpts)
}

// WrappedAssets is a free data retrieval call binding the contract method 0xb6694c2a.
//
// Solidity: function wrappedAssets(bytes32 ) view returns(address)
func (_Wormhole *WormholeCaller) WrappedAssets(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Wormhole.contract.Call(opts, &out, "wrappedAssets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WrappedAssets is a free data retrieval call binding the contract method 0xb6694c2a.
//
// Solidity: function wrappedAssets(bytes32 ) view returns(address)
func (_Wormhole *WormholeSession) WrappedAssets(arg0 [32]byte) (common.Address, error) {
	return _Wormhole.Contract.WrappedAssets(&_Wormhole.CallOpts, arg0)
}

// WrappedAssets is a free data retrieval call binding the contract method 0xb6694c2a.
//
// Solidity: function wrappedAssets(bytes32 ) view returns(address)
func (_Wormhole *WormholeCallerSession) WrappedAssets(arg0 [32]byte) (common.Address, error) {
	return _Wormhole.Contract.WrappedAssets(&_Wormhole.CallOpts, arg0)
}

// LockAssets is a paid mutator transaction binding the contract method 0x70713960.
//
// Solidity: function lockAssets(address asset, uint256 amount, bytes32 recipient, uint8 target_chain, uint32 nonce, bool refund_dust) returns()
func (_Wormhole *WormholeTransactor) LockAssets(opts *bind.TransactOpts, asset common.Address, amount *big.Int, recipient [32]byte, target_chain uint8, nonce uint32, refund_dust bool) (*types.Transaction, error) {
	return _Wormhole.contract.Transact(opts, "lockAssets", asset, amount, recipient, target_chain, nonce, refund_dust)
}

// LockAssets is a paid mutator transaction binding the contract method 0x70713960.
//
// Solidity: function lockAssets(address asset, uint256 amount, bytes32 recipient, uint8 target_chain, uint32 nonce, bool refund_dust) returns()
func (_Wormhole *WormholeSession) LockAssets(asset common.Address, amount *big.Int, recipient [32]byte, target_chain uint8, nonce uint32, refund_dust bool) (*types.Transaction, error) {
	return _Wormhole.Contract.LockAssets(&_Wormhole.TransactOpts, asset, amount, recipient, target_chain, nonce, refund_dust)
}

// LockAssets is a paid mutator transaction binding the contract method 0x70713960.
//
// Solidity: function lockAssets(address asset, uint256 amount, bytes32 recipient, uint8 target_chain, uint32 nonce, bool refund_dust) returns()
func (_Wormhole *WormholeTransactorSession) LockAssets(asset common.Address, amount *big.Int, recipient [32]byte, target_chain uint8, nonce uint32, refund_dust bool) (*types.Transaction, error) {
	return _Wormhole.Contract.LockAssets(&_Wormhole.TransactOpts, asset, amount, recipient, target_chain, nonce, refund_dust)
}

// LockETH is a paid mutator transaction binding the contract method 0x58d62e46.
//
// Solidity: function lockETH(bytes32 recipient, uint8 target_chain, uint32 nonce) payable returns()
func (_Wormhole *WormholeTransactor) LockETH(opts *bind.TransactOpts, recipient [32]byte, target_chain uint8, nonce uint32) (*types.Transaction, error) {
	return _Wormhole.contract.Transact(opts, "lockETH", recipient, target_chain, nonce)
}

// LockETH is a paid mutator transaction binding the contract method 0x58d62e46.
//
// Solidity: function lockETH(bytes32 recipient, uint8 target_chain, uint32 nonce) payable returns()
func (_Wormhole *WormholeSession) LockETH(recipient [32]byte, target_chain uint8, nonce uint32) (*types.Transaction, error) {
	return _Wormhole.Contract.LockETH(&_Wormhole.TransactOpts, recipient, target_chain, nonce)
}

// LockETH is a paid mutator transaction binding the contract method 0x58d62e46.
//
// Solidity: function lockETH(bytes32 recipient, uint8 target_chain, uint32 nonce) payable returns()
func (_Wormhole *WormholeTransactorSession) LockETH(recipient [32]byte, target_chain uint8, nonce uint32) (*types.Transaction, error) {
	return _Wormhole.Contract.LockETH(&_Wormhole.TransactOpts, recipient, target_chain, nonce)
}

// SubmitVAA is a paid mutator transaction binding the contract method 0x3bc0aee6.
//
// Solidity: function submitVAA(bytes vaa) returns()
func (_Wormhole *WormholeTransactor) SubmitVAA(opts *bind.TransactOpts, vaa []byte) (*types.Transaction, error) {
	return _Wormhole.contract.Transact(opts, "submitVAA", vaa)
}

// SubmitVAA is a paid mutator transaction binding the contract method 0x3bc0aee6.
//
// Solidity: function submitVAA(bytes vaa) returns()
func (_Wormhole *WormholeSession) SubmitVAA(vaa []byte) (*types.Transaction, error) {
	return _Wormhole.Contract.SubmitVAA(&_Wormhole.TransactOpts, vaa)
}

// SubmitVAA is a paid mutator transaction binding the contract method 0x3bc0aee6.
//
// Solidity: function submitVAA(bytes vaa) returns()
func (_Wormhole *WormholeTransactorSession) SubmitVAA(vaa []byte) (*types.Transaction, error) {
	return _Wormhole.Contract.SubmitVAA(&_Wormhole.TransactOpts, vaa)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Wormhole *WormholeTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Wormhole.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Wormhole *WormholeSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Wormhole.Contract.Fallback(&_Wormhole.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Wormhole *WormholeTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Wormhole.Contract.Fallback(&_Wormhole.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Wormhole *WormholeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Wormhole.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Wormhole *WormholeSession) Receive() (*types.Transaction, error) {
	return _Wormhole.Contract.Receive(&_Wormhole.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Wormhole *WormholeTransactorSession) Receive() (*types.Transaction, error) {
	return _Wormhole.Contract.Receive(&_Wormhole.TransactOpts)
}

// WormholeLogGuardianSetChangedIterator is returned from FilterLogGuardianSetChanged and is used to iterate over the raw logs and unpacked data for LogGuardianSetChanged events raised by the Wormhole contract.
type WormholeLogGuardianSetChangedIterator struct {
	Event *WormholeLogGuardianSetChanged // Event containing the contract specifics and raw log

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
func (it *WormholeLogGuardianSetChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WormholeLogGuardianSetChanged)
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
		it.Event = new(WormholeLogGuardianSetChanged)
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
func (it *WormholeLogGuardianSetChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WormholeLogGuardianSetChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WormholeLogGuardianSetChanged represents a LogGuardianSetChanged event raised by the Wormhole contract.
type WormholeLogGuardianSetChanged struct {
	OldGuardianIndex uint32
	NewGuardianIndex uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLogGuardianSetChanged is a free log retrieval operation binding the contract event 0xdfb80683934199683861bf00b64ecdf0984bbaf661bf27983dba382e99297a62.
//
// Solidity: event LogGuardianSetChanged(uint32 oldGuardianIndex, uint32 newGuardianIndex)
func (_Wormhole *WormholeFilterer) FilterLogGuardianSetChanged(opts *bind.FilterOpts) (*WormholeLogGuardianSetChangedIterator, error) {

	logs, sub, err := _Wormhole.contract.FilterLogs(opts, "LogGuardianSetChanged")
	if err != nil {
		return nil, err
	}
	return &WormholeLogGuardianSetChangedIterator{contract: _Wormhole.contract, event: "LogGuardianSetChanged", logs: logs, sub: sub}, nil
}

// WatchLogGuardianSetChanged is a free log subscription operation binding the contract event 0xdfb80683934199683861bf00b64ecdf0984bbaf661bf27983dba382e99297a62.
//
// Solidity: event LogGuardianSetChanged(uint32 oldGuardianIndex, uint32 newGuardianIndex)
func (_Wormhole *WormholeFilterer) WatchLogGuardianSetChanged(opts *bind.WatchOpts, sink chan<- *WormholeLogGuardianSetChanged) (event.Subscription, error) {

	logs, sub, err := _Wormhole.contract.WatchLogs(opts, "LogGuardianSetChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WormholeLogGuardianSetChanged)
				if err := _Wormhole.contract.UnpackLog(event, "LogGuardianSetChanged", log); err != nil {
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

// ParseLogGuardianSetChanged is a log parse operation binding the contract event 0xdfb80683934199683861bf00b64ecdf0984bbaf661bf27983dba382e99297a62.
//
// Solidity: event LogGuardianSetChanged(uint32 oldGuardianIndex, uint32 newGuardianIndex)
func (_Wormhole *WormholeFilterer) ParseLogGuardianSetChanged(log types.Log) (*WormholeLogGuardianSetChanged, error) {
	event := new(WormholeLogGuardianSetChanged)
	if err := _Wormhole.contract.UnpackLog(event, "LogGuardianSetChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WormholeLogTokensLockedIterator is returned from FilterLogTokensLocked and is used to iterate over the raw logs and unpacked data for LogTokensLocked events raised by the Wormhole contract.
type WormholeLogTokensLockedIterator struct {
	Event *WormholeLogTokensLocked // Event containing the contract specifics and raw log

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
func (it *WormholeLogTokensLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WormholeLogTokensLocked)
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
		it.Event = new(WormholeLogTokensLocked)
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
func (it *WormholeLogTokensLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WormholeLogTokensLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WormholeLogTokensLocked represents a LogTokensLocked event raised by the Wormhole contract.
type WormholeLogTokensLocked struct {
	TargetChain   uint8
	TokenChain    uint8
	TokenDecimals uint8
	Token         [32]byte
	Sender        [32]byte
	Recipient     [32]byte
	Amount        *big.Int
	Nonce         uint32
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLogTokensLocked is a free log retrieval operation binding the contract event 0x6bbd554ad75919f71fd91bf917ca6e4f41c10f03ab25751596a22253bb39aab8.
//
// Solidity: event LogTokensLocked(uint8 target_chain, uint8 token_chain, uint8 token_decimals, bytes32 indexed token, bytes32 indexed sender, bytes32 recipient, uint256 amount, uint32 nonce)
func (_Wormhole *WormholeFilterer) FilterLogTokensLocked(opts *bind.FilterOpts, token [][32]byte, sender [][32]byte) (*WormholeLogTokensLockedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Wormhole.contract.FilterLogs(opts, "LogTokensLocked", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &WormholeLogTokensLockedIterator{contract: _Wormhole.contract, event: "LogTokensLocked", logs: logs, sub: sub}, nil
}

// WatchLogTokensLocked is a free log subscription operation binding the contract event 0x6bbd554ad75919f71fd91bf917ca6e4f41c10f03ab25751596a22253bb39aab8.
//
// Solidity: event LogTokensLocked(uint8 target_chain, uint8 token_chain, uint8 token_decimals, bytes32 indexed token, bytes32 indexed sender, bytes32 recipient, uint256 amount, uint32 nonce)
func (_Wormhole *WormholeFilterer) WatchLogTokensLocked(opts *bind.WatchOpts, sink chan<- *WormholeLogTokensLocked, token [][32]byte, sender [][32]byte) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Wormhole.contract.WatchLogs(opts, "LogTokensLocked", tokenRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WormholeLogTokensLocked)
				if err := _Wormhole.contract.UnpackLog(event, "LogTokensLocked", log); err != nil {
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

// ParseLogTokensLocked is a log parse operation binding the contract event 0x6bbd554ad75919f71fd91bf917ca6e4f41c10f03ab25751596a22253bb39aab8.
//
// Solidity: event LogTokensLocked(uint8 target_chain, uint8 token_chain, uint8 token_decimals, bytes32 indexed token, bytes32 indexed sender, bytes32 recipient, uint256 amount, uint32 nonce)
func (_Wormhole *WormholeFilterer) ParseLogTokensLocked(log types.Log) (*WormholeLogTokensLocked, error) {
	event := new(WormholeLogTokensLocked)
	if err := _Wormhole.contract.UnpackLog(event, "LogTokensLocked", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WrappedAssetABI is the input ABI used to generate the binding from.
const WrappedAssetABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetAddress\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"assetChain\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridge\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"_assetChain\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_assetAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// WrappedAssetFuncSigs maps the 4-byte function signature to its string representation.
var WrappedAssetFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"1ba46cfd": "assetAddress()",
	"026b0539": "assetChain()",
	"70a08231": "balanceOf(address)",
	"e78cea92": "bridge()",
	"9dc29fac": "burn(address,uint256)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"a7a2d3fb": "initialize(uint8,bytes32,uint8)",
	"158ef93e": "initialized()",
	"40c10f19": "mint(address,uint256)",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// WrappedAssetBin is the compiled bytecode used for deploying new contracts.
var WrappedAssetBin = "0x60806040526007805460ff1916601217905534801561001d57600080fd5b5061124a8061002d6000396000f3fe608060405234801561001057600080fd5b50600436106101165760003560e01c806340c10f19116100a2578063a457c2d711610071578063a457c2d714610204578063a7a2d3fb14610217578063a9059cbb1461022a578063dd62ed3e1461023d578063e78cea921461025057610116565b806340c10f19146101c157806370a08231146101d657806395d89b41146101e95780639dc29fac146101f157610116565b806318160ddd116100e957806318160ddd146101765780631ba46cfd1461018b57806323b872dd14610193578063313ce567146101a657806339509351146101ae57610116565b8063026b05391461011b57806306fdde0314610139578063095ea7b31461014e578063158ef93e1461016e575b600080fd5b610123610265565b6040516101309190611129565b60405180910390f35b61014161026e565b6040516101309190610e82565b61016161015c366004610d97565b6102ae565b6040516101309190610e6e565b6101616102cc565b61017e6102d5565b6040516101309190610e79565b61017e6102db565b6101616101a1366004610d57565b6102e1565b61012361036e565b6101616101bc366004610d97565b610377565b6101d46101cf366004610d97565b6103cb565b005b61017e6101e4366004610d08565b610411565b610141610430565b6101d46101ff366004610d97565b6104c6565b610161610212366004610d97565b6104ff565b6101d4610225366004610dc1565b61056d565b610161610238366004610d97565b610605565b61017e61024b366004610d23565b610619565b610258610644565b6040516101309190610e5a565b60005460ff1681565b6000546060906102809060ff16610658565b610288610730565b604051602001610299929190610dfe565b60405160208183030381529060405290505b90565b60006102c26102bb61084a565b848461084e565b5060015b92915050565b60025460ff1681565b60055490565b60015481565b60006102ee848484610902565b610364846102fa61084a565b61035f856040518060600160405280602881526020016111c8602891396001600160a01b038a1660009081526004602052604081209061033861084a565b6001600160a01b03168152602081019190915260400160002054919063ffffffff610a1816565b61084e565b5060019392505050565b60075460ff1690565b60006102c261038461084a565b8461035f856004600061039561084a565b6001600160a01b03908116825260208083019390935260409182016000908120918c16815292529020549063ffffffff610a4416565b60025461010090046001600160a01b031633146104035760405162461bcd60e51b81526004016103fa906110ad565b60405180910390fd5b61040d8282610a70565b5050565b6001600160a01b0381166000908152600360205260409020545b919050565b60068054604080516020601f60026000196101006001881615020190951694909404938401819004810282018101909252828152606093909290918301828280156104bc5780601f10610491576101008083540402835291602001916104bc565b820191906000526020600020905b81548152906001019060200180831161049f57829003601f168201915b5050505050905090565b60025461010090046001600160a01b031633146104f55760405162461bcd60e51b81526004016103fa90611024565b61040d8282610b30565b60006102c261050c61084a565b8461035f856040518060600160405280602581526020016111f0602591396004600061053661084a565b6001600160a01b03908116825260208083019390935260409182016000908120918d1681529252902054919063ffffffff610a1816565b60025460ff16156105905760405162461bcd60e51b81526004016103fa90610f71565b6000805460ff1990811660ff861617909155600183815560028054610100600160a81b0319166101003302179092161790556040805180820190915260038082526215d5d560ea1b60209092019182526105ec91600691610c48565b506007805460ff191660ff929092169190911790555050565b60006102c261061261084a565b8484610902565b6001600160a01b03918216600090815260046020908152604080832093909416825291909152205490565b60025461010090046001600160a01b031681565b60608161067d57506040805180820190915260018152600360fc1b602082015261042b565b8160005b811561069557600101600a82049150610681565b60608167ffffffffffffffff811180156106ae57600080fd5b506040519080825280601f01601f1916602001820160405280156106d9576020820181803683370190505b50905060001982015b851561072757600a860660300160f81b8282806001900393508151811061070557fe5b60200101906001600160f81b031916908160001a905350600a860495506106e2565b50949350505050565b604080518082018252601081526f181899199a1a9b1b9c1cb0b131b232b360811b6020820152600154825160428082526080820190945260609384919060208201818036833701905050905060005b60208110156108425783600484836020811061079757fe5b1a60f81b6001600160f81b031916901c60f81c60ff16815181106107b757fe5b602001015160f81c60f81b8282600202815181106107d157fe5b60200101906001600160f81b031916908160001a905350838382602081106107f557fe5b825191901a600f1690811061080657fe5b602001015160f81c60f81b82826002026001018151811061082357fe5b60200101906001600160f81b031916908160001a90535060010161077f565b509250505090565b3390565b6001600160a01b0383166108745760405162461bcd60e51b81526004016103fa90611069565b6001600160a01b03821661089a5760405162461bcd60e51b81526004016103fa90610ef8565b6001600160a01b0380841660008181526004602090815260408083209487168084529490915290819020849055517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925906108f5908590610e79565b60405180910390a3505050565b6001600160a01b0383166109285760405162461bcd60e51b81526004016103fa90610fdf565b6001600160a01b03821661094e5760405162461bcd60e51b81526004016103fa90610eb5565b610991816040518060600160405280602681526020016111a2602691396001600160a01b038616600090815260036020526040902054919063ffffffff610a1816565b6001600160a01b0380851660009081526003602052604080822093909355908416815220546109c6908263ffffffff610a4416565b6001600160a01b0380841660008181526003602052604090819020939093559151908516907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef906108f5908590610e79565b60008184841115610a3c5760405162461bcd60e51b81526004016103fa9190610e82565b505050900390565b600082820183811015610a695760405162461bcd60e51b81526004016103fa90610f3a565b9392505050565b6001600160a01b038216610a965760405162461bcd60e51b81526004016103fa906110f2565b600554610aa9908263ffffffff610a4416565b6005556001600160a01b038216600090815260036020526040902054610ad5908263ffffffff610a4416565b6001600160a01b0383166000818152600360205260408082209390935591519091907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90610b24908590610e79565b60405180910390a35050565b6001600160a01b038216610b565760405162461bcd60e51b81526004016103fa90610f9e565b610b9981604051806060016040528060228152602001611180602291396001600160a01b038516600090815260036020526040902054919063ffffffff610a1816565b6001600160a01b038316600090815260036020526040902055600554610bc5908263ffffffff610c0616565b6005556040516000906001600160a01b038416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef90610b24908590610e79565b6000610a6983836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610a18565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610c8957805160ff1916838001178555610cb6565b82800160010185558215610cb6579182015b82811115610cb6578251825591602001919060010190610c9b565b50610cc2929150610cc6565b5090565b6102ab91905b80821115610cc25760008155600101610ccc565b80356001600160a01b03811681146102c657600080fd5b803560ff811681146102c657600080fd5b600060208284031215610d19578081fd5b610a698383610ce0565b60008060408385031215610d35578081fd5b610d3f8484610ce0565b9150610d4e8460208501610ce0565b90509250929050565b600080600060608486031215610d6b578081fd5b8335610d7681611167565b92506020840135610d8681611167565b929592945050506040919091013590565b60008060408385031215610da9578182fd5b610db38484610ce0565b946020939093013593505050565b600080600060608486031215610dd5578283fd5b610ddf8585610cf7565b925060208401359150610df58560408601610cf7565b90509250925092565b60007202bb7b936b437b632902bb930b83832b210169606d1b82528351610e2c816013850160208801611137565b808301602d60f81b601382015284519150610e4e826014830160208801611137565b01601401949350505050565b6001600160a01b0391909116815260200190565b901515815260200190565b90815260200190565b6000602082528251806020840152610ea1816040850160208701611137565b601f01601f19169190910160400192915050565b60208082526023908201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260408201526265737360e81b606082015260800190565b60208082526022908201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604082015261737360f01b606082015260800190565b6020808252601b908201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604082015260600190565b602080825260139082015272185b1c9958591e481a5b9a5d1a585b1a5e9959606a1b604082015260600190565b60208082526021908201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736040820152607360f81b606082015260800190565b60208082526025908201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604082015264647265737360d81b606082015260800190565b60208082526025908201527f6275726e2063616e206f6e6c792062652063616c6c6564206279207468652062604082015264726964676560d81b606082015260800190565b60208082526024908201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646040820152637265737360e01b606082015260800190565b60208082526025908201527f6d696e742063616e206f6e6c792062652063616c6c6564206279207468652062604082015264726964676560d81b606082015260800190565b6020808252601f908201527f45524332303a206d696e7420746f20746865207a65726f206164647265737300604082015260600190565b60ff91909116815260200190565b60005b8381101561115257818101518382015260200161113a565b83811115611161576000848401525b50505050565b6001600160a01b038116811461117c57600080fd5b5056fe45524332303a206275726e20616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220aba6d88d0970b39f338df3ac25bf269be2bc913c5de9e35d40ceaf797625a83164736f6c634300060a0033"

// DeployWrappedAsset deploys a new Ethereum contract, binding an instance of WrappedAsset to it.
func DeployWrappedAsset(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *WrappedAsset, error) {
	parsed, err := abi.JSON(strings.NewReader(WrappedAssetABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WrappedAssetBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WrappedAsset{WrappedAssetCaller: WrappedAssetCaller{contract: contract}, WrappedAssetTransactor: WrappedAssetTransactor{contract: contract}, WrappedAssetFilterer: WrappedAssetFilterer{contract: contract}}, nil
}

// WrappedAsset is an auto generated Go binding around an Ethereum contract.
type WrappedAsset struct {
	WrappedAssetCaller     // Read-only binding to the contract
	WrappedAssetTransactor // Write-only binding to the contract
	WrappedAssetFilterer   // Log filterer for contract events
}

// WrappedAssetCaller is an auto generated read-only Go binding around an Ethereum contract.
type WrappedAssetCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedAssetTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WrappedAssetTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedAssetFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WrappedAssetFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WrappedAssetSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WrappedAssetSession struct {
	Contract     *WrappedAsset     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WrappedAssetCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WrappedAssetCallerSession struct {
	Contract *WrappedAssetCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// WrappedAssetTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WrappedAssetTransactorSession struct {
	Contract     *WrappedAssetTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// WrappedAssetRaw is an auto generated low-level Go binding around an Ethereum contract.
type WrappedAssetRaw struct {
	Contract *WrappedAsset // Generic contract binding to access the raw methods on
}

// WrappedAssetCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WrappedAssetCallerRaw struct {
	Contract *WrappedAssetCaller // Generic read-only contract binding to access the raw methods on
}

// WrappedAssetTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WrappedAssetTransactorRaw struct {
	Contract *WrappedAssetTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWrappedAsset creates a new instance of WrappedAsset, bound to a specific deployed contract.
func NewWrappedAsset(address common.Address, backend bind.ContractBackend) (*WrappedAsset, error) {
	contract, err := bindWrappedAsset(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WrappedAsset{WrappedAssetCaller: WrappedAssetCaller{contract: contract}, WrappedAssetTransactor: WrappedAssetTransactor{contract: contract}, WrappedAssetFilterer: WrappedAssetFilterer{contract: contract}}, nil
}

// NewWrappedAssetCaller creates a new read-only instance of WrappedAsset, bound to a specific deployed contract.
func NewWrappedAssetCaller(address common.Address, caller bind.ContractCaller) (*WrappedAssetCaller, error) {
	contract, err := bindWrappedAsset(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedAssetCaller{contract: contract}, nil
}

// NewWrappedAssetTransactor creates a new write-only instance of WrappedAsset, bound to a specific deployed contract.
func NewWrappedAssetTransactor(address common.Address, transactor bind.ContractTransactor) (*WrappedAssetTransactor, error) {
	contract, err := bindWrappedAsset(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WrappedAssetTransactor{contract: contract}, nil
}

// NewWrappedAssetFilterer creates a new log filterer instance of WrappedAsset, bound to a specific deployed contract.
func NewWrappedAssetFilterer(address common.Address, filterer bind.ContractFilterer) (*WrappedAssetFilterer, error) {
	contract, err := bindWrappedAsset(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WrappedAssetFilterer{contract: contract}, nil
}

// bindWrappedAsset binds a generic wrapper to an already deployed contract.
func bindWrappedAsset(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WrappedAssetABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedAsset *WrappedAssetRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedAsset.Contract.WrappedAssetCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedAsset *WrappedAssetRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedAsset.Contract.WrappedAssetTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedAsset *WrappedAssetRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedAsset.Contract.WrappedAssetTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WrappedAsset *WrappedAssetCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WrappedAsset.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WrappedAsset *WrappedAssetTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WrappedAsset.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WrappedAsset *WrappedAssetTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WrappedAsset.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedAsset *WrappedAssetCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedAsset.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedAsset *WrappedAssetSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrappedAsset.Contract.Allowance(&_WrappedAsset.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_WrappedAsset *WrappedAssetCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _WrappedAsset.Contract.Allowance(&_WrappedAsset.CallOpts, owner, spender)
}

// AssetAddress is a free data retrieval call binding the contract method 0x1ba46cfd.
//
// Solidity: function assetAddress() view returns(bytes32)
func (_WrappedAsset *WrappedAssetCaller) AssetAddress(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _WrappedAsset.contract.Call(opts, &out, "assetAddress")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AssetAddress is a free data retrieval call binding the contract method 0x1ba46cfd.
//
// Solidity: function assetAddress() view returns(bytes32)
func (_WrappedAsset *WrappedAssetSession) AssetAddress() ([32]byte, error) {
	return _WrappedAsset.Contract.AssetAddress(&_WrappedAsset.CallOpts)
}

// AssetAddress is a free data retrieval call binding the contract method 0x1ba46cfd.
//
// Solidity: function assetAddress() view returns(bytes32)
func (_WrappedAsset *WrappedAssetCallerSession) AssetAddress() ([32]byte, error) {
	return _WrappedAsset.Contract.AssetAddress(&_WrappedAsset.CallOpts)
}

// AssetChain is a free data retrieval call binding the contract method 0x026b0539.
//
// Solidity: function assetChain() view returns(uint8)
func (_WrappedAsset *WrappedAssetCaller) AssetChain(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WrappedAsset.contract.Call(opts, &out, "assetChain")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// AssetChain is a free data retrieval call binding the contract method 0x026b0539.
//
// Solidity: function assetChain() view returns(uint8)
func (_WrappedAsset *WrappedAssetSession) AssetChain() (uint8, error) {
	return _WrappedAsset.Contract.AssetChain(&_WrappedAsset.CallOpts)
}

// AssetChain is a free data retrieval call binding the contract method 0x026b0539.
//
// Solidity: function assetChain() view returns(uint8)
func (_WrappedAsset *WrappedAssetCallerSession) AssetChain() (uint8, error) {
	return _WrappedAsset.Contract.AssetChain(&_WrappedAsset.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedAsset *WrappedAssetCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WrappedAsset.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedAsset *WrappedAssetSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrappedAsset.Contract.BalanceOf(&_WrappedAsset.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_WrappedAsset *WrappedAssetCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _WrappedAsset.Contract.BalanceOf(&_WrappedAsset.CallOpts, account)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_WrappedAsset *WrappedAssetCaller) Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WrappedAsset.contract.Call(opts, &out, "bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_WrappedAsset *WrappedAssetSession) Bridge() (common.Address, error) {
	return _WrappedAsset.Contract.Bridge(&_WrappedAsset.CallOpts)
}

// Bridge is a free data retrieval call binding the contract method 0xe78cea92.
//
// Solidity: function bridge() view returns(address)
func (_WrappedAsset *WrappedAssetCallerSession) Bridge() (common.Address, error) {
	return _WrappedAsset.Contract.Bridge(&_WrappedAsset.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedAsset *WrappedAssetCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _WrappedAsset.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedAsset *WrappedAssetSession) Decimals() (uint8, error) {
	return _WrappedAsset.Contract.Decimals(&_WrappedAsset.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WrappedAsset *WrappedAssetCallerSession) Decimals() (uint8, error) {
	return _WrappedAsset.Contract.Decimals(&_WrappedAsset.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_WrappedAsset *WrappedAssetCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _WrappedAsset.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_WrappedAsset *WrappedAssetSession) Initialized() (bool, error) {
	return _WrappedAsset.Contract.Initialized(&_WrappedAsset.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_WrappedAsset *WrappedAssetCallerSession) Initialized() (bool, error) {
	return _WrappedAsset.Contract.Initialized(&_WrappedAsset.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedAsset *WrappedAssetCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrappedAsset.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedAsset *WrappedAssetSession) Name() (string, error) {
	return _WrappedAsset.Contract.Name(&_WrappedAsset.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WrappedAsset *WrappedAssetCallerSession) Name() (string, error) {
	return _WrappedAsset.Contract.Name(&_WrappedAsset.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedAsset *WrappedAssetCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _WrappedAsset.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedAsset *WrappedAssetSession) Symbol() (string, error) {
	return _WrappedAsset.Contract.Symbol(&_WrappedAsset.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WrappedAsset *WrappedAssetCallerSession) Symbol() (string, error) {
	return _WrappedAsset.Contract.Symbol(&_WrappedAsset.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedAsset *WrappedAssetCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WrappedAsset.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedAsset *WrappedAssetSession) TotalSupply() (*big.Int, error) {
	return _WrappedAsset.Contract.TotalSupply(&_WrappedAsset.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WrappedAsset *WrappedAssetCallerSession) TotalSupply() (*big.Int, error) {
	return _WrappedAsset.Contract.TotalSupply(&_WrappedAsset.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrappedAsset *WrappedAssetTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrappedAsset *WrappedAssetSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.Contract.Approve(&_WrappedAsset.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_WrappedAsset *WrappedAssetTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.Contract.Approve(&_WrappedAsset.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_WrappedAsset *WrappedAssetTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_WrappedAsset *WrappedAssetSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.Contract.Burn(&_WrappedAsset.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_WrappedAsset *WrappedAssetTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.Contract.Burn(&_WrappedAsset.TransactOpts, account, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WrappedAsset *WrappedAssetTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WrappedAsset *WrappedAssetSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.Contract.DecreaseAllowance(&_WrappedAsset.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_WrappedAsset *WrappedAssetTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.Contract.DecreaseAllowance(&_WrappedAsset.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WrappedAsset *WrappedAssetTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WrappedAsset *WrappedAssetSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.Contract.IncreaseAllowance(&_WrappedAsset.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_WrappedAsset *WrappedAssetTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.Contract.IncreaseAllowance(&_WrappedAsset.TransactOpts, spender, addedValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xa7a2d3fb.
//
// Solidity: function initialize(uint8 _assetChain, bytes32 _assetAddress, uint8 decimals) returns()
func (_WrappedAsset *WrappedAssetTransactor) Initialize(opts *bind.TransactOpts, _assetChain uint8, _assetAddress [32]byte, decimals uint8) (*types.Transaction, error) {
	return _WrappedAsset.contract.Transact(opts, "initialize", _assetChain, _assetAddress, decimals)
}

// Initialize is a paid mutator transaction binding the contract method 0xa7a2d3fb.
//
// Solidity: function initialize(uint8 _assetChain, bytes32 _assetAddress, uint8 decimals) returns()
func (_WrappedAsset *WrappedAssetSession) Initialize(_assetChain uint8, _assetAddress [32]byte, decimals uint8) (*types.Transaction, error) {
	return _WrappedAsset.Contract.Initialize(&_WrappedAsset.TransactOpts, _assetChain, _assetAddress, decimals)
}

// Initialize is a paid mutator transaction binding the contract method 0xa7a2d3fb.
//
// Solidity: function initialize(uint8 _assetChain, bytes32 _assetAddress, uint8 decimals) returns()
func (_WrappedAsset *WrappedAssetTransactorSession) Initialize(_assetChain uint8, _assetAddress [32]byte, decimals uint8) (*types.Transaction, error) {
	return _WrappedAsset.Contract.Initialize(&_WrappedAsset.TransactOpts, _assetChain, _assetAddress, decimals)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_WrappedAsset *WrappedAssetTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_WrappedAsset *WrappedAssetSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.Contract.Mint(&_WrappedAsset.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_WrappedAsset *WrappedAssetTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.Contract.Mint(&_WrappedAsset.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WrappedAsset *WrappedAssetTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WrappedAsset *WrappedAssetSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.Contract.Transfer(&_WrappedAsset.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_WrappedAsset *WrappedAssetTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.Contract.Transfer(&_WrappedAsset.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WrappedAsset *WrappedAssetTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WrappedAsset *WrappedAssetSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.Contract.TransferFrom(&_WrappedAsset.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_WrappedAsset *WrappedAssetTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WrappedAsset.Contract.TransferFrom(&_WrappedAsset.TransactOpts, sender, recipient, amount)
}

// WrappedAssetApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WrappedAsset contract.
type WrappedAssetApprovalIterator struct {
	Event *WrappedAssetApproval // Event containing the contract specifics and raw log

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
func (it *WrappedAssetApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedAssetApproval)
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
		it.Event = new(WrappedAssetApproval)
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
func (it *WrappedAssetApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedAssetApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedAssetApproval represents a Approval event raised by the WrappedAsset contract.
type WrappedAssetApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedAsset *WrappedAssetFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WrappedAssetApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrappedAsset.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WrappedAssetApprovalIterator{contract: _WrappedAsset.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedAsset *WrappedAssetFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WrappedAssetApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WrappedAsset.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedAssetApproval)
				if err := _WrappedAsset.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WrappedAsset *WrappedAssetFilterer) ParseApproval(log types.Log) (*WrappedAssetApproval, error) {
	event := new(WrappedAssetApproval)
	if err := _WrappedAsset.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WrappedAssetTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WrappedAsset contract.
type WrappedAssetTransferIterator struct {
	Event *WrappedAssetTransfer // Event containing the contract specifics and raw log

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
func (it *WrappedAssetTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WrappedAssetTransfer)
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
		it.Event = new(WrappedAssetTransfer)
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
func (it *WrappedAssetTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WrappedAssetTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WrappedAssetTransfer represents a Transfer event raised by the WrappedAsset contract.
type WrappedAssetTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedAsset *WrappedAssetFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WrappedAssetTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedAsset.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WrappedAssetTransferIterator{contract: _WrappedAsset.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedAsset *WrappedAssetFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WrappedAssetTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WrappedAsset.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WrappedAssetTransfer)
				if err := _WrappedAsset.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WrappedAsset *WrappedAssetFilterer) ParseTransfer(log types.Log) (*WrappedAssetTransfer, error) {
	event := new(WrappedAssetTransfer)
	if err := _WrappedAsset.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
