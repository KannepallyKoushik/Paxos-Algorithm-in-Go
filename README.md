# Paxos-Algorithm-in-Go

## Abstract: what is Paxos protocol ?
    1. Paxos is a Distributed Consensus Algorithm . It operates as a network of logical processes , of Different types.
    2.A group of machines (let’s have in mind a server farm) must agree on a value.The proposers will try to propose value. 
      The protocol must ensure that at most one value is selected.

## Working of Paxos
Phase 1. 

- A proposer selects a proposal number n and sends a prepare request with number n to a majority of acceptors.
- If an acceptor receives a prepare request with number n greater than that of any prepare request to which it has already responded, 
  then it responds to the request with a promise not to accept any more proposals numbered less than n and with the highest-numbered proposal (if any) 
  that it has accepted.
  
Phase 2. 

-  If the proposer receives a response to its prepare requests (mumbered n) from a majority of acceptors, then it sends an accept 
   request to each of those arepturs for a proposal numbered with a value u, where e is the value of the highest-numbered proposal among the responses,
   or is any value if the responses reported no proposals.
- If an acceptor receives an accept request for a proposal numberd n ,it accepts the proposal unless it has already responded to a prepare 
  request having a number greater than n.
  
There are also learners, who checks if a quorum (lets say a majority) of accepters have accepted a same proposal.
  
![Paxos](paxos-understand.png)

## Execution
1. Open 6 Terminals(command prompt) and move to the directory where we have **server.go** files (Or place each server.go file in distinct computers which are in same network)
2. Run each server.go file is seperate terminals immidiately one after the other(We have 5 server.go files with numbering as such)
3. Now in 6th Terminal traverse to client folder and run client.go file
```command to run : go run filename.go```
4. It is a menu driven application where We have to send request through client terminal and Consensus Algorithm starts working
5. Enter the Server Number to which you want to send request to (Ranges from 1-5)
