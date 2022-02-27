package teamwork

import (
	"encoding/json"
	"errors"
)

type Option map[string]any

func (o Option) asOptionFunc() func(Option) Option {
	return func(starter Option) Option {
		if starter == nil {
			return o
		}
		for k, v := range o {
			starter[k] = v
		}
		return starter
	}
}

// --- UNEXPORTED HELPERS ---

func applyAttrs[T any](p T, opts Option) error {
	attrBytes, _ := json.Marshal(opts)
	if err := json.Unmarshal(attrBytes, p); err != nil {
		return errors.New("failed to apply options on pointer")
	}
	return nil
}

func resolveAttrs(options ...Option) Option {
	attrs := Option{}
	for _, option := range options {
		attrs = option.asOptionFunc()(attrs)
	}
	return attrs
}

func withOptions[T any](p T, options ...Option) (T, error) {
	attrs := resolveAttrs(options...)
	err := applyAttrs(p, attrs)
	return p, err
}
