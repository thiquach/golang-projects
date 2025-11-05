// convert strings to slices and print
package main

import "fmt"

func main() {
    const sample = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
    
    sampleSlice := []byte(sample)
    fmt.Printf("sampleSlice % x \n", sampleSlice)

    fmt.Println("Println - without formating :")
    fmt.Println(sampleSlice)

    fmt.Println("Byte loop:")
    for i := 0; i < len(sampleSlice); i++ {
        fmt.Printf("%x ", sampleSlice[i])
    }
    fmt.Printf("\n")

    fmt.Println("Printf with %x:")
    fmt.Printf("%x\n", sampleSlice)

    fmt.Println("Printf with % x:")
    fmt.Printf("% x\n", sampleSlice)

    fmt.Println("Printf with %q:")
    fmt.Printf("%q\n", sampleSlice)

    fmt.Println("*********** %q loop:")
    for i := 0; i < len(sampleSlice); i++ {
        fmt.Printf("%q ", sampleSlice[i])
    }
    fmt.Printf("\n")

    fmt.Println("Printf with %+q:")
    fmt.Printf("%+q\n", sampleSlice)
}