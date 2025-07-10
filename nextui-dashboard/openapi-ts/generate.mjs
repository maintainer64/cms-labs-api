import {createClient} from '@hey-api/openapi-ts';

const basePath = `helpers/api`;

async function generate() {
    const directories = [
        'backend',
        'clabgate',
    ];
    for (let directory of directories) {
        await createClient({
            input: `../${directory}/docs/swagger.json`, output: {
                path: `${basePath}/${directory}`, lint: false, format: false,
            }, plugins: ['legacy/axios'],
        });
    }
}

generate()
