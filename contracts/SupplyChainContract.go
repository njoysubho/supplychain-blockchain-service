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
	InvoiceId     *big.Int
	Item          string
	Unit          string
	Amount        *big.Int
	AmountPerUnit *big.Int
	SalesDate     *big.Int
	BeneficiaryId *big.Int
	BuyerId       *big.Int
	Status        string
}

// ContractsABI is the input ABI used to generate the binding from.
const ContractsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"BeneficiaryAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"}],\"name\":\"BuyerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"approvedBy\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"}],\"name\":\"InvoiceApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beneficiaryId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"buyerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"}],\"name\":\"ItemSold\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"benificiaryId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"}],\"name\":\"approveBeneficiary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"buyerId\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"}],\"name\":\"approveBuyer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"item\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"unit\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"beneficiaryId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyerId\",\"type\":\"uint256\"}],\"name\":\"createInvoice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"benificiaryId\",\"type\":\"string\"}],\"name\":\"getBenificiary\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"beneficiaryId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pan\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankAccount\",\"type\":\"string\"}],\"internalType\":\"structKisanSupplyChain.Beneficiary\",\"name\":\"ben\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"buyerId\",\"type\":\"string\"}],\"name\":\"getBuyer\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"buyerId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pan\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tan\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankAccount\",\"type\":\"string\"}],\"internalType\":\"structKisanSupplyChain.Buyer\",\"name\":\"buyer\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"}],\"name\":\"getSalesDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"item\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"unit\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salesDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"beneficiaryId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyerId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"internalType\":\"structKisanSupplyChain.Sales\",\"name\":\"sales\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"beneficiaryId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pan\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankAccount\",\"type\":\"string\"}],\"name\":\"registerBeneficiary\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"benId\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"buyerId\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pan\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tan\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankAccount\",\"type\":\"string\"}],\"name\":\"registerBuyer\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"bId\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractsBin is the compiled bytecode used for deploying new contracts.
var ContractsBin = "0x60806040526000805534801561001457600080fd5b50611d01806100246000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80638ffd24111161005b5780638ffd24111461014d5780639d3f69661461017d578063aaaced71146101ad578063d1526800146101dd57610088565b8063103692d41461008d57806330be8365146100bd578063345e0c25146100ed57806359779fdf1461011d575b600080fd5b6100a760048036038101906100a291906113f0565b61020d565b6040516100b49190611af8565b60405180910390f35b6100d760048036038101906100d291906113f0565b610572565b6040516100e49190611b1a565b60405180910390f35b610107600480360381019061010291906116f2565b610979565b6040516101149190611abb565b60405180910390f35b61013760048036038101906101329190611746565b610a5e565b6040516101449190611b3c565b60405180910390f35b61016760048036038101906101629190611520565b610cae565b6040516101749190611ad6565b60405180910390f35b61019760048036038101906101929190611431565b610dff565b6040516101a49190611ad6565b60405180910390f35b6101c760048036038101906101c291906116f2565b610f2c565b6040516101d49190611abb565b60405180910390f35b6101f760048036038101906101f29190611639565b611011565b6040516102049190611b5e565b60405180910390f35b610215611231565b6001826040516102259190611a67565b90815260200160405180910390206040518060a0016040529081600082018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102d75780601f106102ac576101008083540402835291602001916102d7565b820191906000526020600020905b8154815290600101906020018083116102ba57829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103795780601f1061034e57610100808354040283529160200191610379565b820191906000526020600020905b81548152906001019060200180831161035c57829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561041b5780601f106103f05761010080835404028352916020019161041b565b820191906000526020600020905b8154815290600101906020018083116103fe57829003601f168201915b50505050508152602001600382018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104bd5780601f10610492576101008083540402835291602001916104bd565b820191906000526020600020905b8154815290600101906020018083116104a057829003601f168201915b50505050508152602001600482018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561055f5780601f106105345761010080835404028352916020019161055f565b820191906000526020600020905b81548152906001019060200180831161054257829003601f168201915b5050505050815250509050809050919050565b61057a611260565b60028260405161058a9190611a67565b90815260200160405180910390206040518060c0016040529081600082018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561063c5780601f106106115761010080835404028352916020019161063c565b820191906000526020600020905b81548152906001019060200180831161061f57829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106de5780601f106106b3576101008083540402835291602001916106de565b820191906000526020600020905b8154815290600101906020018083116106c157829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107805780601f1061075557610100808354040283529160200191610780565b820191906000526020600020905b81548152906001019060200180831161076357829003601f168201915b50505050508152602001600382018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108225780601f106107f757610100808354040283529160200191610822565b820191906000526020600020905b81548152906001019060200180831161080557829003601f168201915b50505050508152602001600482018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108c45780601f10610899576101008083540402835291602001916108c4565b820191906000526020600020905b8154815290600101906020018083116108a757829003601f168201915b50505050508152602001600582018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109665780601f1061093b57610100808354040283529160200191610966565b820191906000526020600020905b81548152906001019060200180831161094957829003601f168201915b5050505050815250509050809050919050565b60006040518060400160405280600b81526020017f42656e417070726f7665640000000000000000000000000000000000000000008152506004846040516109c19190611a67565b9081526020016040518091039020600084815260200190815260200160002090805190602001906109f3929190611296565b506040518060400160405280600b81526020017f42656e417070726f766564000000000000000000000000000000000000000000815250600360008481526020019081526020016000206008019080519060200190610a53929190611296565b506001905092915050565b610a66611316565b600360008381526020019081526020016000206040518061012001604052908160008201548152602001600182018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b285780601f10610afd57610100808354040283529160200191610b28565b820191906000526020600020905b815481529060010190602001808311610b0b57829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610bca5780601f10610b9f57610100808354040283529160200191610bca565b820191906000526020600020905b815481529060010190602001808311610bad57829003601f168201915b505050505081526020016003820154815260200160048201548152602001600582015481526020016006820154815260200160078201548152602001600882018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c9e5780601f10610c7357610100808354040283529160200191610c9e565b820191906000526020600020905b815481529060010190602001808311610c8157829003601f168201915b5050505050815250509050919050565b6060869050610cbb611260565b6040518060c0016040528089815260200188815260200187815260200186815260200185815260200184815250905080600289604051610cfb9190611a67565b90815260200160405180910390206000820151816000019080519060200190610d25929190611296565b506020820151816001019080519060200190610d42929190611296565b506040820151816002019080519060200190610d5f929190611296565b506060820151816003019080519060200190610d7c929190611296565b506080820151816004019080519060200190610d99929190611296565b5060a0820151816005019080519060200190610db6929190611296565b509050507fa02f3df0d0f1c777036d971f90f2b342e5f8879472d31d1e8affe7af7e70e5e688604051610de99190611ad6565b60405180910390a1879150509695505050505050565b6060859050610e0c611231565b6040518060a0016040528088815260200187815260200186815260200185815260200184815250905080600188604051610e469190611a67565b90815260200160405180910390206000820151816000019080519060200190610e70929190611296565b506020820151816001019080519060200190610e8d929190611296565b506040820151816002019080519060200190610eaa929190611296565b506060820151816003019080519060200190610ec7929190611296565b506080820151816004019080519060200190610ee4929190611296565b509050507fe7b80b875912d24f5138e1d9ac8b502f7f307f9cf5ad51ae52df7d85dbf134f987604051610f179190611ad6565b60405180910390a18191505095945050505050565b60006040518060400160405280600b81526020017f427579417070726f766564000000000000000000000000000000000000000000815250600584604051610f749190611a67565b908152602001604051809103902060008481526020019081526020016000209080519060200190610fa6929190611296565b506040518060400160405280600b81526020017f427579417070726f766564000000000000000000000000000000000000000000815250600360008481526020019081526020016000206008019080519060200190611006929190611296565b506001905092915050565b600061101b6111ed565b9050611025611316565b6040518061012001604052808381526020018981526020018881526020018781526020018681526020014281526020018581526020018481526020016040518060400160405280601281526020017f436f6e63696c6174696f6e50656e64696e670000000000000000000000000000815250815250905080600360008481526020019081526020016000206000820151816000015560208201518160010190805190602001906110d6929190611296565b5060408201518160020190805190602001906110f3929190611296565b50606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e08201518160070155610100820151816008019080519060200190611143929190611296565b509050506040518060400160405280601281526020017f436f6e63696c6174696f6e50656e64696e6700000000000000000000000000008152506006600084815260200190815260200160002090805190602001906111a3929190611296565b507f205b6613171927df8c76164c9183dd985bb389bab25cd473c7cfbc5a8273dcb98484846040516111d793929190611b79565b60405180910390a1819150509695505050505050565b600044426000808154809291906001019190505560405160200161121393929190611a7e565b6040516020818303038152906040528051906020012060001c905090565b6040518060a0016040528060608152602001606081526020016060815260200160608152602001606081525090565b6040518060c001604052806060815260200160608152602001606081526020016060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106112d757805160ff1916838001178555611305565b82800160010185558215611305579182015b828111156113045782518255916020019190600101906112e9565b5b5090506113129190611362565b5090565b6040518061012001604052806000815260200160608152602001606081526020016000815260200160008152602001600081526020016000815260200160008152602001606081525090565b61138491905b80821115611380576000816000905550600101611368565b5090565b90565b600082601f83011261139857600080fd5b81356113ab6113a682611bdd565b611bb0565b915080825260208301602083018583830111156113c757600080fd5b6113d2838284611c57565b50505092915050565b6000813590506113ea81611cb4565b92915050565b60006020828403121561140257600080fd5b600082013567ffffffffffffffff81111561141c57600080fd5b61142884828501611387565b91505092915050565b600080600080600060a0868803121561144957600080fd5b600086013567ffffffffffffffff81111561146357600080fd5b61146f88828901611387565b955050602086013567ffffffffffffffff81111561148c57600080fd5b61149888828901611387565b945050604086013567ffffffffffffffff8111156114b557600080fd5b6114c188828901611387565b935050606086013567ffffffffffffffff8111156114de57600080fd5b6114ea88828901611387565b925050608086013567ffffffffffffffff81111561150757600080fd5b61151388828901611387565b9150509295509295909350565b60008060008060008060c0878903121561153957600080fd5b600087013567ffffffffffffffff81111561155357600080fd5b61155f89828a01611387565b965050602087013567ffffffffffffffff81111561157c57600080fd5b61158889828a01611387565b955050604087013567ffffffffffffffff8111156115a557600080fd5b6115b189828a01611387565b945050606087013567ffffffffffffffff8111156115ce57600080fd5b6115da89828a01611387565b935050608087013567ffffffffffffffff8111156115f757600080fd5b61160389828a01611387565b92505060a087013567ffffffffffffffff81111561162057600080fd5b61162c89828a01611387565b9150509295509295509295565b60008060008060008060c0878903121561165257600080fd5b600087013567ffffffffffffffff81111561166c57600080fd5b61167889828a01611387565b965050602087013567ffffffffffffffff81111561169557600080fd5b6116a189828a01611387565b95505060406116b289828a016113db565b94505060606116c389828a016113db565b93505060806116d489828a016113db565b92505060a06116e589828a016113db565b9150509295509295509295565b6000806040838503121561170557600080fd5b600083013567ffffffffffffffff81111561171f57600080fd5b61172b85828601611387565b925050602061173c858286016113db565b9150509250929050565b60006020828403121561175857600080fd5b6000611766848285016113db565b91505092915050565b61177881611c41565b82525050565b600061178982611c09565b6117938185611c14565b93506117a3818560208601611c66565b6117ac81611ca3565b840191505092915050565b60006117c282611c09565b6117cc8185611c25565b93506117dc818560208601611c66565b6117e581611ca3565b840191505092915050565b60006117fb82611c09565b6118058185611c36565b9350611815818560208601611c66565b80840191505092915050565b600060a083016000830151848203600086015261183e828261177e565b91505060208301518482036020860152611858828261177e565b91505060408301518482036040860152611872828261177e565b9150506060830151848203606086015261188c828261177e565b915050608083015184820360808601526118a6828261177e565b9150508091505092915050565b600060c08301600083015184820360008601526118d0828261177e565b915050602083015184820360208601526118ea828261177e565b91505060408301518482036040860152611904828261177e565b9150506060830151848203606086015261191e828261177e565b91505060808301518482036080860152611938828261177e565b91505060a083015184820360a0860152611952828261177e565b9150508091505092915050565b6000610120830160008301516119786000860182611a32565b5060208301518482036020860152611990828261177e565b915050604083015184820360408601526119aa828261177e565b91505060608301516119bf6060860182611a32565b5060808301516119d26080860182611a32565b5060a08301516119e560a0860182611a32565b5060c08301516119f860c0860182611a32565b5060e0830151611a0b60e0860182611a32565b50610100830151848203610100860152611a25828261177e565b9150508091505092915050565b611a3b81611c4d565b82525050565b611a4a81611c4d565b82525050565b611a61611a5c82611c4d565b611c99565b82525050565b6000611a7382846117f0565b915081905092915050565b6000611a8a8286611a50565b602082019150611a9a8285611a50565b602082019150611aaa8284611a50565b602082019150819050949350505050565b6000602082019050611ad0600083018461176f565b92915050565b60006020820190508181036000830152611af081846117b7565b905092915050565b60006020820190508181036000830152611b128184611821565b905092915050565b60006020820190508181036000830152611b3481846118b3565b905092915050565b60006020820190508181036000830152611b56818461195f565b905092915050565b6000602082019050611b736000830184611a41565b92915050565b6000606082019050611b8e6000830186611a41565b611b9b6020830185611a41565b611ba86040830184611a41565b949350505050565b6000604051905081810181811067ffffffffffffffff82111715611bd357600080fd5b8060405250919050565b600067ffffffffffffffff821115611bf457600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b60008115159050919050565b6000819050919050565b82818337600083830152505050565b60005b83811015611c84578082015181840152602081019050611c69565b83811115611c93576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b611cbd81611c4d565b8114611cc857600080fd5b5056fea264697066735822122041da1ca9275fe9de4ab7609c796c65f3f370c2a7f8d82d3141440074479d7f8164736f6c63430006080033"

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

// GetSalesDetails is a free data retrieval call binding the contract method 0x59779fdf.
//
// Solidity: function getSalesDetails(uint256 invoiceId) view returns((uint256,string,string,uint256,uint256,uint256,uint256,uint256,string) sales)
func (_Contracts *ContractsCaller) GetSalesDetails(opts *bind.CallOpts, invoiceId *big.Int) (KisanSupplyChainSales, error) {
	var (
		ret0 = new(KisanSupplyChainSales)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "getSalesDetails", invoiceId)
	return *ret0, err
}

// GetSalesDetails is a free data retrieval call binding the contract method 0x59779fdf.
//
// Solidity: function getSalesDetails(uint256 invoiceId) view returns((uint256,string,string,uint256,uint256,uint256,uint256,uint256,string) sales)
func (_Contracts *ContractsSession) GetSalesDetails(invoiceId *big.Int) (KisanSupplyChainSales, error) {
	return _Contracts.Contract.GetSalesDetails(&_Contracts.CallOpts, invoiceId)
}

// GetSalesDetails is a free data retrieval call binding the contract method 0x59779fdf.
//
// Solidity: function getSalesDetails(uint256 invoiceId) view returns((uint256,string,string,uint256,uint256,uint256,uint256,uint256,string) sales)
func (_Contracts *ContractsCallerSession) GetSalesDetails(invoiceId *big.Int) (KisanSupplyChainSales, error) {
	return _Contracts.Contract.GetSalesDetails(&_Contracts.CallOpts, invoiceId)
}

// ApproveBeneficiary is a paid mutator transaction binding the contract method 0x345e0c25.
//
// Solidity: function approveBeneficiary(string benificiaryId, uint256 invoiceId) returns(bool success)
func (_Contracts *ContractsTransactor) ApproveBeneficiary(opts *bind.TransactOpts, benificiaryId string, invoiceId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approveBeneficiary", benificiaryId, invoiceId)
}

// ApproveBeneficiary is a paid mutator transaction binding the contract method 0x345e0c25.
//
// Solidity: function approveBeneficiary(string benificiaryId, uint256 invoiceId) returns(bool success)
func (_Contracts *ContractsSession) ApproveBeneficiary(benificiaryId string, invoiceId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveBeneficiary(&_Contracts.TransactOpts, benificiaryId, invoiceId)
}

// ApproveBeneficiary is a paid mutator transaction binding the contract method 0x345e0c25.
//
// Solidity: function approveBeneficiary(string benificiaryId, uint256 invoiceId) returns(bool success)
func (_Contracts *ContractsTransactorSession) ApproveBeneficiary(benificiaryId string, invoiceId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveBeneficiary(&_Contracts.TransactOpts, benificiaryId, invoiceId)
}

// ApproveBuyer is a paid mutator transaction binding the contract method 0xaaaced71.
//
// Solidity: function approveBuyer(string buyerId, uint256 invoiceId) returns(bool success)
func (_Contracts *ContractsTransactor) ApproveBuyer(opts *bind.TransactOpts, buyerId string, invoiceId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approveBuyer", buyerId, invoiceId)
}

// ApproveBuyer is a paid mutator transaction binding the contract method 0xaaaced71.
//
// Solidity: function approveBuyer(string buyerId, uint256 invoiceId) returns(bool success)
func (_Contracts *ContractsSession) ApproveBuyer(buyerId string, invoiceId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveBuyer(&_Contracts.TransactOpts, buyerId, invoiceId)
}

// ApproveBuyer is a paid mutator transaction binding the contract method 0xaaaced71.
//
// Solidity: function approveBuyer(string buyerId, uint256 invoiceId) returns(bool success)
func (_Contracts *ContractsTransactorSession) ApproveBuyer(buyerId string, invoiceId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveBuyer(&_Contracts.TransactOpts, buyerId, invoiceId)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0xd1526800.
//
// Solidity: function createInvoice(string item, string unit, uint256 amount, uint256 amountPerUnit, uint256 beneficiaryId, uint256 buyerId) returns(uint256 invoiceId)
func (_Contracts *ContractsTransactor) CreateInvoice(opts *bind.TransactOpts, item string, unit string, amount *big.Int, amountPerUnit *big.Int, beneficiaryId *big.Int, buyerId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "createInvoice", item, unit, amount, amountPerUnit, beneficiaryId, buyerId)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0xd1526800.
//
// Solidity: function createInvoice(string item, string unit, uint256 amount, uint256 amountPerUnit, uint256 beneficiaryId, uint256 buyerId) returns(uint256 invoiceId)
func (_Contracts *ContractsSession) CreateInvoice(item string, unit string, amount *big.Int, amountPerUnit *big.Int, beneficiaryId *big.Int, buyerId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CreateInvoice(&_Contracts.TransactOpts, item, unit, amount, amountPerUnit, beneficiaryId, buyerId)
}

// CreateInvoice is a paid mutator transaction binding the contract method 0xd1526800.
//
// Solidity: function createInvoice(string item, string unit, uint256 amount, uint256 amountPerUnit, uint256 beneficiaryId, uint256 buyerId) returns(uint256 invoiceId)
func (_Contracts *ContractsTransactorSession) CreateInvoice(item string, unit string, amount *big.Int, amountPerUnit *big.Int, beneficiaryId *big.Int, buyerId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.CreateInvoice(&_Contracts.TransactOpts, item, unit, amount, amountPerUnit, beneficiaryId, buyerId)
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
	ApprovedBy *big.Int
	InvoiceId  *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInvoiceApproved is a free log retrieval operation binding the contract event 0x427e01b094cee80465b93663a800e401f0eb28213a4fbd6d1cc217f066036654.
//
// Solidity: event InvoiceApproved(uint256 approvedBy, uint256 invoiceId)
func (_Contracts *ContractsFilterer) FilterInvoiceApproved(opts *bind.FilterOpts) (*ContractsInvoiceApprovedIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "InvoiceApproved")
	if err != nil {
		return nil, err
	}
	return &ContractsInvoiceApprovedIterator{contract: _Contracts.contract, event: "InvoiceApproved", logs: logs, sub: sub}, nil
}

// WatchInvoiceApproved is a free log subscription operation binding the contract event 0x427e01b094cee80465b93663a800e401f0eb28213a4fbd6d1cc217f066036654.
//
// Solidity: event InvoiceApproved(uint256 approvedBy, uint256 invoiceId)
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

// ParseInvoiceApproved is a log parse operation binding the contract event 0x427e01b094cee80465b93663a800e401f0eb28213a4fbd6d1cc217f066036654.
//
// Solidity: event InvoiceApproved(uint256 approvedBy, uint256 invoiceId)
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
	BeneficiaryId *big.Int
	BuyerId       *big.Int
	InvoiceId     *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterItemSold is a free log retrieval operation binding the contract event 0x205b6613171927df8c76164c9183dd985bb389bab25cd473c7cfbc5a8273dcb9.
//
// Solidity: event ItemSold(uint256 beneficiaryId, uint256 buyerId, uint256 invoiceId)
func (_Contracts *ContractsFilterer) FilterItemSold(opts *bind.FilterOpts) (*ContractsItemSoldIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "ItemSold")
	if err != nil {
		return nil, err
	}
	return &ContractsItemSoldIterator{contract: _Contracts.contract, event: "ItemSold", logs: logs, sub: sub}, nil
}

// WatchItemSold is a free log subscription operation binding the contract event 0x205b6613171927df8c76164c9183dd985bb389bab25cd473c7cfbc5a8273dcb9.
//
// Solidity: event ItemSold(uint256 beneficiaryId, uint256 buyerId, uint256 invoiceId)
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

// ParseItemSold is a log parse operation binding the contract event 0x205b6613171927df8c76164c9183dd985bb389bab25cd473c7cfbc5a8273dcb9.
//
// Solidity: event ItemSold(uint256 beneficiaryId, uint256 buyerId, uint256 invoiceId)
func (_Contracts *ContractsFilterer) ParseItemSold(log types.Log) (*ContractsItemSold, error) {
	event := new(ContractsItemSold)
	if err := _Contracts.contract.UnpackLog(event, "ItemSold", log); err != nil {
		return nil, err
	}
	return event, nil
}
