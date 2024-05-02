#  (2024-05-02)


### Bug Fixes

* **app:** Fixed bech32 on account keeper to not be hardcoded  ([#54](https://github.com/dymensionxyz/rollapp-wasm/issues/54)) ([883653a](https://github.com/dymensionxyz/rollapp-wasm/commit/883653af7053450af80719e1cfd93e8309ba7a7d))
* merge conflict ([#13](https://github.com/dymensionxyz/rollapp-wasm/issues/13)) ([2cc8431](https://github.com/dymensionxyz/rollapp-wasm/commit/2cc8431a3dc57a60efece2a485c7298c08d22ecb))
* updated IBC Keeper to use the sequencer keeper. ([#63](https://github.com/dymensionxyz/rollapp-wasm/issues/63)) ([6c4a2b6](https://github.com/dymensionxyz/rollapp-wasm/commit/6c4a2b674527476ad08e790dfd4b41ef18f086e3))


### Features

* add hub genesis module ([#43](https://github.com/dymensionxyz/rollapp-wasm/issues/43)) ([73b3ceb](https://github.com/dymensionxyz/rollapp-wasm/commit/73b3cebef6c159494f0a4074ef5edb804b82bf0c))
* Add wasm module for rollapp-wasm ([#10](https://github.com/dymensionxyz/rollapp-wasm/issues/10)) ([9829d4a](https://github.com/dymensionxyz/rollapp-wasm/commit/9829d4a10b9f7928c98151b7295b20f0d54a8ad0))
* **app:** Add modules authz and feegrant ([#60](https://github.com/dymensionxyz/rollapp-wasm/issues/60)) ([a4451ea](https://github.com/dymensionxyz/rollapp-wasm/commit/a4451eaebd11eb49c89a40c239f6dd8593f201d1))
* **be:** integrate block explorer Json-RPC server ([#41](https://github.com/dymensionxyz/rollapp-wasm/issues/41)) ([51fd3e3](https://github.com/dymensionxyz/rollapp-wasm/commit/51fd3e36a0404d68325c64f79f65a15afc3be82a))
* **ci:** add auto update changelog workflow ([25a3ec8](https://github.com/dymensionxyz/rollapp-wasm/commit/25a3ec87506915de2330203bf48c340f3625d983))
* **ci:** Add auto update changelog workflow ([#61](https://github.com/dymensionxyz/rollapp-wasm/issues/61)) ([ed9c6da](https://github.com/dymensionxyz/rollapp-wasm/commit/ed9c6da98f33a9842ae83007b46bc074f67d2152))
* **ci:** update changelog workflow ([93199d6](https://github.com/dymensionxyz/rollapp-wasm/commit/93199d6063afef7c1a7b5431b461b8674cf71273))



## PR #64 - Feat(ci): fix changelog workflow
- ## Description

<!-- Add a description of the changes that this PR introduces and the files that
are the most critical to review.
-->

----

Closes #XXX

**All** items are required. Please add a note to the item if the item is not applicable and
please add links to any relevant follow-up issues.

PR review checkboxes:

I have...

- [ ]  Added a relevant changelog entry to the  section in 
- [ ]  Targeted PR against the correct branch
- [ ]  included the correct [type prefix](https://github.com/commitizen/conventional-commit-types/blob/v3.0.0/index.json) in the PR title
- [ ]  Linked to the GitHub issue with discussion and accepted design
- [ ]  Targets only one GitHub issue
- [ ]  Wrote unit and integration tests
- [ ]  Wrote relevant migration scripts if necessary
- [ ]  All CI checks have passed
- [ ]  Added relevant  [comments](https://blog.golang.org/godoc-documenting-go-code)
- [ ]  Updated the scripts for local run, e.g genesis_config_commands.sh if the PR changes parameters
- [ ]  Add an issue in the [e2e-tests repo](https://github.com/dymensionxyz/e2e-tests) if necessary

SDK Checklist
- [ ] Import/Export Genesis
- [ ] Registered Invariants
- [ ] Registered Events
- [ ] Updated openapi.yaml
- [ ] No usage of go 
- [ ] No usage of 
- [ ] Used fixed point arithmetic and not float arithmetic
- [ ] Avoid panicking in Begin/End block as much as possible
- [ ] No unexpected math Overflow
- [ ] Used  and not 
- [ ] Out-of-block compute is bounded
- [ ] No serialized ID at the end of store keys
- [ ] UInt to byte conversion should use BigEndian

Full security checklist [here](https://www.faulttolerant.xyz/2024-01-16-cosmos-security-1/)


----;

For Reviewer:

- [ ]  Confirmed the correct [type prefix](https://github.com/commitizen/conventional-commit-types/blob/v3.0.0/index.json) in the PR title
- [ ]  Reviewers assigned
- [ ]  Confirmed all author checklist items have been addressed

---;

After reviewer approval:

- [ ]  In case the PR targets the main branch, PR should not be squash merge in order to keep meaningful git history.
- [ ]  In case the PR targets a release branch, PR must be rebased.


### PR #64: Feat(ci): fix changelog workflow

