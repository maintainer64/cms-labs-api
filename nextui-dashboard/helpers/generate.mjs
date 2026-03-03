import {generateApi} from 'swagger-typescript-api';
import {readFile, writeFile} from "fs/promises";
import path from "node:path";
import services from "./generator/services.mjs";
import HooksGeneratorOpenapiRPC from "./generator/hook-use-query.mjs";

const basePath = path.resolve("helpers/api/types");


async function generate() {
    const lines = [];
    for (let service of services) {
        const input = path.resolve(service.fullPathOpenApi)
        if (service.generateTypesFile) {
            await generateApi({
                input: input,
                fileName: `${service.generateTypesFile}-types.ts`,
                generateClient: false,
                output: basePath,
            });
            lines.push(`export * from "./${service.generateTypesFile}-types";`)
        }
        const generator = new HooksGeneratorOpenapiRPC(service)
        const openapi = (await readFile(input)).toString();
        await generator.generate(openapi);
    }
    await writeFile(`${basePath}/index.ts`, lines.join("\n"))
}

generate()
