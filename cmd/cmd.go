package cmd

var (
	Version string
)

func Execute() error {
	return root_cmd.Execute()
}
