package plan

import (
	"sigs.k8s.io/yaml"

	"github.com/authgear/authgear-server/pkg/lib/config"
	"github.com/authgear/authgear-server/pkg/lib/config/configsource"
	"github.com/authgear/authgear-server/pkg/lib/infra/db/globaldb"
	"github.com/authgear/authgear-server/pkg/portal/lib/plan"
	"github.com/authgear/authgear-server/pkg/portal/model"
	"github.com/authgear/authgear-server/pkg/util/clock"
)

type Service struct {
	Handle            *globaldb.Handle
	Store             *plan.Store
	ConfigSourceStore *configsource.Store
	Clock             clock.Clock
}

func (s *Service) CreatePlan(name string) error {
	return s.Handle.WithTx(func() (err error) {
		p := model.NewPlan(name)
		return s.Store.Create(p)
	})
}

// UpdatePlan update the plan feature config and also the app which
// have tha same plan name, returns the number of updated app
func (s Service) UpdatePlan(name string, featureConfig *config.FeatureConfig) (appCount int, err error) {
	err = s.Handle.WithTx(func() (err error) {
		p, err := s.Store.GetPlan(name)
		if err != nil {
			return err
		}
		p.RawFeatureConfig = featureConfig
		return s.Store.Update(p)
	})
	if err != nil {
		return
	}

	// update apps feature config
	featureConfigYAML, e := yaml.Marshal(featureConfig)
	if e != nil {
		err = e
		return
	}

	err = s.Handle.WithTx(func() (err error) {
		consrcs, err := s.ConfigSourceStore.ListByPlan(name)
		if err != nil {
			return err
		}
		for _, consrc := range consrcs {
			// json.Marshal handled base64 encoded of the YAML file
			// https://golang.org/pkg/encoding/json/#Marshal
			// Array and slice values encode as JSON arrays,
			// except that []byte encodes as a base64-encoded string,
			// and a nil slice encodes as the null JSON value.
			consrc.Data[configsource.AuthgearFeatureYAML] = featureConfigYAML
			consrc.UpdatedAt = s.Clock.NowUTC()
			err = s.ConfigSourceStore.UpdateDatabaseSource(consrc)
			if err != nil {
				return err
			}
		}
		appCount = len(consrcs)
		return nil
	})
	return
}
