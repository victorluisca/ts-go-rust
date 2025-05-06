import getConfig, { Operation } from "../config";

test("simple print all", function () {
    const config = getConfig({});

    expect(config.operation).toEqual(Operation.Print);
    expect(config.arguments).toEqual([]);
});

test("print key", function () {
    const config = getConfig({
        arguments: ["foo"],
    });

    expect(config.operation).toEqual(Operation.Print);
    expect(config.arguments).toEqual(["foo"]);
});

test("add key", function () {
    const config = getConfig({
        arguments: ["add", "foo", "bar"],
    });

    expect(config.operation).toEqual(Operation.Add);
    expect(config.arguments).toEqual(["foo", "bar"]);
});

test("del key", function () {
    const config = getConfig({
        arguments: ["del", "foo"],
    });

    expect(config.operation).toEqual(Operation.Delete);
    expect(config.arguments).toEqual(["foo"]);
});
