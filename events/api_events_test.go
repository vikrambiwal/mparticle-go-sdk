package events

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http/httptrace"
	"testing"
	"time"
)

func TestFullBatch(t *testing.T) {
	client := NewAPIClient(NewConfiguration())

	ctx := context.WithValue(
		context.Background(),
		ContextBasicAuth,
		BasicAuth{
			APIKey:    "REPLACE WITH KEY",
			APISecret: "REPLACE WITH SECRET",
		},
	)

	batch := Batch{Environment: DevelopmentEnvironment} //or "ProductionEnvironment"
	//set context
	batch.BatchContext = &BatchContext{
		//configure data plan
		DataPlan: &DataPlanContext{
			PlanID:      "freddy_s_plan",
			PlanVersion: 1,
		},
		//configure batching
		Batching: &BatchingContext{
			Bypass:       true,
			PartitionKey: "6eea30aa-6ccc-4b30-a691-3afa211d54ae",
			RequestIDs:   []string{"foo", "bar"},
		},
	}

	//set user identities
	batch.UserIdentities = &UserIdentities{
		CustomerID: "go1234",
		Email:      "go-example@foo.com",
	}

	//set device identities
	batch.DeviceInfo = &DeviceInformation{
		IOSAdvertisingID: "607258d9-c28b-43ad-95ed-e9593025d5a1",
	}

	//set user attributes
	batch.UserAttributes = make(map[string]interface{})
	batch.UserAttributes["foo"] = "bar"
	batch.UserAttributes["foo-array"] = []string{"bar1", "bar2"}

	customEvent := NewCustomEvent()
	customEvent.Data.EventName = "My Custom Event Name"
	customEvent.Data.CustomEventType = OtherCustomEventType
	customEvent.Data.CustomAttributes = make(map[string]string)
	customEvent.Data.CustomAttributes["foo"] = "bar"

	screenEvent := NewScreenViewEvent()
	screenEvent.Data.ScreenName = "My Screen Name"

	sessionStartEvent := NewSessionStartEvent()
	sessionStartEvent.Data.SessionID = 12345678
	sessionStartEvent.Data.TimestampUnixtimeMS = int64(time.Now().Unix())

	sessionEndEvent := NewSessionEndEvent()
	sessionEndEvent.Data.SessionID = 12345678
	sessionEndEvent.Data.SessionDurationMS = int64(time.Now().Unix())

	gdprConsentState := GdprConsentState{
		Document:            "document_agreement.v2",
		Consented:           true,
		TimestampUnixtimeMS: int64(time.Now().Unix()),
	}

	consentState := &ConsentState{
		GDPR: make(map[string]GdprConsentState),
	}
	consentState.GDPR["performance"] = gdprConsentState
	//batch.ConsentState = consentState

	totalProductAmount := 10.00
	totalProducts := int32(2)
	product := Product{
		TotalProductAmount: &totalProductAmount,
		ID:                 "product-id",
		Name:               "product-name",
		Quantity:           &totalProducts,
	}

	commerceEvent := NewCommerceEvent()
	totalPurchaseAmount := 123.12
	commerceEvent.Data.ProductAction = &ProductAction{
		Action:        PurchaseAction,
		TotalAmount:   &totalPurchaseAmount,
		TransactionID: "foo-transaction-id",
		Products:      []Product{product},
	}

	batch.Events = []Event{customEvent, screenEvent, commerceEvent}

	trace := &httptrace.ClientTrace{
		DNSDone: func(dnsInfo httptrace.DNSDoneInfo) {
			fmt.Printf("DNS Info: %+v\n", dnsInfo)
		},
		GotConn: func(connInfo httptrace.GotConnInfo) {
			fmt.Printf("Got Conn: %+v\n", connInfo)
		},
	}

	out, err := json.Marshal(batch)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(out))
	newCtx := httptrace.WithClientTrace(ctx, trace)
	result, err := client.EventsAPI.UploadEvents(newCtx, batch)
	if result != nil && result.StatusCode == 202 {
		fmt.Println("Upload successful")
	} else {
		t.Errorf(
			"Error while uploading!\nstatus: %v\nresponse body: %#v",
			err.(GenericError).Error(),
			err.(GenericError).Model())
	}
}
