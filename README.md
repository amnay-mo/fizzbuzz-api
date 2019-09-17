# FizzBuzz API

[![Build Status](https://travis-ci.com/amnay-mo/fizzbuzz-api.svg?token=zxyRL3yiy4zMBuyCG2my&branch=dev)](https://travis-ci.org/amnay-mo/fizzbuzz-api)

A REST API that does FizzBuzz!

## Building

Just run:

```
make lint test build
```

## Running

Just run the binary!
You may specify the HTTP port and the Redis backend's address with the `APP_PORT` and `REDIS_ADDR` env var:

```
REDIS_ADDR=localhost:6379 APP_PORT=9000 ./fizzbuzz-api
```

## Usage

- This is how you call the fizzbuzz endpoint:

```
GET /fizzbuzz?fizzNumber=2&buzzNumber=3&limit=10&fizzWord=Fizz&buzzWord=Buzz
```

- You may also get stats on the most queried fizzbuzz parameters:

```
GET /fizzbuzz/stats
```

## License

MIT
