#  (2024-12-20)


### Bug Fixes

* **antehandler:** fixed missing keeper in ante handler initialization ([#161](https://github.com/dymensionxyz/rollapp-wasm/issues/161)) ([b9ad42b](https://github.com/dymensionxyz/rollapp-wasm/commit/b9ad42b29fcb1901a7bf7a47986413fb7f705648))
* **app:** Fixed bech32 on account keeper to not be hardcoded  ([#54](https://github.com/dymensionxyz/rollapp-wasm/issues/54)) ([883653a](https://github.com/dymensionxyz/rollapp-wasm/commit/883653af7053450af80719e1cfd93e8309ba7a7d))
* **deps:** bumped dymint to `d51b961e7` to fix stuck submission bug ([#130](https://github.com/dymensionxyz/rollapp-wasm/issues/130)) ([4acf9e8](https://github.com/dymensionxyz/rollapp-wasm/commit/4acf9e80f1b1189a89dc1b39acb5706418a2157b))
* **deps:** bumped rdk to 3fe31b2db to fix denom-metadata transfer on unrelayed packets ([#131](https://github.com/dymensionxyz/rollapp-wasm/issues/131)) ([42ae0cf](https://github.com/dymensionxyz/rollapp-wasm/commit/42ae0cffb427defee32392663599bf1d2c34a482))
* **deps:** dymint bump to fix da grpc issue ([#136](https://github.com/dymensionxyz/rollapp-wasm/issues/136)) ([d6e15ba](https://github.com/dymensionxyz/rollapp-wasm/commit/d6e15ba4b90e6530cabf3dbd9c825b38ac0f6f2b))
* **deps:** fix tokenfactory init genesis validation  ([#173](https://github.com/dymensionxyz/rollapp-wasm/issues/173)) ([8db29dd](https://github.com/dymensionxyz/rollapp-wasm/commit/8db29dd28c77dbb1d4eb306ca7e0a5680b8f2ad7))
* **deps:** update rdk to fix tokenfactory denom-metadata override ([#188](https://github.com/dymensionxyz/rollapp-wasm/issues/188)) ([cbadaa9](https://github.com/dymensionxyz/rollapp-wasm/commit/cbadaa9ff6370e0537ff26cee3cf0b153910513b))
* **deps:** updated dymint and rdk by removing the replace ([#132](https://github.com/dymensionxyz/rollapp-wasm/issues/132)) ([a34bc94](https://github.com/dymensionxyz/rollapp-wasm/commit/a34bc942d86d658a11038c69e860c973e96a1053))
* **encoding:** support eip712 encoding  ([#203](https://github.com/dymensionxyz/rollapp-wasm/issues/203)) ([6168f5d](https://github.com/dymensionxyz/rollapp-wasm/commit/6168f5d587476ceab30e99cf29df7635bdacfbce))
* **ethsecp256k1:** register missing codecs for supporting ethsecp256k1 algo ([#138](https://github.com/dymensionxyz/rollapp-wasm/issues/138)) ([0f65b88](https://github.com/dymensionxyz/rollapp-wasm/commit/0f65b88b33ed669aa47bbb7183a62eb8e4068803))
* failing tests ([#149](https://github.com/dymensionxyz/rollapp-wasm/issues/149)) ([9044cc8](https://github.com/dymensionxyz/rollapp-wasm/commit/9044cc886d9d5813c6672e2a5a53119f148f0325))
* lint fix deprecated packages ([#167](https://github.com/dymensionxyz/rollapp-wasm/issues/167)) ([684b2ea](https://github.com/dymensionxyz/rollapp-wasm/commit/684b2ea00488db7b6bb4153096114747b5f66c39))
* merge conflict ([#13](https://github.com/dymensionxyz/rollapp-wasm/issues/13)) ([2cc8431](https://github.com/dymensionxyz/rollapp-wasm/commit/2cc8431a3dc57a60efece2a485c7298c08d22ecb))
* removed MsgSend from allowed consensus msgs ([#146](https://github.com/dymensionxyz/rollapp-wasm/issues/146)) ([21753d2](https://github.com/dymensionxyz/rollapp-wasm/commit/21753d26066ec4320e892aeef577c2899a8c756d))
* **rollappparams:** Validate gas price gov proposal param change ([#204](https://github.com/dymensionxyz/rollapp-wasm/issues/204)) ([8cb11b0](https://github.com/dymensionxyz/rollapp-wasm/commit/8cb11b01b2792a52688835ab4eda21a58d8dd194))
* **scripts:** adjusted IBC script to work with whitelisted relayers ([#154](https://github.com/dymensionxyz/rollapp-wasm/issues/154)) ([8df5b11](https://github.com/dymensionxyz/rollapp-wasm/commit/8df5b11ec32421434d1a98095ea0cdc0f976591b))
* **scripts:** fix init script to support mock backends ([#104](https://github.com/dymensionxyz/rollapp-wasm/issues/104)) ([65a57ca](https://github.com/dymensionxyz/rollapp-wasm/commit/65a57ca2b00141485ce7e32ab82d6a965b9d336b))
* updated IBC Keeper to use the sequencer keeper. ([#63](https://github.com/dymensionxyz/rollapp-wasm/issues/63)) ([6c4a2b6](https://github.com/dymensionxyz/rollapp-wasm/commit/6c4a2b674527476ad08e790dfd4b41ef18f086e3))
* **upgrade:** drs upgrade from 1 to 3 fix  ([#199](https://github.com/dymensionxyz/rollapp-wasm/issues/199)) ([b2e0573](https://github.com/dymensionxyz/rollapp-wasm/commit/b2e0573ffff19607e51dfac681d026bf9b78a9dd))
* **version:** bump dymint to c0e39f93d729 ([#166](https://github.com/dymensionxyz/rollapp-wasm/issues/166)) ([e798c0c](https://github.com/dymensionxyz/rollapp-wasm/commit/e798c0c89591b950be24e32f4d7ebc3f44d634ba))


### Features

* (tokenfactory): Wire module ([#158](https://github.com/dymensionxyz/rollapp-wasm/issues/158)) ([92712f7](https://github.com/dymensionxyz/rollapp-wasm/commit/92712f74b0d76fd1dd2f098b5a9a3de8122f45bf))
* add ethsecp256k1 as the default signing algo ([#80](https://github.com/dymensionxyz/rollapp-wasm/issues/80)) ([7362c6f](https://github.com/dymensionxyz/rollapp-wasm/commit/7362c6f89ba701d3103a5c25bbe45f01de0321f6))
* add hub genesis module ([#43](https://github.com/dymensionxyz/rollapp-wasm/issues/43)) ([73b3ceb](https://github.com/dymensionxyz/rollapp-wasm/commit/73b3cebef6c159494f0a4074ef5edb804b82bf0c))
* Add wasm module for rollapp-wasm ([#10](https://github.com/dymensionxyz/rollapp-wasm/issues/10)) ([9829d4a](https://github.com/dymensionxyz/rollapp-wasm/commit/9829d4a10b9f7928c98151b7295b20f0d54a8ad0))
* **ante:** added global min gas price checks ([#168](https://github.com/dymensionxyz/rollapp-wasm/issues/168)) ([47eb30c](https://github.com/dymensionxyz/rollapp-wasm/commit/47eb30c8cb8fdf1fe3c7819b38f82614072a4526))
* **ante:** skip fees for IBC messages  ([#127](https://github.com/dymensionxyz/rollapp-wasm/issues/127)) ([ff8e895](https://github.com/dymensionxyz/rollapp-wasm/commit/ff8e895578215eb62acb6582bfe0a0b8902326f6))
* **ante:** whitelisted relayers ([#148](https://github.com/dymensionxyz/rollapp-wasm/issues/148)) ([d815566](https://github.com/dymensionxyz/rollapp-wasm/commit/d81556668196e2c9ac133d9c8f32298e19c14afe))
* **app:** add denommetadata IBC middleware to app ([#94](https://github.com/dymensionxyz/rollapp-wasm/issues/94)) ([9a4756e](https://github.com/dymensionxyz/rollapp-wasm/commit/9a4756e0cd12bd2faa43531377ff51c15f5ce58f))
* **app:** Add modules authz and feegrant ([#60](https://github.com/dymensionxyz/rollapp-wasm/issues/60)) ([a4451ea](https://github.com/dymensionxyz/rollapp-wasm/commit/a4451eaebd11eb49c89a40c239f6dd8593f201d1))
* **app:** return genesis bridge data in InitChainResponse ([#162](https://github.com/dymensionxyz/rollapp-wasm/issues/162)) ([68c538a](https://github.com/dymensionxyz/rollapp-wasm/commit/68c538afed46cbde7e5cc10b981298d3c29173d9))
* **be:** integrate block explorer Json-RPC server ([#41](https://github.com/dymensionxyz/rollapp-wasm/issues/41)) ([51fd3e3](https://github.com/dymensionxyz/rollapp-wasm/commit/51fd3e36a0404d68325c64f79f65a15afc3be82a))
* **build:** Add ability to override bech32 with env var ([#192](https://github.com/dymensionxyz/rollapp-wasm/issues/192)) ([25ce045](https://github.com/dymensionxyz/rollapp-wasm/commit/25ce0453d1b2ee3429c9d108cadc2e7a5850e25e))
* **callback:** add callback and cwerrors module for scheduled execution of wasm contracts ([#99](https://github.com/dymensionxyz/rollapp-wasm/issues/99)) ([7936ae2](https://github.com/dymensionxyz/rollapp-wasm/commit/7936ae2bfc57ac138989ae39eb968d3b07916bf1))
* changed DRS to be int instead of commit hash ([#151](https://github.com/dymensionxyz/rollapp-wasm/issues/151)) ([fd1f992](https://github.com/dymensionxyz/rollapp-wasm/commit/fd1f992eadd01db79ec17cc511e753eb57b89ca6))
* **ci:** Add auto update changelog workflow ([#61](https://github.com/dymensionxyz/rollapp-wasm/issues/61)) ([ed9c6da](https://github.com/dymensionxyz/rollapp-wasm/commit/ed9c6da98f33a9842ae83007b46bc074f67d2152))
* **ci:** Add setup script and push hook ([#86](https://github.com/dymensionxyz/rollapp-wasm/issues/86)) ([d4dc3e4](https://github.com/dymensionxyz/rollapp-wasm/commit/d4dc3e4d73a72ab0e99cefc79c82eb0dcd79b187))
* consensus messages ([#141](https://github.com/dymensionxyz/rollapp-wasm/issues/141)) ([692fb70](https://github.com/dymensionxyz/rollapp-wasm/commit/692fb7096d6a7cb73ace726d1cddb5c276f02de5))
* **deps:** bump to last rdk version + tests fix ([#190](https://github.com/dymensionxyz/rollapp-wasm/issues/190)) ([65f00f6](https://github.com/dymensionxyz/rollapp-wasm/commit/65f00f6307fa316f4b200d9fff6b2f575bdba232))
* **deps:** bumped rdk to support tokenless feature ([#195](https://github.com/dymensionxyz/rollapp-wasm/issues/195)) ([6aea896](https://github.com/dymensionxyz/rollapp-wasm/commit/6aea896ca88789e3bc16afae80b5379e466836e9))
* **deps:** new genesis-bridge flow and rollappparams upgrade fix ([#183](https://github.com/dymensionxyz/rollapp-wasm/issues/183)) ([b023b94](https://github.com/dymensionxyz/rollapp-wasm/commit/b023b94c150ae10caefef9cc08667af7844de535))
* **deps:** validate-genesis-bridge command ([#177](https://github.com/dymensionxyz/rollapp-wasm/issues/177)) ([93d7626](https://github.com/dymensionxyz/rollapp-wasm/commit/93d7626b84437ce97b7233f6269319edea60d9a9))
* **fork:** add changes to include bump fork l2 ([#157](https://github.com/dymensionxyz/rollapp-wasm/issues/157)) ([4bedb8d](https://github.com/dymensionxyz/rollapp-wasm/commit/4bedb8ddecac5015803f4128cf23390f9c624191))
* **gasless:** wire the `x/gasless` module ([#102](https://github.com/dymensionxyz/rollapp-wasm/issues/102)) ([823dee3](https://github.com/dymensionxyz/rollapp-wasm/commit/823dee3cf52f205ccff47aa94e5669aa13c2ff9c))
* **genesis transfers:** wires genesis transfers ([#97](https://github.com/dymensionxyz/rollapp-wasm/issues/97)) ([a00a8c6](https://github.com/dymensionxyz/rollapp-wasm/commit/a00a8c6d96668bf917c2ca7a0597b1b62ad7a3a7))
* **genesis_bridge:** update to use new genesis bridge ([#143](https://github.com/dymensionxyz/rollapp-wasm/issues/143)) ([97b6b38](https://github.com/dymensionxyz/rollapp-wasm/commit/97b6b38240b2b234ed2fda9137f5c1d8001199b1))
* **genesis-templates:** add genesis templates for drs 3 ([#196](https://github.com/dymensionxyz/rollapp-wasm/issues/196)) ([e232e65](https://github.com/dymensionxyz/rollapp-wasm/commit/e232e65e4fd5441169dd5dc9efc8e6580dd2aed2))
* **hubgenesis:** add genesis checksum to genesis-info ([#152](https://github.com/dymensionxyz/rollapp-wasm/issues/152)) ([a8716c7](https://github.com/dymensionxyz/rollapp-wasm/commit/a8716c7549776b8d314340c67ddabe94549287bb))
* **makefile:** create genesis template with DRS from make file ([#176](https://github.com/dymensionxyz/rollapp-wasm/issues/176)) ([bcd3815](https://github.com/dymensionxyz/rollapp-wasm/commit/bcd38157b1466b2d8ba0abb1bfdaff3e0aa8f330))
* remove token factory module and move it to rdk ([#170](https://github.com/dymensionxyz/rollapp-wasm/issues/170)) ([7bf7515](https://github.com/dymensionxyz/rollapp-wasm/commit/7bf75158170cfb4732386c7349a36dc3704b017b))
* **sequencers:** wired sequencer rewards module. ([#119](https://github.com/dymensionxyz/rollapp-wasm/issues/119)) ([b6f0165](https://github.com/dymensionxyz/rollapp-wasm/commit/b6f01657c38dc47233515ac4f100213afce25028))
* set bech32 prefix without changing source code ([#68](https://github.com/dymensionxyz/rollapp-wasm/issues/68)) ([82c81a2](https://github.com/dymensionxyz/rollapp-wasm/commit/82c81a2e521669e2f0f48f34c9c8d56ed46d4196))
* **upgrade:** add drs 2 upgrade handler ([#186](https://github.com/dymensionxyz/rollapp-wasm/issues/186)) ([0408ebb](https://github.com/dymensionxyz/rollapp-wasm/commit/0408ebb448bcd5e4273246ae734bb753beb198cb))
* **upgrade:** add drs validation to allow gov based upgrades ([#182](https://github.com/dymensionxyz/rollapp-wasm/issues/182)) ([5a08f00](https://github.com/dymensionxyz/rollapp-wasm/commit/5a08f0094ce01669efe90f823818a3c0685330a9))
* **upgrade:** add upgrade handler for DRS4 ([#202](https://github.com/dymensionxyz/rollapp-wasm/issues/202)) ([0581f47](https://github.com/dymensionxyz/rollapp-wasm/commit/0581f476e1c207c6983d5acfc6d7c7584f1a1495))
* **wasm:** authz handler for signless contract execution ([#171](https://github.com/dymensionxyz/rollapp-wasm/issues/171)) ([8012a5f](https://github.com/dymensionxyz/rollapp-wasm/commit/8012a5fa66fc049e79207263a8d33218d02e1060))
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



