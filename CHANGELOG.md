# Changelog

## [3.2.7](https://github.com/RobertYoung/manutd-ticket-checker/compare/v3.2.6...v3.2.7) (2022-11-04)


### Bug Fixes

* replace parenthesis in home assistant ids ([910f660](https://github.com/RobertYoung/manutd-ticket-checker/commit/910f660ab49797b823aa2c21d5aa24132d99277c))

## [3.2.6](https://github.com/RobertYoung/manutd-ticket-checker/compare/v3.2.5...v3.2.6) (2022-10-29)


### Bug Fixes

* adds timeout argument to help with rod hanging ([0c228e5](https://github.com/RobertYoung/manutd-ticket-checker/commit/0c228e52dfb3b0bb8be13eb17557f225309feece))

## [3.2.5](https://github.com/RobertYoung/manutd-ticket-checker/compare/v3.2.4...v3.2.5) (2022-10-26)


### Bug Fixes

* removes timeouts and adds debugging ([bf776fb](https://github.com/RobertYoung/manutd-ticket-checker/commit/bf776fba6a9158f90e0450f9d50de907b85059e5))

## [3.2.4](https://github.com/RobertYoung/manutd-ticket-checker/compare/v3.2.3...v3.2.4) (2022-10-20)


### Bug Fixes

* fixes add to basket css selector ([5e07db2](https://github.com/RobertYoung/manutd-ticket-checker/commit/5e07db261461613075cc28a184bc5e399182da1d))

## [3.2.3](https://github.com/RobertYoung/manutd-ticket-checker/compare/v3.2.2...v3.2.3) (2022-10-20)


### Bug Fixes

* improves error handling ([e6059d2](https://github.com/RobertYoung/manutd-ticket-checker/commit/e6059d2ccde41cf26d40b443139c1bcd52e88e79))

## [3.2.2](https://github.com/RobertYoung/manutd-ticket-checker/compare/v3.2.1...v3.2.2) (2022-10-19)


### Bug Fixes

* check for errors if it can't find an element ([297ddab](https://github.com/RobertYoung/manutd-ticket-checker/commit/297ddab3048ff7e07506dd58fd13e52ef8bf34ad))

## [3.2.1](https://github.com/RobertYoung/manutd-ticket-checker/compare/v3.2.0...v3.2.1) (2022-10-16)


### Bug Fixes

* fixes panic if no available events are found ([3806602](https://github.com/RobertYoung/manutd-ticket-checker/commit/38066022172fd1cf43d3b1b095e517f1cec6fa15))

## [3.2.0](https://github.com/RobertYoung/manutd-ticket-checker/compare/v3.1.0...v3.2.0) (2022-10-15)


### Features

* checks the seats are available by trying to add to basket ([f6644dd](https://github.com/RobertYoung/manutd-ticket-checker/commit/f6644dd40a437b697cd1418a770241332cf3ef0c))

## [3.1.0](https://github.com/RobertYoung/manutd-ticket-checker/compare/v3.0.4...v3.1.0) (2022-10-01)


### Features

* don't show events that have a max price of 0 ([5a5e7cd](https://github.com/RobertYoung/manutd-ticket-checker/commit/5a5e7cd8974ff34b2f467b2d2fbf8e2f1e798b9f))

## [3.0.4](https://github.com/RobertYoung/manutd-ticket-checker/compare/v3.0.3...v3.0.4) (2022-09-30)


### Bug Fixes

* fixes notification sent at override in csv file ([b2f3c5f](https://github.com/RobertYoung/manutd-ticket-checker/commit/b2f3c5f980c2cc4c2e19be392ed44af245d8ccce))

## [3.0.3](https://github.com/RobertYoung/manutd-ticket-checker/compare/v3.0.2...v3.0.3) (2022-09-29)


### Bug Fixes

* reading a file when it doesn't exist ([77cff79](https://github.com/RobertYoung/manutd-ticket-checker/commit/77cff797ac868c1120fd4582d800e5db2734f508))

## [3.0.2](https://github.com/RobertYoung/manutd-ticket-checker/compare/v3.0.1...v3.0.2) (2022-09-29)


### Bug Fixes

* fixes overwriting a csv file ([f037f0e](https://github.com/RobertYoung/manutd-ticket-checker/commit/f037f0e35ab1210f00ccd8ceb640f92f9987cdf5))

## [3.0.1](https://github.com/RobertYoung/manutd-ticket-checker/compare/v3.0.0...v3.0.1) (2022-09-27)


### Bug Fixes

* adding missing mutc files ([8b20617](https://github.com/RobertYoung/manutd-ticket-checker/commit/8b206172e0af2f5a309cf3bcfb28ac4aea66792e))

## [3.0.0](https://github.com/RobertYoung/manutd-ticket-checker/compare/v2.5.3...v3.0.0) (2022-09-27)


### ??? BREAKING CHANGES

* upgrades to go 1.19

### Features

* adds max price and notification throttle flags ([06c840b](https://github.com/RobertYoung/manutd-ticket-checker/commit/06c840b005b1e2db4edf8c252f12c04d521c428c))
* don't mark an event to be notified more than once every 24 hours ([7b93196](https://github.com/RobertYoung/manutd-ticket-checker/commit/7b93196d18bb37387da4d6106864023e2f4946f5))
* sends notification per match found ([5e0d4d7](https://github.com/RobertYoung/manutd-ticket-checker/commit/5e0d4d748cb423a114e84c4cffc4590f4d3f20db))
* store event data to csv ([d67c044](https://github.com/RobertYoung/manutd-ticket-checker/commit/d67c0449061301188a6765a2937e18991a0c6a3b))
* upgrades to go 1.19 ([e12d063](https://github.com/RobertYoung/manutd-ticket-checker/commit/e12d06359f3b607d415261107fe0b747d203c0f1))

## [2.5.3](https://github.com/RobertYoung/manutd-ticket-checker/compare/v2.5.2...v2.5.3) (2022-09-25)


### Bug Fixes

* remove multiple platforms from linux/amd64 ([85114e0](https://github.com/RobertYoung/manutd-ticket-checker/commit/85114e05c5f13c38d4e5a58e6e17a73830a9548f))

## [2.5.2](https://github.com/RobertYoung/manutd-ticket-checker/compare/v2.5.1...v2.5.2) (2022-09-25)


### Bug Fixes

* panic if haas isn't enabled when sending a notification ([6f2ec4c](https://github.com/RobertYoung/manutd-ticket-checker/commit/6f2ec4cac984c6e3a3c473075eba3e31b8494151))
* release arm64 under a different tag ([6e28be9](https://github.com/RobertYoung/manutd-ticket-checker/commit/6e28be9612e651078fc447a2c08853ed6b4e0a91))

## [2.5.1](https://github.com/RobertYoung/manutd-ticket-checker/compare/v2.5.0...v2.5.1) (2022-09-25)


### Bug Fixes

* build images together but under a different os/arch ([26dc39a](https://github.com/RobertYoung/manutd-ticket-checker/commit/26dc39ad26cd8d0dcb6db50a982b8c8dde991939))

## [2.5.0](https://github.com/RobertYoung/manutd-ticket-checker/compare/v2.4.0...v2.5.0) (2022-09-25)


### Features

* pin version of go to 1.17 ([f130dd1](https://github.com/RobertYoung/manutd-ticket-checker/commit/f130dd13f0a8c9338d0aa5513c064759fa95f2d4))


### Bug Fixes

* adds docker buildx and qemu github actions ([08f2c09](https://github.com/RobertYoung/manutd-ticket-checker/commit/08f2c090102938f93b7b1d821b78fa905a20e360))

## [2.4.0](https://github.com/RobertYoung/manutd-ticket-checker/compare/v2.3.0...v2.4.0) (2022-09-25)


### Features

* log into dockerhub and updates docker build arch ([3a7357f](https://github.com/RobertYoung/manutd-ticket-checker/commit/3a7357f85685c8335096c164a6b3ec1db05144b7))

## [2.3.0](https://github.com/RobertYoung/manutd-ticket-checker/compare/v2.2.1...v2.3.0) (2022-09-25)


### Features

* adds support for docker ([7a266dc](https://github.com/RobertYoung/manutd-ticket-checker/commit/7a266dc7e2da8974f7bf82aaf39cffb766bebe90))
* sets up docker for goreleaser ([5e6dd54](https://github.com/RobertYoung/manutd-ticket-checker/commit/5e6dd54ae3e3a75aa5d7594a789e8867331ae2a3))

## [2.2.1](https://github.com/RobertYoung/manutd-ticket-checker/compare/v2.2.0...v2.2.1) (2022-09-24)


### Bug Fixes

* incorrect module name ([18825d1](https://github.com/RobertYoung/manutd-ticket-checker/commit/18825d18216c4dff27d50e16c024c2d499a41bd4))

## [2.2.0](https://github.com/RobertYoung/manutd-ticket-checker/compare/v2.1.0...v2.2.0) (2022-09-24)


### Features

* migrate module to github ([041d283](https://github.com/RobertYoung/manutd-ticket-checker/commit/041d283c74bdb234e287a1b5c299fd2b0197c0e5))

## [2.1.0](https://github.com/RobertYoung/manutd-ticket-checker/compare/v2.0.0...v2.1.0) (2022-09-24)


### Features

* adds more unavailable matches and imrpoves readme ([17bdcbb](https://github.com/RobertYoung/manutd-ticket-checker/commit/17bdcbb9f8b1df2a4c6fd16548868ac8d0318806))

## [2.0.0](https://github.com/RobertYoung/manutd-ticket-checker/compare/v1.6.0...v2.0.0) (2022-09-24)


### ??? BREAKING CHANGES

* push entity state to home assistant

### Features

* push entity state to home assistant ([3f483bb](https://github.com/RobertYoung/manutd-ticket-checker/commit/3f483bbaa2b1de1e89534226288e2096479d2770))
* send notification if available match found ([c20a276](https://github.com/RobertYoung/manutd-ticket-checker/commit/c20a2769afa5a93c085100b49c743c8b564c6341))

## [1.6.0](https://github.com/RobertYoung/manutd-ticket-checker/compare/v1.5.0...v1.6.0) (2022-09-24)


### Features

* adds haas api integration ([8476146](https://github.com/RobertYoung/manutd-ticket-checker/commit/8476146b78a74d4bd91a8b0c141a55227f151a59))

## [1.5.0](https://github.com/RobertYoung/manutd-ticket-checker/compare/v1.4.0...v1.5.0) (2022-09-24)


### Features

* use ga token ([4f9023e](https://github.com/RobertYoung/manutd-ticket-checker/commit/4f9023e5c897a2e730d61d7163c6c4773b41d4ed))

## [1.4.0](https://github.com/RobertYoung/manutd-ticket-checker/compare/v1.3.0...v1.4.0) (2022-09-24)


### Features

* run release-tag on release ([16cabdc](https://github.com/RobertYoung/manutd-ticket-checker/commit/16cabdcd5d5ea17a0abf78d7c4424598fb51595a))

## [1.3.0](https://github.com/RobertYoung/manutd-ticket-checker/compare/v1.2.1...v1.3.0) (2022-09-24)


### Features

* upload release artifacts ([c87e2ba](https://github.com/RobertYoung/manutd-ticket-checker/commit/c87e2babdc0f7fbe292cacdc2e85159ce33c908d))

## [1.2.1](https://github.com/RobertYoung/manutd-ticket-checker/compare/v1.2.0...v1.2.1) (2022-09-24)


### Bug Fixes

* removes release please ([ad763c8](https://github.com/RobertYoung/manutd-ticket-checker/commit/ad763c8490ad234c172ce6dca56ceb41dcbf90ae))

## [1.2.0](https://github.com/RobertYoung/manutd-ticket-checker/compare/v1.1.0...v1.2.0) (2022-09-24)


### Features

* trigger a release ([20773db](https://github.com/RobertYoung/manutd-ticket-checker/commit/20773db69549c6785661be28622527f59579430e))

## [1.1.0](https://github.com/RobertYoung/manutd-ticket-checker/compare/v1.0.0...v1.1.0) (2022-09-24)


### Features

* adds description and cli usage ([d6c9efe](https://github.com/RobertYoung/manutd-ticket-checker/commit/d6c9efef8cb072dd9807aa761380c730a6138219))

## 1.0.0 (2022-09-24)


### Features

* adds premier-league-only flag ([6f7ffec](https://github.com/RobertYoung/manutd-ticket-checker/commit/6f7ffecf38703de831e43c67eeaba3ad37c4606a))
* creates structs for pages and items ([b6360d6](https://github.com/RobertYoung/manutd-ticket-checker/commit/b6360d66a0910efbfd5e86d86c6432564f823d43))
* find page by url ([3618a61](https://github.com/RobertYoung/manutd-ticket-checker/commit/3618a612b07afb9a238109c331967d48cc25a313))
* finds events with available tickets and clicks buy now ([15005a2](https://github.com/RobertYoung/manutd-ticket-checker/commit/15005a2dbad420dfb6503865321ced856810046c))
* finds min and max price for an event ([aafe382](https://github.com/RobertYoung/manutd-ticket-checker/commit/aafe382c26a96d8651a5356e3288523969e0fa32))
* searchs ticket page and lists events ([a0969b1](https://github.com/RobertYoung/manutd-ticket-checker/commit/a0969b1dbe40f99f69ebfd3d84285025f4935a37))
