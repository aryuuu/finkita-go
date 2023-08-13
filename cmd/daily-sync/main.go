package main

import (
	// "context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	// "github.com/aryuuu/finkita/domain"
	"github.com/aryuuu/finkita/internal/configs"

	// "github.com/aryuuu/finkita/internal/repositories"
	// "github.com/aryuuu/finkita/internal/service"
	_ "github.com/lib/pq"
	"github.com/tebeka/selenium"
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
	if webDriver, err = selenium.NewRemote(caps, "http://localhost:4444"); err != nil {
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

	// dbCon := createDBConnection()
	// accountRepo := repositories.NewAccountRepo(dbCon)
	// accountService := service.NewAccountService(accountRepo)

	// // get accounts
	// accounts, err := accountService.GetAccountsWithPassword(context.Background())
	// if err != nil {
	// 	log.Printf("error getting all accounts: %v", err)
	// }

	// // log.Printf("accounts: %v", accounts)
	// if len(accounts) == 0 {
	// 	return
	// }

	// login to each account
	myUsername := configs.Scraper.MyUsername
	myPassword := configs.Scraper.MyPassword
	err = login(webDriver, myUsername, myPassword)
	if err != nil {
		return
	}

	defer logOut(webDriver)

	err = openAccountsPage(webDriver)
	if err != nil {
		return
	}

	err = openMutationsPage(webDriver)
	if err != nil {
		return
	}

	err = openOPRPage(webDriver)
	if err != nil {
		return
	}

	// // TODO: handle this
	// // select account
	// selectRekeningButtonXPath := "//a[@id='AccountMenuList']"
	// selectRekeningButton, err := webDriver.FindElement(selenium.ByXPATH, selectRekeningButtonXPath)
	// if err != nil {
	// 	log.Printf("failed to find mutasi button: %v", err)
	// 	return
	// }
	// err = selectRekeningButton.Click()
	// if err != nil {
	// 	log.Printf("failed to click mutasi button: %v", err)
	// 	return
	// }

	err = openCurrentMonthStatementPage(webDriver)
	if err != nil {
		return
	}

	// mutations := []domain.Mutation{}
	// for {
	// }

}

func login(webDriver selenium.WebDriver, username string, password string) error {
	userIDTextInput, err := webDriver.FindElement(selenium.ByID, userIDInputID)
	if err != nil {
		log.Printf("failed to get user id text input element: %v", err)
		return err
	}

	passwordTextInput, err := webDriver.FindElement(selenium.ByID, passwordInputID)
	if err != nil {
		log.Printf("failed to get password text input element: %v", err)
		return err
	}

	err = userIDTextInput.SendKeys(username)
	if err != nil {
		log.Printf("failed to send keys to user id: %v", err)
		return err
	}
	err = passwordTextInput.SendKeys(password)
	if err != nil {
		log.Printf("failed to send keys to password: %v", err)
		return err
	}

	loginButtonXPath := "//input[@type='submit']"
	loginButton, err := webDriver.FindElement(selenium.ByXPATH, loginButtonXPath)
	if err != nil {
		log.Printf("failed to get login button element: %v", err)
		return err
	}

	err = loginButton.Click()
	if err != nil {
		log.Printf("failed to click login button: %v", err)
		return err
	}

	time.Sleep(2 * time.Second)

	myPageTitle, err := webDriver.Title()
	if err != nil {
		log.Printf("failed to get my page title")
		return err
	}
	log.Printf("myPageTitle: %s", myPageTitle)

	return nil
}

func openAccountsPage(webDriver selenium.WebDriver) error {
	// open rekening page
	rekeningButtonXPath := "//a[@id='MBMenuList']"
	rekeningButton, err := webDriver.FindElements(selenium.ByXPATH, rekeningButtonXPath)
	if err != nil {
		log.Printf("failed to find rekening button: %v", err)
		return err
	}

	if len(rekeningButton) == 0 {
		log.Printf("no rekening button found")
		return errors.New("no rekening button found")
	}

	err = rekeningButton[0].Click()
	if err != nil {
		log.Printf("failed to click rekening button: %v", err)
		return err
	}

	return nil
}

func openMutationsPage(webDriver selenium.WebDriver) error {
	// open mutasi page
	mutasiButtonXPath := "//a[@id='AccountMenuList']"
	mutasiButton, err := webDriver.FindElements(selenium.ByXPATH, mutasiButtonXPath)
	if err != nil {
		log.Printf("failed to find mutasi button: %v", err)
		return err
	}

	if len(mutasiButton) < 3 {
		log.Printf("no mutasi button found")
		return errors.New("no mutasi button found")
	}

	err = mutasiButton[2].Click()
	if err != nil {
		log.Printf("failed to click mutasi button: %v", err)
		return err
	}

	return nil
}

func openOPRPage(webDriver selenium.WebDriver) error {
	// open select account dropdown
	// MAIN_ACCOUNT_TYPE both id and name, type input
	selectRekeningXPath := "//select[@id='MAIN_ACCOUNT_TYPE']"
	selectRekening, err := webDriver.FindElement(selenium.ByXPATH, selectRekeningXPath)
	if err != nil {
		log.Printf("failed to find rekening selection: %v", err)
		return err
	}
	err = selectRekening.Click()
	if err != nil {
		log.Printf("failed to click rekening selection: %v", err)
		return err
	}

	oprOptionXPath := "//option[@value='OPR']"
	oprOption, err := webDriver.FindElement(selenium.ByXPATH, oprOptionXPath)
	if err != nil {
		log.Printf("failed to find opr option: %v", err)
		return err
	}

	err = oprOption.Click()
	if err != nil {
		log.Printf("failed to click opr option: %v", err)
		return err
	}

	log.Println("OPR selection success")

	// lanjut button
	lanjutButtonID := "AccountIDSelectRq"
	lanjutButton, err := webDriver.FindElement(selenium.ByID, lanjutButtonID)
	if err != nil {
		log.Printf("failed to find lanjut button: %v", err)
		return err
	}

	err = lanjutButton.Click()
	if err != nil {
		log.Printf("failed to click lanjut button: %v", err)
		return err
	}

	return nil
}

func openCurrentMonthStatementPage(webDriver selenium.WebDriver) error {
	// click open current month statement page
	txnDateRadioButtonXPath := "//input[@id='Search_Option_6']"
	txnDataRadioButton, err := webDriver.FindElement(selenium.ByXPATH, txnDateRadioButtonXPath)
	if err != nil {
		log.Printf("failed to find txn radio button: %v", err)
		return err
	}
	err = txnDataRadioButton.Click()
	if err != nil {
		log.Printf("failed to click txn date radio button: %v", err)
		return err
	}

	// input date range
	startYearAndMonth := time.Now().Format("2006-Jan")
	startDate := fmt.Sprintf("%s-01", startYearAndMonth)
	startDateInputTextXPath := "//input[@id='txnSrcFromDate']"
	endDateInputTextXPath := "//input[@id='txnSrcToDate']"
	startDateInputText, err := webDriver.FindElement(selenium.ByXPATH, startDateInputTextXPath)
	if err != nil {
		log.Printf("failed to find start date input text: %v", err)
		return err
	}

	err = startDateInputText.SendKeys(startDate)
	if err != nil {
		log.Printf("failed to send keys to start date input text: %v", err)
		return err
	}

	endDateInputText, err := webDriver.FindElement(selenium.ByXPATH, endDateInputTextXPath)
	if err != nil {
		log.Printf("failed to find end date input text: %v", err)
		return err
	}

	currentDate := time.Now().Format("2006-Jan-02")
	err = endDateInputText.SendKeys(currentDate)
	if err != nil {
		log.Printf("failed to send keys to end date input text: %v", err)
		return err
	}

	// next button 
	nextButtonId  := "FullStmtInqRq"
	nextButton, err := webDriver.FindElement(selenium.ByID, nextButtonId)
	if err != nil {
		log.Printf("failed to find txn radio button: %v", err)
		return err
	}

	err = nextButton.Click()
	if err != nil {
		log.Printf("failed to find txn radio button: %v", err)
		return err
	}

	log.Println("we are in baby")
	return nil
}

func logOut(webDriver selenium.WebDriver) {
	// logout
	logoutButtonXPath := "//input[@name='LogOut']"
	logoutButton, err := webDriver.FindElement(selenium.ByXPATH, logoutButtonXPath)
	if err != nil {
		log.Printf("failed to find logout button: %v", err)
		return
	}

	err = logoutButton.Click()
	if err != nil {
		log.Printf("failed to click logout button: %v", err)
		return
	}

	// logout
	logoutConfirmationButtonXPath := "//input[@name='__LOGOUT__']"
	logoutConfirmationButton, err := webDriver.FindElement(selenium.ByXPATH, logoutConfirmationButtonXPath)
	if err != nil {
		log.Printf("failed to find logout confirmation button: %v", err)
		return
	}

	err = logoutConfirmationButton.Click()
	if err != nil {
		log.Printf("failed to click logout confirmation button: %v", err)
		return
	}

	log.Print("logged out successfully")
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
