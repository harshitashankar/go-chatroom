package protocol

import(
	"bufio"
	"log"
	"io"
)

type CommandReader struct{
	reader *bufio.Reader
	//--maybe try  type CommandReader *bufio.Reader 

}

func NewCommandReader(reader io.Reader) *CommandReader {
	return &CommandReader{
		reader: bufio.NewReader(reader),
	}
}

func(r *CommandReader) Read() (interface{}, error){
	//Read the first part
	//Reads until finds the delimiter(' ')
	commandName, err := r.reader.ReadString(' ')

	if err !=nil {
		return nil, err
	}

	switch commandName{
	case "MESSAGE ":
		user, err := r.reader.ReadString(' ')
		if err != nil{
			return nil, err
		}

		message, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		return MessageCommand{
			Name: user[:len(user)-1],
			Message: message[:len(message)-1],
		}, nil

	case "SEND ":
		message, err := r.reader.ReadString('\n')
		if err!= nil {
			return nil, err
		}

		return SendCommand{
			Message: message[:len(message)-1],
		}, nil
	
	case "NAME ":
		user, err := r.reader.ReadString('\n')
		if err != nil {
			return nil, err
		}

		return NameCommand{
			Name: user[:len(user)-1],
		}, nil
	
	default:
		log.Printf("Unknown Command: %v", commandName)
	}

	return nil, UnknownCommand
}

func(r *CommandReader) ReadAll()([] interface{}, error) {
	commands := []interface{}{}
	//i think its struct of interface{} basically struct of any type

	for{
		command, err := r.Read()

		if command != nil {
			commands = append(commands, command)
		}

		if err == io.EOF {
			break
		} else if err !=nil {
			return commands, err
		}
	}

	return commands, nil
}