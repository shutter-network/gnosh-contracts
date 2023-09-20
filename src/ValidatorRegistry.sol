// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "src/IValidatorRegistry.sol";

contract ValidatorRegistry is IValidatorRegistry {
    function register(
        bytes memory registrationMessage,
        bytes memory registrationSignature
    ) external {
        emit Registration(registrationMessage, registrationSignature);
    }

    function deregister(
        bytes memory deregistrationMessage,
        bytes memory deregistrationSignature
    ) external {
        emit Deregistration(deregistrationMessage, deregistrationSignature);
    }
}
