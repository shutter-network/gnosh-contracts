// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/ISequencer.sol";
import "../src/Sequencer.sol";

contract SequencerTest is Test {
    Sequencer public sequencer;

    event TransactionSubmitted(
        uint64 eon,
        bytes32 identityPrefix,
        address sender,
        bytes encryptedTransaction,
        uint256 gasLimit
    );
    event DecryptionProgressSubmitted(bytes message);

    function setUp() public {
        sequencer = new Sequencer();
    }

    function testSubmitTransaction() public {
        uint64 eon = 5;
        bytes32 identityPrefix = hex"001122";
        address sender = makeAddr("sender");
        bytes memory encryptedTransaction = "aabbcc";
        uint256 gasLimit = 8;

        vm.expectEmit(address(sequencer));
        emit TransactionSubmitted(
            eon,
            identityPrefix,
            sender,
            encryptedTransaction,
            gasLimit
        );

        vm.fee(20);
        hoax(sender);
        sequencer.submitEncryptedTransaction{value: 160}(
            eon,
            identityPrefix,
            encryptedTransaction,
            gasLimit
        );
    }

    function testSubmitTransactionInsufficientFee() public {
        uint64 eon = 5;
        bytes32 identityPrefix = hex"001122";
        address sender = makeAddr("sender");
        bytes memory encryptedTransaction = "aabbcc";
        uint256 gasLimit = 8;

        vm.expectRevert(InsufficientFee.selector);

        vm.fee(20);
        hoax(sender);
        sequencer.submitEncryptedTransaction{value: 159}(
            eon,
            identityPrefix,
            encryptedTransaction,
            gasLimit
        );
    }

    function testDecryptionProgressSubmitted() public {
        bytes memory message = hex"aabbcc";
        vm.expectEmit(address(sequencer));
        emit DecryptionProgressSubmitted(message);
        sequencer.submitDecryptionProgress(message);
    }
}
