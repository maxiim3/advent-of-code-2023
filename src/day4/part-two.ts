import * as fs from "fs";


function run() {
    let file = fs.readFileSync("./src/day4/input.txt", {encoding: "utf-8"})

    let scratchcards = file.split("\n")
    scratchcards.pop()

    let points = Array.from({length: scratchcards.length}).map(() => 1)

    const winners = scratchcards.map((card) => {

        const [, readLine] = card.split(":");
        const [leftHandSide, rightHandSide] = readLine.split("|");

        const numberFound: number[] = []
        const winningNumbers = leftHandSide
            .trim()
            .split(" ")
            .map(number => parseInt(number))
            .filter(number => !Number.isNaN(number))

        rightHandSide
            .trim()
            .split(" ")
            .map(number => parseInt(number))
            .filter(number => {
                    if (Number.isNaN(number)) return

                    let include = winningNumbers.includes(number)
                    if (include) {
                        numberFound.push(number)
                    }
                }
            )


        return numberFound

    })

    let iteration = scratchcards.length

    console.log(points)
    for (let index = 0; index < iteration; index++) {
        let matchingNumbers = winners[index].length

        console.log(`Card number ${index + 1}, has ${matchingNumbers} matching numbers and points are ${points[index]}`)
        for (let n = 1; n <= matchingNumbers; n++) {
            console.log(`Adding 1 to card ${index + n + 1}`)
            points[index + n] += points[index]
            console.log(points)
            // console.log(points[index + n -1])

        }
    }
    console.log(points)
    let totalOfCards = points.reduce((acc, x) => acc + x)
    console.log(totalOfCards)

}


run()
