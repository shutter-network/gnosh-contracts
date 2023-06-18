// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Test.sol";
import "../src/IKeyBroadcastContract.sol";
import "../src/KeyBroadcastContract.sol";
import "../src/KeyperSetManager.sol";
import "../src/KeyperSet.sol";

contract KeyBroadcastTest is Test {
    KeyBroadcastContract public keyBroadcastContract;
    KeyperSetManager public keyperSetManager;
    KeyperSet public keyperSet;

    event EonKeyBroadcast(uint64 eon, bytes key);

    function setUp() public {
        keyperSetManager = new KeyperSetManager();
        keyBroadcastContract = new KeyBroadcastContract(
            address(keyperSetManager)
        );
        keyperSet = new KeyperSet();
        keyperSet.setFinalized();
        keyperSetManager.addKeyperSet(100, address(keyperSet));

        keyperSet = new KeyperSet();
        keyperSet.setFinalized();
        keyperSetManager.addKeyperSet(200, address(keyperSet));
    }

    function testBroadcastEonKeyEmpty() public {
        vm.expectRevert("key is empty");
        bytes memory key = bytes("");
        vm.prank(address(keyperSet));
        keyBroadcastContract.broadcastEonKey(1, key);
    }

    function testBroadcastEonKeyNotAllowed() public {
        vm.expectRevert("not allowed to broadcast eon key");
        bytes memory key = bytes("foo bar");
        vm.prank(address(keyperSet));
        keyBroadcastContract.broadcastEonKey(0, key);
    }

    function testBroadcastEonKeyDuplicate() public {
        bytes memory key = bytes("foo bar");
        vm.prank(address(keyperSet));
        keyBroadcastContract.broadcastEonKey(1, key);

        vm.expectRevert("key already broadcast for this eon");
        vm.prank(address(keyperSet));
        keyBroadcastContract.broadcastEonKey(1, key);
    }

    function testBroadcastEonKeyEmitsEvent() public {
        vm.expectEmit(address(keyBroadcastContract));
        bytes memory key = bytes("foo bar");
        emit EonKeyBroadcast(1, key);
        vm.prank(address(keyperSet));
        keyBroadcastContract.broadcastEonKey(1, key);
    }

    function testGetEonKey() public {
        assertEq(keyBroadcastContract.getEonKey(1), bytes(""));
        vm.prank(address(keyperSet));
        keyBroadcastContract.broadcastEonKey(1, bytes("foo bar"));
        assertEq(keyBroadcastContract.getEonKey(1), bytes("foo bar"));
    }
}
