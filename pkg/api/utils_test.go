package api

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/projectsyn/lieutenant-operator/pkg/apis/syn/v1alpha1"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/AlekSi/pointer"
)

func TestRepoConversionDefaultAuto(t *testing.T) {
	apiRepo := &GitRepo{
		Type: nil,
		Url:  nil,
	}
	repoName := "c-dshfjuhrtu"
	repoTemplate := newGitRepoTemplate(apiRepo, repoName)
	assert.Nil(t, repoTemplate)
	assert.Nil(t, newGitRepoTemplate(nil, repoName))
}

func TestRepoConversionUnmanagedo(t *testing.T) {
	apiRepo := &GitRepo{
		Type: pointer.ToString("unmanaged"),
		Url:  pointer.ToString("ssh://git@some.host/path/to/repo.git"),
	}
	repoTemplate := newGitRepoTemplate(apiRepo, "some-name")
	assert.Empty(t, repoTemplate.RepoName)
	assert.Empty(t, repoTemplate.Path)
}

func TestRepoConversionSpecSubGroupPath(t *testing.T) {
	repoName := "myName"
	repoPath := "path/to"
	apiRepo := &GitRepo{
		Type: pointer.ToString("auto"),
		Url:  pointer.ToString("ssh://git@some.host/" + repoPath + "/" + repoName + ".git"),
	}
	repoTemplate := newGitRepoTemplate(apiRepo, "some-name")
	assert.Equal(t, repoName, repoTemplate.RepoName)
	assert.Equal(t, repoPath, repoTemplate.Path)
	assert.Empty(t, repoTemplate.APISecretRef.Name)
}

func TestRepoConversionSpecPath(t *testing.T) {
	repoName := "myName"
	repoPath := "path"
	apiRepo := &GitRepo{
		Type: pointer.ToString("auto"),
		Url:  pointer.ToString("ssh://git@some.host/" + repoPath + "/" + repoName + ".git"),
	}
	repoTemplate := newGitRepoTemplate(apiRepo, "some-name")
	assert.Equal(t, repoName, repoTemplate.RepoName)
	assert.Equal(t, repoPath, repoTemplate.Path)
	assert.Empty(t, repoTemplate.APISecretRef.Name)
}

func TestRepoConversionFail(t *testing.T) {
	apiRepo := &GitRepo{
		Url: pointer.ToString("://git@some.host/group/example.git"),
	}
	repoTemplate := newGitRepoTemplate(apiRepo, "some-name")
	assert.Nil(t, repoTemplate)

	repoTemplate = newGitRepoTemplate(&GitRepo{
		Url: pointer.ToString("ssh://git@some.host/example.git"),
	}, "test")
	assert.Nil(t, repoTemplate)
}

func TestGenerateClusterID(t *testing.T) {
	assertGeneratedID(t, ClusterIDPrefix, func() (s string) {
		id, err := GenerateClusterID()
		require.NoError(t, err)
		return string(id.Id)
	})
}

func TestGenerateTenantID(t *testing.T) {
	assertGeneratedID(t, TenantIDPrefix, func() (s string) {
		id, err := GenerateTenantID()
		require.NoError(t, err)
		return id.Id.String()
	})
}

func assertGeneratedID(t *testing.T, prefix string, supplier func() string) {
	// Verify generated ID so that it conforms to https: //kubernetes.io/docs/concepts/overview/working-with-objects/names/#names
	// Regex pattern tested on regexr.com
	r := regexp.MustCompile("^[a-z]-[a-z0-9]{3,}(-|_)[a-z0-9]{3,}(-|_)[0-9]+$")
	// Run the randomizer a few times
	for i := 1; i <= 1000; i++ {
		id := supplier()
		require.LessOrEqualf(t, len(id), 63, "Iteration %d: too long for a DNS-compatible name: %s", i, id)
		require.Regexpf(t, r, id, "Iteration %d: not in the form of 'adjective-noun-number' %s", i, id)
		require.True(t, strings.HasPrefix(id, prefix))
	}
}

var tenantTests = map[string]struct{
  properties TenantProperties
  spec v1alpha1.TenantSpec
}{
	"empty": {
		TenantProperties{},
		v1alpha1.TenantSpec{},
	},
	"global git URL": {
		TenantProperties{
			GlobalGitRepoURL: pointer.ToString("ssh://git@example.com/foo/bar.git"),
		},
		v1alpha1.TenantSpec{
			GlobalGitRepoURL: "ssh://git@example.com/foo/bar.git",
		},
	},
	"global git revision": {
		TenantProperties{
			GlobalGitRepoRevision: pointer.ToString("v1.2.3"),
		},
		v1alpha1.TenantSpec{
			GlobalGitRepoRevision: "v1.2.3",
		},
	},
	"git revision": {
		TenantProperties{
			GitRepo: &RevisionedGitRepo{Revision: Revision{pointer.ToString("v1.2.3")}},
		},
		v1alpha1.TenantSpec{
			GitRepoRevision: "v1.2.3",
		},
	},
}

func TestNewCRDFromAPITenant(t *testing.T) {
	for name, test := range tenantTests {
		t.Run(name, func(t *testing.T) {
			apiTenant := Tenant{
				TenantId{
					Id: Id(fmt.Sprintf("t-%s", t.Name())),
				},
				test.properties,
			}
			tenant := NewCRDFromAPITenant(apiTenant)
			assert.Equal(t, test.spec, tenant.Spec)
		})
	}
}

func TestNewAPITenantFromCRD(t *testing.T) {
	for name, test := range tenantTests {
		t.Run(name, func(t *testing.T) {
			tenant := v1alpha1.Tenant{
				Spec: test.spec,
			}
			apiTenant := NewAPITenantFromCRD(tenant)
			if test.properties.GitRepo == nil {
				test.properties.GitRepo = &RevisionedGitRepo{}
			}
			assert.Equal(t, test.properties, apiTenant.TenantProperties)
		})
	}
}

var clusterTests = map[string]struct{
	properties ClusterProperties
	spec v1alpha1.ClusterSpec
}{
	"empty": {
		ClusterProperties{},
		v1alpha1.ClusterSpec{},
	},
	"global git revision": {
		ClusterProperties{
			GlobalGitRepoRevision: pointer.ToString("v1.2.3"),
		},
		v1alpha1.ClusterSpec{
			GlobalGitRepoRevision: "v1.2.3",
		},
	},
	"tenant git revision": {
		ClusterProperties{
			TenantGitRepoRevision: pointer.ToString("v1.2.3"),
		},
		v1alpha1.ClusterSpec{
			TenantGitRepoRevision: "v1.2.3",
		},
	},
}

func TestNewCRDFromAPICluster(t *testing.T) {
	for name, test := range clusterTests {
		t.Run(name, func(t *testing.T) {
			apiCluster := Cluster{
				ClusterId{
					Id: Id(fmt.Sprintf("c-%s", t.Name())),
				},
				ClusterTenant{fmt.Sprintf("t-%s", t.Name())},
				test.properties,
			}
			cluster := NewCRDFromAPICluster(apiCluster)
			if len(test.spec.TenantRef.Name) == 0 {
				test.spec.TenantRef.Name = fmt.Sprintf("t-%s", t.Name())
			}
			assert.Equal(t, test.spec, cluster.Spec)
		})
	}
}

func TestNewAPIClusterFromCRD(t *testing.T) {
	for name, test := range clusterTests {
		t.Run(name, func(t *testing.T) {
			cluster := v1alpha1.Cluster{
				Spec: test.spec,
			}
			apiCluster := NewAPIClusterFromCRD(cluster)
			if test.properties.GitRepo == nil {
				test.properties.GitRepo = &GitRepo{}
			}
			assert.Equal(t, test.properties, apiCluster.ClusterProperties)
		})
	}
}
