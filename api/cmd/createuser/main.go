package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/joshuakinkade/web-journal/auth"
	"github.com/joshuakinkade/web-journal/db"
	"github.com/joshuakinkade/web-journal/dtos"
	"github.com/joshuakinkade/web-journal/users"
	"github.com/joshuakinkade/web-journal/validation"
	"golang.org/x/term"
)

const (
	colorReset  = "\033[0m"
	colorBold   = "\033[1m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorCyan   = "\033[36m"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	envPath, err := findEnv()
	if err != nil {
		fatal("Could not find .env file in this directory or any parent directory.")
	}
	godotenv.Load(envPath)

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		fatal("DB_PORT is not set or not a valid number.")
	}

	err = db.Init(db.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     port,
		DBName:   os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		fatal(fmt.Sprintf("Could not connect to database: %v", err))
	}

	fmt.Println()
	fmt.Printf("%s%s Create User %s\n", colorBold, colorCyan, colorReset)
	fmt.Println(strings.Repeat("─", 40))
	fmt.Println()

	// Email
	email := prompt(reader, "Email")
	if email == "" {
		fatal("Email is required.")
	}

	// Display name
	displayName := prompt(reader, "Display Name")
	if displayName == "" {
		fatal("Display name is required.")
	}

	// Password
	password := promptPassword("Password")
	if password == "" {
		fatal("Password is required.")
	}

	confirm := promptPassword("Confirm Password")
	if password != confirm {
		fatal("Passwords do not match.")
	}

	// Admin
	isAdmin := promptBool(reader, "Grant admin role?")

	// Summary
	fmt.Println()
	fmt.Println(strings.Repeat("─", 40))
	fmt.Printf("%s%sSummary%s\n", colorBold, colorYellow, colorReset)
	fmt.Printf("  %-16s %s\n", "Email:", email)
	fmt.Printf("  %-16s %s\n", "Display Name:", displayName)
	fmt.Printf("  %-16s %s\n", "Admin:", boolLabel(isAdmin))
	fmt.Println(strings.Repeat("─", 40))
	fmt.Println()

	if !promptBool(reader, "Create this user?") {
		fmt.Printf("%sAborted.%s\n\n", colorYellow, colorReset)
		os.Exit(0)
	}

	svc := users.NewService(auth.Manager{}, users.NewRepository())
	input := dtos.NewUserInput{
		Email:           email,
		DisplayName:     displayName,
		Password:        password,
		ConfirmPassword: confirm,
	}
	user, err := svc.CreateUser(input, isAdmin)
	if err != nil {
		var ve validation.Error
		if errors.As(err, &ve) {
			fmt.Printf("\n%s%sValidation Errors:%s\n", colorBold, colorRed, colorReset)
			for field, msg := range ve.Errors {
				fmt.Printf("  %s%-16s%s %s\n", colorYellow, field+":", colorReset, msg)
			}
			fmt.Println()
			os.Exit(1)
		}
		fatal(fmt.Sprintf("Failed to create user: %v", err))
	}

	fmt.Println()
	fmt.Printf("%s%s User created successfully! %s\n", colorBold, colorGreen, colorReset)
	fmt.Println(strings.Repeat("─", 40))
	fmt.Printf("  %-16s %s\n", "ID:", user.ID())
	fmt.Printf("  %-16s %s\n", "Email:", user.Email())
	fmt.Printf("  %-16s %s\n", "Display Name:", user.DisplayName())
	fmt.Printf("  %-16s %s\n", "Slug:", user.Slug())
	fmt.Printf("  %-16s %s\n", "Admin:", boolLabel(isAdmin))
	fmt.Println(strings.Repeat("─", 40))
	fmt.Println()
}

// prompt prints a labeled prompt and reads a line of input.
func prompt(r *bufio.Reader, label string) string {
	fmt.Printf("%s%s:%s ", colorBold, label, colorReset)
	input, _ := r.ReadString('\n')
	return strings.TrimSpace(input)
}

// promptPassword reads a password without echoing it to the terminal.
func promptPassword(label string) string {
	fmt.Printf("%s%s:%s ", colorBold, label, colorReset)
	bytes, err := term.ReadPassword(int(syscall.Stdin))
	fmt.Println()
	if err != nil {
		fatal(fmt.Sprintf("Failed to read password: %v", err))
	}
	return strings.TrimSpace(string(bytes))
}

// promptBool asks a yes/no question and returns the result.
func promptBool(r *bufio.Reader, label string) bool {
	for {
		fmt.Printf("%s%s%s [y/n]: ", colorBold, label, colorReset)
		input, _ := r.ReadString('\n')
		switch strings.ToLower(strings.TrimSpace(input)) {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Printf("%sPlease enter y or n.%s\n", colorYellow, colorReset)
		}
	}
}

// fatal prints an error message and exits.
func fatal(msg string) {
	fmt.Printf("\n%s%sError:%s %s\n\n", colorBold, colorRed, colorReset, msg)
	os.Exit(1)
}

// findEnv walks up from the working directory until it finds a .env file.
func findEnv() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	for {
		candidate := filepath.Join(dir, ".env")
		if _, err := os.Stat(candidate); err == nil {
			return candidate, nil
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			// reached filesystem root
			return "", fmt.Errorf(".env not found")
		}
		dir = parent
	}
}

func boolLabel(v bool) string {
	if v {
		return fmt.Sprintf("%sYes%s", colorGreen, colorReset)
	}
	return fmt.Sprintf("%sNo%s", colorRed, colorReset)
}
