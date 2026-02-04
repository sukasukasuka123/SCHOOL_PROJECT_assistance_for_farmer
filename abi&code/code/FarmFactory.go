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

// FarmFactoryDiseaseTreeRemovalRecord is an auto generated low-level Go binding around an user-defined struct.
type FarmFactoryDiseaseTreeRemovalRecord struct {
	RemovalId             uint32
	LinkedMarkingId       uint32
	Timestamp             uint32
	GridCode              string
	ConfirmedDiseaseCount uint8
	FalsePositiveCount    uint8
	RemovalGPS            [][32]byte
	RemovalEvidenceIPFS   string
	FarmerAddress         common.Address
	ReplantScheduled      bool
}

// FarmFactoryFarm is an auto generated low-level Go binding around an user-defined struct.
type FarmFactoryFarm struct {
	FarmId    common.Address
	Name      string
	Owner     common.Address
	GridCount *big.Int
}

// FarmFactoryMarkingOperationLog is an auto generated low-level Go binding around an user-defined struct.
type FarmFactoryMarkingOperationLog struct {
	OperationId          uint32
	LinkedDetectionId    uint32
	Timestamp            uint32
	GridCode             string
	MarkingPointCount    uint8
	MarkingGPS           [][32]byte
	MarkerType           string
	MarkerDosagePerPoint uint16
	IsolationPesticide   string
	IsolationDosage      uint16
	OperatorDroneId      common.Address
	OperationImageIPFS   string
}

// FarmFactoryReplantingRecord is an auto generated low-level Go binding around an user-defined struct.
type FarmFactoryReplantingRecord struct {
	ReplantId          uint32
	LinkedRemovalId    uint32
	Timestamp          uint32
	GridCode           string
	SeedlingCount      uint8
	ReplantGPS         [][32]byte
	SeedlingSource     string
	QuarantineCertHash string
	SeedlingVariety    string
	ReplantImageIPFS   string
	FarmerAddress      common.Address
}

// FarmFactoryRiskDetectionRecord is an auto generated low-level Go binding around an user-defined struct.
type FarmFactoryRiskDetectionRecord struct {
	RecordId          uint32
	Timestamp         uint32
	GridCode          string
	GpsCenter         [32]byte
	RiskLevel         uint8
	NdviValue         uint16
	SpectralData      [3]uint16
	PsyllidDensity    uint8
	EvidenceImageIPFS string
	OperatorDroneId   common.Address
}

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

// FarmFactoryMetaData contains all meta data concerning the FarmFactory contract.
var FarmFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_gridFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"FarmCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"}],\"name\":\"GridAddedToFarm\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"operationId\",\"type\":\"uint32\"}],\"name\":\"GridMarked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"recordId\",\"type\":\"uint32\"}],\"name\":\"RiskDetected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"removalId\",\"type\":\"uint32\"}],\"name\":\"TreesRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"replantId\",\"type\":\"uint32\"}],\"name\":\"TreesReplanted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"uint32\",\"name\":\"recordId\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"operationType\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"detail\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"evidenceIPFS\",\"type\":\"string\"}],\"name\":\"addGridMaintenance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"enumGridObjectFactory.GridTerrainType\",\"name\":\"terrainType\",\"type\":\"uint8\"}],\"name\":\"addGridToFarm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"operationId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkedDetectionId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"markingPointCount\",\"type\":\"uint8\"},{\"internalType\":\"bytes32[]\",\"name\":\"markingGPS\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"markerType\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"markerDosagePerPoint\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"isolationPesticide\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"isolationDosage\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"operatorDroneId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"operationImageIPFS\",\"type\":\"string\"}],\"internalType\":\"structFarmFactory.MarkingOperationLog\",\"name\":\"record\",\"type\":\"tuple\"}],\"name\":\"addMarkingOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"replantId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkedRemovalId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"seedlingCount\",\"type\":\"uint8\"},{\"internalType\":\"bytes32[]\",\"name\":\"replantGPS\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"seedlingSource\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"quarantineCertHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"seedlingVariety\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"replantImageIPFS\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"farmerAddress\",\"type\":\"address\"}],\"internalType\":\"structFarmFactory.ReplantingRecord\",\"name\":\"record\",\"type\":\"tuple\"}],\"name\":\"addReplanting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"recordId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"gpsCenter\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"riskLevel\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"ndviValue\",\"type\":\"uint16\"},{\"internalType\":\"uint16[3]\",\"name\":\"spectralData\",\"type\":\"uint16[3]\"},{\"internalType\":\"uint8\",\"name\":\"psyllidDensity\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"evidenceImageIPFS\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"operatorDroneId\",\"type\":\"address\"}],\"internalType\":\"structFarmFactory.RiskDetectionRecord\",\"name\":\"record\",\"type\":\"tuple\"}],\"name\":\"addRiskDetection\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"uint32\",\"name\":\"removalId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkedMarkingId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"confirmedDiseaseCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"falsePositiveCount\",\"type\":\"uint8\"},{\"internalType\":\"bytes32[]\",\"name\":\"removalGPS\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"removalEvidenceIPFS\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"farmerAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"replantScheduled\",\"type\":\"bool\"}],\"internalType\":\"structFarmFactory.DiseaseTreeRemovalRecord\",\"name\":\"record\",\"type\":\"tuple\"}],\"name\":\"addTreeRemoval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"createFarm\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"farmGridCodeToId\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"farmGridList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"farms\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gridCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"}],\"name\":\"getFarm\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gridCount\",\"type\":\"uint256\"}],\"internalType\":\"structFarmFactory.Farm\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"}],\"name\":\"getFarmGrids\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"}],\"name\":\"getGridIdByCode\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"}],\"name\":\"getGridInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"gridId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"enumGridObjectFactory.GridTerrainType\",\"name\":\"terrainType\",\"type\":\"uint8\"},{\"internalType\":\"enumGridObjectFactory.GridStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"enumGridObjectFactory.CitrusDiseaseType\",\"name\":\"diseaseType\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"lastStatusUpdateTime\",\"type\":\"uint32\"}],\"internalType\":\"structGridObjectFactory.GridInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"}],\"name\":\"getMarkingOperations\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"operationId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkedDetectionId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"markingPointCount\",\"type\":\"uint8\"},{\"internalType\":\"bytes32[]\",\"name\":\"markingGPS\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"markerType\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"markerDosagePerPoint\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"isolationPesticide\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"isolationDosage\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"operatorDroneId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"operationImageIPFS\",\"type\":\"string\"}],\"internalType\":\"structFarmFactory.MarkingOperationLog[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"}],\"name\":\"getRemovals\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"removalId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkedMarkingId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"confirmedDiseaseCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"falsePositiveCount\",\"type\":\"uint8\"},{\"internalType\":\"bytes32[]\",\"name\":\"removalGPS\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"removalEvidenceIPFS\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"farmerAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"replantScheduled\",\"type\":\"bool\"}],\"internalType\":\"structFarmFactory.DiseaseTreeRemovalRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"}],\"name\":\"getReplantings\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"replantId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkedRemovalId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"seedlingCount\",\"type\":\"uint8\"},{\"internalType\":\"bytes32[]\",\"name\":\"replantGPS\",\"type\":\"bytes32[]\"},{\"internalType\":\"string\",\"name\":\"seedlingSource\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"quarantineCertHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"seedlingVariety\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"replantImageIPFS\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"farmerAddress\",\"type\":\"address\"}],\"internalType\":\"structFarmFactory.ReplantingRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"}],\"name\":\"getRiskDetections\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"recordId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"gpsCenter\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"riskLevel\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"ndviValue\",\"type\":\"uint16\"},{\"internalType\":\"uint16[3]\",\"name\":\"spectralData\",\"type\":\"uint16[3]\"},{\"internalType\":\"uint8\",\"name\":\"psyllidDensity\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"evidenceImageIPFS\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"operatorDroneId\",\"type\":\"address\"}],\"internalType\":\"structFarmFactory.RiskDetectionRecord[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gridFactory\",\"outputs\":[{\"internalType\":\"contractIGridObjectFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"markingOperations\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"operationId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkedDetectionId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"markingPointCount\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"markerType\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"markerDosagePerPoint\",\"type\":\"uint16\"},{\"internalType\":\"string\",\"name\":\"isolationPesticide\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"isolationDosage\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"operatorDroneId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"operationImageIPFS\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"removals\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"removalId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkedMarkingId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"confirmedDiseaseCount\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"falsePositiveCount\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"removalEvidenceIPFS\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"farmerAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"replantScheduled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"replantings\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"replantId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"linkedRemovalId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"seedlingCount\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"seedlingSource\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"quarantineCertHash\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"seedlingVariety\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"replantImageIPFS\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"farmerAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"riskDetections\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"recordId\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"timestamp\",\"type\":\"uint32\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"gpsCenter\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"riskLevel\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"ndviValue\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"psyllidDensity\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"evidenceImageIPFS\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"operatorDroneId\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"enumGridObjectFactory.CitrusDiseaseType\",\"name\":\"newType\",\"type\":\"uint8\"}],\"name\":\"updateGridDiseaseType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"farmId\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"gridCode\",\"type\":\"string\"},{\"internalType\":\"enumGridObjectFactory.GridStatus\",\"name\":\"newStatus\",\"type\":\"uint8\"}],\"name\":\"updateGridStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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
// Solidity: function farms(address ) view returns(address farmId, string name, address owner, uint256 gridCount)
func (_FarmFactory *FarmFactoryCaller) Farms(opts *bind.CallOpts, arg0 common.Address) (struct {
	FarmId    common.Address
	Name      string
	Owner     common.Address
	GridCount *big.Int
}, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "farms", arg0)

	outstruct := new(struct {
		FarmId    common.Address
		Name      string
		Owner     common.Address
		GridCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FarmId = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Owner = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.GridCount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Farms is a free data retrieval call binding the contract method 0x421adfa0.
//
// Solidity: function farms(address ) view returns(address farmId, string name, address owner, uint256 gridCount)
func (_FarmFactory *FarmFactorySession) Farms(arg0 common.Address) (struct {
	FarmId    common.Address
	Name      string
	Owner     common.Address
	GridCount *big.Int
}, error) {
	return _FarmFactory.Contract.Farms(&_FarmFactory.CallOpts, arg0)
}

// Farms is a free data retrieval call binding the contract method 0x421adfa0.
//
// Solidity: function farms(address ) view returns(address farmId, string name, address owner, uint256 gridCount)
func (_FarmFactory *FarmFactoryCallerSession) Farms(arg0 common.Address) (struct {
	FarmId    common.Address
	Name      string
	Owner     common.Address
	GridCount *big.Int
}, error) {
	return _FarmFactory.Contract.Farms(&_FarmFactory.CallOpts, arg0)
}

// GetFarm is a free data retrieval call binding the contract method 0x3bfc187c.
//
// Solidity: function getFarm(address farmId) view returns((address,string,address,uint256))
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
// Solidity: function getFarm(address farmId) view returns((address,string,address,uint256))
func (_FarmFactory *FarmFactorySession) GetFarm(farmId common.Address) (FarmFactoryFarm, error) {
	return _FarmFactory.Contract.GetFarm(&_FarmFactory.CallOpts, farmId)
}

// GetFarm is a free data retrieval call binding the contract method 0x3bfc187c.
//
// Solidity: function getFarm(address farmId) view returns((address,string,address,uint256))
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
// Solidity: function getGridInfo(address farmId, string gridCode) view returns((address,string,address,uint8,uint8,uint8,uint32))
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
// Solidity: function getGridInfo(address farmId, string gridCode) view returns((address,string,address,uint8,uint8,uint8,uint32))
func (_FarmFactory *FarmFactorySession) GetGridInfo(farmId common.Address, gridCode string) (GridObjectFactoryGridInfo, error) {
	return _FarmFactory.Contract.GetGridInfo(&_FarmFactory.CallOpts, farmId, gridCode)
}

// GetGridInfo is a free data retrieval call binding the contract method 0x98c0a3e5.
//
// Solidity: function getGridInfo(address farmId, string gridCode) view returns((address,string,address,uint8,uint8,uint8,uint32))
func (_FarmFactory *FarmFactoryCallerSession) GetGridInfo(farmId common.Address, gridCode string) (GridObjectFactoryGridInfo, error) {
	return _FarmFactory.Contract.GetGridInfo(&_FarmFactory.CallOpts, farmId, gridCode)
}

// GetMarkingOperations is a free data retrieval call binding the contract method 0x7190033b.
//
// Solidity: function getMarkingOperations(address farmId) view returns((uint32,uint32,uint32,string,uint8,bytes32[],string,uint16,string,uint16,address,string)[])
func (_FarmFactory *FarmFactoryCaller) GetMarkingOperations(opts *bind.CallOpts, farmId common.Address) ([]FarmFactoryMarkingOperationLog, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "getMarkingOperations", farmId)

	if err != nil {
		return *new([]FarmFactoryMarkingOperationLog), err
	}

	out0 := *abi.ConvertType(out[0], new([]FarmFactoryMarkingOperationLog)).(*[]FarmFactoryMarkingOperationLog)

	return out0, err

}

// GetMarkingOperations is a free data retrieval call binding the contract method 0x7190033b.
//
// Solidity: function getMarkingOperations(address farmId) view returns((uint32,uint32,uint32,string,uint8,bytes32[],string,uint16,string,uint16,address,string)[])
func (_FarmFactory *FarmFactorySession) GetMarkingOperations(farmId common.Address) ([]FarmFactoryMarkingOperationLog, error) {
	return _FarmFactory.Contract.GetMarkingOperations(&_FarmFactory.CallOpts, farmId)
}

// GetMarkingOperations is a free data retrieval call binding the contract method 0x7190033b.
//
// Solidity: function getMarkingOperations(address farmId) view returns((uint32,uint32,uint32,string,uint8,bytes32[],string,uint16,string,uint16,address,string)[])
func (_FarmFactory *FarmFactoryCallerSession) GetMarkingOperations(farmId common.Address) ([]FarmFactoryMarkingOperationLog, error) {
	return _FarmFactory.Contract.GetMarkingOperations(&_FarmFactory.CallOpts, farmId)
}

// GetRemovals is a free data retrieval call binding the contract method 0x3e669f9b.
//
// Solidity: function getRemovals(address farmId) view returns((uint32,uint32,uint32,string,uint8,uint8,bytes32[],string,address,bool)[])
func (_FarmFactory *FarmFactoryCaller) GetRemovals(opts *bind.CallOpts, farmId common.Address) ([]FarmFactoryDiseaseTreeRemovalRecord, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "getRemovals", farmId)

	if err != nil {
		return *new([]FarmFactoryDiseaseTreeRemovalRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]FarmFactoryDiseaseTreeRemovalRecord)).(*[]FarmFactoryDiseaseTreeRemovalRecord)

	return out0, err

}

// GetRemovals is a free data retrieval call binding the contract method 0x3e669f9b.
//
// Solidity: function getRemovals(address farmId) view returns((uint32,uint32,uint32,string,uint8,uint8,bytes32[],string,address,bool)[])
func (_FarmFactory *FarmFactorySession) GetRemovals(farmId common.Address) ([]FarmFactoryDiseaseTreeRemovalRecord, error) {
	return _FarmFactory.Contract.GetRemovals(&_FarmFactory.CallOpts, farmId)
}

// GetRemovals is a free data retrieval call binding the contract method 0x3e669f9b.
//
// Solidity: function getRemovals(address farmId) view returns((uint32,uint32,uint32,string,uint8,uint8,bytes32[],string,address,bool)[])
func (_FarmFactory *FarmFactoryCallerSession) GetRemovals(farmId common.Address) ([]FarmFactoryDiseaseTreeRemovalRecord, error) {
	return _FarmFactory.Contract.GetRemovals(&_FarmFactory.CallOpts, farmId)
}

// GetReplantings is a free data retrieval call binding the contract method 0xb69a0bf0.
//
// Solidity: function getReplantings(address farmId) view returns((uint32,uint32,uint32,string,uint8,bytes32[],string,string,string,string,address)[])
func (_FarmFactory *FarmFactoryCaller) GetReplantings(opts *bind.CallOpts, farmId common.Address) ([]FarmFactoryReplantingRecord, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "getReplantings", farmId)

	if err != nil {
		return *new([]FarmFactoryReplantingRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]FarmFactoryReplantingRecord)).(*[]FarmFactoryReplantingRecord)

	return out0, err

}

// GetReplantings is a free data retrieval call binding the contract method 0xb69a0bf0.
//
// Solidity: function getReplantings(address farmId) view returns((uint32,uint32,uint32,string,uint8,bytes32[],string,string,string,string,address)[])
func (_FarmFactory *FarmFactorySession) GetReplantings(farmId common.Address) ([]FarmFactoryReplantingRecord, error) {
	return _FarmFactory.Contract.GetReplantings(&_FarmFactory.CallOpts, farmId)
}

// GetReplantings is a free data retrieval call binding the contract method 0xb69a0bf0.
//
// Solidity: function getReplantings(address farmId) view returns((uint32,uint32,uint32,string,uint8,bytes32[],string,string,string,string,address)[])
func (_FarmFactory *FarmFactoryCallerSession) GetReplantings(farmId common.Address) ([]FarmFactoryReplantingRecord, error) {
	return _FarmFactory.Contract.GetReplantings(&_FarmFactory.CallOpts, farmId)
}

// GetRiskDetections is a free data retrieval call binding the contract method 0x6b4152ee.
//
// Solidity: function getRiskDetections(address farmId) view returns((uint32,uint32,string,bytes32,uint8,uint16,uint16[3],uint8,string,address)[])
func (_FarmFactory *FarmFactoryCaller) GetRiskDetections(opts *bind.CallOpts, farmId common.Address) ([]FarmFactoryRiskDetectionRecord, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "getRiskDetections", farmId)

	if err != nil {
		return *new([]FarmFactoryRiskDetectionRecord), err
	}

	out0 := *abi.ConvertType(out[0], new([]FarmFactoryRiskDetectionRecord)).(*[]FarmFactoryRiskDetectionRecord)

	return out0, err

}

// GetRiskDetections is a free data retrieval call binding the contract method 0x6b4152ee.
//
// Solidity: function getRiskDetections(address farmId) view returns((uint32,uint32,string,bytes32,uint8,uint16,uint16[3],uint8,string,address)[])
func (_FarmFactory *FarmFactorySession) GetRiskDetections(farmId common.Address) ([]FarmFactoryRiskDetectionRecord, error) {
	return _FarmFactory.Contract.GetRiskDetections(&_FarmFactory.CallOpts, farmId)
}

// GetRiskDetections is a free data retrieval call binding the contract method 0x6b4152ee.
//
// Solidity: function getRiskDetections(address farmId) view returns((uint32,uint32,string,bytes32,uint8,uint16,uint16[3],uint8,string,address)[])
func (_FarmFactory *FarmFactoryCallerSession) GetRiskDetections(farmId common.Address) ([]FarmFactoryRiskDetectionRecord, error) {
	return _FarmFactory.Contract.GetRiskDetections(&_FarmFactory.CallOpts, farmId)
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

// MarkingOperations is a free data retrieval call binding the contract method 0xa305ed02.
//
// Solidity: function markingOperations(address , uint256 ) view returns(uint32 operationId, uint32 linkedDetectionId, uint32 timestamp, string gridCode, uint8 markingPointCount, string markerType, uint16 markerDosagePerPoint, string isolationPesticide, uint16 isolationDosage, address operatorDroneId, string operationImageIPFS)
func (_FarmFactory *FarmFactoryCaller) MarkingOperations(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	OperationId          uint32
	LinkedDetectionId    uint32
	Timestamp            uint32
	GridCode             string
	MarkingPointCount    uint8
	MarkerType           string
	MarkerDosagePerPoint uint16
	IsolationPesticide   string
	IsolationDosage      uint16
	OperatorDroneId      common.Address
	OperationImageIPFS   string
}, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "markingOperations", arg0, arg1)

	outstruct := new(struct {
		OperationId          uint32
		LinkedDetectionId    uint32
		Timestamp            uint32
		GridCode             string
		MarkingPointCount    uint8
		MarkerType           string
		MarkerDosagePerPoint uint16
		IsolationPesticide   string
		IsolationDosage      uint16
		OperatorDroneId      common.Address
		OperationImageIPFS   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.OperationId = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.LinkedDetectionId = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.GridCode = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.MarkingPointCount = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.MarkerType = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.MarkerDosagePerPoint = *abi.ConvertType(out[6], new(uint16)).(*uint16)
	outstruct.IsolationPesticide = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.IsolationDosage = *abi.ConvertType(out[8], new(uint16)).(*uint16)
	outstruct.OperatorDroneId = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)
	outstruct.OperationImageIPFS = *abi.ConvertType(out[10], new(string)).(*string)

	return *outstruct, err

}

// MarkingOperations is a free data retrieval call binding the contract method 0xa305ed02.
//
// Solidity: function markingOperations(address , uint256 ) view returns(uint32 operationId, uint32 linkedDetectionId, uint32 timestamp, string gridCode, uint8 markingPointCount, string markerType, uint16 markerDosagePerPoint, string isolationPesticide, uint16 isolationDosage, address operatorDroneId, string operationImageIPFS)
func (_FarmFactory *FarmFactorySession) MarkingOperations(arg0 common.Address, arg1 *big.Int) (struct {
	OperationId          uint32
	LinkedDetectionId    uint32
	Timestamp            uint32
	GridCode             string
	MarkingPointCount    uint8
	MarkerType           string
	MarkerDosagePerPoint uint16
	IsolationPesticide   string
	IsolationDosage      uint16
	OperatorDroneId      common.Address
	OperationImageIPFS   string
}, error) {
	return _FarmFactory.Contract.MarkingOperations(&_FarmFactory.CallOpts, arg0, arg1)
}

// MarkingOperations is a free data retrieval call binding the contract method 0xa305ed02.
//
// Solidity: function markingOperations(address , uint256 ) view returns(uint32 operationId, uint32 linkedDetectionId, uint32 timestamp, string gridCode, uint8 markingPointCount, string markerType, uint16 markerDosagePerPoint, string isolationPesticide, uint16 isolationDosage, address operatorDroneId, string operationImageIPFS)
func (_FarmFactory *FarmFactoryCallerSession) MarkingOperations(arg0 common.Address, arg1 *big.Int) (struct {
	OperationId          uint32
	LinkedDetectionId    uint32
	Timestamp            uint32
	GridCode             string
	MarkingPointCount    uint8
	MarkerType           string
	MarkerDosagePerPoint uint16
	IsolationPesticide   string
	IsolationDosage      uint16
	OperatorDroneId      common.Address
	OperationImageIPFS   string
}, error) {
	return _FarmFactory.Contract.MarkingOperations(&_FarmFactory.CallOpts, arg0, arg1)
}

// Removals is a free data retrieval call binding the contract method 0x0ddbe234.
//
// Solidity: function removals(address , uint256 ) view returns(uint32 removalId, uint32 linkedMarkingId, uint32 timestamp, string gridCode, uint8 confirmedDiseaseCount, uint8 falsePositiveCount, string removalEvidenceIPFS, address farmerAddress, bool replantScheduled)
func (_FarmFactory *FarmFactoryCaller) Removals(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	RemovalId             uint32
	LinkedMarkingId       uint32
	Timestamp             uint32
	GridCode              string
	ConfirmedDiseaseCount uint8
	FalsePositiveCount    uint8
	RemovalEvidenceIPFS   string
	FarmerAddress         common.Address
	ReplantScheduled      bool
}, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "removals", arg0, arg1)

	outstruct := new(struct {
		RemovalId             uint32
		LinkedMarkingId       uint32
		Timestamp             uint32
		GridCode              string
		ConfirmedDiseaseCount uint8
		FalsePositiveCount    uint8
		RemovalEvidenceIPFS   string
		FarmerAddress         common.Address
		ReplantScheduled      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RemovalId = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.LinkedMarkingId = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.GridCode = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.ConfirmedDiseaseCount = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.FalsePositiveCount = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.RemovalEvidenceIPFS = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.FarmerAddress = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.ReplantScheduled = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// Removals is a free data retrieval call binding the contract method 0x0ddbe234.
//
// Solidity: function removals(address , uint256 ) view returns(uint32 removalId, uint32 linkedMarkingId, uint32 timestamp, string gridCode, uint8 confirmedDiseaseCount, uint8 falsePositiveCount, string removalEvidenceIPFS, address farmerAddress, bool replantScheduled)
func (_FarmFactory *FarmFactorySession) Removals(arg0 common.Address, arg1 *big.Int) (struct {
	RemovalId             uint32
	LinkedMarkingId       uint32
	Timestamp             uint32
	GridCode              string
	ConfirmedDiseaseCount uint8
	FalsePositiveCount    uint8
	RemovalEvidenceIPFS   string
	FarmerAddress         common.Address
	ReplantScheduled      bool
}, error) {
	return _FarmFactory.Contract.Removals(&_FarmFactory.CallOpts, arg0, arg1)
}

// Removals is a free data retrieval call binding the contract method 0x0ddbe234.
//
// Solidity: function removals(address , uint256 ) view returns(uint32 removalId, uint32 linkedMarkingId, uint32 timestamp, string gridCode, uint8 confirmedDiseaseCount, uint8 falsePositiveCount, string removalEvidenceIPFS, address farmerAddress, bool replantScheduled)
func (_FarmFactory *FarmFactoryCallerSession) Removals(arg0 common.Address, arg1 *big.Int) (struct {
	RemovalId             uint32
	LinkedMarkingId       uint32
	Timestamp             uint32
	GridCode              string
	ConfirmedDiseaseCount uint8
	FalsePositiveCount    uint8
	RemovalEvidenceIPFS   string
	FarmerAddress         common.Address
	ReplantScheduled      bool
}, error) {
	return _FarmFactory.Contract.Removals(&_FarmFactory.CallOpts, arg0, arg1)
}

// Replantings is a free data retrieval call binding the contract method 0x5c351940.
//
// Solidity: function replantings(address , uint256 ) view returns(uint32 replantId, uint32 linkedRemovalId, uint32 timestamp, string gridCode, uint8 seedlingCount, string seedlingSource, string quarantineCertHash, string seedlingVariety, string replantImageIPFS, address farmerAddress)
func (_FarmFactory *FarmFactoryCaller) Replantings(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	ReplantId          uint32
	LinkedRemovalId    uint32
	Timestamp          uint32
	GridCode           string
	SeedlingCount      uint8
	SeedlingSource     string
	QuarantineCertHash string
	SeedlingVariety    string
	ReplantImageIPFS   string
	FarmerAddress      common.Address
}, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "replantings", arg0, arg1)

	outstruct := new(struct {
		ReplantId          uint32
		LinkedRemovalId    uint32
		Timestamp          uint32
		GridCode           string
		SeedlingCount      uint8
		SeedlingSource     string
		QuarantineCertHash string
		SeedlingVariety    string
		ReplantImageIPFS   string
		FarmerAddress      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReplantId = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.LinkedRemovalId = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.Timestamp = *abi.ConvertType(out[2], new(uint32)).(*uint32)
	outstruct.GridCode = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.SeedlingCount = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.SeedlingSource = *abi.ConvertType(out[5], new(string)).(*string)
	outstruct.QuarantineCertHash = *abi.ConvertType(out[6], new(string)).(*string)
	outstruct.SeedlingVariety = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.ReplantImageIPFS = *abi.ConvertType(out[8], new(string)).(*string)
	outstruct.FarmerAddress = *abi.ConvertType(out[9], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Replantings is a free data retrieval call binding the contract method 0x5c351940.
//
// Solidity: function replantings(address , uint256 ) view returns(uint32 replantId, uint32 linkedRemovalId, uint32 timestamp, string gridCode, uint8 seedlingCount, string seedlingSource, string quarantineCertHash, string seedlingVariety, string replantImageIPFS, address farmerAddress)
func (_FarmFactory *FarmFactorySession) Replantings(arg0 common.Address, arg1 *big.Int) (struct {
	ReplantId          uint32
	LinkedRemovalId    uint32
	Timestamp          uint32
	GridCode           string
	SeedlingCount      uint8
	SeedlingSource     string
	QuarantineCertHash string
	SeedlingVariety    string
	ReplantImageIPFS   string
	FarmerAddress      common.Address
}, error) {
	return _FarmFactory.Contract.Replantings(&_FarmFactory.CallOpts, arg0, arg1)
}

// Replantings is a free data retrieval call binding the contract method 0x5c351940.
//
// Solidity: function replantings(address , uint256 ) view returns(uint32 replantId, uint32 linkedRemovalId, uint32 timestamp, string gridCode, uint8 seedlingCount, string seedlingSource, string quarantineCertHash, string seedlingVariety, string replantImageIPFS, address farmerAddress)
func (_FarmFactory *FarmFactoryCallerSession) Replantings(arg0 common.Address, arg1 *big.Int) (struct {
	ReplantId          uint32
	LinkedRemovalId    uint32
	Timestamp          uint32
	GridCode           string
	SeedlingCount      uint8
	SeedlingSource     string
	QuarantineCertHash string
	SeedlingVariety    string
	ReplantImageIPFS   string
	FarmerAddress      common.Address
}, error) {
	return _FarmFactory.Contract.Replantings(&_FarmFactory.CallOpts, arg0, arg1)
}

// RiskDetections is a free data retrieval call binding the contract method 0xc48502b5.
//
// Solidity: function riskDetections(address , uint256 ) view returns(uint32 recordId, uint32 timestamp, string gridCode, bytes32 gpsCenter, uint8 riskLevel, uint16 ndviValue, uint8 psyllidDensity, string evidenceImageIPFS, address operatorDroneId)
func (_FarmFactory *FarmFactoryCaller) RiskDetections(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	RecordId          uint32
	Timestamp         uint32
	GridCode          string
	GpsCenter         [32]byte
	RiskLevel         uint8
	NdviValue         uint16
	PsyllidDensity    uint8
	EvidenceImageIPFS string
	OperatorDroneId   common.Address
}, error) {
	var out []interface{}
	err := _FarmFactory.contract.Call(opts, &out, "riskDetections", arg0, arg1)

	outstruct := new(struct {
		RecordId          uint32
		Timestamp         uint32
		GridCode          string
		GpsCenter         [32]byte
		RiskLevel         uint8
		NdviValue         uint16
		PsyllidDensity    uint8
		EvidenceImageIPFS string
		OperatorDroneId   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RecordId = *abi.ConvertType(out[0], new(uint32)).(*uint32)
	outstruct.Timestamp = *abi.ConvertType(out[1], new(uint32)).(*uint32)
	outstruct.GridCode = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.GpsCenter = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.RiskLevel = *abi.ConvertType(out[4], new(uint8)).(*uint8)
	outstruct.NdviValue = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.PsyllidDensity = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.EvidenceImageIPFS = *abi.ConvertType(out[7], new(string)).(*string)
	outstruct.OperatorDroneId = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// RiskDetections is a free data retrieval call binding the contract method 0xc48502b5.
//
// Solidity: function riskDetections(address , uint256 ) view returns(uint32 recordId, uint32 timestamp, string gridCode, bytes32 gpsCenter, uint8 riskLevel, uint16 ndviValue, uint8 psyllidDensity, string evidenceImageIPFS, address operatorDroneId)
func (_FarmFactory *FarmFactorySession) RiskDetections(arg0 common.Address, arg1 *big.Int) (struct {
	RecordId          uint32
	Timestamp         uint32
	GridCode          string
	GpsCenter         [32]byte
	RiskLevel         uint8
	NdviValue         uint16
	PsyllidDensity    uint8
	EvidenceImageIPFS string
	OperatorDroneId   common.Address
}, error) {
	return _FarmFactory.Contract.RiskDetections(&_FarmFactory.CallOpts, arg0, arg1)
}

// RiskDetections is a free data retrieval call binding the contract method 0xc48502b5.
//
// Solidity: function riskDetections(address , uint256 ) view returns(uint32 recordId, uint32 timestamp, string gridCode, bytes32 gpsCenter, uint8 riskLevel, uint16 ndviValue, uint8 psyllidDensity, string evidenceImageIPFS, address operatorDroneId)
func (_FarmFactory *FarmFactoryCallerSession) RiskDetections(arg0 common.Address, arg1 *big.Int) (struct {
	RecordId          uint32
	Timestamp         uint32
	GridCode          string
	GpsCenter         [32]byte
	RiskLevel         uint8
	NdviValue         uint16
	PsyllidDensity    uint8
	EvidenceImageIPFS string
	OperatorDroneId   common.Address
}, error) {
	return _FarmFactory.Contract.RiskDetections(&_FarmFactory.CallOpts, arg0, arg1)
}

// AddGridMaintenance is a paid mutator transaction binding the contract method 0x298ad68b.
//
// Solidity: function addGridMaintenance(address farmId, string gridCode, uint32 recordId, string operationType, string detail, string evidenceIPFS) returns()
func (_FarmFactory *FarmFactoryTransactor) AddGridMaintenance(opts *bind.TransactOpts, farmId common.Address, gridCode string, recordId uint32, operationType string, detail string, evidenceIPFS string) (*types.Transaction, error) {
	return _FarmFactory.contract.Transact(opts, "addGridMaintenance", farmId, gridCode, recordId, operationType, detail, evidenceIPFS)
}

// AddGridMaintenance is a paid mutator transaction binding the contract method 0x298ad68b.
//
// Solidity: function addGridMaintenance(address farmId, string gridCode, uint32 recordId, string operationType, string detail, string evidenceIPFS) returns()
func (_FarmFactory *FarmFactorySession) AddGridMaintenance(farmId common.Address, gridCode string, recordId uint32, operationType string, detail string, evidenceIPFS string) (*types.Transaction, error) {
	return _FarmFactory.Contract.AddGridMaintenance(&_FarmFactory.TransactOpts, farmId, gridCode, recordId, operationType, detail, evidenceIPFS)
}

// AddGridMaintenance is a paid mutator transaction binding the contract method 0x298ad68b.
//
// Solidity: function addGridMaintenance(address farmId, string gridCode, uint32 recordId, string operationType, string detail, string evidenceIPFS) returns()
func (_FarmFactory *FarmFactoryTransactorSession) AddGridMaintenance(farmId common.Address, gridCode string, recordId uint32, operationType string, detail string, evidenceIPFS string) (*types.Transaction, error) {
	return _FarmFactory.Contract.AddGridMaintenance(&_FarmFactory.TransactOpts, farmId, gridCode, recordId, operationType, detail, evidenceIPFS)
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

// AddMarkingOperation is a paid mutator transaction binding the contract method 0x8d0496b7.
//
// Solidity: function addMarkingOperation(address farmId, (uint32,uint32,uint32,string,uint8,bytes32[],string,uint16,string,uint16,address,string) record) returns()
func (_FarmFactory *FarmFactoryTransactor) AddMarkingOperation(opts *bind.TransactOpts, farmId common.Address, record FarmFactoryMarkingOperationLog) (*types.Transaction, error) {
	return _FarmFactory.contract.Transact(opts, "addMarkingOperation", farmId, record)
}

// AddMarkingOperation is a paid mutator transaction binding the contract method 0x8d0496b7.
//
// Solidity: function addMarkingOperation(address farmId, (uint32,uint32,uint32,string,uint8,bytes32[],string,uint16,string,uint16,address,string) record) returns()
func (_FarmFactory *FarmFactorySession) AddMarkingOperation(farmId common.Address, record FarmFactoryMarkingOperationLog) (*types.Transaction, error) {
	return _FarmFactory.Contract.AddMarkingOperation(&_FarmFactory.TransactOpts, farmId, record)
}

// AddMarkingOperation is a paid mutator transaction binding the contract method 0x8d0496b7.
//
// Solidity: function addMarkingOperation(address farmId, (uint32,uint32,uint32,string,uint8,bytes32[],string,uint16,string,uint16,address,string) record) returns()
func (_FarmFactory *FarmFactoryTransactorSession) AddMarkingOperation(farmId common.Address, record FarmFactoryMarkingOperationLog) (*types.Transaction, error) {
	return _FarmFactory.Contract.AddMarkingOperation(&_FarmFactory.TransactOpts, farmId, record)
}

// AddReplanting is a paid mutator transaction binding the contract method 0x945dcd17.
//
// Solidity: function addReplanting(address farmId, (uint32,uint32,uint32,string,uint8,bytes32[],string,string,string,string,address) record) returns()
func (_FarmFactory *FarmFactoryTransactor) AddReplanting(opts *bind.TransactOpts, farmId common.Address, record FarmFactoryReplantingRecord) (*types.Transaction, error) {
	return _FarmFactory.contract.Transact(opts, "addReplanting", farmId, record)
}

// AddReplanting is a paid mutator transaction binding the contract method 0x945dcd17.
//
// Solidity: function addReplanting(address farmId, (uint32,uint32,uint32,string,uint8,bytes32[],string,string,string,string,address) record) returns()
func (_FarmFactory *FarmFactorySession) AddReplanting(farmId common.Address, record FarmFactoryReplantingRecord) (*types.Transaction, error) {
	return _FarmFactory.Contract.AddReplanting(&_FarmFactory.TransactOpts, farmId, record)
}

// AddReplanting is a paid mutator transaction binding the contract method 0x945dcd17.
//
// Solidity: function addReplanting(address farmId, (uint32,uint32,uint32,string,uint8,bytes32[],string,string,string,string,address) record) returns()
func (_FarmFactory *FarmFactoryTransactorSession) AddReplanting(farmId common.Address, record FarmFactoryReplantingRecord) (*types.Transaction, error) {
	return _FarmFactory.Contract.AddReplanting(&_FarmFactory.TransactOpts, farmId, record)
}

// AddRiskDetection is a paid mutator transaction binding the contract method 0x7f6a7ead.
//
// Solidity: function addRiskDetection(address farmId, (uint32,uint32,string,bytes32,uint8,uint16,uint16[3],uint8,string,address) record) returns()
func (_FarmFactory *FarmFactoryTransactor) AddRiskDetection(opts *bind.TransactOpts, farmId common.Address, record FarmFactoryRiskDetectionRecord) (*types.Transaction, error) {
	return _FarmFactory.contract.Transact(opts, "addRiskDetection", farmId, record)
}

// AddRiskDetection is a paid mutator transaction binding the contract method 0x7f6a7ead.
//
// Solidity: function addRiskDetection(address farmId, (uint32,uint32,string,bytes32,uint8,uint16,uint16[3],uint8,string,address) record) returns()
func (_FarmFactory *FarmFactorySession) AddRiskDetection(farmId common.Address, record FarmFactoryRiskDetectionRecord) (*types.Transaction, error) {
	return _FarmFactory.Contract.AddRiskDetection(&_FarmFactory.TransactOpts, farmId, record)
}

// AddRiskDetection is a paid mutator transaction binding the contract method 0x7f6a7ead.
//
// Solidity: function addRiskDetection(address farmId, (uint32,uint32,string,bytes32,uint8,uint16,uint16[3],uint8,string,address) record) returns()
func (_FarmFactory *FarmFactoryTransactorSession) AddRiskDetection(farmId common.Address, record FarmFactoryRiskDetectionRecord) (*types.Transaction, error) {
	return _FarmFactory.Contract.AddRiskDetection(&_FarmFactory.TransactOpts, farmId, record)
}

// AddTreeRemoval is a paid mutator transaction binding the contract method 0xf85a68c8.
//
// Solidity: function addTreeRemoval(address farmId, (uint32,uint32,uint32,string,uint8,uint8,bytes32[],string,address,bool) record) returns()
func (_FarmFactory *FarmFactoryTransactor) AddTreeRemoval(opts *bind.TransactOpts, farmId common.Address, record FarmFactoryDiseaseTreeRemovalRecord) (*types.Transaction, error) {
	return _FarmFactory.contract.Transact(opts, "addTreeRemoval", farmId, record)
}

// AddTreeRemoval is a paid mutator transaction binding the contract method 0xf85a68c8.
//
// Solidity: function addTreeRemoval(address farmId, (uint32,uint32,uint32,string,uint8,uint8,bytes32[],string,address,bool) record) returns()
func (_FarmFactory *FarmFactorySession) AddTreeRemoval(farmId common.Address, record FarmFactoryDiseaseTreeRemovalRecord) (*types.Transaction, error) {
	return _FarmFactory.Contract.AddTreeRemoval(&_FarmFactory.TransactOpts, farmId, record)
}

// AddTreeRemoval is a paid mutator transaction binding the contract method 0xf85a68c8.
//
// Solidity: function addTreeRemoval(address farmId, (uint32,uint32,uint32,string,uint8,uint8,bytes32[],string,address,bool) record) returns()
func (_FarmFactory *FarmFactoryTransactorSession) AddTreeRemoval(farmId common.Address, record FarmFactoryDiseaseTreeRemovalRecord) (*types.Transaction, error) {
	return _FarmFactory.Contract.AddTreeRemoval(&_FarmFactory.TransactOpts, farmId, record)
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
// Solidity: event FarmCreated(address indexed farmId, string name, address owner)
func (_FarmFactory *FarmFactoryFilterer) FilterFarmCreated(opts *bind.FilterOpts, farmId []common.Address) (*FarmFactoryFarmCreatedIterator, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}

	logs, sub, err := _FarmFactory.contract.FilterLogs(opts, "FarmCreated", farmIdRule)
	if err != nil {
		return nil, err
	}
	return &FarmFactoryFarmCreatedIterator{contract: _FarmFactory.contract, event: "FarmCreated", logs: logs, sub: sub}, nil
}

// WatchFarmCreated is a free log subscription operation binding the contract event 0xcfb9790749b5d99475809cf22fdf8e7b1992ea93e7bd0513964e36ef06a1e60d.
//
// Solidity: event FarmCreated(address indexed farmId, string name, address owner)
func (_FarmFactory *FarmFactoryFilterer) WatchFarmCreated(opts *bind.WatchOpts, sink chan<- *FarmFactoryFarmCreated, farmId []common.Address) (event.Subscription, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}

	logs, sub, err := _FarmFactory.contract.WatchLogs(opts, "FarmCreated", farmIdRule)
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
// Solidity: event FarmCreated(address indexed farmId, string name, address owner)
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

// FarmFactoryGridMarkedIterator is returned from FilterGridMarked and is used to iterate over the raw logs and unpacked data for GridMarked events raised by the FarmFactory contract.
type FarmFactoryGridMarkedIterator struct {
	Event *FarmFactoryGridMarked // Event containing the contract specifics and raw log

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
func (it *FarmFactoryGridMarkedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmFactoryGridMarked)
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
		it.Event = new(FarmFactoryGridMarked)
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
func (it *FarmFactoryGridMarkedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmFactoryGridMarkedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmFactoryGridMarked represents a GridMarked event raised by the FarmFactory contract.
type FarmFactoryGridMarked struct {
	FarmId      common.Address
	GridCode    string
	OperationId uint32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterGridMarked is a free log retrieval operation binding the contract event 0x15e824fffe155561f129f13afd27efd7051d81764028f28b090eb3ac94aae688.
//
// Solidity: event GridMarked(address indexed farmId, string gridCode, uint32 operationId)
func (_FarmFactory *FarmFactoryFilterer) FilterGridMarked(opts *bind.FilterOpts, farmId []common.Address) (*FarmFactoryGridMarkedIterator, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}

	logs, sub, err := _FarmFactory.contract.FilterLogs(opts, "GridMarked", farmIdRule)
	if err != nil {
		return nil, err
	}
	return &FarmFactoryGridMarkedIterator{contract: _FarmFactory.contract, event: "GridMarked", logs: logs, sub: sub}, nil
}

// WatchGridMarked is a free log subscription operation binding the contract event 0x15e824fffe155561f129f13afd27efd7051d81764028f28b090eb3ac94aae688.
//
// Solidity: event GridMarked(address indexed farmId, string gridCode, uint32 operationId)
func (_FarmFactory *FarmFactoryFilterer) WatchGridMarked(opts *bind.WatchOpts, sink chan<- *FarmFactoryGridMarked, farmId []common.Address) (event.Subscription, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}

	logs, sub, err := _FarmFactory.contract.WatchLogs(opts, "GridMarked", farmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmFactoryGridMarked)
				if err := _FarmFactory.contract.UnpackLog(event, "GridMarked", log); err != nil {
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

// ParseGridMarked is a log parse operation binding the contract event 0x15e824fffe155561f129f13afd27efd7051d81764028f28b090eb3ac94aae688.
//
// Solidity: event GridMarked(address indexed farmId, string gridCode, uint32 operationId)
func (_FarmFactory *FarmFactoryFilterer) ParseGridMarked(log types.Log) (*FarmFactoryGridMarked, error) {
	event := new(FarmFactoryGridMarked)
	if err := _FarmFactory.contract.UnpackLog(event, "GridMarked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmFactoryRiskDetectedIterator is returned from FilterRiskDetected and is used to iterate over the raw logs and unpacked data for RiskDetected events raised by the FarmFactory contract.
type FarmFactoryRiskDetectedIterator struct {
	Event *FarmFactoryRiskDetected // Event containing the contract specifics and raw log

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
func (it *FarmFactoryRiskDetectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmFactoryRiskDetected)
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
		it.Event = new(FarmFactoryRiskDetected)
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
func (it *FarmFactoryRiskDetectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmFactoryRiskDetectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmFactoryRiskDetected represents a RiskDetected event raised by the FarmFactory contract.
type FarmFactoryRiskDetected struct {
	FarmId   common.Address
	GridCode string
	RecordId uint32
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRiskDetected is a free log retrieval operation binding the contract event 0x88d35c0c5c69c3ab2b4856a7657141cebc6e8a71b0d5c13cda445e25cb5afff1.
//
// Solidity: event RiskDetected(address indexed farmId, string gridCode, uint32 recordId)
func (_FarmFactory *FarmFactoryFilterer) FilterRiskDetected(opts *bind.FilterOpts, farmId []common.Address) (*FarmFactoryRiskDetectedIterator, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}

	logs, sub, err := _FarmFactory.contract.FilterLogs(opts, "RiskDetected", farmIdRule)
	if err != nil {
		return nil, err
	}
	return &FarmFactoryRiskDetectedIterator{contract: _FarmFactory.contract, event: "RiskDetected", logs: logs, sub: sub}, nil
}

// WatchRiskDetected is a free log subscription operation binding the contract event 0x88d35c0c5c69c3ab2b4856a7657141cebc6e8a71b0d5c13cda445e25cb5afff1.
//
// Solidity: event RiskDetected(address indexed farmId, string gridCode, uint32 recordId)
func (_FarmFactory *FarmFactoryFilterer) WatchRiskDetected(opts *bind.WatchOpts, sink chan<- *FarmFactoryRiskDetected, farmId []common.Address) (event.Subscription, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}

	logs, sub, err := _FarmFactory.contract.WatchLogs(opts, "RiskDetected", farmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmFactoryRiskDetected)
				if err := _FarmFactory.contract.UnpackLog(event, "RiskDetected", log); err != nil {
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

// ParseRiskDetected is a log parse operation binding the contract event 0x88d35c0c5c69c3ab2b4856a7657141cebc6e8a71b0d5c13cda445e25cb5afff1.
//
// Solidity: event RiskDetected(address indexed farmId, string gridCode, uint32 recordId)
func (_FarmFactory *FarmFactoryFilterer) ParseRiskDetected(log types.Log) (*FarmFactoryRiskDetected, error) {
	event := new(FarmFactoryRiskDetected)
	if err := _FarmFactory.contract.UnpackLog(event, "RiskDetected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmFactoryTreesRemovedIterator is returned from FilterTreesRemoved and is used to iterate over the raw logs and unpacked data for TreesRemoved events raised by the FarmFactory contract.
type FarmFactoryTreesRemovedIterator struct {
	Event *FarmFactoryTreesRemoved // Event containing the contract specifics and raw log

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
func (it *FarmFactoryTreesRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmFactoryTreesRemoved)
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
		it.Event = new(FarmFactoryTreesRemoved)
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
func (it *FarmFactoryTreesRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmFactoryTreesRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmFactoryTreesRemoved represents a TreesRemoved event raised by the FarmFactory contract.
type FarmFactoryTreesRemoved struct {
	FarmId    common.Address
	GridCode  string
	RemovalId uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTreesRemoved is a free log retrieval operation binding the contract event 0x134658a3a2ee16de2338aaeb4f7ff81beb8354523c314f3880afdba5e9fc20d6.
//
// Solidity: event TreesRemoved(address indexed farmId, string gridCode, uint32 removalId)
func (_FarmFactory *FarmFactoryFilterer) FilterTreesRemoved(opts *bind.FilterOpts, farmId []common.Address) (*FarmFactoryTreesRemovedIterator, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}

	logs, sub, err := _FarmFactory.contract.FilterLogs(opts, "TreesRemoved", farmIdRule)
	if err != nil {
		return nil, err
	}
	return &FarmFactoryTreesRemovedIterator{contract: _FarmFactory.contract, event: "TreesRemoved", logs: logs, sub: sub}, nil
}

// WatchTreesRemoved is a free log subscription operation binding the contract event 0x134658a3a2ee16de2338aaeb4f7ff81beb8354523c314f3880afdba5e9fc20d6.
//
// Solidity: event TreesRemoved(address indexed farmId, string gridCode, uint32 removalId)
func (_FarmFactory *FarmFactoryFilterer) WatchTreesRemoved(opts *bind.WatchOpts, sink chan<- *FarmFactoryTreesRemoved, farmId []common.Address) (event.Subscription, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}

	logs, sub, err := _FarmFactory.contract.WatchLogs(opts, "TreesRemoved", farmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmFactoryTreesRemoved)
				if err := _FarmFactory.contract.UnpackLog(event, "TreesRemoved", log); err != nil {
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

// ParseTreesRemoved is a log parse operation binding the contract event 0x134658a3a2ee16de2338aaeb4f7ff81beb8354523c314f3880afdba5e9fc20d6.
//
// Solidity: event TreesRemoved(address indexed farmId, string gridCode, uint32 removalId)
func (_FarmFactory *FarmFactoryFilterer) ParseTreesRemoved(log types.Log) (*FarmFactoryTreesRemoved, error) {
	event := new(FarmFactoryTreesRemoved)
	if err := _FarmFactory.contract.UnpackLog(event, "TreesRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FarmFactoryTreesReplantedIterator is returned from FilterTreesReplanted and is used to iterate over the raw logs and unpacked data for TreesReplanted events raised by the FarmFactory contract.
type FarmFactoryTreesReplantedIterator struct {
	Event *FarmFactoryTreesReplanted // Event containing the contract specifics and raw log

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
func (it *FarmFactoryTreesReplantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FarmFactoryTreesReplanted)
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
		it.Event = new(FarmFactoryTreesReplanted)
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
func (it *FarmFactoryTreesReplantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FarmFactoryTreesReplantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FarmFactoryTreesReplanted represents a TreesReplanted event raised by the FarmFactory contract.
type FarmFactoryTreesReplanted struct {
	FarmId    common.Address
	GridCode  string
	ReplantId uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTreesReplanted is a free log retrieval operation binding the contract event 0x7549d6e1b1faf802eb6a72cdee9c26f01438d33fd6febf142bea24a48bc8fde3.
//
// Solidity: event TreesReplanted(address indexed farmId, string gridCode, uint32 replantId)
func (_FarmFactory *FarmFactoryFilterer) FilterTreesReplanted(opts *bind.FilterOpts, farmId []common.Address) (*FarmFactoryTreesReplantedIterator, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}

	logs, sub, err := _FarmFactory.contract.FilterLogs(opts, "TreesReplanted", farmIdRule)
	if err != nil {
		return nil, err
	}
	return &FarmFactoryTreesReplantedIterator{contract: _FarmFactory.contract, event: "TreesReplanted", logs: logs, sub: sub}, nil
}

// WatchTreesReplanted is a free log subscription operation binding the contract event 0x7549d6e1b1faf802eb6a72cdee9c26f01438d33fd6febf142bea24a48bc8fde3.
//
// Solidity: event TreesReplanted(address indexed farmId, string gridCode, uint32 replantId)
func (_FarmFactory *FarmFactoryFilterer) WatchTreesReplanted(opts *bind.WatchOpts, sink chan<- *FarmFactoryTreesReplanted, farmId []common.Address) (event.Subscription, error) {

	var farmIdRule []interface{}
	for _, farmIdItem := range farmId {
		farmIdRule = append(farmIdRule, farmIdItem)
	}

	logs, sub, err := _FarmFactory.contract.WatchLogs(opts, "TreesReplanted", farmIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FarmFactoryTreesReplanted)
				if err := _FarmFactory.contract.UnpackLog(event, "TreesReplanted", log); err != nil {
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

// ParseTreesReplanted is a log parse operation binding the contract event 0x7549d6e1b1faf802eb6a72cdee9c26f01438d33fd6febf142bea24a48bc8fde3.
//
// Solidity: event TreesReplanted(address indexed farmId, string gridCode, uint32 replantId)
func (_FarmFactory *FarmFactoryFilterer) ParseTreesReplanted(log types.Log) (*FarmFactoryTreesReplanted, error) {
	event := new(FarmFactoryTreesReplanted)
	if err := _FarmFactory.contract.UnpackLog(event, "TreesReplanted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
