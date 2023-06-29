// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "src/IKeyBroadcastContract.sol";
import "src/IKeyperSetManager.sol";

contract KeyBroadcastContract is IKeyBroadcastContract {
    mapping(uint64 => bytes) private keys;
    address private keyperSetManagerAddress;

    constructor(address _keyperSetManagerAddress) {
        keyperSetManagerAddress = _keyperSetManagerAddress;
    }

    function broadcastEonKey(uint64 eon, bytes memory key) external {
        if (key.length == 0) {
            revert InvalidKey();
        }
        if (keys[eon].length > 0) {
            revert AlreadyHaveKey();
        }
        if (
            msg.sender !=
            IKeyperSetManager(keyperSetManagerAddress).getKeyperSetAddress(eon)
        ) {
            revert NotAllowed();
        }

        keys[eon] = key;
        emit EonKeyBroadcast(eon, key);
    }

    function getEonKey(uint64 eon) external view returns (bytes memory) {
        return keys[eon];
    }
}
