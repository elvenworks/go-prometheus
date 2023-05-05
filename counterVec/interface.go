package counterVec

type ICounterVec interface {
	CreateHit(orgUid string)
	CreateFailure(orgUid string)
}
