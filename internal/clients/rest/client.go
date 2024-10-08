package rest

import (
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/clients/rest/handlers"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/clients/rest/hooks"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/internal/clients/rest/httptransport"
	"github.com/saladtechnologies/salad-cloud-imds-sdk-go/pkg/saladcloudimdssdkconfig"
)

type RestClient[T any] struct {
	handlers *handlers.HandlerChain[T]
}

func NewRestClient[T any](config saladcloudimdssdkconfig.Config) *RestClient[T] {
	retryHandler := handlers.NewRetryHandler[T]()
	responseValidationHandler := handlers.NewResponseValidationHandler[T]()
	unmarshalHandler := handlers.NewUnmarshalHandler[T]()
	requestValidationHandler := handlers.NewRequestValidationHandler[T]()
	hookHandler := handlers.NewHookHandler[T](hooks.NewCustomHook())
	terminatingHandler := handlers.NewTerminatingHandler[T]()

	handlers := handlers.BuildHandlerChain[T]().
		AddHandler(retryHandler).
		AddHandler(responseValidationHandler).
		AddHandler(unmarshalHandler).
		AddHandler(requestValidationHandler).
		AddHandler(hookHandler).
		AddHandler(terminatingHandler)

	return &RestClient[T]{
		handlers: handlers,
	}
}

func (client *RestClient[T]) Call(request httptransport.Request) (*httptransport.Response[T], *httptransport.ErrorResponse[T]) {
	return client.handlers.CallApi(request)
}
