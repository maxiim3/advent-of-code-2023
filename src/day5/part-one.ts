import * as fs from "fs";


async function extractData(blocks: string[]) {

    const convertToNumbers = blocks
        .map((input) => input
            .split(" ")
            .map(data => parseInt(data))
        )

    const output = convertToNumbers.map(inputs => {
        const [destinationStart, sourceCategory, range] = inputs

        return {
            rangeInput: [sourceCategory, sourceCategory + range],
            rangeOutput: [destinationStart, destinationStart + range],

        } as const
    })

    return output
}

async function extractInformationFromFile() {
    let file = fs.readFileSync("./src/day5/input.txt", {encoding: "utf-8"})

    let lines = file.split("\n\n")
    let blocks = lines.map(chunk => chunk.split(":")[1])

    const seeds = lines[0]
        .trim()
        .split(":")[1]
        .split(" ")
        .map(n => parseInt(n))
        .filter(n => !Number.isNaN(n))

    blocks.shift()

    const splitted = blocks.map(chunk => chunk.split("\n").filter(input => input.length > 0))

    return {
        seeds,
        seed_to_soil: await extractData(splitted[0]),
        soil_to_fertilizer: await extractData(splitted[1]),
        fertilizer_to_water: await extractData(splitted[2]),
        water_to_light: await extractData(splitted[3]),
        light_to_temperature: await extractData(splitted[4]),
        temperature_to_humidity: await extractData(splitted[5]),
        humidity_to_location: await extractData(splitted[6])
    } as const

}


function mapSourceToDesitnation(collection: Awaited<ReturnType<typeof extractData>>, origin: number) {


    const retrieveListFromCollection = collection.find(({rangeInput}) => {
        const [start, end] = rangeInput
        return origin >= start && origin < end
    })

    if (!retrieveListFromCollection) return origin

    const relativeIndex = origin - retrieveListFromCollection.rangeInput[0]

    return retrieveListFromCollection.rangeOutput[0] + relativeIndex
}

async function run() {

    const {
        seeds,
        seed_to_soil,
        soil_to_fertilizer,
        fertilizer_to_water,
        water_to_light,
        light_to_temperature,
        temperature_to_humidity,
        humidity_to_location
    } = await extractInformationFromFile()

    const miniLocation = seeds.reduce((minLocation, seed) => {

        console.log(seed_to_soil)

        const soil = mapSourceToDesitnation(seed_to_soil, seed)

        const fertilizer = mapSourceToDesitnation(soil_to_fertilizer, soil)
        const water = mapSourceToDesitnation(fertilizer_to_water, fertilizer)
        const light = mapSourceToDesitnation(water_to_light, water)
        const temperature = mapSourceToDesitnation(light_to_temperature, light)
        const humidity = mapSourceToDesitnation(temperature_to_humidity, temperature)
        const location = mapSourceToDesitnation(humidity_to_location, humidity)

        console.log(`seed ${seed}, location ${location}, current Minimal : ${minLocation}`)
        return minLocation > location ? location : minLocation
    })

    console.log(miniLocation)
}


run().then(r => console.log("READY"))
