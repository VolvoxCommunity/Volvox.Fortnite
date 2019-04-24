package framework

/**

 * Created by cxnky on 24/04/2019 at 15:43
 * framework
 * https://github.com/cxnky/

**/

type (
	// Command defines a function with a context as the parameter
	Command func(Context)
	// CmdMap is the list of commands that have been registered along with their Command object
	CmdMap map[string]Command
	// CommandHandler is the object that is responsible for holding all of the commands and information
	CommandHandler struct {
		cmds CmdMap
	}
)

// NewCommandHandler creates a new instance of the command handler
func NewCommandHandler() *CommandHandler {
	return &CommandHandler{make(CmdMap)}
}

// GetCommands returns a list of the registered commands
func (h CommandHandler) GetCommands() CmdMap {
	return h.cmds
}

// Get is responsible for getting/attempting to get an individual command
func (h CommandHandler) Get(name string) (*Command, bool) {
	cmd, found := h.cmds[name]
	return &cmd, found
}

// RegisterCommand is responsible for registering a command in the command handler
func (h CommandHandler) RegisterCommand(name string, command Command) {
	h.cmds[name] = command
	if len(name) > 1 {
		h.cmds[name[:1]] = command
	}
}
