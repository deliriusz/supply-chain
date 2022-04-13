// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

library ProductLib {
   enum LifecycleState {
      InProduction,
      Created,
      Payed,
      Delivered
   }

   struct Product {
      LifecycleState state;
      string name;
      uint initialPrice;
   }
}