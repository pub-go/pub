package dto

import "code.gopub.tech/pub/dal/model"

type LoginInfo struct {
	Username  string
	UserAgent string
}

type UserInfo struct {
	*model.User
	Roles []string
	roles map[string]struct{}
}

func (u *UserInfo) HasRole(role string) bool {
	if len(u.roles) != len(u.Roles) {
		m := make(map[string]struct{}, len(u.Roles))
		for _, r := range u.Roles {
			m[r] = struct{}{}
		}
		u.roles = m
	}
	_, ok := u.roles[role]
	return ok
}
