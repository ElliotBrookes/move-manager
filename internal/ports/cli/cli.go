package ports

import (
	"ElliotBrookes/move-manager/internal/application"
	"os"
)

type CliPort struct {
}

func (p CliPort) Handle() {
	switch os.Args[1] {

	case "add_journey":
		// Call adapter

	}

}

func HandleCliInvoc(a []string) error {
	// Ensure all required args are present
	argList, err := RequiredCliArgs(a)
	if err != nil {
		return err
	}

	// Call application to calc move
	err := application.CalcMove(ra)
	if err != nil {
		return err
	}

	return nil
}
