package admin

import (
	"fmt"
	"path"
	"time"

	"github.com/GoogleCloudPlatform/kubernetes/pkg/util"

	configapi "github.com/openshift/origin/pkg/cmd/server/api"
	"github.com/openshift/origin/pkg/cmd/server/bootstrappolicy"
)

const (
	CAFilePrefix     = "ca"
	MasterFilePrefix = "master"
)

type ClientCertInfo struct {
	CertLocation    configapi.CertInfo
	UnqualifiedUser string
	User            string
	Groups          util.StringSet
}

func DefaultSignerName() string {
	return fmt.Sprintf("%s@%d", "openshift-signer", time.Now().Unix())
}

func DefaultRootCAFile(certDir string) string {
	return DefaultCertFilename(certDir, CAFilePrefix)
}

func DefaultKubeletClientCAFile(certDir string) string {
	return DefaultRootCAFile(certDir)
}

func DefaultKubeletClientCerts(certDir string) []ClientCertInfo {
	return []ClientCertInfo{
		DefaultMasterKubeletClientCertInfo(certDir),
	}
}

func DefaultMasterKubeletClientCertInfo(certDir string) ClientCertInfo {
	return ClientCertInfo{
		CertLocation: configapi.CertInfo{
			CertFile: path.Join(certDir, MasterFilePrefix+".kubelet-client.crt"),
			KeyFile:  path.Join(certDir, MasterFilePrefix+".kubelet-client.key"),
		},
		User: "system:master",
	}
}

func DefaultEtcdClientCAFile(certDir string) string {
	return DefaultRootCAFile(certDir)
}

func DefaultEtcdClientCerts(certDir string) []ClientCertInfo {
	return []ClientCertInfo{
		DefaultMasterEtcdClientCertInfo(certDir),
	}
}

func DefaultMasterEtcdClientCertInfo(certDir string) ClientCertInfo {
	return ClientCertInfo{
		CertLocation: configapi.CertInfo{
			CertFile: path.Join(certDir, MasterFilePrefix+".etcd-client.crt"),
			KeyFile:  path.Join(certDir, MasterFilePrefix+".etcd-client.key"),
		},
		User: "system:master",
	}
}

func DefaultAPIClientCAFile(certDir string) string {
	return DefaultRootCAFile(certDir)
}

func DefaultAPIClientCerts(certDir string) []ClientCertInfo {
	return []ClientCertInfo{
		DefaultDeployerClientCertInfo(certDir),
		DefaultOpenshiftLoopbackClientCertInfo(certDir),
		DefaultKubeClientClientCertInfo(certDir),
		DefaultClusterAdminClientCertInfo(certDir),
		DefaultRouterClientCertInfo(certDir),
		DefaultRegistryClientCertInfo(certDir),
	}
}

func DefaultRouterClientCertInfo(certDir string) ClientCertInfo {
	return ClientCertInfo{
		CertLocation: configapi.CertInfo{
			CertFile: DefaultCertFilename(certDir, bootstrappolicy.RouterUnqualifiedUsername),
			KeyFile:  DefaultKeyFilename(certDir, bootstrappolicy.RouterUnqualifiedUsername),
		},
		UnqualifiedUser: bootstrappolicy.RouterUnqualifiedUsername,
		User:            bootstrappolicy.RouterUsername,
		Groups:          util.NewStringSet(bootstrappolicy.RouterGroup),
	}
}

func DefaultRegistryClientCertInfo(certDir string) ClientCertInfo {
	return ClientCertInfo{
		CertLocation: configapi.CertInfo{
			CertFile: DefaultCertFilename(certDir, bootstrappolicy.RegistryUnqualifiedUsername),
			KeyFile:  DefaultKeyFilename(certDir, bootstrappolicy.RegistryUnqualifiedUsername),
		},
		UnqualifiedUser: bootstrappolicy.RegistryUnqualifiedUsername,
		User:            bootstrappolicy.RegistryUsername,
		Groups:          util.NewStringSet(bootstrappolicy.RegistryGroup),
	}
}

func DefaultDeployerClientCertInfo(certDir string) ClientCertInfo {
	return ClientCertInfo{
		CertLocation: configapi.CertInfo{
			CertFile: DefaultCertFilename(certDir, "openshift-deployer"),
			KeyFile:  DefaultKeyFilename(certDir, "openshift-deployer"),
		},
		UnqualifiedUser: "openshift-deployer",
		User:            "system:openshift-deployer",
		Groups:          util.NewStringSet("system:deployers"),
	}
}

func DefaultOpenshiftLoopbackClientCertInfo(certDir string) ClientCertInfo {
	return ClientCertInfo{
		CertLocation: configapi.CertInfo{
			CertFile: DefaultCertFilename(certDir, "openshift-client"),
			KeyFile:  DefaultKeyFilename(certDir, "openshift-client"),
		},
		UnqualifiedUser: "openshift-client",
		User:            "system:openshift-client",
	}
}

func DefaultKubeClientClientCertInfo(certDir string) ClientCertInfo {
	return ClientCertInfo{
		CertLocation: configapi.CertInfo{
			CertFile: DefaultCertFilename(certDir, "kube-client"),
			KeyFile:  DefaultKeyFilename(certDir, "kube-client"),
		},
		UnqualifiedUser: "kube-client",
		User:            "system:kube-client",
	}
}

func DefaultClusterAdminClientCertInfo(certDir string) ClientCertInfo {
	return ClientCertInfo{
		CertLocation: configapi.CertInfo{
			CertFile: DefaultCertFilename(certDir, "admin"),
			KeyFile:  DefaultKeyFilename(certDir, "admin"),
		},
		UnqualifiedUser: "admin",
		User:            "system:admin",
		Groups:          util.NewStringSet("system:cluster-admins"),
	}
}

func DefaultServerCerts(certDir string) []configapi.CertInfo {
	return []configapi.CertInfo{
		DefaultMasterServingCertInfo(certDir),
		DefaultAssetServingCertInfo(certDir),
		DefaultEtcdServingCertInfo(certDir),
	}
}

func DefaultMasterServingCertInfo(certDir string) configapi.CertInfo {
	return configapi.CertInfo{
		CertFile: path.Join(certDir, MasterFilePrefix+".server.crt"),
		KeyFile:  path.Join(certDir, MasterFilePrefix+".server.key"),
	}
}

func DefaultAssetServingCertInfo(certDir string) configapi.CertInfo {
	// Use master certs for assets also
	return DefaultMasterServingCertInfo(certDir)
}

func DefaultEtcdServingCertInfo(certDir string) configapi.CertInfo {
	return configapi.CertInfo{
		CertFile: path.Join(certDir, "etcd.server.crt"),
		KeyFile:  path.Join(certDir, "etcd.server.key"),
	}
}

func DefaultNodeDir(nodeName string) string {
	return "node-" + nodeName
}

func DefaultNodeServingCertInfo(nodeDir string) configapi.CertInfo {
	return configapi.CertInfo{
		CertFile: path.Join(nodeDir, "server.crt"),
		KeyFile:  path.Join(nodeDir, "server.key"),
	}
}
func DefaultNodeClientCertInfo(nodeDir string) configapi.CertInfo {
	return configapi.CertInfo{
		CertFile: path.Join(nodeDir, "master-client.crt"),
		KeyFile:  path.Join(nodeDir, "master-client.key"),
	}
}
func DefaultNodeKubeConfigFile(nodeDir string) string {
	return path.Join(nodeDir, "node.kubeconfig")
}

func DefaultCAFilename(certDir, prefix string) string {
	return path.Join(certDir, prefix+".crt")
}
func DefaultCertFilename(certDir, prefix string) string {
	return path.Join(certDir, prefix+".crt")
}
func DefaultKeyFilename(certDir, prefix string) string {
	return path.Join(certDir, prefix+".key")
}
func DefaultSerialFilename(certDir, prefix string) string {
	return path.Join(certDir, prefix+".serial.txt")
}
func DefaultKubeConfigFilename(certDir, prefix string) string {
	return path.Join(certDir, prefix+".kubeconfig")
}
