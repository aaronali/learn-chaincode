/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at
  http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/


//todo:  need to make consitent status.  need better way to take them out of the process when closed.  maybe use an enum
//todo: add application security to get user names
//todo:  make user into account



package main

import (
	"errors"
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/op/go-logging"
)

var logger = shim.NewLogger("myChaincodeApi")


// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
	bl ChaincodeBusinessLayer
}


//********************************************************************************************************
//****      Blockchain API functions                                                                  ****
//********************************************************************************************************



//Init the blockchain.  populate a 2x2 grid of potential events for users to buy
func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Printf("Init called, initializing chaincode")
	logger.Warningf("myChainCodeDAO : %s", "init")
	//initialize our repositories
	t.bl.initObjects(stub)
	
	t.bl.writeOut("in init")
	
	//Register some users.  this would normally happen via the UI but we will do it here to simplify
	_, err := t.bl.registerUser("BANK")
	if err != nil {
		return nil, err
	}
	
	_, err = t.bl.registerUser("David")
	if err != nil {
		return nil, err
	}
	
	_, err = t.bl.registerUser("Wesley")
	if err != nil {
		return nil, err
	}
	
	_, err = t.bl.registerUser("Aaron")
	if err != nil {
		return nil, err
	}
	
	_, err = t.bl.registerUser("Anu")
	if err != nil {
		return nil, err
	}
	
	_, err = t.bl.registerUser("Mustafa")
	if err != nil {
		return nil, err
	}
	
	_, err = t.bl.registerUser("user_type1_0")
	if err != nil {
		return nil, err
	}
	
	//register our securities and offer them for sale
	_, err = t.bl.registerSecurity("Gold", "Gold killed")
	if err != nil {
		return nil, err
	}
	_, err = t.bl.registerSecurity("Silver","Silver Shares")
	if err != nil {
		return nil, err
	}
	
	_, err = t.bl.registerSecurity("Nickle", "Nickle Shares")
	if err != nil {
		return nil, err
	}
	_, err = t.bl.registerSecurity("Bronze", "Bronze Shares")
	if err != nil {
		return nil, err
	}
	
	 
	
	//the bank does an IPO
	_, err = t.bl.registerTrade("ask", "BANK", "Gold", defaultPrice, 100, "")
	if err != nil {
		return nil, err
	}
	
	
	_, err = t.bl.registerTrade("ask", "BANK", "Silver", defaultPrice, 100, "")
	if err != nil {
		return nil, err
	}
	
	_, err = t.bl.registerTrade("ask", "BANK", "Bronze", defaultPrice, 100, "")
	if err != nil {
		return nil, err
	}
	
	_, err = t.bl.registerTrade("ask", "BANK", "Nickle", defaultPrice, 100, "")
	if err != nil {
		return nil, err
	}
	//
		_, err = t.bl.registerTrade("bid", "user_type1_0", "Gold", defaultPrice, 100, "")
	if err != nil {
		return nil, err
	}
	
	
	_, err = t.bl.registerTrade("bid", "user_type1_0", "Silver", defaultPrice, 100, "")
	if err != nil {
		return nil, err
	}
	
	_, err = t.bl.registerTrade("bid", "user_type1_0", "Bronze", defaultPrice, 100, "")
	if err != nil {
		return nil, err
	}
	
	_, err = t.bl.registerTrade("bid", "user_type1_0", "Nickle",defaultPrice, 100, "")
	if err != nil {
		return nil, err
	}
		
	
	_, err = t.bl.registerTrade("bid", "user_type1_0", "Nickle", defaultPrice, 100, "")
	if err != nil {
		return nil, err
	}
	
	_, err = t.bl.registerTrade("bid", "Aaron", "Nickle", defaultPrice, 100, "")
	if err != nil {
		return nil, err
	}
	
	t.bl.writeOut("Before dividend")
	//offer payoff anyone with Jaime,Killed (Aaron)
	_, err = t.bl.dividend("Nickle", 50)
	if err != nil {
		t.bl.writeOut("in init: after dividend in err != nil")
		return nil, err
	}
	
	
	t.bl.writeOut("Before return")
	return nil, nil
}



// Invoke callback representing the invocation of a chaincode
func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Printf("Invoke called, determining function")
	
	t.bl.initObjects(stub) //for some reason the stub changes each call
	// Handle different functions
	if function == "init" {
		fmt.Printf("Function is init")
		return t.Init(stub, function, args)
	} else if function == "ask" || function == "bid" {
		// offers squares up for sale as initial public offering
		fmt.Printf("Function is trade")
		
		if len(args) != 5 {
			return nil, errors.New("Incorrect number of arguments. Expecting registerTrade(tradeType, userid, security, price, units, expiry)")
		}
		
		price, err := strconv.ParseFloat(args[2], 64)
		if err != nil {
			return nil, err
		}
		
		units, err := strconv.Atoi(args[3])
		if err != nil {
			return nil, err
		}
		
		return t.bl.registerTrade(function, args[0], args[1], price, units, args[4])
	} else if function == "dividend" {
		// enter an an character event happening in the show.  pays out to users holding squares
		fmt.Printf("Function is ask")
		
		amount, err := strconv.Atoi(args[1])
		if err != nil {
			return nil, err
		}
		return t.bl.dividend(args[0], amount)
	} else if function == "registeruser" {
		// matches trades and excecutes any matches
		fmt.Printf("Function is registeruser")
		userID := args[0]
		return t.bl.registerUser(userID)
	}
	
	return nil, errors.New("Received unknown function invocation")
}




// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Printf("Query called, determining function") 
	
	t.bl.initObjects(stub)
	// Handle different functions
	if function == "holdings" {
		// query a users holdings
		return t.bl.holdings(args[0])  //userID
	} else if function == "ballance" {
		// query a users cash on hand
		return t.bl.ballance(stub, args[0])	//userID
	} else if function == "users" {
		// query for list of users
		return t.bl.users()
	} else if function == "securities" {
		// query for list of securities
		return t.bl.securities()
	} else {
		fmt.Printf("Function is query")
		return nil, errors.New("Invalid query function name. Expecting holdings, ballance, users or securities")
	}	
	
	return nil, nil
}





func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
	
	logger.SetLevel(shim.LogInfo)

    logLevel, _ := shim.LogLevel(os.Getenv("SHIM_LOGGING_LEVEL"))
    shim.SetLoggingLevel(logLevel)
}  