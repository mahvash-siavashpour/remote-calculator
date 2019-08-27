# Remote Calculator

This is a remote calculator in golang that uses a client server connection.

## File Description

There are two versions of the client and server codes. HTTP and the old version which transfers data via buffer reader and writter.

### How to use Old Version

On the client side you have to enter the operands and the operator in the order that is shown on the input line.And server sends back the result to client.

### How to use HTTP Version

This is the format of the json that client sends to the server :
```
Operand1 : number
Operator : sign
Operand2 : number
```

And the server sends back the result in a single string.

