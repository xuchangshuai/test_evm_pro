pragma solidity >=0.7.0 <0.9.0;


contract OutputContract{
    
    fallback () payable external {}
    
    function getBalance() view public returns(uint balance){
        return address(this).balance;
    }
}