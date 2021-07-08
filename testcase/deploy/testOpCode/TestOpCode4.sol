pragma solidity >=0.7.0 <0.9.0;

contract Person {
    uint public age = 10;


    function increaseAge(uint num) public returns (uint){
        return ++age;
    }

    function getAge() public returns (uint){
        return age;
    }

}