package store

type Factory interface {
	Links() LinkStore
	LinkTraces() LinkTraceStore
}
