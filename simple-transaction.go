/*

*/

package main

import (
    "errors"
    "fmt"

    "github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
}

var customer_id string
var device_id string
var trx_type string
var trx_total string

// Initialize
func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    var err error

    if len(args) != 4 {
        return nil, errors.New("Incorrect number of arguments.  Expecting 4")
    }

    customer_id = args[0]
    device_id = args[1]
    trx_type = args[2]
    trx_total = args[3]

    if err != nil {
        return nil, errors.New("Could not populate variables")
    }

    fmt.Println(customer_id,device_id,trx_type,trx_total)

    err = stub.PutState(customer_id, []byte(customer_id))
    if err != nil {
        return nil, err
    }

    err = stub.PutState(device_id, []byte(device_id))
    if err != nil {
        return nil, err
    }

    err = stub.PutState(trx_type, []byte(trx_type))
    if err != nil {
        return nil, err
    }

    err = stub.PutState(trx_total, []byte(trx_total))
    if err != nil {
        return nil, err
    }
    return nil, nil        
}

// Invoke
func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    var err error

    timestamp, err2 := stub.GetTxTimestamp()

    if err2 != nil {
        fmt.Println("Error getting transaction timestamp.", err2)
    }

    fmt.Println("Transaction Time: ", timestamp, customer_id, device_id, trx_type, trx_total)
    
    return nil, err

}

// Query
func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
    return nil, nil
}

// Main
func main() {
    err := shim.Start(new(SimpleChaincode))
    if err != nil {
        fmt.Printf("Error starting Simple chaincode: %s", err)
    }
}
