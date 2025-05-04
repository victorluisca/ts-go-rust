import cli from "command-line-args";

export type ProjectorOptions = {
    pwd?: string;
    config?: string;
    arguments?: string[];
};

export default function getOptions(): ProjectorOptions {
    return cli([
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
        {
            name: "arguments",
            type: String,
            defaultOption: true,
            multiple: true,
        },
    ]) as ProjectorOptions;
}
