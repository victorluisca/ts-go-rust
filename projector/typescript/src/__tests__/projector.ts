import { Operation } from "../config";
import Projector from "../projector";

function createData() {
    return {
        projector: {
            "/": {
                foo: "bar1",
                bar: "baz",
            },
            "/foo": {
                foo: "bar2",
            },
            "/foo/bar": {
                foo: "bar3",
            },
        },
    };
}

function getProjector(pwd: string, data = createData()): Projector {
    return new Projector(
        {
            arguments: [],
            operation: Operation.Print,
            pwd,
            config: "",
        },
        data
    );
}

test("getValueAll", function () {
    const projector = getProjector("/foo/bar");
    expect(projector.getValueAll()).toEqual({
        bar: "baz",
        foo: "bar3",
    });
});

test("getValue", function () {
    let projector = getProjector("/foo/bar");
    expect(projector.getValue("foo")).toEqual("bar3");

    projector = getProjector("/foo");
    expect(projector.getValue("foo")).toEqual("bar2");
    expect(projector.getValue("bar")).toEqual("baz");
});

test("setValue", function () {
    let data = createData();
    let projector = getProjector("/foo/bar", data);

    projector.setValue("foo", "baz");
    expect(projector.getValue("foo")).toEqual("baz");

    projector.setValue("bar", "foo");
    expect(projector.getValue("bar")).toEqual("foo");

    projector = getProjector("/", data);
    expect(projector.getValue("bar")).toEqual("baz");
});

test("deleteValue", function () {
    const projector = getProjector("/foo/bar");

    projector.deleteValue("bar");
    expect(projector.getValue("bar")).toEqual("baz");

    projector.deleteValue("foo");
    expect(projector.getValue("foo")).toEqual("bar2");
});
