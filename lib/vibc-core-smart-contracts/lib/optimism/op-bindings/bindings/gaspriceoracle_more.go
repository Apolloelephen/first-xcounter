// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const GasPriceOracleStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var GasPriceOracleStorageLayout = new(solc.StorageLayout)

var GasPriceOracleDeployedBin = "0x608060405234801561001057600080fd5b50600436106100be5760003560e01c806354fd4d5011610076578063de26c4a11161005b578063de26c4a114610157578063f45e65d81461016a578063fe173b971461015157600080fd5b806354fd4d50146101085780636ef25c3a1461015157600080fd5b8063313ce567116100a7578063313ce567146100e657806349948e0e146100ed578063519b4bd31461010057600080fd5b80630c18c162146100c35780632e0f2625146100de575b600080fd5b6100cb610172565b6040519081526020015b60405180910390f35b6100cb600681565b60066100cb565b6100cb6100fb3660046103fd565b6101fc565b6100cb61025d565b6101446040518060400160405280600581526020017f312e312e3000000000000000000000000000000000000000000000000000000081525081565b6040516100d591906104cc565b486100cb565b6100cb6101653660046103fd565b6102be565b6100cb61036d565b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16638b239f736040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101f7919061053f565b905090565b600080610208836102be565b9050600061021461025d565b61021e9083610587565b9050600061022e6006600a6106e6565b9050600061023a61036d565b6102449084610587565b9050600061025283836106f9565b979650505050505050565b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16635cf249696040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101d3573d6000803e3d6000fd5b80516000908190815b81811015610341578481815181106102e1576102e1610734565b01602001517fff00000000000000000000000000000000000000000000000000000000000000166000036103215761031a600484610763565b925061032f565b61032c601084610763565b92505b806103398161077b565b9150506102c7565b50600061034c610172565b6103569084610763565b905061036481610440610763565b95945050505050565b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16639e8c49666040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101d3573d6000803e3d6000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60006020828403121561040f57600080fd5b813567ffffffffffffffff8082111561042757600080fd5b818401915084601f83011261043b57600080fd5b81358181111561044d5761044d6103ce565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908382118183101715610493576104936103ce565b816040528281528760208487010111156104ac57600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b818110156104f9578581018301518582016040015282016104dd565b8181111561050b576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60006020828403121561055157600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156105bf576105bf610558565b500290565b600181815b8085111561061d57817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0482111561060357610603610558565b8085161561061057918102915b93841c93908002906105c9565b509250929050565b600082610634575060016106e0565b81610641575060006106e0565b816001811461065757600281146106615761067d565b60019150506106e0565b60ff84111561067257610672610558565b50506001821b6106e0565b5060208310610133831016604e8410600b84101617156106a0575081810a6106e0565b6106aa83836105c4565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156106dc576106dc610558565b0290505b92915050565b60006106f28383610625565b9392505050565b60008261072f577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000821982111561077657610776610558565b500190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036107ac576107ac610558565b506001019056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(GasPriceOracleStorageLayoutJSON), GasPriceOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["GasPriceOracle"] = GasPriceOracleStorageLayout
	deployedBytecodes["GasPriceOracle"] = GasPriceOracleDeployedBin
	immutableReferences["GasPriceOracle"] = false
}
