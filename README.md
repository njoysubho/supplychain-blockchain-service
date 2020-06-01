# Smart Contract for Farmer Supplychain

This repo contains Solidity Smart Contract (which is still evolving) and a Golang REST
API on top of Solidity Contract. 

Smart Contract depict a scenario where Farmers sell their produce to procurer.
The contract stores invoice and state of approval for those invoices .

Golang APIs provide clients to interact with underlying ethereum contract using 
golang-ethereum client.

The contract is deployed on Rinkeby test metwork and all transaction is happened via a test
wallet .

The golang app is deployed on GKE using Github Action .

This project uses infura to deploy to Rinkeby.

The app will not work without a wallet and rinkeby api key .