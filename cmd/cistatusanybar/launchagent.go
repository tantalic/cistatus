package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"text/template"

	"github.com/urfave/cli"
)

func installLaunchAgent(conf config) error {
	logger := conf.Logger()
	logger.Println("Installing launch agent")

	t, err := template.New("launchd.tmpl").Parse(launchAgentTemplate)
	if err != nil {
		return cli.NewExitError("error generating launchd agent config", 3)
	}

	exe, err := filepath.Abs(os.Args[0])
	if err != nil {
		return cli.NewExitError("unable to determine path to cistatusanybar", 4)
	}

	path, err := launchAgentPath(conf)
	if err != nil {
		return cli.NewExitError("unable to determine path to save launchagent", 5)
	}

	err = os.MkdirAll(filepath.Dir(path), 0700)
	if err != nil {
		return err
	}

	logPath, err := logPath(conf)
	if err != nil {
		return cli.NewExitError("unable to determine path to save logs", 6)
	}

	err = os.MkdirAll(filepath.Dir(logPath), 0700)
	if err != nil {
		return err
	}

	data := launchdData{
		Config:     conf,
		Executable: exe,
		LogPath:    logPath,
	}

	var b bytes.Buffer
	err = t.Execute(&b, data)
	if err != nil {
		return cli.NewExitError("error generating launchd agent config", 6)
	}
	err = ioutil.WriteFile(path, b.Bytes(), 0644)
	if err != nil {
		return cli.NewExitError("error wrtting launchd agent config", 7)
	}

	loadCmd := exec.Command("launchctl", "load", path)
	return loadCmd.Run()
}

func launchAgentPath(conf config) (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", cli.NewExitError("unable to determine current user", 8)
	}

	return filepath.Join(u.HomeDir, "Library/LaunchAgents", "com.tantalic.cistatusanybar.agent.plist"), nil
}

func logPath(conf config) (string, error) {
	u, err := user.Current()
	if err != nil {
		return "", cli.NewExitError("unable to determine current user", 9)
	}

	filename := "com.tantalic.cistatusanybar.log"

	return filepath.Join(u.HomeDir, "Library/Logs", filename), nil
}

type launchdData struct {
	Executable string
	Config     config
	LogPath    string
}

const launchAgentTemplate = `<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
	<dict>
		<key>Label</key>
		<string>com.tantalic.cistatusanybar</string>
		<key>ProgramArguments</key>
		<array>
			<string>{{ .Executable }}</string>
			<string>{{.Config.StatusHost}}</string>
		</array>
		<key>EnvironmentVariables</key>
		<dict>
			{{ if .Config.StatusPort -}}
			<key>STATUS_PORT</key>
			<string>{{.Config.StatusPort}}</string>
			{{- end }}
			{{- if .Config.AnyBarHost }}
			<key>ANYBAR_HOST</key>
			<string>{{.Config.AnyBarHost}}</string>
			{{- end }}
			{{- if .Config.AnyBarPort }}
			<key>ANYBAR_PORT</key>
			<string>{{.Config.AnyBarPort}}</string>
			{{- end }}
			{{- if .Config.StartAnyBar }}
			<key>START_ANYBAR</key>
			<string>{{.Config.StartAnyBar}}</string>
			{{- end }}
			{{- if .LogPath }}
			<key>VERBOSE</key>
			<string>true</string>
			{{- end }}
		</dict>
		{{ if .LogPath -}}
		<key>StandardOutPath</key>
		<string>{{ .LogPath }}</string>
		{{- end }}
		<key>RunAtLoad</key>
		<true/>
		<key>KeepAlive</key>
		<true/>
		<key>ThrottleInterval</key>
		<integer>30</integer>
	</dict>
</plist>
`
