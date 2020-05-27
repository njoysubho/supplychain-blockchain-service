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
	BeneficiaryId *big.Int
	Name          string
	Uid           string
	Pan           string
	BankAccount   string
}

// KisanSupplyChainBuyer is an auto generated low-level Go binding around an user-defined struct.
type KisanSupplyChainBuyer struct {
	BuyerId     *big.Int
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
const ContractsABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"approvedBy\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"}],\"name\":\"InvoiceApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"beneficiaryId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"buyerId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"}],\"name\":\"ItemSold\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"benificiaryId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"}],\"name\":\"approveBeneficiary\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"buyerId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"}],\"name\":\"approveBuyer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"item\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"unit\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"beneficiaryId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyerId\",\"type\":\"uint256\"}],\"name\":\"createInvoice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"benificiaryId\",\"type\":\"uint256\"}],\"name\":\"getBenificiary\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"beneficiaryId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pan\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankAccount\",\"type\":\"string\"}],\"internalType\":\"structKisanSupplyChain.Beneficiary\",\"name\":\"ben\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"buyerId\",\"type\":\"uint256\"}],\"name\":\"getBuyer\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"buyerId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pan\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tan\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankAccount\",\"type\":\"string\"}],\"internalType\":\"structKisanSupplyChain.Buyer\",\"name\":\"buyer\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"}],\"name\":\"getSalesDetails\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"invoiceId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"item\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"unit\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountPerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salesDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"beneficiaryId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"buyerId\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"status\",\"type\":\"string\"}],\"internalType\":\"structKisanSupplyChain.Sales\",\"name\":\"sales\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pan\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankAccount\",\"type\":\"string\"}],\"name\":\"registerBeneficiary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"benificiaryId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"uid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"pan\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tan\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"bankAccount\",\"type\":\"string\"}],\"name\":\"registerBuyer\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"buyerId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ContractsBin is the compiled bytecode used for deploying new contracts.
var ContractsBin = "0x60806040526000805534801561001457600080fd5b5061198b806100246000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c80639069b3ca1161005b5780639069b3ca1461014d578063a88493041461017d578063ccb12f04146101ad578063d1526800146101dd57610088565b80630b13b6e31461008d57806359779fdf146100bd5780635bf608b8146100ed57806388b9c4491461011d575b600080fd5b6100a760048036038101906100a2919061148a565b61020d565b6040516100b49190611783565b60405180910390f35b6100d760048036038101906100d29190611461565b6102e7565b6040516100e491906117e2565b60405180910390f35b61010760048036038101906101029190611461565b610537565b60405161011491906117c0565b60405180910390f35b610137600480360381019061013291906111f6565b61089b565b6040516101449190611804565b60405180910390f35b6101676004803603810190610162919061148a565b610979565b6040516101749190611783565b60405180910390f35b61019760048036038101906101929190611461565b610a53565b6040516101a4919061179e565b60405180910390f35b6101c760048036038101906101c291906112b9565b610d15565b6040516101d49190611804565b60405180910390f35b6101f760048036038101906101f291906113a8565b610e17565b6040516102049190611804565b60405180910390f35b60006040518060400160405280600b81526020017f427579417070726f766564000000000000000000000000000000000000000000815250600560008581526020019081526020016000206000848152602001908152602001600020908051906020019061027c929190611037565b506040518060400160405280600b81526020017f427579417070726f7665640000000000000000000000000000000000000000008152506003600084815260200190815260200160002060080190805190602001906102dc929190611037565b506001905092915050565b6102ef6110b7565b600360008381526020019081526020016000206040518061012001604052908160008201548152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103b15780601f10610386576101008083540402835291602001916103b1565b820191906000526020600020905b81548152906001019060200180831161039457829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104535780601f1061042857610100808354040283529160200191610453565b820191906000526020600020905b81548152906001019060200180831161043657829003601f168201915b505050505081526020016003820154815260200160048201548152602001600582015481526020016006820154815260200160078201548152602001600882018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105275780601f106104fc57610100808354040283529160200191610527565b820191906000526020600020905b81548152906001019060200180831161050a57829003601f168201915b5050505050815250509050919050565b61053f611103565b600260008381526020019081526020016000206040518060c001604052908160008201548152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106005780601f106105d557610100808354040283529160200191610600565b820191906000526020600020905b8154815290600101906020018083116105e357829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106a25780601f10610677576101008083540402835291602001916106a2565b820191906000526020600020905b81548152906001019060200180831161068557829003601f168201915b50505050508152602001600382018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107445780601f1061071957610100808354040283529160200191610744565b820191906000526020600020905b81548152906001019060200180831161072757829003601f168201915b50505050508152602001600482018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107e65780601f106107bb576101008083540402835291602001916107e6565b820191906000526020600020905b8154815290600101906020018083116107c957829003601f168201915b50505050508152602001600582018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108885780601f1061085d57610100808354040283529160200191610888565b820191906000526020600020905b81548152906001019060200180831161086b57829003601f168201915b5050505050815250509050809050919050565b60006108a5610ff3565b90506108af611139565b6040518060a001604052808381526020018781526020018681526020018581526020018481525090508060016000848152602001908152602001600020600082015181600001556020820151816001019080519060200190610912929190611037565b50604082015181600201908051906020019061092f929190611037565b50606082015181600301908051906020019061094c929190611037565b506080820151816004019080519060200190610969929190611037565b5090505081915050949350505050565b60006040518060400160405280600b81526020017f42656e417070726f76656400000000000000000000000000000000000000000081525060046000858152602001908152602001600020600084815260200190815260200160002090805190602001906109e8929190611037565b506040518060400160405280600b81526020017f42656e417070726f766564000000000000000000000000000000000000000000815250600360008481526020019081526020016000206008019080519060200190610a48929190611037565b506001905092915050565b610a5b611139565b600160008381526020019081526020016000206040518060a001604052908160008201548152602001600182018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b1c5780601f10610af157610100808354040283529160200191610b1c565b820191906000526020600020905b815481529060010190602001808311610aff57829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610bbe5780601f10610b9357610100808354040283529160200191610bbe565b820191906000526020600020905b815481529060010190602001808311610ba157829003601f168201915b50505050508152602001600382018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c605780601f10610c3557610100808354040283529160200191610c60565b820191906000526020600020905b815481529060010190602001808311610c4357829003601f168201915b50505050508152602001600482018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d025780601f10610cd757610100808354040283529160200191610d02565b820191906000526020600020905b815481529060010190602001808311610ce557829003601f168201915b5050505050815250509050809050919050565b6000610d1f610ff3565b9050610d29611103565b6040518060c001604052808381526020018881526020018781526020018681526020018581526020018481525090508060026000848152602001908152602001600020600082015181600001556020820151816001019080519060200190610d92929190611037565b506040820151816002019080519060200190610daf929190611037565b506060820151816003019080519060200190610dcc929190611037565b506080820151816004019080519060200190610de9929190611037565b5060a0820151816005019080519060200190610e06929190611037565b509050508191505095945050505050565b6000610e21610ff3565b9050610e2b6110b7565b6040518061012001604052808381526020018981526020018881526020018781526020018681526020014281526020018581526020018481526020016040518060400160405280601281526020017f436f6e63696c6174696f6e50656e64696e67000000000000000000000000000081525081525090508060036000848152602001908152602001600020600082015181600001556020820151816001019080519060200190610edc929190611037565b506040820151816002019080519060200190610ef9929190611037565b50606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e08201518160070155610100820151816008019080519060200190610f49929190611037565b509050506040518060400160405280601281526020017f436f6e63696c6174696f6e50656e64696e670000000000000000000000000000815250600660008481526020019081526020016000209080519060200190610fa9929190611037565b507f205b6613171927df8c76164c9183dd985bb389bab25cd473c7cfbc5a8273dcb9848484604051610fdd9392919061181f565b60405180910390a1819150509695505050505050565b600044426000808154809291906001019190505560405160200161101993929190611746565b6040516020818303038152906040528051906020012060001c905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061107857805160ff19168380011785556110a6565b828001600101855582156110a6579182015b828111156110a557825182559160200191906001019061108a565b5b5090506110b39190611168565b5090565b6040518061012001604052806000815260200160608152602001606081526020016000815260200160008152602001600081526020016000815260200160008152602001606081525090565b6040518060c001604052806000815260200160608152602001606081526020016060815260200160608152602001606081525090565b6040518060a0016040528060008152602001606081526020016060815260200160608152602001606081525090565b61118a91905b8082111561118657600081600090555060010161116e565b5090565b90565b600082601f83011261119e57600080fd5b81356111b16111ac82611883565b611856565b915080825260208301602083018583830111156111cd57600080fd5b6111d88382846118e1565b50505092915050565b6000813590506111f08161193e565b92915050565b6000806000806080858703121561120c57600080fd5b600085013567ffffffffffffffff81111561122657600080fd5b6112328782880161118d565b945050602085013567ffffffffffffffff81111561124f57600080fd5b61125b8782880161118d565b935050604085013567ffffffffffffffff81111561127857600080fd5b6112848782880161118d565b925050606085013567ffffffffffffffff8111156112a157600080fd5b6112ad8782880161118d565b91505092959194509250565b600080600080600060a086880312156112d157600080fd5b600086013567ffffffffffffffff8111156112eb57600080fd5b6112f78882890161118d565b955050602086013567ffffffffffffffff81111561131457600080fd5b6113208882890161118d565b945050604086013567ffffffffffffffff81111561133d57600080fd5b6113498882890161118d565b935050606086013567ffffffffffffffff81111561136657600080fd5b6113728882890161118d565b925050608086013567ffffffffffffffff81111561138f57600080fd5b61139b8882890161118d565b9150509295509295909350565b60008060008060008060c087890312156113c157600080fd5b600087013567ffffffffffffffff8111156113db57600080fd5b6113e789828a0161118d565b965050602087013567ffffffffffffffff81111561140457600080fd5b61141089828a0161118d565b955050604061142189828a016111e1565b945050606061143289828a016111e1565b935050608061144389828a016111e1565b92505060a061145489828a016111e1565b9150509295509295509295565b60006020828403121561147357600080fd5b6000611481848285016111e1565b91505092915050565b6000806040838503121561149d57600080fd5b60006114ab858286016111e1565b92505060206114bc858286016111e1565b9150509250929050565b6114cf816118cb565b82525050565b60006114e0826118af565b6114ea81856118ba565b93506114fa8185602086016118f0565b6115038161192d565b840191505092915050565b600060a0830160008301516115266000860182611711565b506020830151848203602086015261153e82826114d5565b9150506040830151848203604086015261155882826114d5565b9150506060830151848203606086015261157282826114d5565b9150506080830151848203608086015261158c82826114d5565b9150508091505092915050565b600060c0830160008301516115b16000860182611711565b50602083015184820360208601526115c982826114d5565b915050604083015184820360408601526115e382826114d5565b915050606083015184820360608601526115fd82826114d5565b9150506080830151848203608086015261161782826114d5565b91505060a083015184820360a086015261163182826114d5565b9150508091505092915050565b6000610120830160008301516116576000860182611711565b506020830151848203602086015261166f82826114d5565b9150506040830151848203604086015261168982826114d5565b915050606083015161169e6060860182611711565b5060808301516116b16080860182611711565b5060a08301516116c460a0860182611711565b5060c08301516116d760c0860182611711565b5060e08301516116ea60e0860182611711565b5061010083015184820361010086015261170482826114d5565b9150508091505092915050565b61171a816118d7565b82525050565b611729816118d7565b82525050565b61174061173b826118d7565b611923565b82525050565b6000611752828661172f565b602082019150611762828561172f565b602082019150611772828461172f565b602082019150819050949350505050565b600060208201905061179860008301846114c6565b92915050565b600060208201905081810360008301526117b8818461150e565b905092915050565b600060208201905081810360008301526117da8184611599565b905092915050565b600060208201905081810360008301526117fc818461163e565b905092915050565b60006020820190506118196000830184611720565b92915050565b60006060820190506118346000830186611720565b6118416020830185611720565b61184e6040830184611720565b949350505050565b6000604051905081810181811067ffffffffffffffff8211171561187957600080fd5b8060405250919050565b600067ffffffffffffffff82111561189a57600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b60008115159050919050565b6000819050919050565b82818337600083830152505050565b60005b8381101561190e5780820151818401526020810190506118f3565b8381111561191d576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b611947816118d7565b811461195257600080fd5b5056fea2646970667358221220808a23494ed7db6db60572841a952a4757f2cc04ac89b3527331fc628da9081f64736f6c63430006080033"

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

// GetBenificiary is a free data retrieval call binding the contract method 0xa8849304.
//
// Solidity: function getBenificiary(uint256 benificiaryId) view returns((uint256,string,string,string,string) ben)
func (_Contracts *ContractsCaller) GetBenificiary(opts *bind.CallOpts, benificiaryId *big.Int) (KisanSupplyChainBeneficiary, error) {
	var (
		ret0 = new(KisanSupplyChainBeneficiary)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "getBenificiary", benificiaryId)
	return *ret0, err
}

// GetBenificiary is a free data retrieval call binding the contract method 0xa8849304.
//
// Solidity: function getBenificiary(uint256 benificiaryId) view returns((uint256,string,string,string,string) ben)
func (_Contracts *ContractsSession) GetBenificiary(benificiaryId *big.Int) (KisanSupplyChainBeneficiary, error) {
	return _Contracts.Contract.GetBenificiary(&_Contracts.CallOpts, benificiaryId)
}

// GetBenificiary is a free data retrieval call binding the contract method 0xa8849304.
//
// Solidity: function getBenificiary(uint256 benificiaryId) view returns((uint256,string,string,string,string) ben)
func (_Contracts *ContractsCallerSession) GetBenificiary(benificiaryId *big.Int) (KisanSupplyChainBeneficiary, error) {
	return _Contracts.Contract.GetBenificiary(&_Contracts.CallOpts, benificiaryId)
}

// GetBuyer is a free data retrieval call binding the contract method 0x5bf608b8.
//
// Solidity: function getBuyer(uint256 buyerId) view returns((uint256,string,string,string,string,string) buyer)
func (_Contracts *ContractsCaller) GetBuyer(opts *bind.CallOpts, buyerId *big.Int) (KisanSupplyChainBuyer, error) {
	var (
		ret0 = new(KisanSupplyChainBuyer)
	)
	out := ret0
	err := _Contracts.contract.Call(opts, out, "getBuyer", buyerId)
	return *ret0, err
}

// GetBuyer is a free data retrieval call binding the contract method 0x5bf608b8.
//
// Solidity: function getBuyer(uint256 buyerId) view returns((uint256,string,string,string,string,string) buyer)
func (_Contracts *ContractsSession) GetBuyer(buyerId *big.Int) (KisanSupplyChainBuyer, error) {
	return _Contracts.Contract.GetBuyer(&_Contracts.CallOpts, buyerId)
}

// GetBuyer is a free data retrieval call binding the contract method 0x5bf608b8.
//
// Solidity: function getBuyer(uint256 buyerId) view returns((uint256,string,string,string,string,string) buyer)
func (_Contracts *ContractsCallerSession) GetBuyer(buyerId *big.Int) (KisanSupplyChainBuyer, error) {
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

// ApproveBeneficiary is a paid mutator transaction binding the contract method 0x9069b3ca.
//
// Solidity: function approveBeneficiary(uint256 benificiaryId, uint256 invoiceId) returns(bool success)
func (_Contracts *ContractsTransactor) ApproveBeneficiary(opts *bind.TransactOpts, benificiaryId *big.Int, invoiceId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approveBeneficiary", benificiaryId, invoiceId)
}

// ApproveBeneficiary is a paid mutator transaction binding the contract method 0x9069b3ca.
//
// Solidity: function approveBeneficiary(uint256 benificiaryId, uint256 invoiceId) returns(bool success)
func (_Contracts *ContractsSession) ApproveBeneficiary(benificiaryId *big.Int, invoiceId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveBeneficiary(&_Contracts.TransactOpts, benificiaryId, invoiceId)
}

// ApproveBeneficiary is a paid mutator transaction binding the contract method 0x9069b3ca.
//
// Solidity: function approveBeneficiary(uint256 benificiaryId, uint256 invoiceId) returns(bool success)
func (_Contracts *ContractsTransactorSession) ApproveBeneficiary(benificiaryId *big.Int, invoiceId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveBeneficiary(&_Contracts.TransactOpts, benificiaryId, invoiceId)
}

// ApproveBuyer is a paid mutator transaction binding the contract method 0x0b13b6e3.
//
// Solidity: function approveBuyer(uint256 buyerId, uint256 invoiceId) returns(bool success)
func (_Contracts *ContractsTransactor) ApproveBuyer(opts *bind.TransactOpts, buyerId *big.Int, invoiceId *big.Int) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "approveBuyer", buyerId, invoiceId)
}

// ApproveBuyer is a paid mutator transaction binding the contract method 0x0b13b6e3.
//
// Solidity: function approveBuyer(uint256 buyerId, uint256 invoiceId) returns(bool success)
func (_Contracts *ContractsSession) ApproveBuyer(buyerId *big.Int, invoiceId *big.Int) (*types.Transaction, error) {
	return _Contracts.Contract.ApproveBuyer(&_Contracts.TransactOpts, buyerId, invoiceId)
}

// ApproveBuyer is a paid mutator transaction binding the contract method 0x0b13b6e3.
//
// Solidity: function approveBuyer(uint256 buyerId, uint256 invoiceId) returns(bool success)
func (_Contracts *ContractsTransactorSession) ApproveBuyer(buyerId *big.Int, invoiceId *big.Int) (*types.Transaction, error) {
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

// RegisterBeneficiary is a paid mutator transaction binding the contract method 0x88b9c449.
//
// Solidity: function registerBeneficiary(string name, string uid, string pan, string bankAccount) returns(uint256 benificiaryId)
func (_Contracts *ContractsTransactor) RegisterBeneficiary(opts *bind.TransactOpts, name string, uid string, pan string, bankAccount string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerBeneficiary", name, uid, pan, bankAccount)
}

// RegisterBeneficiary is a paid mutator transaction binding the contract method 0x88b9c449.
//
// Solidity: function registerBeneficiary(string name, string uid, string pan, string bankAccount) returns(uint256 benificiaryId)
func (_Contracts *ContractsSession) RegisterBeneficiary(name string, uid string, pan string, bankAccount string) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterBeneficiary(&_Contracts.TransactOpts, name, uid, pan, bankAccount)
}

// RegisterBeneficiary is a paid mutator transaction binding the contract method 0x88b9c449.
//
// Solidity: function registerBeneficiary(string name, string uid, string pan, string bankAccount) returns(uint256 benificiaryId)
func (_Contracts *ContractsTransactorSession) RegisterBeneficiary(name string, uid string, pan string, bankAccount string) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterBeneficiary(&_Contracts.TransactOpts, name, uid, pan, bankAccount)
}

// RegisterBuyer is a paid mutator transaction binding the contract method 0xccb12f04.
//
// Solidity: function registerBuyer(string name, string uid, string pan, string tan, string bankAccount) returns(uint256 buyerId)
func (_Contracts *ContractsTransactor) RegisterBuyer(opts *bind.TransactOpts, name string, uid string, pan string, tan string, bankAccount string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "registerBuyer", name, uid, pan, tan, bankAccount)
}

// RegisterBuyer is a paid mutator transaction binding the contract method 0xccb12f04.
//
// Solidity: function registerBuyer(string name, string uid, string pan, string tan, string bankAccount) returns(uint256 buyerId)
func (_Contracts *ContractsSession) RegisterBuyer(name string, uid string, pan string, tan string, bankAccount string) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterBuyer(&_Contracts.TransactOpts, name, uid, pan, tan, bankAccount)
}

// RegisterBuyer is a paid mutator transaction binding the contract method 0xccb12f04.
//
// Solidity: function registerBuyer(string name, string uid, string pan, string tan, string bankAccount) returns(uint256 buyerId)
func (_Contracts *ContractsTransactorSession) RegisterBuyer(name string, uid string, pan string, tan string, bankAccount string) (*types.Transaction, error) {
	return _Contracts.Contract.RegisterBuyer(&_Contracts.TransactOpts, name, uid, pan, tan, bankAccount)
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
