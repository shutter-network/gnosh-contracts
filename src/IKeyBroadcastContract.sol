// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

interface IKeyBroadcastContract {
    function broadcastEonKey(uint64 eon, bytes memory key) external;

    function getEonKey(uint64 eon) external view returns (bytes memory);

    event EonKeyBroadcast(uint64 eon, bytes key);
}
