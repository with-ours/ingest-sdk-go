# Changelog

## 1.4.0 (2026-05-06)

Full Changelog: [v1.3.1...v1.4.0](https://github.com/with-ours/ingest-sdk-go/compare/v1.3.1...v1.4.0)

### Features

* **api:** api update ([bafaad7](https://github.com/with-ours/ingest-sdk-go/commit/bafaad7f8a8c12c7355e371d30d5489324491168))
* **api:** manual updates ([fb954d1](https://github.com/with-ours/ingest-sdk-go/commit/fb954d152f2218592d872cf2cfde1ddf785e1a1d))

## 1.3.1 (2026-05-01)

Full Changelog: [v1.3.0...v1.3.1](https://github.com/with-ours/ingest-sdk-go/compare/v1.3.0...v1.3.1)

### Chores

* avoid embedding reflect.Type for dead code elimination ([596cb92](https://github.com/with-ours/ingest-sdk-go/commit/596cb92b8e651c1e22d744c891d4bb7e497ff165))

## 1.3.0 (2026-04-29)

Full Changelog: [v1.2.1...v1.3.0](https://github.com/with-ours/ingest-sdk-go/compare/v1.2.1...v1.3.0)

### Features

* **api:** api update ([faf6737](https://github.com/with-ours/ingest-sdk-go/commit/faf673779164018959162a849cbb71cc172b919a))
* **api:** api update ([bedcea2](https://github.com/with-ours/ingest-sdk-go/commit/bedcea241945b71ac03cfcdd5f314d931790dd66))
* **go:** add default http client with timeout ([7cceb57](https://github.com/with-ours/ingest-sdk-go/commit/7cceb570e10eafd16291fa03da4b8b74d6b6d1d5))
* support setting headers via env ([1c1106c](https://github.com/with-ours/ingest-sdk-go/commit/1c1106cfb6891aa65bda2d38732d4d91d79d6def))

## 1.2.1 (2026-04-23)

Full Changelog: [v1.2.0...v1.2.1](https://github.com/with-ours/ingest-sdk-go/compare/v1.2.0...v1.2.1)

### Chores

* **internal:** more robust bootstrap script ([d82aa2e](https://github.com/with-ours/ingest-sdk-go/commit/d82aa2eafacc61a71531c54d674d51aef6f49c57))

## 1.2.0 (2026-03-28)

Full Changelog: [v1.1.3...v1.2.0](https://github.com/with-ours/ingest-sdk-go/compare/v1.1.3...v1.2.0)

### Features

* **api:** api update ([d394cc0](https://github.com/with-ours/ingest-sdk-go/commit/d394cc067f3244e661cc8e90ae67cabf408c2066))
* **internal:** support comma format in multipart form encoding ([a843265](https://github.com/with-ours/ingest-sdk-go/commit/a8432652dc0da619c4cca1176d85d2bdffb22a99))


### Bug Fixes

* prevent duplicate ? in query params ([daf1b4a](https://github.com/with-ours/ingest-sdk-go/commit/daf1b4af8aff6cdc38a8d8130c63f54a57e47deb))


### Chores

* **ci:** support opting out of skipping builds on metadata-only commits ([1c2b573](https://github.com/with-ours/ingest-sdk-go/commit/1c2b5738ab6d2496563087c4af2c16632f5a9736))
* **client:** fix multipart serialisation of Default() fields ([963a604](https://github.com/with-ours/ingest-sdk-go/commit/963a604971d46ebf602cdbca4ce13020c256eefb))
* **internal:** support default value struct tag ([21e8e96](https://github.com/with-ours/ingest-sdk-go/commit/21e8e967c5cc46b298c1dc606e880e593d1ddc45))
* remove unnecessary error check for url parsing ([42d8b43](https://github.com/with-ours/ingest-sdk-go/commit/42d8b43e1ce1275181e3249bf0a55d92b80618fb))
* update docs for api:"required" ([56f7e78](https://github.com/with-ours/ingest-sdk-go/commit/56f7e784e3c32a1428592161d6fc0af8310bd73e))

## 1.1.3 (2026-03-25)

Full Changelog: [v1.1.2...v1.1.3](https://github.com/with-ours/ingest-sdk-go/compare/v1.1.2...v1.1.3)

### Chores

* **ci:** skip lint on metadata-only changes ([277f854](https://github.com/with-ours/ingest-sdk-go/commit/277f854a9bbd3ec79c0f110dc55666a19ed3f4c4))
* **internal:** update gitignore ([88895df](https://github.com/with-ours/ingest-sdk-go/commit/88895df12ffdb37472958219e541825dbf7c3bb1))

## 1.1.2 (2026-03-17)

Full Changelog: [v1.1.1...v1.1.2](https://github.com/with-ours/ingest-sdk-go/compare/v1.1.1...v1.1.2)

### Chores

* **internal:** tweak CI branches ([ab3df17](https://github.com/with-ours/ingest-sdk-go/commit/ab3df1788f43bdaf5a04eea6253cee92876b990a))

## 1.1.1 (2026-03-11)

Full Changelog: [v1.1.0...v1.1.1](https://github.com/with-ours/ingest-sdk-go/compare/v1.1.0...v1.1.1)

### Chores

* **ci:** skip uploading artifacts on stainless-internal branches ([6e6d134](https://github.com/with-ours/ingest-sdk-go/commit/6e6d1342c641791bb4e594eb1862d3a5e7fd07f4))
* **internal:** minor cleanup ([89fc92f](https://github.com/with-ours/ingest-sdk-go/commit/89fc92ffb114d44cff9181851346cb8b10b5888f))
* **internal:** use explicit returns ([6c4985f](https://github.com/with-ours/ingest-sdk-go/commit/6c4985f22bb96a113a7e9d3ec682a2f043358987))
* **internal:** use explicit returns in more places ([d7cd131](https://github.com/with-ours/ingest-sdk-go/commit/d7cd131aa14323b0a2f84ac40e9baad95ea6e220))

## 1.1.0 (2026-03-05)

Full Changelog: [v1.0.0...v1.1.0](https://github.com/with-ours/ingest-sdk-go/compare/v1.0.0...v1.1.0)

### Features

* **api:** api update ([c534253](https://github.com/with-ours/ingest-sdk-go/commit/c534253401c0f3120088cf0b5ce30c728bbdc74a))


### Bug Fixes

* fix request delays for retrying to be more respectful of high requested delays ([7c13098](https://github.com/with-ours/ingest-sdk-go/commit/7c1309801d391e4911516a2a6fefc8396e484fe7))


### Chores

* **internal:** codegen related update ([fe04004](https://github.com/with-ours/ingest-sdk-go/commit/fe040047303041d043ec07216024506f4e31806f))

## 1.0.0 (2026-02-25)

Full Changelog: [v0.9.1...v1.0.0](https://github.com/with-ours/ingest-sdk-go/compare/v0.9.1...v1.0.0)

### Features

* **api:** api update ([0833332](https://github.com/with-ours/ingest-sdk-go/commit/08333324900d2d9dd7726d3357e855c6b5e115ff))


### Bug Fixes

* allow canceling a request while it is waiting to retry ([79b8da6](https://github.com/with-ours/ingest-sdk-go/commit/79b8da62992b41fe5a24cc0a7ad6f1f2cfd12255))


### Chores

* **internal:** move custom custom `json` tags to `api` ([1ec0bd3](https://github.com/with-ours/ingest-sdk-go/commit/1ec0bd39d7d27409f92a6f2a645edf1c6c985a95))
* **internal:** remove mock server code ([aa66202](https://github.com/with-ours/ingest-sdk-go/commit/aa66202b3d25a4be6aef72c3a662f19ae529b0d0))
* update mock server docs ([5062790](https://github.com/with-ours/ingest-sdk-go/commit/50627906625584034a0c924d9ca01b636dd2f692))

## 0.9.1 (2026-02-11)

Full Changelog: [v0.9.0...v0.9.1](https://github.com/with-ours/ingest-sdk-go/compare/v0.9.0...v0.9.1)

### Bug Fixes

* **encoder:** correctly serialize NullStruct ([b90433d](https://github.com/with-ours/ingest-sdk-go/commit/b90433d205437cfdb0dcbdfffd778747c71aa286))

## 0.9.0 (2026-02-04)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/with-ours/ingest-sdk-go/compare/v0.8.0...v0.9.0)

### Features

* **api:** api update ([0cd1384](https://github.com/with-ours/ingest-sdk-go/commit/0cd1384bb7b4e3b13a66b860bfbf7dd4023e1b4a))

## 0.8.0 (2026-01-24)

Full Changelog: [v0.7.3...v0.8.0](https://github.com/with-ours/ingest-sdk-go/compare/v0.7.3...v0.8.0)

### Features

* **client:** add a convenient param.SetJSON helper ([c623fd6](https://github.com/with-ours/ingest-sdk-go/commit/c623fd647911cac3dff99f12df9326c142ce58ca))

## 0.7.3 (2026-01-17)

Full Changelog: [v0.7.2...v0.7.3](https://github.com/with-ours/ingest-sdk-go/compare/v0.7.2...v0.7.3)

### Bug Fixes

* **docs:** add missing pointer prefix to api.md return types ([b7bed3b](https://github.com/with-ours/ingest-sdk-go/commit/b7bed3b3ac32790fcf5f9b5db257b6f062df6e7b))


### Chores

* **internal:** update `actions/checkout` version ([c2af041](https://github.com/with-ours/ingest-sdk-go/commit/c2af0411b29434588b96de45f7d8d543fa982c70))

## 0.7.2 (2026-01-06)

Full Changelog: [v0.7.1...v0.7.2](https://github.com/with-ours/ingest-sdk-go/compare/v0.7.1...v0.7.2)

### Chores

* **internal:** codegen related update ([f76e161](https://github.com/with-ours/ingest-sdk-go/commit/f76e161f968f0648ccb5703e6d5c592d7c17a1e5))

## 0.7.1 (2025-12-19)

Full Changelog: [v0.7.0...v0.7.1](https://github.com/with-ours/ingest-sdk-go/compare/v0.7.0...v0.7.1)

### Bug Fixes

* skip usage tests that don't work with Prism ([c48cf04](https://github.com/with-ours/ingest-sdk-go/commit/c48cf04fbec25f1b1f93dc666343ac7c1906ab05))


### Chores

* add float64 to valid types for RegisterFieldValidator ([f227dad](https://github.com/with-ours/ingest-sdk-go/commit/f227dadaa984d923c41652498ed732ac6332be50))

## 0.7.0 (2025-12-17)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/with-ours/ingest-sdk-go/compare/v0.6.0...v0.7.0)

### Features

* **api:** api update ([2c2c245](https://github.com/with-ours/ingest-sdk-go/commit/2c2c2450e6b40e99e3bdc0df5fcec012af3377bf))

## 0.6.0 (2025-12-12)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/with-ours/ingest-sdk-go/compare/v0.5.0...v0.6.0)

### Features

* **encoder:** support bracket encoding form-data object members ([aa97ee0](https://github.com/with-ours/ingest-sdk-go/commit/aa97ee05fe962dafd936a75aa87542ad609610dd))

## 0.5.0 (2025-12-11)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/with-ours/ingest-sdk-go/compare/v0.4.0...v0.5.0)

### Features

* **api:** api update ([14618bc](https://github.com/with-ours/ingest-sdk-go/commit/14618bc2ac1f2535fafcb47306c5b3cfbe08421d))

## 0.4.0 (2025-12-11)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/with-ours/ingest-sdk-go/compare/v0.3.0...v0.4.0)

### Features

* **api:** api update ([c56dd4f](https://github.com/with-ours/ingest-sdk-go/commit/c56dd4f50290a29de0fc0e4d23e4825ff3fdb563))

## 0.3.0 (2025-12-10)

Full Changelog: [v0.2.1...v0.3.0](https://github.com/with-ours/ingest-sdk-go/compare/v0.2.1...v0.3.0)

### Features

* **api:** api update ([831932e](https://github.com/with-ours/ingest-sdk-go/commit/831932e8037d558784f5124eea18d4aa60d0949a))

## 0.2.1 (2025-12-06)

Full Changelog: [v0.2.0...v0.2.1](https://github.com/with-ours/ingest-sdk-go/compare/v0.2.0...v0.2.1)

### Bug Fixes

* **mcp:** correct code tool API endpoint ([1fadb43](https://github.com/with-ours/ingest-sdk-go/commit/1fadb43ced15b3fdfe3956ffd1d1699eb99d54e9))
* rename param to avoid collision ([f4ee438](https://github.com/with-ours/ingest-sdk-go/commit/f4ee43869fb88863276d7fcc2a73f98c44761d53))


### Chores

* elide duplicate aliases ([6bff750](https://github.com/with-ours/ingest-sdk-go/commit/6bff75047a42dacc11ff08f967ee975f88fe66a2))
* **internal:** codegen related update ([7c4cb0e](https://github.com/with-ours/ingest-sdk-go/commit/7c4cb0eadb21bca5f16235d24c244b4eef4a11fc))

## 0.2.0 (2025-11-26)

Full Changelog: [v0.1.1...v0.2.0](https://github.com/with-ours/ingest-sdk-go/compare/v0.1.1...v0.2.0)

### Features

* **api:** api update ([c672303](https://github.com/with-ours/ingest-sdk-go/commit/c67230360ae5ba13bbd33a855a2fa328ff5a80a1))
* **api:** api update ([94e8861](https://github.com/with-ours/ingest-sdk-go/commit/94e8861f223f63063fd49a622ce08e0f940b12b0))


### Chores

* fix empty interfaces ([44190ee](https://github.com/with-ours/ingest-sdk-go/commit/44190ee9a26a1334a7332a0add412dff58175d67))

## 0.1.1 (2025-11-19)

Full Changelog: [v0.1.0...v0.1.1](https://github.com/with-ours/ingest-sdk-go/compare/v0.1.0...v0.1.1)

### Chores

* bump gjson version ([fdfb43d](https://github.com/with-ours/ingest-sdk-go/commit/fdfb43dd6cdf5491ead43366702f9a54b87a7f46))

## 0.1.0 (2025-11-05)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/with-ours/ingest-sdk-go/compare/v0.0.1...v0.1.0)

### Features

* **api:** manual updates ([389ba5b](https://github.com/with-ours/ingest-sdk-go/commit/389ba5b10b91669bf7adecb385ebab890e80e672))
* **api:** manual updates ([ce39a2f](https://github.com/with-ours/ingest-sdk-go/commit/ce39a2f222bfa8f01e92d4a303214b15c2514cf4))


### Chores

* configure new SDK language ([cb1bfbc](https://github.com/with-ours/ingest-sdk-go/commit/cb1bfbc4e1a12acd9e526d4adc46c33c5a801bf4))
* **internal:** grammar fix (it's -&gt; its) ([00920f9](https://github.com/with-ours/ingest-sdk-go/commit/00920f9bf376563198fa9eb46d52a0996d5f2eaa))
* update SDK settings ([753eca6](https://github.com/with-ours/ingest-sdk-go/commit/753eca68fd028bc8b44bbe5aaaf69105cbe13ab0))
