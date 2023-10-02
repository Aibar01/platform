package platform

import (
	"fmt"
	"github.com/Aibar01/platform/attrs"
	"github.com/Aibar01/platform/client"
)

type Service struct {
	*client.Client
}

func NewService(serviceName string, contextAttrs attrs.ContextAttrs) *Service {
	baseUrl := fmt.Sprintf("http://%s:8000", serviceName)

	return &Service{
		Client: client.NewClient(baseUrl, contextAttrs),
	}
}
