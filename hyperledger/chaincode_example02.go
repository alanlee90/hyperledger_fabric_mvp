package main


import (
	"fmt"
	"strconv"
	"encoding/json"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
	"bytes"
)

type SimpleChaincode struct {

}

type ValInfo struct {
	Id	string	`json:"id"`
	Val	int	`json:"val,string"`
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Chaincode Practice Init")
	return shim.Success(nil)
}

func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("ex02 Invoke")
	function, args := stub.GetFunctionAndParameters()
	if function == "makeIdAndVal" {
		// Make payment of X units from A to B
		return t.invoke(stub, args)
	}else if function == "query" {
		// the old "Query" is now implemented in invoke
		return t.query(stub, args)
	}else if function == "moveVal" {
		// Move val from A to B or B to A
		return t.moveVal(stub, args)
	}else if function == "queryById" {
		// query info by Id
		return t.queryById(stub, args)
	}

	return shim.Error("Invalid invoke function name.")
}

func (t *SimpleChaincode) makeIdAndVal(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	id := args[0]
	val, err := strconv.Atoi(args[1])
	fmt.Println("log>Input Id value : " + id)
	valInfo := &ValInfo{id, val}
	valInfoBytes, err := json.Marshal(valInfo)
	if err != nil {
		return shim.Error(err.Error())
	}

	err = stub.PutState(id, valInfoBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("putState complete")
	return shim.Success(nil)
}

func (t *SimpleChaincode) queryById(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	fmt.Println("queryById 함수 호출")
	id := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"id\":\"%s\"}}", id)
	queryResults, err := getQueryResultForQueryString(stub, queryString)

	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func (t *SimpleChaincode) moveVal(stub shim.ChaincodeStubInterface, args []string) pb.Response{
	fmt.Println("call moveVal Chaincode")
	from := args[0]
	to := args[1]
	X, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("Invalid transaction amount, expectiong a integer value")
	}
	FromBytes, err := stub.GetState(from)
	if err != nil {
		return shim.Error("Failed to get fromId:" + err.Error())
	} else if FromBytes == nil {
		return shim.Error("Data does not exist")
	}
	toBytes, err := stub.GetState(to)
	if err != nil {
		return shim.Error("Failed to get fromId:" + err.Error())
	} else if toBytes == nil {
		return shim.Error("Data does not exist")
	}
	fromVal := ValInfo{}
	err = json.Unmarshal(FromBytes,&fromVal)
	if err != nil {
		return shim.Error(err.Error())
	}
	toVal := ValInfo{}
	err = json.Unmarshal(toBytes,&toVal)
	if err != nil {
		return shim.Error(err.Error())
	}
	fromVal.Val = fromVal.val - X
	toVal.Val = toVal.Val + X
	fromValBytes, _ := json.Marshal(fromVal)
	err = stub.PutState(from, fromValBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	toValBytes, _ := json.Marshal(toVal)
	err = stub.PutState(to, toValBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (t *SimpleChaincode) query(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("call query method")
	queryString: = "{\"selector\":{}}"
	fmt.Println("queryString" + queryString)
	queryResults, err := getQueryResultForQueryString(stub, queryString)
}