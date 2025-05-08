import path from "path";
import { Config } from "./config";
import fs from "fs";

export type ProjectorData = {
    projector: {
        // pwd
        [key: string]: {
            // key -> value
            [key: string]: string;
        };
    };
};

const defaultData = {
    projector: {},
};

export default class Projector {
    constructor(private config: Config, private data: ProjectorData) {}

    getValueAll(): { [key: string]: string } {
        let current = this.config.pwd;
        let previous = "";
        const paths = [];

        do {
            previous = current;
            paths.push(current);
            current = path.dirname(current);
        } while (current != previous);

        return paths.reverse().reduce((acc, path) => {
            const value = this.data.projector[path];
            if (value) {
                Object.assign(acc, value);
            }

            return acc;
        }, {});
    }

    getValue(key: string): string | undefined {
        let current = this.config.pwd;
        let previous = "";
        let output: string | undefined;

        do {
            const value = this.data.projector[current]?.[key];
            if (value) {
                output = value;
                break;
            }
            previous = current;
            current = path.dirname(current);
        } while (current != previous);

        return output;
    }

    setValue(key: string, value: string) {
        let dir = this.data.projector[this.config.pwd];
        if (!dir) {
            dir = this.data.projector[this.config.pwd] = {};
        }
        dir[key] = value;
    }

    deleteValue(key: string) {
        const dir = this.data.projector[this.config.pwd];
        if (dir) {
            delete dir[key];
        }
    }

    save() {
        const configPath = path.dirname(this.config.config);
        if (!fs.existsSync(configPath)) {
            fs.mkdirSync(configPath, { recursive: true });
        }
        fs.writeFileSync(this.config.config, JSON.stringify(this.data));
    }

    static fromConfig(config: Config): Projector {
        if (fs.existsSync(config.config)) {
            let data: ProjectorData;
            try {
                data = JSON.parse(fs.readFileSync(config.config).toString());
            } catch (e) {
                data = defaultData;
            }
            return new Projector(config, data);
        }
        return new Projector(config, defaultData);
    }
}
