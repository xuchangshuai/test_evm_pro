pragma solidity >=0.7.0 <0.9.0;

contract Test {

    function add() public pure {
        uint256 a = 2;
        uint256 b = 1;
        a + b;
    }

    function sub() public pure {
        uint256 a = 2;
        uint256 b = 1;
        a - b;
    }

    function mul() public pure {
        uint256 a = 2;
        uint256 b = 1;
        a * b;
    }

    function div() public pure {
        uint256 a = 2;
        uint256 b = 1;
        a / b;
    }

    function sdiv() public pure {
        int256 a = 2;
        int256 b = 1;
        a / b;
    }

    function mod() public pure {
        uint256 a = 2;
        uint256 b = 1;
        a % b;
    }


    function smod() public pure {
        int256 a = 2;
        int256 b = 1;
        a % b;
    }

    function exp() public pure {
        uint256 a = 2;
        uint256 b = 1;
        a ** b;
    }

    function and() public pure {
        int256 a = 2;
        int256 b = 1;
        int c;
        c = a & b;
    }

    function or() public pure {
        int256 a = 2;
        int256 b = 1;
        int256 c;
        c = a | b;
    }

    function xor() public pure {
        int256 a = 2;
        int256 b = 6;
        int256 c;
        c = a ^ b;
    }

    function not() public pure {
        int256 a = 2;
        int256 c;
        c = ~a;
    }

    function byteFun() public pure {
        bytes32 a = "0x1";
        bytes32 c;
        c = a[1];
    }

    function lt() public pure {
        uint256 a = 2;
        uint256 b = 2;
        bool c = a < b;
    }

    function gt() public pure {
        uint256 a = 2;
        uint256 b = 2;
        bool c = a > b;
    }

    function slt() public pure {
        int256 a = 2;
        int256 b = 2;
        bool c = a < b;
    }

    function sgt() public pure {
        int256 a = 2;
        int256 b = 2;
        bool c = a > b;
    }

    function eq() public pure {
        int256 a = 2;
        int256 b = 2;
        bool c = a == b;
    }

}









