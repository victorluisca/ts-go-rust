import path from "path";
import { ProjectorOptions } from "./opts";

export enum Operation {
    Print,
    Add,
    Delete,
}

export type Config = {
    arguments: string[];
    operation: Operation;
    pwd: string;
    config: string;
};

function getPwd(opts: ProjectorOptions): string {
    if (opts.pwd) {
        return opts.pwd;
    }

    return process.cwd();
}

function getConfig(opts: ProjectorOptions): string {
    if (opts.config) {
        return opts.config;
    }

    const location = process.env["XDG_CONFIG_HOME"];

    if (!location) {
        throw new Error("unable to determine config location");
    }

    return path.join(location, "projector", "projector.json");
}

function getOperation(opts: ProjectorOptions): Operation {
    if (!opts.arguments || opts.arguments.length === 0) {
        return Operation.Print;
    }

    if (opts.arguments[0] == "add") {
        return Operation.Add;
    }

    if (opts.arguments[0] == "del") {
        return Operation.Delete;
    }

    return Operation.Print;
}

function getArgs(opts: ProjectorOptions) {
    if (!opts.arguments || opts.arguments.length === 0) {
        return [];
    }

    const operation = getOperation(opts);
    if (operation === Operation.Print) {
        if (opts.arguments.length > 1) {
            throw new Error(
                `expected 0 or 1 arguments but got ${opts.arguments.length}`
            );
        }
        return opts.arguments;
    }

    if (operation === Operation.Add) {
        if (opts.arguments.length !== 3) {
            throw new Error(
                `expected 2 arguments but got ${opts.arguments.length - 1}`
            );
        }
        return opts.arguments.slice(1);
    }

    if (opts.arguments.length !== 2) {
        throw new Error(
            `expected 1 argument but got ${opts.arguments.length - 1}`
        );
    }
    return opts.arguments.slice(1);
}

export default function config(opts: ProjectorOptions): Config {
    return {
        pwd: getPwd(opts),
        config: getConfig(opts),
        arguments: getArgs(opts),
        operation: getOperation(opts),
    };
}
