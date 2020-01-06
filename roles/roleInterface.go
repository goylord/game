package roles

import (
	"game/config"
	"image"
	_ "image/png"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Role struct {
	maxRunFrames                  int // 奔跑动画帧数
	maxStandFrames                int // 站立动画帧数
	maxAttackFrames               int // 攻击动画帧数
	maxJumpFrames                 int
	runFrameWidth                 int
	standFrameWidth               int
	attackFrameWidth              int
	jumpFrameWidth                int
	runFrameHeight                int
	standFrameHeight              int
	attackFrameHeight             int
	jumpFrameHeight               int
	roleRunAnimationCollection    map[int]*ebiten.Image // 奔跑动画集合
	roleStandAnimationCollection  map[int]*ebiten.Image // 站立动画集合
	roleAttackAnimationCollection map[int]*ebiten.Image // 攻击动画集合
	roleJumpAnimationCollection   map[int]*ebiten.Image // 攻击动画集合
	RoleAnimationCollection       map[string]map[int]*ebiten.Image
	roleStatus                    string // 角色状态 stand, attack, running, jumpping
	jumpDirection                 string
	x                             float64
	y                             float64
	count                         int
	direction                     bool
	frameCount                    int
}
type RoleInterface interface {
	InitRoleAnimation(string, string, string, string) error // 初始化角色动画信息，传入三个图片地址
	SetPosition(float64, float64)
	Run()
	RunBack()
	Stand()
	Jump()
	Attack()
	StopAttack()
	Count()
	ResetCount()
	GetStatus() string
	GetDirection() bool
	GetPosition() (float64, float64)
	GetAnimationFrames() int
	FrameCountPlus()
	GetFramesCount() int
	ResetCountFrames()
}

func (role *Role) InitRoleAnimation(runImageFile, standImageFile, attackImageFile string, jumpImageFile string) (err error) {
	runnerImage, _, err := ebitenutil.NewImageFromFile(runImageFile, ebiten.FilterDefault)
	standImage, _, err := ebitenutil.NewImageFromFile(standImageFile, ebiten.FilterDefault)
	attackImage, _, err := ebitenutil.NewImageFromFile(attackImageFile, ebiten.FilterDefault)
	jumpImage, _, err := ebitenutil.NewImageFromFile(jumpImageFile, ebiten.FilterDefault)

	if err != nil {
		return
	}
	// 初始化动画集合
	sy := 0
	role.roleRunAnimationCollection = make(map[int]*ebiten.Image)
	role.roleStandAnimationCollection = make(map[int]*ebiten.Image)
	role.roleAttackAnimationCollection = make(map[int]*ebiten.Image)
	role.roleJumpAnimationCollection = make(map[int]*ebiten.Image)
	role.RoleAnimationCollection = make(map[string]map[int]*ebiten.Image)
	for i := 0; i < role.maxRunFrames; i++ {
		role.roleRunAnimationCollection[i] = runnerImage.SubImage(image.Rect(i*role.runFrameWidth, sy, (i+1)*role.runFrameWidth, sy+role.runFrameHeight)).(*ebiten.Image)
	}
	for i := 0; i < role.maxRunFrames; i++ {
		role.roleStandAnimationCollection[i] = standImage.SubImage(image.Rect(i*role.standFrameWidth, sy, (i+1)*role.standFrameWidth, sy+role.standFrameHeight)).(*ebiten.Image)
	}
	for i := 0; i < role.maxRunFrames; i++ {
		role.roleAttackAnimationCollection[i] = attackImage.SubImage(image.Rect(i*role.attackFrameWidth, sy, (i+1)*role.attackFrameWidth, sy+role.attackFrameHeight)).(*ebiten.Image)
	}
	for i := 0; i < role.maxJumpFrames; i++ {
		role.roleJumpAnimationCollection[i] = jumpImage.SubImage(image.Rect(i*role.jumpFrameWidth, sy, (i+1)*role.jumpFrameWidth, sy+role.jumpFrameHeight)).(*ebiten.Image)
	}
	role.RoleAnimationCollection["running"] = role.roleRunAnimationCollection
	role.RoleAnimationCollection["stand"] = role.roleStandAnimationCollection
	role.RoleAnimationCollection["attack"] = role.roleAttackAnimationCollection
	role.RoleAnimationCollection["jumpping"] = role.roleJumpAnimationCollection
	return
}
func (role *Role) SetPosition(x, y float64) {
	role.x = x
	role.y = y
}
func (role *Role) Count() {
	role.count++
}
func (role *Role) ResetCount() {
	role.count = 0
}
func (role *Role) GetStatus() string {
	return role.roleStatus
}
func (role *Role) GetDirection() bool {
	return role.direction
}
func (role *Role) GetPosition() (float64, float64) {
	return role.x, role.y
}

// 向前
func (role *Role) Run() {
	if role.count != 0 { // 判断是否正在加锁
		if role.roleStatus == "jumpping" {
			role.jumpDirection = "right"
		}
		return
	}
	if !role.direction {
		role.x = role.x - float64(role.runFrameWidth)
	}
	role.roleStatus = "running"
	role.direction = true
	role.x += 3
}

// 向后
func (role *Role) RunBack() {
	if role.count != 0 { // 判断是否正在加锁
		if role.roleStatus == "jumpping" {
			role.jumpDirection = "left"
		}
		return
	}
	if role.direction {
		role.x = role.x + float64(role.runFrameWidth)
	}
	role.roleStatus = "running"
	role.direction = false
	role.x -= 3
}

// 站立
func (role *Role) Stand() {
	if role.count != 0 { // 判断是否正在加锁
		return
	}
	role.roleStatus = "stand"
}
func (role *Role) Jump() {
	if role.count != 0 { // 判断是否正在加锁
		if role.roleStatus == "jumpping" {
			role.StopJump()
		}
		return
	}
	role.roleStatus = "jumpping"
	role.Count()
	role.ResetCountFrames()
	role.y -= float64(role.jumpFrameHeight - role.runFrameHeight)
}
func (role *Role) StopJump() {
	if role.roleStatus == "jumpping" {
		role.Count()
		if role.count < (config.WindowFrames * (role.maxJumpFrames - 1) / 2) {
			role.y -= 3
		}
		if role.count > (config.WindowFrames * (role.maxJumpFrames) / 2) {
			role.y += 3
		}
		if role.jumpDirection == "left" {
			role.x -= 3
		} else if role.jumpDirection == "right" {
			role.x += 3
		}
		if role.count >= config.WindowFrames*(role.maxJumpFrames-1) {
			role.jumpDirection = ""
			role.ResetCount()
			role.Stand()
			role.y += float64(role.jumpFrameHeight - role.runFrameHeight)
		}
	}
}

// 攻击
func (role *Role) Attack() {
	if role.count != 0 {
		return
	}
	role.roleStatus = "attack"
	role.Count()
	role.ResetCountFrames()
	// 修复攻击图片过高时产生误差
	role.y -= float64(role.attackFrameHeight - role.runFrameHeight)
}
func (role *Role) StopAttack() {
	if role.roleStatus == "attack" {
		role.Count()
		if role.count >= config.WindowFrames*(role.maxAttackFrames-1) {
			role.ResetCount()
			role.y += float64(role.attackFrameHeight - role.runFrameHeight)
			role.Stand()
		}
	}
}
func (role *Role) GetAnimationFrames() int {
	if role.roleStatus == "stand" {
		return role.maxStandFrames
	} else if role.roleStatus == "running" {
		return role.maxRunFrames
	} else if role.roleStatus == "attack" {
		return role.maxAttackFrames
	} else if role.roleStatus == "jumpping" {
		return role.maxJumpFrames
	}
	return role.maxRunFrames
}
func (role *Role) FrameCountPlus() {
	role.frameCount++
	if role.frameCount == config.WindowFrames*role.GetAnimationFrames() {
		role.ResetCountFrames()
	}
}
func (role *Role) GetFramesCount() int {
	return role.frameCount
}
func (role *Role) ResetCountFrames() {
	role.frameCount = 0
}
