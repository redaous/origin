package install

import (
	"k8s.io/apimachinery/pkg/runtime"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"

	"github.com/openshift/origin/pkg/service/admission/apis/externalipranger"
	"github.com/openshift/origin/pkg/service/admission/apis/externalipranger/v1"
)

func InstallLegacyInternal(scheme *runtime.Scheme) {
	utilruntime.Must(externalipranger.Install(scheme))
	utilruntime.Must(v1.Install(scheme))
	utilruntime.Must(externalipranger.DeprecatedInstallLegacy(scheme))
	utilruntime.Must(v1.DeprecatedInstall(scheme))
}
