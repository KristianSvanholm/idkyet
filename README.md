# idk yet

Encode go structs into bytes and back!

## Why??
Because JSON is terrible!

Consider the following data:
```go

    type Account struct {
    	Email string
    	Pwd   string
    	Age   int8
    	Bal   int64
    }

    Account {
        Email : "john@email.com",
        Pwd : "SomeHash",
        Age : 25,
        Bal : 25600
    }

```

The JSON for this equates to `64` bytes of data, lots of which is JSON garbage data such as quotes, colons, commas and the keyname data.  
The pure byte data however, is only `33` bytes long. A ~48.5% reduction!

Great Success!

## Support

This project currently supports the following Golang datatypes
- int8
- int16
- int32
- int64
- float32
- float64

## Planned support

- Arrays
- Structs within structs within structs
- Optional fields
- Ignored fields

