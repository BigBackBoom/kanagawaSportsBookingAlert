package crawler

import (
	"AutoReservationSys/navigator"
	"github.com/sclevine/agouti"
	"golang.org/x/xerrors"
)

type ScheduleCrawler struct {
	driver *agouti.WebDriver
}

func (s *ScheduleCrawler) StartCrawling() error {

	err := s.initDriver()

	if err != nil {
		return xerrors.Errorf("Failed to StartCrawling: %w", err)
	}

	err = s.navigateToSchedulePage()

	if err != nil {
		return xerrors.Errorf("Failed to StartCrawling: %w", err)
	}

	err = s.close()

	if err != nil {
		return xerrors.Errorf("Failed to StartCrawling: %w", err)
	}

	return nil
}

func (s *ScheduleCrawler) initDriver() error {

	driver := agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless", // headlessモードの指定
			"--no-sandbox",
		}),
		agouti.Debug,
	)

	//driver := agouti.ChromeDriver()
	if err := driver.Start(); err != nil {
		return xerrors.Errorf("Failed to start driver: %w", err)
	}

	s.driver = driver
	return nil
}

func (s *ScheduleCrawler) navigateToSchedulePage() error {

	// open login page and login
	ln := navigator.NewLoginPageNavigator(s.driver)
	ln.Navigate()

	// start booking page navigation
	bn := navigator.NewSchedulePageNavigator(ln.Driver, ln.Page)
	err := bn.Navigate()

	if err != nil {
		return xerrors.Errorf("failed to navigateToSchedulePage: %w", err)
	}

	return nil
}

func (s *ScheduleCrawler) close() error {
	if err := s.driver.Stop(); err != nil {
		return xerrors.Errorf("Failed to stop web driver: %w", err)
	}
	return nil
}
