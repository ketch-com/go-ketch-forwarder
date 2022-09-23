package types

type StandardObject interface {
	GetApiVersion() string
	GetKind() Kind
	GetMetadata() *Metadata
}

type StandardRequest[T any] interface {
	StandardObject
	GetRequest() T
}

type StandardResponse[T any] interface {
	StandardObject
	GetResponse() T
}

type SettableStandardObject[T any] interface {
	SetApiVersion(apiVersion string)
	SetKind(kind Kind)
	SetMetadata(metadata *Metadata)
}

type SettableStandardRequest[T any] interface {
	SettableStandardObject[T]
	SetRequest(request *T)
}

type SettableStandardResponse[T any] interface {
	SettableStandardObject[T]
	SetResponse(response *T)
}
