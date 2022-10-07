package app

import "gitlab.xxx.com/xxx-xxx/go-kit/logger"
func initLoggerApplicationHook(app *Application) error {
	l, err := logger.New(logger.Options{
		Level:   app.Config.LogLevel,
		Outputs: []string{app.Config.LogFile},
	})

	if err != nil {
		return err
	}

	logger.SetLogger(l)
	app.Logger = l.Sugar()
	return nil
}
