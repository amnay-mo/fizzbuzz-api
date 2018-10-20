# FizzBuzz API

A REST API that does FizzBuzz!

## Building

Just run:

```
make lint test build
```

## Running

Just run the binary!
You may also specify the listen port with the `APP_PORT` env var:

```
APP_PORT=9000 ./foobar-api
```

## API Endpoints

There is only one:

```
/fizzbuzz?fizzNumber=2&buzzNumber=3&limit=10&fizzWord=Fizz&buzzWord=Buzz
```

## License

MIT
