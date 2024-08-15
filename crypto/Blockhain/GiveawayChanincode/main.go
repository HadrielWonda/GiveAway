package main

import (
	"fmt"

	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
)

type giveawayChaincode struct {
}

func (c *giveawayChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	return peer.Response{
		Status:  200,
		Message: "successfully initiated",
		Payload: nil,
	}
}
func (c *giveawayChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {

	funName, args := stub.GetFunctionAndParameters()
	if funName == "setSecret" {
		return putSecret(stub, args)
	}
	if funName == "userRegistration" {
		return userRegistration(stub, args)
	}
	if funName == "NGORegistration" {
		return NGORegistration(stub, args)
	}
	if funName == "getUser" {
		return getUser(stub, args)
	}
	if funName == "getNgo" {
		return getNgo(stub, args)
	}
	if funName == "userGateway" {
		return userGateway(stub, args)
	}
	if funName == "ngoGateway" {
		return ngoGateway(stub, args)
	}
	if funName == "getAllMission" {
		return getAllMission(stub)
	}
	return peer.Response{
		Status:  200,
		Message: "successfully initiated",
		Payload: nil,
	}
}
func main() {
	err := shim.Start(new(giveawayChaincode))
	if err != nil {
		fmt.Println(err.Error())
	}
}
