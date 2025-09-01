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
	// RETURN TRUE FOR WP DISABLED(0) FALSE FOR ANY OTHER VALUE IT RETURNS FALSE
	cmd := exec.Command("crossystem", "wpsw_cur") // if i had one wish id bring my dog baack :pray:
	out, err := cmd.Output()
	if err != nil {
		log.Fatalf("CROSSYSTEM FUCKING EXPLODED %v ", err)
	}
	blackout := strings.TrimSpace(string(out)) // trims whitespace
	value, err := strconv.Atoi(blackout)       //int conversion
	if err != nil {
		log.Printf("Error with strconv.Atoi(blackout) idfk what caused this lowkrey")
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
	wp := true //grabWP()
	Logo()
	choice := options(wp)
	modularity(choice)

}

func options(wp bool) int {

	fmt.Println("Hello! welcome to Nyn!\nWhat tool would you like to use!		WP DISABLED?:", wp)
	if wp { //wp == true idk how this even works lowkey
		fmt.Printf(" Defog (1)		GrabVPD (2)		\"Re enroll\" (set enrollment) (3)")
	} else {
		fmt.Printf(Red + "Defog (1)" + Normal + "		GrabVPD (2)		Re enroll (set enrollment) (3)")
	}
	fmt.Println("\nInput text! (Single number):")
	var choice int
	_, err := fmt.Scanf("%d", &choice)
	if err != nil {
		log.Fatalf("DUMBASS INVALID INPUT, %v", err)
		return -67
	}
	return choice
}

func modularity(choice int) { // im going fucking insane oh my god dddddd
	switch choice {
	case 1:
		// i sleep but bascially make this go down a for or switch that will basciaslly record the GBB you would like to use and then blow your shit out
	case 2:
		fmt.Printf("5")
	case 3:
		fmt.Printf("5")
	default:
		log.Fatal("you did something bad or i did something REALLY bad (report if you entered valid options without space or newline)")
	}
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
