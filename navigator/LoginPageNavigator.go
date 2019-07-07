package navigator

import (
	"github.com/sclevine/agouti"
	"log"
)

const (
	USERNAME = ""
	PASSWORD = ""
)

type LoginPageNavigator struct {
	Driver *agouti.WebDriver
	Page   *agouti.Page
}

func NewLoginPageNavigator(wd *agouti.WebDriver) LoginPageNavigator {
	return LoginPageNavigator{wd, nil}
}

func (n *LoginPageNavigator) Navigate() {
	var err error = nil
	n.Page, err = n.Driver.NewPage(agouti.Browser("chrome"))

	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	if err := n.Page.Navigate("https://yoyaku.city.yokohama.lg.jp"); err != nil {
		log.Fatalf("Failed to navigate:%v", err)
	}

	uid := n.Page.FirstByXPath("//*[@id=\"main001\"]/div[1]/div/div/div/div[1]/input")
	ps := n.Page.FirstByXPath("//*[@id=\"main001\"]/div[1]/div/div/div/div[2]/input")

	if err := uid.Fill(USERNAME); err != nil {
		log.Fatalf("Failed to fill uid: %v", err)
	}

	if err := ps.Fill(PASSWORD); err != nil {
		log.Fatalf("Failed to fill ps: %v", err)
	}

	bt := n.Page.FindByID("navi_login_r")
	if err := bt.Click(); err != nil {
		log.Fatalf("Failed to submit: %v", err)
	}

}
