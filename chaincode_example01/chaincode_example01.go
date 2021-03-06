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

package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/openblockchain/obc-peer/openchain/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

var A, B string
var Aval, Bval, X int

// Run callback representing the invocation of a chaincode
// This chaincode will manage two accounts A and B and will transfer X units from A to B upon invoke
func (t *SimpleChaincode) Run(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	var err error

	// Handle different functions
	if function == "init" {
		if len(args) != 4 {
			return nil, errors.New("Incorrect number of arguments. Expecting 4")
		}

		// Initialize the chaincode
		A = args[0]
		Aval, err = strconv.Atoi(args[1])
		if err != nil {
			return nil, errors.New("Expecting integer value for asset holding")
		}
		B = args[2]
		Bval, err = strconv.Atoi(args[3])
		if err != nil {
			return nil, errors.New("Expecting integer value for asset holding")
		}
		fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)

		/************
				// Write the state to the ledger
				err = stub.PutState(A, []byte(strconv.Itoa(Aval))
				if err != nil {
					return nil, err
				}

				stub.PutState(B, []byte(strconv.Itoa(Bval))
				err = stub.PutState(B, []byte(strconv.Itoa(Bval))
				if err != nil {
					return nil, err
				}
		************/
	} else if function == "invoke" {
		// Transaction makes payment of X units from A to B
		X, err = strconv.Atoi(args[0])
		Aval = Aval - X
		Bval = Bval + X
		fmt.Printf("Aval = %d, Bval = %d\n", Aval, Bval)
	}

	return nil, nil
}

// Query callback representing the query of a chaincode
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	return nil, nil
}

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
