# Singular Example

> Native Go contract

This simple example is designed to show you how to read & write state variables.

The `Singular` smart contract supports two main methods:

* `get` - which returns the current value of the the singular cell (state variable).

* `set` - which takes a value as argument and overwrites the value in the cell

This is similar to any key/value store. Examples of key/value stores can be: etcd , Redis, memcached and more..

### How to test this example?

1. Deploy the singular contract onto the blockchain
   using the following command
   
        $ ./gamma-cli deploy Singular go/examples/singular/singular.go

   an example response from `gamma-cli` would be:
   
        {"TransactionReceipt":{"Txhash":"TUuY0i+omRLfHrBaay3PqOjdvwQDExUZlkrdhg2r+8U=","ExecutionResult":1,"OutputArguments":null},"TransactionStatus":1,"BlockHeight":136,"BlockTimestamp":0}
    
    an `ExecutionResult` with the value `1` means our contract has been deployed successfully. 

2. Get the singular cell value is possible using the following `gamma-cli` call

        $ ./gamma-cli run call go/examples/singular/jsons/get.json
        
   and the response should look similar to
   
        {"OutputArguments":[{"Name":"string","Type":"string","Value":"nil"}],"CallResult":1,"BlockHeight":187,"BlockTimestamp":0}
        
The `OutputArguments` field within the `JSON` response contains the return values, in our case a string value of `nil`
which is the initial value of our singular cell.

3. Setting a different value into the singular cell is achieved by using the `set` method from our newly deployed 
contract.

        ./gamma-cli run send go/examples/singular/jsons/set.json
        
and the response to our `set` call should like similar to

    {"TransactionReceipt":{"Txhash":"ZNfpEFDXLCaTvsRWIJaGpdb4NxdegNnsktzi3dtotHY=","ExecutionResult":1,"OutputArguments":null},"TransactionStatus":1,"BlockHeight":233,"BlockTimestamp":0}
    
Again, the most important part from this `JSON` as the `ExecutionResult` value which need to equate to `1` for our operation to be successful.
    
