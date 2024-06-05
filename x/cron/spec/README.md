<!--
order: 0
title: Cron Overview
parent:
  title: "cron"
-->

# `cron`

## Contents

## Abstract

`x/cron` is an implementation of a Cosmos SDK module,

This document specifies the internal `x/cron` module on the network.

The `x/cron` module provides functionality for scheduling and executing tasks, including executing sudo contract calls during specific phases, such as begin blockers. By integrating scheduled contract executions, `x/cron` enhances the functionality of smart contracts, ensuring critical operations are performed automatically and reliably.

1. **[Concepts](01_concepts.md)**
2. **[State](02_state.md)**
3. **[Clients](03_clients.md)**

## Flow

### Step 1: Upload the contract
  #### ex: `wasmrollappd tx wasm store "$DIR/../../x/wasm/keeper/testdata/coinflip.wasm`
### Step 2: Register the contract according to the game type
  #### ex: `wasmrollappd tx cron register-contract coin-flip rol14hj2tavq8f.... 1`
### result : cronjob for the specific contract is activated