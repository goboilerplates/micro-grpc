package runner

import (
	"log"

	"github.com/goboilerplates/micro-rpc/example/goclient/request"
)

// RunGetSampleRequests runs all requests of GetSampleRequester.
func RunGetSampleRequests(requester request.GetSampleRequester) {
	log.Println("All = ", requester.All())
	log.Println("AllByName = ", requester.AllByName())

	log.Println("ServerStreamAll = ", requester.ServerStreamAll())
	log.Println("ServerStreamAllByName = ", requester.ServerStreamAllByName())

	log.Println("ClientStreamAll = ", requester.ClientStreamAll())
	log.Println("ClientStreamAllByName = ", requester.ClientStreamAllByName())

	log.Println("StreamAll = ", requester.StreamAll())
	log.Println("StreamAllByName = ", requester.StreamAllByName())
}
