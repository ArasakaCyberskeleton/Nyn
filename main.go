package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

const (
	Normal = "\033[0m"
	Red    = "\033[31m"
)

func grabVPD() {
	cmd := exec.Command("vpd", "-l")
	cmd.Stdout = os.Stdout // forward to terminal
	cmd.Stderr = os.Stderr
	cmd.Run()
	fmt.Println("vpdran")
}

func evilFog() {
	cmd := exec.Command("futility", "gbb", "--set", "--flash", "--flags=0x8031") // can be set to like 180b1 i gotta add custom gbb later
	cmd1 := exec.Command("vpd", "-i", "RW_VPD", "-s", "check_enrollment=0")
	cmd2 := exec.Command("crossystem", "clear_tpm_owner_request=1")
	cmd3 := exec.Command("vpd", "-i", "RW_VPD", "-s", "block_devmode=0")
	cmd.Run()
	cmd1.Run()
	cmd2.Run()
	cmd3.Run()

}

func grabWP() bool {
	cmd := exec.Command("crossystem", "wpsw_cur") // if i had one wish id bring my dog baack :pray:
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("CROSSYSTEM FUCKING EXPLODED %v ", err)
	}
	blackout := strings.TrimSpace(string(out)) // trims whitespace
	value, err := strconv.Atoi(blackout)       //int conversion
	if err != nil {
		log.Printf("Error with strconv.Atoi(blackout)")
	}

	if value == 0 {
		return true
	}
	return false
}

func goodFog() {
	cmd := exec.Command("crossystem", "block_devmode=0")
	cmd1 := exec.Command("vpd", "-i", "RW_VPD", "-s", "block_devmode=0")
	cmd5 := exec.Command("vpd", "-i", "RW_VPD", "-s", "check_enrollment=1")

	cmd.Run()
	cmd1.Run()
	cmd5.Run()
}

func main() {
	wp := grabWP()

	Logo()
	options(wp)

}

func options(wp bool) {
	fmt.Println("Hello! welcome to Nyn!\n What tool would you like to use!\n WP DISABLED?: ", wp)
	if wp { //wp == true idk how this even works lowkey
		fmt.Printf("Defog (1) \n GrabVPD (2) \n \"Re enroll\" (set enrollment) (3)")
	} else {
		fmt.Printf(Red + "Defog (1)" + Normal + " \n GrabVPD (2) \n \"Re enroll\" (set enrollment) (3)")
	}
	fmt.Println("")
}

func Logo() {
	fmt.Println("╔╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╗")
	fmt.Println("╠╬╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╬╣")
	fmt.Println("╠╣         ,--.                      ,--.╠╣")
	fmt.Println("╠╣       ,--.'|                    ,--.'|╠╣")
	fmt.Println("╠╣   ,--,:  : |        ,---,   ,--,:  : |╠╣")
	fmt.Println("╠╣,`--.'`|  ' :       /_ ./|,`--.'`|  ' :╠╣")
	fmt.Println("╠╣|   :  :  | | ,---, |  ' :|   :  :  | |╠╣")
	fmt.Println("╠╣:   |   \\ | :/___/ \\.  : |:   |   \\ | :╠╣")
	fmt.Println("╠╣|   : '  '; | .  \\  \\ ,' '|   : '  '; |╠╣")
	fmt.Println("╠╣'   ' ;.    ;  \\  ;  `  ,''   ' ;.    ;╠╣")
	fmt.Println("╠╣|   | | \\   |   \\  \\    ' |   | | \\   |╠╣")
	fmt.Println("╠╣'   : |  ; .'    '  \\   | '   : |  ; .'╠╣")
	fmt.Println("╠╣|   | '`--'       \\  ;  ; |   | '`--'  ╠╣")
	fmt.Println("╠╣'   : |            :  \\  \\'   : |      ╠╣")
	fmt.Println("╠╣;   |.'             \\  ' ;;   |.'      ╠╣")
	fmt.Println("╠╣'---'                `--` '---'        ╠╣")
	fmt.Println("╠╬╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╦╬╣")
	fmt.Println("╚╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╩╝")
}
