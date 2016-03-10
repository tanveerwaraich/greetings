//Package greeting to wirte greeting program
package greetings

//var greetingword is declared and initailized
var greetingword = "Hello"

//func Greetingmessage display greeting message
func Greetingmessage(name string) string {
	return greetingword + "-" + name
}
