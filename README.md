# go-totp
simple library to generate/check time-based one time pins (TOTP). expects secrets as base32 strings.

## Usage
checking a totp
```go
// get secret and input from user
if matches, err := CheckTotp(secret, tot.Code); !matches || err != nil {
	// handle failure
}
```
generating a totp
```go
// get secret
code, err := GenerateTotp(secret)
```
