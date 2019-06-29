package roles

var fenLinWan = Role{
	maxRunFrames 	  : 8,
	maxStandFrames    : 3,
	maxAttackFrames   : 6,
	maxJumpFrames     : 10,
	runFrameWidth 	  : 66,
	standFrameWidth   : 49,
	attackFrameWidth  : 107,
	jumpFrameWidth    : 71,
	runFrameHeight 	  : 66,
	standFrameHeight  : 66,
	attackFrameHeight : 66,
	jumpFrameHeight   : 58,
	roleStatus        : "stand",
	x                 : 100,
	y                 : 100,
	count             : 0,
	direction         : true,
	frameCount        : 0,
}
func init() {
	_ = fenLinWan.InitRoleAnimation("/Users/qushanzu/go/src/game/img/go.png", "/Users/qushanzu/go/src/game/img/stand.png", "/Users/qushanzu/go/src/game/img/attack.png", "/Users/qushanzu/go/src/game/img/jump.png")
} 
// 获得角色
func GetRole(roleName string) *Role {
	if roleName == "fenglinwan" {
		return &fenLinWan
	}
	return &fenLinWan
}