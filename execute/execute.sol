pragma solidity ^0.4.24;

// The input could be byte code for another contract
// If we could deploy that contract from this contract
// then call some function, then destroy it, it would work
// The keyword new will not work because the solidity compiler
// needs to know a-priori what the abi is of the contract.

// However contract creation is merely sending data in a transaction to the zero
// address which when executed returns the bytecode of the contract, thus we should
// be able to send some contract deploy byte code to the 0x

// Just the interface of the deployed contract, required
interface Adder {
    function add(int a, int b) public returns (int);
}

interface StringReverse {
    function stringReverse(string input) public returns(string);
}

contract Tester {
    event TestPass(bool res);

    struct Submission {
        bool pass;  // pass/fail
        address submitter; // submitter address
    }

    uint public submissionCount;    
    // This declares a state variable that
    Submission[] public submissions;

    function create(bytes code) private returns (address addr){
        // Creates a contract based on code, 
        // returns address created there
        assembly {
            addr := create(0, add(code, 0x20), mload(code))
            //jumpi(invalidJumpLabel,iszero(extcodesize(addr)))
        }
        return addr;
    }

    function testAdder(bytes code) public {
        address deployed = create(code); 
        if (deployed == 0) throw;
        // Could be some kind of error if they bytecode
        // the user passed in does not adhere to the
        // interface of Adder
        Adder deployedAdder = Adder(deployed);
        submissionCount += 1;
        if (deployedAdder.add(7, 3) == 10 && deployedAdder.add(-3, 3) == 0 && deployedAdder.add(-3, -7) == -10) {
            emit TestPass(true);
            submissions.push(Submission(true, msg.sender));
        } else {
            emit TestPass(false);
            submissions.push(Submission(false, msg.sender));
        }
    }

    function testStringReverse(bytes code) public {
        address deployed = create(code); 
        if (deployed == 0) throw;
        StringReverse deployedStrRev = StringReverse(deployed);
        // TODO: add more extensive tests
        submissionCount += 1;
        if (keccak256(deployedStrRev.stringReverse("abcdef")) == keccak256("fedcba")) {
            emit TestPass(true);
            submissions.push(Submission(true, msg.sender));
        } else {
            emit TestPass(false);
            submissions.push(Submission(false, msg.sender));
        }
    }
}
