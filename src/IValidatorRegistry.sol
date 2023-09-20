// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IValidatorRegistry {
    function register(
        bytes memory registrationMessage,
        bytes memory registrationSignature
    ) external;

    function deregister(
        bytes memory deregistrationMessage,
        bytes memory deregistrationSignature
    ) external;

    event Registration(bytes registrationMessage, bytes registrationSignature);
    event Deregistration(
        bytes deregistrationMessage,
        bytes deregistrationSignature
    );
}
