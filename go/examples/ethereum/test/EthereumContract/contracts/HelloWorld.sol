pragma solidity >=0.4.21 <0.6.0;

contract HelloWorld {
    event HelloSaid(bytes10 name, uint256 amount);

    function sayHello() public pure returns (string memory) {
        return("hello world");
    }

    function emitHello() public {
        emit HelloSaid("John Snow", 3850);
    }
}
