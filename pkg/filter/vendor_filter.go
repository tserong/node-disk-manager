package filter

import (
	"github.com/harvester/node-disk-manager/pkg/block"
	"github.com/harvester/node-disk-manager/pkg/utils"
)

const (
	vendorFilterName            = "vendor filter"
	vendorFilterDefaultLonghorn = "longhorn"
)

var (
	defaultExcludedVendors = []string{vendorFilterDefaultLonghorn}
)

type vendorFilter struct {
	vendors []string
}

func RegisterVendorFilter(filters ...string) *Filter {
	vf := &vendorFilter{}
	for _, filter := range filters {
		if filter != "" {
			vf.vendors = append(vf.vendors, filter)
		}
	}
	return &Filter{
		Name:       vendorFilterName,
		DiskFilter: vf,
	}
}

// Match returns true if vendor of the disk is matched
func (vf *vendorFilter) Match(blockDevice *block.Disk) bool {
	if blockDevice.Vendor != "" && utils.MatchesIgnoredCase(vf.vendors, blockDevice.Vendor) {
		return true
	}
	// HA! longhorn also appears here (e.g.: E:ID_PATH=ip-10.52.0.21:3260-iscsi-iqn.2019-10.io.longhorn:pvc-c781f54b-3774-4a67-bdca-d3465383cc60-lun-1)
	return blockDevice.BusPath != "" && utils.ContainsIgnoredCase(vf.vendors, blockDevice.BusPath)
}
