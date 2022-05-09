package discovery

import (
	"context"
	"testing"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc-b/resolver"
	"google.golang.org/grpc-b/serviceconfig"
)

type mockLogger struct {
	level log.Level
	key   string
	val   string
}

func (l *mockLogger) Log(level log.Level, keyvals ...interface{}) error {
	l.level = level
	l.key = keyvals[0].(string)
	l.val = keyvals[1].(string)
	return nil
}

func TestWithLogger(t *testing.T) {
	b := &builder{}
	WithLogger(&mockLogger{})(b)
}

func TestWithInsecure(t *testing.T) {
	b := &builder{}
	WithInsecure(true)(b)
	assert.True(t, b.insecure)
}

func TestWithTimeout(t *testing.T) {
	o := &builder{}
	v := time.Duration(123)
	WithTimeout(v)(o)
	assert.Equal(t, v, o.timeout)
}

type mockDiscovery struct{}

func (m *mockDiscovery) GetService(ctx context.Context, serviceName string) ([]*registry.ServiceInstance, error) {
	return nil, nil
}

func (m *mockDiscovery) Watch(ctx context.Context, serviceName string) (registry.Watcher, error) {
	return &testWatch{}, nil
}

func TestBuilder_Scheme(t *testing.T) {
	b := NewBuilder(&mockDiscovery{})
	assert.Equal(t, "discovery", b.Scheme())
}

type mockConn struct{}

func (m *mockConn) UpdateState(resolver.State) error {
	return nil
}

func (m *mockConn) ReportError(error) {}

func (m *mockConn) NewAddress(addresses []resolver.Address) {}

func (m *mockConn) NewServiceConfig(serviceConfig string) {}

func (m *mockConn) ParseServiceConfig(serviceConfigJSON string) *serviceconfig.ParseResult {
	return nil
}

func TestBuilder_Build(t *testing.T) {
	b := NewBuilder(&mockDiscovery{})
	_, err := b.Build(resolver.Target{Scheme: resolver.GetDefaultScheme(), Endpoint: "gprc://authority/endpoint"}, &mockConn{}, resolver.BuildOptions{})
	assert.NoError(t, err)
}
