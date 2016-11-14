package main

import (
	"crypto/rand"
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type Trade struct {
	Entity    string  `json:"entity"`
	Char      string  `json:"char"`
	Event     string  `json:"event"`
	Action    string  `json:"action"`
	Price     float64 `json:"price"`
	Units     int     `json:"units"`
	Status    string  `json:"status"`
	Expiry    string  `json:"expiry"`
	Fulfilled int     `json:"fulfilled"`
}

type HappeningRegister struct {
	Char  string `json:"char"`
	Event string `json:"event"`
	User  string `json:"user"`
}

type AccountUser struct {
	userId    string
	AccountId string
	key       string
	lasKey    string
}

type User struct {
	StringKey     string
	StringKeyLast string
	UserID        string `json:"userID"`
	Status        string `json:"status"`
	Cash          int    `json:"cash"`
}

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
	users map[string]AccountUser
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	t.users = make(map[string]AccountUser)
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}
	err := stub.PutState("hello_world", []byte(args[0]))
	err1 := stub.PutState("mainWindow", []byte("<div>Username</div><div>Account Balance</div>"))
	if err != nil {
		return nil, err
	}
	if err1 != nil {
		return nil, err1
	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	} else if function == "registarLogin" {
		return t.write(stub, args)
	}

	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	//userKey = args[1]
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	}
	if function == "login" { //read a variable
		return t.login(stub, args)
	}
	if function == "getView" {
		return t.getView(stub, args)
	}
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var key, value string
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	key = args[0] //rename for funsies
	value = args[1]
	err = stub.PutState(key, []byte(value)) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name and hased password")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}

func (t *SimpleChaincode) getView(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var err error
	val, err := stub.GetState("mainWindow")
	return val, err
}

func (t *SimpleChaincode) login(stub *shim.ChaincodeStub, args []string) ([]byte, error) {
	var username string
	var ok bool
	username = args[0]
	accountId := args[1]
	userId := args[2]
	var accountUser AccountUser
	var err error
	accountUser, ok = t.users[username]
	if ok {

		fmt.Println("good")
	} else {
		var temp AccountUser
		temp.AccountId = accountId
		temp.userId = userId
		fmt.Println(temp)
		t.users[username] = temp
		fmt.Println(t.users)
	}
	err = nil
	u := uuid()
	var b = []byte(u)
	var val []byte
	val, err = stub.GetState("CurrentUsers")
	fmt.Print(val)
	fmt.Print(accountUser)
	//val, err := stub.PutState("")
	return b, err
}

func (t *SimpleChaincode) registerUserWithEnrollID(id string, enrollID string, role int, memberMetadata string, opt ...string) (string, error) {
	var err error
	///	tok, err = c///a.registerUserWithEnrollID(id, string, role, memberMetadata, opt)
	return "nil", err
}

func (t *SimpleChaincode) registarLogin(stub *shim.ChaincodeStub, args []string) (string, error) {
	t.login(stub, args)
	var err error
	var val []byte

	val, err = stub.GetState("CurrentUsers")
	fmt.Print(val)
	return "nil", err
}

// need to make a persistance class / data abstraction
func (t *SimpleChaincode) push(stub *shim.ChaincodeStub, structureName string, value []byte) ([]byte, error) {
	fmt.Printf("Running Push")
	index, err := t.getNextIndex(stub, "Last"+structureName)
	if err != nil {
		return nil, err
	}

	// Write the state back to the ledger
	var key string

	key = structureName + string(index)

	err = stub.PutState(key, []byte(value))
	if err != nil {
		return nil, err
	}

	return index, nil
}

func (t *SimpleChaincode) getNextIndex(stub *shim.ChaincodeStub, lastIDString string) ([]byte, error) {
	fmt.Printf("Running getNextIndex")

	var id int
	var err error
	lastID, err := stub.GetState(lastIDString)
	if err != nil {
		id = 1
	} else {
		temp, err := strconv.Atoi(string(lastID))
		if err != nil {
			return nil, err
		}
		id = temp + 1
	}

	idString := []byte(strconv.Itoa(id)) //not really an id "string".  the byte array / string in this language is a pain
	err = stub.PutState(lastIDString, idString)
	if err != nil {
		return nil, err
	}

	return idString, nil
}

func uuid() (uuid string) {

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	uuid = fmt.Sprintf("%X-%X-%X-%X-%X", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	fmt.Println(uuid)
	return
}
