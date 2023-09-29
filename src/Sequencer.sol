// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "src/ISequencer.sol";

contract Sequencer is ISequencer {
    function submitEncryptedTransaction(
        uint64 eon,
        bytes32 identityPrefix,
        bytes memory encryptedTransaction,
        uint256 gasLimit
    ) external payable {
        if (msg.value < block.basefee * gasLimit) {
            revert InsufficientFee();
        }

        emit TransactionSubmitted(
            eon,
            identityPrefix,
            msg.sender,
            encryptedTransaction,
            gasLimit
        );
    }

    function submitDecryptionProgress(bytes memory message) external {
        emit DecryptionProgressSubmitted(message);
    }
}
