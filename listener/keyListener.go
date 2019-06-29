package listener

import (
	"github.com/hajimehoshi/ebiten/inpututil"
	"github.com/hajimehoshi/ebiten"
	"game/roles"
)

func JudgeKeyPress(role *roles.Role) {
	d := inpututil.KeyPressDuration(ebiten.KeyD)
	a := inpututil.KeyPressDuration(ebiten.KeyA)
	j := inpututil.KeyPressDuration(ebiten.KeyJ)
	k := inpututil.KeyPressDuration(ebiten.KeyK)
	if d >= 1 { // 按了D
		role.Run()
	} else if (a >= 1) { // 按了A
		role.RunBack()
	} else {
		role.Stand()
	}
	if j >= 1 {
		role.Attack()
	} else {
		role.StopAttack()
	}
	if k >= 1 {
		role.Jump()
	} else {
		role.StopJump()
	}

}