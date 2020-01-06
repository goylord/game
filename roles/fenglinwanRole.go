package roles

// Current project path
var prefixImgPath string = "/Users/apple/go/src/game"

var fenLinWan = Role{
	maxRunFrames:      8,
	maxStandFrames:    3,
	maxAttackFrames:   6,
	maxJumpFrames:     10,
	runFrameWidth:     66,
	standFrameWidth:   49,
	attackFrameWidth:  107,
	jumpFrameWidth:    71,
	runFrameHeight:    66,
	standFrameHeight:  66,
	attackFrameHeight: 66,
	jumpFrameHeight:   58,
	roleStatus:        "stand",
	x:                 100,
	y:                 100,
	count:             0,
	direction:         true,
	frameCount:        0,
}

func init() {
	_ = fenLinWan.InitRoleAnimation(prefixImgPath+"/img/go.png", prefixImgPath+"/img/stand.png", prefixImgPath+"/img/attack.png", prefixImgPath+"/img/jump.png")
}

var superMan = Role{
	maxRunFrames:      8,
	maxStandFrames:    3,
	maxAttackFrames:   8,
	maxJumpFrames:     8,
	runFrameWidth:     85,
	standFrameWidth:   72,
	attackFrameWidth:  143,
	jumpFrameWidth:    179,
	runFrameHeight:    66,
	standFrameHeight:  66,
	attackFrameHeight: 113,
	jumpFrameHeight:   113,
	roleStatus:        "stand",
	x:                 200,
	y:                 100,
	count:             0,
	direction:         true,
	frameCount:        0,
}

func init() {
	_ = superMan.InitRoleAnimation(prefixImgPath+"/img/hero.png", prefixImgPath+"/img/heroStand.png", prefixImgPath+"/img/heroAttack.png", prefixImgPath+"/img/heroJump2.png")
}

// 获得角色
func GetRole(roleName string) *Role {
	if roleName == "fenglinwan" {
		return &fenLinWan
	}
	if roleName == "superman" {
		return &superMan
	}
	return &fenLinWan
}
