import * as fs from "fs";


/////////////////////////// TYPES //////////////////////
type Record = {
    time: number,
    distance_record: number
}

/////////////////////// Utils ///////////////////////////
function readFile() {
    let file = fs.readFileSync("./src/day6/input.txt", {encoding: "utf-8"})

    return file

}

/////////////////////// Mapping File to data record //////////////
function extractTimeAndRecordDistanceFromFile(file: string) {
    let columns = 0
    const split_lines = file.split("\n");
    split_lines.pop();

    const convertEachLineToNumbers = split_lines.map((line) => {
        const [_, rightHandSide] = line.split(":");
        const data = rightHandSide.split(/\s+/).filter(d => d.length)
        columns = data.length
        return data.map(d => parseInt(d))
    })

    let records: Record[] = []

    for (let i = 0; i < columns; i++) {
        records.push({
            time: convertEachLineToNumbers[0][i],
            distance_record: convertEachLineToNumbers[1][i]
        })
    }
    return records
}

////////////////// Main Implementation /////////////////////
function run() {
    const file = readFile()
    console.log("Day 6 baby")
    const records = extractTimeAndRecordDistanceFromFile(file);

    const total = records.reduce((acc, record) => {
        let solutions = calculateNumberOfSolutions(record.time, record.distance_record)

        return acc * solutions
    }, 1)

    const half = records.reduce((acc, record) => {
        let half = calculateNumberOfSolutionsHalf(record.time, record.distance_record)

        return acc * half
    }, 1)


    console.log(total, half)

}


///////////////// Business Functions ///////////////////////////
function calculateNumberOfSolutions(time: number, record: number) {
    let times_that_beats_the_record: [number, number][] = []

    for (let i = 1; i <= time; i++) {
        let distance = calculateBoatDistance(i, time)
        if (distance > record) {
            times_that_beats_the_record.push([i, distance])
        }
    }
    // console.log(times_that_beats_the_record)
    return times_that_beats_the_record.length
}

function calculateNumberOfSolutionsHalf(time: number, record: number) {
    let times_that_beats_the_record: [number, number][] = []

    for (let i = 1; i <= time / 2; i++) {
        let distance = calculateBoatDistance(i, time)
        if (distance > record) {
            times_that_beats_the_record.push([i, distance])
        }
    }
    // console.log(times_that_beats_the_record)
    return time % 2 === 0 ? times_that_beats_the_record.length * 2 - 1 : times_that_beats_the_record.length * 2
}

function calculateBoatDistance(hold: number, time: number) {
    let delta_t = time - hold
    return hold * delta_t
}

run()
