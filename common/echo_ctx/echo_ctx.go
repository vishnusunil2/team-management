package echo_ctx

import "github.com/labstack/echo/v4"

const (
	User = "user"
	Org  = "org"
)

type UserProfileCtx struct {
	UserID string
	OrgId  string
}

func NewUserProfileCtx(userID, orgID string) *UserProfileCtx {
	return &UserProfileCtx{UserID: userID, OrgId: orgID}
}
func GetOrgId(ctx echo.Context) string {
	return GetUserProfile(ctx).OrgId
}
func SetOrgId(ctx echo.Context, orgId string) {
	ctx.Set(Org, orgId)
}
func SetUserProfile(ctx echo.Context, userProfile *UserProfileCtx) {
	ctx.Set(User, userProfile)
	if userProfile.OrgId != "" {
		SetOrgId(ctx, userProfile.OrgId)
	}
}
func GetUserProfile(ctx echo.Context) *UserProfileCtx {
	{
		return ctx.Get(User).(*UserProfileCtx)
	}
}

func GetUserId(ctx echo.Context) string { return GetUserProfile(ctx).UserID }
