package docgen

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	goCdLogger "github.com/nikhilsbhat/gocd-sdk-go/pkg/logger"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

const (
	docGenName          = "UrFaveCliDocGen"
	docFolderPermission = 0o755
)

// shamelessly copied code from https://github.com/urfave/cli/issues/340#issuecomment-334389849 with few modification.
// thankful to https://github.com/Southclaws for this.
func getDocs(app *cli.App, logger *logrus.Logger) string {
	buffer := bytes.Buffer{}

	buffer.WriteString(fmt.Sprintf("# `%s`\n\n", app.Name))
	buffer.WriteString(fmt.Sprintf("%s\n%s\n", app.Usage, app.Version))

	if app.Description != "" {
		buffer.WriteString(app.Description)
		buffer.WriteString("\n\n")
	}

	logger.Info("generating documents for subcommands")
	buffer.WriteString("## Subcommands\n\n")

	for _, command := range app.Commands {
		logger.Infof("generating documents for subcommand '%s'", command.Name)

		buffer.WriteString(fmt.Sprintf("### `%s`\n\n", command.Name))
		if command.Usage != "" {
			logger.Infof("generating documents on usage subcommand '%s'", command.Name)
			buffer.WriteString(command.Usage)
			buffer.WriteString("\n\n")
		}
		if command.Description != "" {
			logger.Infof("generating documents on description of subcommand '%s'", command.Name)
			buffer.WriteString(command.Description)
			buffer.WriteString("\n\n")
		}
		if len(command.Flags) > 0 {
			logger.Infof("generating documents on flags used by subcommand '%s'", command.Name)
			buffer.WriteString("#### Flags\n\n")
			for _, flag := range command.Flags {
				buffer.WriteString(fmt.Sprintf("- `%s`\n", flag.String()))
			}
			buffer.WriteString("\n\n")
		}
	}

	if len(app.Flags) > 0 {
		logger.Infof("generating documents on global flags of app '%s'", app.Name)
		buffer.WriteString("## Global Flags\n\n")
		for _, flag := range app.Flags {
			buffer.WriteString(fmt.Sprintf("- `%s`\n", flag.String()))
		}
		buffer.WriteString("\n")
	}

	buffer.WriteString("Authors\n")
	for _, author := range app.Authors {
		buffer.WriteString(fmt.Sprintf(" - %s\n", author.String()))
	}

	buffer.WriteString("\n###### Auto generated by nikhilsbhat/urfavecli-docgen on " + time.Now().Format("2-Jan-2006") + "\n")

	return buffer.String()
}

// GenerateDocs generates markdown documentation for the commands in app.
func GenerateDocs(app *cli.App, file string) error {
	docsRootPath, err := filepath.Abs("doc")
	if err != nil {
		return err
	}

	docsPath := filepath.Join(docsRootPath, fmt.Sprintf("%s.md", file))

	logger := logrus.New()
	logger.SetLevel(goCdLogger.GetLoglevel("info"))
	logger.WithField(docGenName, true)
	logger.SetFormatter(&logrus.JSONFormatter{})

	logger.Infof("generating cli documents for '%s'", docGenName)
	logger.Infof("documets would be generated under '%s'", docsPath)

	docString := getDocs(app, logger)

	logger.Infof("documents for cli '%s' were rendered, proceeding to write the same to path '%s'", docGenName, docsPath)

	_, err = os.Stat(docsRootPath)
	if err != nil && errors.Is(err, os.ErrNotExist) {
		logger.Infof("creating directory 'doc' to place document files")
		if err = os.MkdirAll(docsRootPath, docFolderPermission); err != nil {
			return fmt.Errorf("creating document directory errored with: '%w'", err)
		}
	} else {
		logger.Infof("skipping the creation of directory 'doc' as it was already created.")
	}

	logger.Infof("proceeding to write the rendered document to '%s'.", docsPath)

	_, err = os.Stat(docsPath)
	if errors.Is(err, os.ErrNotExist) { //nolint:gocritic
		logger.Info("writing rendered document to file")
		if _, err = os.Create(docsPath); err != nil {
			return fmt.Errorf("writing document to file '%s' errored with: '%w'", docsPath, err)
		}
	} else if err != nil {
		return fmt.Errorf("stating document to file '%s' errored with: '%w'", docsPath, err)
	} else {
		logger.Infof("found an existing document with the same name. Updating it with the latest updates.")
	}

	docFile, err := os.OpenFile(docsPath, os.O_CREATE|os.O_WRONLY, os.ModeAppend) //nolint:nosnakecase
	if err != nil {
		return fmt.Errorf("reading the document file '%s' failed with: '%w'", docsPath, err)
	}

	if _, err := docFile.WriteString(docString); err != nil {
		return fmt.Errorf("updating document file '%s' with newer information errored with: '%w'", docsPath, err)
	}

	logger.Infof("documents were successfully rendered to '%s'", docsPath)

	return nil
}
