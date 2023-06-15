// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;

import "openzeppelin/contracts/access/Ownable.sol";
import "src/IKeyperSet.sol";

contract KeyperSet is IKeyperSet, Ownable {
    bool finalized;
    address[] members;
    uint64 threshold;

    function isFinalized() external view returns (bool) {
        return finalized;
    }

    function getNumMembers() external view returns (uint64) {
        return uint64(members.length);
    }

    function getMember(uint64 index) external view returns (address) {
        return members[index];
    }

    function getMembers() external view returns (address[] memory) {
        return members;
    }

    function getThreshold() external view returns (uint64) {
        return threshold;
    }

    function addMembers(address[] calldata newMembers) public onlyOwner {
        require(!finalized, "cannot add members to finalized keyper set");
        for (uint64 j = 0; j < newMembers.length; j++) {
            members.push(newMembers[j]);
        }
    }

    function setThreshold(uint64 _threshold) public onlyOwner {
        require(!finalized, "cannot set threshold of finalized keyper set");
        threshold = _threshold;
    }

    function setFinalized() public onlyOwner {
        finalized = true;
    }
}
