package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/aryuuu/finkita/internal/configs"
	"github.com/aryuuu/finkita/internal/repositories"
	"github.com/aryuuu/finkita/internal/service"
	_ "github.com/lib/pq"
	selenium "sourcegraph.com/sourcegraph/go-selenium"
)

const (
	userIDInputID   = "CorpId"
	passwordInputID = "PassWord"
)

func main() {
	var webDriver selenium.WebDriver
	var err error
	args := []string{
		"--ignore-certificate-errors",
		"--disable-extensions",
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"--headless",
	}
	chromeCaps := map[string]interface{}{
		"excludeSwitches": [1]string{"enable-automation"},
		"args":            args,
	}

	caps := selenium.Capabilities(map[string]interface{}{
		"browserName":   "chrome",
		"chromeOptions": chromeCaps,
	})
	if webDriver, err = selenium.NewRemote(caps, "http://localhost:4444/wd/hub"); err != nil {
		fmt.Printf("Failed to open session: %s\n", err)
		return
	}
	defer webDriver.Quit()

	err = webDriver.Get("https://ibank.bni.co.id/MBAWeb/FMB;jsessionid=00001cFOoVSaowNNAQUec1kioX7:1a1li5jho?page=Thin_SignOnRetRq.xml")
	if err != nil {
		fmt.Printf("Failed to load page: %s\n", err)
		return
	}

	if title, err := webDriver.Title(); err == nil {
		fmt.Printf("Page title: %s\n", title)
	} else {
		fmt.Printf("Failed to get page title: %s", err)
		return
	}

	dbCon := createDBConnection()
	accountRepo := repositories.NewAccountRepo(dbCon)
	accountService := service.NewAccountService(accountRepo)

	// get accounts
	accounts, err := accountService.GetAccountsWithPassword(context.Background())
	if err != nil {
		log.Printf("error getting all accounts: %v", err)
	}

	// log.Printf("accounts: %v", accounts)
	if len(accounts) == 0 {
		return
	}

	// login to each account
	myAccount := accounts[0]
	userIDTextInput, err := webDriver.FindElement(selenium.ById, userIDInputID)
	log.Printf("userIDTextInput: %+v", userIDTextInput)
	if err != nil {
		log.Printf("failed to get user id text input element: %v", err)
		return
	}

	passwordTextInput, err := webDriver.FindElement(selenium.ById, passwordInputID)
	log.Printf("passwordTextInput: %+v", passwordTextInput)
	if err != nil {
		log.Printf("failed to get password text input element: %v", err)
		return
	}

	err = userIDTextInput.SendKeys(myAccount.UserID)
	if err != nil {
		log.Printf("failed to send keys to user id: %v", err)
		return
	}
	err = passwordTextInput.SendKeys(myAccount.Password)
	if err != nil {
		log.Printf("failed to send keys to password: %v", err)
		return
	}

	loginButtonXPath := "//input[@type='submit']"
	loginButton, err := webDriver.FindElement(selenium.ByXPATH, loginButtonXPath)
	if err != nil {
		log.Printf("failed to get login button element: %v", err)
		return
	}

	err = loginButton.Click()
	if err != nil {
		log.Printf("failed to click login button: %v", err)
		return
	}

	time.Sleep(2 * time.Second)

	myPageTitle, err := webDriver.Title()
	if err != nil {
		log.Printf("failed to get my page title")
		return
	}

	log.Printf("myPageTitle: %s", myPageTitle)

	// get info
	// logout

	// webDriver.FindElement(selenium.ByXPATH, "")

	// var elem selenium.WebElement
	// elem, err = webDriver.FindElement(selenium.ByCSSSelector, ".repo .name")
	// if err != nil {
	// 	fmt.Printf("Failed to find element: %s\n", err)
	// 	return
	// }

	// if text, err := elem.Text(); err == nil {
	// 	fmt.Printf("Repository: %s\n", text)
	// } else {
	// 	fmt.Printf("Failed to get text of element: %s\n", err)
	// 	return
	// }

	// output:
	// Page title: go-selenium - Sourcegraph
	// Repository: go-selenium
}

func createDBConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		configs.Postgres.Host, configs.Postgres.Port, configs.Postgres.Username, configs.Postgres.Password, configs.Postgres.Database)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Printf("failed to create connection to postgres db: %v\n", err)
		panic(err)
	}

	return db
}
