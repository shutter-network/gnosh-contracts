// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/IValidatorRegistry.sol";
import "../src/ValidatorRegistry.sol";

contract ValidatorRegistryTest is Test {
    event Registration(bytes registrationMessage, bytes registrationSignature);
    event Deregistration(
        bytes deregistrationMessage,
        bytes deregistrationSignature
    );

    IValidatorRegistry registry;

    function setUp() public {
        registry = new ValidatorRegistry();
    }

    function testRegistration() public {
        bytes memory message = bytes("msg");
        bytes memory signature = bytes("sig");

        vm.expectEmit(address(registry));
        emit Registration(message, signature);

        registry.register(message, signature);
    }

    function testDeregistration() public {
        bytes memory message = bytes("msg");
        bytes memory signature = bytes("sig");

        vm.expectEmit(address(registry));
        emit Deregistration(message, signature);

        registry.deregister(message, signature);
    }
}
