pragma solidity >=0.7.0 <0.9.0;

contract InputContract{
    
    address payable public _outputAddress;
    
    
    constructor(address payable outputAddress)public{
        _outputAddress = outputAddress;
    }
    
    function transferAccounts() public payable returns(bool success){
        _outputAddress.transfer(msg.value-1000000000);
        return true;
    }
    
    function getBalance() view public returns(uint balance){
        return address(this).balance;
    }
}