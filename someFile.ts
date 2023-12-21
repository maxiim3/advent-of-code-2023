function addToArr(arr: (number | string)[]) {
    console.log("Before ", arr)
    arr.unshift("Added value");
    arr.push(0)
    console.log("After ", arr)

    return arr
}


const numbers = Array(3).fill(1)
const strings = Array(3).fill("+")
const numStrings = Array(3).fill(0).map((_, i) => i % 2 === 0 ? "+" : 1)

addToArr(numbers)
addToArr(strings)
addToArr(numStrings)

