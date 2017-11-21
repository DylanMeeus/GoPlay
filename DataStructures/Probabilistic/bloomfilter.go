package main    // testing
import(
    "fmt"
    "crypto/sha1"
    "crypto/md5"
    "encoding/binary"
    "bytes"
    "hash"
)

type bloomhash func(string) uint64

type bloomfilter struct{
    m []bool      // m as in the literature
    length uint64 // avoid figuring out the length of the array over and over, as this _must be_ static!
    data []string // actual data
    hashfunctions []bloomhash
}



// bloom filter
func createBloomFilter(size uint64) bloomfilter{
    bits := make([]bool,size)
    emptydata := make([]string,0)
    hashfunctions := []bloomhash{sha1sum, md5sum}

    return bloomfilter{m:bits, length:size, data:emptydata, hashfunctions:hashfunctions}
}

func (filter *bloomfilter) add(element string){
    for i := 0; i < len(filter.hashfunctions); i++ {
        hashfunction := filter.hashfunctions[i]
        sum := hashfunction(element)
        index := sum % filter.length
        filter.m[index] = true
    }
    filter.data = append(filter.data,element)
}

// advised to not do this in actual situations, bloom filters can have millions of entries!
func (filter *bloomfilter) printMemoryState(){
    fmt.Print((*filter).m)
}

func (filter *bloomfilter) contains(element string) bool {
    contains := true
    for i := 0; i < len(filter.hashfunctions); i++ {
        hashfunction := filter.hashfunctions[i]
        sum := hashfunction(element)
        index := sum % filter.length
        contains = filter.m[index] && contains
    }
    if !contains {
        return contains
    } else{
        // deep search the data to deal with hash collisions
        for i := 0; i < len(filter.data); i++ {
            if filter.data[i] == element {
                return true
            }
        }
        return false
    }
}

// hashing functions

func sha1sum(input string) uint64{
    shahasher := sha1.New()
    return getHashSum(&shahasher, input)
}

func md5sum(input string) uint64{
    md5hasher := md5.New()
    return getHashSum(&md5hasher, input)
}

func getHashSum(hasher *hash.Hash, input string) uint64{
    (*hasher).Write([]byte(input))
    bits := ((*hasher).Sum(nil))
    buffer := bytes.NewBuffer(bits)
    result, _ := binary.ReadUvarint(buffer)
    return result
}




// main

func main(){
    filter := createBloomFilter(50)
    filter.add("Richard")
    fmt.Println(filter.contains("Richard"))
    fmt.Println(filter.contains("Feynman"))
}