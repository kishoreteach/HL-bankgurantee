package main

import (
	"fmt"	
	"errors"
	"encoding/json"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)


type assetWorkflow struct {
	
}

const (

	PROPOSE		= "PROPOSE"
	ISSUED		= "ISSUED"
	CANCEL		= "CANCEL"
	ACCEPTED	= "ACCEPTED"	
	REDEEM		= "REDEEM"
	COUNTER		= "COUNTER"
	

)


type letterOfGrnt struct {	
	Id		string	`json:"id"`
	ClientID	string	`json:"clientid"`
	CreationDate	string	`json:"creationDate"`
	ExpirationDate	string	`json:"expirationDate"`
	Beneficiary	string	`json:"beneficiary"`
	Amount		int	`json:"amount"`
	DocumentHash	string	`json:"documenthash"`
	Status		string	`json:"status"`	
	Comments	string	`json:"comment"`
}

func (t *assetWorkflow) Init(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("Initializing Asset Workflow")
	return shim.Success(nil)
}

func (t *assetWorkflow) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	
	fmt.Println("Asset Workflow Invoke")
	
	function, args := stub.GetFunctionAndParameters()
		
	if function == "proposeLetterOfGurantee" {
		
		return t.proposeLetterOfGurantee(stub, args)
		
	} else if function == "issueLetterOfGurantee" {

		return t.issueLetterOfGurantee(stub, args)

        } else if function == "viewLetterOfGurantee" {

		return t.viewLetterOfGurantee(stub, args)

	} else if function == "counterLetterOfGurantee" {

		return t.counterLetterOfGurantee(stub, args)

	} else if function == "acceptLetterOfGurantee" {

		return t.acceptLetterOfGurantee(stub, args)

	} else if function == "cancelLetterOfGurantee" {

		return t.cancelLetterOfGurantee(stub, args)

	} else if function == "redeemLetterOfGurantee" {

		return t.redeemLetterOfGurantee(stub, args)
	}
		
	return shim.Success(nil)
}


func (s *assetWorkflow) proposeLetterOfGurantee(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("Invokde Proposal flow")	
	var err error
	var Key string
		
	
	if len(args) < 3 {

		err = errors.New(fmt.Sprintf("Incorrect number of arguments. Expecting at least 3: {id, clientid, Expiry Date}. Found %d", len(args)))

		return shim.Error(err.Error())
	}

	// create the key
	Key, err = getBankLettereKey(stub, args[0],args[1])	

	if err != nil {
		return shim.Error(err.Error())
	}

	var id = args[0]
	var clientid = args[1]
	var creationdate = args[2]
	var expirationdate = args[3]
	var beneficiary = args[4]
	var documenthash = args[6]
	var status = ISSUED
	var comments = args[8]
	
	var amount, err1 = strconv.Atoi(string(args[5]))
	if err1 != nil {
		return shim.Error(err1.Error())
	}
	
	letterOfGuarantee_1 := &letterOfGrnt{id,clientid,creationdate,expirationdate,beneficiary,amount,documenthash,status,comments}
	
	letterOfGuaranteeJSONasBytes, err := json.Marshal(letterOfGuarantee_1)
	
	if err != nil {
		return shim.Error(err.Error())
	}
	
	err = stub.PutState(Key, letterOfGuaranteeJSONasBytes)
	
	if err != nil {

		return shim.Error(err.Error())

	}
	
	return shim.Success(nil)
}




func (s *assetWorkflow) issueLetterOfGurantee(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	fmt.Println("Invokde issue flow")
	fmt.Println("Invokde Proposal flow")	
	var err error
	var letterOfGrntBytes []byte
	var valAsbytes []byte
	var letterOfGrntJSON letterOfGrnt



	valAsbytes, err = stub.GetState(args[0]) //get the marble from chaincode state
	
	if err != nil {
		return shim.Error(err.Error())
	}

	// Unmarshal the JSON
	err = json.Unmarshal([]byte(valAsbytes), &letterOfGrntJSON)
	
	if err != nil {
		return shim.Error(err.Error())
	}

	if letterOfGrntJSON.Status == ISSUED {

	fmt.Printf("Bank Letter of Guarantee %s already issued", args[0])

	} else {

		letterOfGrntJSON.Status = ISSUED

		letterOfGrntBytes, err = json.Marshal(letterOfGrntJSON)
		err = stub.PutState(args[0], letterOfGrntBytes)

		if err != nil {
			return shim.Error("Error marshaling E/L structure")
		}
	}
	return shim.Success(nil)	
}

func (s *assetWorkflow) viewLetterOfGurantee(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	
	fmt.Println("Inside query flow")
	if len(args) != 1 {

		return shim.Error("Incorrect number of arguments during view letter call. Expecting 1")
	}
	
	letterOfGuaranteeBytes, _ := stub.GetState(args[0])

	return shim.Success(letterOfGuaranteeBytes)

		
}

func (s *assetWorkflow) counterLetterOfGurantee(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("Invokde counter flow")	

	return shim.Success(nil)	
}

func (s *assetWorkflow) acceptLetterOfGurantee(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("Invokde accept flow")		

	return shim.Success(nil)	
}

func (s *assetWorkflow) redeemLetterOfGurantee(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("Invokde redeem flow")	
	
	return shim.Success(nil)	
}

func (s *assetWorkflow) cancelLetterOfGurantee(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("Invokde cancel flow")	
	
	return shim.Success(nil)	
}

func getBankLettereKey(stub shim.ChaincodeStubInterface, bankname string, custid string) (string, error) {

	ltrkey, err := stub.CreateCompositeKey("BankLetter", []string{bankname,custid})
	fmt.Printf("inside key creation: %s", ltrkey)
	if err != nil {

		return "", err

	} else {

		return ltrkey, nil
	}
}

func main() {

	err := shim.Start(new(assetWorkflow))

	if err != nil {

		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}




