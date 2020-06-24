pragma solidity ^0.6.0;
pragma experimental ABIEncoderV2;

import "@openzeppelin/contracts-ethereum-package/contracts/access/AccessControl.sol";
import "@openzeppelin/contracts-ethereum-package/contracts/utils/Pausable.sol";


contract SupplyChainContract is AccessControlUpgradeSafe, PausableUpgradeSafe {
    uint seed;
    bytes32 public constant ADMIN_ROLE = keccak256("ADMIN_ROLE");

    function init() public initializer {
        __AccessControl_init();
        _setupRole(DEFAULT_ADMIN_ROLE, _msgSender());
        seed = 0;
    }

    function _msgSender() internal view override returns (address payable) {
        return msg.sender;
    }

    function getSender() public view returns(address){
        return msg.sender;
    }
    function _msgData() internal view override returns (bytes memory) {
        this;
        // silence state mutability warning without generating bytecode - see https://github.com/ethereum/solidity/issues/2691
        return msg.data;
    }

    struct Seller {
        string sellerId;
        string name;
        string uid;
        string pan;
        string bankAccount;
        bool isActive;
    }

    struct Buyer {
        string buyerId;
        string name;
        string uid;
        string pan;
        string tan;
        string bankAccount;
        bool isActive;
    }

    struct Sales {
        string invoiceId;
        string item;
        string unit;
        uint256 amount;
        uint256 amountPerUnit;
        uint256 salesDate;
        string beneficiaryId;
        string buyerId;
        string status;
        bool isValid;
    }

    event ItemSold(string sellerId, string buyerId, string invoiceId);
    event InvoiceApproved(string approvedBy, string invoiceId);
    event SellerAdded(string id);
    event BuyerAdded(string id);

    mapping(string => Seller) sellerDetailsOf;
    mapping(string => Buyer) buyerDetailsOf;
    mapping(string => Sales) salesDetailOf;
    mapping(string => mapping(string => string)) sellerApprovals;
    mapping(string => mapping(string => string)) buyerApprovals;
    mapping(string => string) invoiceApprovalStatus;

    function random() private returns (uint){
        return uint(keccak256(abi.encodePacked(block.difficulty, block.timestamp, seed++)));
    }

    function registerSeller(
        string memory sellerId,
        string memory name,
        string memory uid,
        string memory pan,
        string memory bankAccount) whenNotPaused
    public returns (string memory sId){
        require(hasRole(DEFAULT_ADMIN_ROLE, _msgSender()), "Only an admin can register");
        require(!sellerDetailsOf[sellerId].isActive, "Seller already exists");
        require(bytes(sellerId).length!=0,"Id cannot be empty");
        require(bytes(name).length!=0,"Name cannot be empty");
        require(bytes(pan).length!=0,"Pan cannot be empty");
        require(bytes(bankAccount).length!=0,"BankAccount cannot be empty");
        sId = sellerId;
        Seller memory newSeller = Seller(sellerId, name, uid, pan, bankAccount, true);
        sellerDetailsOf[sellerId] = newSeller;
        emit SellerAdded(sId);
        return sId;
}

    function registerBuyer(string memory buyerId,
        string memory name,
        string memory uid,
        string memory pan,
        string memory tan,
        string memory bankAccount) whenNotPaused
    public returns (string memory bId){
        require(hasRole(DEFAULT_ADMIN_ROLE,_msgSender()), "Only an admin can register");
        require(!buyerDetailsOf[buyerId].isActive, "Buyer already exists");
        require(bytes(buyerId).length!=0,"Id cannot be empty");
        require(bytes(name).length!=0,"Name cannot be empty");
        require(bytes(pan).length!=0,"Pan cannot be empty");
        require(bytes(tan).length!=0,"Tan cannot be empty");
        require(bytes(bankAccount).length!=0,"BankAccount cannot be empty");
        bId = buyerId;
        Buyer memory newBuyer = Buyer(buyerId, name, uid, pan, tan, bankAccount, true);
        buyerDetailsOf[buyerId] = newBuyer;
        emit BuyerAdded(buyerId);
        return buyerId;
    }

    function createInvoice(string memory id,
        string memory item,
        string memory unit,
        uint256 amount,
        uint256 amountPerUnit,
        string memory sellerId,
        string memory buyerId
    ) whenNotPaused
    public returns (string  memory invoiceId){
        require(buyerDetailsOf[buyerId].isActive, "Buyer does not exist");
        require(sellerDetailsOf[sellerId].isActive, "Seller does not  exists");
        invoiceId = id;
        Sales memory invoice = Sales(invoiceId, item, unit, amount, amountPerUnit, now, sellerId, buyerId, 'ConcilationPending', true);
        salesDetailOf[invoiceId] = invoice;
        invoiceApprovalStatus[invoiceId] = 'ConcilationPending';
        emit ItemSold(sellerId, buyerId, invoiceId);
        return invoiceId;
    }

    function approveBySeller(string memory sellerId, string memory invoiceId) whenNotPaused public returns (bool success){
        require(sellerDetailsOf[sellerId].isActive, "Seller does not  exists");
        require(salesDetailOf[invoiceId].isValid, "InvoiceId not valid");
        sellerApprovals[sellerId][invoiceId] = 'BenApproved';
        salesDetailOf[invoiceId].status = 'BenApproved';
        return true;
    }

    function approveByBuyer(string memory buyerId, string memory invoiceId) whenNotPaused public returns (bool success){
        require(buyerDetailsOf[buyerId].isActive, "Buyer does not  exists");
        require(salesDetailOf[invoiceId].isValid, "InvoiceId not valid");
        buyerApprovals[buyerId][invoiceId] = 'BuyerApproved';
        salesDetailOf[invoiceId].status = 'BuyerApproved';
        return true;

    }

    function getSeller(string memory sellerId) whenNotPaused public view returns (Seller memory seller){
        seller = sellerDetailsOf[sellerId];
        return seller;
    }

    function getBuyer(string memory buyerId) whenNotPaused public view returns (Buyer memory buyer){
        buyer = buyerDetailsOf[buyerId];
        return buyer;
    }

    function getSalesDetails(string memory invoiceId) whenNotPaused public view returns (Sales memory sales){
        return salesDetailOf[invoiceId];
    }

}