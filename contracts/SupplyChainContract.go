// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// KisanSupplyChainBeneficiary is an auto generated low-level Go binding around an user-defined struct.
type KisanSupplyChainBeneficiary struct {
	BeneficiaryId string
	Name          string
	Uid           string
	Pan           string
	BankAccount   string
}

// KisanSupplyChainBuyer is an auto generated low-level Go binding around an user-defined struct.
type KisanSupplyChainBuyer struct {
	BuyerId     string
	Name        string
	Uid         string
	Pan         string
	Tan         string
	BankAccount string
}

// KisanSupplyChainSales is an auto generated low-level Go binding around an user-defined struct.
type KisanSupplyChainSales struct {
	InvoiceId     string
	Item          string
	Unit          string
	Amount        *big.Int
	AmountPerUnit *big.Int
	SalesDate     *big.Int
	BeneficiaryId string
	BuyerId       string
	Status        string
}

// ContractsABI is the input ABI used to generate the binding from.
const ContractsABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"BeneficiaryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"BuyerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"approvedBy\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"invoiceId\",\"type\":\"string\"}],\"name\":\"InvoiceApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"beneficiaryId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"buyerId\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"invoiceId\",\"type\":\"string\"}],\"name\":\"ItemSold\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"addAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"benificiaryId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"invoiceId\",\"type\":\"string\"}],\"name\":\"approveBeneficiary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"buyerId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"invoiceId\",\"type\":\"string\"}],\"name\":\"approveBuyer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"item\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"unit\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPerUnit\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"beneficiaryId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"buyerId\",\"type\":\"string\"}],\"name\":\"createInvoice\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"invoiceId\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"benificiaryId\",\"type\":\"string\"}],\"name\":\"getBenificiary\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"beneficiaryId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pan\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankAccount\",\"type\":\"string\"}],\"internalType\":\"structKisanSupplyChain.Beneficiary\",\"name\":\"ben\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"buyerId\",\"type\":\"string\"}],\"name\":\"getBuyer\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"buyerId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pan\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tan\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankAccount\",\"type\":\"string\"}],\"internalType\":\"structKisanSupplyChain.Buyer\",\"name\":\"buyer\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"invoiceId\",\"type\":\"string\"}],\"name\":\"getSalesDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"invoiceId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"item\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"unit\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salesDate\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"beneficiaryId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"buyerId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"internalType\":\"structKisanSupplyChain.Sales\",\"name\":\"sales\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mayBeAdmin\",\"type\":\"address\"}],\"name\":\"isAdmin\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"beneficiaryId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pan\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankAccount\",\"type\":\"string\"}],\"name\":\"registerBeneficiary\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"benId\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"buyerId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pan\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tan\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankAccount\",\"type\":\"string\"}],\"name\":\"registerBuyer\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"bId\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractsBin is the compiled bytecode used for deploying new contracts.
var ContractsBin = "0x6080604052600080556000600160006101000a81548160ff02191690831515021790555034801561002f57600080fd5b506001600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff021916908315150217905550612468806100976000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c8063704802751161007157806370480275146101d95780638456cb59146102095780638ffd2411146102135780639d3f696614610243578063dd9ddfca14610273578063f7b188a5146102a3576100b4565b8063103692d4146100b95780631378ded0146100e95780631fa377441461011957806320f64aa61461014957806324d7806c1461017957806330be8365146101a9575b600080fd5b6100d360048036038101906100ce9190611b20565b6102ad565b6040516100e09190612272565b60405180910390f35b61010360048036038101906100fe9190611dd5565b610633565b6040516101109190612204565b60405180910390f35b610133600480360381019061012e9190611b61565b610879565b60405161014091906121e9565b60405180910390f35b610163600480360381019061015e9190611b20565b610995565b60405161017091906122b6565b60405180910390f35b610193600480360381019061018e9190611af7565b610dd9565b6040516101a091906121e9565b60405180910390f35b6101c360048036038101906101be9190611b20565b610e50565b6040516101d09190612294565b60405180910390f35b6101f360048036038101906101ee9190611af7565b611278565b60405161020091906121e9565b60405180910390f35b610211611359565b005b61022d60048036038101906102289190611cbc565b6113f3565b60405161023a9190612204565b60405180910390f35b61025d60048036038101906102589190611bcd565b6115c2565b60405161026a9190612204565b60405180910390f35b61028d60048036038101906102889190611b61565b61176d565b60405161029a91906121e9565b60405180910390f35b6102ab611889565b005b6102b5611923565b60011515600160009054906101000a900460ff16151514156102d657600080fd5b6003826040516102e691906121d2565b90815260200160405180910390206040518060a0016040529081600082018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103985780601f1061036d57610100808354040283529160200191610398565b820191906000526020600020905b81548152906001019060200180831161037b57829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561043a5780601f1061040f5761010080835404028352916020019161043a565b820191906000526020600020905b81548152906001019060200180831161041d57829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104dc5780601f106104b1576101008083540402835291602001916104dc565b820191906000526020600020905b8154815290600101906020018083116104bf57829003601f168201915b50505050508152602001600382018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561057e5780601f106105535761010080835404028352916020019161057e565b820191906000526020600020905b81548152906001019060200180831161056157829003601f168201915b50505050508152602001600482018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106205780601f106105f557610100808354040283529160200191610620565b820191906000526020600020905b81548152906001019060200180831161060357829003601f168201915b5050505050815250509050809050919050565b606060011515600160009054906101000a900460ff161515141561065657600080fd5b879050610661611952565b6040518061012001604052808381526020018981526020018881526020018781526020018681526020014281526020018581526020018481526020016040518060400160405280601281526020017f436f6e63696c6174696f6e50656e64696e6700000000000000000000000000008152508152509050806005836040516106e991906121d2565b9081526020016040518091039020600082015181600001908051906020019061071392919061199e565b50602082015181600101908051906020019061073092919061199e565b50604082015181600201908051906020019061074d92919061199e565b50606082015181600301556080820151816004015560a0820151816005015560c082015181600601908051906020019061078892919061199e565b5060e08201518160070190805190602001906107a592919061199e565b506101008201518160080190805190602001906107c392919061199e565b509050506040518060400160405280601281526020017f436f6e63696c6174696f6e50656e64696e67000000000000000000000000000081525060088360405161080d91906121d2565b9081526020016040518091039020908051906020019061082e92919061199e565b507fc18e13caa82804b8ecce2d9483f0afc4bb7a3cb524beb949ad7f67951c61317d84848460405161086293929190612226565b60405180910390a181915050979650505050505050565b600060011515600160009054906101000a900460ff161515141561089c57600080fd5b6040518060400160405280600b81526020017f42656e417070726f7665640000000000000000000000000000000000000000008152506006846040516108e291906121d2565b9081526020016040518091039020836040516108fe91906121d2565b9081526020016040518091039020908051906020019061091f92919061199e565b506040518060400160405280600b81526020017f42656e417070726f76656400000000000000000000000000000000000000000081525060058360405161096691906121d2565b9081526020016040518091039020600801908051906020019061098a92919061199e565b506001905092915050565b61099d611952565b60011515600160009054906101000a900460ff16151514156109be57600080fd5b6005826040516109ce91906121d2565b908152602001604051809103902060405180610120016040529081600082018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a815780601f10610a5657610100808354040283529160200191610a81565b820191906000526020600020905b815481529060010190602001808311610a6457829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b235780601f10610af857610100808354040283529160200191610b23565b820191906000526020600020905b815481529060010190602001808311610b0657829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610bc55780601f10610b9a57610100808354040283529160200191610bc5565b820191906000526020600020905b815481529060010190602001808311610ba857829003601f168201915b50505050508152602001600382015481526020016004820154815260200160058201548152602001600682018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c855780601f10610c5a57610100808354040283529160200191610c85565b820191906000526020600020905b815481529060010190602001808311610c6857829003601f168201915b50505050508152602001600782018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d275780601f10610cfc57610100808354040283529160200191610d27565b820191906000526020600020905b815481529060010190602001808311610d0a57829003601f168201915b50505050508152602001600882018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610dc95780601f10610d9e57610100808354040283529160200191610dc9565b820191906000526020600020905b815481529060010190602001808311610dac57829003601f168201915b5050505050815250509050919050565b600060011515600160009054906101000a900460ff1615151415610dfc57600080fd5b600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050919050565b610e58611a1e565b60011515600160009054906101000a900460ff1615151415610e7957600080fd5b600482604051610e8991906121d2565b90815260200160405180910390206040518060c0016040529081600082018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f3b5780601f10610f1057610100808354040283529160200191610f3b565b820191906000526020600020905b815481529060010190602001808311610f1e57829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610fdd5780601f10610fb257610100808354040283529160200191610fdd565b820191906000526020600020905b815481529060010190602001808311610fc057829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561107f5780601f106110545761010080835404028352916020019161107f565b820191906000526020600020905b81548152906001019060200180831161106257829003601f168201915b50505050508152602001600382018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111215780601f106110f657610100808354040283529160200191611121565b820191906000526020600020905b81548152906001019060200180831161110457829003601f168201915b50505050508152602001600482018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111c35780601f10611198576101008083540402835291602001916111c3565b820191906000526020600020905b8154815290600101906020018083116111a657829003601f168201915b50505050508152602001600582018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156112655780601f1061123a57610100808354040283529160200191611265565b820191906000526020600020905b81548152906001019060200180831161124857829003601f168201915b5050505050815250509050809050919050565b600060011515600160009054906101000a900460ff161515141561129b57600080fd5b60011515600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515146112f857600080fd5b6001600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555060019050919050565b60011515600160009054906101000a900460ff161515141561137a57600080fd5b60011515600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff161515146113d757600080fd5b60018060006101000a81548160ff021916908315150217905550565b606060011515600160009054906101000a900460ff161515141561141657600080fd5b60011515600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151461147357600080fd5b86905061147e611a1e565b6040518060c00160405280898152602001888152602001878152602001868152602001858152602001848152509050806004896040516114be91906121d2565b908152602001604051809103902060008201518160000190805190602001906114e892919061199e565b50602082015181600101908051906020019061150592919061199e565b50604082015181600201908051906020019061152292919061199e565b50606082015181600301908051906020019061153f92919061199e565b50608082015181600401908051906020019061155c92919061199e565b5060a082015181600501908051906020019061157992919061199e565b509050507fa02f3df0d0f1c777036d971f90f2b342e5f8879472d31d1e8affe7af7e70e5e6886040516115ac9190612204565b60405180910390a1879150509695505050505050565b606060011515600160009054906101000a900460ff16151514156115e557600080fd5b60011515600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151461164257600080fd5b85905061164d611923565b6040518060a001604052808881526020018781526020018681526020018581526020018481525090508060038860405161168791906121d2565b908152602001604051809103902060008201518160000190805190602001906116b192919061199e565b5060208201518160010190805190602001906116ce92919061199e565b5060408201518160020190805190602001906116eb92919061199e565b50606082015181600301908051906020019061170892919061199e565b50608082015181600401908051906020019061172592919061199e565b509050507fe7b80b875912d24f5138e1d9ac8b502f7f307f9cf5ad51ae52df7d85dbf134f9876040516117589190612204565b60405180910390a18191505095945050505050565b600060011515600160009054906101000a900460ff161515141561179057600080fd5b6040518060400160405280600b81526020017f427579417070726f7665640000000000000000000000000000000000000000008152506007846040516117d691906121d2565b9081526020016040518091039020836040516117f291906121d2565b9081526020016040518091039020908051906020019061181392919061199e565b506040518060400160405280600b81526020017f427579417070726f76656400000000000000000000000000000000000000000081525060058360405161185a91906121d2565b9081526020016040518091039020600801908051906020019061187e92919061199e565b506001905092915050565b60011515600160009054906101000a900460ff161515146118a957600080fd5b60011515600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615151461190657600080fd5b6000600160006101000a81548160ff021916908315150217905550565b6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6040518061012001604052806060815260200160608152602001606081526020016000815260200160008152602001600081526020016060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106119df57805160ff1916838001178555611a0d565b82800160010185558215611a0d579182015b82811115611a0c5782518255916020019190600101906119f1565b5b509050611a1a9190611a54565b5090565b6040518060c001604052806060815260200160608152602001606081526020016060815260200160608152602001606081525090565b611a7691905b80821115611a72576000816000905550600101611a5a565b5090565b90565b600081359050611a8881612404565b92915050565b600082601f830112611a9f57600080fd5b8135611ab2611aad82612305565b6122d8565b91508082526020830160208301858383011115611ace57600080fd5b611ad98382846123b1565b50505092915050565b600081359050611af18161241b565b92915050565b600060208284031215611b0957600080fd5b6000611b1784828501611a79565b91505092915050565b600060208284031215611b3257600080fd5b600082013567ffffffffffffffff811115611b4c57600080fd5b611b5884828501611a8e565b91505092915050565b60008060408385031215611b7457600080fd5b600083013567ffffffffffffffff811115611b8e57600080fd5b611b9a85828601611a8e565b925050602083013567ffffffffffffffff811115611bb757600080fd5b611bc385828601611a8e565b9150509250929050565b600080600080600060a08688031215611be557600080fd5b600086013567ffffffffffffffff811115611bff57600080fd5b611c0b88828901611a8e565b955050602086013567ffffffffffffffff811115611c2857600080fd5b611c3488828901611a8e565b945050604086013567ffffffffffffffff811115611c5157600080fd5b611c5d88828901611a8e565b935050606086013567ffffffffffffffff811115611c7a57600080fd5b611c8688828901611a8e565b925050608086013567ffffffffffffffff811115611ca357600080fd5b611caf88828901611a8e565b9150509295509295909350565b60008060008060008060c08789031215611cd557600080fd5b600087013567ffffffffffffffff811115611cef57600080fd5b611cfb89828a01611a8e565b965050602087013567ffffffffffffffff811115611d1857600080fd5b611d2489828a01611a8e565b955050604087013567ffffffffffffffff811115611d4157600080fd5b611d4d89828a01611a8e565b945050606087013567ffffffffffffffff811115611d6a57600080fd5b611d7689828a01611a8e565b935050608087013567ffffffffffffffff811115611d9357600080fd5b611d9f89828a01611a8e565b92505060a087013567ffffffffffffffff811115611dbc57600080fd5b611dc889828a01611a8e565b9150509295509295509295565b600080600080600080600060e0888a031215611df057600080fd5b600088013567ffffffffffffffff811115611e0a57600080fd5b611e168a828b01611a8e565b975050602088013567ffffffffffffffff811115611e3357600080fd5b611e3f8a828b01611a8e565b965050604088013567ffffffffffffffff811115611e5c57600080fd5b611e688a828b01611a8e565b9550506060611e798a828b01611ae2565b9450506080611e8a8a828b01611ae2565b93505060a088013567ffffffffffffffff811115611ea757600080fd5b611eb38a828b01611a8e565b92505060c088013567ffffffffffffffff811115611ed057600080fd5b611edc8a828b01611a8e565b91505092959891949750929550565b611ef48161237b565b82525050565b6000611f0582612331565b611f0f818561233c565b9350611f1f8185602086016123c0565b611f28816123f3565b840191505092915050565b6000611f3e82612331565b611f48818561234d565b9350611f588185602086016123c0565b611f61816123f3565b840191505092915050565b6000611f7782612331565b611f81818561235e565b9350611f918185602086016123c0565b80840191505092915050565b600060a0830160008301518482036000860152611fba8282611efa565b91505060208301518482036020860152611fd48282611efa565b91505060408301518482036040860152611fee8282611efa565b915050606083015184820360608601526120088282611efa565b915050608083015184820360808601526120228282611efa565b9150508091505092915050565b600060c083016000830151848203600086015261204c8282611efa565b915050602083015184820360208601526120668282611efa565b915050604083015184820360408601526120808282611efa565b9150506060830151848203606086015261209a8282611efa565b915050608083015184820360808601526120b48282611efa565b91505060a083015184820360a08601526120ce8282611efa565b9150508091505092915050565b60006101208301600083015184820360008601526120f98282611efa565b915050602083015184820360208601526121138282611efa565b9150506040830151848203604086015261212d8282611efa565b915050606083015161214260608601826121c3565b50608083015161215560808601826121c3565b5060a083015161216860a08601826121c3565b5060c083015184820360c08601526121808282611efa565b91505060e083015184820360e086015261219a8282611efa565b9150506101008301518482036101008601526121b68282611efa565b9150508091505092915050565b6121cc816123a7565b82525050565b60006121de8284611f6c565b915081905092915050565b60006020820190506121fe6000830184611eeb565b92915050565b6000602082019050818103600083015261221e8184611f33565b905092915050565b600060608201905081810360008301526122408186611f33565b905081810360208301526122548185611f33565b905081810360408301526122688184611f33565b9050949350505050565b6000602082019050818103600083015261228c8184611f9d565b905092915050565b600060208201905081810360008301526122ae818461202f565b905092915050565b600060208201905081810360008301526122d081846120db565b905092915050565b6000604051905081810181811067ffffffffffffffff821117156122fb57600080fd5b8060405250919050565b600067ffffffffffffffff82111561231c57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600061237482612387565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156123de5780820151818401526020810190506123c3565b838111156123ed576000848401525b50505050565b6000601f19601f8301169050919050565b61240d81612369565b811461241857600080fd5b50565b612424816123a7565b811461242f57600080fd5b5056fea26469706673582212204295c34ac22a5b56c3573b2936184d71a4ada7998cecc28553647f9532022fbd64736f6c63430006080033"

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// GetBenificiary is a free data retrieval call binding the contract method 0x103692d4.
//
// Solidity: function getBenificiary(string benificiaryId) view returns((string,string,string,string,string) ben)
func (_Contracts *ContractsCaller) GetBenificiary(opts *bind.CallOpts, benificiaryId string) (KisanSupplyChainBeneficiary, error) {
	var (
		ret0 = new(KisanSupplyChainBeneficiary)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "getBenificiary", benificiaryId)
	return *ret0, err
}

// GetBenificiary is a free data retrieval call binding the contract method 0x103692d4.
//
// Solidity: function getBenificiary(string benificiaryId) view returns((string,string,string,string,string) ben)
func (_Contracts *ContractsSession) GetBenificiary(benificiaryId string) (KisanSupplyChainBeneficiary, error) {
	return _Contracts.Contract.GetBenificiary(&_Contracts.CallOpts, benificiaryId)
}

// GetBenificiary is a free data retrieval call binding the contract method 0x103692d4.
//
// Solidity: function getBenificiary(string benificiaryId) view returns((string,string,string,string,string) ben)
func (_Contracts *ContractsCallerSession) GetBenificiary(benificiaryId string) (KisanSupplyChainBeneficiary, error) {
	return _Contracts.Contract.GetBenificiary(&_Contracts.CallOpts, benificiaryId)
}

// GetBuyer is a free data retrieval call binding the contract method 0x30be8365.
//
// Solidity: function getBuyer(string buyerId) view returns((string,string,string,string,string,string) buyer)
func (_Contracts *ContractsCaller) GetBuyer(opts *bind.CallOpts, buyerId string) (KisanSupplyChainBuyer, error) {
	var (
		ret0 = new(KisanSupplyChainBuyer)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "getBuyer", buyerId)
	return *ret0, err
}

// GetBuyer is a free data retrieval call binding the contract method 0x30be8365.
//
// Solidity: function getBuyer(string buyerId) view returns((string,string,string,string,string,string) buyer)
func (_Contracts *ContractsSession) GetBuyer(buyerId string) (KisanSupplyChainBuyer, error) {
	return _Contracts.Contract.GetBuyer(&_Contracts.CallOpts, buyerId)
}

// GetBuyer is a free data retrieval call binding the contract method 0x30be8365.
//
// Solidity: function getBuyer(string buyerId) view returns((string,string,string,string,string,string) buyer)
func (_Contracts *ContractsCallerSession) GetBuyer(buyerId string) (KisanSupplyChainBuyer, error) {
	return _Contracts.Contract.GetBuyer(&_Contracts.CallOpts, buyerId)
}

// GetSalesDetails is a free data retrieval call binding the contract method 0x20f64aa6.
//
// Solidity: function getSalesDetails(string invoiceId) view returns((string,string,string,uint256,uint256,uint256,string,string,string) sales)
func (_Contracts *ContractsCaller) GetSalesDetails(opts *bind.CallOpts, invoiceId string) (KisanSupplyChainSales, error) {
	var (
		ret0 = new(KisanSupplyChainSales)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "getSalesDetails", invoiceId)
	return *ret0, err
}

// GetSalesDetails is a free data retrieval call binding the contract method 0x20f64aa6.
//
// Solidity: function getSalesDetails(string invoiceId) view returns((string,string,string,uint256,uint256,uint256,string,string,string) sales)
func (_Contracts *ContractsSession) GetSalesDetails(invoiceId string) (KisanSupplyChainSales, error) {
	return _Contracts.Contract.GetSalesDetails(&_Contracts.CallOpts, invoiceId)
}

// GetSalesDetails is a free data retrieval call binding the contract method 0x20f64aa6.
//
// Solidity: function getSalesDetails(string invoiceId) view returns((string,string,string,uint256,uint256,uint256,string,string,string) sales)
func (_Contracts *ContractsCallerSession) GetSalesDetails(invoiceId string) (KisanSupplyChainSales, error) {
	return _Contracts.Contract.GetSalesDetails(&_Contracts.CallOpts, invoiceId)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address mayBeAdmin) view returns(bool result)
func (_Contracts *ContractsCaller) IsAdmin(opts *bind.CallOpts, mayBeAdmin common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "isAdmin", mayBeAdmin)
	return *ret0, err
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address mayBeAdmin) view returns(bool result)
func (_Contracts *ContractsSession) IsAdmin(mayBeAdmin common.Address) (bool, error) {
	return _Contracts.Contract.IsAdmin(&_Contracts.CallOpts, mayBeAdmin)
}

// IsAdmin is a free data retrieval call binding the contract method 0x24d7806c.
//
// Solidity: function isAdmin(address mayBeAdmin) view returns(bool result)
func (_Contracts *ContractsCallerSession) IsAdmin(mayBeAdmin common.Address) (bool, error) {
	return _Contracts.Contract.IsAdmin(&_Contracts.CallOpts, mayBeAdmin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns(bool success)
func (_Contracts *ContractsTransactor) AddAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "addAdmin", admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns(bool success)
func (_Contracts *ContractsSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddAdmin(&_Contracts.TransactOpts, admin)
}

// AddAdmin is a paid mutator transaction binding the contract method 0x70480275.
//
// Solidity: function addAdmin(address admin) returns(bool success)
func (_Contracts *ContractsTransactorSession) AddAdmin(admin common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.AddAdmin(&_Contracts.TransactOpts, admin)
}

// ApproveBeneficiary is a paid mutator transaction binding the contract method 0x1fa37744.
//
// Solidity: function approveBeneficiary(string benificiaryId, string invoiceId) returns(bool success)
func (_Contracts *ContractsTransactor) ApproveBeneficiary(opts *bind.TransactOpts, benificiaryId string, invoiceId string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approveBeneficiary", benificiaryId, invoiceId)
}

// ApproveBeneficiary is a paid mutator transaction binding the contract method 0x1fa37744.
//
// Solidity: function approveBeneficiary(string benificiaryId, string invoiceId) returns(bool success)
func (_Contracts *ContractsSession) ApproveBeneficiary(benificiaryId string, invoiceId string) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveBeneficiary(&_Contracts.TransactOpts, benificiaryId, invoiceId)
}

// ApproveBeneficiary is a paid mutator transaction binding the contract method 0x1fa37744.
//
// Solidity: function approveBeneficiary(string benificiaryId, string invoiceId) returns(bool success)
func (_Contracts *ContractsTransactorSession) ApproveBeneficiary(benificiaryId string, invoiceId string) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveBeneficiary(&_Contracts.TransactOpts, benificiaryId, invoiceId)
}

// ApproveBuyer is a paid mutator transaction binding the contract method 0xdd9ddfca.
//
// Solidity: function approveBuyer(string buyerId, string invoiceId) returns(bool success)
func (_Contracts *ContractsTransactor) ApproveBuyer(opts *bind.TransactOpts, buyerId string, invoiceId string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approveBuyer", buyerId, invoiceId)
}

// ApproveBuyer is a paid mutator transaction binding the contract method 0xdd9ddfca.
//
// Solidity: function approveBuyer(string buyerId, string invoiceId) returns(bool success)
func (_Contracts *ContractsSession) ApproveBuyer(buyerId string, invoiceId string) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveBuyer(&_Contracts.TransactOpts, buyerId, invoiceId)
}

// ApproveBuyer is a paid mutator transaction binding the contract method 0xdd9ddfca.
//
// Solidity: function approveBuyer(string buyerId, string invoiceId) returns(bool success)
func (_Contracts *ContractsTransactorSession) ApproveBuyer(buyerId string, invoiceId string) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveBuyer(&_Contracts.TransactOpts, buyerId, invoiceId)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0x1378ded0.
//
// Solidity: function createInvoice(string id, string item, string unit, uint256 amount, uint256 amountPerUnit, string beneficiaryId, string buyerId) returns(string invoiceId)
func (_Contracts *ContractsTransactor) CreateInvoice(opts *bind.TransactOpts, id string, item string, unit string, amount *big.Int, amountPerUnit *big.Int, beneficiaryId string, buyerId string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createInvoice", id, item, unit, amount, amountPerUnit, beneficiaryId, buyerId)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0x1378ded0.
//
// Solidity: function createInvoice(string id, string item, string unit, uint256 amount, uint256 amountPerUnit, string beneficiaryId, string buyerId) returns(string invoiceId)
func (_Contracts *ContractsSession) CreateInvoice(id string, item string, unit string, amount *big.Int, amountPerUnit *big.Int, beneficiaryId string, buyerId string) (*types.Transaction, error) {
	return _Contracts.Contract.CreateInvoice(&_Contracts.TransactOpts, id, item, unit, amount, amountPerUnit, beneficiaryId, buyerId)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0x1378ded0.
//
// Solidity: function createInvoice(string id, string item, string unit, uint256 amount, uint256 amountPerUnit, string beneficiaryId, string buyerId) returns(string invoiceId)
func (_Contracts *ContractsTransactorSession) CreateInvoice(id string, item string, unit string, amount *big.Int, amountPerUnit *big.Int, beneficiaryId string, buyerId string) (*types.Transaction, error) {
	return _Contracts.Contract.CreateInvoice(&_Contracts.TransactOpts, id, item, unit, amount, amountPerUnit, beneficiaryId, buyerId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Contracts *ContractsTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Contracts *ContractsSession) Pause() (*types.Transaction, error) {
	return _Contracts.Contract.Pause(&_Contracts.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Contracts *ContractsTransactorSession) Pause() (*types.Transaction, error) {
	return _Contracts.Contract.Pause(&_Contracts.TransactOpts)
}

// RegisterBeneficiary is a paid mutator transaction binding the contract method 0x9d3f6966.
//
// Solidity: function registerBeneficiary(string beneficiaryId, string name, string uid, string pan, string bankAccount) returns(string benId)
func (_Contracts *ContractsTransactor) RegisterBeneficiary(opts *bind.TransactOpts, beneficiaryId string, name string, uid string, pan string, bankAccount string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerBeneficiary", beneficiaryId, name, uid, pan, bankAccount)
}

// RegisterBeneficiary is a paid mutator transaction binding the contract method 0x9d3f6966.
//
// Solidity: function registerBeneficiary(string beneficiaryId, string name, string uid, string pan, string bankAccount) returns(string benId)
func (_Contracts *ContractsSession) RegisterBeneficiary(beneficiaryId string, name string, uid string, pan string, bankAccount string) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterBeneficiary(&_Contracts.TransactOpts, beneficiaryId, name, uid, pan, bankAccount)
}

// RegisterBeneficiary is a paid mutator transaction binding the contract method 0x9d3f6966.
//
// Solidity: function registerBeneficiary(string beneficiaryId, string name, string uid, string pan, string bankAccount) returns(string benId)
func (_Contracts *ContractsTransactorSession) RegisterBeneficiary(beneficiaryId string, name string, uid string, pan string, bankAccount string) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterBeneficiary(&_Contracts.TransactOpts, beneficiaryId, name, uid, pan, bankAccount)
}

// RegisterBuyer is a paid mutator transaction binding the contract method 0x8ffd2411.
//
// Solidity: function registerBuyer(string buyerId, string name, string uid, string pan, string tan, string bankAccount) returns(string bId)
func (_Contracts *ContractsTransactor) RegisterBuyer(opts *bind.TransactOpts, buyerId string, name string, uid string, pan string, tan string, bankAccount string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerBuyer", buyerId, name, uid, pan, tan, bankAccount)
}

// RegisterBuyer is a paid mutator transaction binding the contract method 0x8ffd2411.
//
// Solidity: function registerBuyer(string buyerId, string name, string uid, string pan, string tan, string bankAccount) returns(string bId)
func (_Contracts *ContractsSession) RegisterBuyer(buyerId string, name string, uid string, pan string, tan string, bankAccount string) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterBuyer(&_Contracts.TransactOpts, buyerId, name, uid, pan, tan, bankAccount)
}

// RegisterBuyer is a paid mutator transaction binding the contract method 0x8ffd2411.
//
// Solidity: function registerBuyer(string buyerId, string name, string uid, string pan, string tan, string bankAccount) returns(string bId)
func (_Contracts *ContractsTransactorSession) RegisterBuyer(buyerId string, name string, uid string, pan string, tan string, bankAccount string) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterBuyer(&_Contracts.TransactOpts, buyerId, name, uid, pan, tan, bankAccount)
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_Contracts *ContractsTransactor) UnPause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "unPause")
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_Contracts *ContractsSession) UnPause() (*types.Transaction, error) {
	return _Contracts.Contract.UnPause(&_Contracts.TransactOpts)
}

// UnPause is a paid mutator transaction binding the contract method 0xf7b188a5.
//
// Solidity: function unPause() returns()
func (_Contracts *ContractsTransactorSession) UnPause() (*types.Transaction, error) {
	return _Contracts.Contract.UnPause(&_Contracts.TransactOpts)
}

// ContractsBeneficiaryAddedIterator is returned from FilterBeneficiaryAdded and is used to iterate over the raw logs and unpacked data for BeneficiaryAdded events raised by the Contracts contract.
type ContractsBeneficiaryAddedIterator struct {
	Event *ContractsBeneficiaryAdded // Event containing the contract specifics and raw log

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
func (it *ContractsBeneficiaryAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsBeneficiaryAdded)
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
		it.Event = new(ContractsBeneficiaryAdded)
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
func (it *ContractsBeneficiaryAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsBeneficiaryAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsBeneficiaryAdded represents a BeneficiaryAdded event raised by the Contracts contract.
type ContractsBeneficiaryAdded struct {
	Id  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBeneficiaryAdded is a free log retrieval operation binding the contract event 0xe7b80b875912d24f5138e1d9ac8b502f7f307f9cf5ad51ae52df7d85dbf134f9.
//
// Solidity: event BeneficiaryAdded(string id)
func (_Contracts *ContractsFilterer) FilterBeneficiaryAdded(opts *bind.FilterOpts) (*ContractsBeneficiaryAddedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "BeneficiaryAdded")
	if err != nil {
		return nil, err
	}
	return &ContractsBeneficiaryAddedIterator{contract: _Contracts.contract, event: "BeneficiaryAdded", logs: logs, sub: sub}, nil
}

// WatchBeneficiaryAdded is a free log subscription operation binding the contract event 0xe7b80b875912d24f5138e1d9ac8b502f7f307f9cf5ad51ae52df7d85dbf134f9.
//
// Solidity: event BeneficiaryAdded(string id)
func (_Contracts *ContractsFilterer) WatchBeneficiaryAdded(opts *bind.WatchOpts, sink chan<- *ContractsBeneficiaryAdded) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "BeneficiaryAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsBeneficiaryAdded)
				if err := _Contracts.contract.UnpackLog(event, "BeneficiaryAdded", log); err != nil {
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

// ParseBeneficiaryAdded is a log parse operation binding the contract event 0xe7b80b875912d24f5138e1d9ac8b502f7f307f9cf5ad51ae52df7d85dbf134f9.
//
// Solidity: event BeneficiaryAdded(string id)
func (_Contracts *ContractsFilterer) ParseBeneficiaryAdded(log types.Log) (*ContractsBeneficiaryAdded, error) {
	event := new(ContractsBeneficiaryAdded)
	if err := _Contracts.contract.UnpackLog(event, "BeneficiaryAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractsBuyerAddedIterator is returned from FilterBuyerAdded and is used to iterate over the raw logs and unpacked data for BuyerAdded events raised by the Contracts contract.
type ContractsBuyerAddedIterator struct {
	Event *ContractsBuyerAdded // Event containing the contract specifics and raw log

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
func (it *ContractsBuyerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsBuyerAdded)
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
		it.Event = new(ContractsBuyerAdded)
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
func (it *ContractsBuyerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsBuyerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsBuyerAdded represents a BuyerAdded event raised by the Contracts contract.
type ContractsBuyerAdded struct {
	Id  string
	Raw types.Log // Blockchain specific contextual infos
}

// FilterBuyerAdded is a free log retrieval operation binding the contract event 0xa02f3df0d0f1c777036d971f90f2b342e5f8879472d31d1e8affe7af7e70e5e6.
//
// Solidity: event BuyerAdded(string id)
func (_Contracts *ContractsFilterer) FilterBuyerAdded(opts *bind.FilterOpts) (*ContractsBuyerAddedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "BuyerAdded")
	if err != nil {
		return nil, err
	}
	return &ContractsBuyerAddedIterator{contract: _Contracts.contract, event: "BuyerAdded", logs: logs, sub: sub}, nil
}

// WatchBuyerAdded is a free log subscription operation binding the contract event 0xa02f3df0d0f1c777036d971f90f2b342e5f8879472d31d1e8affe7af7e70e5e6.
//
// Solidity: event BuyerAdded(string id)
func (_Contracts *ContractsFilterer) WatchBuyerAdded(opts *bind.WatchOpts, sink chan<- *ContractsBuyerAdded) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "BuyerAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsBuyerAdded)
				if err := _Contracts.contract.UnpackLog(event, "BuyerAdded", log); err != nil {
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

// ParseBuyerAdded is a log parse operation binding the contract event 0xa02f3df0d0f1c777036d971f90f2b342e5f8879472d31d1e8affe7af7e70e5e6.
//
// Solidity: event BuyerAdded(string id)
func (_Contracts *ContractsFilterer) ParseBuyerAdded(log types.Log) (*ContractsBuyerAdded, error) {
	event := new(ContractsBuyerAdded)
	if err := _Contracts.contract.UnpackLog(event, "BuyerAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractsInvoiceApprovedIterator is returned from FilterInvoiceApproved and is used to iterate over the raw logs and unpacked data for InvoiceApproved events raised by the Contracts contract.
type ContractsInvoiceApprovedIterator struct {
	Event *ContractsInvoiceApproved // Event containing the contract specifics and raw log

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
func (it *ContractsInvoiceApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsInvoiceApproved)
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
		it.Event = new(ContractsInvoiceApproved)
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
func (it *ContractsInvoiceApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsInvoiceApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsInvoiceApproved represents a InvoiceApproved event raised by the Contracts contract.
type ContractsInvoiceApproved struct {
	ApprovedBy string
	InvoiceId  string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInvoiceApproved is a free log retrieval operation binding the contract event 0xac9dd2880cd48342ef259f5ee4d387680bfb36fdde2cc9e7d489542853580192.
//
// Solidity: event InvoiceApproved(string approvedBy, string invoiceId)
func (_Contracts *ContractsFilterer) FilterInvoiceApproved(opts *bind.FilterOpts) (*ContractsInvoiceApprovedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "InvoiceApproved")
	if err != nil {
		return nil, err
	}
	return &ContractsInvoiceApprovedIterator{contract: _Contracts.contract, event: "InvoiceApproved", logs: logs, sub: sub}, nil
}

// WatchInvoiceApproved is a free log subscription operation binding the contract event 0xac9dd2880cd48342ef259f5ee4d387680bfb36fdde2cc9e7d489542853580192.
//
// Solidity: event InvoiceApproved(string approvedBy, string invoiceId)
func (_Contracts *ContractsFilterer) WatchInvoiceApproved(opts *bind.WatchOpts, sink chan<- *ContractsInvoiceApproved) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "InvoiceApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsInvoiceApproved)
				if err := _Contracts.contract.UnpackLog(event, "InvoiceApproved", log); err != nil {
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

// ParseInvoiceApproved is a log parse operation binding the contract event 0xac9dd2880cd48342ef259f5ee4d387680bfb36fdde2cc9e7d489542853580192.
//
// Solidity: event InvoiceApproved(string approvedBy, string invoiceId)
func (_Contracts *ContractsFilterer) ParseInvoiceApproved(log types.Log) (*ContractsInvoiceApproved, error) {
	event := new(ContractsInvoiceApproved)
	if err := _Contracts.contract.UnpackLog(event, "InvoiceApproved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ContractsItemSoldIterator is returned from FilterItemSold and is used to iterate over the raw logs and unpacked data for ItemSold events raised by the Contracts contract.
type ContractsItemSoldIterator struct {
	Event *ContractsItemSold // Event containing the contract specifics and raw log

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
func (it *ContractsItemSoldIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsItemSold)
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
		it.Event = new(ContractsItemSold)
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
func (it *ContractsItemSoldIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsItemSoldIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsItemSold represents a ItemSold event raised by the Contracts contract.
type ContractsItemSold struct {
	BeneficiaryId string
	BuyerId       string
	InvoiceId     string
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterItemSold is a free log retrieval operation binding the contract event 0xc18e13caa82804b8ecce2d9483f0afc4bb7a3cb524beb949ad7f67951c61317d.
//
// Solidity: event ItemSold(string beneficiaryId, string buyerId, string invoiceId)
func (_Contracts *ContractsFilterer) FilterItemSold(opts *bind.FilterOpts) (*ContractsItemSoldIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ItemSold")
	if err != nil {
		return nil, err
	}
	return &ContractsItemSoldIterator{contract: _Contracts.contract, event: "ItemSold", logs: logs, sub: sub}, nil
}

// WatchItemSold is a free log subscription operation binding the contract event 0xc18e13caa82804b8ecce2d9483f0afc4bb7a3cb524beb949ad7f67951c61317d.
//
// Solidity: event ItemSold(string beneficiaryId, string buyerId, string invoiceId)
func (_Contracts *ContractsFilterer) WatchItemSold(opts *bind.WatchOpts, sink chan<- *ContractsItemSold) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "ItemSold")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsItemSold)
				if err := _Contracts.contract.UnpackLog(event, "ItemSold", log); err != nil {
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

// ParseItemSold is a log parse operation binding the contract event 0xc18e13caa82804b8ecce2d9483f0afc4bb7a3cb524beb949ad7f67951c61317d.
//
// Solidity: event ItemSold(string beneficiaryId, string buyerId, string invoiceId)
func (_Contracts *ContractsFilterer) ParseItemSold(log types.Log) (*ContractsItemSold, error) {
	event := new(ContractsItemSold)
	if err := _Contracts.contract.UnpackLog(event, "ItemSold", log); err != nil {
		return nil, err
	}
	return event, nil
}
