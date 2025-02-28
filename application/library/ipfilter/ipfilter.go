package ipfilter

import (
	"net/netip"

	"github.com/webx-top/echo"

	"github.com/coscms/webcore/library/ipfilter"
	"github.com/nging-plugins/frpmanager/application/dbschema"
)

var Factory = ipfilter.NewFactory[uint64, uint]()

func IsAllowed(ctx echo.Context, user *dbschema.NgingFrpUser, ip netip.Addr) bool {
	return Factory.IsAllowed(ctx, user.Id, 0, user.IpBlacklist, user.IpWhitelist, ip)
}

func Validate(ctx echo.Context, ip string) error {
	return ipfilter.ValidateRows(ctx, ip)
}
