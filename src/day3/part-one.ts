// Path: day-3/index.ts

import * as fs from "fs";

type DataCell = string
type Matrix = DataCell[][]

const file = fs.readFileSync("./src/day3/input.txt", { encoding: "utf-8" });

const lines = file.split("\n")
lines.pop()

const matrix: Matrix = lines.map(line => line.split(""))

function isSymbol(pattern: string) {
    const regExp = /[.\dA-Za-z]/gi;
    const matched = regExp.exec(pattern)
    if (matched) return false
    else return true
}


function isNumber(pattern: string) {
    const matched = Number(pattern)
    if (Number.isNaN(matched)) return false
    else return true
}

class Vector {
    x: number;
    y: number;
    isNumber: ReturnType<typeof isNumber> | null
    isSymbol: ReturnType<typeof isSymbol> | null
    value: string | null

    constructor(y: number, x: number) {
        this.x = x
        this.y = y

        if (!matrix[x] || !matrix[x][y]) {
            this.isNumber = null
            this.isSymbol = null
            this.value = null
        } else {
            this.value = matrix[this.x][this.y]!
            this.isNumber = isNumber(this.value!)
            this.isSymbol = isSymbol(this.value)
        }
    }

}


type Vectors = {
    [key: string]: Vector
}

class Chunk {
    data: string
    isPartNumber: boolean

    constructor() {
        this.data = ""
        this.isPartNumber = false
    }

    append(value: string) {
        const data = this.data.split("")
        data.push(value)
        this.data = data.join("");
    }

}

class Stream {
    stream: Chunk[]

    constructor() {
        this.stream = [new Chunk()]
    }

    createChunk() {
        const chunk = new Chunk()
        this.stream.push(chunk)
        return chunk
    }

    handleLastChunk(coordinates: Vectors) {
        const hasAdjascentSymbol = Object.entries(coordinates).some(([key, value]) => key !== "current" && value.isSymbol)

        if (coordinates.current.isNumber) {
            const lastChunk = this.stream[this.stream.length - 1]
            lastChunk.append(coordinates.current.value!)
            if (hasAdjascentSymbol) {
                lastChunk.isPartNumber = true
            }
        }
    }
}

function runPartOne(matrix: Matrix) {
    const packages: Stream[] = []
    matrix.forEach((rowValue, rowIndex) => {
        const stream = new Stream()
        packages.push(stream)

        rowValue.forEach((cell, index) => {
            const vectors: Vectors = {
                current: new Vector(index, rowIndex),
                next: new Vector(index + 1, rowIndex),
                top: new Vector(index, rowIndex - 1),
                bottom: new Vector(index, rowIndex + 1),
                topLeft: new Vector(index - 1, rowIndex - 1),
                topRight: new Vector(index + 1, rowIndex - 1),
                bottomLeft: new Vector(index - 1, rowIndex + 1),
                bottomRight: new Vector(index + 1, rowIndex + 1),
                previous: new Vector(index - 1, rowIndex)
            } as const


            let hasChanged = false;
            if (
                (index === 0)
                || (vectors.previous.isNumber && !vectors.current.isNumber)
                || (!vectors.previous.isNumber && vectors.current.isNumber)
            ) {

                hasChanged = true
            }

            if (hasChanged && vectors.current.isNumber) {
                stream.createChunk()
            }
            if (vectors.current.isNumber) {
                stream.handleLastChunk(vectors)
            }
        })
    }
    )
    let partNumbers: number[] = []
    packages.forEach(stream => {
        return stream.stream.filter(chunk => {
            return chunk.isPartNumber
        }).forEach(chunk => partNumbers.push(parseInt(chunk.data)))
    })
    console.log(partNumbers)

    const result = partNumbers.reduce((acc, curr) => acc + curr)

    console.log(result)

}


runPartOne(matrix)


























































