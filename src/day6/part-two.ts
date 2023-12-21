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
    const split_lines = file.split("\n");
    split_lines.pop();

    const convertEachLineToNumbers = split_lines
        .map((line) => {
            const [, rightHandSide] = line.split(":");
            const data = rightHandSide
                .split(/\s+/)
                .filter(d => d.length)
                .join("")

            return parseInt(data)
        })

    let records: Record = {
        time: convertEachLineToNumbers[0],
        distance_record: convertEachLineToNumbers[1]
    }

    return records
}


///////////////// Business Functions ///////////////////////////
function calculateNumberOfSolutionsHalf(time: number, record: number) {
    let times_that_beats_the_record: [number, number][] = []

    for (let i = 1; i <= time / 2; i++) {
        let distance = calculateBoatDistance(i, time)
        if (distance > record) {
            times_that_beats_the_record.push([i, distance])
        }
    }
    return time % 2 === 0 ? times_that_beats_the_record.length * 2 - 1 : times_that_beats_the_record.length * 2
}


function calculateBoatDistance(hold: number, time: number) {
    let delta_t = time - hold
    return hold * delta_t
}

////////////////// Main Implementation /////////////////////
function run() {
    console.log("Day 6 baby")
    const file = readFile()
    const record = extractTimeAndRecordDistanceFromFile(file);

    console.time("label")
    let solutions = calculateNumberOfSolutionsHalf(record.time, record.distance_record)
    // let solutions = calculateBestTime(record.time, record.distance_record)
    console.timeEnd("label")

    console.log(solutions)
}

run()
