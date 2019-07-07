/**
 * Copyright 2019 the orbs-contract-sdk authors
 * This file is part of the orbs-contract-sdk library in the Orbs project.
 *
 * This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
 * The above notice should be included in all copies or substantial portions of the software.
 */

module.exports = async function(done) {
  try {
    const HelloWorld = artifacts.require('./contracts/HelloWorld.sol');
    const instance = await HelloWorld.new();
    const output = {
      "ContractAddress": instance.address,
      "ContractJsonAbi": JSON.stringify(instance.abi)
    };
    console.log(JSON.stringify(output, null, 2));
  } catch(e) {
    console.error(e);
  }
  done();
};