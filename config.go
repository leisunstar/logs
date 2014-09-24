package logs

type config struct {
	// level
	Level       string `conf:"level"`
	DatetimeFmt string `conf:"datetime_fmt"`
	// debug
	DebugEnable bool   `conf:"debug_enable"`
	DebugType   string `conf:"debug_type"`
	DebugOut    string `conf:"debug_out"`
	// info
	InfoEnable bool   `conf:"info_enable"`
	InfoType   string `conf:"info_type"`
	InfoOut    string `conf:"info_out"`
	// warning
	WarningEnable bool   `conf:"warning_enable"`
	WarningType   string `conf:"warning_type"`
	WarningOut    string `conf:"warning_out"`
	// error
	ErrorEnable bool   `conf:"error_enable"`
	ErrorType   string `conf:"error_type"`
	ErrorOut    string `conf:"error_out"`
}
