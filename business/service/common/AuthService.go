package common

type AuthService interface {
	CachePerms(userId uint)
}
