
package domain

type Entity interface {
	GetID() string
}

type AggregateRoot interface {
	Entity
}

type Repository interface {
	GetDB() interface{}
}

type Specification interface {
	IsSatisfiedBy(entity interface{}) bool
}
