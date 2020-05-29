pragma solidity >=0.4.22;
pragma experimental ABIEncoderV2;
contract KisanSupplyChain{
    uint seed=0;
    struct Beneficiary {
        string beneficiaryId;
        string name;
        string uid;
        string pan;
        string bankAccount;
    }

    struct Buyer{
        string buyerId;
        string name;
        string uid;
        string pan;
        string tan;
        string bankAccount;
    }

    struct Sales{
        uint invoiceId;
        string item;
        string unit;
        uint256 amount;
        uint256 amountPerUnit;
        uint256 salesDate;
        uint beneficiaryId;
        uint buyerId;
        string status;
    }

    event ItemSold(uint beneficiaryId,uint buyerId,uint invoiceId);
    event InvoiceApproved(uint approvedBy,uint invoiceId);
    event BeneficiaryAdded(string id);
    event BuyerAdded(string id);

    mapping(string => Beneficiary) beneficiaryDetailsOf;
    mapping(string => Buyer) buyerDetailsOf;
    mapping(uint => Sales) salesDetaildOf;
    mapping(string => mapping(uint=>string)) beneficiaryApprovals;
    mapping(string => mapping(uint=>string)) buyerApprovals;
    mapping(uint => string) invoiceApprovalStatus;

    function random() private  returns (uint){
        return uint(keccak256(abi.encodePacked(block.difficulty,block.timestamp,seed++)));
    }
    function registerBeneficiary(
        string memory beneficiaryId,
        string memory name,
        string memory uid,
        string memory pan,
        string memory bankAccount)
    public returns (string memory benId){
        benId = beneficiaryId;
        Beneficiary memory newBenificiary = Beneficiary(beneficiaryId,name,uid,pan,bankAccount);
        beneficiaryDetailsOf[beneficiaryId] = newBenificiary;
        emit BeneficiaryAdded(beneficiaryId);
        return benId;
    }

    function registerBuyer(     string memory buyerId,
        string memory name,
        string memory uid,
        string memory pan,
        string memory tan,
        string memory bankAccount)
    public returns (string memory bId){
        bId=buyerId;
        Buyer memory newBuyer = Buyer(buyerId,name,uid,pan,tan,bankAccount);
        buyerDetailsOf[buyerId] = newBuyer;
        emit BuyerAdded(buyerId);
        return buyerId;
    }

    function createInvoice(string memory item,
        string memory unit,
        uint256 amount,
        uint256 amountPerUnit,
        uint beneficiaryId,
        uint buyerId
    )
    public returns (uint  invoiceId){
        invoiceId = random();
        Sales memory invoice = Sales(invoiceId,item,unit,amount,amountPerUnit,now,beneficiaryId,buyerId,'ConcilationPending');
        salesDetaildOf[invoiceId] = invoice;
        invoiceApprovalStatus[invoiceId] = 'ConcilationPending';
        emit ItemSold(beneficiaryId,buyerId,invoiceId);
        return invoiceId;
    }

    function approveBeneficiary(string memory benificiaryId,uint invoiceId) public returns(bool success){
        beneficiaryApprovals[benificiaryId][invoiceId]='BenApproved';
        salesDetaildOf[invoiceId].status = 'BenApproved';
        return true;

    }

    function approveBuyer(string memory  buyerId,uint invoiceId) public returns(bool success){
        buyerApprovals[buyerId][invoiceId]='BuyApproved';
        salesDetaildOf[invoiceId].status = 'BuyApproved';
        return true;

    }
    function getBenificiary(string memory benificiaryId) public view returns (Beneficiary memory ben){
        ben = beneficiaryDetailsOf[benificiaryId];
        return ben;
    }

    function getBuyer(string memory buyerId) public view returns (Buyer memory buyer){
        buyer = buyerDetailsOf[buyerId];
        return buyer;
    }

    function getSalesDetails(uint invoiceId) public view returns (Sales memory sales){
        return salesDetaildOf[invoiceId];
    }

}