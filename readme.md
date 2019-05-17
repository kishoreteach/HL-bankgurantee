
Please note this project is in WIP and is being constantly updated so may not work as expected. 

## Table of contents
* [General info](#general-info)
* [Setup](#setup)

## General info

The project uses the Blockchain business network as defined for 
first-network(https://hyperledger-fabric.readthedocs.io/en/release-1.4/build_network.html)


## Setup
To run this project, 

navigate to HL-bankgurantee/first-network/ and start ./byfn.sh , this will start the business network
The network has 2 Orgs (org1 and org2) and each org has 2 peers

To invoke blockchain logic for bank gurantee please run following cli commands

1) Install the chaincode
docker exec cli peer chaincode install -n letter10 -v 1.0 -l golang -p github.com/chaincode/letterofGrnt/go/




