package cliflag

import (
	"errors"
	"fmt"
	"strings"

	cli "gopkg.in/urfave/cli.v2"
)

// RequireAll - Checks for all flags and reports errors if some are missing.
func RequireAll(c *cli.Context, required []cli.Flag) error {
	// Check for all flags
	// Append errors to slice
	var errs []string
	for _, flag := range required {
		if len(c.String(flag.Names()[0])) == 0 {
			errs = append(errs, fmt.Sprintf("--%s is missing", flag.Names()[0]))
		}
	}

	// If some errors were found join them
	// with a new line and return error
	if len(errs) != 0 {
		errs = append([]string{"Errors:"}, errs...)
		return errors.New(strings.Join(errs, "\n"))
	}
	return nil
}
