package reqs

type InstallReq struct {
	SiteTitle string `form:"site_title" binding:"required"`
	Username  string `form:"username"  binding:"required"`
	Email     string `form:"email"  binding:"required"`
	Password  string `form:"password"  binding:"required,len=64"`
	Salt      string `form:"salt"  binding:"required"`
}

type LoginReq struct {
	Username string `form:"username"  binding:"required"`
	Password string `form:"password"  binding:"required,len=64"`
	Salt     string `form:"salt"  binding:"required"`
}
