package helmify

import (
	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	"io"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

// Processor - converts k8s object to helm template
type Processor interface {
	// Process - converts k8s object to Helm template.
	// return false if not able to process given object type
	Process(chartInfo ChartInfo, unstructured *unstructured.Unstructured) (bool, Template, error)
}

// Template - represents Helm template in 'templates' directory
type Template interface {
	// Filename - returns template filename
	Filename() string
	// Values - returns set of values used in template
	Values() Values
	// Write - writes helm template into given writer
	Write(writer io.Writer) error
}

// Values - represents helm template values.yaml
type Values map[string]interface{}

func (v *Values) Merge(values Values) error {
	if err := mergo.Merge(v, values, mergo.WithAppendSlice); err != nil {
		return errors.Wrap(err, "unable to merge helm values")
	}
	return nil
}

// Output - converts Template into helm chart on disk
type Output interface {
	Create(chartInfo ChartInfo, templates []Template) error
}

type ChartInfo struct {
	ChartName         string
	OperatorName      string
	OperatorNamespace string
}
