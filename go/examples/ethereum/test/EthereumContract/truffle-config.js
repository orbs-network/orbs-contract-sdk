/**
 * Copyright 2019 the orbs-contract-sdk authors
 * This file is part of the orbs-contract-sdk library in the Orbs project.
 *
 * This source code is licensed under the MIT license found in the LICENSE file in the root directory of this source tree.
 * The above notice should be included in all copies or substantial portions of the software.
 */

module.exports = {
  networks: {
    ganache: {
      host: "127.0.0.1",
      port: 7545,
      network_id: "5777"
    },
  },
  mocha: {
  },
  compilers: {
    solc: {
    }
  }
};