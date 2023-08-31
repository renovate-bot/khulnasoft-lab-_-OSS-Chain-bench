package dependencies

import (
	thirdpartypackages "github.com/khulnasoft-lab/oss-chain-bench/internal/checks/dependencies/third-party-packages"
	validatepackages "github.com/khulnasoft-lab/oss-chain-bench/internal/checks/dependencies/validate_packages"
	"github.com/khulnasoft-lab/oss-chain-bench/internal/models/checkmodels"
)

func GetChecks() []*checkmodels.Check {
	checks := []*checkmodels.Check{}
	checks = append(checks, thirdpartypackages.GetChecks()...)
	checks = append(checks, validatepackages.GetChecks()...)
	return checks
}
