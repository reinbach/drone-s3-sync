package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
)

func createCredentials() error {
	f, err := os.Create(path.Join(os.Getenv("HOME"), ".aws/credentials"))
	if err != nil {
		return err
	}
	defer f.Close()
	buf := bufio.NewWriter(f)
	err = writeCredentials(buf)
	if err != nil {
		return err
	}
	buf.Flush()
	return nil
}

func writeCredentials(w io.Writer) error {
	accessKey := os.Getenv("AWS_ACCESS_KEY_ID")
	secretKey := os.Getenv("AWS_SECRET_ACCESS_KEY")
	_, err := io.WriteString(
		w,
		fmt.Sprintf(
			`[default]
aws_access_key_id=%s
aws_secret_access_key=%s
`,
			accessKey,
			secretKey,
		),
	)

	return err
}

func sync() error {
	args := []string{"s3", "sync"}
	args = append(args, os.Getenv("PLUGIN_SOURCE"))

	target := os.Getenv("PLUGIN_TARGET")
	if !strings.HasPrefix(target, "/") {
		target = fmt.Sprintf("/%s", target)
	}
	dest := fmt.Sprintf("s3://%s%s", os.Getenv("PLUGIN_BUCKET"), target)
	args = append(args, dest)

	args = append(args, "--region")
	args = append(args, os.Getenv("PLUGIN_REGION"))

	acl := os.Getenv("PLUGIN_ACL")
	if acl != "" {
		args = append(args, "--acl")
		args = append(args, acl)
	}

	delete := os.Getenv("PLUGIN_DELETE")
	if delete == "true" {
		args = append(args, "--delete")
	}

	include := os.Getenv("PLUGIN_INCLUDE")
	if include != "" {
		args = append(args, "--include")
		args = append(args, include)
	}

	exclude := os.Getenv("PLUGIN_EXCLUDE")
	if exclude != "" {
		args = append(args, "--exclude")
		args = append(args, exclude)
	}

	cmd := exec.Command("aws", args...)
	cmd.Dir = os.Getenv("DRONE_WORKSPACE")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	fmt.Println("$", strings.Join(cmd.Args, " "))
	return cmd.Run()
}

func main() {
	err := createCredentials()
	if err != nil {
		log.Fatal(err)
	}
	err = sync()
	if err != nil {
		log.Fatal(err)
	}
}
