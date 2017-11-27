package datamodels

import "sync"

var (
	singleton Aggregator
	lock      sync.RWMutex
)

// Acknowledger is the job to consumer, good parts only
type Acknowledger interface {
	// Ack delegates an acknowledgement through the Acknowledger interface that the client or server has finished work on a delivery.
	Ack(multiple bool) error
	// Nack negatively acknowledge the delivery of message(s) identified by the delivery tag from either the client or server.
	Nack(multiple, requeue bool) error
	// Reject delegates a negatively acknowledgement through the Acknowledger interface.
	Reject(requeue bool) error
}

// TableModel is the model for counting data and aggregate them into on query
type TableModel struct {
	Supplier string // All
	Source   string // All
	Demand   string // All
	Time     int64  // All

	//impression Job : 1 for every bidRequest supplier hit exchange
	SupRequestIn int64
	//impression Job : count of seatBids supplier send to exchange
	SupAdIn int64
	//winner job : number of win ads *1
	SupAdOut int64
	//winner job : price the ad won
	SupBidOut float64
	//show job : number of ad shown *1
	SupAdDeliver int64
	//show job : price of won ad
	SupBidDeliver float64

	// demand Job : total demand bid send to exchange
	SupDemBidIn float64
	// demand job : 1 the bidRequest send to demand
	SupDemRequestOut int64
	// demand Job: the number of bid seat exchange send to demand
	SupDemAdOut int64
	// winner job : the winner price ad won
	SupDemBidWin float64
	// winner job:  the count of demand ad won
	SupDemAdWin int64
	// demand Job :the number of seat bid demand provide for exchange
	SupDemAdIn int64
	// show Job : the total ad count delivered to publisher
	SupDemAdDeliver int64
	// show job : the total bid price delivered to publisher
	SupDemBidDeliver float64

	Profit       float64 //show
	Click        int64
	Acknowledger Acknowledger
	WorkerID     string
}

// Aggregator is a helper type to handle the entire process, and hey, its mock-able!
type Aggregator interface {
	Channel() chan<- TableModel
}

// RegisterAggregator to register an aggregator
func RegisterAggregator(a Aggregator) {
	lock.Lock()
	defer lock.Unlock()

	singleton = a
}

// ActiveAggregator return the current aggregator
func ActiveAggregator() Aggregator {
	lock.RLock()
	defer lock.RUnlock()

	return singleton
}
