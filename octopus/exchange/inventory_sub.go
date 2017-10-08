package exchange

// Site is the inventory of type of website.
// the inventory has its own attributes so site attributes are passed inside that
// Design note :
// XXX: not supporting search for now
// XXX: not supporting mobile ready
type Site interface {
	Inventory
	// Page the address of page that user see
	Page() string
	// Ref is the referrer of the page
	Ref() string
}

// App is the application version of the inventory
// Design notes:
// XXX : Not supporting store url
// XXX : Not supporting Version and Paid status
type App interface {
	Inventory
	// App bundle or package name
	Bundle() string
}
