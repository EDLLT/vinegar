// Copyright vinegar-development 2023

package main

import (
	"fmt"
	"os"
)

const (
	PLAYERURL = "https://www.roblox.com/download/client"
	STUDIOURL = "https://www.roblox.com/download/studio"
)

func usage() {
	fmt.Println("usage: vinegar [delete|kill|reset]")
	fmt.Println("       vinegar [player|studio] [args...]")
	if !InFlatpak() {
		fmt.Println("       vinegar [dxvk] install|uninstall")
	}
	os.Exit(1)
}

func main() {
	args := os.Args[1:]
	argsCount := len(args)

	if argsCount < 1 {
		usage()
	}

	CheckDirs(Dirs.Log, Dirs.Pfx)

	switch args[0] {
	case "delete":
		DeleteDirs(Dirs.Data, Dirs.Cache)
	case "dxvk":
		if !InFlatpak() {
			if argsCount < 2 {
				usage()
			}

			switch args[1] {
			case "install":
				DxvkInstall()
			case "uninstall":
				DxvkUninstall()
			}
		} else {
			fmt.Println("DXVK is already installed in the Flatpak and cannot be altered!")
		}
	case "exec":
		Exec("wine", args[1:]...)
	case "kill":
		PfxKill()
	case "player":
		RobloxLaunch("RobloxPlayerLauncher.exe", PLAYERURL, true, args[1:]...)
	case "studio":
		RobloxLaunch("RobloxStudioLauncherBeta.exe", STUDIOURL, false, args[1:]...)
	case "reset":
		DeleteDirs(Dirs.Pfx, Dirs.Log)
		// Automatic creation of the directories after it has been deleted
		CheckDirs(Dirs.Pfx, Dirs.Log)
	default:
		usage()
	}
}