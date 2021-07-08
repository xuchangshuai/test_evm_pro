pragma solidity >=0.7.0 <0.9.0;

contract CallTest {

    function callByFun(address addr) public returns (bool, bytes memory){
        bytes4 methodId = bytes4(keccak256("increaseAge(uint256)"));
        return addr.call(abi.encode(methodId, 1));
    }
}