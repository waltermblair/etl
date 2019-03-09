package main

import (
        "context"
        "github.com/aws/aws-lambda-go/lambda"
        "strings"
)

type Person struct {
        Name struct {
                First string `json:"first"`
                Last  string `json:"last"`
        } `json:"name"`
        Balance string `json:"balance"`
        Email   string `json:"email"`
}

func trimBalance(bal string) string {
        for _, c := range []string{"$", ","} {
                bal = strings.Replace(bal, c, "", -1)
        }
        return bal
}

func HandleRequest(ctx context.Context, p Person) (Person, error) {
        return Person{p.Name, trimBalance(p.Balance), p.Email}, nil
}

func main() {
        lambda.Start(HandleRequest)
}
