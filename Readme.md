# base64fix

Go’s `encoding/base64` package does not fully decode base64 encoded data which isn’t padded. This is problematic with strings encoded with a base64 URL encoding strategy in which padding is optional per the specification and isn’t provided by some encoders—especially in cases where the data length is known, either by a HTTP header or by being parsed from a JSON data payload.

This tiny library wraps Go’s `encoding/base64` decoding with quick check to ensure that the data you’re passing in is padded and, if needed, will fix it by adding the correct padding.

## Usage

To use, first import this library:

```
import "github.com/duncan/base64fix"
```

Then, instead of decoding base64 data using the following which won’t fully decode the base64 data and which will return a `CorruptInputError`:

```go
s, err := base64.URLEncoding.DecodeString("YWJjZGU")
```

This will return `abc` and an error that indicates a problem. Instead of this, you can use the following:

```go
s, err := base64fix.URLEncoding.DecodeString("YWJjZGU")
```

This correctly decodes `abcde` with no error.

In addition to `DecodeString`, an implementation that wraps the default `Decode` function is provided:

```go
i, err := StdEncoding.Decode(d, []byte(s))
```

If the behavior of Go’s `encoding/base64` package changes and becomes more accepting of unpadded content, you can easily move back to the default implementation by changing `base64fix` to `base64` and dropping the import.

## Discussion of padding requirements

[RFC 4648](http://www.faqs.org/rfcs/rfc4648.html) has this to say in Section 3.2 about padding in base-encoded data:

> In some circumstances, the use of padding ("=") in base-encoded data
   is not required or used.  In the general case, when assumptions about
   the size of transported data cannot be made, padding is required to
   yield correct decoded data.

Furthermore, it says in Section 5:

> The pad character "=" is typically percent-encoded when used in an
   URI [9], but if the data length is known implicitly, this can be
   avoided by skipping the padding; see section 3.2.

In other words, it’s complicated. But, in the final analysis, getting things done in the real world when clients send URL encoded base64 data without padding means being robust and dealing with it accordingly.

## Questions

**What about fixing it in Go’s core library?**

Probably not going to happen any time soon. The issue of the core library’s strict decoding has been discussed since at least 2012—see [issue #4237 on github.com/golang/go](https://github.com/golang/go/issues/4237). It has been postponed multiple times by the core maintainers and doesn’t look like it’ll be addressed anytime soon.

**What about encoding? Why not a wrapper to remove padding?**

Accepting base64 data without padding is in line with [Postel’s law](http://en.wikipedia.org/wiki/Robustness_principle), but so is being strict with output. Even in cases where padding is optional, there’s no need to remove it if it’s already been generated. Since this is the thinnest possible wrapper over the core library, there’s no need to address strict encoding.

## License

Licensed under the Apache License, Version 2.0
