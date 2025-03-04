package greetings
import( 
	"fmt"
	"errors"
	"math/rand"
)
func Hello(name string) (string,error) {
	if name == "" {
		return "", errors.New("name cannot be empty")
	}
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}
func Hellos(names []string) (map[string]string, error) {
    messages := make(map[string]string, len(names))
    for _, name := range names {
        message, err := Hello(name)
        if err != nil {
            return nil, err
        }
        messages[name] = message
    }
    return messages, nil
}

func randomFormat() string {
	formats := []string{
        "Hello, %s!",
        "Salut, %s!",
        "Hola, %s!",
        "Bonjour, %s!",
        "Ciao, %s!",
    }
    return formats[rand.Intn(len(formats))]
}