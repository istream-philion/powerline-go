package main

import (
	"bytes"
	"log"
	//"bufio"
	"os"
	//"github.com/shirou/gopsutil/load"
 "time"
)
import "text/template"


func applyTemplate(p *powerline, templateStr string, buffer *bytes.Buffer) error {
	if templateStr == "DEFAULT" {
		templateStr = "{{venv}} {{user}} {{host}} {{cwd}} {{aws}} {{time}}\n"
	}

	var funcMap = template.FuncMap{
		"aws":       func() string { return os.Getenv("AWS_PROFILE") },
		"cwd":       func() string { return os.Getenv("PWD") },
		"docker":    func() string { return os.Getenv("DOCKER_HOST") },
		//"dotenv":    segmentDotEnv,
		//"exit":      segmentExitCode,
		"fill":      segFiller,
		//"git":       segmentGit,
		//"gitlite":   segmentGitLite,
		//"hg":        segmentHg,
		"host":        func() string { h, _ := os.Hostname(); return h },
		//"jobs":      segmentJobs,
		//"kube":      func() string { segmentKube(p); p. },
		//"load":      load.Avg(),
		//"newline":   segmentNewline,
		//"perms":     segmentPerms,
		//"root":      segmentRoot,
		//"shell-var": segmentShellVar,
		//"ssh":       segmentSsh,
		//"termtitle": segmentTermTitle,
		"time":      func() string { return time.Now().Format("15:04:05") },
		//"node":      segmentNode,
		"user":      func() string { return os.Getenv("USER") },
		"venv":      func() string { return os.Getenv("VIRTUAL_ENV") },
		"vgo":       func() string { return os.Getenv("VIRTUALGO") },
		//"nix-shell": func() string { return os.Getenv("IN_NIX_SHELL") },
	}
	tmpl, err := template.New("prompt").Funcs(funcMap).Parse(templateStr)
  if err != nil {
  	log.Fatalf("parsing: %s", err)
  }

  // Run the template to verify the output.
  //err = tmpl.Execute(os.Stdout, "the go programming language")
  //if err != nil {
  //	log.Fatalf("execution: %s", err)
  //}
	err = tmpl.Execute(os.Stdout, "yoyo")
	//err = tmpl.Execute(bufio.NewWriter(buffer), "yoyo")
	if err != nil {
		log.Fatalf("execution: %s", err)
	}
	log.Printf("buffer: %s", buffer)
	return err
}

func segFiller() string {
	return "                     . . .                                   "
}

//func segmentVirtualGo() string {
//	return os.Getenv("VIRTUALGO")
//}
