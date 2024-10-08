// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "forge-std/Script.sol";
import "../src/KeyBroadcastContract.sol";
import "../src/KeyperSet.sol";
import "../src/KeyperSetManager.sol";
import "../src/Sequencer.sol";
import "../src/ValidatorRegistry.sol";

contract Deploy is Script {
    function run() external {
        uint256 deployKey = vm.envUint("DEPLOY_KEY");
        address deployerAddress = vm.addr(deployKey);
        console.log("Deployer:", deployerAddress);
        vm.startBroadcast(deployKey);
        deploySequencer();
        deployValidatorRegistry();
        vm.stopBroadcast();
    }

    function deploySequencer() public returns (Sequencer) {
        Sequencer s = new Sequencer();
        console.log("Sequencer:", address(s));
        return s;
    }

    function deployValidatorRegistry() public returns (ValidatorRegistry) {
        ValidatorRegistry vr = new ValidatorRegistry();
        console.log("ValidatorRegistry:", address(vr));
        return vr;
    }
}
