import * as fs from "fs";


function createAlmanacSection(blocks: string[]): AlmanacSection[] {

    const convertToNumbers = blocks
        .map((input) => input
            .split(" ")
            .map(data => parseInt(data))
        )

    return convertToNumbers.map(linesOfInput => {
        const [destinationStart, sourceCategory, range] = linesOfInput

        return {
            sourceRange: [sourceCategory, sourceCategory + range],
            destinationRange: [destinationStart, destinationStart + range],

        } as const
    })
}

type AlmanacSection = {
    sourceRange: readonly [number, number],
    destinationRange: readonly [number, number]
}


async function extractInformationFromFile() {
    let file = fs.readFileSync("./src/day5/input.txt", {encoding: "utf-8"})

    let lines = file.split("\n\n")
    let chunks = lines.map(chunk => chunk.split(":")[1])

    const seeds = lines[0]
        .trim()
        .split(":")[1]
        .split(" ")
        .map(n => parseInt(n))
        .filter(n => !Number.isNaN(n))

    const rangeOfSeeds = []
    for (let i = 0; i < seeds.length; i += 2) {
        rangeOfSeeds.push({
            seedStart: seeds[i],
            seedEnd: seeds[i] + seeds[i + 1],
            range: seeds[i + 1]
        })
    }

    chunks.shift()

    const blocks = chunks.map(chunk => chunk.split("\n").filter(input => input.length > 0))

    return {
        seeds,
        rangeOfSeeds,
        blocks
    } as const

}


function calculateDestinationIndex(almanac: AlmanacSection[], origin: number) {

    const listWithOrigin = almanac.find(({sourceRange}) => {
        const [start, end] = sourceRange
        return origin >= start && origin < end
    })

    if (!listWithOrigin) return origin

    const {sourceRange, destinationRange} = listWithOrigin
    const [sourceStart,] = sourceRange
    const [destinationStart,] = destinationRange

    const inputIndex = origin - sourceStart

    return destinationStart + inputIndex
}


function pipeMappings(input: number, seed_to_soil: AlmanacSection[], soil_to_fertilizer: AlmanacSection[], fertilizer_to_water: AlmanacSection[], water_to_light: AlmanacSection[], light_to_temperature: AlmanacSection[], temperature_to_humidity: AlmanacSection[], humidity_to_location: AlmanacSection[]) {


    const soil = calculateDestinationIndex(seed_to_soil, input)
    const fertilizer = calculateDestinationIndex(soil_to_fertilizer, soil)
    const water = calculateDestinationIndex(fertilizer_to_water, fertilizer)
    const light = calculateDestinationIndex(water_to_light, water)
    const temperature = calculateDestinationIndex(light_to_temperature, light)
    const humidity = calculateDestinationIndex(temperature_to_humidity, temperature)
    return calculateDestinationIndex(humidity_to_location, humidity)

}

async function run() {
    const {
        rangeOfSeeds,
        blocks
    } = await extractInformationFromFile()

    let seed_to_soil = createAlmanacSection(blocks[0])
    let soil_to_fertilizer = createAlmanacSection(blocks[1])
    let fertilizer_to_water = createAlmanacSection(blocks[2])
    let water_to_light = createAlmanacSection(blocks[3])
    let light_to_temperature = createAlmanacSection(blocks[4])
    let temperature_to_humidity = createAlmanacSection(blocks[5])
    let humidity_to_location = createAlmanacSection(blocks[6])

    let minimalLocation: number | null = null

    console.log("FIRE 🔥")
    rangeOfSeeds.forEach(({seedStart, range, seedEnd}) => {
        console.log("🍪")
        console.time("pipe")
        for (let seed = seedStart; seed < seedEnd; seed++) {
            const out = pipeMappings(seed, seed_to_soil, soil_to_fertilizer, fertilizer_to_water, water_to_light, light_to_temperature, temperature_to_humidity, humidity_to_location)

            minimalLocation = minimalLocation === null ? out : Math.min(minimalLocation, out)
        }
        console.timeEnd("pipe")
        console.log(minimalLocation)
    })
    console.log(minimalLocation)

}

await run()

