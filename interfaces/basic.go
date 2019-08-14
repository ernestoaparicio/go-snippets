package interfaces

import "fmt"

type HelpString string

func (hs HelpString) Help() string {
	return string(hs)
}

type UnHelpString struct {}

func (uhs *UnHelpString) Help() string{}{
	return "I cannot help you"
}

func RunBasic() {
	h := HelpString("Help me")
	fmt.Println(h.Help())

	explicit := interface{ Help() string }.Help(h)
	fmt.Println(explicit)

}
