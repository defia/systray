package main

import (
	"fmt"
	"log"
	"os"

	"github.com/defia/systray"
	"github.com/defia/systray/example/icon"
)

func main() {
	// Should be called at the very beginning of main().
	systray.Run(onReady)
}

func onReady() {
	systray.HideConsole()
	systray.SetIcon(icon.Data)
	systray.SetTitle("Awesome App")
	systray.SetTooltip("Lantern")
	mQuit := systray.AddMenuItem("Quit", "Quit the whole app")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
		os.Exit(0)
		fmt.Println("Quit now...")

	}()
	log.SetFlags(log.Lshortfile)
	// We can manipulate the systray in other goroutines
	go func() {

		systray.SetIcon(icon.Data)
		systray.SetTitle("defia ftw!")
		systray.SetTooltip("defia yeah!")
		//		mChange := systray.AddMenuItem("Change Me", "Change Me")
		mChecked := systray.AddMenuItem("Show", "Check Me")
		//		mEnabled := systray.AddMenuItem("Enabled", "Enabled")
		//		systray.AddMenuItem("Ignored", "Ignored")
		//		mUrl := systray.AddMenuItem("Open Lantern.org", "my home")
		mQuit := systray.AddMenuItem("退出", "Quit the whole app")
		for {
			select {
			//			case <-mChange.ClickedCh:
			//				mChange.SetTitle("I've Changed")

			case <-mChecked.ClickedCh:
				if mChecked.Checked() {
					systray.HideConsole()
					mChecked.Uncheck()
					//mChecked.SetTitle("show")
				} else {
					systray.ShowConsole()
					mChecked.Check()
					//mChecked.SetTitle("hide")
				}
				//			case <-mEnabled.ClickedCh:
				//				mEnabled.SetTitle("Disabled")
				//				mEnabled.Disable()
				//			case <-mUrl.ClickedCh:
				//				open.Run("https://www.getlantern.org")
			case <-mQuit.ClickedCh:
				systray.Quit()
				fmt.Println("Quit2 now...")
				return
			}
		}
	}()
}
