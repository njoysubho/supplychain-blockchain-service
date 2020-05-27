pragma solidity >=0.4.22;
pragma experimental ABIEncoderV2;
contract KisanSupplyChain{
    uint seed=0;
   struct Beneficiary {
       uint beneficiaryId;
       string name;
       string uid;
       string pan;
       string bankAccount;
   }

   struct Buyer{
       uint buyerId;
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


   mapping(uint => Beneficiary) beneficiaryDetailsOf;
   mapping(uint => Buyer) buyerDetailsOf;
   mapping(uint => Sales) salesDetaildOf;
   mapping(uint => mapping(uint=>string)) beneficiaryApprovals;
   mapping(uint => mapping(uint=>string)) buyerApprovals;
   mapping(uint => string) invoiceApprovalStatus;

   function random() private  returns (uint){
        return uint(keccak256(abi.encodePacked(block.difficulty,block.timestamp,seed++)));
    }
   function registerBeneficiary(string memory name,
                                string memory uid,
                                string memory pan,
                                string memory bankAccount)
    public returns (uint benificiaryId){
        benificiaryId = random();
        Beneficiary memory newBenificiary = Beneficiary(benificiaryId,name,uid,pan,bankAccount);
        beneficiaryDetailsOf[benificiaryId] = newBenificiary;
        return benificiaryId;
    }

    function registerBuyer(string memory name,
                                string memory uid,
                                string memory pan,
                                string memory tan,
                                string memory bankAccount)
    public returns (uint buyerId){
        buyerId = random();
        Buyer memory newBuyer = Buyer(buyerId,name,uid,pan,tan,bankAccount);
        buyerDetailsOf[buyerId] = newBuyer;
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

    function approveBeneficiary(uint benificiaryId,uint invoiceId) public returns(bool success){
        beneficiaryApprovals[benificiaryId][invoiceId]='BenApproved';
        salesDetaildOf[invoiceId].status = 'BenApproved';
        return true;

    }

    function approveBuyer(uint buyerId,uint invoiceId) public returns(bool success){
        buyerApprovals[buyerId][invoiceId]='BuyApproved';
        salesDetaildOf[invoiceId].status = 'BuyApproved';
        return true;

    }
    function getBenificiary(uint benificiaryId) public view returns (Beneficiary memory ben){
        ben = beneficiaryDetailsOf[benificiaryId];
        return ben;
    }

    function getBuyer(uint buyerId) public view returns (Buyer memory buyer){
        buyer = buyerDetailsOf[buyerId];
        return buyer;
    }

    function getSalesDetails(uint invoiceId) public view returns (Sales memory sales){
        return salesDetaildOf[invoiceId];
    }

}