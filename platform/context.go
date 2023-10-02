package platform

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/Aibar01/platform/attrs"
	"github.com/Aibar01/platform/auth"
	"net/http"
	"sync"
)

type PlatformContext struct {
	contextAttrs attrs.ContextAttrs
	services     map[string]*Service
	User         auth.UserInterface
	Consumer     *auth.Consumer
}

func (p *PlatformContext) Service(serviceName string) *Service {
	if p.services == nil {
		p.services = make(map[string]*Service)
	}

	if service, ok := p.services[serviceName]; ok {
		return service
	}

	service := NewService(serviceName, p.contextAttrs)
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	p.services[serviceName] = service

	return service
}

func (p *PlatformContext) Context() map[string]any {
	data, err := json.Marshal(p.contextAttrs)

	if err != nil {
		fmt.Println(err)
	}

	decodedData := make(map[string]any)
	err = json.Unmarshal(data, &decodedData)

	if err != nil {
		fmt.Println(err)
	}

	return decodedData
}

func (p *PlatformContext) Validate(header http.Header) error {
	contextAttrs := map[string]bool{
		"X-Correlation-Id": true,
		"Authorization":    true,
		"X-Authorization":  false,
		"Accept-Language":  true,
	}
	data := make(map[string]string)

	for attr, required := range contextAttrs {
		if attr == "Authorization" {
			user, err := auth.ExtractUserFromToken(header.Get(attr))
			if err != nil {
				return err
			}
			p.User = user
			consumer, err := auth.NewConsumerFromToken(header.Get(attr))
			if err != nil {
				return err
			}
			p.Consumer = consumer
		} else if header.Get(attr) == "" && required {
			return errors.New(fmt.Sprintf("%s must be in the request", attr))
		}
		data[attr] = header.Get(attr)
	}

	jsonData, err := json.Marshal(data)

	if err != nil {
		return err
	}

	if err = json.Unmarshal(jsonData, &p.contextAttrs); err != nil {
		return err
	}

	return nil
}
