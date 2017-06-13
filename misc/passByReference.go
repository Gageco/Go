func main() {
    i := 1
    fmt.Println("initial:", i)

    zeroval(i)
    fmt.Println("zeroval:", i)

    // The `&i` syntax gives the memory address of `i`,
    // i.e. a pointer to `i`.
    zeroptr(&i)
    fmt.Println("zeroptr:", i)

    // Pointers can be printed too.
    fmt.Println("pointer:", &i)
}
