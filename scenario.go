package gocuke

import "github.com/cucumber/messages-go/v16"

// Scenario is a special step argument type that can be used in a step definition
// or hook method.
type Scenario interface {
	Name() string
	Tags() []string
	private()
}

type scenario struct {
	pickle *messages.Pickle
}

// Name returns the scenario name.
func (s scenario) Name() string {
	return s.pickle.Name
}

// Tags returns the scenario tags.
func (s scenario) Tags() []string {
	tags := make([]string, len(s.pickle.Tags))
	for i, tag := range s.pickle.Tags {
		tags[i] = tag.Name
	}
	return tags
}

func (s scenario) private() {}

var _ Scenario = scenario{}