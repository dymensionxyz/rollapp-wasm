#  (2024-10-29)


### Bug Fixes

* **app:** Fixed bech32 on account keeper to not be hardcoded  ([#54](https://github.com/dymensionxyz/rollapp-wasm/issues/54)) ([883653a](https://github.com/dymensionxyz/rollapp-wasm/commit/883653af7053450af80719e1cfd93e8309ba7a7d))
* **deps:** bumped dymint to `d51b961e7` to fix stuck submission bug ([#130](https://github.com/dymensionxyz/rollapp-wasm/issues/130)) ([4acf9e8](https://github.com/dymensionxyz/rollapp-wasm/commit/4acf9e80f1b1189a89dc1b39acb5706418a2157b))
* **deps:** bumped rdk to 3fe31b2db to fix denom-metadata transfer on unrelayed packets ([#131](https://github.com/dymensionxyz/rollapp-wasm/issues/131)) ([42ae0cf](https://github.com/dymensionxyz/rollapp-wasm/commit/42ae0cffb427defee32392663599bf1d2c34a482))
* **deps:** dymint bump to fix da grpc issue ([#136](https://github.com/dymensionxyz/rollapp-wasm/issues/136)) ([d6e15ba](https://github.com/dymensionxyz/rollapp-wasm/commit/d6e15ba4b90e6530cabf3dbd9c825b38ac0f6f2b))
* **deps:** updated dymint and rdk by removing the replace ([#132](https://github.com/dymensionxyz/rollapp-wasm/issues/132)) ([a34bc94](https://github.com/dymensionxyz/rollapp-wasm/commit/a34bc942d86d658a11038c69e860c973e96a1053))
* **ethsecp256k1:** register missing codecs for supporting ethsecp256k1 algo ([#138](https://github.com/dymensionxyz/rollapp-wasm/issues/138)) ([0f65b88](https://github.com/dymensionxyz/rollapp-wasm/commit/0f65b88b33ed669aa47bbb7183a62eb8e4068803))
* failing tests ([#149](https://github.com/dymensionxyz/rollapp-wasm/issues/149)) ([9044cc8](https://github.com/dymensionxyz/rollapp-wasm/commit/9044cc886d9d5813c6672e2a5a53119f148f0325))
* merge conflict ([#13](https://github.com/dymensionxyz/rollapp-wasm/issues/13)) ([2cc8431](https://github.com/dymensionxyz/rollapp-wasm/commit/2cc8431a3dc57a60efece2a485c7298c08d22ecb))
* removed MsgSend from allowed consensus msgs ([#146](https://github.com/dymensionxyz/rollapp-wasm/issues/146)) ([21753d2](https://github.com/dymensionxyz/rollapp-wasm/commit/21753d26066ec4320e892aeef577c2899a8c756d))
* **scripts:** fix init script to support mock backends ([#104](https://github.com/dymensionxyz/rollapp-wasm/issues/104)) ([65a57ca](https://github.com/dymensionxyz/rollapp-wasm/commit/65a57ca2b00141485ce7e32ab82d6a965b9d336b))
* updated IBC Keeper to use the sequencer keeper. ([#63](https://github.com/dymensionxyz/rollapp-wasm/issues/63)) ([6c4a2b6](https://github.com/dymensionxyz/rollapp-wasm/commit/6c4a2b674527476ad08e790dfd4b41ef18f086e3))


### Features

* add ethsecp256k1 as the default signing algo ([#80](https://github.com/dymensionxyz/rollapp-wasm/issues/80)) ([7362c6f](https://github.com/dymensionxyz/rollapp-wasm/commit/7362c6f89ba701d3103a5c25bbe45f01de0321f6))
* add hub genesis module ([#43](https://github.com/dymensionxyz/rollapp-wasm/issues/43)) ([73b3ceb](https://github.com/dymensionxyz/rollapp-wasm/commit/73b3cebef6c159494f0a4074ef5edb804b82bf0c))
* Add wasm module for rollapp-wasm ([#10](https://github.com/dymensionxyz/rollapp-wasm/issues/10)) ([9829d4a](https://github.com/dymensionxyz/rollapp-wasm/commit/9829d4a10b9f7928c98151b7295b20f0d54a8ad0))
* **ante:** skip fees for IBC messages  ([#127](https://github.com/dymensionxyz/rollapp-wasm/issues/127)) ([ff8e895](https://github.com/dymensionxyz/rollapp-wasm/commit/ff8e895578215eb62acb6582bfe0a0b8902326f6))
* **ante:** whitelisted relayers ([#148](https://github.com/dymensionxyz/rollapp-wasm/issues/148)) ([d815566](https://github.com/dymensionxyz/rollapp-wasm/commit/d81556668196e2c9ac133d9c8f32298e19c14afe))
* **app:** add denommetadata IBC middleware to app ([#94](https://github.com/dymensionxyz/rollapp-wasm/issues/94)) ([9a4756e](https://github.com/dymensionxyz/rollapp-wasm/commit/9a4756e0cd12bd2faa43531377ff51c15f5ce58f))
* **app:** Add modules authz and feegrant ([#60](https://github.com/dymensionxyz/rollapp-wasm/issues/60)) ([a4451ea](https://github.com/dymensionxyz/rollapp-wasm/commit/a4451eaebd11eb49c89a40c239f6dd8593f201d1))
* **be:** integrate block explorer Json-RPC server ([#41](https://github.com/dymensionxyz/rollapp-wasm/issues/41)) ([51fd3e3](https://github.com/dymensionxyz/rollapp-wasm/commit/51fd3e36a0404d68325c64f79f65a15afc3be82a))
* **callback:** add callback and cwerrors module for scheduled execution of wasm contracts ([#99](https://github.com/dymensionxyz/rollapp-wasm/issues/99)) ([7936ae2](https://github.com/dymensionxyz/rollapp-wasm/commit/7936ae2bfc57ac138989ae39eb968d3b07916bf1))
* **ci:** Add auto update changelog workflow ([#61](https://github.com/dymensionxyz/rollapp-wasm/issues/61)) ([ed9c6da](https://github.com/dymensionxyz/rollapp-wasm/commit/ed9c6da98f33a9842ae83007b46bc074f67d2152))
* **ci:** Add setup script and push hook ([#86](https://github.com/dymensionxyz/rollapp-wasm/issues/86)) ([d4dc3e4](https://github.com/dymensionxyz/rollapp-wasm/commit/d4dc3e4d73a72ab0e99cefc79c82eb0dcd79b187))
* consensus messages ([#141](https://github.com/dymensionxyz/rollapp-wasm/issues/141)) ([692fb70](https://github.com/dymensionxyz/rollapp-wasm/commit/692fb7096d6a7cb73ace726d1cddb5c276f02de5))
* **gasless:** wire the `x/gasless` module ([#102](https://github.com/dymensionxyz/rollapp-wasm/issues/102)) ([823dee3](https://github.com/dymensionxyz/rollapp-wasm/commit/823dee3cf52f205ccff47aa94e5669aa13c2ff9c))
* **genesis transfers:** wires genesis transfers ([#97](https://github.com/dymensionxyz/rollapp-wasm/issues/97)) ([a00a8c6](https://github.com/dymensionxyz/rollapp-wasm/commit/a00a8c6d96668bf917c2ca7a0597b1b62ad7a3a7))
* **genesis_bridge:** update to use new genesis bridge ([#143](https://github.com/dymensionxyz/rollapp-wasm/issues/143)) ([97b6b38](https://github.com/dymensionxyz/rollapp-wasm/commit/97b6b38240b2b234ed2fda9137f5c1d8001199b1))
* **sequencers:** wired sequencer rewards module. ([#119](https://github.com/dymensionxyz/rollapp-wasm/issues/119)) ([b6f0165](https://github.com/dymensionxyz/rollapp-wasm/commit/b6f01657c38dc47233515ac4f100213afce25028))
* set bech32 prefix without changing source code ([#68](https://github.com/dymensionxyz/rollapp-wasm/issues/68)) ([82c81a2](https://github.com/dymensionxyz/rollapp-wasm/commit/82c81a2e521669e2f0f48f34c9c8d56ed46d4196))
* wire `x/time-upgrade` module ([#125](https://github.com/dymensionxyz/rollapp-wasm/issues/125)) ([b333cb8](https://github.com/dymensionxyz/rollapp-wasm/commit/b333cb8a57d65e3524a5117e355bbb03aa4f1838))
* wire rollapp params module  ([#122](https://github.com/dymensionxyz/rollapp-wasm/issues/122)) ([7bb47e8](https://github.com/dymensionxyz/rollapp-wasm/commit/7bb47e8c23351d13ab91d6292d71e2c1bc1ae590))



#  (2024-05-20)


### Bug Fixes

* **app:** Fixed bech32 on account keeper to not be hardcoded  ([#54](https://github.com/dymensionxyz/rollapp-wasm/issues/54)) ([883653a](https://github.com/dymensionxyz/rollapp-wasm/commit/883653af7053450af80719e1cfd93e8309ba7a7d))
* merge conflict ([#13](https://github.com/dymensionxyz/rollapp-wasm/issues/13)) ([2cc8431](https://github.com/dymensionxyz/rollapp-wasm/commit/2cc8431a3dc57a60efece2a485c7298c08d22ecb))
* updated IBC Keeper to use the sequencer keeper. ([#63](https://github.com/dymensionxyz/rollapp-wasm/issues/63)) ([6c4a2b6](https://github.com/dymensionxyz/rollapp-wasm/commit/6c4a2b674527476ad08e790dfd4b41ef18f086e3))


### Features

* add ethsecp256k1 as the default signing algo ([#80](https://github.com/dymensionxyz/rollapp-wasm/issues/80)) ([7362c6f](https://github.com/dymensionxyz/rollapp-wasm/commit/7362c6f89ba701d3103a5c25bbe45f01de0321f6))
* add hub genesis module ([#43](https://github.com/dymensionxyz/rollapp-wasm/issues/43)) ([73b3ceb](https://github.com/dymensionxyz/rollapp-wasm/commit/73b3cebef6c159494f0a4074ef5edb804b82bf0c))
* Add wasm module for rollapp-wasm ([#10](https://github.com/dymensionxyz/rollapp-wasm/issues/10)) ([9829d4a](https://github.com/dymensionxyz/rollapp-wasm/commit/9829d4a10b9f7928c98151b7295b20f0d54a8ad0))
* **app:** Add modules authz and feegrant ([#60](https://github.com/dymensionxyz/rollapp-wasm/issues/60)) ([a4451ea](https://github.com/dymensionxyz/rollapp-wasm/commit/a4451eaebd11eb49c89a40c239f6dd8593f201d1))
* **be:** integrate block explorer Json-RPC server ([#41](https://github.com/dymensionxyz/rollapp-wasm/issues/41)) ([51fd3e3](https://github.com/dymensionxyz/rollapp-wasm/commit/51fd3e36a0404d68325c64f79f65a15afc3be82a))
* **ci:** add auto update changelog workflow ([5bc7247](https://github.com/dymensionxyz/rollapp-wasm/commit/5bc7247f4ecd073f9410024a7ce0944c126b1aaa))
* **ci:** Add auto update changelog workflow ([#61](https://github.com/dymensionxyz/rollapp-wasm/issues/61)) ([ed9c6da](https://github.com/dymensionxyz/rollapp-wasm/commit/ed9c6da98f33a9842ae83007b46bc074f67d2152))
* **ci:** Add setup script and push hook ([#86](https://github.com/dymensionxyz/rollapp-wasm/issues/86)) ([d4dc3e4](https://github.com/dymensionxyz/rollapp-wasm/commit/d4dc3e4d73a72ab0e99cefc79c82eb0dcd79b187))
* set bech32 prefix without changing source code ([#68](https://github.com/dymensionxyz/rollapp-wasm/issues/68)) ([82c81a2](https://github.com/dymensionxyz/rollapp-wasm/commit/82c81a2e521669e2f0f48f34c9c8d56ed46d4196))



