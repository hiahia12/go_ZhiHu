package user

type Group struct {
}

func (g *Group) Sign() *SignApi {
	return &insSign
}

func (g *Group) Write() *WriteApi {
	return &insWrite
}

func (g *Group) Like() *LikeApi {
	return &insLike
}

func (g *Group) Follow() *FollowApi {
	return &insFollow
}

func (g *Group) Change() *ChangApi {
	return &insChange
}

func (g *Group) Favourite() *FavouriteApi {
	return &insFavourite
}
