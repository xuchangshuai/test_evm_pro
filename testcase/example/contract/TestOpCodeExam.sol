pragma solidity >=0.7.0 <0.9.0;

contract simple {

    uint num = 0;

    constructor(){
        num = 123;
    }

    function add(uint i) public payable returns (uint){
        uint m = 111;
        num = num * i + m;
        return num;
    }

}