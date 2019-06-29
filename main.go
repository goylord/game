package main
import (
	"game/roles"
	"game/config"
	"game/listener"
	"log"
	"fmt"
	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)


var (
	fenlinwan *roles.Role
	drawImage *ebiten.Image
	roleMainOP  *ebiten.DrawImageOptions
)

func update(screen *ebiten.Image) error {
	if (ebiten.IsDrawingSkipped()) {
		return nil
	}
	fenlinwan.FrameCountPlus()
	listener.JudgeKeyPress(fenlinwan)
	i := (fenlinwan.GetFramesCount() / config.WindowFrames) % fenlinwan.GetAnimationFrames()
	roleMainOP.GeoM.Reset()
	if !fenlinwan.GetDirection() {
		roleMainOP.GeoM.Scale(-1, 1) //反向
	}
	roleMainOP.GeoM.Translate(fenlinwan.GetPosition())
	screen.DrawImage(fenlinwan.RoleAnimationCollection[fenlinwan.GetStatus()][i], roleMainOP)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.CurrentTPS()))
	return nil
}
func main() {
	roleMainOP = &ebiten.DrawImageOptions{}
	fenlinwan = roles.GetRole("superman")
	fenlinwan.SetPosition(100, config.ScreentHeight - 120)
	if err := ebiten.Run(update, config.ScreentWidth, config.ScreentHeight, 1, "枫林晚"); err != nil {
		log.Fatal("报错", err)
	}
}