# go-totp
simple library to generate/check time-based one time pins (TOTP)

## Usage
```go
  // get input from user
	if matches, err := CheckTotp(secret, tot.Code); !matches || err != nil {
		// handle failure
	}
```
