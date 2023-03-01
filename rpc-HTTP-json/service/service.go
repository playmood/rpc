package service

const (
	SERVICE_NAME = "HelloService"
)

type Sum struct {
	A int
	B int
}

type HelloService interface {
	Hello(request string, response *string) error
	Calc(request Sum, response *int) error
}
