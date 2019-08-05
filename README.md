# _Google Analytics_ OMG Microservice

[![Open Microservice Guide](https://img.shields.io/badge/OMG%20Enabled-👍-green.svg?)](https://microservice.guide)
[![Build Status](https://travis-ci.com/omg-services/google-analytics.svg?branch=master)](https://travis-ci.com/omg-services/google-analytics)
[![codecov](https://codecov.io/gh/omg-services/google-analytics/branch/master/graph/badge.svg)](https://codecov.io/gh/omg-services/google-analytics)


An OMG service for google analytics, It is web analytics service that allows you to analyze in-depth detail about the visitors on your website. It provides valuable insights that can help you to shape the success strategy of your business.

## Direct usage in [Storyscript](https://storyscript.io/):


##### Account List
```coffee
google-analytics accountList
{"items": ["list"],"itemsPerPage": 1000,"kind": "kindResource","startIndex": 1,"totalResults": 1,"username": "accountUsername"}
```
##### WebProperties List
```coffee
google-analytics webPropertiesList accountId:'accountId'
{"items": ["list"],"itemsPerPage": 1000,"kind": "kindResource","startIndex": 1,"totalResults": 1,"username": "accountUsername"}
```
##### Profile List
```coffee
google-analytics profileList accountId:'accountId' webPropertyId:'webPropertyId'
{"items": ["list"],"itemsPerPage": 1000,"kind": "kindResource","startIndex": 1,"totalResults": 1,"username": "accountUsername"}
```
##### RealTime Data
```coffee
google-analytics realtime profileId:'profileId'
{"columnHeaders": ["list"],"id": "idLink","kind": "analytics#realtimeData","profileInfo": {    "profileDetails"},"query": {"queryDetails"},"rows": ["rowCount"],"selfLink": "selfLink",  "totalResults": 100,"totalsForAllResults": {"totalsForAllResultsDetails"}}
```

Curious to [learn more](https://docs.storyscript.io/)?

✨🍰✨

## Usage with [OMG CLI](https://www.npmjs.com/package/omg)

##### Account List
```shell
$ omg run accountList -e KEY=<BASE64_DATA_OF_JSON_KEY_FILE>
```
##### WebProperties List
```shell
$ omg run webPropertiesList -a accountId=<ACCOUNT_ID> -e KEY=<BASE64_DATA_OF_JSON_KEY_FILE>
```
##### Profile List
```shell
$ omg run profileList -a accountId=<ACCOUNT_ID> -a webPropertyId=<WEBPROPERTY_ID> -e KEY=<BASE64_DATA_OF_JSON_KEY_FILE>
```
##### RealTime Data
```shell
$ omg run realtime -a profileId=<PROFILE_ID> -e KEY=<BASE64_DATA_OF_JSON_KEY_FILE>
```

**Note**: The OMG CLI requires [Docker](https://docs.docker.com/install/) to be installed.

## License
[MIT License](https://github.com/omg-services/google-analytics/blob/master/LICENSE).
