package follow

type Follow struct {
	ID            int    `json:"id"`
	FollowerID    int    `json:"followerId"`
	FollowerName  string `json:"followerName"`
	FollowingID   int    `json:"followingId"`
	FollowingName string `json:"followingName"`
	Status        bool   `json:"status"`
}
