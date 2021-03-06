package scope

import (
	core "github.com/rinnguyen1614/rin-echo-core"
	"github.com/rinnguyen1614/rin-echo-core/setting/adapter"
)

var (
	// Represents a setting that can be configured/changed for each User.
	UserSettingProviderName = "U"
	// Represents a setting that can be configured/changed for the application level.
	GlobalSettingProviderName = "G"
)

type ScopeProvider interface {
	Name() string
	WithContext(ctx core.Context) ScopeProvider
	GetOrInit(name string) string
	GetMulti(names []string) map[string]string
	GetAll() map[string]string
	Set(name, value string) error
	Delete(name string) error
}

type scopeProvider struct {
	name    string
	adapter adapter.Adapter
}

func newScopeProvider(name string, adapter adapter.Adapter) *scopeProvider {
	return &scopeProvider{
		name:    name,
		adapter: adapter,
	}
}

func (s scopeProvider) Name() string {
	return s.name
}
