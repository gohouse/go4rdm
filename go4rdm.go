package go4rdm

import (
	"errors"
	"fmt"
	"fyne.io/fyne/dialog"
	"github.com/gen2brain/dlgs"
	"github.com/gohouse/go4rdm/config"
	"github.com/gohouse/go4rdm/data"
	"github.com/gohouse/go4rdm/ui"
	"log"
	"os/exec"
	"runtime"
)

var (
	VERSION int64 = 1
)

//func main() {
//	NewGo4rdm().Run()
//}

type Go4rdm struct {
	conf *config.Config
	rds  *data.Redis
}

func NewGo4rdm() *Go4rdm {
	conf := config.NewConfig()
	return &Go4rdm{conf: conf, rds: data.NewRedis(conf)}
}

func (g *Go4rdm) Run() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("[[panic recover]]: ", err)
		}
	}()
	// 统计一下
	go data.NewApi().Statistic()
	// 检查版本
	if g.checkVersion() {
		return
	}
	u := ui.NewUI(g.conf, g.rds)
	defer func() {
		if err := recover(); err != nil {
			dialog.ShowError(errors.New(fmt.Sprint(err)), u.Window)
		}
	}()
	u.Build()
}

func (g *Go4rdm) checkVersion() bool {
	v := data.NewApi().GetVersion()
	if v.Num > VERSION {
		info, _ := dlgs.Question("update to new version "+v.NumText, v.Notes, false)
		if info {
			err := OpenUrl(v.Url)
			if err != nil {
				err2 := OpenUrl(v.Url)
				if err2 != nil {
					dlgs.Info("error", fmt.Sprintf("auto open error, please visit %s for download the new version", v.Url))
				}
			}
			return true
		}
	}
	return false
}

func OpenUrl(url string) error {
	var cmd string
	var args []string
	switch runtime.GOOS {
	case "darwin":
		cmd = "open"
	case "windows":
		cmd = "cmd"
		args = append(args, `/c`, `start`)
	default:
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
