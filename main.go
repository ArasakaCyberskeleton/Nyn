package main

import (
	"crypto/rand"
	"encoding/hex"
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
	time.Sleep(time.Second)                                                      // tuf?
	defer fmt.Printf("(assuming no errors) set flagStr / gbb: %s", flagStr+"\n") //debug
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
		log.Fatalf("Error with strconv.Atoi(blackout) idfk what caused this lowkrey %v ", err)
	}

	if value == 0 {
		return true
	}
	return false
}

func Reprovision4cheap(gbbreset bool) {

	cmd := exec.Command("crossystem", "block_devmode=1")
	cmd1 := exec.Command("vpd", "-i", "RW_VPD", "-s", "block_devmode=1")
	cmd2 := exec.Command("crossystem", "disable_dev_request=1")
	cmd3 := exec.Command("tpm_manager_client", "take_ownership")
	cmd4 := exec.Command("cryptohome", "--action=set_firmware_management_parameters", "--flags=0x01")
	cmd5 := exec.Command("vpd", "-i", "RW_VPD", "-s", "check_enrollment=1")
	gbbnew := exec.Command("futility", "gbb", "--set", "--flash", "--flags=0x0")

	cmd.Run()
	cmd1.Run()
	cmd2.Run()
	cmd3.Run()
	cmd4.Run()
	cmd5.Run()
	if gbbreset {
		gbbnew.Run()
		fmt.Println("GBB SET 0x0 you are now re enrolled.")
	} else {
		fmt.Println("Should be re enrolled!")
	}
}

func main() {
	wp := grabWP()
	Logo()
	choice := options(wp)
	modularity(choice, wp) // holy modular chat is this modular on a scale of whole chromeos in a single file to modular how modular is this is this modular enough?

}

func options(wp bool) int {

	fmt.Println("Hello! welcome to Nyn!\nWhat tool would you like to use!		WP DISABLED?:", wp)
	if wp { //wp == true idk how this even works lowkey
		fmt.Printf(" Defog (1)		\"Re enroll\" (set enrollment) (2)		Vital Product Data General (3)")
	} else {
		fmt.Printf(Red + "Defog (1)" + Normal + "		GrabVPD (2)		Re enroll (set enrollment) (3)")
	}
	fmt.Println("\nInput text! (Single number):")
	var choice int
	_, err := fmt.Scanln(&choice)
	if err != nil {
		log.Fatalf("DUMBASS INVALID INPUT, %v", err)
		return -67 //6-7% sure this value means nothing
	}
	return choice
}

func modularity(choice int, wp bool) { // im going fucking insane oh my god dddddd
	switch choice {
	case 1:
		fmt.Printf("Welcome to defogging! This will not fully work if your WP is enabled. \nThis will set gbb flags and among things.")
		var choice1 int
		fmt.Printf("Choose a GBB\n		1. 0x8031\n		2. 0x80b1\n		3. 0x8091(not suggested)\n		4. Custom GBB\n(0-4): ")
		_, err := fmt.Scanln(&choice1)
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
			fmt.Printf(Red + "ALL INPUT IS UNCHECKED DOUBLE CHECK YOUR ANSWER!\n" + "DO NOT INCLUDE 0x ONLY HEX ie 8031" + Normal + "\nCustom GBB flag: ")
			var choice2 string
			_, err := fmt.Scanln(&choice2)
			if err != nil {
				log.Fatalf("damn idfk%v", err)
			}
			defog(choice2)
		default:
			log.Fatalf("Something went wrong! Or your stupid and entered a invalid input.")
		}

	case 2:
		var input string
		fmt.Printf(Red + "THIS WILL RE-ENROLL YOUR CHROMEBOOK\n DO YOU WANT TO COUNTINE?\n(y-n): " + Normal)
		_, err := fmt.Scanln(&input)
		if err != nil {
			log.Fatalf("something went wrong!%v ", err)
		}
		processedInput := strings.ToLower(strings.TrimSpace(input))
		if processedInput == "y" {
			fmt.Printf(Normal + "Would you like to reset GBB flags to 0x0?\n(y-n): " + Normal)

			var input1 string
			_, err := fmt.Scanln(&input1)

			if err != nil {
				log.Fatalf("something went wrong!%v ", err)
			}

			processedInput1 := strings.ToLower(strings.TrimSpace(input1))

			switch processedInput1 {
			case "y":
				Reprovision4cheap(true)
			case "n":
				Reprovision4cheap(false)
			}

		} else {
			fmt.Printf("Aborting so tuff!")
		}

	case 3:

		if wp { //wp == true idk how this even works lowkey
			fmt.Printf(" Change Secret (1)		Change SN (2)		Vital Product Data General (3)")
		} else {
			fmt.Printf(Red + "Change Secret (1) " + Normal + "		Change SN (2)		Vital Product Data General (3) ")
		}
		fmt.Println("\nInput text! (Single number):")

		var choice2 int
		_, err := fmt.Scanln(&choice2)
		if err != nil {
			log.Fatalf("DUMBASS INVALID INPUT, %v", err)
		}

		switch choice2 {
		case 1:
			fmt.Printf("Are you sure you want to change your secret?\n This is a serious action that can seriously effect your chromebook!\nit is HIGHLEY suggested you run vpd -l and take a picture incase your info gets corrupted.\n(y-n):  ")
			var choice string
			_, err := fmt.Scanln(&choice)
			if err != nil {
				log.Fatalf("DUMBASS INVALID INPUT, %v", err)
			}
			if choice == "y" {
				newsecret := make([]byte, 32) // fatass 256 bit fattty
				_, err := rand.Read(newsecret)
				if err != nil {
					log.Panicf("what the fuck happened like ACTUALLY %v", err)
					log.Fatalf("")
				}
				if err != nil {
					log.Panicf("VPD FAILED LISTING YOUR SECRET:", err)
					log.Fatalf("")
				}
				cmd := exec.Command("vpd", "-g", "stable_device_secret_DO_NOT_SHARE")
				output, _ := cmd.Output()
				oldsecret := strings.TrimSpace(string(output))

				secretHex := hex.EncodeToString(newsecret)

				fmt.Printf("Your current secret pre change is PLEASE SAVE!: %s", oldsecret)
				time.Sleep(2 * time.Second)
				fmt.Printf("Your current secret pre change is PLEASE SAVE!: %s", oldsecret)
				time.Sleep(2 * time.Second)
				fmt.Printf("Your current secret pre change is PLEASE SAVE!: %s", oldsecret)
				time.Sleep(time.Second)
				fmt.Printf("New secret (hex): %s", secretHex)
				writeCmd := exec.Command("sudo", "vpd", "-i", "RO_VPD", "-s", fmt.Sprintf("stable_device_secret_DO_NOT_SHARE=%s", secretHex))
				if err := writeCmd.Run(); err != nil {
					log.Fatalf("Failed writing new secret: %v", err)
				}
			}
		case 2:
			log.Fatalln("not finished :(")
		case 3:
			cmd := exec.Command("vpd", "-l")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			cmd.Run()
			fmt.Println("vpdran")
		}

	default:
		log.Fatalf("Invalid value re run program!") // dumbass entered the wrong input laugh!}

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

// note to self later
