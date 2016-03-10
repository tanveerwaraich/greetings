//Package greeting to wirte greeting program
package greetings

//var GreetingWord is declared and initailized
var GreetingWord = "Welcome "

//func GreetingMessage display greeting message
func GreetingMessage(name string) string {
	return GreetingWord + name
}
