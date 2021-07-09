pragma solidity >=0.7.0 <0.9.0;

contract Test {

    function add(uint256 a, uint256 b) public pure returns (uint256){
        return a + b;
    }

    function sub(uint256 a, uint256 b) public pure returns (uint256) {
        return a - b;
    }

    function mul(uint256 a, uint256 b) public pure returns (uint256) {
        return a * b;
    }

    function div(uint256 a, uint256 b) public pure returns (uint256) {
        return a / b;
    }

    function sdiv(int256 a, int256 b) public pure returns (int256) {
        return a / b;
    }

    function mod(uint256 a, uint256 b) public pure returns (uint256) {
        return a % b;
    }

    function smod(int256 a, int256 b) public pure returns (int256) {
        return a % b;
    }


    function exp(uint256 a, uint256 b) public pure returns (uint256) {
        return a ** b;
    }

    function lt(uint256 a, uint256 b) public pure returns (bool){
        return a < b;
    }

    function gt(uint256 a, uint256 b) public pure returns (bool){
        return a > b;
    }

    function slt(int256 a, int256 b) public pure returns (bool){
        return a < b;
    }

    function sgt(int256 a, int256 b) public pure returns (bool){
        return a > b;
    }

    function eq(int256 a, int256 b) public pure returns (bool){
        return a == b;
    }

    function iszero(int256 a) public pure returns (bool){
        return a == 0;
    }

    function and(int256 a, int256 b) public pure returns (int256){
        return a & b;
    }

    function or(int256 a, int256 b) public pure returns (int256){
        return a | b;
    }

    function xor(int256 a, int256 b) public pure returns (int256){
        return a ^ b;
    }

    function not(int256 a) public pure returns (int256){
        return ~a;
    }

    function shl(uint256 a) public pure returns (uint256){
        return a << 2;
    }

    function shr(uint256 a) public pure returns (uint256){
        return a >> 2;
    }

    function sha(string memory a,string memory b) public pure returns(bool){
        return keccak256(abi.encode(a)) == keccak256(abi.encode(b));
    }

}









