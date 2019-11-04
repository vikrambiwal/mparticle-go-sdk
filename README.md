<img src="https://static.mparticle.com/sdk/mp_logo_black.svg" width="280">

## Go Events SDK

This is the mParticle Go SDK for the server-based Events API - use it to send your data to the [mParticle platform](https://www.mparticle.com/) and off to 250+ integrations!

## Getting Started 

### Get the Dependency

```sh
go get github.com/mParticle/mparticle-go-sdk/events
```

### Import the `events` Package

```go
import "github.com/mParticle/mparticle-go-sdk/events"
```

## Create a `Batch`

All data that passes through mParticle does so in the form of a "batch." A batch describes identities, attributes, events, and other information related to a *single user*. This SDK lets you upload either single batches or multiple batches at a time.

The full schema of a batch is documented in the [mParticle Events API overview](https://docs.mparticle.com/developers/server/http/). The models in this SDK directly match the JSON referenced in the overview.

```go
batch := events.Batch{Environment: events.DevelopmentEnvironment} //or "ProductionEnvironment"

//set user identities
batch.UserIdentities = &events.UserIdentities{
    CustomerID: "go1234",
    Email:      "go-example@foo.com",
}

//set device identities
batch.DeviceInfo = &events.DeviceInformation{
    IOSAdvertisingID: "607258d9-c28b-43ad-95ed-e9593025d5a1",
}

//set user attributes
batch.UserAttributes = make(map[string]interface{})
batch.UserAttributes["foo"] = "bar"
batch.UserAttributes["foo-array"] = []string{"bar1", "bar2"}
```

> It's critical to include either user or device identities with your server-side data

## Create Events

All mParticle events have a similar structure:

- `event_type`: this is the type of event, such as `custom_event` and `commerce_event`
- `data`: this contains common properties of all events, as well as properties specific to each `event_type`

The following are common properties that all events share, as represented by the `CommonEventData` struct:

```javascript
{
    "event_id" : 6004677780660780000,
    "source_message_id" : "e8335d31-2031-4bff-afec-17ffc1784697",
    "session_id" : 4957918240501247982,
    "session_uuid" : "91b86d0c-86cb-4124-a8b2-edee107de454",
    "timestamp_unixtime_ms" : 1402521613976,
    "location" : {},
    "device_current_state" : {},
    "custom_attributes": {},
    "custom_flags": {}
}
```

The Go Events SDK represents this structure via an event and an event-data struct for each unique event type. For convenience, every event struct provides a constructor to autopopulate the correct event type and data object, for example `NewCustomEvent()`, `NewScreenViewEvent()`, and `NewCommerceEvent()`:

### Custom Events

```go
customEvent := events.NewCustomEvent()
customEvent.Data.EventName = "My Custom Event Name"
customEvent.Data.CustomEventType = events.OtherCustomEventType
customEvent.Data.CustomAttributes = make(map[string]string)
customEvent.Data.CustomAttributes["foo"] = "bar"
```

### Screen Events

```go
screenEvent := events.NewScreenViewEvent()
screenEvent.Data.ScreenName = "My Screen Name"
```

### Commerce Events

```go
totalProductAmount := 123.12
product := events.Product{
    TotalProductAmount: &totalProductAmount,
    ID:                 "product-id",
    Name:               "product-name",
}
	
commerceEvent := events.NewCommerceEvent()
totalPurchaseAmount := 123.12
commerceEvent.Data.ProductAction = &events.ProductAction{
    Action:        events.PurchaseAction,
    TotalAmount:   &totalProductAmount,
    TransactionID: "foo-transaction-id",
    Products:      []events.Product{product},
}
```

## Full Upload Example

The SDK supports both multi-batch ("bulk") or single-batch uploads:

```go
client := events.NewAPIClient(events.NewConfiguration())

ctx := context.WithValue(
    context.Background(),
    events.ContextBasicAuth,
    events.BasicAuth{
        APIKey:    "REPLACE WITH API KEY",
        APISecret: "REPLACE WITH API SECRET",
    },
)
batch := events.Batch{Environment: events.DevelopmentEnvironment} //or "ProductionEnvironment"

//set user identities
batch.UserIdentities = &events.UserIdentities{
    CustomerID: "go1234",
    Email:      "go-example@foo.com",
}

//set device identities
batch.DeviceInfo = &events.DeviceInformation{
    IOSAdvertisingID: "607258d9-c28b-43ad-95ed-e9593025d5a1",
}

//set user attributes
batch.UserAttributes = make(map[string]interface{})
batch.UserAttributes["foo"] = "bar"
batch.UserAttributes["foo-array"] = []string{"bar1", "bar2"}

customEvent := events.NewCustomEvent()
customEvent.Data.EventName = "My Custom Event Name"
customEvent.Data.CustomEventType = events.OtherCustomEventType
customEvent.Data.CustomAttributes = make(map[string]string)
customEvent.Data.CustomAttributes["foo"] = "bar"

screenEvent := events.NewScreenViewEvent()
screenEvent.Data.ScreenName = "My Screen Name"

totalProductAmount := 123.12
product := events.Product{
    TotalProductAmount: &totalProductAmount,
    ID:                 "product-id",
    Name:               "product-name",
}

commerceEvent := events.NewCommerceEvent()
totalPurchaseAmount := 123.12
commerceEvent.Data.ProductAction = &events.ProductAction{
    Action:        events.PurchaseAction,
    TotalAmount:   &totalPurchaseAmount,
    TransactionID: "foo-transaction-id",
    Products:      []events.Product{product},
}

batch.Events = []events.Event{customEvent, screenEvent, commerceEvent}

result, err := client.EventsAPI.UploadEvents(ctx, batch)
if result != nil && result.StatusCode == 202 {
    fmt.Println("Upload successful")
} else {
    t.Errorf(
        "Error while uploading!\nstatus: %v\nresponse body: %#v",
        err.(events.GenericError).Error(),
        err.(events.GenericError).Model())
}
```