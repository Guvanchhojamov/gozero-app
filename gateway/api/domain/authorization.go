package domain

type Authorization interface {
	ParseToken(accessToken string) (userId uint32, err error)
}
