const {accounts, contract, web3} = require('@openzeppelin/test-environment');
const assert = require('assert');
const chai = require('chai');
const chaiAsPromised = require("chai-as-promised");
chai.use(chaiAsPromised);
const expect = chai.expect
const should = chai.should();
const {
    BN,           // Big Number support
    constants,    // Common constants, like the zero address and largest integers
    expectEvent,  // Assertions for emitted events
    expectRevert, // Assertions for transactions that should fail
} = require('@openzeppelin/test-helpers');
const Scm = contract.fromArtifact('SupplyChainContract');

describe('scm-test', () => {
    const [owner, admin] = accounts
    beforeEach(async () => {
        this.scmContract = await Scm.new({from: owner})
        await this.scmContract.init({from: owner})
        await this.scmContract.grantRole(web3.utils.asciiToHex("DEFAULT_ADMIN_ROLE"), admin
            , {from: owner})
    });
describe('buyer tests',()=>{

    it('should be deployed', async () => {
        assert.ok(this.scmContract.address)
    });

    it("should register buyer", async () => {
        await this.scmContract.registerBuyer("buyer1", "buyer1name", "buyer1uid", "pan", "tan"
            , "bankaccount",{from:owner});
    });

    it('should not let register buyer to non owner', async ()=> {
        await this.scmContract.registerBuyer("buyer1", "buyer1name", "buyer1uid", "pan", "tan"
            , "bankaccount",{from:admin}).should.be.rejected
    });

    it("should not register duplicate buyer", async () => {
        await this.scmContract.registerBuyer("buyer1", "buyer1name", "buyer1uid", "pan", "tan"
            , "bankaccount",{from:owner});
        await this.scmContract.registerBuyer("buyer1", "buyer1name", "buyer1uid", "pan", "tan"
            , "bankaccount",{from:owner}).should.be.rejected;
    });

    it("should not register buyer with empty id", async () => {
        await this.scmContract.registerBuyer("", "buyer1name", "buyer1uid", "pan", "tan"
            , "bankaccount",{from:owner}).should.be.rejected;
    });
    it("should not register buyer with empty name", async () => {
        await this.scmContract.registerBuyer("buyerId", "", "buyer1uid", "pan", "tan"
            , "bankaccount",{from:owner}).should.be.rejected;
    });


    it("should not register buyer with empty pan", async () => {
        await this.scmContract.registerBuyer("buyerId", "name", "buyer1uid", "", "tan"
            , "bankaccount",{from:owner}).should.be.rejected;
    });


    it("should not register buyer with empty tan", async () => {
        await this.scmContract.registerBuyer("buyerId", "name", "buyer1uid", "pan", ""
            , "bankaccount",{from:owner}).should.be.rejected;
    });


    it("should not register buyer with empty bankAccount", async () => {
        await this.scmContract.registerBuyer("buyerId", "name", "buyer1uid", "pan", "tan"
            , "",{from:owner}).should.be.rejected;
    });
});

describe("seller test",()=>{

    it("should register seller", async () => {
        await this.scmContract.registerSeller("seller1", "seller1name", "seller1uid", "pan","bankAccount",{from:owner});
    });

    it('should not let register seller to non owner', async ()=> {
        await this.scmContract.registerSeller("seller1", "seller1name", "seller1uid", "pan"
            , "bankaccount",{from:admin}).should.be.rejected
    });

    it("should not register duplicate seller", async () => {
        await this.scmContract.registerSeller("seller1", "seller1name", "seller1uid", "pan","bankAccount",{from:owner});
        await this.scmContract.registerSeller("seller1", "seller1name", "seller1uid", "pan"
            , "bankaccount",{from:owner}).should.be.rejected;
    });

    it("should not register seller with empty id", async () => {
        await this.scmContract.registerSeller("", "seller1name", "seller1uid", "pan"
            , "bankaccount",{from:owner}).should.be.rejected;
    });
    it("should not register seller with empty name", async () => {
        await this.scmContract.registerSeller("seller1", "", "seller1uid", "pan"
            , "bankaccount",{from:owner}).should.be.rejected;
    });


    it("should not register seller with empty pan", async () => {
        await this.scmContract.registerSeller("seller1", "seller1name", "seller1uid", ""
            , "bankaccount",{from:owner}).should.be.rejected;
    });

    it("should not register buyer with empty bankAccount", async () => {
        await this.scmContract.registerSeller("seller1", "seller1name", "seller1uid", "pan"
            , "",{from:owner}).should.be.rejected;
    });
});

describe("test invoice",()=>{
    it('should create invoice for valid transaction',  async ()=> {
        await this.scmContract.registerSeller("seller1", "seller1name", "seller1uid", "pan","bankAccount",{from:owner});
        await this.scmContract.registerBuyer("buyer1", "buyer1name", "buyer1uid", "pan","tan","bankAccount",{from:owner});
        const invoice = await this.scmContract.createInvoice("invoiceId","XEBF","100","2500","25","seller1","buyer1",{from:owner})
        expectEvent(invoice,'ItemSold',{
            sellerId:'seller1',
            buyerId: 'buyer1',
            invoiceId:'invoiceId'
        })
    });

    it('should not create invoice for unknown seller',  async ()=> {
        await this.scmContract.registerBuyer("buyer1", "buyer1name", "buyer1uid", "pan","tan","bankAccount",{from:owner});
        await this.scmContract.createInvoice("invoiceId","XEBF","100","2500","25","seller1","buyer1",{from:owner}).should.be.rejected;
    });
});

});


