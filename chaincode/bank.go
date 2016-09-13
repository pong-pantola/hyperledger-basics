package main

import (
  "errors"
  "fmt"
  "strconv"
  "github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

func main() {
  err := shim.Start(new(SimpleChaincode))
  if err != nil {
    fmt.Printf("Error starting Simple chaincode: %s", err)
  }
}


/*
Init is called when chaincode is deployed.
*/
func (t *SimpleChaincode) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

  if function == "createAccount" {

    //Function createAccount creates two accounts.

    var account1Name, account2Name string //name of account1 and account2
    var account1Balance, account2Balance int //balance of account1 and account2
    var err error

    if len(args) != 4 {
      return nil, errors.New("Function createAccount expects 4 arguments.")    
    }
 
    //initialize the name and balance of the 1st account
    account1Name = args[0]
    account1Balance, err = strconv.Atoi(args[1])
    if err != nil {
      return nil, errors.New("The argument args[1] of function createAccount is expected to be an integer value.")
    }

    //initialize the name and balance of the 2nd account
    account2Name = args[2]
    account2Balance, err = strconv.Atoi(args[3])
    if err != nil {
      return nil, errors.New("The argument args[3] of function createAccount is expected to be an integer value.")
    }


    fmt.Printf("Creating 1st account:\n")
    fmt.Printf("Name: %s\n", account1Name)
    fmt.Printf("Balance: %d\n", account1Balance)

    //save the balance of the 1st account to the ledger using its name as the key
    //Note: account1Balance is integer. It is converted to string.  String is converted to array of bytes.
    err = stub.PutState(account1Name, []byte(strconv.Itoa(account1Balance)))
    if err != nil {
      return nil, err
    }else{
      fmt.Printf("Balance of 1st Account successfully saved in the ledger.\n\n")
    }

    fmt.Printf("Creating 2nd account:\n")
    fmt.Printf("Name: %s\n", account2Name)
    fmt.Printf("Balance: %d\n", account2Balance)

    //save the balance of the 2nd account to the ledger using its name as the key
    //Note: account2Balance is integer. It is converted to string.  String is converted to array of bytes.
    err = stub.PutState(account2Name, []byte(strconv.Itoa(account2Balance)))
    if err != nil {
      return nil, err
    }else{
      fmt.Printf("Balance of 2nd Account successfully saved in the ledger.\n\n")
    }

    return nil, nil

  }
  
  return nil, errors.New("Unknown function " + function + ".")
}



func (t *SimpleChaincode) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

  if function == "getBalance" {

    //Function getBalance gets the balance of an account name.

    var accountName string //name of account to be queried
    var err error

    if len(args) != 1 {
      return nil, errors.New("Function getBalance expects 1 argument.")    
    }
 
    //initialize the name of account to be queried
    accountName = args[0]


    fmt.Printf("Getting the balance of Account %s\n", accountName)

    //get the balance of the account from the ledger using its name as the key
    var accountBalanceInArrBytes []byte //account balance in array of bytes
    accountBalanceInArrBytes, err = stub.GetState(accountName)

    if err != nil {
      return nil, err
    }else{
      fmt.Printf("Balance: %s\n", string(accountBalanceInArrBytes))
    }

    return accountBalanceInArrBytes, nil

  }

  return nil, errors.New("Unknown function " + function + ".")
}



func (t *SimpleChaincode) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {

  if function == "transferFunds" {

    //Function transferFunds transfer funds from one account to another.

    var srcAccountName, dstAccountName string //name of source account and destination account
    var amount int //amount to transfer
    var err error

    if len(args) != 3 {
      return nil, errors.New("Function transferFunds expects 3 arguments.")    
    }
 
    //initialize the name of the source account
    srcAccountName = args[0]

    //initialize the name of the destination account
    dstAccountName = args[1]

    //initialize the amount to transfer
    amount, err = strconv.Atoi(args[2])
    if err != nil {
      return nil, errors.New("The argument args[2] of function transferFunds is expected to be an integer value.")
    }


    //get the balance of the source account from the ledger using its name as the key
    var srcAccountBalanceInArrBytes []byte //balance of source account in array of bytes
    srcAccountBalanceInArrBytes, err = stub.GetState(srcAccountName)

    if err != nil {
      return nil, err
    }

    //convert array of bytes to string
    var srcAccountInString string //balance of source account in string
    srcAccountInString =  string(srcAccountBalanceInArrBytes)
    fmt.Printf("Balance of Source Account: %s\n", srcAccountInString)

    //convert string to integer
    var srcAccountBalance int //balance of source account in integer
    srcAccountBalance, _ = strconv.Atoi(srcAccountInString)



    //get the balance of the destination account from the ledger using its name as the key
    var dstAccountBalanceInArrBytes []byte //balance of destination account in array of bytes
    dstAccountBalanceInArrBytes, err = stub.GetState(dstAccountName)

    if err != nil {
      return nil, err
    }

    //convert array of bytes to string
    var dstAccountInString string //balance of destination account in string
    dstAccountInString =  string(dstAccountBalanceInArrBytes)
    fmt.Printf("Balance of Destination Account: %s\n", dstAccountInString)

    var dstAccountBalance int //balance of destination account in integer
    dstAccountBalance, _ = strconv.Atoi(dstAccountInString)


    //decrease the balance of the source account by the amount to transfer
    srcAccountBalance -= amount

    //increase the balance of destination account by the amount to transfer
    dstAccountBalance += amount


    //save the updated balance of the source account to the ledger using its name as the key
    //Note: srcAccountBalance is integer. It is converted to string.  String is converted to array of bytes.
    err = stub.PutState(srcAccountName, []byte(strconv.Itoa(srcAccountBalance)))
    if err != nil {
      return nil, err
    }else{
      fmt.Printf("Updated balance of Source Account successfully saved in the ledger.\n\n")
    }

    //save the updated balance of the destination account to the ledger using its name as the key
    //Note: dstAccountBalance is integer. It is converted to string.  String is converted to array of bytes.
    err = stub.PutState(dstAccountName, []byte(strconv.Itoa(dstAccountBalance)))
    if err != nil {
      return nil, err
    }else{
      fmt.Printf("Updated balance of Destination Account successfully saved in the ledger.\n\n")
    }

    return nil, nil

  }
  
  return nil, errors.New("Unknown function " + function + ".")
}

