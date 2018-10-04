Decentralized Hackerrank:
- Submitting contract byte code to a Tester contract which validates whether
your code passes the tests 
- You can build reputation by solving these challenges
- Your code needs to adhere to the interface of the challege

Example challenge:

Write an implementation of an adder function according to this interface:

contract Adder {
    function add(uint a, uint b) returns (uint){}
}

Then you send a transaction with the compiled version of your implementation
and we execute it. You receive pass/fail as a boolean from a TestPass event.

TODO:
- Get this working with ganach-cli, then try writing a simple web server where someone can enter 
their code in the browser and the webserver will send their code in for testing, would be cool if
it could display all the results of people current reputations (based on the challenges they have passed)
    - On second thought maybe I don't even need ganach-cli, I can just have a webserver which is manipulating
a simulated backend

- Figure out how to self destruct the submission contract after it gets tested so
we dont just fill up the chain with submission contracts
- Reputation system
- A feature could be users can actually create their own challenges
- Minimal UI to submit the code




