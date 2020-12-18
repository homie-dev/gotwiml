# gotwiml

gotwiml is a library for generating TwiML written in golang

### Demo

A simple example of use is as follows

```go
func main() {
    resp := twiml.NewVoiceResponse().
        Say("hello world!", attr.Voice(voice.Alice))

    xml, _ := resp.ToXMLPretty("  ")
    fmt.Println(xml)
}
```
```xml
<?xml version="1.0" encoding="UTF-8"?>
<Response>
  <Say voice="alice">hello world!</Say>
</Response>
```

The Functional Options pattern is used so that attributes can be specified flexibly

```go
// say "hello world!"
resp := twiml.NewVoiceResponse().
    Say("hello world!")
```
```go
// say "hello world!" in a British accent by alice.
resp := twiml.NewVoiceResponse().
    Say("hello world!",
        attr.Voice(voice.Alice),
        attr.Language(language.EnGb),
    )
```

If there is more than one content to be executed, it is possible to write them in a method chain.

```go
// Multiple instructions using a method chain
resp := twiml.NewVoiceResponse().
    Say("hello world!").
    Dial("+81 90 0000 0000", attr.AnswerOnBridge(true))
```

Nested elements can also be added with `AppendXXX`.

```go
// Generating nested TwiML
dial := twiml.NewDial(
    attr.AnswerOnBridge(true),
    attr.Record(recording.FromAnswerDual),
).Client("john")
resp := twiml.NewVoiceResponse().AppendDial(dial)
```

### License
MIT License, see [LICENSE](https://github.com/homie-dev/gotwiml/blob/master/LICENSE).