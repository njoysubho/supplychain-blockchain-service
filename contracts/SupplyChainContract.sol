pragma solidity >=0.4.22;
pragma experimental ABIEncoderV2;
contract KisanSupplyChain{
    uint seed=0;
    bool paused = false;
    mapping(address=>bool) admins;
    modifier whenNotPaused(){
        require(paused != true);
        _;
    }

    modifier whenPaused(){
        require(paused == true);
        _;
    }

    modifier OnlyAdmin(){
        require(admins[msg.sender]==true);
        _;
    }

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
        string invoiceId;
        string item;
        string unit;
        uint256 amount;
        uint256 amountPerUnit;
        uint256 salesDate;
        string beneficiaryId;
        string buyerId;
        string status;
    }

    event ItemSold(string beneficiaryId,string buyerId,string invoiceId);
    event InvoiceApproved(string approvedBy,string invoiceId);
    event BeneficiaryAdded(string id);
    event BuyerAdded(string id);

    mapping(string => Beneficiary) beneficiaryDetailsOf;
    mapping(string => Buyer) buyerDetailsOf;
    mapping(string => Sales) salesDetaildOf;
    mapping(string => mapping(string=>string)) beneficiaryApprovals;
    mapping(string => mapping(string=>string)) buyerApprovals;
    mapping(string => string) invoiceApprovalStatus;

    constructor() public{
        admins[msg.sender] = true;
    }


    function random() private  returns (uint){
        return uint(keccak256(abi.encodePacked(block.difficulty,block.timestamp,seed++)));
    }
    function registerBeneficiary(
        string memory beneficiaryId,
        string memory name,
        string memory uid,
        string memory pan,
        string memory bankAccount) whenNotPaused OnlyAdmin
    public returns (string memory benId){
        benId = beneficiaryId;
        Beneficiary memory newBenificiary = Beneficiary(beneficiaryId,name,uid,pan,bankAccount);
        beneficiaryDetailsOf[beneficiaryId] = newBenificiary;
        emit BeneficiaryAdded(beneficiaryId);
        return benId;
    }

    function registerBuyer( string memory buyerId,
        string memory name,
        string memory uid,
        string memory pan,
        string memory tan,
        string memory bankAccount) whenNotPaused OnlyAdmin
    public returns (string memory bId){
        bId=buyerId;
        Buyer memory newBuyer = Buyer(buyerId,name,uid,pan,tan,bankAccount);
        buyerDetailsOf[buyerId] = newBuyer;
        emit BuyerAdded(buyerId);
        return buyerId;
    }

    function createInvoice(string memory id,
        string memory item,
        string memory unit,
        uint256 amount,
        uint256 amountPerUnit,
        string memory beneficiaryId,
        string memory buyerId
    ) whenNotPaused
    public returns (string  memory invoiceId){
        invoiceId = id;
        Sales memory invoice = Sales(invoiceId,item,unit,amount,amountPerUnit,now,beneficiaryId,buyerId,'ConcilationPending');
        salesDetaildOf[invoiceId] = invoice;
        invoiceApprovalStatus[invoiceId] = 'ConcilationPending';
        emit ItemSold(beneficiaryId,buyerId,invoiceId);
        return invoiceId;
    }

    function approveBeneficiary(string memory benificiaryId,string memory invoiceId) whenNotPaused public returns(bool success){
        beneficiaryApprovals[benificiaryId][invoiceId]='BenApproved';
        salesDetaildOf[invoiceId].status = 'BenApproved';
        return true;

    }

    function approveBuyer(string memory  buyerId,string memory invoiceId) whenNotPaused public returns(bool success){
        buyerApprovals[buyerId][invoiceId]='BuyApproved';
        salesDetaildOf[invoiceId].status = 'BuyApproved';
        return true;

    }
    function getBenificiary(string memory benificiaryId) whenNotPaused public view returns (Beneficiary memory ben){
        ben = beneficiaryDetailsOf[benificiaryId];
        return ben;
    }

    function getBuyer(string memory buyerId) whenNotPaused public view returns (Buyer memory buyer){
        buyer = buyerDetailsOf[buyerId];
        return buyer;
    }

    function getSalesDetails(string memory invoiceId) whenNotPaused public view returns (Sales memory sales){
        return salesDetaildOf[invoiceId];
    }

    function pause() whenNotPaused OnlyAdmin public {
        paused=true;
    }

    function unPause() whenPaused OnlyAdmin public {
        paused=false;
    }

    function addAdmin(address admin) whenNotPaused OnlyAdmin public returns(bool success){
        admins[admin]=true;
        return true;
    }

    function isAdmin(address mayBeAdmin) whenNotPaused public view returns (bool result){
        return admins[mayBeAdmin];
    }

}