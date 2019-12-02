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
You may configure the app using env vars:

- `APP_PORT`: listen port
- `REDIS_ADDR`: redis backend addr in the form of `<HOST>:<PORT>`
- `MAX_LIMIT`: the maximum limit parameter the application will accept; if exceeded, it will return a status `412`

Example:

```
REDIS_ADDR=localhost:6379 APP_PORT=9000 MAX_LIMIT=1024 ./fizzbuzz-api
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
