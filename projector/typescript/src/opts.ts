import cli from "command-line-args";

export type ProjectorOptions = {
    arguments?: string[];
    pwd?: string;
    config?: string;
};

export default function getOptions(): ProjectorOptions {
    return cli([
        {
            name: "arguments",
            type: String,
            defaultOption: true,
            multiple: true,
        },
        {
            name: "pwd",
            type: String,
            alias: "p",
        },
        {
            name: "config",
            type: String,
            alias: "c",
        },
    ]) as ProjectorOptions;
}
