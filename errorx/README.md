# errorx

errorx contains gogox's error package with additional features

## How to Use

`errorx.Error` implements the `error` interface. There are some notably additional features:

### Code

`errorx.Error.Code` is intended to be used as unique identifier of each gogox's error. Below is the example of intended use:

```go
err := service.DoSomething()

if err != nil {
  // Parse error to gogox error. It will return gogox error if successful, otherwise return not ok.
  gogoxErr, ok := errorx.ParseError(err)
  if !ok {
    // as it is not errorx's error, assume that it is the underlying error from dependencies that is not expected.
    // hence we can safely assume it is internal error.
    gogoxErr := errorx.New(errorx.CodeInternal, err.Error())
  }

  switch(gogoxErr.Code) {
  case "your.error.code.not_found":
    w.WriteHeader(404)
  case "your.other.error.code.unauthorized":
    w.WriteHeader(401)
  case "user.invalid.name":
    w.WriteHeader(422)
  default:
    w.WriteHeader(500)
  }
}
```

### LogError

`LogError` is intended to be use for logging. You can store more verbose error message (like including the `user_id` or `email`) for easier debugging process in your observability platform.

If not provided, `logMessage` is automatically built from `Code` and `Message`.

### StackTrace

For each `New` and `Wrap` (and their respective variants like `Newf` and `Wrapf`) call, it will store the stack caller. Use `PrintStackTrace` to print the stack trace.

### Details

You can add more `Details` for your gogox `Error`. Details are commonly used to store field information. Example: Invalid parameter of `name` or `email` field.
