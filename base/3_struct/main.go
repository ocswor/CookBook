package main

import "fmt"

type AnimalCategory struct {
	kingdom string // 界。
	phylum  string // 门。
	class   string // 纲。
	order   string // 目。
	family  string // 科。
	genus   string // 属。
	species string // 种。
}

func (ac AnimalCategory) String() string {
	return fmt.Sprintf("%s%s%s%s%s%s%s",
		ac.kingdom, ac.phylum, ac.class, ac.order,
		ac.family, ac.genus, ac.species)
}

func (ac AnimalCategory) Fake() string {
	return fmt.Sprintf("fake %s%s%s%s%s%s%s",
		ac.kingdom, ac.phylum, ac.class, ac.order,
		ac.family, ac.genus, ac.species)
}

type Animal struct {
	scientificName string // 学名。
	AnimalCategory    // 动物基本分类。
}

func (a Animal) Category() string {
	return a.AnimalCategory.String()
}


type Cat struct {
	name string
	Animal
}
func (cat Cat) String() string {
	return fmt.Sprintf("%s (category: %s, name: %q)",
		cat.scientificName, cat.Animal.AnimalCategory, cat.name)
}

func (cat *Cat) SetName(name string) {
	cat.name = name
}
func (cat Cat) GetName() string{
	return cat.name
}

func main() {
	category := AnimalCategory{species: "cat"}
	fmt.Printf("The animal category: %s\n", category)


	// 我们要做的只是把类型当做字段嵌入进来，然后坐享其成地使用嵌入字段所拥有的一切。
	cat := Animal{scientificName:"cat",AnimalCategory: category}
	fmt.Printf("The animal category: %s\n", cat.Fake())

	cat2 := Cat{}
	fmt.Printf("cat %v",cat)
	fmt.Printf("cat %+v",cat)
	fmt.Printf("cat %#v",cat)

	cat2.SetName("cat")
	fmt.Printf("cat2 %s",cat2.GetName())
}
