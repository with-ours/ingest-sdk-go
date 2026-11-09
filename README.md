# Ours Privacy Go API Library

<!-- x-release-please-start-version -->

<a href="https://pkg.go.dev/github.com/with-ours/ingest-sdk-go"><img src="https://pkg.go.dev/badge/github.com/with-ours/ingest-sdk-go.svg" alt="Go Reference"></a>

<!-- x-release-please-end -->

The Ours Privacy Go library provides convenient access to the [Ours Privacy REST API](https://docs.oursprivacy.com)
from applications written in Go.

It is generated with [Stainless](https://www.stainless.com/).

## Installation

<!-- x-release-please-start-version -->

```go
import (
	"github.com/with-ours/ingest-sdk-go" // imported as oursprivacy
)
```

<!-- x-release-please-end -->

Or to pin the version:

<!-- x-release-please-start-version -->

```sh
go get -u 'github.com/with-ours/ingest-sdk-go@v1.25.0'
```

<!-- x-release-please-end -->

## Requirements

This library requires Go 1.26+. Use a currently supported Go release with the latest security patches for production.

## Usage

`token` is the Source token from your Ours Privacy dashboard. Events must be allow-listed: an event that isn't on your [Event Allowlist](https://docs.oursprivacy.com/docs/allow-listing-events) isn't forwarded to your destinations.

See the [SDK guide](https://docs.oursprivacy.com/docs/go) and the runnable code in [`examples`](examples).

The full API of this library can be found in [api.md](api.md).

```go
package main

import (
	"context"
	"fmt"

	"github.com/with-ours/ingest-sdk-go"
)

func main() {
	client := oursprivacy.NewClient()
	response, err := client.Track.Event(context.TODO(), oursprivacy.TrackEventParams{
		Token:      "YOUR_SOURCE_TOKEN",
		Event:      "signup_completed",
		DistinctID: oursprivacy.String("evt_01J9ZK3Q7X"),
		EventProperties: map[string]oursprivacy.TrackEventParamsEventPropertyUnion{
			"plan": {
				OfString: oursprivacy.String("pro"),
			},
			"value": {
				OfFloat: oursprivacy.Float(49),
			},
			"currency": {
				OfString: oursprivacy.String("USD"),
			},
		},
		ExternalID: oursprivacy.String("crm_contact_1042"),
		IdentityContext: oursprivacy.TrackEventParamsIdentityContext{
			IP:        "203.0.113.10",
			UserAgent: "Mozilla/5.0",
		},
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", response.Success)
}

```

### Identify a visitor

Set properties on a visitor, creating the visitor if there's no match.

```go
package main

import (
	"context"
	"fmt"

	"github.com/with-ours/ingest-sdk-go"
)

func main() {
	client := oursprivacy.NewClient()
	response, err := client.Visitor.Upsert(context.TODO(), oursprivacy.VisitorUpsertParams{
		Token: "YOUR_SOURCE_TOKEN",
		UserProperties: oursprivacy.VisitorUpsertParamsUserProperties{
			FirstName: oursprivacy.String("Sample"),
			LastName:  oursprivacy.String("User"),
		},
		Email: oursprivacy.String("user@example.com"),
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", response.Success)
}

```

### Send a batch of events

Send many events in one request, each with its own required `distinctId`.

```go
package main

import (
	"context"
	"fmt"

	"github.com/with-ours/ingest-sdk-go"
)

func main() {
	client := oursprivacy.NewClient()
	batch, err := client.Batch.New(context.TODO(), oursprivacy.BatchNewParams{
		Token: "YOUR_SOURCE_TOKEN",
		Events: []oursprivacy.BatchNewParamsEvent{{
			Event:      "signup_completed",
			DistinctID: "evt_01J9ZK3Q7X",
			ExternalID: oursprivacy.String("crm_contact_1042"),
		}, {
			Event:      "plan_upgraded",
			DistinctID: "evt_01J9ZK3Q8Y",
			ExternalID: oursprivacy.String("crm_contact_1042"),
		}},
	})
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", batch.Accepted)
}

```

A partial success returns HTTP 207 with a result for each row. Retry only the rows whose `code` is `queue_failed`, with the same `distinctId`. You can backfill events up to one year old; older `time` values are clamped.

### Get an experiment assignment

Look up a visitor's variant, using the experiment's `exp_*` token rather than the Source token.

```go
package main

import (
	"context"
	"fmt"

	"github.com/with-ours/ingest-sdk-go"
)

func main() {
	client := oursprivacy.NewClient()
	response, err := client.Experiments.Assignment(
		context.TODO(),
		"checkout_redesign",
		oursprivacy.ExperimentAssignmentParams{
			Token:     "exp_YOUR_EXPERIMENT_TOKEN",
			VisitorID: "visitor_1042",
		},
	)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("%+v\n", response)
}

```

### Further reading

The docs cover each endpoint in full.

- [Track](https://docs.oursprivacy.com/docs/reference/track)
- [Identify](https://docs.oursprivacy.com/docs/reference/identify)
- [Batch](https://docs.oursprivacy.com/docs/reference/batch)
- [Rate limits](https://docs.oursprivacy.com/docs/reference/rate-limits)
- [Server-side experimentation](https://docs.oursprivacy.com/docs/experimentation/server-side)

### Request fields

Types such as `ExampleParams`, `Point`, `FooParams` and `AnimalUnion` in this section and those after it are illustrative, not part of this SDK.

The oursprivacy library uses the [`omitzero`](https://tip.golang.org/doc/go1.24#encodingjsonpkgencodingjson)
semantics from the Go 1.24+ `encoding/json` release for request fields.

Required primitive fields (`int64`, `string`, etc.) feature the tag <code>\`api:"required"\`</code>. These
fields are always serialized, even their zero values.

Optional primitive types are wrapped in a `param.Opt[T]`. These fields can be set with the provided constructors, `oursprivacy.String(string)`, `oursprivacy.Int(int64)`, etc.

Any `param.Opt[T]`, map, slice, struct or string enum uses the
tag <code>\`json:"...,omitzero"\`</code>. Its zero value is considered omitted.

The `param.IsOmitted(any)` function can confirm the presence of any `omitzero` field.

```go
p := oursprivacy.ExampleParams{
	ID:   "id_xxx",                  // required property
	Name: oursprivacy.String("..."), // optional property

	Point: oursprivacy.Point{
		X: 0,                  // required field will serialize as 0
		Y: oursprivacy.Int(1), // optional field will serialize as 1
		// ... omitted non-required fields will not be serialized
	},

	Origin: oursprivacy.Origin{}, // the zero value of [Origin] is considered omitted
}
```

To send `null` instead of a `param.Opt[T]`, use `param.Null[T]()`.
To send `null` instead of a struct `T`, use `param.NullStruct[T]()`.

```go
p.Name = param.Null[string]()       // 'null' instead of string
p.Point = param.NullStruct[Point]() // 'null' instead of struct

param.IsNull(p.Name)  // true
param.IsNull(p.Point) // true
```

Request structs contain a `.SetExtraFields(map[string]any)` method which can send non-conforming
fields in the request body. Extra fields overwrite any struct fields with a matching
key. For security reasons, only use `SetExtraFields` with trusted data.

To send a custom value instead of a struct, use `param.Override[T](value)`.

```go
// In cases where the API specifies a given type,
// but you want to send something else, use [SetExtraFields]:
p.SetExtraFields(map[string]any{
	"x": 0.01, // send "x" as a float instead of int
})

// Send a number instead of an object
custom := param.Override[oursprivacy.FooParams](12)
```

### Request unions

Unions are represented as a struct with fields prefixed by "Of" for each of its variants,
only one field can be non-zero. The non-zero field will be serialized.

Sub-properties of the union can be accessed via methods on the union struct.
These methods return a mutable pointer to the underlying data, if present.

```go
// Only one field can be non-zero, use param.IsOmitted() to check if a field is set
type AnimalUnionParam struct {
	OfCat *Cat `json:",omitzero,inline`
	OfDog *Dog `json:",omitzero,inline`
}

animal := AnimalUnionParam{
	OfCat: &Cat{
		Name: "Whiskers",
		Owner: PersonParam{
			Address: AddressParam{Street: "3333 Coyote Hill Rd", Zip: 0},
		},
	},
}

// Mutating a field
if address := animal.GetOwner().GetAddress(); address != nil {
	address.ZipCode = 94304
}
```

### Response objects

All fields in response structs are ordinary value types (not pointers or wrappers).
Response structs also include a special `JSON` field containing metadata about
each property.

```go
type Animal struct {
	Name   string `json:"name,nullable"`
	Owners int    `json:"owners"`
	Age    int    `json:"age"`
	JSON   struct {
		Name        respjson.Field
		Owner       respjson.Field
		Age         respjson.Field
		ExtraFields map[string]respjson.Field
	} `json:"-"`
}
```

To handle optional data, use the `.Valid()` method on the JSON field.
`.Valid()` returns true if a field is not `null`, not present, or couldn't be marshaled.

If `.Valid()` is false, the corresponding field will simply be its zero value.

```go
raw := `{"owners": 1, "name": null}`

var res Animal
json.Unmarshal([]byte(raw), &res)

// Accessing regular fields

res.Owners // 1
res.Name   // ""
res.Age    // 0

// Optional field checks

res.JSON.Owners.Valid() // true
res.JSON.Name.Valid()   // false
res.JSON.Age.Valid()    // false

// Raw JSON values

res.JSON.Owners.Raw()                  // "1"
res.JSON.Name.Raw() == "null"          // true
res.JSON.Name.Raw() == respjson.Null   // true
res.JSON.Age.Raw() == ""               // true
res.JSON.Age.Raw() == respjson.Omitted // true
```

These `.JSON` structs also include an `ExtraFields` map containing
any properties in the json response that were not specified
in the struct. This can be useful for API features not yet
present in the SDK.

```go
body := res.JSON.ExtraFields["my_unexpected_field"].Raw()
```

### Response Unions

In responses, unions are represented by a flattened struct containing all possible fields from each of the
object variants.
To convert it to a variant use the `.AsFooVariant()` method or the `.AsAny()` method if present.

If a response value union contains primitive values, primitive fields will be alongside
the properties but prefixed with `Of` and feature the tag `json:"...,inline"`.

```go
type AnimalUnion struct {
	// From variants [Dog], [Cat]
	Owner Person `json:"owner"`
	// From variant [Dog]
	DogBreed string `json:"dog_breed"`
	// From variant [Cat]
	CatBreed string `json:"cat_breed"`
	// ...

	JSON struct {
		Owner respjson.Field
		// ...
	} `json:"-"`
}

// If animal variant
if animal.Owner.Address.ZipCode == "" {
	panic("missing zip code")
}

// Switch on the variant
switch variant := animal.AsAny().(type) {
case Dog:
case Cat:
default:
	panic("unexpected type")
}
```

### RequestOptions

This library uses the functional options pattern. Functions defined in the
`option` package return a `RequestOption`, which is a closure that mutates a
`RequestConfig`. These options can be supplied to the client or at individual
requests. For example:

```go
client := oursprivacy.NewClient(
	// Adds a header to every request made by the client
	option.WithHeader("X-Some-Header", "custom_header_info"),
)

client.Track.Event(context.TODO(), ...,
	// Override the header
	option.WithHeader("X-Some-Header", "some_other_custom_header_info"),
	// Add an undocumented field to the request body, using sjson syntax
	option.WithJSONSet("some.json.path", map[string]string{"my": "object"}),
)
```

The request option `option.WithDebugLog(nil)` may be helpful while debugging.

See the [full list of request options](https://pkg.go.dev/github.com/with-ours/ingest-sdk-go/option).

### Errors

When the API returns a non-success status code, we return an error with type
`*oursprivacy.Error`. This contains the `StatusCode`, `*http.Request`, and
`*http.Response` values of the request, as well as the JSON of the error body
(much like other response objects in the SDK).

To handle errors, we recommend that you use the `errors.As` pattern:

```go
_, err := client.Track.Event(context.TODO(), oursprivacy.TrackEventParams{
	Token: "REPLACE_ME",
	Event: "REPLACE_ME",
})
if err != nil {
	var apierr *oursprivacy.Error
	if errors.As(err, &apierr) {
		println(string(apierr.DumpRequest(true)))  // Prints the serialized HTTP request, with secrets replaced
		println(string(apierr.DumpResponse(true))) // Prints the serialized HTTP response, with secrets replaced
	}
	panic(err.Error()) // GET "/track": 400 Bad Request { ... }
}
```

When other errors occur, they are returned unwrapped; for example,
if HTTP transport fails, you might receive `*url.Error` wrapping `*net.OpError`.

### Timeouts

Requests do not time out by default; use context to configure a timeout for a request lifecycle.

Note that if a request is [retried](#retries), the context timeout does not start over.
To set a per-retry timeout, use `option.WithRequestTimeout()`.

```go
// This sets the timeout for the request, including all the retries.
ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
defer cancel()
client.Track.Event(
	ctx,
	oursprivacy.TrackEventParams{
		Token: "REPLACE_ME",
		Event: "REPLACE_ME",
	},
	// This sets the per-retry timeout
	option.WithRequestTimeout(20*time.Second),
)
```

### Retries

Certain errors will be automatically retried 2 times by default, with a short exponential backoff.
We retry by default all connection errors, 408 Request Timeout, 409 Conflict, 429 Rate Limit,
and >=500 Internal errors.

You can use the `WithMaxRetries` option to configure or disable this:

```go
// Configure the default for all requests:
client := oursprivacy.NewClient(
	option.WithMaxRetries(0), // default is 2
)

// Override per-request:
client.Track.Event(
	context.TODO(),
	oursprivacy.TrackEventParams{
		Token: "REPLACE_ME",
		Event: "REPLACE_ME",
	},
	option.WithMaxRetries(5),
)
```

### Accessing raw response data (e.g. response headers)

You can access the raw HTTP response data by using the `option.WithResponseInto()` request option. This is useful when
you need to examine response headers, status codes, or other details.

```go
// Create a variable to store the HTTP response
var response *http.Response
response, err := client.Track.Event(
	context.TODO(),
	oursprivacy.TrackEventParams{
		Token: "REPLACE_ME",
		Event: "REPLACE_ME",
	},
	option.WithResponseInto(&response),
)
if err != nil {
	// handle error
}
fmt.Printf("%+v\n", response)

fmt.Printf("Status Code: %d\n", response.StatusCode)
fmt.Printf("Headers: %+#v\n", response.Header)
```

### Making custom/undocumented requests

This library is typed for convenient access to the documented API. If you need to access undocumented
endpoints, params, or response properties, the library can still be used.

#### Undocumented endpoints

To make requests to undocumented endpoints, you can use `client.Get`, `client.Post`, and other HTTP verbs.
`RequestOptions` on the client, such as retries, will be respected when making these requests.

```go
var (
    // params can be an io.Reader, a []byte, an encoding/json serializable object,
    // or a "…Params" struct defined in this library.
    params map[string]any

    // result can be an []byte, *http.Response, a encoding/json deserializable object,
    // or a model defined in this library.
    result *http.Response
)
err := client.Post(context.Background(), "/unspecified", params, &result)
if err != nil {
    …
}
```

#### Undocumented request params

To make requests using undocumented parameters, you may use either the `option.WithQuerySet()`
or the `option.WithJSONSet()` methods.

```go
params := FooNewParams{
    ID:   "id_xxxx",
    Data: FooNewParamsData{
        FirstName: oursprivacy.String("John"),
    },
}
client.Foo.New(context.Background(), params, option.WithJSONSet("data.last_name", "Doe"))
```

#### Undocumented response properties

To access undocumented response properties, you may either access the raw JSON of the response as a string
with `result.JSON.RawJSON()`, or get the raw JSON of a particular field on the result with
`result.JSON.Foo.Raw()`.

Any fields that are not present on the response struct will be saved and can be accessed by `result.JSON.ExtraFields()` which returns the extra fields as a `map[string]Field`.

### Middleware

We provide `option.WithMiddleware` which applies the given
middleware to requests.

```go
func Logger(req *http.Request, next option.MiddlewareNext) (res *http.Response, err error) {
	// Before the request
	start := time.Now()
	LogReq(req)

	// Forward the request to the next handler
	res, err = next(req)

	// Handle stuff after the request
	end := time.Now()
	LogRes(res, err, start - end)

    return res, err
}

client := oursprivacy.NewClient(
	option.WithMiddleware(Logger),
)
```

When multiple middlewares are provided as variadic arguments, the middlewares
are applied left to right. If `option.WithMiddleware` is given
multiple times, for example first in the client then the method, the
middleware in the client will run first and the middleware given in the method
will run next.

You may also replace the default `http.Client` with
`option.WithHTTPClient(client)`. Only one http client is
accepted (this overwrites any previous client) and receives requests after any
middleware has been applied.

## Semantic versioning

This package generally follows [SemVer](https://semver.org/spec/v2.0.0.html) conventions, though certain backwards-incompatible changes may be released as minor versions:

1. Changes to library internals which are technically public but not intended or documented for external use. _(Please open a GitHub issue to let us know if you are relying on such internals.)_
2. Changes that we do not expect to impact the vast majority of users in practice.

We take backwards-compatibility seriously and work hard to ensure you can rely on a smooth upgrade experience.

We are keen for your feedback; please open an [issue](https://www.github.com/with-ours/ingest-sdk-go/issues) with questions, bugs, or suggestions.

## Contributing

See [the contributing documentation](./CONTRIBUTING.md).
