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
contract Adder {
    function add(uint a, uint b) returns (uint){}
}

contract Tester {
    event TestPass(bool res);

    function create(bytes code) private returns (address addr){
        // Creates a contract based on code, 
        // returns address created there
        assembly {
            addr := create(0, add(code, 0x20), mload(code))
            //jumpi(invalidJumpLabel,iszero(extcodesize(addr)))
        }
    }

    function test(bytes code) public {
        // Create the add contract
// keccak("test(bytes)") 2f570a234f56174a0be5cf2fff788ff394b02e8140a68e91a993c49f6c1e0219
        address deployed = create(code); 
        if(deployed == 0) throw;
//         // Could be some kind of error if they bytecode
//         // the user passed in does not adhere to the
//         // interface of Adder
        Adder deployedAdder = Adder(deployed);
        // Call add and return t/f if tests pass
        if (deployedAdder.add(10, 10) == 20) {
            emit TestPass(true);
        } else {
            emit TestPass(false);
        }
    }
}
