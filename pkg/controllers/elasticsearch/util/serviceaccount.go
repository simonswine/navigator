package util

import (
	"github.com/jetstack/navigator/pkg/apis/navigator/v1alpha1"
)

func ServiceAccountName(c *v1alpha1.ElasticsearchCluster) string {
	return ResourceBaseName(c)
}
