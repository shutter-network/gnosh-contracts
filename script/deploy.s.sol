// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/KeyBroadcastContract.sol";
import "../src/KeyperSet.sol";
import "../src/KeyperSetManager.sol";
import "../src/Sequencer.sol";
import "../src/ValidatorRegistry.sol";

contract DeployAll is Script {
    KeyBroadcastContract public keyBroadcastContract;
    KeyperSet public keyperSet;
    KeyperSetManager public keyperSetManager;
    Sequencer public sequencer;
    ValidatorRegistry public validatorRegistry;

    function deployKeyperSet() public {
        address broadcaster = vm.envAddress("BROADCASTER");
        keyperSet = new KeyperSet();
        address[] memory members = new address[](3);
        members[0] = address(1);
        members[1] = address(2);
        members[2] = address(3);
        keyperSet.addMembers(members);
        keyperSet.setKeyBroadcaster(broadcaster);
        keyperSet.setFinalized();
    }

    function deployKeyperSetManager() public {
        keyperSetManager = new KeyperSetManager();
        keyperSetManager.addKeyperSet(0, address(keyperSet));
    }

    function deployKeyBroadcastContract() public {
        keyBroadcastContract = new KeyBroadcastContract(
            address(keyperSetManager)
        );
    }

    function deploySequencer() public {
        sequencer = new Sequencer();
    }

    function deployValidatorRegistry() public {
        validatorRegistry = new ValidatorRegistry();
    }

    function run() external {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);
        deployKeyperSet();
        deployKeyperSetManager();
        deployKeyBroadcastContract();
        deploySequencer();
        deployValidatorRegistry();
        vm.stopBroadcast();
    }
}
