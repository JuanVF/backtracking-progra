package sockets

import (
	"fmt"

	"github.com/gorilla/websocket"
)

// Singleton object for SocketHandler
var handler *SocketHandler

// Void is a type to send void functions as arguments
type Void func(*websocket.Conn)

// SocketHandler will handle all the connections from
// users through websockets, also will contain all the actions
// can be performed by what the user sends
type SocketHandler struct {
	connections map[*websocket.Conn]*Sockets
	actions     map[int]Void // ID must be the Message ID
	errorAction Void         //This will be return when the action you ask for does not exists
}

// GetSocketHandlerInstance will return the instance of the SocketHandler
// this will be a Singleton object
func GetInstance() *SocketHandler {
	if handler == nil {
		handler = &SocketHandler{
			connections: make(map[*websocket.Conn]*Sockets),
			actions:     make(map[int]Void),
			errorAction: func(*websocket.Conn) {
				// By default will not do anything
				fmt.Printf("Error: function doesnt exists requested by socket...\n")
			},
		}
	}

	return handler
}

// SetErrorAction allow you to set the error action
func (s *SocketHandler) SetErrorAction(action Void) {
	s.errorAction = action
}

// AddConn will add a user to the handler
func (s *SocketHandler) AddConn(ws *websocket.Conn) error {
	if s.connections[ws] != nil {
		return fmt.Errorf("User already exists...\n")
	}

	s.connections[ws] = NewSocket(ws)

	return nil
}

// RemoveConn will remove a user from the handler
func (s *SocketHandler) RemoveConn(ws *websocket.Conn) error {
	if s.connections[ws] == nil {
		return fmt.Errorf("User does not exists...\n")
	}

	ws.Close()
	delete(s.connections, ws)

	return nil
}

// SendToAll will send a message to all the users
// if an error occurs it will delete the user from the connection
func (s *SocketHandler) SendToAll(msg Message) (sendError error) {
	for user := range s.connections {
		err := user.WriteJSON(msg)

		if err != nil {
			delete(s.connections, user)
			sendError = err
			continue
		}
	}

	return
}

// SendTo will send a message to a certain user
// if an error occurs it will delete the user from the connection
func (s *SocketHandler) SendTo(msg Message, ws *websocket.Conn) error {
	if s.connections[ws] == nil {
		return fmt.Errorf("Socket Handler: %v does not exists...\n", ws)
	}

	err := ws.WriteJSON(msg)

	if err != nil {
		delete(s.connections, ws)
	}

	return err
}

// Send will send a message to various users
func (s *SocketHandler) Send(msg Message, users []*websocket.Conn) (sendError error) {
	for _, user := range users {
		// We will send the message to all the users that exists
		// but this will throw an error anyways
		if s.connections[user] == nil {
			sendError = fmt.Errorf("Socket Handler: %v does not exists...\n", user)

			continue
		}

		err := user.WriteJSON(msg)

		if err != nil {
			delete(s.connections, user)
			sendError = err

			continue
		}
	}

	return
}

// AddAction allow you to add a function to be executed by
// the user given a certain key
func (s *SocketHandler) AddAction(key int, action Void) error {
	if s.actions[key] != nil {
		return fmt.Errorf("Socket Handler: Action already exists...\n")
	}

	s.actions[key] = action

	return nil
}

// AddActions allow you to add several actions to the handler
func (s *SocketHandler) AddActions(actions map[int]Void) error {
	for key, action := range actions {
		err := s.AddAction(key, action)

		if err != nil {
			return err
		}
	}

	return nil
}

// RemoveAction allow you to quit an action
func (s *SocketHandler) RemoveAction(key int) (action Void, err error) {
	if s.actions[key] == nil {
		return nil, fmt.Errorf("Socket Handler: Action doesnt exists...\n")
	}

	action = s.actions[key]

	delete(s.actions, key)

	return
}

// RemoveActions allow you to quit several actions
func (s *SocketHandler) RemoveActions(keys []int) (actions []Void, err error) {
	actions = make([]Void, 0)

	for key := range keys {
		action, err := s.RemoveAction(key)

		if err != nil {
			return nil, err
		}

		actions = append(actions, action)
	}

	return
}

// SetAction allow you to modify a certain action
func (s *SocketHandler) SetAction(key int, action Void) error {
	if s.actions[key] == nil {
		return fmt.Errorf("Socket Handler: Action doesnt exists...\n")
	}

	s.actions[key] = action

	return nil
}

// GetAction will return an action based on a key
// if it doesnt exists it will return the error action
func (s *SocketHandler) GetAction(key int) Void {
	action := s.actions[key]

	if action == nil {
		return s.errorAction
	}

	return action
}
