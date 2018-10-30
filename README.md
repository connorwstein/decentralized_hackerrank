Decentralized Hackerrank:
- Submitting contract byte code to a Tester contract which validates whether
your code passes the tests 
- You can build reputation by solving these challenges
- Your code needs to adhere to the interface of the challege
- The incentive structure: you pay per submission and then the winner gets the total
of all those submissions

Example challenge:

Write an implementation of an adder function according to this interface:

contract Adder {
    function add(uint a, uint b) returns (uint){}
}

Then you send a transaction with the compiled version of your implementation
and we execute it. You receive pass/fail as a boolean from a TestPass event.

TODO:
- Makefile
- More challenges, better UI
- Come up with a good testing strategy for the web application part 
- Resolve front running with hash based reservation
- Figure out how to self destruct the submission contract after it gets tested so
we dont just fill up the chain with submission contracts
- Reputation system
- A feature could be users can actually create their own challenges

Notes
- Obviously creating an entire contract for each submission is not scalable - some way to make use of plasma here? 
- ganache-cli does not seem to be able to handle selfdestructing a contract created with 
the create assembly code. Exact same code works fine in the remix IDE, not sure why.
