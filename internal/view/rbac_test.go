package view_test

import (
	"testing"

	"github.com/derailed/k9s/internal/config"
	"github.com/derailed/k9s/internal/view"
	"github.com/stretchr/testify/assert"
)

func TestRbacNew(t *testing.T) {
	cfg := config.NewConfig(ks{})
	app := view.NewApp(cfg)
	v := view.NewRbac(app, "", "fred", view.ClusterRole)
	v.Init(makeCtx())

	assert.Equal(t, "Rbac", v.Name())
	assert.Equal(t, 10, len(v.Hints()))
}