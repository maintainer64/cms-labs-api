import templateUseQuery from "./use-query-template.mjs";
import templateUseMutation from "./use-mutation-template.mjs";
import path from "node:path";
import * as fs from "node:fs";

class HooksGeneratorOpenapiRPC {
    constructor(service) {
        this.config = {
            baseApiRpc: service.baseApiRpc,
            baseApiRpcConst: service.baseApiRpcConst,
            queriesFolder: path.resolve('helpers/queries')
        };
    }

    async generate(openapiContent) {
        const openapi = await this.parseOpenAPI(openapiContent);
        return await this.generateHookQuery(openapi);
    }

    async parseOpenAPI(content) {
        if (typeof content === 'string') {
            return JSON.parse(content);
        }
        return content;
    }

    async generateHookQuery(openapi) {
        for (const [pathKey, pathItem] of Object.entries(openapi.paths || {})) {
            const methodName = this.extractMethodName(pathKey);
            if (!methodName) continue;
            if (!pathItem.post) continue;
            const requestRPC = this.getNormalizeDefinitionModel(pathItem?.post?.parameters?.[0]?.schema);
            const responseRPC = this.getNormalizeDefinitionModel(pathItem?.post?.responses?.['200']?.schema);
            if (!requestRPC || !responseRPC) continue;
            const templateParams = {
                path: this.config.baseApiRpcConst,
                transport: "transportWithAuth",
                requestRPC: requestRPC,
                responseRPC: responseRPC,
                methodName: methodName,
                methodFolder: this.getFoldersByMethodName(methodName),
                methodJs: this.getNormalizeMethodName(methodName),
                methodFilename: this.getFilenameByMethodName(methodName),
            }
            const contentUseMutation = templateUseMutation(templateParams);
            await this.writeFile(
                path.resolve(
                    this.config.queriesFolder,
                    templateParams.methodFolder || 'base',
                    `use-mutation-${templateParams.methodFilename}.ts`
                ),
                contentUseMutation,
            );
            const contentUseQuery = templateUseQuery(templateParams);
            await this.writeFile(
                path.resolve(
                    this.config.queriesFolder,
                    templateParams.methodFolder || 'base',
                    `use-query-${templateParams.methodFilename}.ts`
                ),
                contentUseQuery,
            );
        }
        return null;
    }

    getNormalizeDefinitionModel(schema) {
        if (!schema?.['$ref']) return undefined
        return schema?.['$ref'].replace(
            /^#\/definitions\/([^.]+)\.(.+)$/,
            (match, prefix, className) => {
                const capitalizedPrefix = prefix
                    .split('_')
                    .map(word => word.charAt(0).toUpperCase() + word.slice(1))
                    .join('');

                return capitalizedPrefix + className;
            }
        );
    }

    getNormalizeMethodName(methodName) {
        return methodName.replaceAll("_", ".").replaceAll(" ", ".").split(".").map(
            word => word.charAt(0).toUpperCase() + word.slice(1)
        )
            .join('');
    }

    getFilenameByMethodName(methodName) {
        return methodName.toLowerCase().replaceAll("_", "-").replaceAll(" ", "-").replaceAll(".", "-");
    }

    getFoldersByMethodName(methodName) {
        const folderFirst = methodName.toLowerCase().split(".")
        return folderFirst?.[0] || 'base';
    }

    extractMethodName(path) {
        for (const basePath of this.config.baseApiRpc) {
            if (path.startsWith(basePath)) {
                let methodPart = path.slice(basePath.length);
                methodPart = methodPart.replace(/^\/+|\/+$/g, '');
                methodPart = methodPart.replace(/\//g, '.');

                if (methodPart.length > 0) {
                    return methodPart;
                }
            }
        }
        return null;
    }

    async writeFile(filePath, content) {
        try {
            // Получаем директорию из пути к файлу
            const dir = path.dirname(filePath);

            // Создаем все необходимые директории рекурсивно
            await fs.mkdir(dir, {recursive: true}, () => {
            });

            // Проверяем существование файла
            fs.access(filePath, fs.constants.F_OK, (err) => {
                if (err) {
                    // Файл не существует, создаем и записываем
                    fs.writeFile(filePath, content, () => {
                        console.log(`Файл ${filePath} успешно создан`);
                    });
                } else {
                    return false;
                }
            });
        } catch (error) {
            console.error(`Ошибка при работе с файлом ${filePath}:`, error.message);
            throw error;
        }
    }
}

export default HooksGeneratorOpenapiRPC;
