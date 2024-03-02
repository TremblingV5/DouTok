package services

type etcdConfig interface {
	GetAddr() string
}

type baseConfig interface {
	GetAddr() string
	GetName() string
}

type otelConfig interface {
	IsEnable() bool
	GetAddr() string
}
