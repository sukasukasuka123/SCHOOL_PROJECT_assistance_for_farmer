// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package FarmRecordManager

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

// FarmRecordManagerDiseaseTreeRemovalRecord is an auto generated low-level Go binding around an user-defined struct.
type FarmRecordManagerDiseaseTreeRemovalRecord struct {
	Timestamp               uint32
	LinkedMarkingId         uint32
	GridCodeHash            [32]byte
	ConfirmedDiseaseCount   uint8
	FalsePositiveCount      uint8
	RemovalEvidenceIPFSHash [32]byte
	Farmer                  common.Address
	ReplantScheduled        bool
}

// FarmRecordManagerMarkingOperationLog is an auto generated low-level Go binding around an user-defined struct.
type FarmRecordManagerMarkingOperationLog struct {
	Timestamp              uint32
	LinkedDetectionId      uint32
	GridCodeHash           [32]byte
	MarkingPointCount      uint8
	MarkerTypeCode         uint8
	MarkerDosagePerPoint   uint16
	IsolationPesticideCode uint8
	IsolationDosage        uint16
	OperatorDrone          common.Address
	OperationIPFSHash      [32]byte
}

// FarmRecordManagerReplantingRecord is an auto generated low-level Go binding around an user-defined struct.
type FarmRecordManagerReplantingRecord struct {
	Timestamp           uint32
	LinkedRemovalId     uint32
	GridCodeHash        [32]byte
	SeedlingCount       uint8
	SeedlingVarietyCode uint8
	QuarantineCertHash  [32]byte
	ReplantIPFSHash     [32]byte
	Farmer              common.Address
}

// FarmRecordManagerRiskDetectionRecord is an auto generated low-level Go binding around an user-defined struct.
type FarmRecordManagerRiskDetectionRecord struct {
	Timestamp        uint32
	GridCodeHash     [32]byte
	GpsCenter        [32]byte
	RiskLevel        uint8
	NdviValue        uint16
	SpectralData     [3]uint16
	PsyllidDensity   uint8
	EvidenceIPFSHash [32]byte
	OperatorDrone    common.Address
}

// FarmRecordManagerMetaData contains all meta data concerning the FarmRecordManager contract.
var FarmRecordManagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_farmFactory\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_userContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_droneFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"operationId\",\"type\":\"uint32\"}],\"name\":\"GridMarked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"recordId\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"riskLevel\",\"type\":\"uint8\"}],\"name\":\"RiskDetected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"removalId\",\"type\":\"uint32\"}],\"name\":\"TreesRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint32\",\"name\":\"replantId\",\"type\":\"uint32\"}],\"name\":\"TreesReplanted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"linkedDetectionId\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"markingPointCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"markerTypeCode\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"markerDosagePerPoint\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"isolationPesticideCode\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"isolationDosage\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"operationIPFSHash\",\"type\":\"bytes32\"}],\"name\":\"addMarkingOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"linkedRemovalId\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"seedlingCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"seedlingVarietyCode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"quarantineCertHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"replantIPFSHash\",\"type\":\"bytes32\"}],\"name\":\"addReplanting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"gpsCenter\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"riskLevel\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"ndviValue\",\"type\":\"uint16\"},{\"internalType\":\"uint16[3]\",\"name\":\"spectralData\",\"type\":\"uint16[3]\"},{\"internalType\":\"uint8\",\"name\":\"psyllidDensity\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"evidenceIPFSHash\",\"type\":\"bytes32\"}],\"name\":\"addRiskDetection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"linkedMarkingId\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"confirmedDiseaseCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"falsePositiveCount\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"removalEvidenceIPFSHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"replantScheduled\",\"type\":\"bool\"}],\"name\":\"addTreeRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"droneFactory\",\"outputs\":[{\"internalType\":\"contractDroneObjectFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"farmFactory\",\"outputs\":[{\"internalType\":\"contractIFarmFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"operationId\",\"type\":\"uint32\"}],\"name\":\"getMarkingOperation\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkedDetectionId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"gridCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"markingPointCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"markerTypeCode\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"markerDosagePerPoint\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"isolationPesticideCode\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"isolationDosage\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"operatorDrone\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"operationIPFSHash\",\"type\":\"bytes32\"}],\"internalType\":\"structFarmRecordManager.MarkingOperationLog\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"count\",\"type\":\"uint32\"}],\"name\":\"getRecentMarkingOperations\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkedDetectionId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"gridCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"markingPointCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"markerTypeCode\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"markerDosagePerPoint\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"isolationPesticideCode\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"isolationDosage\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"operatorDrone\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"operationIPFSHash\",\"type\":\"bytes32\"}],\"internalType\":\"structFarmRecordManager.MarkingOperationLog[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"count\",\"type\":\"uint32\"}],\"name\":\"getRecentRemovals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkedMarkingId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"gridCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"confirmedDiseaseCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"falsePositiveCount\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"removalEvidenceIPFSHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"farmer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"replantScheduled\",\"type\":\"bool\"}],\"internalType\":\"structFarmRecordManager.DiseaseTreeRemovalRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"count\",\"type\":\"uint32\"}],\"name\":\"getRecentReplantings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkedRemovalId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"gridCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"seedlingCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"seedlingVarietyCode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"quarantineCertHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"replantIPFSHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"farmer\",\"type\":\"address\"}],\"internalType\":\"structFarmRecordManager.ReplantingRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"count\",\"type\":\"uint32\"}],\"name\":\"getRecentRiskDetections\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"gridCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"gpsCenter\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"riskLevel\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"ndviValue\",\"type\":\"uint16\"},{\"internalType\":\"uint16[3]\",\"name\":\"spectralData\",\"type\":\"uint16[3]\"},{\"internalType\":\"uint8\",\"name\":\"psyllidDensity\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"evidenceIPFSHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operatorDrone\",\"type\":\"address\"}],\"internalType\":\"structFarmRecordManager.RiskDetectionRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"}],\"name\":\"getRecordCounts\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"risks\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"markings\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"removalsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"replantingsCount\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"removalId\",\"type\":\"uint32\"}],\"name\":\"getRemoval\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkedMarkingId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"gridCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"confirmedDiseaseCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"falsePositiveCount\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"removalEvidenceIPFSHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"farmer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"replantScheduled\",\"type\":\"bool\"}],\"internalType\":\"structFarmRecordManager.DiseaseTreeRemovalRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"replantId\",\"type\":\"uint32\"}],\"name\":\"getReplanting\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkedRemovalId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"gridCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"seedlingCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"seedlingVarietyCode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"quarantineCertHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"replantIPFSHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"farmer\",\"type\":\"address\"}],\"internalType\":\"structFarmRecordManager.ReplantingRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"recordId\",\"type\":\"uint32\"}],\"name\":\"getRiskDetection\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"gridCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"gpsCenter\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"riskLevel\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"ndviValue\",\"type\":\"uint16\"},{\"internalType\":\"uint16[3]\",\"name\":\"spectralData\",\"type\":\"uint16[3]\"},{\"internalType\":\"uint8\",\"name\":\"psyllidDensity\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"evidenceIPFSHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operatorDrone\",\"type\":\"address\"}],\"internalType\":\"structFarmRecordManager.RiskDetectionRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"markingOperationCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"markingOperations\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkedDetectionId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"gridCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"markingPointCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"markerTypeCode\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"markerDosagePerPoint\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"isolationPesticideCode\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"isolationDosage\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"operatorDrone\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"operationIPFSHash\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"removalCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"removals\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkedMarkingId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"gridCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"confirmedDiseaseCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"falsePositiveCount\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"removalEvidenceIPFSHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"farmer\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"replantScheduled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"replantingCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"replantings\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkedRemovalId\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"gridCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"seedlingCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"seedlingVarietyCode\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"quarantineCertHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"replantIPFSHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"farmer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"riskDetectionCount\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"name\":\"riskDetections\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"bytes32\",\"name\":\"gridCodeHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"gpsCenter\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"riskLevel\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"ndviValue\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"psyllidDensity\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"evidenceIPFSHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"operatorDrone\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"userContract\",\"outputs\":[{\"internalType\":\"contractUser\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// FarmRecordManagerABI is the input ABI used to generate the binding from.
// Deprecated: Use FarmRecordManagerMetaData.ABI instead.
var FarmRecordManagerABI = FarmRecordManagerMetaData.ABI

// FarmRecordManager is an auto generated Go binding around an Ethereum contract.
type FarmRecordManager struct {
	FarmRecordManagerCaller     // Read-only binding to the contract
	FarmRecordManagerTransactor // Write-only binding to the contract
	FarmRecordManagerFilterer   // Log filterer for contract events
}

// FarmRecordManagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type FarmRecordManagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmRecordManagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FarmRecordManagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmRecordManagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FarmRecordManagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FarmRecordManagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FarmRecordManagerSession struct {
	Contract     *FarmRecordManager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FarmRecordManagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FarmRecordManagerCallerSession struct {
	Contract *FarmRecordManagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// FarmRecordManagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FarmRecordManagerTransactorSession struct {
	Contract     *FarmRecordManagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// FarmRecordManagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type FarmRecordManagerRaw struct {
	Contract *FarmRecordManager // Generic contract binding to access the raw methods on
}

// FarmRecordManagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FarmRecordManagerCallerRaw struct {
	Contract *FarmRecordManagerCaller // Generic read-only contract binding to access the raw methods on
}

// FarmRecordManagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FarmRecordManagerTransactorRaw struct {
	Contract *FarmRecordManagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFarmRecordManager creates a new instance of FarmRecordManager, bound to a specific deployed contract.
func NewFarmRecordManager(address common.Address, backend bind.ContractBackend) (*FarmRecordManager, error) {
	contract, err := bindFarmRecordManager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FarmRecordManager{FarmRecordManagerCaller: FarmRecordManagerCaller{contract: contract}, FarmRecordManagerTransactor: FarmRecordManagerTransactor{contract: contract}, FarmRecordManagerFilterer: FarmRecordManagerFilterer{contract: contract}}, nil
}

// NewFarmRecordManagerCaller creates a new read-only instance of FarmRecordManager, bound to a specific deployed contract.
func NewFarmRecordManagerCaller(address common.Address, caller bind.ContractCaller) (*FarmRecordManagerCaller, error) {
	contract, err := bindFarmRecordManager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FarmRecordManagerCaller{contract: contract}, nil
}

// NewFarmRecordManagerTransactor creates a new write-only instance of FarmRecordManager, bound to a specific deployed contract.
func NewFarmRecordManagerTransactor(address common.Address, transactor bind.ContractTransactor) (*FarmRecordManagerTransactor, error) {
	contract, err := bindFarmRecordManager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FarmRecordManagerTransactor{contract: contract}, nil
}

// NewFarmRecordManagerFilterer creates a new log filterer instance of FarmRecordManager, bound to a specific deployed contract.
func NewFarmRecordManagerFilterer(address common.Address, filterer bind.ContractFilterer) (*FarmRecordManagerFilterer, error) {
	contract, err := bindFarmRecordManager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FarmRecordManagerFilterer{contract: contract}, nil
}

// bindFarmRecordManager binds a generic wrapper to an already deployed contract.
func bindFarmRecordManager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FarmRecordManagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FarmRecordManager *FarmRecordManagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FarmRecordManager.Contract.FarmRecordManagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FarmRecordManager *FarmRecordManagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarmRecordManager.Contract.FarmRecordManagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FarmRecordManager *FarmRecordManagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FarmRecordManager.Contract.FarmRecordManagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FarmRecordManager *FarmRecordManagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FarmRecordManager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FarmRecordManager *FarmRecordManagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FarmRecordManager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FarmRecordManager *FarmRecordManagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FarmRecordManager.Contract.contract.Transact(opts, method, params...)
}

// DroneFactory is a free data retrieval call binding the contract method 0xc7b26142.
//
// Solidity: function droneFactory() view returns(address)
func (_FarmRecordManager *FarmRecordManagerCaller) DroneFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "droneFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DroneFactory is a free data retrieval call binding the contract method 0xc7b26142.
//
// Solidity: function droneFactory() view returns(address)
func (_FarmRecordManager *FarmRecordManagerSession) DroneFactory() (common.Address, error) {
	return _FarmRecordManager.Contract.DroneFactory(&_FarmRecordManager.CallOpts)
}

// DroneFactory is a free data retrieval call binding the contract method 0xc7b26142.
//
// Solidity: function droneFactory() view returns(address)
func (_FarmRecordManager *FarmRecordManagerCallerSession) DroneFactory() (common.Address, error) {
	return _FarmRecordManager.Contract.DroneFactory(&_FarmRecordManager.CallOpts)
}

// FarmFactory is a free data retrieval call binding the contract method 0x6d25e175.
//
// Solidity: function farmFactory() view returns(address)
func (_FarmRecordManager *FarmRecordManagerCaller) FarmFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "farmFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FarmFactory is a free data retrieval call binding the contract method 0x6d25e175.
//
// Solidity: function farmFactory() view returns(address)
func (_FarmRecordManager *FarmRecordManagerSession) FarmFactory() (common.Address, error) {
	return _FarmRecordManager.Contract.FarmFactory(&_FarmRecordManager.CallOpts)
}

// FarmFactory is a free data retrieval call binding the contract method 0x6d25e175.
//
// Solidity: function farmFactory() view returns(address)
func (_FarmRecordManager *FarmRecordManagerCallerSession) FarmFactory() (common.Address, error) {
	return _FarmRecordManager.Contract.FarmFactory(&_FarmRecordManager.CallOpts)
}

// GetMarkingOperation is a free data retrieval call binding the contract method 0x70d39c5e.
//
// Solidity: function getMarkingOperation(address farmId, uint32 operationId) view returns((uint32,uint32,bytes32,uint8,uint8,uint16,uint8,uint16,address,bytes32))
func (_FarmRecordManager *FarmRecordManagerCaller) GetMarkingOperation(opts *bind.CallOpts, farmId common.Address, operationId uint32) (FarmRecordManagerMarkingOperationLog, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "getMarkingOperation", farmId, operationId)

	if err != nil {
		return *new(FarmRecordManagerMarkingOperationLog), err
	}

	out0 := *abi.ConvertType(out[0], new(FarmRecordManagerMarkingOperationLog)).(*FarmRecordManagerMarkingOperationLog)

	return out0, err

}

// GetMarkingOperation is a free data retrieval call binding the contract method 0x70d39c5e.
//
// Solidity: function getMarkingOperation(address farmId, uint32 operationId) view returns((uint32,uint32,bytes32,uint8,uint8,uint16,uint8,uint16,address,bytes32))
func (_FarmRecordManager *FarmRecordManagerSession) GetMarkingOperation(farmId common.Address, operationId uint32) (FarmRecordManagerMarkingOperationLog, error) {
	return _FarmRecordManager.Contract.GetMarkingOperation(&_FarmRecordManager.CallOpts, farmId, operationId)
}

// GetMarkingOperation is a free data retrieval call binding the contract method 0x70d39c5e.
//
// Solidity: function getMarkingOperation(address farmId, uint32 operationId) view returns((uint32,uint32,bytes32,uint8,uint8,uint16,uint8,uint16,address,bytes32))
func (_FarmRecordManager *FarmRecordManagerCallerSession) GetMarkingOperation(farmId common.Address, operationId uint32) (FarmRecordManagerMarkingOperationLog, error) {
	return _FarmRecordManager.Contract.GetMarkingOperation(&_FarmRecordManager.CallOpts, farmId, operationId)
}

// GetRecentMarkingOperations is a free data retrieval call binding the contract method 0xe0cc2be9.
//
// Solidity: function getRecentMarkingOperations(address farmId, uint32 count) view returns((uint32,uint32,bytes32,uint8,uint8,uint16,uint8,uint16,address,bytes32)[])
func (_FarmRecordManager *FarmRecordManagerCaller) GetRecentMarkingOperations(opts *bind.CallOpts, farmId common.Address, count uint32) ([]FarmRecordManagerMarkingOperationLog, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "getRecentMarkingOperations", farmId, count)

	if err != nil {
		return *new([]FarmRecordManagerMarkingOperationLog), err
	}

	out0 := *abi.ConvertType(out[0], new([]FarmRecordManagerMarkingOperationLog)).(*[]FarmRecordManagerMarkingOperationLog)

	return out0, err

}

// GetRecentMarkingOperations is a free data retrieval call binding the contract method 0xe0cc2be9.
//
// Solidity: function getRecentMarkingOperations(address farmId, uint32 count) view returns((uint32,uint32,bytes32,uint8,uint8,uint16,uint8,uint16,address,bytes32)[])
func (_FarmRecordManager *FarmRecordManagerSession) GetRecentMarkingOperations(farmId common.Address, count uint32) ([]FarmRecordManagerMarkingOperationLog, error) {
	return _FarmRecordManager.Contract.GetRecentMarkingOperations(&_FarmRecordManager.CallOpts, farmId, count)
}

// GetRecentMarkingOperations is a free data retrieval call binding the contract method 0xe0cc2be9.
//
// Solidity: function getRecentMarkingOperations(address farmId, uint32 count) view returns((uint32,uint32,bytes32,uint8,uint8,uint16,uint8,uint16,address,bytes32)[])
func (_FarmRecordManager *FarmRecordManagerCallerSession) GetRecentMarkingOperations(farmId common.Address, count uint32) ([]FarmRecordManagerMarkingOperationLog, error) {
	return _FarmRecordManager.Contract.GetRecentMarkingOperations(&_FarmRecordManager.CallOpts, farmId, count)
}

// GetRecentRemovals is a free data retrieval call binding the contract method 0xe26bb78d.
//
// Solidity: function getRecentRemovals(address farmId, uint32 count) view returns((uint32,uint32,bytes32,uint8,uint8,bytes32,address,bool)[])
func (_FarmRecordManager *FarmRecordManagerCaller) GetRecentRemovals(opts *bind.CallOpts, farmId common.Address, count uint32) ([]FarmRecordManagerDiseaseTreeRemovalRecord, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "getRecentRemovals", farmId, count)

	if err != nil {
		return *new([]FarmRecordManagerDiseaseTreeRemovalRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]FarmRecordManagerDiseaseTreeRemovalRecord)).(*[]FarmRecordManagerDiseaseTreeRemovalRecord)

	return out0, err

}

// GetRecentRemovals is a free data retrieval call binding the contract method 0xe26bb78d.
//
// Solidity: function getRecentRemovals(address farmId, uint32 count) view returns((uint32,uint32,bytes32,uint8,uint8,bytes32,address,bool)[])
func (_FarmRecordManager *FarmRecordManagerSession) GetRecentRemovals(farmId common.Address, count uint32) ([]FarmRecordManagerDiseaseTreeRemovalRecord, error) {
	return _FarmRecordManager.Contract.GetRecentRemovals(&_FarmRecordManager.CallOpts, farmId, count)
}

// GetRecentRemovals is a free data retrieval call binding the contract method 0xe26bb78d.
//
// Solidity: function getRecentRemovals(address farmId, uint32 count) view returns((uint32,uint32,bytes32,uint8,uint8,bytes32,address,bool)[])
func (_FarmRecordManager *FarmRecordManagerCallerSession) GetRecentRemovals(farmId common.Address, count uint32) ([]FarmRecordManagerDiseaseTreeRemovalRecord, error) {
	return _FarmRecordManager.Contract.GetRecentRemovals(&_FarmRecordManager.CallOpts, farmId, count)
}

// GetRecentReplantings is a free data retrieval call binding the contract method 0x06bc31ef.
//
// Solidity: function getRecentReplantings(address farmId, uint32 count) view returns((uint32,uint32,bytes32,uint8,uint8,bytes32,bytes32,address)[])
func (_FarmRecordManager *FarmRecordManagerCaller) GetRecentReplantings(opts *bind.CallOpts, farmId common.Address, count uint32) ([]FarmRecordManagerReplantingRecord, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "getRecentReplantings", farmId, count)

	if err != nil {
		return *new([]FarmRecordManagerReplantingRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]FarmRecordManagerReplantingRecord)).(*[]FarmRecordManagerReplantingRecord)

	return out0, err

}

// GetRecentReplantings is a free data retrieval call binding the contract method 0x06bc31ef.
//
// Solidity: function getRecentReplantings(address farmId, uint32 count) view returns((uint32,uint32,bytes32,uint8,uint8,bytes32,bytes32,address)[])
func (_FarmRecordManager *FarmRecordManagerSession) GetRecentReplantings(farmId common.Address, count uint32) ([]FarmRecordManagerReplantingRecord, error) {
	return _FarmRecordManager.Contract.GetRecentReplantings(&_FarmRecordManager.CallOpts, farmId, count)
}

// GetRecentReplantings is a free data retrieval call binding the contract method 0x06bc31ef.
//
// Solidity: function getRecentReplantings(address farmId, uint32 count) view returns((uint32,uint32,bytes32,uint8,uint8,bytes32,bytes32,address)[])
func (_FarmRecordManager *FarmRecordManagerCallerSession) GetRecentReplantings(farmId common.Address, count uint32) ([]FarmRecordManagerReplantingRecord, error) {
	return _FarmRecordManager.Contract.GetRecentReplantings(&_FarmRecordManager.CallOpts, farmId, count)
}

// GetRecentRiskDetections is a free data retrieval call binding the contract method 0x337b7aed.
//
// Solidity: function getRecentRiskDetections(address farmId, uint32 count) view returns((uint32,bytes32,bytes32,uint8,uint16,uint16[3],uint8,bytes32,address)[])
func (_FarmRecordManager *FarmRecordManagerCaller) GetRecentRiskDetections(opts *bind.CallOpts, farmId common.Address, count uint32) ([]FarmRecordManagerRiskDetectionRecord, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "getRecentRiskDetections", farmId, count)

	if err != nil {
		return *new([]FarmRecordManagerRiskDetectionRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]FarmRecordManagerRiskDetectionRecord)).(*[]FarmRecordManagerRiskDetectionRecord)

	return out0, err

}

// GetRecentRiskDetections is a free data retrieval call binding the contract method 0x337b7aed.
//
// Solidity: function getRecentRiskDetections(address farmId, uint32 count) view returns((uint32,bytes32,bytes32,uint8,uint16,uint16[3],uint8,bytes32,address)[])
func (_FarmRecordManager *FarmRecordManagerSession) GetRecentRiskDetections(farmId common.Address, count uint32) ([]FarmRecordManagerRiskDetectionRecord, error) {
	return _FarmRecordManager.Contract.GetRecentRiskDetections(&_FarmRecordManager.CallOpts, farmId, count)
}

// GetRecentRiskDetections is a free data retrieval call binding the contract method 0x337b7aed.
//
// Solidity: function getRecentRiskDetections(address farmId, uint32 count) view returns((uint32,bytes32,bytes32,uint8,uint16,uint16[3],uint8,bytes32,address)[])
func (_FarmRecordManager *FarmRecordManagerCallerSession) GetRecentRiskDetections(farmId common.Address, count uint32) ([]FarmRecordManagerRiskDetectionRecord, error) {
	return _FarmRecordManager.Contract.GetRecentRiskDetections(&_FarmRecordManager.CallOpts, farmId, count)
}

// GetRecordCounts is a free data retrieval call binding the contract method 0x4ef09211.
//
// Solidity: function getRecordCounts(address farmId) view returns(uint32 risks, uint32 markings, uint32 removalsCount, uint32 replantingsCount)
func (_FarmRecordManager *FarmRecordManagerCaller) GetRecordCounts(opts *bind.CallOpts, farmId common.Address) (struct {
	Risks            uint32
	Markings         uint32
	RemovalsCount    uint32
	ReplantingsCount uint32
}, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "getRecordCounts", farmId)

	outstruct := new(struct {
		Risks            uint32
		Markings         uint32
		RemovalsCount    uint32
		ReplantingsCount uint32
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Risks = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Markings = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.RemovalsCount = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.ReplantingsCount = *abi.ConvertType(out[3], new(uint32)).(*uint32)

	return *outstruct, err

}

// GetRecordCounts is a free data retrieval call binding the contract method 0x4ef09211.
//
// Solidity: function getRecordCounts(address farmId) view returns(uint32 risks, uint32 markings, uint32 removalsCount, uint32 replantingsCount)
func (_FarmRecordManager *FarmRecordManagerSession) GetRecordCounts(farmId common.Address) (struct {
	Risks            uint32
	Markings         uint32
	RemovalsCount    uint32
	ReplantingsCount uint32
}, error) {
	return _FarmRecordManager.Contract.GetRecordCounts(&_FarmRecordManager.CallOpts, farmId)
}

// GetRecordCounts is a free data retrieval call binding the contract method 0x4ef09211.
//
// Solidity: function getRecordCounts(address farmId) view returns(uint32 risks, uint32 markings, uint32 removalsCount, uint32 replantingsCount)
func (_FarmRecordManager *FarmRecordManagerCallerSession) GetRecordCounts(farmId common.Address) (struct {
	Risks            uint32
	Markings         uint32
	RemovalsCount    uint32
	ReplantingsCount uint32
}, error) {
	return _FarmRecordManager.Contract.GetRecordCounts(&_FarmRecordManager.CallOpts, farmId)
}

// GetRemoval is a free data retrieval call binding the contract method 0x8cd5c437.
//
// Solidity: function getRemoval(address farmId, uint32 removalId) view returns((uint32,uint32,bytes32,uint8,uint8,bytes32,address,bool))
func (_FarmRecordManager *FarmRecordManagerCaller) GetRemoval(opts *bind.CallOpts, farmId common.Address, removalId uint32) (FarmRecordManagerDiseaseTreeRemovalRecord, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "getRemoval", farmId, removalId)

	if err != nil {
		return *new(FarmRecordManagerDiseaseTreeRemovalRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(FarmRecordManagerDiseaseTreeRemovalRecord)).(*FarmRecordManagerDiseaseTreeRemovalRecord)

	return out0, err

}

// GetRemoval is a free data retrieval call binding the contract method 0x8cd5c437.
//
// Solidity: function getRemoval(address farmId, uint32 removalId) view returns((uint32,uint32,bytes32,uint8,uint8,bytes32,address,bool))
func (_FarmRecordManager *FarmRecordManagerSession) GetRemoval(farmId common.Address, removalId uint32) (FarmRecordManagerDiseaseTreeRemovalRecord, error) {
	return _FarmRecordManager.Contract.GetRemoval(&_FarmRecordManager.CallOpts, farmId, removalId)
}

// GetRemoval is a free data retrieval call binding the contract method 0x8cd5c437.
//
// Solidity: function getRemoval(address farmId, uint32 removalId) view returns((uint32,uint32,bytes32,uint8,uint8,bytes32,address,bool))
func (_FarmRecordManager *FarmRecordManagerCallerSession) GetRemoval(farmId common.Address, removalId uint32) (FarmRecordManagerDiseaseTreeRemovalRecord, error) {
	return _FarmRecordManager.Contract.GetRemoval(&_FarmRecordManager.CallOpts, farmId, removalId)
}

// GetReplanting is a free data retrieval call binding the contract method 0xddffe65c.
//
// Solidity: function getReplanting(address farmId, uint32 replantId) view returns((uint32,uint32,bytes32,uint8,uint8,bytes32,bytes32,address))
func (_FarmRecordManager *FarmRecordManagerCaller) GetReplanting(opts *bind.CallOpts, farmId common.Address, replantId uint32) (FarmRecordManagerReplantingRecord, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "getReplanting", farmId, replantId)

	if err != nil {
		return *new(FarmRecordManagerReplantingRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(FarmRecordManagerReplantingRecord)).(*FarmRecordManagerReplantingRecord)

	return out0, err

}

// GetReplanting is a free data retrieval call binding the contract method 0xddffe65c.
//
// Solidity: function getReplanting(address farmId, uint32 replantId) view returns((uint32,uint32,bytes32,uint8,uint8,bytes32,bytes32,address))
func (_FarmRecordManager *FarmRecordManagerSession) GetReplanting(farmId common.Address, replantId uint32) (FarmRecordManagerReplantingRecord, error) {
	return _FarmRecordManager.Contract.GetReplanting(&_FarmRecordManager.CallOpts, farmId, replantId)
}

// GetReplanting is a free data retrieval call binding the contract method 0xddffe65c.
//
// Solidity: function getReplanting(address farmId, uint32 replantId) view returns((uint32,uint32,bytes32,uint8,uint8,bytes32,bytes32,address))
func (_FarmRecordManager *FarmRecordManagerCallerSession) GetReplanting(farmId common.Address, replantId uint32) (FarmRecordManagerReplantingRecord, error) {
	return _FarmRecordManager.Contract.GetReplanting(&_FarmRecordManager.CallOpts, farmId, replantId)
}

// GetRiskDetection is a free data retrieval call binding the contract method 0x74f0d363.
//
// Solidity: function getRiskDetection(address farmId, uint32 recordId) view returns((uint32,bytes32,bytes32,uint8,uint16,uint16[3],uint8,bytes32,address))
func (_FarmRecordManager *FarmRecordManagerCaller) GetRiskDetection(opts *bind.CallOpts, farmId common.Address, recordId uint32) (FarmRecordManagerRiskDetectionRecord, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "getRiskDetection", farmId, recordId)

	if err != nil {
		return *new(FarmRecordManagerRiskDetectionRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(FarmRecordManagerRiskDetectionRecord)).(*FarmRecordManagerRiskDetectionRecord)

	return out0, err

}

// GetRiskDetection is a free data retrieval call binding the contract method 0x74f0d363.
//
// Solidity: function getRiskDetection(address farmId, uint32 recordId) view returns((uint32,bytes32,bytes32,uint8,uint16,uint16[3],uint8,bytes32,address))
func (_FarmRecordManager *FarmRecordManagerSession) GetRiskDetection(farmId common.Address, recordId uint32) (FarmRecordManagerRiskDetectionRecord, error) {
	return _FarmRecordManager.Contract.GetRiskDetection(&_FarmRecordManager.CallOpts, farmId, recordId)
}

// GetRiskDetection is a free data retrieval call binding the contract method 0x74f0d363.
//
// Solidity: function getRiskDetection(address farmId, uint32 recordId) view returns((uint32,bytes32,bytes32,uint8,uint16,uint16[3],uint8,bytes32,address))
func (_FarmRecordManager *FarmRecordManagerCallerSession) GetRiskDetection(farmId common.Address, recordId uint32) (FarmRecordManagerRiskDetectionRecord, error) {
	return _FarmRecordManager.Contract.GetRiskDetection(&_FarmRecordManager.CallOpts, farmId, recordId)
}

// MarkingOperationCount is a free data retrieval call binding the contract method 0xb96f5c24.
//
// Solidity: function markingOperationCount(address ) view returns(uint32)
func (_FarmRecordManager *FarmRecordManagerCaller) MarkingOperationCount(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "markingOperationCount", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MarkingOperationCount is a free data retrieval call binding the contract method 0xb96f5c24.
//
// Solidity: function markingOperationCount(address ) view returns(uint32)
func (_FarmRecordManager *FarmRecordManagerSession) MarkingOperationCount(arg0 common.Address) (uint32, error) {
	return _FarmRecordManager.Contract.MarkingOperationCount(&_FarmRecordManager.CallOpts, arg0)
}

// MarkingOperationCount is a free data retrieval call binding the contract method 0xb96f5c24.
//
// Solidity: function markingOperationCount(address ) view returns(uint32)
func (_FarmRecordManager *FarmRecordManagerCallerSession) MarkingOperationCount(arg0 common.Address) (uint32, error) {
	return _FarmRecordManager.Contract.MarkingOperationCount(&_FarmRecordManager.CallOpts, arg0)
}

// MarkingOperations is a free data retrieval call binding the contract method 0x979f964f.
//
// Solidity: function markingOperations(address , uint32 ) view returns(uint32 timestamp, uint32 linkedDetectionId, bytes32 gridCodeHash, uint8 markingPointCount, uint8 markerTypeCode, uint16 markerDosagePerPoint, uint8 isolationPesticideCode, uint16 isolationDosage, address operatorDrone, bytes32 operationIPFSHash)
func (_FarmRecordManager *FarmRecordManagerCaller) MarkingOperations(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	Timestamp              uint32
	LinkedDetectionId      uint32
	GridCodeHash           [32]byte
	MarkingPointCount      uint8
	MarkerTypeCode         uint8
	MarkerDosagePerPoint   uint16
	IsolationPesticideCode uint8
	IsolationDosage        uint16
	OperatorDrone          common.Address
	OperationIPFSHash      [32]byte
}, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "markingOperations", arg0, arg1)

	outstruct := new(struct {
		Timestamp              uint32
		LinkedDetectionId      uint32
		GridCodeHash           [32]byte
		MarkingPointCount      uint8
		MarkerTypeCode         uint8
		MarkerDosagePerPoint   uint16
		IsolationPesticideCode uint8
		IsolationDosage        uint16
		OperatorDrone          common.Address
		OperationIPFSHash      [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.LinkedDetectionId = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.GridCodeHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.MarkingPointCount = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.MarkerTypeCode = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.MarkerDosagePerPoint = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.IsolationPesticideCode = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.IsolationDosage = *abi.ConvertType(out[7], new(uint16)).(*uint16)
	outstruct.OperatorDrone = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)
	outstruct.OperationIPFSHash = *abi.ConvertType(out[9], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// MarkingOperations is a free data retrieval call binding the contract method 0x979f964f.
//
// Solidity: function markingOperations(address , uint32 ) view returns(uint32 timestamp, uint32 linkedDetectionId, bytes32 gridCodeHash, uint8 markingPointCount, uint8 markerTypeCode, uint16 markerDosagePerPoint, uint8 isolationPesticideCode, uint16 isolationDosage, address operatorDrone, bytes32 operationIPFSHash)
func (_FarmRecordManager *FarmRecordManagerSession) MarkingOperations(arg0 common.Address, arg1 uint32) (struct {
	Timestamp              uint32
	LinkedDetectionId      uint32
	GridCodeHash           [32]byte
	MarkingPointCount      uint8
	MarkerTypeCode         uint8
	MarkerDosagePerPoint   uint16
	IsolationPesticideCode uint8
	IsolationDosage        uint16
	OperatorDrone          common.Address
	OperationIPFSHash      [32]byte
}, error) {
	return _FarmRecordManager.Contract.MarkingOperations(&_FarmRecordManager.CallOpts, arg0, arg1)
}

// MarkingOperations is a free data retrieval call binding the contract method 0x979f964f.
//
// Solidity: function markingOperations(address , uint32 ) view returns(uint32 timestamp, uint32 linkedDetectionId, bytes32 gridCodeHash, uint8 markingPointCount, uint8 markerTypeCode, uint16 markerDosagePerPoint, uint8 isolationPesticideCode, uint16 isolationDosage, address operatorDrone, bytes32 operationIPFSHash)
func (_FarmRecordManager *FarmRecordManagerCallerSession) MarkingOperations(arg0 common.Address, arg1 uint32) (struct {
	Timestamp              uint32
	LinkedDetectionId      uint32
	GridCodeHash           [32]byte
	MarkingPointCount      uint8
	MarkerTypeCode         uint8
	MarkerDosagePerPoint   uint16
	IsolationPesticideCode uint8
	IsolationDosage        uint16
	OperatorDrone          common.Address
	OperationIPFSHash      [32]byte
}, error) {
	return _FarmRecordManager.Contract.MarkingOperations(&_FarmRecordManager.CallOpts, arg0, arg1)
}

// RemovalCount is a free data retrieval call binding the contract method 0x8e964d82.
//
// Solidity: function removalCount(address ) view returns(uint32)
func (_FarmRecordManager *FarmRecordManagerCaller) RemovalCount(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "removalCount", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RemovalCount is a free data retrieval call binding the contract method 0x8e964d82.
//
// Solidity: function removalCount(address ) view returns(uint32)
func (_FarmRecordManager *FarmRecordManagerSession) RemovalCount(arg0 common.Address) (uint32, error) {
	return _FarmRecordManager.Contract.RemovalCount(&_FarmRecordManager.CallOpts, arg0)
}

// RemovalCount is a free data retrieval call binding the contract method 0x8e964d82.
//
// Solidity: function removalCount(address ) view returns(uint32)
func (_FarmRecordManager *FarmRecordManagerCallerSession) RemovalCount(arg0 common.Address) (uint32, error) {
	return _FarmRecordManager.Contract.RemovalCount(&_FarmRecordManager.CallOpts, arg0)
}

// Removals is a free data retrieval call binding the contract method 0x1086430b.
//
// Solidity: function removals(address , uint32 ) view returns(uint32 timestamp, uint32 linkedMarkingId, bytes32 gridCodeHash, uint8 confirmedDiseaseCount, uint8 falsePositiveCount, bytes32 removalEvidenceIPFSHash, address farmer, bool replantScheduled)
func (_FarmRecordManager *FarmRecordManagerCaller) Removals(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	Timestamp               uint32
	LinkedMarkingId         uint32
	GridCodeHash            [32]byte
	ConfirmedDiseaseCount   uint8
	FalsePositiveCount      uint8
	RemovalEvidenceIPFSHash [32]byte
	Farmer                  common.Address
	ReplantScheduled        bool
}, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "removals", arg0, arg1)

	outstruct := new(struct {
		Timestamp               uint32
		LinkedMarkingId         uint32
		GridCodeHash            [32]byte
		ConfirmedDiseaseCount   uint8
		FalsePositiveCount      uint8
		RemovalEvidenceIPFSHash [32]byte
		Farmer                  common.Address
		ReplantScheduled        bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.LinkedMarkingId = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.GridCodeHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.ConfirmedDiseaseCount = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.FalsePositiveCount = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.RemovalEvidenceIPFSHash = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Farmer = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.ReplantScheduled = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// Removals is a free data retrieval call binding the contract method 0x1086430b.
//
// Solidity: function removals(address , uint32 ) view returns(uint32 timestamp, uint32 linkedMarkingId, bytes32 gridCodeHash, uint8 confirmedDiseaseCount, uint8 falsePositiveCount, bytes32 removalEvidenceIPFSHash, address farmer, bool replantScheduled)
func (_FarmRecordManager *FarmRecordManagerSession) Removals(arg0 common.Address, arg1 uint32) (struct {
	Timestamp               uint32
	LinkedMarkingId         uint32
	GridCodeHash            [32]byte
	ConfirmedDiseaseCount   uint8
	FalsePositiveCount      uint8
	RemovalEvidenceIPFSHash [32]byte
	Farmer                  common.Address
	ReplantScheduled        bool
}, error) {
	return _FarmRecordManager.Contract.Removals(&_FarmRecordManager.CallOpts, arg0, arg1)
}

// Removals is a free data retrieval call binding the contract method 0x1086430b.
//
// Solidity: function removals(address , uint32 ) view returns(uint32 timestamp, uint32 linkedMarkingId, bytes32 gridCodeHash, uint8 confirmedDiseaseCount, uint8 falsePositiveCount, bytes32 removalEvidenceIPFSHash, address farmer, bool replantScheduled)
func (_FarmRecordManager *FarmRecordManagerCallerSession) Removals(arg0 common.Address, arg1 uint32) (struct {
	Timestamp               uint32
	LinkedMarkingId         uint32
	GridCodeHash            [32]byte
	ConfirmedDiseaseCount   uint8
	FalsePositiveCount      uint8
	RemovalEvidenceIPFSHash [32]byte
	Farmer                  common.Address
	ReplantScheduled        bool
}, error) {
	return _FarmRecordManager.Contract.Removals(&_FarmRecordManager.CallOpts, arg0, arg1)
}

// ReplantingCount is a free data retrieval call binding the contract method 0xf5243cf0.
//
// Solidity: function replantingCount(address ) view returns(uint32)
func (_FarmRecordManager *FarmRecordManagerCaller) ReplantingCount(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "replantingCount", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ReplantingCount is a free data retrieval call binding the contract method 0xf5243cf0.
//
// Solidity: function replantingCount(address ) view returns(uint32)
func (_FarmRecordManager *FarmRecordManagerSession) ReplantingCount(arg0 common.Address) (uint32, error) {
	return _FarmRecordManager.Contract.ReplantingCount(&_FarmRecordManager.CallOpts, arg0)
}

// ReplantingCount is a free data retrieval call binding the contract method 0xf5243cf0.
//
// Solidity: function replantingCount(address ) view returns(uint32)
func (_FarmRecordManager *FarmRecordManagerCallerSession) ReplantingCount(arg0 common.Address) (uint32, error) {
	return _FarmRecordManager.Contract.ReplantingCount(&_FarmRecordManager.CallOpts, arg0)
}

// Replantings is a free data retrieval call binding the contract method 0x1c8dde84.
//
// Solidity: function replantings(address , uint32 ) view returns(uint32 timestamp, uint32 linkedRemovalId, bytes32 gridCodeHash, uint8 seedlingCount, uint8 seedlingVarietyCode, bytes32 quarantineCertHash, bytes32 replantIPFSHash, address farmer)
func (_FarmRecordManager *FarmRecordManagerCaller) Replantings(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	Timestamp           uint32
	LinkedRemovalId     uint32
	GridCodeHash        [32]byte
	SeedlingCount       uint8
	SeedlingVarietyCode uint8
	QuarantineCertHash  [32]byte
	ReplantIPFSHash     [32]byte
	Farmer              common.Address
}, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "replantings", arg0, arg1)

	outstruct := new(struct {
		Timestamp           uint32
		LinkedRemovalId     uint32
		GridCodeHash        [32]byte
		SeedlingCount       uint8
		SeedlingVarietyCode uint8
		QuarantineCertHash  [32]byte
		ReplantIPFSHash     [32]byte
		Farmer              common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.LinkedRemovalId = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.GridCodeHash = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.SeedlingCount = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.SeedlingVarietyCode = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.QuarantineCertHash = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.ReplantIPFSHash = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)
	outstruct.Farmer = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Replantings is a free data retrieval call binding the contract method 0x1c8dde84.
//
// Solidity: function replantings(address , uint32 ) view returns(uint32 timestamp, uint32 linkedRemovalId, bytes32 gridCodeHash, uint8 seedlingCount, uint8 seedlingVarietyCode, bytes32 quarantineCertHash, bytes32 replantIPFSHash, address farmer)
func (_FarmRecordManager *FarmRecordManagerSession) Replantings(arg0 common.Address, arg1 uint32) (struct {
	Timestamp           uint32
	LinkedRemovalId     uint32
	GridCodeHash        [32]byte
	SeedlingCount       uint8
	SeedlingVarietyCode uint8
	QuarantineCertHash  [32]byte
	ReplantIPFSHash     [32]byte
	Farmer              common.Address
}, error) {
	return _FarmRecordManager.Contract.Replantings(&_FarmRecordManager.CallOpts, arg0, arg1)
}

// Replantings is a free data retrieval call binding the contract method 0x1c8dde84.
//
// Solidity: function replantings(address , uint32 ) view returns(uint32 timestamp, uint32 linkedRemovalId, bytes32 gridCodeHash, uint8 seedlingCount, uint8 seedlingVarietyCode, bytes32 quarantineCertHash, bytes32 replantIPFSHash, address farmer)
func (_FarmRecordManager *FarmRecordManagerCallerSession) Replantings(arg0 common.Address, arg1 uint32) (struct {
	Timestamp           uint32
	LinkedRemovalId     uint32
	GridCodeHash        [32]byte
	SeedlingCount       uint8
	SeedlingVarietyCode uint8
	QuarantineCertHash  [32]byte
	ReplantIPFSHash     [32]byte
	Farmer              common.Address
}, error) {
	return _FarmRecordManager.Contract.Replantings(&_FarmRecordManager.CallOpts, arg0, arg1)
}

// RiskDetectionCount is a free data retrieval call binding the contract method 0x1637ba8e.
//
// Solidity: function riskDetectionCount(address ) view returns(uint32)
func (_FarmRecordManager *FarmRecordManagerCaller) RiskDetectionCount(opts *bind.CallOpts, arg0 common.Address) (uint32, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "riskDetectionCount", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// RiskDetectionCount is a free data retrieval call binding the contract method 0x1637ba8e.
//
// Solidity: function riskDetectionCount(address ) view returns(uint32)
func (_FarmRecordManager *FarmRecordManagerSession) RiskDetectionCount(arg0 common.Address) (uint32, error) {
	return _FarmRecordManager.Contract.RiskDetectionCount(&_FarmRecordManager.CallOpts, arg0)
}

// RiskDetectionCount is a free data retrieval call binding the contract method 0x1637ba8e.
//
// Solidity: function riskDetectionCount(address ) view returns(uint32)
func (_FarmRecordManager *FarmRecordManagerCallerSession) RiskDetectionCount(arg0 common.Address) (uint32, error) {
	return _FarmRecordManager.Contract.RiskDetectionCount(&_FarmRecordManager.CallOpts, arg0)
}

// RiskDetections is a free data retrieval call binding the contract method 0x85518709.
//
// Solidity: function riskDetections(address , uint32 ) view returns(uint32 timestamp, bytes32 gridCodeHash, bytes32 gpsCenter, uint8 riskLevel, uint16 ndviValue, uint8 psyllidDensity, bytes32 evidenceIPFSHash, address operatorDrone)
func (_FarmRecordManager *FarmRecordManagerCaller) RiskDetections(opts *bind.CallOpts, arg0 common.Address, arg1 uint32) (struct {
	Timestamp        uint32
	GridCodeHash     [32]byte
	GpsCenter        [32]byte
	RiskLevel        uint8
	NdviValue        uint16
	PsyllidDensity   uint8
	EvidenceIPFSHash [32]byte
	OperatorDrone    common.Address
}, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "riskDetections", arg0, arg1)

	outstruct := new(struct {
		Timestamp        uint32
		GridCodeHash     [32]byte
		GpsCenter        [32]byte
		RiskLevel        uint8
		NdviValue        uint16
		PsyllidDensity   uint8
		EvidenceIPFSHash [32]byte
		OperatorDrone    common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Timestamp = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.GridCodeHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.GpsCenter = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.RiskLevel = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.NdviValue = *abi.ConvertType(out[4], new(uint16)).(*uint16)
	outstruct.PsyllidDensity = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.EvidenceIPFSHash = *abi.ConvertType(out[6], new([32]byte)).(*[32]byte)
	outstruct.OperatorDrone = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// RiskDetections is a free data retrieval call binding the contract method 0x85518709.
//
// Solidity: function riskDetections(address , uint32 ) view returns(uint32 timestamp, bytes32 gridCodeHash, bytes32 gpsCenter, uint8 riskLevel, uint16 ndviValue, uint8 psyllidDensity, bytes32 evidenceIPFSHash, address operatorDrone)
func (_FarmRecordManager *FarmRecordManagerSession) RiskDetections(arg0 common.Address, arg1 uint32) (struct {
	Timestamp        uint32
	GridCodeHash     [32]byte
	GpsCenter        [32]byte
	RiskLevel        uint8
	NdviValue        uint16
	PsyllidDensity   uint8
	EvidenceIPFSHash [32]byte
	OperatorDrone    common.Address
}, error) {
	return _FarmRecordManager.Contract.RiskDetections(&_FarmRecordManager.CallOpts, arg0, arg1)
}

// RiskDetections is a free data retrieval call binding the contract method 0x85518709.
//
// Solidity: function riskDetections(address , uint32 ) view returns(uint32 timestamp, bytes32 gridCodeHash, bytes32 gpsCenter, uint8 riskLevel, uint16 ndviValue, uint8 psyllidDensity, bytes32 evidenceIPFSHash, address operatorDrone)
func (_FarmRecordManager *FarmRecordManagerCallerSession) RiskDetections(arg0 common.Address, arg1 uint32) (struct {
	Timestamp        uint32
	GridCodeHash     [32]byte
	GpsCenter        [32]byte
	RiskLevel        uint8
	NdviValue        uint16
	PsyllidDensity   uint8
	EvidenceIPFSHash [32]byte
	OperatorDrone    common.Address
}, error) {
	return _FarmRecordManager.Contract.RiskDetections(&_FarmRecordManager.CallOpts, arg0, arg1)
}

// UserContract is a free data retrieval call binding the contract method 0x5000596b.
//
// Solidity: function userContract() view returns(address)
func (_FarmRecordManager *FarmRecordManagerCaller) UserContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FarmRecordManager.contract.Call(opts, &out, "userContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UserContract is a free data retrieval call binding the contract method 0x5000596b.
//
// Solidity: function userContract() view returns(address)
func (_FarmRecordManager *FarmRecordManagerSession) UserContract() (common.Address, error) {
	return _FarmRecordManager.Contract.UserContract(&_FarmRecordManager.CallOpts)
}

// UserContract is a free data retrieval call binding the contract method 0x5000596b.
//
// Solidity: function userContract() view returns(address)
func (_FarmRecordManager *FarmRecordManagerCallerSession) UserContract() (common.Address, error) {
	return _FarmRecordManager.Contract.UserContract(&_FarmRecordManager.CallOpts)
}

// AddMarkingOperation is a paid mutator transaction binding the contract method 0xbebb241e.
//
// Solidity: function addMarkingOperation(address farmId, string gridCode, uint32 linkedDetectionId, uint8 markingPointCount, uint8 markerTypeCode, uint16 markerDosagePerPoint, uint8 isolationPesticideCode, uint16 isolationDosage, bytes32 operationIPFSHash) returns()
func (_FarmRecordManager *FarmRecordManagerTransactor) AddMarkingOperation(opts *bind.TransactOpts, farmId common.Address, gridCode string, linkedDetectionId uint32, markingPointCount uint8, markerTypeCode uint8, markerDosagePerPoint uint16, isolationPesticideCode uint8, isolationDosage uint16, operationIPFSHash [32]byte) (*types.Transaction, error) {
	return _FarmRecordManager.contract.Transact(opts, "addMarkingOperation", farmId, gridCode, linkedDetectionId, markingPointCount, markerTypeCode, markerDosagePerPoint, isolationPesticideCode, isolationDosage, operationIPFSHash)
}

// AddMarkingOperation is a paid mutator transaction binding the contract method 0xbebb241e.
//
// Solidity: function addMarkingOperation(address farmId, string gridCode, uint32 linkedDetectionId, uint8 markingPointCount, uint8 markerTypeCode, uint16 markerDosagePerPoint, uint8 isolationPesticideCode, uint16 isolationDosage, bytes32 operationIPFSHash) returns()
func (_FarmRecordManager *FarmRecordManagerSession) AddMarkingOperation(farmId common.Address, gridCode string, linkedDetectionId uint32, markingPointCount uint8, markerTypeCode uint8, markerDosagePerPoint uint16, isolationPesticideCode uint8, isolationDosage uint16, operationIPFSHash [32]byte) (*types.Transaction, error) {
	return _FarmRecordManager.Contract.AddMarkingOperation(&_FarmRecordManager.TransactOpts, farmId, gridCode, linkedDetectionId, markingPointCount, markerTypeCode, markerDosagePerPoint, isolationPesticideCode, isolationDosage, operationIPFSHash)
}

// AddMarkingOperation is a paid mutator transaction binding the contract method 0xbebb241e.
//
// Solidity: function addMarkingOperation(address farmId, string gridCode, uint32 linkedDetectionId, uint8 markingPointCount, uint8 markerTypeCode, uint16 markerDosagePerPoint, uint8 isolationPesticideCode, uint16 isolationDosage, bytes32 operationIPFSHash) returns()
func (_FarmRecordManager *FarmRecordManagerTransactorSession) AddMarkingOperation(farmId common.Address, gridCode string, linkedDetectionId uint32, markingPointCount uint8, markerTypeCode uint8, markerDosagePerPoint uint16, isolationPesticideCode uint8, isolationDosage uint16, operationIPFSHash [32]byte) (*types.Transaction, error) {
	return _FarmRecordManager.Contract.AddMarkingOperation(&_FarmRecordManager.TransactOpts, farmId, gridCode, linkedDetectionId, markingPointCount, markerTypeCode, markerDosagePerPoint, isolationPesticideCode, isolationDosage, operationIPFSHash)
}

// AddReplanting is a paid mutator transaction binding the contract method 0x7ef46672.
//
// Solidity: function addReplanting(address farmId, string gridCode, uint32 linkedRemovalId, uint8 seedlingCount, uint8 seedlingVarietyCode, bytes32 quarantineCertHash, bytes32 replantIPFSHash) returns()
func (_FarmRecordManager *FarmRecordManagerTransactor) AddReplanting(opts *bind.TransactOpts, farmId common.Address, gridCode string, linkedRemovalId uint32, seedlingCount uint8, seedlingVarietyCode uint8, quarantineCertHash [32]byte, replantIPFSHash [32]byte) (*types.Transaction, error) {
	return _FarmRecordManager.contract.Transact(opts, "addReplanting", farmId, gridCode, linkedRemovalId, seedlingCount, seedlingVarietyCode, quarantineCertHash, replantIPFSHash)
}

// AddReplanting is a paid mutator transaction binding the contract method 0x7ef46672.
//
// Solidity: function addReplanting(address farmId, string gridCode, uint32 linkedRemovalId, uint8 seedlingCount, uint8 seedlingVarietyCode, bytes32 quarantineCertHash, bytes32 replantIPFSHash) returns()
func (_FarmRecordManager *FarmRecordManagerSession) AddReplanting(farmId common.Address, gridCode string, linkedRemovalId uint32, seedlingCount uint8, seedlingVarietyCode uint8, quarantineCertHash [32]byte, replantIPFSHash [32]byte) (*types.Transaction, error) {
	return _FarmRecordManager.Contract.AddReplanting(&_FarmRecordManager.TransactOpts, farmId, gridCode, linkedRemovalId, seedlingCount, seedlingVarietyCode, quarantineCertHash, replantIPFSHash)
}

// AddReplanting is a paid mutator transaction binding the contract method 0x7ef46672.
//
// Solidity: function addReplanting(address farmId, string gridCode, uint32 linkedRemovalId, uint8 seedlingCount, uint8 seedlingVarietyCode, bytes32 quarantineCertHash, bytes32 replantIPFSHash) returns()
func (_FarmRecordManager *FarmRecordManagerTransactorSession) AddReplanting(farmId common.Address, gridCode string, linkedRemovalId uint32, seedlingCount uint8, seedlingVarietyCode uint8, quarantineCertHash [32]byte, replantIPFSHash [32]byte) (*types.Transaction, error) {
	return _FarmRecordManager.Contract.AddReplanting(&_FarmRecordManager.TransactOpts, farmId, gridCode, linkedRemovalId, seedlingCount, seedlingVarietyCode, quarantineCertHash, replantIPFSHash)
}

// AddRiskDetection is a paid mutator transaction binding the contract method 0x4807971a.
//
// Solidity: function addRiskDetection(address farmId, string gridCode, bytes32 gpsCenter, uint8 riskLevel, uint16 ndviValue, uint16[3] spectralData, uint8 psyllidDensity, bytes32 evidenceIPFSHash) returns()
func (_FarmRecordManager *FarmRecordManagerTransactor) AddRiskDetection(opts *bind.TransactOpts, farmId common.Address, gridCode string, gpsCenter [32]byte, riskLevel uint8, ndviValue uint16, spectralData [3]uint16, psyllidDensity uint8, evidenceIPFSHash [32]byte) (*types.Transaction, error) {
	return _FarmRecordManager.contract.Transact(opts, "addRiskDetection", farmId, gridCode, gpsCenter, riskLevel, ndviValue, spectralData, psyllidDensity, evidenceIPFSHash)
}

// AddRiskDetection is a paid mutator transaction binding the contract method 0x4807971a.
//
// Solidity: function addRiskDetection(address farmId, string gridCode, bytes32 gpsCenter, uint8 riskLevel, uint16 ndviValue, uint16[3] spectralData, uint8 psyllidDensity, bytes32 evidenceIPFSHash) returns()
func (_FarmRecordManager *FarmRecordManagerSession) AddRiskDetection(farmId common.Address, gridCode string, gpsCenter [32]byte, riskLevel uint8, ndviValue uint16, spectralData [3]uint16, psyllidDensity uint8, evidenceIPFSHash [32]byte) (*types.Transaction, error) {
	return _FarmRecordManager.Contract.AddRiskDetection(&_FarmRecordManager.TransactOpts, farmId, gridCode, gpsCenter, riskLevel, ndviValue, spectralData, psyllidDensity, evidenceIPFSHash)
}

// AddRiskDetection is a paid mutator transaction binding the contract method 0x4807971a.
//
// Solidity: function addRiskDetection(address farmId, string gridCode, bytes32 gpsCenter, uint8 riskLevel, uint16 ndviValue, uint16[3] spectralData, uint8 psyllidDensity, bytes32 evidenceIPFSHash) returns()
func (_FarmRecordManager *FarmRecordManagerTransactorSession) AddRiskDetection(farmId common.Address, gridCode string, gpsCenter [32]byte, riskLevel uint8, ndviValue uint16, spectralData [3]uint16, psyllidDensity uint8, evidenceIPFSHash [32]byte) (*types.Transaction, error) {
	return _FarmRecordManager.Contract.AddRiskDetection(&_FarmRecordManager.TransactOpts, farmId, gridCode, gpsCenter, riskLevel, ndviValue, spectralData, psyllidDensity, evidenceIPFSHash)
}

// AddTreeRemoval is a paid mutator transaction binding the contract method 0xf54dd1a0.
//
// Solidity: function addTreeRemoval(address farmId, string gridCode, uint32 linkedMarkingId, uint8 confirmedDiseaseCount, uint8 falsePositiveCount, bytes32 removalEvidenceIPFSHash, bool replantScheduled) returns()
func (_FarmRecordManager *FarmRecordManagerTransactor) AddTreeRemoval(opts *bind.TransactOpts, farmId common.Address, gridCode string, linkedMarkingId uint32, confirmedDiseaseCount uint8, falsePositiveCount uint8, removalEvidenceIPFSHash [32]byte, replantScheduled bool) (*types.Transaction, error) {
	return _FarmRecordManager.contract.Transact(opts, "addTreeRemoval", farmId, gridCode, linkedMarkingId, confirmedDiseaseCount, falsePositiveCount, removalEvidenceIPFSHash, replantScheduled)
}

// AddTreeRemoval is a paid mutator transaction binding the contract method 0xf54dd1a0.
//
// Solidity: function addTreeRemoval(address farmId, string gridCode, uint32 linkedMarkingId, uint8 confirmedDiseaseCount, uint8 falsePositiveCount, bytes32 removalEvidenceIPFSHash, bool replantScheduled) returns()
func (_FarmRecordManager *FarmRecordManagerSession) AddTreeRemoval(farmId common.Address, gridCode string, linkedMarkingId uint32, confirmedDiseaseCount uint8, falsePositiveCount uint8, removalEvidenceIPFSHash [32]byte, replantScheduled bool) (*types.Transaction, error) {
	return _FarmRecordManager.Contract.AddTreeRemoval(&_FarmRecordManager.TransactOpts, farmId, gridCode, linkedMarkingId, confirmedDiseaseCount, falsePositiveCount, removalEvidenceIPFSHash, replantScheduled)
}

// AddTreeRemoval is a paid mutator transaction binding the contract method 0xf54dd1a0.
//
// Solidity: function addTreeRemoval(address farmId, string gridCode, uint32 linkedMarkingId, uint8 confirmedDiseaseCount, uint8 falsePositiveCount, bytes32 removalEvidenceIPFSHash, bool replantScheduled) returns()
func (_FarmRecordManager *FarmRecordManagerTransactorSession) AddTreeRemoval(farmId common.Address, gridCode string, linkedMarkingId uint32, confirmedDiseaseCount uint8, falsePositiveCount uint8, removalEvidenceIPFSHash [32]byte, replantScheduled bool) (*types.Transaction, error) {
	return _FarmRecordManager.Contract.AddTreeRemoval(&_FarmRecordManager.TransactOpts, farmId, gridCode, linkedMarkingId, confirmedDiseaseCount, falsePositiveCount, removalEvidenceIPFSHash, replantScheduled)
}

// FarmRecordManagerGridMarkedIterator is returned from FilterGridMarked and is used to iterate over the raw logs and unpacked data for GridMarked events raised by the FarmRecordManager contract.
type FarmRecordManagerGridMarkedIterator struct {
	Event *FarmRecordManagerGridMarked // Event containing the contract specifics and raw log

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
func (it *FarmRecordManagerGridMarkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmRecordManagerGridMarked)
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
		it.Event = new(FarmRecordManagerGridMarked)
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
func (it *FarmRecordManagerGridMarkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmRecordManagerGridMarkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmRecordManagerGridMarked represents a GridMarked event raised by the FarmRecordManager contract.
type FarmRecordManagerGridMarked struct {
	FarmId      common.Address
	OperationId uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGridMarked is a free log retrieval operation binding the contract event 0x19a122899037db3fbe3cd7c78dd0e1b5c6282fabf1fb3e5c8fb17fc4d02a0a9d.
//
// Solidity: event GridMarked(address indexed farmId, uint32 indexed operationId)
func (_FarmRecordManager *FarmRecordManagerFilterer) FilterGridMarked(opts *bind.FilterOpts, farmId []common.Address, operationId []uint32) (*FarmRecordManagerGridMarkedIterator, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}
	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}

	logs, sub, err := _FarmRecordManager.contract.FilterLogs(opts, "GridMarked", farmIdRule, operationIdRule)
	if err != nil {
		return nil, err
	}
	return &FarmRecordManagerGridMarkedIterator{contract: _FarmRecordManager.contract, event: "GridMarked", logs: logs, sub: sub}, nil
}

// WatchGridMarked is a free log subscription operation binding the contract event 0x19a122899037db3fbe3cd7c78dd0e1b5c6282fabf1fb3e5c8fb17fc4d02a0a9d.
//
// Solidity: event GridMarked(address indexed farmId, uint32 indexed operationId)
func (_FarmRecordManager *FarmRecordManagerFilterer) WatchGridMarked(opts *bind.WatchOpts, sink chan<- *FarmRecordManagerGridMarked, farmId []common.Address, operationId []uint32) (event.Subscription, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}
	var operationIdRule []interface{}
	for _, operationIdItem := range operationId {
		operationIdRule = append(operationIdRule, operationIdItem)
	}

	logs, sub, err := _FarmRecordManager.contract.WatchLogs(opts, "GridMarked", farmIdRule, operationIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmRecordManagerGridMarked)
				if err := _FarmRecordManager.contract.UnpackLog(event, "GridMarked", log); err != nil {
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

// ParseGridMarked is a log parse operation binding the contract event 0x19a122899037db3fbe3cd7c78dd0e1b5c6282fabf1fb3e5c8fb17fc4d02a0a9d.
//
// Solidity: event GridMarked(address indexed farmId, uint32 indexed operationId)
func (_FarmRecordManager *FarmRecordManagerFilterer) ParseGridMarked(log types.Log) (*FarmRecordManagerGridMarked, error) {
	event := new(FarmRecordManagerGridMarked)
	if err := _FarmRecordManager.contract.UnpackLog(event, "GridMarked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmRecordManagerRiskDetectedIterator is returned from FilterRiskDetected and is used to iterate over the raw logs and unpacked data for RiskDetected events raised by the FarmRecordManager contract.
type FarmRecordManagerRiskDetectedIterator struct {
	Event *FarmRecordManagerRiskDetected // Event containing the contract specifics and raw log

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
func (it *FarmRecordManagerRiskDetectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmRecordManagerRiskDetected)
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
		it.Event = new(FarmRecordManagerRiskDetected)
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
func (it *FarmRecordManagerRiskDetectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmRecordManagerRiskDetectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmRecordManagerRiskDetected represents a RiskDetected event raised by the FarmRecordManager contract.
type FarmRecordManagerRiskDetected struct {
	FarmId    common.Address
	RecordId  uint32
	RiskLevel uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRiskDetected is a free log retrieval operation binding the contract event 0x6a00fb2be0846da567710c3968081b830f3a766a8f99be64c79c38d7e8238b8d.
//
// Solidity: event RiskDetected(address indexed farmId, uint32 indexed recordId, uint8 riskLevel)
func (_FarmRecordManager *FarmRecordManagerFilterer) FilterRiskDetected(opts *bind.FilterOpts, farmId []common.Address, recordId []uint32) (*FarmRecordManagerRiskDetectedIterator, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}
	var recordIdRule []interface{}
	for _, recordIdItem := range recordId {
		recordIdRule = append(recordIdRule, recordIdItem)
	}

	logs, sub, err := _FarmRecordManager.contract.FilterLogs(opts, "RiskDetected", farmIdRule, recordIdRule)
	if err != nil {
		return nil, err
	}
	return &FarmRecordManagerRiskDetectedIterator{contract: _FarmRecordManager.contract, event: "RiskDetected", logs: logs, sub: sub}, nil
}

// WatchRiskDetected is a free log subscription operation binding the contract event 0x6a00fb2be0846da567710c3968081b830f3a766a8f99be64c79c38d7e8238b8d.
//
// Solidity: event RiskDetected(address indexed farmId, uint32 indexed recordId, uint8 riskLevel)
func (_FarmRecordManager *FarmRecordManagerFilterer) WatchRiskDetected(opts *bind.WatchOpts, sink chan<- *FarmRecordManagerRiskDetected, farmId []common.Address, recordId []uint32) (event.Subscription, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}
	var recordIdRule []interface{}
	for _, recordIdItem := range recordId {
		recordIdRule = append(recordIdRule, recordIdItem)
	}

	logs, sub, err := _FarmRecordManager.contract.WatchLogs(opts, "RiskDetected", farmIdRule, recordIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmRecordManagerRiskDetected)
				if err := _FarmRecordManager.contract.UnpackLog(event, "RiskDetected", log); err != nil {
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

// ParseRiskDetected is a log parse operation binding the contract event 0x6a00fb2be0846da567710c3968081b830f3a766a8f99be64c79c38d7e8238b8d.
//
// Solidity: event RiskDetected(address indexed farmId, uint32 indexed recordId, uint8 riskLevel)
func (_FarmRecordManager *FarmRecordManagerFilterer) ParseRiskDetected(log types.Log) (*FarmRecordManagerRiskDetected, error) {
	event := new(FarmRecordManagerRiskDetected)
	if err := _FarmRecordManager.contract.UnpackLog(event, "RiskDetected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmRecordManagerTreesRemovedIterator is returned from FilterTreesRemoved and is used to iterate over the raw logs and unpacked data for TreesRemoved events raised by the FarmRecordManager contract.
type FarmRecordManagerTreesRemovedIterator struct {
	Event *FarmRecordManagerTreesRemoved // Event containing the contract specifics and raw log

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
func (it *FarmRecordManagerTreesRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmRecordManagerTreesRemoved)
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
		it.Event = new(FarmRecordManagerTreesRemoved)
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
func (it *FarmRecordManagerTreesRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmRecordManagerTreesRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmRecordManagerTreesRemoved represents a TreesRemoved event raised by the FarmRecordManager contract.
type FarmRecordManagerTreesRemoved struct {
	FarmId    common.Address
	RemovalId uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTreesRemoved is a free log retrieval operation binding the contract event 0x9b1e9a57d40236b579bd0b680c1c2f74f3e89abcdd941240245eb9b90204d167.
//
// Solidity: event TreesRemoved(address indexed farmId, uint32 indexed removalId)
func (_FarmRecordManager *FarmRecordManagerFilterer) FilterTreesRemoved(opts *bind.FilterOpts, farmId []common.Address, removalId []uint32) (*FarmRecordManagerTreesRemovedIterator, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}
	var removalIdRule []interface{}
	for _, removalIdItem := range removalId {
		removalIdRule = append(removalIdRule, removalIdItem)
	}

	logs, sub, err := _FarmRecordManager.contract.FilterLogs(opts, "TreesRemoved", farmIdRule, removalIdRule)
	if err != nil {
		return nil, err
	}
	return &FarmRecordManagerTreesRemovedIterator{contract: _FarmRecordManager.contract, event: "TreesRemoved", logs: logs, sub: sub}, nil
}

// WatchTreesRemoved is a free log subscription operation binding the contract event 0x9b1e9a57d40236b579bd0b680c1c2f74f3e89abcdd941240245eb9b90204d167.
//
// Solidity: event TreesRemoved(address indexed farmId, uint32 indexed removalId)
func (_FarmRecordManager *FarmRecordManagerFilterer) WatchTreesRemoved(opts *bind.WatchOpts, sink chan<- *FarmRecordManagerTreesRemoved, farmId []common.Address, removalId []uint32) (event.Subscription, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}
	var removalIdRule []interface{}
	for _, removalIdItem := range removalId {
		removalIdRule = append(removalIdRule, removalIdItem)
	}

	logs, sub, err := _FarmRecordManager.contract.WatchLogs(opts, "TreesRemoved", farmIdRule, removalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmRecordManagerTreesRemoved)
				if err := _FarmRecordManager.contract.UnpackLog(event, "TreesRemoved", log); err != nil {
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

// ParseTreesRemoved is a log parse operation binding the contract event 0x9b1e9a57d40236b579bd0b680c1c2f74f3e89abcdd941240245eb9b90204d167.
//
// Solidity: event TreesRemoved(address indexed farmId, uint32 indexed removalId)
func (_FarmRecordManager *FarmRecordManagerFilterer) ParseTreesRemoved(log types.Log) (*FarmRecordManagerTreesRemoved, error) {
	event := new(FarmRecordManagerTreesRemoved)
	if err := _FarmRecordManager.contract.UnpackLog(event, "TreesRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmRecordManagerTreesReplantedIterator is returned from FilterTreesReplanted and is used to iterate over the raw logs and unpacked data for TreesReplanted events raised by the FarmRecordManager contract.
type FarmRecordManagerTreesReplantedIterator struct {
	Event *FarmRecordManagerTreesReplanted // Event containing the contract specifics and raw log

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
func (it *FarmRecordManagerTreesReplantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmRecordManagerTreesReplanted)
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
		it.Event = new(FarmRecordManagerTreesReplanted)
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
func (it *FarmRecordManagerTreesReplantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmRecordManagerTreesReplantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmRecordManagerTreesReplanted represents a TreesReplanted event raised by the FarmRecordManager contract.
type FarmRecordManagerTreesReplanted struct {
	FarmId    common.Address
	ReplantId uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTreesReplanted is a free log retrieval operation binding the contract event 0xdc2f55a2706d870c6f77e220103ce48c45987347d38676fde648fecea4b49942.
//
// Solidity: event TreesReplanted(address indexed farmId, uint32 indexed replantId)
func (_FarmRecordManager *FarmRecordManagerFilterer) FilterTreesReplanted(opts *bind.FilterOpts, farmId []common.Address, replantId []uint32) (*FarmRecordManagerTreesReplantedIterator, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}
	var replantIdRule []interface{}
	for _, replantIdItem := range replantId {
		replantIdRule = append(replantIdRule, replantIdItem)
	}

	logs, sub, err := _FarmRecordManager.contract.FilterLogs(opts, "TreesReplanted", farmIdRule, replantIdRule)
	if err != nil {
		return nil, err
	}
	return &FarmRecordManagerTreesReplantedIterator{contract: _FarmRecordManager.contract, event: "TreesReplanted", logs: logs, sub: sub}, nil
}

// WatchTreesReplanted is a free log subscription operation binding the contract event 0xdc2f55a2706d870c6f77e220103ce48c45987347d38676fde648fecea4b49942.
//
// Solidity: event TreesReplanted(address indexed farmId, uint32 indexed replantId)
func (_FarmRecordManager *FarmRecordManagerFilterer) WatchTreesReplanted(opts *bind.WatchOpts, sink chan<- *FarmRecordManagerTreesReplanted, farmId []common.Address, replantId []uint32) (event.Subscription, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}
	var replantIdRule []interface{}
	for _, replantIdItem := range replantId {
		replantIdRule = append(replantIdRule, replantIdItem)
	}

	logs, sub, err := _FarmRecordManager.contract.WatchLogs(opts, "TreesReplanted", farmIdRule, replantIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmRecordManagerTreesReplanted)
				if err := _FarmRecordManager.contract.UnpackLog(event, "TreesReplanted", log); err != nil {
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

// ParseTreesReplanted is a log parse operation binding the contract event 0xdc2f55a2706d870c6f77e220103ce48c45987347d38676fde648fecea4b49942.
//
// Solidity: event TreesReplanted(address indexed farmId, uint32 indexed replantId)
func (_FarmRecordManager *FarmRecordManagerFilterer) ParseTreesReplanted(log types.Log) (*FarmRecordManagerTreesReplanted, error) {
	event := new(FarmRecordManagerTreesReplanted)
	if err := _FarmRecordManager.contract.UnpackLog(event, "TreesReplanted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
