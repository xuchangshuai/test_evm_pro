pragma solidity >=0.7.0 <0.9.0;

contract TestCoin {

    address public minter;

    // mapping(address=>uint256) public balances;
    constructor(){
        minter = msg.sender;
    }

    function getAddress() public view returns (address){
        return address(this);
    }

    function getBalance(address addr) public view returns (uint256){
        return address(addr).balance;
    }

    function getOrigin() public view returns (address){
        return tx.origin;
    }

    function getCaller() public view returns (address){
        return msg.sender;
    }

    function getValue() public payable returns (uint256){
        return msg.value;
    }

    function getData() public pure returns (bytes calldata){
        return msg.data;
    }

    function getCallDataCopy() public pure returns (bytes memory){
        bytes memory res = msg.data;
        return res;
    }

    function getGasPrice() public view returns (uint256){
        return tx.gasprice;
    }


    function getBlockHash(uint256 num) public view returns (bytes32){
        return blockhash(num);
    }


    function getCoinBase() public view returns (address){
        return block.coinbase;
    }

    function getTimestamp() public view returns (uint256){
        return block.timestamp;
    }


    function getBlockNumber() public view returns (uint256){
        return block.number;
    }


    function getBlockDifficulty() public view returns (uint256){
        return block.difficulty;
    }


    function pop() public pure {
        int256 a = 1;
        delete a;
    }

    function mload() public pure returns (int256){
        int256 a = 100;
        return a;
    }

    function mstore() public pure {
        int256 a = 100;
    }


}