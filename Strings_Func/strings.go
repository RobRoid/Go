package main

import (
	"fmt"
	"strings"
)

// Creating "Person" Structure
type Person struct {
	Name string
	Age  int
}

func main() {
	prints()
	defer fmt.Println("END")

	// Start
	var s1 = Person{}
	s1.Name = "Muhammad Fauzan"
	s1.Age = 18

	fmt.Println("Name =", s1.Name)

	con1 := strings.Contains(s1.Name, "Dani")
	fmt.Printf(`Is this word contains "Dani" = %t`, con1)
	fmt.Println()

	con2 := strings.Contains(s1.Name, "zan")
	fmt.Printf(`Is this word contains "zan" = %t`, con2)
	fmt.Println()
	fmt.Println()

	pre1 := strings.HasPrefix(s1.Name, "zan")
	fmt.Printf(`Is this word contains "zan" = %t`, pre1)
	fmt.Println()

	pre2 := strings.HasPrefix(s1.Name, "Muh")
	fmt.Printf(`Is this word contains "Muh" = %t`, pre2)
	fmt.Println()
	fmt.Println()

	suf1 := strings.HasSuffix(s1.Name, "Muh")
	fmt.Printf(`Is this word contains "Muh" = %t`, suf1)
	fmt.Println()

	suf2 := strings.HasSuffix(s1.Name, "zan")
	fmt.Printf(`Is this word contains "zan" = %t`, suf2)
	fmt.Println()
	fmt.Println()

	cun1 := strings.Count(s1.Name, "u")
	fmt.Printf(`How many letters that using the letter "u" = %d`, cun1)
	fmt.Println()

	cun2 := strings.Count(s1.Name, "m")
	fmt.Printf(`How many letters that using the letter "m" = %d`, cun2)
	fmt.Println()
	fmt.Println()

	ind1 := strings.Index(s1.Name, "M")
	fmt.Printf(`Where is the Index position of "M" = %d`, ind1)
	fmt.Println()

	ind2 := strings.Index(s1.Name, "F")
	fmt.Printf(`Where is the Index position of "F" = %d`, ind2)
	fmt.Println()
	fmt.Println()

	var animalOrPerson = "A bat tale"
	fmt.Printf("Name = %s\n", animalOrPerson)

	rep1 := strings.Replace(animalOrPerson, "b", "c", 1)
	fmt.Printf(`Replacing the letter "b" to "c" = %s`, rep1)
	fmt.Println()

	rep2 := strings.Replace(animalOrPerson, "t", "d", 2)
	fmt.Printf(`Replacing the letter "t" to "d" = %s`, rep2)
	fmt.Println()
	fmt.Println()
	// See what I'm doing there? No? ok.

	repe := strings.Repeat("Fau", 2)
	fmt.Printf(`Repeating "Fau" 2 times = %s`, repe)
	fmt.Println()
	fmt.Println()

	spl1 := strings.Split(s1.Name, "")
	fmt.Printf(`Seperating "%s" from each other = %s`, s1.Name, spl1)
	fmt.Println()

	spl2 := strings.Split(animalOrPerson, " ")
	fmt.Printf(`Seperating "%s" with nothing = %q`, animalOrPerson, spl2)
	fmt.Println()
	fmt.Println()

	var name = []string{"Fauzan", "Dani", "Dimas", "Bryan"}
	jon := strings.Join(name, ", ")
	fmt.Printf(`Joining "%s" with a comma = %s`, name, jon)
	fmt.Println()
	fmt.Println()

	low := strings.ToLower(s1.Name)
	fmt.Printf(`Lowercasing "%s" = %s`, s1.Name, low)
	fmt.Println()

	upp := strings.ToUpper(name[1])
	fmt.Printf(`Uppercasing "%s" = %s`, name[1], upp)
	fmt.Println()
	fmt.Println()
	// End
}

// Just ignore this.
func prints() {
	fmt.Println()
	fmt.Println("----------------------------")
	fmt.Println(`The use of "strings" package`)
	fmt.Println("----------------------------")
	fmt.Println()
	fmt.Println("START")
	fmt.Println()
}
