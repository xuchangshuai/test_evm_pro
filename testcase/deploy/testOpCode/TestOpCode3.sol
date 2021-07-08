pragma solidity >=0.7.0 <0.9.0;

contract testStorege {

    int256 a = 1;

    function sloadAndSstore() public view {
        int256 b;
        b = a;
    }

    function jump() public pure {
        int256 a;
        for (int256 i = 0; i <= 3; i++) {
            if (i == 2) {
                a = i;
            }

        }
    }


    function push() public pure {
        uint8 c1 = 1;
        uint16 c2 = 1;
        uint24 c3 = 1;
        uint32 c4 = 1;
        uint40 c5 = 1;
        uint48 c6 = 1;
        uint56 c7 = 1;
        uint64 c8 = 1;
        uint72 c9 = 1;
        uint80 c10 = 1;
        uint88 c11 = 1;
        uint96 c12 = 1;
    }


    function dup() public pure {
        uint8 c1 = 1;
        uint16 c2 = c1;
        uint24 c3 = c2;
        uint32 c4 = c3;
        uint40 c5 = c4;
        uint48 c6 = c5;
        uint56 c7 = c6;
        uint64 c8 = c7;
        uint72 c9 = c8;
        uint80 c10 = c9;
        uint88 c11 = c10;
        uint96 c12 = c11;
    }

    function swap() public pure {
        uint8 c1 = 1;
        uint8 c2 = 8;
        c1;
        c2 == c2;
        c1;
    }

}