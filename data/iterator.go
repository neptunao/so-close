package data

// Iterator is generic iterator for all kinds of separated data
type Iterator interface {
	Next() (interface{}, bool)
	Err() error
	Close() error
}

// SourceConnector provides a way to get iterator for data source
type SourceConnector interface {
	Connect() (Iterator, error)
}
