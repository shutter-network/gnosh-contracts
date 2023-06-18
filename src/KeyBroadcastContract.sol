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
        require(key.length > 0, "key is empty");
        require(keys[eon].length == 0, "key already broadcast for this eon");
        require(
            msg.sender ==
                IKeyperSetManager(keyperSetManagerAddress).getKeyperSetAddress(
                    eon
                ),
            "not allowed to broadcast eon key"
        );

        keys[eon] = key;
        emit EonKeyBroadcast(eon, key);
    }

    function getEonKey(uint64 eon) external view returns (bytes memory) {
        return keys[eon];
    }
}
