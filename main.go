package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	Normal = "\033[0m"
	Red    = "\033[31m"
)

func defog(GBBflag string) { //boii get it sh1mmerar reafanerce boiiii deofg :skull:
	flagStr := fmt.Sprintf("0x%s", GBBflag) // ts lowkey took me too long to do
	flashrom := exec.Command("flashrom", "--wp-disable")
	cmd := exec.Command("futility", "gbb", "--set", "--flash", "--flags="+flagStr) // can be set to like 180b1 i gotta add custom gbb later
	cmd1 := exec.Command("vpd", "-i", "RW_VPD", "-s", "check_enrollment=0")
	cmd2 := exec.Command("crossystem", "block_devmode=0")
	cmd3 := exec.Command("vpd", "-i", "RW_VPD", "-s", "block_devmode=0")
	flashrom.Run()
	cmd.Run()
	cmd1.Run()
	cmd2.Run()
	cmd3.Run()
	if err := cmd.Run(); err != nil {
		fmt.Println("futility failed:", err)
	}
	if err := cmd1.Run(); err != nil {
		fmt.Println("vpd check_enrollment failed:", err)
	}
	if err := cmd2.Run(); err != nil {
		fmt.Println("crossystem failed:", err)
	}
	if err := cmd3.Run(); err != nil {
		fmt.Println("vpd block_devmode failed:", err) // what you say about err
	}
	time.Sleep(2 * time.Second)                                             // tuf?
	defer fmt.Printf("(assuming no errors) set flagStr / gbb: %s", flagStr) //debug
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

func evilFog() {
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
		return -67 //6-7% sure this value means nothing
	}
	return choice
}

func modularity(choice int) { // im going fucking insane oh my god dddddd
	switch choice {
	case 1:
		fmt.Printf("Welcome to defogging! This will not fully work if your WP is enabled. This will set gbb flags and other things, Please check source code for more info!")
		var choice1 int
		fmt.Printf("Choose a GBB\n		1. 0x8031\n		2. 0x80b1\n		3. 0x8091(not suggested)\n		4. Custom GBB")
		_, err := fmt.Scanf("%d", &choice1)
		if err != nil {
			log.Fatalf("DUMBASS INVALID INPUT, %v", err)
		}
		switch choice1 {
		case 1:
			defog("8031")
		case 2:
			defog("80b1")
		case 3:
			defog("8091")
		case 4:
			fmt.Printf(Red + "ALL INPUT IS UNCHECKED DOUBLE CHECK YOUR ANSWER!\n" + "DO NOT INCLUDE 0x ONLY HEX ie 8031" + Normal + "Custom GBB flag?")
			var choice2 string
			_, err := fmt.Scanf("%s", &choice2)
			if err != nil {
				log.Fatalf("damn idfk")
			}
			defog(choice2)
		}

	case 2:
		fmt.Printf("5")
	case 3:
		fmt.Printf("5")
	default:
		log.Fatalf("Invalid value re run program!") // dumbass entered the wrong input laugh!
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

func grabVPD() {
	cmd := exec.Command("vpd", "-l")
	cmd.Stdout = os.Stdout // forward to terminal
	cmd.Stderr = os.Stderr
	cmd.Run()
	fmt.Println("vpdran")
}
