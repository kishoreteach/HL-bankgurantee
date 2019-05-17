
Please note this project is in WIP

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

2) Instantiate the chaincode
docker exec cli peer chaincode instantiate -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n letter10 -l golang -v 1.0 -c '{"Args":["init","a","100","b","200"]}' -P 'AND ('\''Org1MSP.peer'\'','\''Org2MSP.peer'\'')'

3) Invoke the Letter Publish method

docker exec cli peer chaincode invoke -o orderer.example.com:7050 --tls true --cafile /opt/gopath/src/github.com/hyperledger/fabric/peer/crypto/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem -C mychannel -n letter10  -c '{"Function": "proposeLetterOfGurantee", "Args": ["001", "cust001", "22ndjan2018","22ndjan2019", "cust001", "100","doc", "status","comm"]}'

4) Invoke the Query

peer chaincode query -C mychannel -n letter10 -c '{"Args":["viewLetterOfGurantee","BankLetter001cust001"]}'

Please note this project is in WIP





